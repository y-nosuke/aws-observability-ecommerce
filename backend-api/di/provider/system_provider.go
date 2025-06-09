package provider

import (
	"github.com/google/wire"

	systemHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/system/presentation/rest/handler"
)

// SystemProviderSet はシステム関連のProvider Set
var SystemProviderSet = wire.NewSet(
	// Handler層
	systemHandler.NewHealthHandler,
)
