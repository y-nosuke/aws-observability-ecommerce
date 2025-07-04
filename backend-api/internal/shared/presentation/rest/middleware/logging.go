package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// LoggingMiddleware 構造化ログミドルウェアを作成
func LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// ===== リクエスト情報の取得 =====
			req := c.Request()

			// ヘッダ
			reqHeader := req.Header

			// パスパラメータ
			pathParams := map[string]string{}
			for _, name := range c.ParamNames() {
				pathParams[name] = c.Param(name)
			}

			// クエリパラメータ
			queryParams := c.QueryParams()

			// ボディ
			reqBodyBytes, ok := c.Get(ContextKeyRequestBody).(*[]byte)
			if !ok {
				reqBodyBytes = &[]byte{}
			}

			// HTTPリクエストログを出力
			var request string
			if req.URL.RawQuery != "" {
				request = fmt.Sprintf("%s %s?%s", req.Method, req.URL.Path, req.URL.RawQuery)
			} else {
				request = fmt.Sprintf("%s %s", req.Method, req.URL.Path)
			}
			slog.InfoContext(req.Context(),
				fmt.Sprintf("HTTP request start: %s", request),
				"method", req.Method,
				"path", req.URL.Path,
				"path_params", formatMap(pathParams),
				"query_params", formatMap(queryParams),
				"request_headers", formatMap(reqHeader),
				"request_body", formatBodyForLogging(*reqBodyBytes, req.Header.Get("Content-Type")),
				"request_body_size", len(*reqBodyBytes),
				"user_agent", req.UserAgent(),
				"remote_ip", c.RealIP(),
				"x_forwarded_for", req.Header.Get("X-Forwarded-For"),
				"referer", req.Referer(),
				"content_type", req.Header.Get("Content-Type"),
				"content_length", req.ContentLength,
				"accept", req.Header.Get("Accept"),
				"query", req.URL.RawQuery,
			)

			// ===== ハンドラー実行 =====
			err := next(c)

			// ===== レスポンス情報の取得 =====
			res := c.Response()

			// 処理時間を計測
			duration := time.Since(start)

			// レスポンスヘッダ
			resHeader := res.Header()

			// ボディ
			resBodyBytes, ok := c.Get(ContextKeyResponseBody).(*[]byte)
			if !ok {
				resBodyBytes = &[]byte{}
			}

			// HTTPレスポンスログを出力
			slog.InfoContext(req.Context(),
				fmt.Sprintf("HTTP request end: %s", request),
				"method", req.Method,
				"path", req.URL.Path,
				"status", res.Status,
				"response_headers", formatMap(resHeader),
				"response_body", formatBodyForLogging(*resBodyBytes, res.Header().Get("Content-Type")),
				"response_body_size", len(*resBodyBytes),
				"content_type", res.Header().Get("Content-Type"),
				"response_size", res.Size,
				"duration_ms", duration.Milliseconds(),
				"cache_hit", false, // TODO: キャッシュ実装後に動的に設定
				"database_queries", 0, // TODO: DB監視実装後に動的に設定
				"external_api_calls", 0, // TODO: 外部API監視実装後に動的に設定
			)

			return err
		}
	}
}

// formatBodyForLogging はボディをログ出力用にフォーマットします
func formatBodyForLogging(bodyBytes []byte, contentType string) string {
	if len(bodyBytes) == 0 {
		return ""
	}

	// バイナリデータかどうかを判定
	if isBinaryContent(contentType) {
		return fmt.Sprintf("[BINARY_DATA: %s, %d bytes]", contentType, len(bodyBytes))
	}

	// テキストデータの場合は文字列として出力（サイズ制限付き）
	const maxBodySize = 1024 // 1KB制限
	if len(bodyBytes) > maxBodySize {
		return fmt.Sprintf("%s... [truncated, total: %d bytes]", string(bodyBytes[:maxBodySize]), len(bodyBytes))
	}

	return string(bodyBytes)
}

// isBinaryContent はContent-Typeからバイナリデータかどうかを判定します
func isBinaryContent(contentType string) bool {
	contentType = strings.ToLower(contentType)

	// バイナリ系Content-Type
	binaryTypes := []string{
		"multipart/form-data",
		"image/",
		"video/",
		"audio/",
		"application/octet-stream",
		"application/pdf",
		"application/zip",
		"application/x-",
	}

	for _, binaryType := range binaryTypes {
		if strings.Contains(contentType, binaryType) {
			return true
		}
	}

	return false
}

// formatMap はマップを文字列に変換します
func formatMap(v interface{}) string {
	switch m := v.(type) {
	case http.Header:
		sb := strings.Builder{}
		for k, vv := range m {
			sb.WriteString(k + ": " + strings.Join(vv, ",") + "; ")
		}
		return sb.String()

	case url.Values:
		sb := strings.Builder{}
		for k, v := range m {
			sb.WriteString(k + ": " + strings.Join(v, ",") + "; ")
		}
		return sb.String()

	case map[string]string:
		sb := strings.Builder{}
		for k, v := range m {
			sb.WriteString(k + ": " + v + "; ")
		}
		return sb.String()

	case map[string][]string:
		sb := strings.Builder{}
		for k, vv := range m {
			sb.WriteString(k + ": " + strings.Join(vv, ",") + "; ")
		}
		return sb.String()

	default:
		return "unknown format"
	}
}
