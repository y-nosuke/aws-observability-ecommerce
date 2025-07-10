package otel

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
)

// WithSpan は関数を実行する際に自動的にスパンを開始・終了します
func WithSpan(ctx context.Context, fn func(context.Context, *Observer) error, attrs ...attribute.KeyValue) error {
	spanCtx, o := Start(ctx, attrs...)
	err := fn(spanCtx, o)
	o.End(err)
	return err
}

// WithSpanValue は戻り値を持つ関数版です
func WithSpanValue[T any](ctx context.Context, fn func(context.Context, *Observer) (T, error), attrs ...attribute.KeyValue) (T, error) {
	spanCtx, o := Start(ctx, attrs...)
	result, err := fn(spanCtx, o)
	o.End(err)
	return result, err
}
