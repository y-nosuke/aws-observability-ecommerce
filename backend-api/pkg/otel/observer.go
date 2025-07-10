package otel

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/utils"
)

type Observer struct {
	ctx      context.Context
	span     trace.Span
	spanName string
}

// StartOption はStart関数のオプションを定義する型
type StartOption func(*startConfig)

// startConfig はStart関数の設定を保持する構造体
type startConfig struct {
	skip       int
	attributes []attribute.KeyValue
}

// WithSkip はスキップレベルを指定するオプション
func WithSkip(skip int) StartOption {
	return func(c *startConfig) {
		c.skip = skip
	}
}

// WithAttributes は属性を指定するオプション
func WithAttributes(attrs ...attribute.KeyValue) StartOption {
	return func(c *startConfig) {
		c.attributes = append(c.attributes, attrs...)
	}
}

// Start はスパンを開始します（Functional Optionパターン使用）
func Start(ctx context.Context, opts ...StartOption) (context.Context, *Observer) {
	// デフォルト設定
	config := &startConfig{
		skip: 1, // デフォルトは1（Start関数の直接の呼び出し元）
	}

	// オプションを適用
	for _, opt := range opts {
		opt(config)
	}

	funcInfo, file, line := utils.ParseFuncInfo(config.skip)

	spanName := fmt.Sprintf("%s.%s", funcInfo.Receiver, funcInfo.FuncName)

	attrs := []attribute.KeyValue{
		attribute.String("file", file),
		attribute.Int("line", line),
	}

	attrs = append(attrs, config.attributes...)

	spanCtx, span := tracer.Start(ctx, spanName, trace.WithAttributes(attrs...))

	slog.InfoContext(spanCtx, fmt.Sprintf("start %s", spanName))

	return spanCtx, &Observer{ctx: ctx, span: span, spanName: spanName}
}

// End はスパンを終了します
func (o *Observer) End(err error) {
	if err != nil {
		o.span.RecordError(err)
		o.span.SetStatus(codes.Error, fmt.Sprintf("%s %s", o.spanName, err.Error()))
	}
	slog.InfoContext(o.ctx, fmt.Sprintf("end %s", o.spanName))
	o.span.End()
}

// SetAttributes はスパンの属性を設定します
func (o *Observer) SetAttributes(kv ...attribute.KeyValue) {
	o.span.SetAttributes(kv...)
}
