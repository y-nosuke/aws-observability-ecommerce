package handler

import (
	"net/http"
	"strconv"

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

	// 商品詳細を取得
	product, err := h.reader.FindProductByID(ctx.Request().Context(), int(id))
	if err != nil {
		// エラーの種類に応じて適切なレスポンスを返す
		if err.Error() == "product not found: "+strconv.Itoa(int(id)) {
			errorResponse := h.mapper.PresentProductNotFound("Product not found", int(id))
			return ctx.JSON(http.StatusNotFound, errorResponse)
		}

		// その他のエラー
		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch product details", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// レスポンスの構築
	response := h.mapper.ToProductResponse(product)
	return ctx.JSON(http.StatusOK, response)
}
