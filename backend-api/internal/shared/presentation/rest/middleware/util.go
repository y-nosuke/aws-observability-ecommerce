package middleware

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// getRoutePattern はEchoのルートパターンを取得します
func getRoutePattern(c echo.Context) string {
	// Echoのルートパターンを取得
	if route := c.Path(); route != "" {
		return route
	}

	// フォールバック: リクエストパスから動的パラメータを推測
	path := c.Request().URL.Path
	return normalizeRoute(path)
}

// normalizeRoute はパスを正規化してルートパターンを生成します
func normalizeRoute(path string) string {
	// 基本的なパス正規化
	if path == "" {
		return "/"
	}

	// 数値IDパラメータを{id}に置換
	parts := strings.Split(path, "/")
	for i, part := range parts {
		// 数値のみの部分を{id}に置換
		if isNumeric(part) && len(part) > 0 {
			parts[i] = "{id}"
		}
	}

	return strings.Join(parts, "/")
}

// isNumeric は文字列が数値かどうかを判定します
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	_, err := strconv.Atoi(s)
	return err == nil
}
