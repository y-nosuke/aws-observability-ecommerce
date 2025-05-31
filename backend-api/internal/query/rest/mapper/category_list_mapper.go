package mapper

import (
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
)

// CategoryListMapper はカテゴリー一覧データのマッピングを担当
type CategoryListMapper struct{}

// NewCategoryListMapper は新しいCategoryListMapperを作成
func NewCategoryListMapper() *CategoryListMapper {
	return &CategoryListMapper{}
}

// ToCategoryListResponse はモデルからOpenAPIレスポンスへ変換
func (m *CategoryListMapper) ToCategoryListResponse(categories []*reader.CategoryWithCount) openapi.CategoryList {
	items := make([]openapi.Category, 0, len(categories))

	for _, c := range categories {
		// 商品数を取得
		productCount := int(c.ProductCount)

		// 親カテゴリIDの処理
		var parentId *int64
		if c.Category.ParentID.Valid {
			id := int64(c.Category.ParentID.Int)
			parentId = &id
		}

		// null.Stringをポインタに変換
		var description *string
		if c.Category.Description.Valid {
			description = &c.Category.Description.String
		}

		var imageURL *string
		if c.Category.ImageURL.Valid {
			imageURL = &c.Category.ImageURL.String
		}

		items = append(items, openapi.Category{
			Id:           int64(c.Category.ID),
			Name:         c.Category.Name,
			Slug:         c.Category.Slug,
			Description:  description,
			ImageUrl:     imageURL,
			ParentId:     parentId,
			ProductCount: &productCount,
		})
	}

	return openapi.CategoryList{
		Items: items,
	}
}

// PresentError はエラーレスポンスを生成
func (m *CategoryListMapper) PresentError(code, message string, details any) openapi.ErrorResponse {
	detailsMap := map[string]any{
		"error": details,
	}
	if dm, ok := details.(map[string]any); ok {
		detailsMap = dm
	}

	return openapi.ErrorResponse{
		Code:    code,
		Message: message,
		Details: &detailsMap,
	}
}

// PresentInternalServerError は内部サーバーエラーレスポンスを生成
func (m *CategoryListMapper) PresentInternalServerError(message string, err error) openapi.ErrorResponse {
	return m.PresentError("internal_server_error", message, err.Error())
}
