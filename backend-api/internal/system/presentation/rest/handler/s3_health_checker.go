package handler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

type S3HealthChecker struct {
	s3Client *s3.Client
	config   config.S3Config
}

func (s *S3HealthChecker) Name() string { return "s3_connectivity" }

func (s *S3HealthChecker) Check(ctx context.Context) error {
	_, err := s.s3Client.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: &s.config.BucketName})
	if err != nil {
		return &HealthCheckError{Msg: "connection failed", OriginalErr: err}
	}

	return nil
}
