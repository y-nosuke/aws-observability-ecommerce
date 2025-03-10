package tracing

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
)

// XRayClient はX-Ray SDKとの連携を管理します
type XRayClient struct {
	client *xray.Client
	logger xraylog.Logger
}

// NewXRayClient はX-Rayクライアントを作成します
func NewXRayClient(awsCfg aws.Config) *XRayClient {
	client := xray.NewFromConfig(awsCfg)

	return &XRayClient{
		client: client,
		logger: xraylog.NewDefaultLogger(xraylog.LogLevelInfo),
	}
}

// 簡略化のため、実装の詳細は省略しています
