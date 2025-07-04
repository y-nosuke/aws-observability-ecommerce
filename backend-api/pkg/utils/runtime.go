package utils

import (
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
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

	// 関数名だけを抽出
	if idx := strings.LastIndex(rel, "."); idx != -1 {
		funcName = rel[idx+1:]
	}

	// レシーバとパッケージパスを分ける
	// 例: internal/product/handler.(*ProductHandler)
	if idx := strings.LastIndex(rel, "/"); idx != -1 {
		left := rel[:idx]
		right := rel[idx+1:]
		if rIdx := strings.Index(right, ")"); rIdx != -1 {
			open := strings.Index(right, "(")
			if open != -1 {
				receiver = right[open : rIdx+1]
			}
		}
		pkgPath = left
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
