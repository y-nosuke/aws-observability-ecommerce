package observability

import (
	"context"
)

// contextKey はコンテキストキーの型
type contextKey string

const (
	BusinessDomainKey contextKey = "app.business_domain"
	EntityTypeKey     contextKey = "app.entity_type"
	EntityIDKey       contextKey = "app.entity_id"
	OperationTypeKey  contextKey = "app.operation_type"
)

// GetDomainFromContext はcontextに設定されたドメイン情報を取得
func GetDomainFromContext(ctx context.Context) string {
	if domain, ok := ctx.Value(BusinessDomainKey).(string); ok && domain != "" {
		return domain
	}

	return "unknown"
}

// GetEntityIDFromContext はcontextからエンティティIDを取得
func GetEntityIDFromContext(ctx context.Context) int {
	if entityID, ok := ctx.Value(EntityIDKey).(int); ok {
		return entityID
	}
	return 0
}

// GetEntityTypeFromContext はcontextからエンティティタイプを取得
func GetEntityTypeFromContext(ctx context.Context) string {
	if entityType, ok := ctx.Value(EntityTypeKey).(string); ok && entityType != "" {
		return entityType
	}
	return "unknown"
}

// GetOperationTypeFromContext はcontextから操作タイプを取得
func GetOperationTypeFromContext(ctx context.Context) string {
	if operationType, ok := ctx.Value(OperationTypeKey).(string); ok && operationType != "" {
		return operationType
	}
	return "unknown"
}
