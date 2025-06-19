package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/XSAM/otelsql"
	_ "github.com/go-sql-driver/mysql"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// DBManager はデータベース接続を管理する構造体
type DBManager struct {
	db *sql.DB
}

// NewDBManager はDBManagerのコンストラクタ（wireプロバイダー）
func NewDBManager(dbConfig config.DatabaseConfig) (*DBManager, error) {
	// データベース接続情報の構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	// otelsqlを使ってデータベース接続を作成（トレーシング対応）
	db, err := otelsql.Open("mysql", dsn,
		otelsql.WithAttributes(
			semconv.DBSystemNameMySQL,
			semconv.DBNamespace(dbConfig.Name),
		),
		otelsql.WithSpanOptions(otelsql.SpanOptions{
			Ping:           true,
			RowsNext:       true,
			DisableErrSkip: true,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open database with otelsql: %w", err)
	}

	// 接続設定
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Minute)

	// 接続の確認
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// DBStats メトリクスを登録
	err = otelsql.RegisterDBStatsMetrics(db, otelsql.WithAttributes(
		semconv.DBSystemNameMySQL,
		semconv.DBNamespace(dbConfig.Name),
	))
	if err != nil {
		log.Printf("Warning: failed to register DB stats metrics: %v", err)
		// メトリクス登録の失敗は致命的ではないので、エラーを返さずに続行
	}

	log.Printf("Connected to database: %s (with OpenTelemetry tracing enabled)", dbConfig.Host)
	return &DBManager{db: db}, nil
}

// DB はデータベース接続を返します
func (m *DBManager) DB() *sql.DB {
	return m.db
}

// Close はデータベース接続を閉じます
func (m *DBManager) Close() error {
	if m.db != nil {
		if err := m.db.Close(); err != nil {
			return fmt.Errorf("failed to close database: %w", err)
		}
		log.Println("Database connection closed")
	}
	return nil
}
