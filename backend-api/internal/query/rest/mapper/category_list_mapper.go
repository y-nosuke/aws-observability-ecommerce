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
		var parentId *int
		if c.Category.ParentID.Valid {
			id := int(c.Category.ParentID.V)
			parentId = &id
		}

		// null.Stringをポインタに変換
		var description *string
		if c.Category.Description.Valid {
			description = &c.Category.Description.V
		}

		var imageURL *string
		if c.Category.ImageURL.Valid {
			imageURL = &c.Category.ImageURL.V
		}

		items = append(items, openapi.Category{
			Id:           int(c.Category.ID),
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
