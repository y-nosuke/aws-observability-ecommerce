package logger

import (
	"errors"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
)

// ErrorAttr はエラーを構造化された属性に変換する
func ErrorAttr(err error) slog.Attr {
	return ErrorAttrWithKey("error", err)
}

// ErrorAttrWithKey は指定されたキーでエラーを構造化された属性に変換する
func ErrorAttrWithKey(key string, err error) slog.Attr {
	if err == nil {
		return slog.String(key, "")
	}

	// エラー情報をマップに変換
	errorInfo := map[string]any{
		"message": err.Error(),
	}

	// スタックトレース情報の追加（開発環境のみ）
	var programCounter [50]uintptr
	n := runtime.Callers(2, programCounter[:])
	frames := runtime.CallersFrames(programCounter[:n])

	stackTrace := make([]string, 0, n)
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			stackTrace = append(stackTrace, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
		}
		if !more {
			break
		}
		if len(stackTrace) >= 5 {
			// スタックトレースは5フレームまで
			break
		}
	}

	if len(stackTrace) > 0 {
		errorInfo["stack_trace"] = stackTrace
	}

	// エラーがラップされている場合は展開
	var unwrapped error
	if errors.Unwrap(err) != nil {
		unwrapped = errors.Unwrap(err)
		errorInfo["cause"] = unwrapped.Error()
	}

	return slog.Any(key, errorInfo)
}
