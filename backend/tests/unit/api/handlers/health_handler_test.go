package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/handlers"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/config"
)

func TestHealthHandlerBasicResponse(t *testing.T) {
	// Echoインスタンスの作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// モック設定
	originalVersion := config.App.Version
	config.App.Version = "test-1.0.0"
	defer func() { config.App.Version = originalVersion }()

	// ハンドラーの作成
	h := handlers.NewHealthHandler()

	// 意図的に少し待機してアップタイムを確認
	time.Sleep(100 * time.Millisecond)

	// リクエストの実行
	err := h.HandleHealthCheck(c)
	assert.NoError(t, err)

	// レスポンスのアサーション
	assert.Equal(t, http.StatusOK, rec.Code)

	// JSONの解析
	var response handlers.HealthResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 具体的な検証
	assert.Equal(t, "ok", response.Status)
	assert.Equal(t, "test-1.0.0", response.Version)
	assert.NotEmpty(t, response.Timestamp)
	assert.GreaterOrEqual(t, response.Uptime, int64(0)) // 0以上であること
}

func TestHealthHandlerResourceInformation(t *testing.T) {
	// Echoインスタンスの作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーの作成
	h := handlers.NewHealthHandler()

	// リクエストの実行
	err := h.HandleHealthCheck(c)
	assert.NoError(t, err)

	// JSONの解析
	var response handlers.HealthResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// リソース情報の詳細な検証
	resources := response.Resources
	assert.NotNil(t, resources, "リソース情報が存在すること")

	// メモリ情報の検証
	memory, ok := resources["memory"].(map[string]interface{})
	assert.True(t, ok, "メモリ情報が正しい型であること")

	assert.Contains(t, memory, "allocated", "アロケーテッドメモリ情報が存在すること")
	assert.Contains(t, memory, "total", "総メモリ情報が存在すること")
	assert.Contains(t, memory, "system", "システムメモリ情報が存在すること")

	// Goルーチン数の検証
	goroutines, ok := resources["goroutines"].(float64)
	assert.True(t, ok, "Goルーチン数が数値であること")
	assert.GreaterOrEqual(t, goroutines, float64(1), "少なくとも1つ以上のGoルーチンが存在すること")
}

func TestHealthHandlerServicesStatus(t *testing.T) {
	// Echoインスタンスの作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーの作成
	h := handlers.NewHealthHandler()

	// リクエストの実行
	err := h.HandleHealthCheck(c)
	assert.NoError(t, err)

	// JSONの解析
	var response handlers.HealthResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// サービスステータスの検証
	services := response.Services
	assert.NotNil(t, services, "サービス情報が存在すること")

	apiStatus, ok := services["api"].(map[string]interface{})
	assert.True(t, ok, "APIサービス情報が存在すること")

	status, ok := apiStatus["status"].(string)
	assert.True(t, ok, "APIステータスが文字列であること")
	assert.Equal(t, "up", status, "APIステータスが'up'であること")
}
