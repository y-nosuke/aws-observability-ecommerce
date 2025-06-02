package handler

import (
	"context"

	awsts "github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
)

type IAMHealthChecker struct {
	AWSFactory *aws.ClientFactory
}

func (i *IAMHealthChecker) Name() string { return "iam_auth" }
func (i *IAMHealthChecker) Check(ctx context.Context) error {
	if i.AWSFactory == nil {
		return &HealthCheckError{Msg: "credentials invalid"}
	}

	awsConf := i.AWSFactory.GetConfig()
	stsClient := awsts.NewFromConfig(*awsConf)
	_, err := stsClient.GetCallerIdentity(ctx, &awsts.GetCallerIdentityInput{})
	if err != nil {
		return &HealthCheckError{Msg: "credentials invalid", OriginalErr: err}
	}

	return nil
}
