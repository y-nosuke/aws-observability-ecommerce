package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
)

// ProductDetailHandler は商品詳細APIのハンドラー
type ProductDetailHandler struct {
	reader *reader.ProductDetailReader
	mapper *mapper.ProductDetailMapper
}

// NewProductDetailHandler は新しいProductDetailHandlerを作成
func NewProductDetailHandler(db boil.ContextExecutor) *ProductDetailHandler {
	return &ProductDetailHandler{
		reader: reader.NewProductDetailReader(db),
		mapper: mapper.NewProductDetailMapper(),
	}
}

// GetProductById は指定されたIDの商品を取得する
func (h *ProductDetailHandler) GetProductById(ctx echo.Context, id openapi.ProductIdParam) error {
	// IDの整合性チェック
	if id <= 0 {
		errorResponse := h.mapper.PresentInvalidParameter("Invalid product ID")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	// 商品詳細取得
	product, err := h.reader.FindProductByID(ctx.Request().Context(), int(id))
	if err != nil {
		// 商品が見つからない場合と内部エラーを区別
		if isNotFoundError(err) {
			errorResponse := h.mapper.PresentProductNotFound("Product not found", int(id))
			return ctx.JSON(http.StatusNotFound, errorResponse)
		}

		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch product details", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// レスポンス変換
	response := h.mapper.ToProductResponse(product)

	return ctx.JSON(http.StatusOK, response)
}

// isNotFoundError はエラーが「見つからない」エラーかどうかを判定
func isNotFoundError(err error) bool {
	// 実際の実装では、具体的なエラータイプをチェック
	// 例: sql.ErrNoRows や独自のNotFoundError等
	return err.Error() == "sql: no rows in result set"
}
