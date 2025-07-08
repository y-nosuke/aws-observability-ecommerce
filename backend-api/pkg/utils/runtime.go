package utils

import (
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"unicode"
)

// FuncInfo は関数の詳細情報（フルパス、モジュール、相対パスなど）を保持
type FuncInfo struct {
	FullPath    string // フル関数名（完全修飾名）
	ModulePath  string // go.mod から取得した module 名
	RelFuncPath string // モジュールを除いた相対パス関数名（例: internal/pkg.Type.Method）
	FuncName    string // 最終的な関数名（例: Method）
	Receiver    string // レシーバ（例: (*ProductHandler)）
	PackagePath string // モジュール配下のパッケージパス（例: internal/product/handler）
}

// cachedModulePath は起動時に一度だけ取得して使い回す
var (
	cachedModulePath string
	once             sync.Once
)

// GetModulePath は go.mod に書かれた module 名を取得（1度だけ実行）
func GetModulePath() string {
	once.Do(func() {
		if info, ok := debug.ReadBuildInfo(); ok && info != nil {
			cachedModulePath = info.Main.Path
		}
	})
	return cachedModulePath
}

// ParseFuncInfo は runtime の pc から関数名を解析し、構造体にして返す
func ParseFuncInfo(skip int) (*FuncInfo, string, int) {
	pc, file, line, ok := runtime.Caller(skip + 1) // +1 で自分を飛ばす
	if !ok {
		return &FuncInfo{FullPath: "unknown"}, "unknown", 0
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return &FuncInfo{FullPath: "unknown"}, "unknown", 0
	}

	full := fn.Name()
	module := GetModulePath()
	rel := full
	if after, ok := strings.CutPrefix(full, module+"/"); ok {
		rel = after
	} else if full == module {
		rel = "" // モジュール関数直下の場合
	}

	// 例: internal/product/handler.(*ProductHandler).UploadProductImage
	funcName := ""
	receiver := ""
	pkgPath := ""

	lastDot := strings.LastIndex(rel, ".")
	if lastDot == -1 {
		// ドットがない場合、これはパッケージ修飾子のない純粋な関数名であると想定
		// 例: main パッケージ内の main 関数など
		funcName = rel
	} else {
		// 関数名とそれ以外（パッケージパス + レシーバ）を分離
		funcName = rel[lastDot+1:]
		qualifier := rel[:lastDot]
		// ポインタレシーバか、値レシーバ/関数かでロジックを分岐
		// 例: internal/product/handler.(*ProductHandler)
		if strings.HasSuffix(qualifier, ")") {
			// ポインタレシーバ: ".(...)" の形式を探す
			if openParen := strings.LastIndex(qualifier, ".("); openParen != -1 {
				pkgPath = qualifier[:openParen]
				receiver = qualifier[openParen+1:]
			} else {
				// 予期しない形式だが、フォールバックとして qualifier をパッケージパスとする
				pkgPath = qualifier
			}
		} else {
			// 値レシーバ または レシーバなしの関数
			// 修飾子内の最後のドットでパッケージと型を分離
			// 例(値レシーバ): app/reception/infra/repo.UserRepositoryImpl -> repo, UserRepositoryImpl
			// 例(関数): app/framework/session -> app/framework/session, (レシーバなし)
			if receiverDot := strings.LastIndex(qualifier, "."); receiverDot != -1 {
				// パッケージパスとレシーバ名の候補に分割
				pathCandidate := qualifier[:receiverDot]
				receiverCandidate := qualifier[receiverDot+1:]
				// パッケージパス部分にスラッシュが含まれるか、
				// またはレシーバ名の候補の最初の文字が大文字である場合、
				// これをレシーバ付きメソッドと判断する
				if strings.Contains(pathCandidate, "/") || (len(receiverCandidate) > 0 && unicode.IsUpper(rune(receiverCandidate[0]))) {
					pkgPath = pathCandidate
					receiver = receiverCandidate
				} else {
					// そうでなければレシーバなしの関数と判断
					pkgPath = qualifier
				}
			} else {
				// ドットがなければレシーバなしの関数
				pkgPath = qualifier
			}
		}
	}

	return &FuncInfo{
		FullPath:    full,
		ModulePath:  module,
		RelFuncPath: rel,
		FuncName:    funcName,
		Receiver:    receiver,
		PackagePath: pkgPath,
	}, file, line
}
