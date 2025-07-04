package handler

import (
	"context"
	"database/sql"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// クライアント用・ログ用メッセージを持つエラー型
// クライアントにはClientMsg、ログにはLogMsgを使う

type HealthCheckError struct {
	Msg         string
	OriginalErr error
}

func (e *HealthCheckError) Error() string {
	if e.OriginalErr != nil {
		return e.OriginalErr.Error()
	}

	return e.Msg
}

type HealthChecker interface {
	Name() string
	Check(ctx context.Context) error
}

// HealthCheckers構造体

type HealthCheckers struct {
	items []HealthChecker
}

func NewHealthCheckers(db *sql.DB, stsClient *sts.Client, s3Client *s3.Client, config config.S3Config, checks []string) *HealthCheckers {
	mapper := map[string]HealthChecker{
		"db":  &DatabaseHealthChecker{DB: db},
		"iam": &IAMHealthChecker{stsClient: stsClient},
		"s3":  &S3HealthChecker{s3Client: s3Client, config: config},
	}

	items := []HealthChecker{&ApiHealthChecker{}}
	for _, s := range []string{"db", "iam", "s3"} {
		if contains(checks, s) {
			items = append(items, mapper[s])
		}
	}

	return &HealthCheckers{
		items: items,
	}
}

func (h *HealthCheckers) Check(ctx context.Context) map[string]error {
	results := map[string]error{}
	for _, item := range h.items {
		results[item.Name()] = item.Check(ctx)
	}

	return results
}

func contains(list []string, target string) bool {
	for _, v := range list {
		if strings.TrimSpace(v) == target {
			return true
		}
	}
	return false
}
