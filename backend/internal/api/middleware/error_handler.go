package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

// ErrorHandler はグローバルなエラーハンドラー
func ErrorHandler(err error, c echo.Context) {
	// コンテキストからロガーを取得
	log := GetLogger(c)

	code := http.StatusInternalServerError
	message := "Internal server error"

	// Echoの標準エラーの場合
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if message, ok = he.Message.(string); ok {
			message = "unknown"
		}
	}

	// エラーの重大度に応じてログレベルを変える
	if code >= 500 {
		log.Error("Request error",
			logger.ErrorAttr(err),
			"status_code", code,
			"path", c.Request().URL.Path,
			"method", c.Request().Method)
	} else if code >= 400 {
		log.Warn("Request warning",
			"error", err.Error(),
			"status_code", code,
			"path", c.Request().URL.Path,
			"method", c.Request().Method)
	}

	// エラーレスポンスを返す
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, map[string]any{
				"error": message,
			})
		}
		if err != nil {
			log.Error("Failed to send error response", logger.ErrorAttr(err))
		}
	}
}
