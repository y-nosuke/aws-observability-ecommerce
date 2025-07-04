package handler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type IAMHealthChecker struct {
	stsClient *sts.Client
}

func (i *IAMHealthChecker) Name() string { return "iam_auth" }

func (i *IAMHealthChecker) Check(ctx context.Context) error {
	if _, err := i.stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}); err != nil {
		return &HealthCheckError{Msg: "credentials invalid", OriginalErr: err}
	}

	return nil
}
