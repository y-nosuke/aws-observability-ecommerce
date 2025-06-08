package logging

import (
	"context"
	"time"
)

// RequestLogData はリクエストログのデータ構造
type RequestLogData struct {
	Method           string
	Path             string
	Query            string
	StatusCode       int
	RequestSize      int64
	ResponseSize     int64
	Duration         time.Duration
	UserAgent        string
	RemoteIP         string
	XForwardedFor    string
	Referer          string
	ContentType      string
	Accept           string
	UserID           string
	SessionID        string
	UserRole         string
	CacheHit         bool
	DatabaseQueries  int
	ExternalAPICalls int
}

// LogRequest はリクエストログを出力します
func (l *StructuredLogger) LogRequest(ctx context.Context, req RequestLogData) {
	fields := []Field{
		{Key: "log_type", Value: "request"},
		{Key: "http", Value: map[string]interface{}{
			"method":              req.Method,
			"path":                req.Path,
			"query":               req.Query,
			"status_code":         req.StatusCode,
			"request_size_bytes":  req.RequestSize,
			"response_size_bytes": req.ResponseSize,
			"duration_ms":         float64(req.Duration.Nanoseconds()) / 1e6,
			"user_agent":          req.UserAgent,
			"remote_ip":           req.RemoteIP,
			"x_forwarded_for":     req.XForwardedFor,
			"referer":             req.Referer,
			"content_type":        req.ContentType,
			"accept":              req.Accept,
		}},
		{Key: "response", Value: map[string]interface{}{
			"cache_hit":          req.CacheHit,
			"database_queries":   req.DatabaseQueries,
			"external_api_calls": req.ExternalAPICalls,
		}},
	}

	if req.UserID != "" {
		fields = append(fields, Field{Key: "user", Value: map[string]interface{}{
			"id":         req.UserID,
			"session_id": req.SessionID,
			"role":       req.UserRole,
		}})
	}

	l.Info(ctx, "HTTP request processed", fields...)
}
