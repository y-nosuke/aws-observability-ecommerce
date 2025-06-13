package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// CategoryListHandler はカテゴリー一覧APIのハンドラー
type CategoryListHandler struct {
	reader *reader.CategoryListReader
	mapper *mapper.CategoryListMapper
}

// NewCategoryListHandler は新しいCategoryListHandlerを作成
func NewCategoryListHandler(db boil.ContextExecutor) *CategoryListHandler {
	return &CategoryListHandler{
		reader: reader.NewCategoryListReader(db),
		mapper: mapper.NewCategoryListMapper(),
	}
}

// ListCategories はカテゴリー一覧を取得する
func (h *CategoryListHandler) ListCategories(ctx echo.Context) error {
	// トレーシングスパンを開始
	tracer := otel.Tracer("aws-observability-ecommerce")
	requestCtx, span := tracer.Start(ctx.Request().Context(), "handler.list_categories", trace.WithAttributes(
		attribute.String("app.layer", "handler"),
		attribute.String("app.domain", "category"),
		attribute.String("app.operation", "list_categories"),
		attribute.String("http.method", ctx.Request().Method),
		attribute.String("http.route", ctx.Path()),
	))
	defer span.End()

	// ログ記録用の操作開始
	completeOp := logger.StartOperation(requestCtx, "list_categories",
		"operation_type", "category_list",
		"layer", "handler")

	// 子スパンでカテゴリー一覧取得
	dataCtx, dataSpan := tracer.Start(requestCtx, "handler.fetch_categories_data")
	categories, err := h.reader.FindCategoriesWithProductCount(dataCtx)
	if err != nil {
		// データ取得エラー処理
		dataSpan.RecordError(err)
		dataSpan.SetStatus(codes.Error, err.Error())
		dataSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusInternalServerError))

		// エラーログ
		logger.WithError(requestCtx, "カテゴリー一覧取得に失敗", err,
			"operation", "fetch_categories",
			"layer", "handler")

		completeOp(false, "error_type", "data_fetch_error")
		
		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch categories", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	dataSpan.SetAttributes(attribute.Int("app.categories_fetched", len(categories)))
	dataSpan.End()

	// 子スパンでレスポンス構築
	_, mapSpan := tracer.Start(requestCtx, "handler.map_categories_response")
	response := h.mapper.ToCategoryListResponse(categories)
	mapSpan.SetAttributes(
		attribute.Int("app.response_items", len(response.Items)), // ✅ Data → Items に修正
		attribute.Int("app.total_categories", len(response.Items)), // ✅ Data → Items に修正
	)
	mapSpan.End()

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("http.response.status_code", http.StatusOK),
		attribute.Int("app.categories_returned", len(response.Items)), // ✅ Data → Items に修正
	)

	// 成功ログ
	logger.Info(requestCtx, "カテゴリー一覧取得が完了",
		"categories_count", len(response.Items), // ✅ Data → Items に修正
		"layer", "handler")

	// 操作完了記録
	completeOp(true,
		"categories_count", len(response.Items), // ✅ Data → Items に修正
		"http_status", http.StatusOK)

	// ビジネスイベントとして記録
	logger.LogBusinessEvent(requestCtx, "categories_listed", "category", "system",
		"categories_count", len(response.Items), // ✅ Data → Items に修正
		"success", true)

	return ctx.JSON(http.StatusOK, response)
}
