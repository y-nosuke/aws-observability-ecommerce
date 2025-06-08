package provider

import (
	"github.com/google/wire"

	queryHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/handler"
)

// QueryProviderSet はクエリ（Read）関連のProvider Set
var QueryProviderSet = wire.NewSet(
	// Handler層のみ（データベース関連は SharedProviderSet で管理）
	queryHandler.NewProductCatalogHandler,
	queryHandler.NewCategoryListHandler,
	queryHandler.NewProductDetailHandler,
)
