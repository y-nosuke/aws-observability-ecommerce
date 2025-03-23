package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// RequestIDMiddleware は各リクエストに一意のIDを割り当てるミドルウェア
func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// リクエストIDがヘッダーに含まれているか確認
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				// 含まれていない場合は新しいIDを生成
				requestID = uuid.New().String()
			}

			// コンテキストにリクエストIDを設定
			SetRequestID(c, requestID)
			// レスポンスヘッダーにもリクエストIDを設定
			c.Response().Header().Set("X-Request-ID", requestID)

			// リクエストIDを含むロガーを作成してコンテキストに設定
			log := GetLogger(c).With("request_id", requestID)
			SetLogger(c, log)

			// 次のハンドラーを呼び出す
			return next(c)
		}
	}
}
