package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/XSAM/otelsql"
	"github.com/stephenafamo/bob"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

var (
	bobDb bob.DB
)

// NewDBConfig はDB接続とクリーンアップ関数を返すプロバイダー
func NewDBConfig(ctx context.Context, dbConfig config.DatabaseConfig) (*sql.DB, func(), error) {
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
			semconv.ServerAddress(dbConfig.Host),
			semconv.ServerPort(dbConfig.Port),
			semconv.DBNamespace(dbConfig.Name),
			semconv.PeerService("MySQL"),
		),
		otelsql.WithSpanOptions(otelsql.SpanOptions{
			Ping:           true,  // 接続確認はトレースする
			RowsNext:       false, // 行ごとのイベントは抑制し、パフォーマンスを優先
			DisableErrSkip: true,  // driver.ErrSkip はエラーとして記録しない
			DisableQuery:   false, // SQLクエリ本文は記録する
		}),
	)
	if err != nil {
		log.Fatalf("db open: %v", err)
	}

	bobDb = bob.NewDB(db)

	// 接続設定
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Minute)

	// 接続の確認
	if err = db.PingContext(ctx); err != nil {
		if closeErr := db.Close(); closeErr != nil {
			return nil, nil, fmt.Errorf("failed to close database: %w", err)
		}
		return nil, nil, fmt.Errorf("failed to ping database: %w", err)
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

	cleanup := func() {
		if err = db.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}

	return db, cleanup, nil
}

// GetDB はDB接続を返す
func GetDB() bob.DB {
	return bobDb
}
