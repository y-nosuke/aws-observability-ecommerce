package product

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"

	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_DSN")+"?parseTime=true")
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping test database: %v", err)
	}

	return db
}

func TestProductRepository_CRUD(t *testing.T) {
	// テストの際は実際のDBに接続します
	// 本来は専用のテスト用DBやトランザクションロールバックを使うべきです
	db := setupTestDB(t)
	defer db.Close()

	repo := New(db)
	ctx := context.Background()

	// テスト用の商品を作成
	p := &models.Product{
		Name:        "Test Product",
		Description: null.StringFrom("Test Description"),
		Price:       decimal.New(1000, 0),
		CategoryID:  1, // 既存のカテゴリーIDを指定
	}

	// Create
	err := repo.Create(ctx, p)
	assert.NoError(t, err)
	assert.NotZero(t, p.ID, "商品IDが設定されていません")

	// FindByID
	found, err := repo.FindByID(ctx, p.ID)
	assert.NoError(t, err)
	assert.Equal(t, p.Name, found.Name)

	// Update
	p.Name = "Updated Test Product"
	err = repo.Update(ctx, p)
	assert.NoError(t, err)

	updated, err := repo.FindByID(ctx, p.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Test Product", updated.Name)

	// FindAll
	products, err := repo.FindAll(ctx, 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)

	// Count
	count, err := repo.Count(ctx)
	assert.NoError(t, err)
	assert.NotZero(t, count)

	// Delete
	err = repo.Delete(ctx, p.ID)
	assert.NoError(t, err)

	// 削除されたことを確認
	_, err = repo.FindByID(ctx, p.ID)
	assert.Error(t, err)
}
