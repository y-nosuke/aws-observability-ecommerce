# パッケージレベルログの使用例

このドキュメントでは、新しいパッケージレベルのログ機能の使い方を例を用いて説明します。

## 🎯 基本的な使い方

### 1. 最もシンプルな使用方法

```go
package main

import (
    "context"
    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
)

func someBusinessLogic(ctx context.Context, userID int64) error {
    // ✅ 理想的な使い方！
    logging.Info(ctx, "処理を開始", "user_id", userID, "operation", "login")
    
    // 何らかの処理...
    
    if err != nil {
        // ✅ エラーも簡単！
        logging.WithError(ctx, "ログイン処理に失敗", err, "user_id", userID)
        return err
    }
    
    logging.Info(ctx, "ログイン成功", "user_id", userID)
    return nil
}
```

### 2. フォーマット付きログ

```go
func processOrder(ctx context.Context, orderID string, userID int64) {
    // ✅ printf風のフォーマット
    logging.InfoF(ctx, "ユーザー %d の注文 %s を処理開始", userID, orderID)
    
    // ✅ エラー時もフォーマット可能
    if err != nil {
        logging.ErrorF(ctx, "注文 %s の処理で失敗: %v", orderID, err)
    }
}
```

### 3. 操作の自動追跡

```go
func uploadFile(ctx context.Context, filename string) error {
    // ✅ 操作の開始と完了を自動追跡
    completeOp := logging.StartOperation(ctx, "file_upload",
        "filename", filename,
        "layer", "service")
    
    // 実際の処理
    err := doUpload(filename)
    
    if err != nil {
        completeOp(false, "error_type", "upload_failed")
        return err
    }
    
    completeOp(true, "file_size", fileSize)
    return nil
}
```

## 🔄 移行前後の比較例

### ユースケース層の移行例

```go
// ====== 移行前（複雑） ======
type GetProductImageUseCase struct {
    imageStorage service.ImageStorage
    logHelper    *logging.LogHelper  // ❌ 複雑なインジェクション
}

func NewGetProductImageUseCase(
    imageStorage service.ImageStorage,
    logger logging.Logger,  // ❌ 引数が増える
) *GetProductImageUseCase {
    return &GetProductImageUseCase{
        imageStorage: imageStorage,
        logHelper:    logging.NewLogHelper(logger),  // ❌ ラッパー作成
    }
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

// ====== 移行後（シンプル） ======
type GetProductImageUseCase struct {
    imageStorage service.ImageStorage
    // ✅ ログ関連の依存を削除！
}

func NewGetProductImageUseCase(
    imageStorage service.ImageStorage,
    // ✅ logger引数を削除！
) *GetProductImageUseCase {
    return &GetProductImageUseCase{
        imageStorage: imageStorage,
    }
}

func (u *GetProductImageUseCase) Execute(ctx context.Context, productID int64, size string) (*dto.GetImageResponse, error) {
    // ✅ 操作の自動追跡
    completeOp := logging.StartOperation(ctx, "get_product_image",
        "product_id", productID,
        "requested_size", size,
        "layer", "usecase")

    imageData, contentType, err := u.imageStorage.GetImageData(ctx, productID, size)
    if err != nil {
        // ✅ シンプルなエラーログ
        logging.WithError(ctx, "画像データの取得に失敗", err,
            "product_id", productID,
            "requested_size", size,
            "layer", "usecase")
        
        completeOp(false, "error_type", "storage_failure")
        return nil, fmt.Errorf("failed to get image data: %w", err)
    }

    // ✅ シンプルな成功ログ
    logging.Info(ctx, "画像データを正常に取得",
        "product_id", productID,
        "content_type", contentType,
        "image_size_bytes", len(imageData),
        "layer", "usecase")

    completeOp(true,
        "content_type", contentType,
        "image_size_bytes", len(imageData))

    return dto.NewGetImageResponse(productID, imageData, contentType), nil
}
```

### ハンドラー層の移行例

```go
// ====== 移行前 ======
type ProductHandler struct {
    getImageUseCase *usecase.GetProductImageUseCase
    logger          logging.Logger  // ❌ 依存注入が必要
}

func (h *ProductHandler) GetProductImage(c echo.Context) error {
    // ❌ 複雑な構造体作成
    fields := []logging.Field{
        {Key: "handler", Value: "GetProductImage"},
        {Key: "product_id", Value: productID},
    }
    h.logger.Info(ctx, "リクエスト処理開始", fields...)
    
    // ... 処理 ...
    
    if err != nil {
        h.logger.Error(ctx, "処理に失敗", err, fields...)
        return err
    }
    
    h.logger.Info(ctx, "処理完了", fields...)
    return c.JSON(http.StatusOK, response)
}

// ====== 移行後 ======
type ProductHandler struct {
    getImageUseCase *usecase.GetProductImageUseCase
    // ✅ logger依存を削除！
}

func (h *ProductHandler) GetProductImage(c echo.Context) error {
    // ✅ シンプルなログ
    logging.Info(ctx, "画像取得リクエスト開始",
        "handler", "GetProductImage",
        "product_id", productID,
        "layer", "handler")
    
    // ... 処理 ...
    
    if err != nil {
        // ✅ シンプルなエラーログ
        logging.WithError(ctx, "画像取得処理に失敗", err,
            "product_id", productID,
            "layer", "handler")
        return err
    }
    
    // ✅ シンプルな成功ログ
    logging.Info(ctx, "画像取得完了",
        "product_id", productID,
        "response_size", len(response),
        "layer", "handler")
    
    return c.JSON(http.StatusOK, response)
}
```

## 🚀 高度な使用例

### 1. ビジネスイベントの記録

```go
func processPayment(ctx context.Context, orderID string, amount int64) error {
    // 支払い処理開始
    logging.Info(ctx, "支払い処理開始",
        "order_id", orderID,
        "amount", amount,
        "layer", "payment_service")
    
    // ... 支払い処理 ...
    
    // ✅ ビジネスイベントとして記録
    logging.LogBusinessEvent(ctx, "payment_completed", "order", orderID,
        "amount", amount,
        "payment_method", "credit_card",
        "processor", "stripe")
    
    return nil
}
```

### 2. パフォーマンス監視

```go
func searchProducts(ctx context.Context, query string) ([]Product, error) {
    start := time.Now()
    
    // データベースクエリ実行
    products, err := db.Search(query)
    dbTime := time.Since(start)
    
    if err != nil {
        logging.WithError(ctx, "商品検索でエラー", err,
            "query", query,
            "db_time_ms", dbTime.Milliseconds())
        return nil, err
    }
    
    // ✅ パフォーマンス情報を含むログ
    logging.Info(ctx, "商品検索完了",
        "query", query,
        "results_count", len(products),
        "db_time_ms", dbTime.Milliseconds(),
        "total_time_ms", time.Since(start).Milliseconds(),
        "cache_hit", false)
    
    return products, nil
}
```

### 3. デバッグ用の簡単ログ

```go
func debugFunction(ctx context.Context) {
    // ✅ 一時的なデバッグログ（後で削除予定）
    logging.QuickInfo(ctx, "デバッグポイント1")
    
    // ... 何らかの処理 ...
    
    if err != nil {
        // ✅ 簡単エラーログ
        logging.QuickError(ctx, "デバッグ中にエラー", err)
    }
}
```

## 📝 推奨事項

### 1. キー名の一貫性

同じ概念には同じキー名を使用:
```go
// ✅ Good: 一貫性のあるキー名
logging.Info(ctx, "ユーザー作成", "user_id", userID)
logging.Info(ctx, "ユーザー更新", "user_id", userID)
logging.Info(ctx, "ユーザー削除", "user_id", userID)

// ❌ Bad: バラバラなキー名
logging.Info(ctx, "ユーザー作成", "user_id", userID)
logging.Info(ctx, "ユーザー更新", "userId", userID)
logging.Info(ctx, "ユーザー削除", "id", userID)
```

### 2. レイヤー識別

アーキテクチャの層を明確に:
```go
// ✅ 各層で明確に識別
logging.Info(ctx, "処理開始", "layer", "handler")
logging.Info(ctx, "処理開始", "layer", "usecase") 
logging.Info(ctx, "処理開始", "layer", "repository")
```

### 3. エラーログの充実

```go
// ✅ エラー時は十分な情報を記録
logging.WithError(ctx, "データベース接続失敗", err,
    "database", "products",
    "operation", "select",
    "retry_count", retryCount,
    "max_retries", maxRetries)
```

これで、`logging.Info(ctx, msg, ...)` の形で簡単にログが出力できるようになりました！
