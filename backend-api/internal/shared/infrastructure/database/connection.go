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

	log.Printf("Connected to database: %s", dbConfig.Host)
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

// 以下は下位互換性のための変数と関数（段階的に削除予定）

var (
	// DB はデータベース接続を保持するグローバル変数（下位互換性のため）
	// 新しいコードではDBManagerを使用してください
	DB *sql.DB
)

// InitDatabase は下位互換性のためのラッパー関数
// 新しいコードではNewDBManagerを使用してください
func InitDatabase(dbConfig config.DatabaseConfig) (*sql.DB, error) {
	manager, err := NewDBManager(dbConfig)
	if err != nil {
		return nil, err
	}

	// グローバル変数にも設定（下位互換性のため）
	DB = manager.DB()
	return manager.DB(), nil
}

// CloseDatabase はデータベース接続を閉じます（下位互換性のため）
// 新しいコードではDBManager.Close()を使用してください
func CloseDatabase() error {
	if DB != nil {
		if err := DB.Close(); err != nil {
			return fmt.Errorf("failed to close database: %w", err)
		}
		log.Println("Database connection closed")
	}
	return nil
}
