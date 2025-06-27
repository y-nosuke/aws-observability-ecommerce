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

// Start はスパンを開始します
func Start(ctx context.Context, kv ...attribute.KeyValue) (context.Context, *Observer) {
	funcInfo, file, line := utils.ParseFuncInfo(1)

	spanName := fmt.Sprintf("%s.%s", funcInfo.Receiver, funcInfo.FuncName)

	attrs := []attribute.KeyValue{
		attribute.String("file", file),
		attribute.Int("line", line),
	}

	attrs = append(attrs, kv...)

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
