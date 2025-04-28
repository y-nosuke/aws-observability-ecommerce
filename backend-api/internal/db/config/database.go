package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/config"
)

var (
	// DB はデータベース接続を保持するグローバル変数
	DB *sql.DB
)

// InitDatabase はデータベース接続を初期化します
func InitDatabase() error {
	// データベース接続情報の構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	// データベース接続の作成
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// 接続設定
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// 接続の確認
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db
	log.Printf("Connected to database: %s", config.Database.Host)
	return nil
}

// CloseDatabase はデータベース接続を閉じます
func CloseDatabase() error {
	if DB != nil {
		if err := DB.Close(); err != nil {
			return fmt.Errorf("failed to close database: %w", err)
		}
		log.Println("Database connection closed")
	}
	return nil
}
