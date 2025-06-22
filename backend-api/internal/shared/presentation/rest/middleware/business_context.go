package middleware

import (
	"context"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// BusinessContextMiddleware はビジネスコンテキスト情報を自動抽出するミドルウェア
func BusinessContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 現在のスパンを取得
			span := trace.SpanFromContext(c.Request().Context())
			if !span.IsRecording() {
				return next(c)
			}

			// ルートパターンの取得
			route := getRoutePattern(c)

			// ビジネスドメイン情報を抽出・記録
			extractBusinessContext(c, span, route)

			return next(c)
		}
	}
}

// extractBusinessContext はルートからビジネスコンテキストを抽出
func extractBusinessContext(c echo.Context, span trace.Span, route string) {
	ctx := c.Request().Context()

	// Product関連のパラメータ抽出
	if productID := extractProductID(c, route); productID > 0 {
		span.SetAttributes(
			attribute.Int64("app.product_id", productID),
			attribute.String("app.entity_type", "product"),
		)

		ctx = context.WithValue(ctx, observability.EntityIDKey, productID)
		ctx = context.WithValue(ctx, observability.EntityTypeKey, "product")
		ctx = context.WithValue(ctx, observability.BusinessDomainKey, "product")

		c.SetRequest(c.Request().WithContext(ctx))
	}

	// Category関連のパラメータ抽出
	if categoryID := extractCategoryID(c, route); categoryID > 0 {
		span.SetAttributes(
			attribute.Int64("app.category_id", categoryID),
			attribute.String("app.entity_type", "category"),
		)

		ctx = context.WithValue(ctx, observability.EntityIDKey, categoryID)
		ctx = context.WithValue(ctx, observability.EntityTypeKey, "category")
		ctx = context.WithValue(ctx, observability.BusinessDomainKey, "category")

		c.SetRequest(c.Request().WithContext(ctx))
	}

	// ページネーション情報の抽出
	extractPaginationInfo(c, span)

	// 検索クエリ情報の抽出
	extractSearchInfo(c, span)

	// Operation type の判定
	operationType := determineOperationType(c.Request().Method, route)
	span.SetAttributes(attribute.String("app.operation_type", operationType))

	// Business domain の判定
	domain := determineDomain(route)
	span.SetAttributes(attribute.String("app.business_domain", domain))
}

// extractProductID はルートからproduct_idを抽出
func extractProductID(c echo.Context, route string) int64 {
	// OpenAPIのルートパラメータから取得
	if idParam := c.Param("id"); idParam != "" {
		// /products/{id} のパターン
		if matched, err := regexp.MatchString(`/products/\{id\}`, route); err == nil && matched {
			if id, err := strconv.ParseInt(idParam, 10, 64); err == nil {
				return id
			}
		}
	}

	// product_idパラメータから取得
	if idParam := c.Param("product_id"); idParam != "" {
		if id, err := strconv.ParseInt(idParam, 10, 64); err == nil {
			return id
		}
	}

	return 0
}

// extractCategoryID はルートからcategory_idを抽出
func extractCategoryID(c echo.Context, route string) int64 {
	// OpenAPIのルートパラメータから取得
	if idParam := c.Param("id"); idParam != "" {
		// /categories/{id} のパターン
		if matched, err := regexp.MatchString(`/categories/\{id\}`, route); err == nil && matched {
			if id, err := strconv.ParseInt(idParam, 10, 64); err == nil {
				return id
			}
		}
	}

	// category_idパラメータから取得（クエリパラメータ）
	if categoryIDStr := c.QueryParam("category_id"); categoryIDStr != "" {
		if id, err := strconv.ParseInt(categoryIDStr, 10, 64); err == nil {
			return id
		}
	}

	return 0
}

// extractPaginationInfo はページネーション情報を抽出
func extractPaginationInfo(c echo.Context, span trace.Span) {
	if pageStr := c.QueryParam("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			span.SetAttributes(attribute.Int("app.request.page", page))
		}
	}

	if pageSizeStr := c.QueryParam("page_size"); pageSizeStr != "" {
		if pageSize, err := strconv.Atoi(pageSizeStr); err == nil {
			span.SetAttributes(attribute.Int("app.request.page_size", pageSize))
		}
	}

	if limitStr := c.QueryParam("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			span.SetAttributes(attribute.Int("app.request.limit", limit))
		}
	}

	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil {
			span.SetAttributes(attribute.Int("app.request.offset", offset))
		}
	}
}

// extractSearchInfo は検索情報を抽出
func extractSearchInfo(c echo.Context, span trace.Span) {
	if keyword := c.QueryParam("keyword"); keyword != "" {
		span.SetAttributes(
			attribute.String("app.search.keyword", keyword),
			attribute.Bool("app.search.enabled", true),
		)
	}

	if sortBy := c.QueryParam("sort_by"); sortBy != "" {
		span.SetAttributes(attribute.String("app.request.sort_by", sortBy))
	}

	if order := c.QueryParam("order"); order != "" {
		span.SetAttributes(attribute.String("app.request.order", order))
	}
}

// determineOperationType はHTTPメソッドとルートから操作タイプを判定
func determineOperationType(method, route string) string {
	switch method {
	case "GET":
		if matched, err := regexp.MatchString(`/\{id\}$`, route); err == nil && matched {
			return "read_single"
		}
		return "read_list"
	case "POST":
		if matched, err := regexp.MatchString(`/\{id\}/image$`, route); err == nil && matched {
			return "upload_image"
		}
		return "create"
	case "PUT", "PATCH":
		return "update"
	case "DELETE":
		return "delete"
	default:
		return "unknown"
	}
}

// determineDomain はルートからビジネスドメインを判定
func determineDomain(route string) string {
	switch {
	case regexp.MustCompile(`/products`).MatchString(route):
		return "product"
	case regexp.MustCompile(`/categories`).MatchString(route):
		return "category"
	case regexp.MustCompile(`/inventory`).MatchString(route):
		return "inventory"
	case regexp.MustCompile(`/health`).MatchString(route):
		return "system"
	default:
		return "unknown"
	}
}
