# 1. Week 3 - Day 2: 商品詳細・カテゴリAPI実装

## 1.1. 目次

- [1. Week 3 - Day 2: 商品詳細・カテゴリAPI実装](#1-week-3---day-2-商品詳細カテゴリapi実装)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. 商品詳細取得APIの実装](#141-商品詳細取得apiの実装)
    - [1.4.2. カテゴリー一覧APIの実装](#142-カテゴリー一覧apiの実装)
    - [1.4.3. カテゴリー別商品一覧APIの実装](#143-カテゴリー別商品一覧apiの実装)
    - [1.4.4. リポジトリとサービスの拡張](#144-リポジトリとサービスの拡張)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. URLパラメータの使用方法とベストプラクティス](#161-urlパラメータの使用方法とベストプラクティス)
    - [1.6.2. リポジトリパターンの効果的な適用](#162-リポジトリパターンの効果的な適用)
    - [1.6.3. レスポンス形式の一貫性と統一性](#163-レスポンス形式の一貫性と統一性)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. N+1問題とEagerローディング](#171-n1問題とeagerローディング)
    - [1.7.2. OpenAPIスキーマとコード生成の連携](#172-openapiスキーマとコード生成の連携)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: パスパラメータが正しく取得できない](#181-問題1-パスパラメータが正しく取得できない)
    - [1.8.2. 問題2: 関連データの取得でN+1問題が発生する](#182-問題2-関連データの取得でn1問題が発生する)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)

## 1.2. 【要点】

- URLパラメータを使用した単一リソース取得の実装パターン
- カテゴリー関連APIの設計と実装による関連データの効率的な取得
- レスポンス形式の一貫性確保とエラーハンドリングの統一化
- SQLBoilerのEagerローディング活用によるN+1問題の回避
- OpenAPI仕様の詳細化とSwagger UIの活用

## 1.3. 【準備】

### 1.3.1. チェックリスト

- [ ] Week 3 - Day 1の実装が完了していること
- [ ] Docker環境が正常に稼働していること
- [ ] データベースが起動し、テストデータが投入されていること
- [ ] OpenAPI仕様ファイル（openapi.yaml）にアクセス可能であること
- [ ] API実装のための基本構造（ルーティング、ハンドラ、リポジトリ）が整備されていること
- [ ] sqlboilerによる生成モデルが存在していること

データベースへの接続とサーバーが正常に起動していることを確認します：

```bash
# サーバーの起動確認
cd backend-api
go run main.go

# 別のターミナルで
curl http://backend-api.localhost/api/health
```

## 1.4. 【手順】

### 1.4.1. 商品詳細取得APIの実装

まず、商品詳細を取得するAPIの実装から始めます。ProductRepositoryに商品詳細取得メソッドを実装します。

1. 商品詳細取得機能を実装するため、リポジトリの実装を確認します：

    ```bash
    cat backend-api/internal/repository/product/product_repository.go
    ```

    リポジトリに既に `FindByID` メソッドがあることを確認したら、このメソッドを使用するProductServiceを拡張します。

2. ProductServiceに商品詳細取得メソッドを追加します：

    `backend-api/internal/service/product_detail_service.go` ファイルに以下のコードを追加します：

    ```go
    package service

    import (
    "context"
    "fmt"
    "database/sql"

    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"
    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository/product"
    )

    // GetProductByID は指定されたIDの商品を取得します
    func (s *ProductService) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
    // リポジトリから商品を取得
    p, err := s.productRepo.FindByID(ctx, id)
    if err != nil {
      if err == sql.ErrNoRows {
      return nil, fmt.Errorf("product not found: %d", id)
      }
      return nil, fmt.Errorf("failed to get product: %w", err)
    }

    // Eager Loadingで関連データを取得
    if p.R == nil || p.R.Category == nil {
      // カテゴリーが読み込まれていない場合はカテゴリーも取得
      if err := p.L.LoadCategory(ctx, s.db, true, p, nil); err != nil {
      return nil, fmt.Errorf("failed to load category: %w", err)
      }
    }

    if p.R == nil || p.R.Inventories == nil {
      // 在庫情報が読み込まれていない場合は在庫も取得
      if err := p.L.LoadInventories(ctx, s.db, true, p, nil); err != nil {
      return nil, fmt.Errorf("failed to load inventory: %w", err)
      }
    }

    return p, nil
    }
    ```

3. 商品詳細取得ハンドラを更新します：

    `backend-api/internal/api/handlers/product.go` メソッドを以下のコードで置き換えます：

    ```go
    // GetProductById は指定されたIDの商品を取得する
    func (h *ProductHandler) GetProductById(ctx echo.Context, id int64) error {
    // IDの整合性チェック
    if id <= 0 {
      errorResponse := openapi.ErrorResponse{
      Code:    "invalid_parameter",
      Message: "Invalid product ID",
      Details: &map[string]interface{}{
        "id": "must be a positive integer",
      },
      }
      return ctx.JSON(http.StatusBadRequest, errorResponse)
    }

    // サービスから商品詳細を取得
    product, err := h.productService.GetProductByID(ctx.Request().Context(), int(id))
    if err != nil {
      // エラーの種類に応じて適切なレスポンスを返す
      if strings.Contains(err.Error(), "product not found") {
      errorResponse := openapi.ErrorResponse{
        Code:    "product_not_found",
        Message: "Product not found",
        Details: &map[string]interface{}{
        "id": id,
        },
      }
      return ctx.JSON(http.StatusNotFound, errorResponse)
      }

      // その他のエラー
      errorResponse := openapi.ErrorResponse{
      Code:    "internal_server_error",
      Message: "Failed to fetch product details",
      Details: &map[string]interface{}{
        "error": err.Error(),
      },
      }
      return ctx.JSON(http.StatusInternalServerError, errorResponse)
    }

    // 在庫状態の取得
    inStock := false
    if product.R != nil && product.R.Inventories != nil && len(product.R.Inventories) > 0 {
      inStock = product.R.Inventories[0].Quantity > 0
    }

    // カテゴリー名の取得
    var categoryName *string
    if product.R != nil && product.R.Category != nil {
      categoryName = &product.R.Category.Name
    }

    // 価格のパース
    price, _ := product.Price.Float64()

    // レスポンスの構築
    response := openapi.Product{
      Id:           int64(product.ID),
      Name:         product.Name,
      Description:  stringPtr(product.Description.String),
      Price:        float32(price),
      ImageUrl:     stringPtr(product.ImageURL.String),
      InStock:      &inStock,
      CategoryId:   int64(product.CategoryID),
      CategoryName: categoryName,
      CreatedAt:    &product.CreatedAt,
      UpdatedAt:    &product.UpdatedAt,
    }

    return ctx.JSON(http.StatusOK, response)
    }
    ```

必要に応じて、stringPtrなどのヘルパー関数がメソッドで使用できることを確認してください。

### 1.4.2. カテゴリー一覧APIの実装

次に、カテゴリー一覧を取得するAPIを実装します。

1. まず、カテゴリーリポジトリを実装します：

    ```bash
    mkdir -p backend-api/internal/repository/category
    touch backend-api/internal/repository/category/category_repository.go
    ```

    カテゴリーリポジトリに以下のコードを追加します：

    ```go
    package category

    import (
    "context"
    "database/sql"
    "fmt"

    "github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/sqlboiler/v4/queries/qm"

    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"
    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository"
    )

    // Repository はカテゴリーリポジトリのインターフェース
    type Repository interface {
    repository.Repository
    FindByID(ctx context.Context, id int) (*models.Category, error)
    FindAll(ctx context.Context) ([]*models.Category, error)
    Create(ctx context.Context, category *models.Category) error
    Update(ctx context.Context, category *models.Category) error
    Delete(ctx context.Context, id int) error
    GetProductCount(ctx context.Context, categoryID int) (int64, error)
    }

    // CategoryRepository はカテゴリーリポジトリの実装
    type CategoryRepository struct {
    repository.RepositoryBase
    }

    // New は新しいカテゴリーリポジトリを作成します
    func New(executor boil.ContextExecutor) Repository {
    return &CategoryRepository{
      RepositoryBase: repository.NewRepositoryBase(executor),
    }
    }

    // WithTx はトランザクションを設定したカテゴリーリポジトリを返します
    func (r *CategoryRepository) WithTx(tx *sql.Tx) repository.Repository {
    return &CategoryRepository{
      RepositoryBase: r.RepositoryBase.WithTx(tx),
    }
    }

    // FindByID は指定IDのカテゴリーを取得します
    func (r *CategoryRepository) FindByID(ctx context.Context, id int) (*models.Category, error) {
    return models.FindCategory(ctx, r.DB(), id)
    }

    // FindAll はカテゴリー一覧を取得します
    func (r *CategoryRepository) FindAll(ctx context.Context) ([]*models.Category, error) {
    // クエリモディファイアの準備
    mods := []qm.QueryMod{
      qm.OrderBy("name ASC"),
    }

    // SQLBoilerを使用してカテゴリーを取得
    categories, err := models.Categories(mods...).All(ctx, r.DB())
    if err != nil {
      return nil, fmt.Errorf("failed to fetch categories: %w", err)
    }

    return categories, nil
    }

    // Create は新しいカテゴリーを作成します
    func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
    return category.Insert(ctx, r.DB(), boil.Infer())
    }

    // Update はカテゴリー情報を更新します
    func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
    _, err := category.Update(ctx, r.DB(), boil.Infer())
    return err
    }

    // Delete はカテゴリーを削除します
    func (r *CategoryRepository) Delete(ctx context.Context, id int) error {
    category, err := r.FindByID(ctx, id)
    if err != nil {
      return err
    }
    if category == nil {
      return fmt.Errorf("category not found: %d", id)
    }
    _, err = category.Delete(ctx, r.DB())
    return err
    }

    // GetProductCount は指定カテゴリーの商品数を取得します
    func (r *CategoryRepository) GetProductCount(ctx context.Context, categoryID int) (int64, error) {
    mods := []qm.QueryMod{
      qm.Where("category_id = ?", categoryID),
    }
    return models.Products(mods...).Count(ctx, r.DB())
    }
    ```

2. 次に、カテゴリーサービスを実装します：

    ```bash
    touch backend-api/internal/service/category_service.go
    ```

    カテゴリーサービスに以下のコードを追加します：

    ```go
    package service

    import (
    "context"
    "database/sql"
    "fmt"

    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"
    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository/category"
    )

    // CategoryService はカテゴリーサービスの実装
    type CategoryService struct {
    db           *sql.DB
    categoryRepo category.Repository
    }

    // NewCategoryService は新しいカテゴリーサービスを作成します
    func NewCategoryService(db *sql.DB, categoryRepo category.Repository) *CategoryService {
    return &CategoryService{
      db:           db,
      categoryRepo: categoryRepo,
    }
    }

    // GetCategories はカテゴリー一覧を取得します
    func (s *CategoryService) GetCategories(ctx context.Context) ([]*models.Category, error) {
    // カテゴリー一覧を取得
    categories, err := s.categoryRepo.FindAll(ctx)
    if err != nil {
      return nil, fmt.Errorf("failed to get categories: %w", err)
    }

    // 各カテゴリーの商品数を取得
    for _, cat := range categories {
      count, err := s.categoryRepo.GetProductCount(ctx, cat.ID)
      if err != nil {
      // エラーが発生しても続行
      continue
      }
      // 商品数を設定（カスタムプロパティとして追加）
      cat.AddPreload("product_count", count)
    }

    return categories, nil
    }

    // GetCategoryByID は指定されたIDのカテゴリーを取得します
    func (s *CategoryService) GetCategoryByID(ctx context.Context, id int) (*models.Category, error) {
    // カテゴリーを取得
    category, err := s.categoryRepo.FindByID(ctx, id)
    if err != nil {
      if err == sql.ErrNoRows {
      return nil, fmt.Errorf("category not found: %d", id)
      }
      return nil, fmt.Errorf("failed to get category: %w", err)
    }

    // 商品数を取得
    count, err := s.categoryRepo.GetProductCount(ctx, category.ID)
    if err == nil {
      // 商品数を設定（カスタムプロパティとして追加）
      category.AddPreload("product_count", count)
    }

    return category, nil
    }
    ```

3. ハンドラに新しいカテゴリーハンドラを追加します：

    ```bash
    touch backend-api/internal/api/handlers/category.go
    ```

    カテゴリーハンドラに以下のコードを追加します：

    ```go
    package handlers

    import (
    "net/http"

    "github.com/labstack/echo/v4"

    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/config"
    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository/category"
    "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/service"
    )

    // CategoryHandler はカテゴリーAPIのハンドラ実装
    type CategoryHandler struct {
    categoryService *service.CategoryService
    }

    // NewCategoryHandler は新しいCategoryHandlerを作成します
    func NewCategoryHandler() *CategoryHandler {
    // リポジトリの初期化
    categoryRepo := category.New(config.DB)

    // サービスの初期化
    categoryService := service.NewCategoryService(config.DB, categoryRepo)

    return &CategoryHandler{
      categoryService: categoryService,
    }
    }

    // ListCategories はカテゴリー一覧を取得します
    func (h *CategoryHandler) ListCategories(ctx echo.Context) error {
    // サービスからカテゴリー一覧を取得
    categories, err := h.categoryService.GetCategories(ctx.Request().Context())
    if err != nil {
      // エラーハンドリング
      errorResponse := openapi.ErrorResponse{
      Code:    "internal_server_error",
      Message: "Failed to fetch categories",
      Details: &map[string]interface{}{
        "error": err.Error(),
      },
      }
      return ctx.JSON(http.StatusInternalServerError, errorResponse)
    }

    // レスポンスの構築
    items := make([]openapi.Category, 0, len(categories))
    for _, c := range categories {
      // カスタムプロパティから商品数を取得
      var productCount *int
      if count, ok := c.Preload("product_count").(int64); ok {
      count := int(count)
      productCount = &count
      }

      items = append(items, openapi.Category{
      Id:           int64(c.ID),
      Name:         c.Name,
      Description:  stringPtr(c.Description.String),
      ImageUrl:     stringPtr(c.ImageURL.String),
      ProductCount: productCount,
      })
    }

    response := openapi.CategoryList{
      Items: items,
    }

    return ctx.JSON(http.StatusOK, response)
    }
    ```

4. ハンドラの登録をhandler.goファイルに追加します：

    ```bash
    vim backend-api/internal/api/handlers/handler.go
    ```

    以下のようにハンドラ構造体とNewHandler関数を更新します：

    ```go
    type Handler struct {
    *HealthHandler
    *ProductHandler
    *CategoryHandler
    }

    // NewHandler はHandlerのインスタンスを生成する
    func NewHandler() *Handler {
    return &Handler{
      HealthHandler:   NewHealthHandler(),
      ProductHandler:  NewProductHandler(),
      CategoryHandler: NewCategoryHandler(),
    }
    }
    ```

### 1.4.3. カテゴリー別商品一覧APIの実装

カテゴリー別商品一覧APIを実装します。

1. ProductHandlerのListProductsByCategoryメソッドを更新します：

    ```bash
    vim backend-api/internal/api/handlers/product.go
    ```

    `ListProductsByCategory` メソッドを以下のコードで更新します：

    ```go
    // ListProductsByCategory は指定されたカテゴリーの商品一覧を取得する
    func (h *ProductHandler) ListProductsByCategory(ctx echo.Context, id int64, params openapi.ListProductsByCategoryParams) error {
    // パラメータの取得
    page := 1
    if params.Page != nil {
      page = *params.Page
    }

    pageSize := 20
    if params.PageSize != nil {
      pageSize = *params.PageSize
    }

    // サービスからカテゴリー別商品一覧を取得
    result, err := h.productService.GetProductsByCategory(ctx.Request().Context(), int(id), page, pageSize)
    if err != nil {
      // エラーの種類に応じて適切なレスポンスを返す
      if strings.Contains(err.Error(), "category not found") {
      errorResponse := openapi.ErrorResponse{
        Code:    "category_not_found",
        Message: "Category not found",
        Details: &map[string]interface{}{
        "id": id,
        },
      }
      return ctx.JSON(http.StatusNotFound, errorResponse)
      }

      // その他のエラー
      errorResponse := openapi.ErrorResponse{
      Code:    "internal_server_error",
      Message: "Failed to fetch products by category",
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
    ```

### 1.4.4. リポジトリとサービスの拡張

ベースリポジトリが未実装の場合は、以下のファイルを作成します：

```bash
touch backend-api/internal/repository/repository.go
```

リポジトリベースクラスに以下のコードを追加します：

```go
package repository

import (
 "database/sql"

 "github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository はリポジトリの基本インターフェース
type Repository interface {
 DB() boil.ContextExecutor
 WithTx(tx *sql.Tx) Repository
}

// RepositoryBase はリポジトリの基本実装
type RepositoryBase struct {
 db boil.ContextExecutor
}

// NewRepositoryBase は新しいRepositoryBaseを作成します
func NewRepositoryBase(executor boil.ContextExecutor) RepositoryBase {
 return RepositoryBase{
  db: executor,
 }
}

// DB はデータベース接続を返します
func (r RepositoryBase) DB() boil.ContextExecutor {
 return r.db
}

// WithTx はトランザクションを設定したRepositoryBaseを返します
func (r RepositoryBase) WithTx(tx *sql.Tx) Repository {
 return RepositoryBase{
  db: tx,
 }
}
```

トランザクション管理用のヘルパー関数を追加します：

```bash
touch backend-api/internal/repository/transaction.go
```

トランザクション管理用のコードを追加します：

```go
package repository

import (
 "context"
 "database/sql"
 "fmt"
)

// RunInTransaction はトランザクション内で関数を実行します
func RunInTransaction(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
 // トランザクションを開始
 tx, err := db.BeginTx(ctx, nil)
 if err != nil {
  return fmt.Errorf("failed to begin transaction: %w", err)
 }

 // 関数の実行
 if err := fn(tx); err != nil {
  // エラーの場合はロールバック
  if rbErr := tx.Rollback(); rbErr != nil {
   return fmt.Errorf("failed to rollback transaction: %w (original error: %v)", rbErr, err)
  }
  return err
 }

 // トランザクションのコミット
 if err := tx.Commit(); err != nil {
  return fmt.Errorf("failed to commit transaction: %w", err)
 }

 return nil
}
```

## 1.5. 【確認ポイント】

実装が正しく完了したことを確認するためのチェックリストです：

- [ ] サーバーが正常に起動し、コンパイルエラーが発生しないこと
- [ ] 商品詳細APIが正常に機能し、正しい商品情報を返すこと

  ```bash
  curl http://backend-api.localhost/api/products/1 | jq
  ```

- [ ] 存在しない商品IDを指定した場合に404エラーが返されること

  ```bash
  curl http://backend-api.localhost/api/products/999 | jq
  ```

- [ ] カテゴリー一覧APIが正常に機能し、全カテゴリーを返すこと

  ```bash
  curl http://backend-api.localhost/api/categories | jq
  ```

- [ ] カテゴリー別商品一覧APIが正常に機能し、指定カテゴリーの商品を返すこと

  ```bash
  curl "http://backend-api.localhost/api/categories/1/products?page=1&pageSize=10" | jq
  ```

- [ ] Swagger UIでAPI仕様が正確に表示されること

  ```bash
  open http://backend-api.localhost/swagger
  ```

## 1.6. 【詳細解説】

### 1.6.1. URLパラメータの使用方法とベストプラクティス

RESTful APIでは、リソースの識別にURLパラメータを使用することが一般的です。Echoフレームワークでは、`:param`形式でパスパラメータを定義し、コントローラー内で`ctx.Param("param")`のように取得できます。

URLパラメータを使用する際のベストプラクティス：

1. **適切なHTTPメソッドの使用**：
   - GET：リソースの取得
   - POST：新しいリソースの作成
   - PUT：リソースの更新
   - DELETE：リソースの削除

2. **リソースの設計**：
   - 名詞を使用する（/products, /categories）
   - 複数形を使用する（/products, not /product）
   - リレーションを表現するためにネストを使用する（/categories/{id}/products）

3. **クエリパラメータとパスパラメータの使い分け**：
   - パスパラメータ：リソースの識別（/products/{id}）
   - クエリパラメータ：フィルタリング、ソート、ページネーション（?page=1&pageSize=10）

4. **エラーハンドリング**：
   - 存在しないリソースに対しては404を返す
   - 不正なパラメータに対しては400を返す
   - サーバーエラーに対しては500を返す

### 1.6.2. リポジトリパターンの効果的な適用

リポジトリパターンは、データアクセスを抽象化するためのパターンです。このパターンを使用することで、ビジネスロジックとデータアクセスを分離し、テストや保守性を向上させることができます。

1. **リポジトリの責務**：
   - データの取得、保存、更新、削除
   - クエリの構築
   - データベース固有のロジックのカプセル化

2. **サービス層の責務**：
   - ビジネスロジックの実装
   - トランザクション管理
   - 複数のリポジトリの協調

3. **依存性の注入**：
   - リポジトリはインターフェースとして定義
   - サービスはリポジトリのインスタンスを受け取る
   - テスト時にはモックリポジトリを注入可能

4. **トランザクション管理**：
   - 複数のデータベース操作を一つのトランザクションで実行
   - WithTxメソッドを使用してトランザクション境界を制御
   - エラー発生時の自動ロールバック

### 1.6.3. レスポンス形式の一貫性と統一性

APIのレスポンス形式は、クライアントとの契約の一部です。一貫性のあるレスポンス形式を提供することで、クライアント側の実装が容易になります。

1. **成功レスポンスの形式統一**：
   - リスト形式：`{"items": [...], "total": 100, "page": 1, ...}`
   - 単一リソース形式：`{"id": 1, "name": "Product", ...}`

2. **エラーレスポンスの形式統一**：
   - エラーコード：`"code": "not_found"`
   - メッセージ：`"message": "Product not found"`
   - 詳細情報：`"details": {"id": 999, ...}`

3. **HTTPステータスコードの適切な使用**：
   - 200 OK：リクエスト成功
   - 201 Created：リソース作成成功
   - 400 Bad Request：クライアントエラー
   - 401 Unauthorized：認証エラー
   - 403 Forbidden：権限エラー
   - 404 Not Found：リソースが存在しない
   - 500 Internal Server Error：サーバーエラー

4. **Content-Type**：
   - `application/json`を使用
   - 適切な文字コード（UTF-8）を指定

## 1.7. 【補足情報】

### 1.7.1. N+1問題とEagerローディング

N+1問題は、ORMを使用する際によく発生する問題で、例えば商品のリストを取得する際に、各商品のカテゴリーを別々のクエリで取得することで大量のクエリが発生する問題です。

```plaintext
SELECT * FROM products;  -- 商品を100件取得
SELECT * FROM categories WHERE id = 1;  -- 商品1のカテゴリーを取得
SELECT * FROM categories WHERE id = 2;  -- 商品2のカテゴリーを取得
...
SELECT * FROM categories WHERE id = 100;  -- 商品100のカテゴリーを取得
```

この問題を解決するために、Eagerローディングを使用します。Eagerローディングでは、関連するデータを一度に取得します。

```plaintext
SELECT * FROM products;  -- 商品を100件取得
SELECT * FROM categories WHERE id IN (1, 2, ..., 100);  -- 全ての商品のカテゴリーを一度に取得
```

SQLBoilerでは、以下のようにEagerローディングを指定できます。

```go
products, err := models.Products(
    qm.Load("Category"),  // カテゴリーを一緒に取得
    qm.Load("Inventories"),  // 在庫情報も一緒に取得
).All(ctx, db)
```

また、既に取得したモデルに対して後からEagerローディングを行うこともできます。

```go
product, err := models.FindProduct(ctx, db, id)
err = product.L.LoadCategory(ctx, db, true, product, nil)
```

### 1.7.2. OpenAPIスキーマとコード生成の連携

OpenAPI仕様はAPIの設計と実装を一貫させるための重要なツールです。OpenAPI仕様からコードを生成することで、以下のメリットがあります。

1. **型安全性**：
   - スキーマに基づいた型が生成される
   - コンパイル時に型エラーを検出できる

2. **開発効率**：
   - APIの仕様変更が自動的にコードに反映される
   - ボイラープレートコードの削減

3. **ドキュメントの自動生成**：
   - Swagger UIによるAPIドキュメントの提供
   - クライアントやフロントエンド開発者との協業が容易になる

4. **コード生成とカスタム実装の分離**：
   - 生成されたインターフェースを実装する形で開発
   - 再生成時にカスタム実装が上書きされない

Go言語では、`oapi-codegen`や`go-swagger`などのツールを使用してOpenAPI仕様からコードを生成できます。これらのツールを使用することで、型定義、サーバースタブ、クライアントコードなどを自動生成できます。

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: パスパラメータが正しく取得できない

**症状**: URLのパスパラメータが正しく取得できず、`nil`や空文字列になる。

**解決策**:

1. ルーティングの定義を確認する：

   ```go
   // 正しい定義
   e.GET("/api/products/:id", handler.GetProductByID)
   ```

2. パラメータの取得方法を確認する：

   ```go
   // 正しい取得方法
   id := ctx.Param("id")
   ```

3. OpenAPI仕様とEchoのルーティングが一致していることを確認する：

   ```yaml
   # OpenAPI仕様
   /products/{id}:
     get:
       parameters:
         - name: id
           in: path
           required: true
           schema:
             type: integer
   ```

4. oapi-codegenを使用している場合は生成されたコードを確認する：

   ```go
   // 生成されたハンドラーインターフェース
   type ServerInterface interface {
       GetProductById(ctx echo.Context, id int64) error
   }
   ```

### 1.8.2. 問題2: 関連データの取得でN+1問題が発生する

**症状**: データベースクエリのログを見ると、商品リストを取得するときに多数のクエリが発行されている。

**解決策**:

1. Eagerローディングを使用する：

   ```go
   products, err := models.Products(
       qm.Load("Category"),
       qm.Load("Inventories"),
   ).All(ctx, db)
   ```

2. クエリの効率を確認する：

   ```bash
   # MySQLのクエリログを有効にする
   mysql -u root -p
   SET GLOBAL general_log = 'ON';
   SET GLOBAL log_output = 'TABLE';

   # ログを確認する
   SELECT * FROM mysql.general_log;
   ```

3. 関連データの取得を最適化する：

   ```go
   // 一度に全てのカテゴリーを取得
   categoryIDs := make([]interface{}, 0, len(products))
   for _, p := range products {
       categoryIDs = append(categoryIDs, p.CategoryID)
   }

   categories, err := models.Categories(
       qm.WhereIn("id IN ?", categoryIDs...),
   ).All(ctx, db)

   // 商品にカテゴリー情報を設定
   categoryMap := make(map[int]*models.Category)
   for _, c := range categories {
       categoryMap[c.ID] = c
   }

   for _, p := range products {
       if c, ok := categoryMap[p.CategoryID]; ok {
           p.R.Category = c
       }
   }
   ```

4. クエリ自体を最適化する：

   ```go
   // JOINを使用して一度に取得
   products, err := models.Products(
       qm.InnerJoin("categories c ON c.id = products.category_id"),
       qm.Select("products.*, c.name as category_name"),
   ).All(ctx, db)
   ```

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **URLパラメータを使用したリソース取得**：
   - RESTful APIの設計原則に従い、リソースの識別にURLパラメータを使用
   - パスパラメータとクエリパラメータの使い分けを理解する

2. **リポジトリパターンとサービス層の役割分担**：
   - リポジトリはデータアクセスのみを担当
   - サービス層はビジネスロジックを実装
   - 依存性の注入を活用したモジュール間の疎結合化

3. **一貫したエラーハンドリング**：
   - エラーの種類に応じた適切なHTTPステータスコードの返却
   - 統一されたエラーレスポンス形式の提供
   - エラー情報の適切なログ記録

4. **関連データの効率的な取得**：
   - N+1問題を回避するためのEagerローディングの活用
   - 適切なインデックス設計によるクエリのパフォーマンス最適化

これらのポイントは、次回以降の実装でも活用されますので、よく理解しておきましょう。

## 1.10. 【次回の準備】

次回（Day 3）では、Lambda関数とS3連携の実装に取り組みます。以下の点について事前に確認しておくと良いでしょう：

1. **AWSの基本概念**：
   - S3（Simple Storage Service）の基本的な使い方
   - Lambda関数の基本構造と実行モデル
   - IAMロールとポリシーの概念

2. **AWS SDKの基本**：
   - Go用AWS SDKのインストールと基本的な使い方
   - S3とLambdaを操作するためのAPIの基本

3. **LocalStackの設定**：
   - LocalStackでS3とLambdaをエミュレートする方法
   - AWS CLIでLocalStackのサービスにアクセスする方法

4. **画像処理の基本**：
   - Goでの画像処理の基本（リサイズ、フォーマット変換など）
   - 画像処理ライブラリ（imaging）の基本的な使い方

これらの知識を事前に確認しておくことで、次回のLambda関数とS3連携の実装をスムーズに進めることができます。
