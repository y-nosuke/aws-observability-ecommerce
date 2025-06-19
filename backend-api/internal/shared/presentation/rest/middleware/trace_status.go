package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// TraceStatusMiddleware はHTTPレスポンスステータスに基づいてトレースステータスを設定するミドルウェア
func TraceStatusMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ハンドラーを実行
			err := next(c)

			// 現在のスパンを取得
			span := trace.SpanFromContext(c.Request().Context())
			if !span.IsRecording() {
				return err
			}

			// レスポンスステータスコードを取得
			statusCode := c.Response().Status

			// ステータスコードに基づいてスパンステータスを設定
			if err != nil {
				// エラーが発生した場合
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			} else if statusCode >= 500 {
				span.SetStatus(codes.Error, http.StatusText(statusCode))
			} else if statusCode >= 400 {
				// 4xxはクライアントエラーなので通常はOKとして扱う
				// ただし、アプリケーションの要件に応じて調整可能
				span.SetStatus(codes.Ok, "")
			} else {
				// 2xx, 3xxの場合は成功
				span.SetStatus(codes.Ok, "")
			}

			return err
		}
	}
}
