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
)

// SetDomainToContext はcontextにドメイン情報を設定
func SetDomainToContext(ctx context.Context, domain string) context.Context {
	return context.WithValue(ctx, BusinessDomainKey, domain)
}

// GetDomainFromContext はcontextに設定されたドメイン情報を取得
func GetDomainFromContext(ctx context.Context) string {
	if domain, ok := ctx.Value(BusinessDomainKey).(string); ok && domain != "" {
		return domain
	}

	return "unknown"
}

// SetEntityTypeToContext はcontextにエンティティタイプを設定
func SetEntityTypeToContext(ctx context.Context, entityType string) context.Context {
	return context.WithValue(ctx, EntityTypeKey, entityType)
}

// GetEntityTypeFromContext はcontextからエンティティタイプを取得
func GetEntityTypeFromContext(ctx context.Context) string {
	if entityType, ok := ctx.Value(EntityTypeKey).(string); ok && entityType != "" {
		return entityType
	}
	return "unknown"
}

// SetEntityIDToContext はcontextにエンティティIDを設定
func SetEntityIDToContext(ctx context.Context, entityID int) context.Context {
	return context.WithValue(ctx, EntityIDKey, entityID)
}

// GetEntityIDFromContext はcontextからエンティティIDを取得
func GetEntityIDFromContext(ctx context.Context) int {
	if entityID, ok := ctx.Value(EntityIDKey).(int); ok {
		return entityID
	}
	return 0
}
