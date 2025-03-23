package middleware

import (
	"github.com/labstack/echo/v4"
)

// AuthLoggerMiddleware は認証情報をロガーに追加するミドルウェア
func AuthLoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 現在のロガーを取得
			log := GetLogger(c)

			// 認証されたユーザーがいれば情報を追加
			// 注意: 実際の認証の実装に合わせて調整する必要があります
			if userID := getUserIDFromContext(c); userID != "" {
				log = log.With("user_id", userID)
				SetLogger(c, log)
			}

			// 管理者ユーザーの場合は追加情報
			if isAdmin := isAdminUser(c); isAdmin {
				log = log.With("user_role", "admin")
				SetLogger(c, log)
			}

			return next(c)
		}
	}
}

// ダミー実装 - 実際の認証システムに合わせて実装する必要があります
func getUserIDFromContext(c echo.Context) string {
	if userID, ok := c.Get("user_id").(string); ok {
		return userID
	}
	return ""
}

// ダミー実装 - 実際の認証システムに合わせて実装する必要があります
func isAdminUser(c echo.Context) bool {
	if role, ok := c.Get("user_role").(string); ok {
		return role == "admin"
	}
	return false
}
