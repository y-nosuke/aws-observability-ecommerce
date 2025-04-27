# 1. Week 3 - Day 1: 商品一覧APIの実装

## 1.1. 目次

- [1. Week 3 - Day 1: 商品一覧APIの実装](#1-week-3---day-1-商品一覧apiの実装)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. データベース接続の設定](#141-データベース接続の設定)
    - [1.4.2. 商品リポジトリの実装](#142-商品リポジトリの実装)
    - [1.4.3. サービス層の実装](#143-サービス層の実装)
    - [1.4.4. プロダクトハンドラーの実装](#144-プロダクトハンドラーの実装)
    - [1.4.5. メイン関数の更新](#145-メイン関数の更新)
    - [APIのテスト](#apiのテスト)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. リポジトリパターンについて](#161-リポジトリパターンについて)
    - [1.6.2. ページネーションの実装方法](#162-ページネーションの実装方法)
    - [1.6.3. エラーハンドリングの設計](#163-エラーハンドリングの設計)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. SQLBoilerを使用したORMの利点](#171-sqlboilerを使用したormの利点)
    - [1.7.2. 一般的なREST API設計のベストプラクティス](#172-一般的なrest-api設計のベストプラクティス)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. データベース接続エラー](#181-データベース接続エラー)
    - [1.8.2. ページネーション計算の誤り](#182-ページネーション計算の誤り)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)

## 1.2. 【要点】

- **リポジトリパターンを活用した商品データアクセス層の実装**：データベース操作を抽象化し、ビジネスロジックから分離
- **ページネーション機能を持つRESTful API実装**：限られたリソースを効率的に取得するためのページングシステム
- **カテゴリーによるフィルタリング機能の実装**：ユーザーが必要な商品のみを表示できる機能
- **適切なエラーハンドリングとレスポンス形式の標準化**：一貫性のあるAPI応答形式の確立
- **AWS CloudWatchとの連携によるAPI監視基盤の構築**：オブザーバビリティの基本となるログ記録の設定

## 1.3. 【準備】

### 1.3.1. チェックリスト

- [ ] Goの開発環境（Go 1.18以上）
- [ ] DockerとDocker Compose（開発環境の統一化のため）
- [ ] MySQLデータベース（ローカル環境またはDocker）
- [ ] AWS CLI（設定済み）
- [ ] Git（バージョン管理）
- [ ] Postman/curlなどのAPIテストツール
- [ ] プロジェクトのクローン済みリポジトリ

## 1.4. 【手順】

### 1.4.1. データベース接続の設定

まず、データベース接続設定を行います。

1. 設定ファイルを作成します。

    ```bash
    mkdir -p backend-api/internal/db/config
    touch backend-api/internal/db/config/database.go
    ```

    次にデータベース接続設定ファイルの内容を作成します:

    ```go
    // backend-api/internal/db/config/database.go
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
    ```

2. `config.go` の `Config` 構造体にデータベース設定を追加します。

    ```bash
    # config.goを編集します
    ```

    既存のファイル `backend-api/internal/config/config.go` を編集し、`Config` 構造体に以下を追加します：

    ```go
    // Config はアプリケーション全体の設定を格納する構造体
    type Config struct {
      App struct {
        Name        string
        Version     string
        Environment string
      }
      Server struct {
        Port int
      }
      // 追加: データベース設定
      Database struct {
        Host     string
        Port     int
        User     string
        Password string
        Name     string
      }
    }
    ```

    `var` セクションにも以下を追加します：

    ```go
    var (
      config   Config
      App      = &config.App
      Server   = &config.Server
      Database = &config.Database  // 追加
    )
    ```

    `Load()` 関数にも以下を追加します：

    ```go
    func Load() error {

      ・・・

      viper.SetDefault("database.host", "localhost")
      viper.SetDefault("database.port", 3306)
      viper.SetDefault("database.user", "root")
      viper.SetDefault("database.password", "password")
      viper.SetDefault("database.name", "ecommerce")

      ・・・

      if err := viper.BindEnv("database.host", "DB_HOST"); err != nil {
        return err
      }
      if err := viper.BindEnv("database.port", "DB_PORT"); err != nil {
        return err
      }
      if err := viper.BindEnv("database.name", "DB_NAME"); err != nil {
        return err
      }
      if err := viper.BindEnv("database.user", "DB_USER"); err != nil {
        return err
      }
      if err := viper.BindEnv("database.password", "DB_PASSWORD"); err != nil {
        return err
      }

      ・・・
    }
    ```

### 1.4.2. 商品リポジトリの実装

1. 既存の商品リポジトリ (`product_repository.go`) に実際のデータベース操作を実装します。

    既存のファイル `backend-api/internal/repository/product/product_repository.go` に実際の実装を追加します。今回は `FindAll` メソッドを中心に実装し、ページネーションとカテゴリーフィルタリングに対応させます：

    ```go
    // FindAll は商品一覧を取得します（ページネーション対応）
    func (r *ProductRepository) FindAll(ctx context.Context, limit, offset int) ([]*models.Product, error) {
      // クエリモディファイアの準備
      mods := []qm.QueryMod{
        qm.Limit(limit),
        qm.Offset(offset),
        qm.OrderBy("created_at DESC"),
        // 商品とカテゴリーを結合して取得
        qm.Load("Category"),
      }

      // SQLBoilerを使用して商品を取得
      products, err := models.Products(mods...).All(ctx, r.DB())
      if err != nil {
        return nil, fmt.Errorf("failed to fetch products: %w", err)
      }

      return products, nil
    }

    // FindByCategory は指定カテゴリーの商品一覧を取得します（ページネーション対応）
    func (r *ProductRepository) FindByCategory(ctx context.Context, categoryID int, limit, offset int) ([]*models.Product, error) {
      // クエリモディファイアの準備
      mods := []qm.QueryMod{
        qm.Where("category_id = ?", categoryID),
        qm.Limit(limit),
        qm.Offset(offset),
        qm.OrderBy("created_at DESC"),
        // 商品とカテゴリーを結合して取得
        qm.Load("Category"),
      }

      // SQLBoilerを使用して商品を取得
      products, err := models.Products(mods...).All(ctx, r.DB())
      if err != nil {
        return nil, fmt.Errorf("failed to fetch products by category: %w", err)
      }

      return products, nil
    }

    // Create は新しい商品を作成します
    func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
      return product.Insert(ctx, r.DB(), boil.Infer())
    }

    ・・・

    // Count は商品の総数を取得します
    func (r *ProductRepository) Count(ctx context.Context) (int64, error) {
      return models.Products().Count(ctx, r.DB())
    }

    // CountByCategory は指定カテゴリーの商品数を取得します
    func (r *ProductRepository) CountByCategory(ctx context.Context, categoryID int) (int64, error) {
      mods := []qm.QueryMod{
        qm.Where("category_id = ?", categoryID),
      }
      return models.Products(mods...).Count(ctx, r.DB())
    }
    ```

### 1.4.3. サービス層の実装

商品サービス層を実装してリポジトリとハンドラーの間の橋渡しを行います。

1. サービス実装を強化します。

    既存のファイル `backend-api/internal/service/product_service.go` に商品一覧を取得するためのサービスメソッドを追加します：

    ```go
    // ProductListResult は商品一覧の取得結果を表す構造体
    type ProductListResult struct {
      Items      []*models.Product
      Total      int64
      Page       int
      PageSize   int
      TotalPages int
    }

    // GetProducts は商品一覧を取得します（ページネーション対応）
    func (s *ProductService) GetProducts(ctx context.Context, page, pageSize int) (*ProductListResult, error) {
      if page < 1 {
        page = 1
      }
      if pageSize < 1 {
        pageSize = 20
      }
      if pageSize > 100 {
        pageSize = 100 // 最大制限
      }

      // オフセットを計算
      offset := (page - 1) * pageSize

      // 商品の総数を取得
      total, err := s.productRepo.Count(ctx)
      if err != nil {
        return nil, fmt.Errorf("failed to count products: %w", err)
      }

      // 総ページ数を計算
      totalPages := int(total) / pageSize
      if int(total)%pageSize > 0 {
        totalPages++
      }

      // 商品一覧を取得
      products, err := s.productRepo.FindAll(ctx, pageSize, offset)
      if err != nil {
        return nil, fmt.Errorf("failed to get products: %w", err)
      }

      return &ProductListResult{
        Items:      products,
        Total:      total,
        Page:       page,
        PageSize:   pageSize,
        TotalPages: totalPages,
      }, nil
    }

    // GetProductsByCategory はカテゴリー別の商品一覧を取得します
    func (s *ProductService) GetProductsByCategory(ctx context.Context, categoryID int, page, pageSize int) (*ProductListResult, error) {
      if page < 1 {
        page = 1
      }
      if pageSize < 1 {
        pageSize = 20
      }
      if pageSize > 100 {
        pageSize = 100 // 最大制限
      }

      // オフセットを計算
      offset := (page - 1) * pageSize

      // カテゴリー別商品の総数を取得
      total, err := s.productRepo.CountByCategory(ctx, categoryID)
      if err != nil {
        return nil, fmt.Errorf("failed to count products by category: %w", err)
      }

      // 総ページ数を計算
      totalPages := int(total) / pageSize
      if int(total)%pageSize > 0 {
        totalPages++
      }

      // カテゴリー別商品一覧を取得
      products, err := s.productRepo.FindByCategory(ctx, categoryID, pageSize, offset)
      if err != nil {
        return nil, fmt.Errorf("failed to get products by category: %w", err)
      }

      return &ProductListResult{
        Items:      products,
        Total:      total,
        Page:       page,
        PageSize:   pageSize,
        TotalPages: totalPages,
      }, nil
    }
    ```

### 1.4.4. プロダクトハンドラーの実装

APIエンドポイントを処理するハンドラーを実装します。
既存のファイル `backend-api/internal/api/handlers/product.go` を更新し、サービス層を利用するよう実装します：

```go
package handlers

import (
 "net/http"

 "github.com/labstack/echo/v4"

 "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
 "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/config"
 "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository/product"
 "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/service"
)

// ProductHandler はoapi-codegenで生成されたサーバーインターフェースを実装する構造体
type ProductHandler struct {
 productService *service.ProductService
}

// NewProductHandler はProductHandlerのインスタンスを生成する
func NewProductHandler() *ProductHandler {
 // リポジトリの初期化
 productRepo := product.New(config.DB)

 // サービスの初期化
 productService := service.NewProductService(config.DB, productRepo)

 return &ProductHandler{
  productService: productService,
 }
}

// ListProducts は商品一覧を取得する
func (h *ProductHandler) ListProducts(ctx echo.Context, params openapi.ListProductsParams) error {
 // パラメータの取得
 page := 1
 if params.Page != nil {
  page = *params.Page
 }

 pageSize := 20
 if params.PageSize != nil {
  pageSize = *params.PageSize
 }

 var result *service.ProductListResult
 var err error

 // カテゴリーIDでフィルタリングするかどうかを判定
 if params.CategoryId != nil {
  categoryID := int(*params.CategoryId)
  result, err = h.productService.GetProductsByCategory(ctx.Request().Context(), categoryID, page, pageSize)
 } else {
  result, err = h.productService.GetProducts(ctx.Request().Context(), page, pageSize)
 }

 if err != nil {
  // エラーハンドリング
  errorResponse := openapi.ErrorResponse{
   Code:    "internal_server_error",
   Message: "Failed to fetch products",
   Details: &map[string]interface{}{
    "error": err.Error(),
   },
  }
  return ctx.JSON(http.StatusInternalServerError, errorResponse)
 }

 // レスポンスの構築
 items := make([]openapi.Product, 0, len(result.Items))
 for _, p := range result.Items {
  // 在庫状態の取得
  inStock := false
  if p.R != nil && p.R.Inventories != nil && len(p.R.Inventories) > 0 {
   inStock = p.R.Inventories[0].Quantity > 0
  }

  // カテゴリー名の取得
  var categoryName *string
  if p.R != nil && p.R.Category != nil {
   categoryName = &p.R.Category.Name
  }

  price, _ := p.Price.Float64()
  items = append(items, openapi.Product{
   Id:           int64(p.ID),
   Name:         p.Name,
   Description:  stringPtr(p.Description.String),
   Price:        float32(price),
   ImageUrl:     stringPtr(p.ImageURL.String),
   InStock:      &inStock,
   CategoryId:   int64(p.CategoryID),
   CategoryName: categoryName,
   CreatedAt:    &p.CreatedAt,
   UpdatedAt:    &p.UpdatedAt,
  })
 }

 response := openapi.ProductList{
  Items:      items,
  Total:      int(result.Total),
  Page:       result.Page,
  PageSize:   result.PageSize,
  TotalPages: result.TotalPages,
 }

 return ctx.JSON(http.StatusOK, response)
}

// GetProductById は指定されたIDの商品を取得する
func (h *ProductHandler) GetProductById(ctx echo.Context, id int64) error {
 // 実装は次回に行います
 if id != 1 {
  errorResponse := openapi.ErrorResponse{
   Code:    "product_not_found",
   Message: "Product not found",
  }
  return ctx.JSON(http.StatusNotFound, errorResponse)
 }

 response := openapi.Product{
  Id:           1,
  Name:         "サンプル商品1",
  Description:  stringPtr("サンプル商品の説明文です。"),
  Price:        1000,
  ImageUrl:     stringPtr("https://example.com/image1.jpg"),
  InStock:      boolPtr(true),
  CategoryId:   1,
  CategoryName: stringPtr("サンプルカテゴリー"),
 }
 return ctx.JSON(http.StatusOK, response)
}

// 以下のメソッドは既存のままとします
// ...
```

### 1.4.5. メイン関数の更新

アプリケーションのエントリーポイントを更新して、データベース接続を初期化するようにします。

既存のファイル `backend-api/main.go` を更新します：

```go
package main

import (
 "context"
 "errors"
 "fmt"
 "log"
 "net/http"
 "os"
 "os/signal"
 "time"

 "github.com/labstack/echo/v4"
 "github.com/labstack/echo/v4/middleware"

 "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/handlers"
 "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/config"
 dbconfig "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/config" // 追加
)

func main() {
 // 設定をロード
 if err := config.Load(); err != nil {
  log.Printf("Failed to load configuration: %v\n", err)
  os.Exit(1)
 }

 // データベース接続の初期化 - 追加
 if err := dbconfig.InitDatabase(); err != nil {
  log.Printf("Failed to initialize database: %v\n", err)
  os.Exit(1)
 }
 // defer でアプリケーション終了時にデータベース接続をクローズ - 追加
 defer func() {
  if err := dbconfig.CloseDatabase(); err != nil {
   log.Printf("Failed to close database: %v\n", err)
  }
 }()

 // Echoインスタンスを作成
 e := echo.New()
 e.HideBanner = true
 e.HidePort = true

 // ミドルウェアの設定
 e.Use(middleware.Recover())
 e.Use(middleware.RequestID())
 e.Use(middleware.Logger()) // 標準のロガーミドルウェアを使用
 e.Use(middleware.CORS())

 // APIグループ
 api := e.Group("/api")

 if err := handlers.RegisterHandlers(api); err != nil {
  log.Fatalf("Failed to register handlers: %v", err)
 }

 e.Static("/swagger", "static/swagger-ui")
 e.File("/swagger", "static/swagger-ui/index.html")
 e.File("/openapi.yaml", "openapi.yaml")

 // コンテキストの初期化（シグナルハンドリング）
 ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
 defer stop()

 // サーバーを起動
 go func() {
  address := fmt.Sprintf(":%d", config.Server.Port)
  if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
   log.Printf("Failed to start server: %v\n", err)
   os.Exit(1)
  }
 }()

 // シグナルを待機
 <-ctx.Done()
 log.Println("Shutdown signal received, gracefully shutting down...")

 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 defer cancel()

 if err := e.Shutdown(ctx); err != nil {
  log.Printf("Failed to shutdown server gracefully: %v\n", err)
 } else {
  log.Printf("Server shutdown gracefully")
 }
}
```

### APIのテスト

コーディングが完了したら、APIをテストします。テスト用のスクリプトを作成します。

```bash
mkdir -p backend-api/scripts
touch backend-api/scripts/test_api.sh
chmod +x backend-api/scripts/test_api.sh
```

```bash
#!/bin/bash
# backend-api/scripts/test_api.sh

# APIのベースURL
BASE_URL="http://backend-api.localhost/api"

# 色の定義
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# ヘルスチェック
echo -e "${GREEN}Testing health endpoint:${NC}"
curl -s "${BASE_URL}/health" | jq
echo ""

# 商品一覧の取得（デフォルトページ）
echo -e "${GREEN}Testing products list (default page):${NC}"
curl -s "${BASE_URL}/products" | jq
echo ""

# 商品一覧の取得（ページサイズ指定）
echo -e "${GREEN}Testing products list with pageSize=5:${NC}"
curl -s "${BASE_URL}/products?pageSize=5" | jq
echo ""

# 商品一覧の取得（ページネーション）
echo -e "${GREEN}Testing products list page 2:${NC}"
curl -s "${BASE_URL}/products?page=2&pageSize=3" | jq
echo ""

# カテゴリー別商品一覧の取得
echo -e "${GREEN}Testing products by category (ID=1):${NC}"
curl -s "${BASE_URL}/products?categoryId=1" | jq
echo ""

# 単一商品の取得
echo -e "${GREEN}Testing get product by ID:${NC}"
curl -s "${BASE_URL}/products/1" | jq
echo ""

# カテゴリー一覧の取得
echo -e "${GREEN}Testing categories list:${NC}"
curl -s "${BASE_URL}/categories" | jq
echo ""

# 存在しない商品の取得（エラーレスポンスのテスト）
echo -e "${GREEN}Testing error response (product not found):${NC}"
curl -s "${BASE_URL}/products/999" | jq
echo ""

echo -e "${GREEN}All tests completed!${NC}"
```

別のターミナルでテストスクリプトを実行します：

```bash
./backend-api/scripts/test_api.sh
```

## 1.5. 【確認ポイント】

- [ ] データベースに正常に接続できている
- [ ] 商品一覧APIがページネーションパラメータ（`page`と`pageSize`）に応じて正しい結果を返す
- [ ] カテゴリーIDでフィルタリングできる
- [ ] レスポンスが正しいJSON形式になっている
- [ ] エラー時に適切なエラーレスポンスを返す
- [ ] ログがCloudWatchに正しく記録されている（AWS環境の場合）
- [ ] 商品の総数と総ページ数が正しく計算されている

## 1.6. 【詳細解説】

### 1.6.1. リポジトリパターンについて

リポジトリパターンは、データアクセスロジックを分離し、データソースの詳細をビジネスロジックから隠蔽するためのパターンです。今回の実装では以下のメリットがあります：

1. **テスト容易性**: モックリポジトリを使用したユニットテストが容易になります
2. **柔軟性**: データベースのタイプを変更しても、上位層に影響を与えません
3. **関心の分離**: データアクセスロジックとビジネスロジックを分離できます
4. **再利用性**: 同じデータアクセスロジックを複数のサービスで再利用できます

### 1.6.2. ページネーションの実装方法

REST APIでのページネーションには主に3つのアプローチがあります：

1. **オフセットベースのページネーション**: 今回採用したアプローチで、`page`と`pageSize`パラメータを使用します。SQLの`LIMIT`と`OFFSET`に対応し、実装が容易ですが、大量データでは効率が低下します。

2. **カーソルベースのページネーション**: 最後に取得したアイテムのIDなどをカーソルとして使用します。大量データに適していますが、実装が複雑です。

3. **キーセットページネーション**: 特定のフィールド値に基づいて次のページを取得します。効率的ですが、複雑なクエリが必要です。

今回は簡単さと一般的な使いやすさからオフセットベースのページネーションを選択しました。

### 1.6.3. エラーハンドリングの設計

エラーハンドリングでは以下の点に注意しています：

1. **一貫性のあるエラーレスポンス形式**: すべてのエラーレスポンスが共通の構造になっており、クライアントが解析しやすくなっています。

2. **エラーコード**: エラータイプを識別するためのコードを含みます。

3. **詳細情報**: デバッグに役立つ追加情報を提供します。

4. **適切なHTTPステータスコード**: エラータイプに適したHTTPステータスコードを使用します。

## 1.7. 【補足情報】

### 1.7.1. SQLBoilerを使用したORMの利点

今回のアプリケーションでは、SQLBoilerというコード生成ベースのORMを使用しています。これには以下の利点があります：

1. **型安全性**: コンパイル時に多くのエラーを検出できます。
2. **パフォーマンス**: 実行時のリフレクションを最小限に抑えることができます。
3. **柔軟性**: 複雑なクエリの構築が容易です。
4. **データベーススキーマとの同期**: スキーマから直接コードを生成するため、不整合が発生しにくいです。

### 1.7.2. 一般的なREST API設計のベストプラクティス

1. **URI設計**: リソース名は複数形を使用し、階層構造を明確にします。例えば、`/products`、`/categories/{id}/products`。

2. **HTTPメソッドの適切な使用**:
   - GET: リソースの取得
   - POST: 新しいリソースの作成
   - PUT: リソースの完全な更新
   - PATCH: リソースの部分的な更新
   - DELETE: リソースの削除

3. **クエリパラメータの一貫性**: ページネーションパラメータ名（`page`、`pageSize`）やフィルターパラメータ名を一貫して使用します。

4. **レスポンスの一貫性**: すべてのエンドポイントで一貫したレスポンス形式を使用します。

## 1.8. 【よくある問題と解決法】

### 1.8.1. データベース接続エラー

**問題**: アプリケーション起動時にデータベース接続エラーが発生する。

**解決策**:

- 環境変数が正しく設定されているか確認
- ホスト名やポート番号が正しいか確認
- データベースユーザーのパーミッションを確認
- ファイアウォール設定を確認

```bash
# データベース接続のトラブルシューティング
mysql -u$DB_USER -p$DB_PASSWORD -h$DB_HOST -P$DB_PORT -e "SELECT 1"
```

### 1.8.2. ページネーション計算の誤り

**問題**: 総ページ数の計算が間違っている、または最後のページが正しく表示されない。

**解決策**:

- 総ページ数の計算ロジックを確認（切り上げが必要）

```go
// 正しい総ページ数の計算
totalPages := (total + pageSize - 1) / pageSize
// または
totalPages := total / pageSize
if total % pageSize > 0 {
    totalPages++
}
```

## 1.9. 【今日の重要なポイント】

1. **レイヤードアーキテクチャによる関心の分離**：ハンドラー → サービス → リポジトリという明確な階層が、コードの保守性と拡張性を向上させます。

2. **ページネーション実装のベストプラクティス**：総数の取得とページ計算という2ステップアプローチにより、クライアントに必要なすべての情報を提供します。

3. **レスポンス形式の標準化**：一貫性のあるレスポンス形式により、クライアント側での処理が容易になります。

## 1.10. 【次回の準備】

1. **単一商品APIの実装**：商品詳細を取得するAPIエンドポイントの実装方法を確認してください。

2. **リレーショナルデータの効率的な取得**：商品と関連する在庫やカテゴリー情報を効率的に取得する方法を考えてください。

3. **エラーハンドリングの強化**：より詳細なエラー情報を提供する方法を検討してください。
