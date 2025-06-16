package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"
)

// TracingMiddleware はHTTPリクエストのトレーシングミドルウェア
func TracingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TracerProviderが設定されていない場合はスキップ
			if otel.GetTracerProvider() == nil {
				return next(c)
			}

			// トレースコンテキストの抽出
			ctx := otel.GetTextMapPropagator().Extract(
				c.Request().Context(),
				propagation.HeaderCarrier(c.Request().Header),
			)

			// スパン開始
			route := getRoutePattern(c)
			spanName := fmt.Sprintf("%s %s", c.Request().Method, route)
			ctx, span := tracer.Start(ctx, spanName, trace.WithSpanKind(trace.SpanKindServer))
			defer span.End()

			// コンテキストを設定
			c.SetRequest(c.Request().WithContext(ctx))

			// HTTP属性を設定
			span.SetAttributes(
				semconv.HTTPRequestMethodKey.String(c.Request().Method),
				semconv.HTTPRoute(route),
				semconv.URLFull(c.Request().URL.String()),
				semconv.URLScheme(c.Scheme()),
				semconv.ServerAddress(c.Request().Host),
				semconv.UserAgentOriginal(c.Request().UserAgent()),
				attribute.String("http.remote_addr", c.RealIP()),
			)

			// リクエストサイズ
			if c.Request().ContentLength > 0 {
				span.SetAttributes(semconv.HTTPRequestBodySize(int(c.Request().ContentLength)))
			}

			// X-Forwarded-For ヘッダーがある場合は追加
			if forwardedFor := c.Request().Header.Get("X-Forwarded-For"); forwardedFor != "" {
				span.SetAttributes(attribute.String("http.x_forwarded_for", forwardedFor))
			}

			// Referer ヘッダーがある場合は追加
			if referer := c.Request().Referer(); referer != "" {
				span.SetAttributes(attribute.String("http.referer", referer))
			}

			// Content-Type ヘッダーがある場合は追加
			if contentType := c.Request().Header.Get("Content-Type"); contentType != "" {
				span.SetAttributes(attribute.String("http.request.content_type", contentType))
			}

			// Accept ヘッダーがある場合は追加
			if accept := c.Request().Header.Get("Accept"); accept != "" {
				span.SetAttributes(attribute.String("http.request.accept", accept))
			}

			// クエリパラメータがある場合は追加
			if rawQuery := c.Request().URL.RawQuery; rawQuery != "" {
				span.SetAttributes(attribute.String("http.request.query", rawQuery))
			}

			// 次のハンドラーを実行
			err := next(c)

			// レスポンス属性を設定
			statusCode := c.Response().Status
			span.SetAttributes(semconv.HTTPResponseStatusCode(statusCode))

			// レスポンスサイズ（可能な場合）
			if size := c.Response().Size; size > 0 {
				span.SetAttributes(semconv.HTTPResponseBodySize(int(size)))
			}

			// エラーハンドリング
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
				span.SetAttributes(
					attribute.Bool("error", true),
					attribute.String("error.type", "handler_error"),
				)
			} else if statusCode >= 400 {
				errorMessage := fmt.Sprintf("HTTP %d", statusCode)
				span.SetStatus(codes.Error, errorMessage)
				span.SetAttributes(
					attribute.Bool("error", true),
					attribute.String("error.type", getErrorType(statusCode)),
				)
			}

			return err
		}
	}
}

// getErrorType はHTTPステータスコードからエラータイプを判定します
func getErrorType(statusCode int) string {
	switch {
	case statusCode >= 400 && statusCode < 500:
		return "client_error"
	case statusCode >= 500:
		return "server_error"
	default:
		return "unknown_error"
	}
}
