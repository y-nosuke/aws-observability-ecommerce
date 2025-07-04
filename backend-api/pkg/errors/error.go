package errors

import (
	"errors"
	"fmt"
	"runtime"
)

const (
	// maxStackDepth はスタックトレースの最大深度
	maxStackDepth = 32
)

// Error はアプリケーション独自のエラー型
type Error struct {
	message string    // ユーザー向けメッセージ
	Cause   error     // 元のエラー
	stack   []uintptr // スタックトレース（PC値のみ、遅延評価）
}

// Frame はスタックフレーム情報
type Frame struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

type StackTrace []Frame

// Error はerrorインターフェースの実装
func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s (cause: %v)", e.message, e.Cause)
	}
	return e.message
}

// Unwrap はerrorsパッケージのUnwrap対応
func (e *Error) Unwrap() error {
	return e.Cause
}

// HasStack はスタックトレースが利用可能かチェック
func (e *Error) HasStack() bool {
	return len(e.stack) > 0
}

// StackTrace はスタックトレースを遅延評価で取得
func (e *Error) StackTrace() StackTrace {
	if !e.HasStack() {
		return nil
	}

	// 既にAppErrorでスタック取得済みの場合は、そちらを優先
	if e.Cause != nil {
		var appErr *Error
		if errors.As(e.Cause, &appErr) && appErr.HasStack() {
			return appErr.StackTrace()
		}
	}

	// runtime.CallersFrames()で効率的に変換
	frames := runtime.CallersFrames(e.stack)
	var result StackTrace

	for {
		frame, more := frames.Next()
		result = append(result, Frame{
			Function: frame.Function,
			File:     frame.File,
			Line:     frame.Line,
		})

		if !more {
			break
		}
	}

	return result
}

// GetStackTrace はスタックトレースを文字列で返す
func (e *Error) GetStackTrace() string {
	frames := e.StackTrace()
	if len(frames) == 0 {
		return ""
	}

	var result string
	for _, frame := range frames {
		result += fmt.Sprintf("  %s\n    %s:%d\n", frame.Function, frame.File, frame.Line)
	}
	return result
}

// GetMessage はメッセージを取得
func (e *Error) GetMessage() string {
	return e.message
}

// New は新しいAppErrorを作成
func New(message string) *Error {
	return &Error{
		message: message,
		stack:   captureStackTrace(2), // 1つ追加でスキップ（Newをスキップ）
	}
}

// Newf はフォーマット付きで新しいAppErrorを作成
func Newf(format string, args ...any) *Error {
	return &Error{
		message: fmt.Sprintf(format, args...),
		stack:   captureStackTrace(2), // 1つ追加でスキップ（Newfをスキップ）
	}
}

// Wrap は既存のエラーをラップ
func Wrap(err error, message string) *Error {
	if err == nil {
		return nil
	}

	appErr := &Error{
		message: message,
		Cause:   err,
	}

	// 既にAppErrorの場合はスタックトレースを保持（重複取得を避ける）
	var existingAppErr *Error
	if errors.As(err, &existingAppErr) && existingAppErr.HasStack() {
		// 既存のスタックトレースを保持
		appErr.stack = existingAppErr.stack
	} else {
		// 新規にスタックトレースを取得（Wrapをスキップ）
		appErr.stack = captureStackTrace(2)
	}

	return appErr
}

// Wrapf はフォーマット付きで既存のエラーをラップ
func Wrapf(err error, format string, args ...any) *Error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

// captureStackTrace はスタックトレースを効率的にキャプチャ
// additionalSkip: 呼び出し元が追加でスキップしたいフレーム数
func captureStackTrace(additionalSkip int) []uintptr {
	stack := make([]uintptr, maxStackDepth)
	// skip = 1 (captureStackTrace) + additionalSkip (New/Wrapなど)
	skip := 1 + additionalSkip
	n := runtime.Callers(skip, stack)

	// 実際のサイズにトリムしてメモリ効率化
	return stack[:n]
}

// IsAppError は指定されたエラーがAppErrorかチェック
func IsAppError(err error) bool {
	var appError *Error
	ok := errors.As(err, &appError)
	return ok
}

// GetAppError は指定されたエラーからAppErrorを取得
func GetAppError(err error) (*Error, bool) {
	var appErr *Error
	ok := errors.As(err, &appErr)
	return appErr, ok
}
