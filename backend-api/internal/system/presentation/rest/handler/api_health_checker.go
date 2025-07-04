package handler

import (
	"context"
)

type ApiHealthChecker struct{}

func (*ApiHealthChecker) Name() string { return "api_server" }

func (*ApiHealthChecker) Check(_ context.Context) error {
	return nil
}
