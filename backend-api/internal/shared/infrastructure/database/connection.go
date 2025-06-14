package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

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

	// データベース接続の作成
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// 接続設定
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Minute)

	// 接続の確認
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// トレーシングラッパーを適用
	tracingDB := NewTracingWrapper(db, dbConfig.Name)

	log.Printf("Connected to database: %s (with tracing enabled)", dbConfig.Host)
	return &DBManager{db: tracingDB.DB}, nil
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
