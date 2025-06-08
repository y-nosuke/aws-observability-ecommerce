package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/errors"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
)

// ErrorHandlingMiddleware はエラーハンドリングミドルウェアを作成
func ErrorHandlingMiddleware(logger logging.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			// レスポンスが既に送信されていない場合のみエラー処理
			if c.Response().Committed {
				return err
			}

			return handleError(c, err, logger)
		}
	}
}

// handleError はエラーを適切に処理してレスポンスを返す
func handleError(c echo.Context, err error, logger logging.Logger) error {
	var statusCode int
	var errorResponse map[string]interface{}

	switch e := err.(type) {
	case *errors.AppError:
		// アプリケーション定義エラー
		statusCode = getStatusCodeFromErrorType(e.Type)
		errorResponse = map[string]interface{}{
			"error": map[string]interface{}{
				"type":    e.Type,
				"message": e.Message,
				"code":    e.Code,
			},
		}

		// 構造化ログでエラーを記録
		errorCtx := logging.ErrorContext{
			Operation:      getOperationFromPath(c.Request().URL.Path),
			ResourceType:   getResourceTypeFromPath(c.Request().URL.Path),
			Severity:       getSeverityFromStatusCode(statusCode),
			BusinessImpact: getBusinessImpactFromError(e),
		}

		// コンテキスト情報があれば追加
		if resourceID := c.Param("id"); resourceID != "" {
			errorCtx.ResourceID = resourceID
		}

		logger.LogError(c.Request().Context(), e, errorCtx)

	case *echo.HTTPError:
		// Echo HTTPエラー
		statusCode = e.Code
		if e.Internal != nil {
			// 内部エラーがある場合は詳細ログを出力
			internalErr := errors.NewAppError("InternalError", e.Error(), "INTERNAL_ERROR").
				WithUnderlying(e.Internal)

			errorCtx := logging.ErrorContext{
				Operation:      getOperationFromPath(c.Request().URL.Path),
				Severity:       "high",
				BusinessImpact: "service_degradation",
			}
			logger.LogError(c.Request().Context(), internalErr, errorCtx)
		}

		errorResponse = map[string]interface{}{
			"error": map[string]interface{}{
				"type":    "HTTPError",
				"message": e.Message,
				"code":    fmt.Sprintf("HTTP_%d", statusCode),
			},
		}

	default:
		// その他の予期しないエラー
		statusCode = http.StatusInternalServerError

		internalErr := errors.NewAppError("UnexpectedError", "An unexpected error occurred", "UNEXPECTED_ERROR").
			WithUnderlying(err)

		errorCtx := logging.ErrorContext{
			Operation:      getOperationFromPath(c.Request().URL.Path),
			Severity:       "critical",
			BusinessImpact: "service_disruption",
		}
		logger.LogError(c.Request().Context(), internalErr, errorCtx)

		errorResponse = map[string]interface{}{
			"error": map[string]interface{}{
				"type":    "InternalServerError",
				"message": "An unexpected error occurred",
				"code":    "INTERNAL_SERVER_ERROR",
			},
		}
	}

	// JSONレスポンスを返す
	return c.JSON(statusCode, errorResponse)
}

// getStatusCodeFromErrorType はエラータイプからHTTPステータスコードを取得
func getStatusCodeFromErrorType(errorType string) int {
	switch errorType {
	case "ValidationError":
		return http.StatusBadRequest
	case "NotFoundError":
		return http.StatusNotFound
	case "UnauthorizedError":
		return http.StatusUnauthorized
	case "ForbiddenError":
		return http.StatusForbidden
	case "ConflictError":
		return http.StatusConflict
	case "DatabaseConnectionError":
		return http.StatusServiceUnavailable
	case "ExternalServiceError":
		return http.StatusBadGateway
	case "TimeoutError":
		return http.StatusGatewayTimeout
	default:
		return http.StatusInternalServerError
	}
}

// getOperationFromPath はパスから操作名を推定
func getOperationFromPath(path string) string {
	if path == "" {
		return "unknown"
	}

	// パスから操作を推定（簡単な実装）
	switch {
	case path == "/api/health":
		return "health_check"
	case path == "/api/products" || path == "/api/products/":
		return "product_list"
	case path == "/api/categories" || path == "/api/categories/":
		return "category_list"
	default:
		if len(path) > 20 {
			return path[:20] + "..."
		}
		return path
	}
}

// getResourceTypeFromPath はパスからリソースタイプを推定
func getResourceTypeFromPath(path string) string {
	switch {
	case path == "/api/health":
		return "health"
	case path == "/api/products" || path == "/api/products/":
		return "product"
	case path == "/api/categories" || path == "/api/categories/":
		return "category"
	default:
		return "unknown"
	}
}

// getSeverityFromStatusCode はステータスコードから影響度を取得
func getSeverityFromStatusCode(statusCode int) string {
	switch {
	case statusCode >= 500:
		return "high"
	case statusCode >= 400:
		return "medium"
	default:
		return "low"
	}
}

// getBusinessImpactFromError はエラーからビジネス影響を取得
func getBusinessImpactFromError(err *errors.AppError) string {
	switch err.Type {
	case "DatabaseConnectionError":
		return "service_disruption"
	case "ValidationError":
		return "user_experience_degradation"
	case "NotFoundError":
		return "content_unavailable"
	case "UnauthorizedError", "ForbiddenError":
		return "access_control_violation"
	case "ExternalServiceError":
		return "feature_unavailable"
	default:
		return "unknown"
	}
}
