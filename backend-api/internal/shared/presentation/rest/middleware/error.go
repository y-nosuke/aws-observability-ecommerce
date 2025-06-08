package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/errors"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// ErrorHandlingMiddleware はエラーハンドリングミドルウェアを作成
func ErrorHandlingMiddleware() echo.MiddlewareFunc {
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

			return handleError(c, err)
		}
	}
}

// handleError はエラーを適切に処理してレスポンスを返す
func handleError(c echo.Context, err error) error {
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

		// エラーログ
		logger.WithError(c.Request().Context(), "アプリケーションエラーが発生", e,
			"error_type", e.Type,
			"error_code", e.Code,
			"operation", getOperationFromPath(c.Request().URL.Path),
			"resource_type", getResourceTypeFromPath(c.Request().URL.Path),
			"severity", getSeverityFromStatusCode(statusCode),
			"business_impact", getBusinessImpactFromError(e),
			"status_code", statusCode,
			"layer", "middleware")

	case *echo.HTTPError:
		// Echo HTTPエラー
		statusCode = e.Code
		if e.Internal != nil {
			// 内部エラーがある場合は詳細ログを出力
			logger.WithError(c.Request().Context(), "HTTPエラー（内部エラー有り）", e.Internal,
				"http_status", statusCode,
				"operation", getOperationFromPath(c.Request().URL.Path),
				"severity", "high",
				"business_impact", "service_degradation",
				"layer", "middleware")
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

		logger.WithError(c.Request().Context(), "予期しないエラーが発生", err,
			"operation", getOperationFromPath(c.Request().URL.Path),
			"severity", "critical",
			"business_impact", "service_disruption",
			"status_code", statusCode,
			"layer", "middleware")

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

// ===== ヘルパー関数 =====

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
