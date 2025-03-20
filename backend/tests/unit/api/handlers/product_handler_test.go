package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/handlers"
)

func TestProductHandlerReturnsProducts(t *testing.T) {
	// Echoインスタンスの作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーの作成
	h := handlers.NewProductHandler()

	// リクエストの実行
	if assert.NoError(t, h.HandleGetProducts(c)) {
		// レスポンスのアサーション
		assert.Equal(t, http.StatusOK, rec.Code)

		// JSONの解析
		var response handlers.PaginatedResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		// レスポンスの内容を検証
		assert.NotNil(t, response.Items)
		assert.Greater(t, response.TotalItems, 0)
		assert.Equal(t, 1, response.Page) // デフォルトは1ページ目
		assert.NotZero(t, response.PageSize)
		assert.NotZero(t, response.TotalPages)

		// 製品アイテムの検証
		products, ok := response.Items.([]interface{})
		assert.True(t, ok)
		if len(products) > 0 {
			product := products[0].(map[string]interface{})
			assert.NotNil(t, product["id"])
			assert.NotNil(t, product["name"])
			assert.NotNil(t, product["price"])
		}
	}
}

func TestProductHandlerPagination(t *testing.T) {
	// Echoインスタンスの作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/products?page=2&page_size=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// クエリパラメータを設定
	q := req.URL.Query()
	q.Set("page", "2")
	q.Set("page_size", "2")
	req.URL.RawQuery = q.Encode()

	// ハンドラーの作成
	h := handlers.NewProductHandler()

	// リクエストの実行
	if assert.NoError(t, h.HandleGetProducts(c)) {
		// JSONの解析
		var response handlers.PaginatedResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		// ページネーションが正しく機能していることを確認
		assert.Equal(t, 2, response.Page)
		assert.Equal(t, 2, response.PageSize)

		// 2ページ目のデータが返されていることを確認
		products, ok := response.Items.([]interface{})
		assert.True(t, ok)
		assert.LessOrEqual(t, len(products), 2) // ページサイズ以下のアイテム数
	}
}

func TestProductHandlerFilterByCategory(t *testing.T) {
	// Echoインスタンスの作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/products?category_id=1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// クエリパラメータを設定
	q := req.URL.Query()
	q.Set("category_id", "1")
	req.URL.RawQuery = q.Encode()

	// ハンドラーの作成
	h := handlers.NewProductHandler()

	// リクエストの実行
	if assert.NoError(t, h.HandleGetProducts(c)) {
		// JSONの解析
		var response handlers.PaginatedResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		// 応答に製品が含まれていることを確認
		products, ok := response.Items.([]interface{})
		assert.True(t, ok)

		// すべての製品がカテゴリーID=1に属していることを確認
		for _, p := range products {
			product, ok := p.(map[string]interface{})
			assert.True(t, ok)
			categoryID := int(product["category_id"].(float64))
			assert.Equal(t, 1, categoryID)
		}
	}
}
