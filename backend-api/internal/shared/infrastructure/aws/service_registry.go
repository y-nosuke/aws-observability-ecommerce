package aws

import (
	"context"
	"fmt"

	configPkg "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// ServiceRegistry はAWSサービスのレジストリ
type ServiceRegistry struct {
	clientFactory   *ClientFactory
	s3ClientWrapper *S3ClientWrapper
}

// NewServiceRegistry は新しいAWSサービスレジストリを作成します
func NewServiceRegistry(ctx context.Context, awsConfig configPkg.AWSConfig) (*ServiceRegistry, error) {
	clientFactory, err := NewClientFactory(ctx, awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create client factory: %w", err)
	}

	s3Client := clientFactory.GetS3Client()
	s3ClientWrapper := NewS3ClientWrapper(s3Client, awsConfig.S3)

	return &ServiceRegistry{
		clientFactory:   clientFactory,
		s3ClientWrapper: s3ClientWrapper,
	}, nil
}

// GetS3ClientWrapper はS3クライアントラッパーを取得します
func (r *ServiceRegistry) GetS3ClientWrapper() *S3ClientWrapper {
	return r.s3ClientWrapper
}

// GetClientFactory はクライアントファクトリを取得します
func (r *ServiceRegistry) GetClientFactory() *ClientFactory {
	return r.clientFactory
}
