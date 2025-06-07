# Phase 4 実装完了: 特化ログ実装

## 📋 実装完了内容

### ✅ Phase 4: 特化ログ実装

- [x] リクエストログ実装
- [x] エラーログ実装
- [x] アプリケーションログ実装
- [x] ログヘルパー機能追加
- [x] ハンドラーへの統合
- [x] 実用例の実装

## 🚀 実装された機能

### 1. リクエストログ実装 (`request_log.go`)

```go
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

func (l *StructuredLogger) LogRequest(ctx context.Context, req RequestLogData)
```

**使用例（ミドルウェア）:**

```go
// StructuredLoggingMiddleware で自動的に出力
logger.LogRequest(c.Request().Context(), logData)
```

**出力JSON例:**

```json
{
  "timestamp": "2025-06-05T10:30:45.123456789Z",
  "level": "info",
  "message": "HTTP request processed",
  "service": {
    "name": "aws-observability-ecommerce",
    "version": "1.0.0",
    "environment": "development"
  },
  "trace": {
    "trace_id": "4bf92f3577b34da6a3ce929d0e0e4736",
    "span_id": "00f067aa0ba902b7"
  },
  "request": {
    "id": "req_abc123def456"
  },
  "log_type": "request",
  "http": {
    "method": "POST",
    "path": "/api/products/123/images",
    "status_code": 200,
    "duration_ms": 45.23,
    "request_size_bytes": 1024,
    "response_size_bytes": 2048
  }
}
```

### 2. エラーログ実装 (`error_log.go`)

```go
type ErrorContext struct {
    Operation      string
    ResourceType   string
    ResourceID     string
    RetryCount     int
    MaxRetries     int
    TimeoutMs      int
    Severity       string
    AffectedUsers  int
    BusinessImpact string
}

func (l *StructuredLogger) LogError(ctx context.Context, err error, errorCtx ErrorContext)
```

**使用例（エラーハンドリングミドルウェア）:**

```go
errorCtx := logging.ErrorContext{
    Operation:      "product_fetch",
    ResourceType:   "product",
    ResourceID:     "123",
    Severity:       "high",
    BusinessImpact: "service_degradation",
}
logger.LogError(ctx, err, errorCtx)
```

**出力JSON例:**

```json
{
  "timestamp": "2025-06-05T10:30:45.123456789Z",
  "level": "error",
  "message": "Application error occurred",
  "log_type": "error",
  "error": {
    "type": "DatabaseConnectionError",
    "message": "connection refused: mysql:3306",
    "code": "DB_CONN_REFUSED",
    "stack_trace": "main.go:45 -> database.go:123",
    "fingerprint": "db_connection_mysql_3306"
  },
  "context": {
    "operation": "product_fetch",
    "resource_type": "product",
    "resource_id": "123"
  },
  "impact": {
    "severity": "high",
    "business_impact": "service_degradation"
  }
}
```

### 3. アプリケーションログ実装 (`application_log.go`)

```go
type ApplicationOperation struct {
    Name             string
    Category         string
    Duration         time.Duration
    Success          bool
    Stage            string
    EntityType       string
    EntityID         string
    Action           string
    Source           string
    Data             map[string]interface{}
    PerformanceData  map[string]interface{}
}

func (l *StructuredLogger) LogApplication(ctx context.Context, op ApplicationOperation)
```

**使用例（商品画像アップロード）:**

```go
// 操作ログの開始
opLogger := logHelper.StartOperation(ctx, "upload_product_image", "product_management").
    WithEntity("product", "123").
    WithAction("upload", "admin_ui").
    WithData("file_name", "image.jpg")

// 処理実行...

// 成功時
opLogger.Complete(ctx)

// 失敗時
opLogger.Fail(ctx, err)
```

**出力JSON例:**

```json
{
  "timestamp": "2025-06-05T10:30:45.123456789Z",
  "level": "info",
  "message": "Application operation completed",
  "log_type": "application",
  "operation": {
    "name": "upload_product_image",
    "category": "product_management",
    "duration_ms": 1200.45,
    "success": true,
    "stage": "completed"
  },
  "business": {
    "entity_type": "product",
    "entity_id": "123",
    "action": "upload",
    "source": "admin_ui"
  },
  "data": {
    "file_name": "image.jpg",
    "file_size_bytes": 2048576,
    "image_format": "jpeg"
  },
  "performance": {
    "s3_upload_duration_ms": 890.12,
    "image_processing_duration_ms": 310.33
  }
}
```

### 4. ログヘルパー機能 (`log_helper.go`)

#### StartOperation - 操作追跡

```go
logHelper := logging.NewLogHelper(logger)
opLogger := logHelper.StartOperation(ctx, "upload_product_image", "product_management")
```

#### ビジネスイベント記録

```go
logHelper.LogBusinessEvent(ctx, "product_image_uploaded", "product", "123", map[string]interface{}{
    "image_url":   response.ImageURL,
    "file_name":   file.Filename,
    "uploaded_by": "admin",
})
```

#### パフォーマンスメトリクス記録

```go
logHelper.LogPerformanceMetric(ctx, "response_time", 120.5, "milliseconds")
```

## 🔧 実装された統合

### 1. ミドルウェア統合

- **RequestIDMiddleware**: 各リクエストに一意IDを生成
- **StructuredLoggingMiddleware**: HTTPリクエストを自動ログ出力
- **ErrorHandlingMiddleware**: エラーを構造化してログ出力

### 2. ハンドラー統合

```go
// ProductHandler にログ機能を統合
type ProductHandler struct {
    uploadProductImageUseCase *usecase.UploadProductImageUseCase
    getProductImageUseCase    *usecase.GetProductImageUseCase
    logHelper                 *logging.LogHelper  // 追加
}

func (h *ProductHandler) UploadProductImage(ctx echo.Context, id openapi.ProductIdParam) error {
    // 操作ログ開始
    opLogger := h.logHelper.StartOperation(ctx.Request().Context(), "upload_product_image", "product_management")

    // 処理...

    opLogger.Complete(ctx.Request().Context())  // 完了ログ
    return ctx.JSON(http.StatusOK, response)
}
```

### 3. main.go統合

```go
func main() {
    // ロガー初期化
    logger, err := logging.NewLogger(config.Observability)

    // ルーター作成時にロガーを渡す
    r := router.NewRouter(logger)

    // アプリケーション開始をログ出力
    logger.LogApplication(ctx, logging.ApplicationOperation{
        Name:     "application_startup",
        Category: "system",
        Success:  true,
        Stage:    "initialization",
    })
}
```

## 📊 ログクエリ例

### Grafana Loki でのクエリ例

#### 1. エラーログの検索

```logql
{service_name="aws-observability-ecommerce", level="error"}
```

#### 2. 特定の操作の追跡

```logql
{service_name="aws-observability-ecommerce", log_type="application"}
| json
| operation_name="upload_product_image"
```

#### 3. レスポンス時間が長いリクエスト

```logql
{service_name="aws-observability-ecommerce", log_type="request"}
| json
| duration_ms > 1000
```

#### 4. 特定のトレースIDでの追跡

```logql
{service_name="aws-observability-ecommerce"}
| json
| trace_id="4bf92f3577b34da6a3ce929d0e0e4736"
```

## 🎯 Phase 4 完了

Phase 4の特化ログ実装が完全に完了しました！

### ✅ 完了した機能

- [x] リクエストログの詳細構造とミドルウェア統合
- [x] エラーログの自動分類とフィンガープリント生成
- [x] アプリケーションログのビルダーパターン実装
- [x] ログヘルパーによる便利機能
- [x] 実際のハンドラーでの実用例
- [x] ミドルウェアと統合ハンドラーの完全統合

### 🚀 Phase 5 に進む準備完了

次はインフラ統合（Docker Compose更新、OTel Collector起動、Loki/Grafana設定）に進めます。
