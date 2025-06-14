package database

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"
	"go.opentelemetry.io/otel/trace"
)

// TracingWrapper はデータベース接続のトレーシングラッパー
type TracingWrapper struct {
	*sql.DB
	dbName string
}

// NewTracingWrapper はトレーシング対応のDB接続を作成
func NewTracingWrapper(db *sql.DB, dbName string) *TracingWrapper {
	return &TracingWrapper{
		DB:     db,
		dbName: dbName,
	}
}

// QueryContext はクエリ実行をトレース
func (tw *TracingWrapper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	ctx, span := tw.startDBSpan(ctx, "db.query", query)
	defer span.End()

	start := time.Now()
	rows, err := tw.DB.QueryContext(ctx, query, args...)
	duration := time.Since(start)

	tw.finishDBSpan(span, duration, err, 0)
	return rows, err
}

// QueryRowContext は単一行クエリ実行をトレース
func (tw *TracingWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	ctx, span := tw.startDBSpan(ctx, "db.query_row", query)
	defer span.End()

	start := time.Now()
	row := tw.DB.QueryRowContext(ctx, query, args...)
	duration := time.Since(start)

	tw.finishDBSpan(span, duration, nil, 1)
	return row
}

// ExecContext は実行をトレース
func (tw *TracingWrapper) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	ctx, span := tw.startDBSpan(ctx, "db.exec", query)
	defer span.End()

	start := time.Now()
	result, err := tw.DB.ExecContext(ctx, query, args...)
	duration := time.Since(start)

	var rowsAffected int64
	if err == nil && result != nil {
		if affected, affectedErr := result.RowsAffected(); affectedErr == nil {
			rowsAffected = affected
		}
	}

	tw.finishDBSpan(span, duration, err, rowsAffected)
	return result, err
}

// PrepareContext はプリペアドステートメント作成をトレース
func (tw *TracingWrapper) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	ctx, span := tw.startDBSpan(ctx, "db.prepare", query)
	defer span.End()

	start := time.Now()
	stmt, err := tw.DB.PrepareContext(ctx, query)
	duration := time.Since(start)

	tw.finishDBSpan(span, duration, err, 0)
	return stmt, err
}

// BeginTx はトランザクション開始をトレース
func (tw *TracingWrapper) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tracer := otel.Tracer("aws-observability-ecommerce")

	ctx, span := tracer.Start(ctx, "db.begin_tx", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	// DB属性を設定
	span.SetAttributes(
		semconv.DBSystemNameMySQL,
		semconv.DBNamespaceKey.String(tw.dbName),
		semconv.DBOperationNameKey.String("BEGIN"),
		attribute.String("db.operation_type", "transaction"),
	)

	start := time.Now()
	tx, err := tw.DB.BeginTx(ctx, opts)
	duration := time.Since(start)

	span.SetAttributes(attribute.Int64("db.duration_ms", duration.Milliseconds()))

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.String("error.type", "transaction_error"))
	}

	return tx, err
}

// startDBSpan はデータベース操作のスパンを開始
func (tw *TracingWrapper) startDBSpan(ctx context.Context, operationName, query string) (context.Context, trace.Span) {
	tracer := otel.Tracer("aws-observability-ecommerce")

	ctx, span := tracer.Start(ctx, operationName, trace.WithSpanKind(trace.SpanKindClient))

	// DB属性を設定
	operation := extractOperation(query)
	table := extractTable(query)

	span.SetAttributes(
		semconv.DBSystemNameMySQL,
		semconv.DBNamespaceKey.String(tw.dbName),
		semconv.DBQueryTextKey.String(query),
		semconv.DBOperationNameKey.String(operation),
		attribute.String("db.sql.table", table),
	)

	return ctx, span
}

// finishDBSpan はデータベース操作のスパンを完了
func (tw *TracingWrapper) finishDBSpan(span trace.Span, duration time.Duration, err error, rowsAffected int64) {
	span.SetAttributes(attribute.Int64("db.duration_ms", duration.Milliseconds()))

	if rowsAffected > 0 {
		span.SetAttributes(attribute.Int64("db.rows_affected", rowsAffected))
	}

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.type", "database_error"),
		)
	}
}

// extractOperation はSQLクエリから操作種別を抽出
func extractOperation(query string) string {
	if len(query) < 6 {
		return "unknown"
	}
	trimmed := strings.TrimSpace(query)
	parts := strings.Fields(trimmed)
	if len(parts) == 0 {
		return "unknown"
	}
	return strings.ToUpper(parts[0])
}

// extractTable はSQLクエリからテーブル名を抽出（簡易版）
func extractTable(query string) string {
	// 簡易的な実装（本格的にはSQLパーサーが必要）
	query = strings.ToLower(strings.TrimSpace(query))

	if strings.HasPrefix(query, "select") {
		if idx := strings.Index(query, "from "); idx != -1 {
			parts := strings.Fields(query[idx+5:])
			if len(parts) > 0 {
				// バッククオートとスペースを除去
				table := strings.Trim(parts[0], "`")
				// JOINがある場合は最初のテーブルのみ
				if commaIdx := strings.Index(table, ","); commaIdx != -1 {
					table = table[:commaIdx]
				}
				return table
			}
		}
	} else if strings.HasPrefix(query, "insert into ") {
		parts := strings.Fields(query[12:])
		if len(parts) > 0 {
			return strings.Trim(parts[0], "`")
		}
	} else if strings.HasPrefix(query, "update ") {
		parts := strings.Fields(query[7:])
		if len(parts) > 0 {
			return strings.Trim(parts[0], "`")
		}
	} else if strings.HasPrefix(query, "delete from ") {
		parts := strings.Fields(query[12:])
		if len(parts) > 0 {
			return strings.Trim(parts[0], "`")
		}
	} else if strings.HasPrefix(query, "replace into ") {
		parts := strings.Fields(query[13:])
		if len(parts) > 0 {
			return strings.Trim(parts[0], "`")
		}
	}

	return "unknown"
}
