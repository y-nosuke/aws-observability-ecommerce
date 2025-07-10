package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
)

// ProductDetailHandler は商品詳細APIのハンドラー
type ProductDetailHandler struct {
	reader *reader.ProductDetailReader
	mapper *mapper.ProductDetailMapper
}

// NewProductDetailHandler は新しいProductDetailHandlerを作成
func NewProductDetailHandler() *ProductDetailHandler {
	return &ProductDetailHandler{
		reader: reader.NewProductDetailReader(),
		mapper: mapper.NewProductDetailMapper(),
	}
}

// GetProductById は指定されたIDの商品を取得する
func (h *ProductDetailHandler) GetProductById(ctx echo.Context, id openapi.ProductIdParam) error {
	return otel.WithSpan(ctx.Request().Context(), func(spanCtx context.Context, o *otel.Observer) error {
		// IDの整合性チェック
		if id <= 0 {
			return fmt.Errorf("id must be a positive integer")
		}

		// 商品詳細取得
		product, err := h.reader.FindProductByID(spanCtx, id)
		if err != nil {
			// 商品が見つからない場合と内部エラーを区別
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("product not found: %w", err)
			}

			return fmt.Errorf("failed to fetch product by id: %w", err)
		}

		// レスポンス変換
		response := h.mapper.ToProductResponse(product)

		return ctx.JSON(http.StatusOK, response)
	})
}
