package middleware

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
)

// PaginationInfo はページネーション情報を保持する構造体
type PaginationInfo struct {
	Page     *int
	PageSize *int
	Limit    *int
	Offset   *int
}

// SearchInfo は検索情報を保持する構造体
type SearchInfo struct {
	Keyword       *string
	SearchEnabled bool
	SortBy        *string
	Order         *string
}

// BusinessContextMiddleware はビジネスコンテキスト情報を自動抽出するミドルウェア
func BusinessContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 現在のスパンを取得
			span := trace.SpanFromContext(c.Request().Context())
			if !span.IsRecording() {
				return next(c)
			}

			ctx := c.Request().Context()

			// ルートパターンの取得
			route := getRoutePattern(c)

			// Business domain の判定
			domain := determineDomain(route)
			span.SetAttributes(attribute.String("app.business_domain", domain))
			observability.SetDomainToContext(ctx, domain)

			if id := extractID(c, domain); id > 0 {
				span.SetAttributes(
					attribute.Int(fmt.Sprintf("app.%s_id", domain), id),
					attribute.String("app.entity_type", domain),
					attribute.Int("app.entity_id", id),
				)

				observability.SetEntityIDToContext(ctx, id)
				observability.SetEntityTypeToContext(ctx, domain)
			}

			c.SetRequest(c.Request().WithContext(ctx))

			// Operation type の判定
			operationType := determineOperationType(c.Request().Method, route)
			span.SetAttributes(attribute.String("app.operation_type", operationType))

			// ページネーション情報の抽出
			paginationInfo := extractPaginationInfo(c)
			if paginationInfo.Page != nil {
				span.SetAttributes(attribute.Int("app.request.page", *paginationInfo.Page))
			}
			if paginationInfo.PageSize != nil {
				span.SetAttributes(attribute.Int("app.request.page_size", *paginationInfo.PageSize))
			}
			if paginationInfo.Limit != nil {
				span.SetAttributes(attribute.Int("app.request.limit", *paginationInfo.Limit))
			}
			if paginationInfo.Offset != nil {
				span.SetAttributes(attribute.Int("app.request.offset", *paginationInfo.Offset))
			}

			// 検索クエリ情報の抽出
			searchInfo := extractSearchInfo(c)
			if searchInfo.Keyword != nil {
				span.SetAttributes(
					attribute.String("app.search.keyword", *searchInfo.Keyword),
					attribute.Bool("app.search.enabled", true),
				)
			}
			if searchInfo.SortBy != nil {
				span.SetAttributes(attribute.String("app.request.sort_by", *searchInfo.SortBy))
			}
			if searchInfo.Order != nil {
				span.SetAttributes(attribute.String("app.request.order", *searchInfo.Order))
			}

			return next(c)
		}
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

// extractID はルートパラメータまたはクエリパラメータからIDを抽出
func extractID(c echo.Context, domain string) int {
	// OpenAPIのルートパラメータから取得
	if idParam := c.Param("id"); idParam != "" {
		if id, err := strconv.Atoi(idParam); err == nil {
			return id
		}
	}

	// idパラメータから取得（クエリパラメータ）
	if idStr := c.QueryParam(fmt.Sprintf("%s_id", domain)); idStr != "" {
		if id, err := strconv.Atoi(idStr); err == nil {
			return id
		}
	}

	return 0
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

// extractPaginationInfo はページネーション情報を抽出
func extractPaginationInfo(c echo.Context) PaginationInfo {
	info := PaginationInfo{}

	if pageStr := c.QueryParam("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			info.Page = &page
		}
	}

	if pageSizeStr := c.QueryParam("page_size"); pageSizeStr != "" {
		if pageSize, err := strconv.Atoi(pageSizeStr); err == nil {
			info.PageSize = &pageSize
		}
	}

	if limitStr := c.QueryParam("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			info.Limit = &limit
		}
	}

	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil {
			info.Offset = &offset
		}
	}

	return info
}

// extractSearchInfo は検索情報を抽出
func extractSearchInfo(c echo.Context) SearchInfo {
	info := SearchInfo{}

	if keyword := c.QueryParam("keyword"); keyword != "" {
		info.Keyword = &keyword
		info.SearchEnabled = true
	}

	if sortBy := c.QueryParam("sort_by"); sortBy != "" {
		info.SortBy = &sortBy
	}

	if order := c.QueryParam("order"); order != "" {
		info.Order = &order
	}

	return info
}
