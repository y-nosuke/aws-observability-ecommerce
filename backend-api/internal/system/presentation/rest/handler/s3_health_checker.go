package handler

import (
	"context"

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

type S3HealthChecker struct {
	AWSFactory *aws.ClientFactory
}

func (s *S3HealthChecker) Name() string { return "s3_connectivity" }
func (s *S3HealthChecker) Check(ctx context.Context) error {
	if s.AWSFactory == nil {
		return &HealthCheckError{Msg: "IAM credentials invalid"}
	}

	bucket := config.AWS.S3.BucketName
	if bucket == "" {
		return &HealthCheckError{Msg: "bucket config not set"}
	}

	s3Client := s.AWSFactory.GetS3Client()
	_, err := s3Client.HeadBucket(ctx, &awss3.HeadBucketInput{Bucket: &bucket})
	if err != nil {
		return &HealthCheckError{Msg: "connection failed", OriginalErr: err}
	}

	return nil
}
