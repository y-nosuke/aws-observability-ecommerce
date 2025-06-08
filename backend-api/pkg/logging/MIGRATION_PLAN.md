# ログシステム改善・移行計画

## 改善のメリット

### ✅ **改善前（現在）の問題**
- 複雑な階層構造 (`Logger` → `LogHelper` → `OperationLogger`)
- 独自構造体の理解が必要 (`Field`, `ApplicationOperation`, `RequestLogData`)
- 使い方が複雑 (`ContinueOperationInLayer`, `WithEntity`, `Complete/Fail`)
- 学習コストが高い

### ✅ **改善後のメリット**
- **シンプルな使い方**: `logger.Info(ctx, msg, ...)`
- **自動化**: 共通フィールド（トレースID、リクエストID等）は自動付与
- **柔軟性**: key-value形式で任意のフィールドを追加可能
- **後方互換性**: 既存の複雑な機能も必要に応じて維持可能

## 段階的移行計画

### 📋 **Phase 1: シンプルロガーの導入（1-2日）**

1. **SimpleLoggerの実装完了** ✅
   - `/pkg/logging/simple_logger.go` 作成済み
   - 基本機能（Info, Error, Warn, Debug）実装済み

2. **DIコンテナへの統合**
   ```go
   // di/provider/shared_provider.go に追加
   func ProvideSimpleLogger(cfg config.ObservabilityConfig) (*logging.SimpleLogger, error) {
       return logging.NewSimpleLogger(cfg)
   }
   ```

3. **新規開発での採用開始**
   - 新しいユースケースやハンドラーでSimpleLoggerを使用
   - 既存コードは触らない

### 📋 **Phase 2: 主要箇所の移行（3-5日）**

1. **ミドルウェアの移行**
   ```go
   // 移行前
   logger.LogRequest(c.Request().Context(), logData)
   
   // 移行後
   logger.LogHTTPRequest(ctx, method, path, status, duration,
       "request_size", requestSize,
       "response_size", responseSize,
       "user_agent", userAgent)
   ```

2. **主要ユースケースの移行**
   - `GetProductImageUseCase`などの移行
   - 段階的に1つずつ移行

3. **エラーハンドリングの統一**
   ```go
   // 移行前
   usecaseLogger.WithData("error_details", err.Error()).Fail(ctx, err)
   
   // 移行後
   logger.WithError(ctx, "操作に失敗", err, "operation", "get_product_image")
   ```

### 📋 **Phase 3: 完全移行（1週間）**

1. **既存コードの全面移行**
2. **不要なファイルの削除**
3. **テストの更新**
4. **ドキュメントの整備**

## 移行作業の実例

### 🔄 **ユースケース移行例**

```go
// ====== 移行前 ======
type GetProductImageUseCase struct {
    imageStorage service.ImageStorage
    logHelper    *logging.LogHelper  // ❌ 複雑
}

func (u *GetProductImageUseCase) Execute(ctx context.Context, productID int64, size string) (*dto.GetImageResponse, error) {
    // ❌ 複雑な初期化
    usecaseLogger := logging.ContinueOperationInLayer(ctx, u.logHelper, "usecase")
    usecaseLogger.WithEntity("product", fmt.Sprint(productID)).
        WithAction("retrieve", "usecase").
        WithData("requested_size", size)

    imageData, contentType, err := u.imageStorage.GetImageData(ctx, productID, size)
    if err != nil {
        // ❌ 複雑なエラーログ
        usecaseLogger.WithData("storage_error", "retrieval_failed").
            WithData("error_details", err.Error()).
            Fail(ctx, err)
        return nil, fmt.Errorf("failed to get image data: %w", err)
    }

    // ❌ 複雑な成功ログ
    usecaseLogger.WithData("content_type", contentType).
        WithData("image_size_bytes", len(imageData)).
        Complete(ctx)

    return dto.NewGetImageResponse(productID, imageData, contentType), nil
}

// ====== 移行後 ======
type GetProductImageUseCase struct {
    imageStorage service.ImageStorage
    logger       *logging.SimpleLogger  // ✅ シンプル
}

func (u *GetProductImageUseCase) Execute(ctx context.Context, productID int64, size string) (*dto.GetImageResponse, error) {
    start := time.Now()

    imageData, contentType, err := u.imageStorage.GetImageData(ctx, productID, size)
    if err != nil {
        // ✅ シンプルなエラーログ
        u.logger.WithError(ctx, "画像データの取得に失敗", err,
            "product_id", productID,
            "requested_size", size,
            "layer", "usecase")
        return nil, fmt.Errorf("failed to get image data: %w", err)
    }

    // ✅ シンプルな成功ログ
    u.logger.LogOperation(ctx, "get_product_image", time.Since(start), true,
        "product_id", productID,
        "content_type", contentType,
        "image_size_bytes", len(imageData),
        "layer", "usecase")

    return dto.NewGetImageResponse(productID, imageData, contentType), nil
}
```

### 🔄 **ミドルウェア移行例**

```go
// ====== 移行前 ======
func StructuredLoggingMiddleware(logger logging.Logger) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // ... リクエスト処理 ...
            
            // ❌ 複雑な構造体作成
            logData := logging.RequestLogData{
                Method:        c.Request().Method,
                Path:          c.Request().URL.Path,
                Query:         c.Request().URL.RawQuery,
                StatusCode:    c.Response().Status,
                RequestSize:   requestSize,
                ResponseSize:  resWrapper.size,
                Duration:      duration,
                UserAgent:     c.Request().UserAgent(),
                RemoteIP:      c.RealIP(),
                // ... 多数のフィールド
            }
            
            // ❌ 専用メソッド呼び出し
            logger.LogRequest(c.Request().Context(), logData)
            return err
        }
    }
}

// ====== 移行後 ======
func SimpleLoggingMiddleware(logger *logging.SimpleLogger) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            start := time.Now()
            
            // ... リクエスト処理 ...
            
            // ✅ シンプルなログ出力
            logger.LogHTTPRequest(ctx, 
                c.Request().Method, 
                c.Request().URL.Path, 
                c.Response().Status, 
                time.Since(start),
                "request_size", requestSize,
                "response_size", responseSize,
                "user_agent", c.Request().UserAgent(),
                "remote_ip", c.RealIP(),
                "query", c.Request().URL.RawQuery)
            
            return err
        }
    }
}
```

## 作業チェックリスト

### Phase 1 ✅
- [x] SimpleLogger実装
- [x] 使用例とドキュメント作成
- [ ] DIコンテナ統合
- [ ] 新規開発での採用開始

### Phase 2
- [ ] ミドルウェア移行
- [ ] 主要ユースケース移行
- [ ] エラーハンドリング統一

### Phase 3
- [ ] 全既存コード移行
- [ ] 不要ファイル削除
- [ ] テスト更新
- [ ] ドキュメント整備

## 🎯 **immediate next step**

まず、DIコンテナへの統合から始めることをお勧めします。これにより新規開発で即座にSimpleLoggerを使用できるようになります。
