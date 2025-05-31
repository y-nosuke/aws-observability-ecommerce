# 1. DDD+クリーンアーキテクチャ対応ディレクトリ構成

- [1. DDD+クリーンアーキテクチャ対応ディレクトリ構成](#1-dddクリーンアーキテクチャ対応ディレクトリ構成)
  - [1.1. 実装上の問題点への対応](#11-実装上の問題点への対応)
    - [1.1.1. SQLBoiler生成ファイルの現実](#111-sqlboiler生成ファイルの現実)
    - [1.1.2. 複数ドメインをまたぐ処理](#112-複数ドメインをまたぐ処理)
    - [1.1.3. 設定とプレゼンテーション層の配置](#113-設定とプレゼンテーション層の配置)
  - [1.2. ディレクトリ構成](#12-ディレクトリ構成)
  - [1.3. 各層の責務と配置基準](#13-各層の責務と配置基準)
    - [1.3.1. Domain層（ドメイン層）](#131-domain層ドメイン層)
    - [1.3.2. Application層（アプリケーション層）](#132-application層アプリケーション層)
    - [1.3.3. Infrastructure層（インフラストラクチャ層）](#133-infrastructure層インフラストラクチャ層)
    - [1.3.4. Presentation層（プレゼンテーション層）](#134-presentation層プレゼンテーション層)
  - [1.4. 新技術追加時の配置判断フロー](#14-新技術追加時の配置判断フロー)
    - [1.4.1. gRPC追加の場合](#141-grpc追加の場合)
    - [1.4.2. GraphQL追加の場合](#142-graphql追加の場合)
    - [1.4.3. CLI追加の場合](#143-cli追加の場合)
    - [1.4.4. 新しいデータベース（Redis等）追加の場合](#144-新しいデータベースredis等追加の場合)
    - [1.4.5. 外部API連携追加の場合](#145-外部api連携追加の場合)
  - [1.5. 判断基準まとめ](#15-判断基準まとめ)
  - [1.6. 主要な設計方針](#16-主要な設計方針)
    - [1.6.1. SQLBoiler生成ファイルの統一配置](#161-sqlboiler生成ファイルの統一配置)
    - [1.6.2. Query側のReader + Mapperパターン](#162-query側のreader--mapperパターン)
    - [1.6.3. APIハンドラーでの使い分け](#163-apiハンドラーでの使い分け)
    - [1.6.4. 設定の共有](#164-設定の共有)
  - [1.7. 利点](#17-利点)
  - [1.8. Lambdaでの利用について](#18-lambdaでの利用について)
    - [1.8.1. Lambda用の構成パターン](#181-lambda用の構成パターン)
      - [1.8.1.1. パターンA: 単一Lambda（モノリシック）](#1811-パターンa-単一lambdaモノリシック)
      - [1.8.1.2. パターンB: マイクロLambda（機能分離）](#1812-パターンb-マイクロlambda機能分離)
    - [1.8.2. Lambda固有の考慮点](#182-lambda固有の考慮点)
      - [1.8.2.1. Presentation層の扱い](#1821-presentation層の扱い)
      - [1.8.2.2. イベント駆動処理](#1822-イベント駆動処理)
      - [1.8.2.3. 依存性注入の調整](#1823-依存性注入の調整)
    - [1.8.3. Lambda使用時の利点](#183-lambda使用時の利点)

## 1.1. 実装上の問題点への対応

### 1.1.1. SQLBoiler生成ファイルの現実

- SQLBoilerはデータベース全体を一括生成（テーブル単位の出し分け不可）
- 全て1箇所に配置し、各ドメインがimportして使用

### 1.1.2. 複数ドメインをまたぐ処理

- CoordinatorパターンはN+1問題とマッピング複雑化を招く
- **Read Model**による専用クエリで対応
- 複数ドメインをまたぐAPIは `internal/query` ディレクトリに配置
  - 例：商品詳細取得（商品 + カテゴリー + 在庫）
  - 例：商品カタログ（商品 + カテゴリー）
  - 例：ダッシュボード（商品 + カテゴリー + 在庫 + 注文）

### 1.1.3. 設定とプレゼンテーション層の配置

- 設定は最上位で一元管理
- プレゼンテーション層も最上位で統一

## 1.2. ディレクトリ構成

```text
backend-api/
├── cmd/
│   └── api/
│       └── main.go                          # アプリケーションのエントリーポイント
├── pkg/                                     # 外部公開可能な共通ライブラリ
│   ├── logger/
│   ├── errors/
│   └── observability/
├── internal/
│   ├── product/                             # 商品ドメイン（詳細構成例）
│   │   ├── application/
│   │   │   ├── usecase/
│   │   │   │   ├── create_product.go        # 商品作成
│   │   │   │   ├── update_product.go        # 商品更新
│   │   │   │   ├── delete_product.go        # 商品削除
│   │   │   │   ├── get_product.go           # 商品取得（単体のみ）
│   │   │   │   └── upload_product_image.go  # 画像アップロード
│   │   │   └── dto/
│   │   │       ├── create_product_request.go
│   │   │       ├── create_product_response.go
│   │   │       ├── update_product_request.go
│   │   │       ├── update_product_response.go
│   │   │       ├── get_product_response.go
│   │   │       ├── delete_product_response.go
│   │   │       ├── upload_image_request.go
│   │   │       └── upload_image_response.go
│   │   ├── domain/
│   │   │   ├── entity/
│   │   │   │   ├── product.go               # 商品エンティティ（ビジネスロジック含む）
│   │   │   │   └── product_test.go
│   │   │   ├── valueobject/
│   │   │   │   ├── product_id.go
│   │   │   │   ├── price.go
│   │   │   │   ├── sku.go
│   │   │   │   ├── product_status.go
│   │   │   │   └── image_url.go
│   │   │   ├── repository/
│   │   │   │   └── product_repository.go    # リポジトリインターフェース
│   │   │   ├── service/
│   │   │   │   ├── product_domain_service.go
│   │   │   │   └── image_storage.go         # 画像ストレージインターフェース
│   │   │   └── factory/
│   │   │       └── product_factory.go
│   │   ├── infrastructure/
│   │   │   ├── persistence/
│   │   │   │   ├── mapper/
│   │   │   │   │   └── product_mapper.go    # model.Product ⇔ domain.Product
│   │   │   │   └── repository/
│   │   │   │       └── product_repository_impl.go # リポジトリ実装
│   │   │   └── external/
│   │   │       └── storage/
│   │   │           └── s3_image_storage_impl.go # S3画像ストレージ実装
│   │   └── presentation/
│   │       └── rest/
│   │           ├── handler/
│   │           │   └── product_handler.go   # 商品API（command専用）
│   │           └── presenter/
│   │               └── product_presenter.go
│   ├── category/                            # カテゴリドメイン（productと同様の構成）
│   │   ├── application/                     # usecase, dto
│   │   ├── domain/                          # entity, valueobject, repository, service, factory
│   │   ├── infrastructure/                  # persistence (mapper, repository)
│   │   └── presentation/                    # rest (handler, presenter)
│   ├── inventory/                           # 在庫ドメイン（productと同様の構成）
│   │   ├── application/                     # usecase, dto
│   │   ├── domain/                          # entity, valueobject, repository, service, factory
│   │   ├── infrastructure/                  # persistence (mapper, repository)
│   │   └── presentation/                    # rest (handler, presenter)
│   ├── query/                               # 複数ドメインをまたぐRead専用サービス
│   │   └── rest/
│   │       ├── handler/
│   │       │   ├── product_catalog_handler.go     # 商品カタログクエリAPI
│   │       │   ├── category_stats_handler.go      # カテゴリ統計クエリAPI
│   │       │   └── dashboard_handler.go           # ダッシュボードクエリAPI
│   │       ├── reader/                      # DBアクセス・データ読み取り
│   │       │   ├── product_catalog_reader.go
│   │       │   ├── category_stats_reader.go
│   │       │   └── dashboard_reader.go
│   │       └── mapper/                      # SQLBoilerモデル → レスポンスDTO変換
│   │           ├── product_catalog_mapper.go
│   │           ├── category_stats_mapper.go
│   │           └── dashboard_mapper.go
│   └── shared/                              # ドメイン間共通
│       ├── application/                     # 横断的アプリケーションサービス
│       │   └── transaction/
│       │       └── transaction_manager.go
│       ├── domain/                          # 共通ドメインプリミティブ
│       │   ├── valueobject/
│       │   │   ├── id.go                    # 共通ID型
│       │   │   ├── email.go
│       │   │   └── phone.go
│       │   └── event/
│       │       ├── domain_event.go
│       │       └── event_dispatcher.go
│       ├── infrastructure/                  # 共通インフラ
│       │   ├── database/
│       │   │   ├── connection.go
│       │   │   └── transaction.go
│       │   ├── model/                       # SQLBoiler生成ファイル（全テーブル）
│       │   │   ├── boil_queries.go
│       │   │   ├── boil_table_names.go
│       │   │   ├── boil_types.go
│       │   │   ├── boil_view_names.go
│       │   │   ├── mysql_upsert.go
│       │   │   ├── categories.go            # カテゴリテーブル
│       │   │   ├── products.go              # 商品テーブル
│       │   │   ├── inventory.go             # 在庫テーブル
│       │   │   └── test.go
│       │   ├── aws/
│       │   │   ├── client_factory.go        # AWS設定を元にクライアント生成
│       │   │   └── s3_client_wrapper.go     # S3の薄いラッパー
│       │   └── config/                      # アプリケーション設定
│       │       ├── app_config.go
│       │       ├── database_config.go
│       │       └── aws_config.go
│       └── presentation/                    # 共通プレゼンテーション要素
│           └── rest/
│               ├── handler/
│               │   └── health_handler.go
│               ├── router/
│               │   └── router.go
│               ├── middleware/
│               │   ├── cors.go
│               │   ├── logging.go
│               │   ├── metrics.go
│               │   └── error_handler.go
│               ├── presenter/
│               │   └── error_presenter.go
│               └── openapi/                          # OpenAPI生成
│                   └── oapi-codegen-config.gen.go    # oapi-codegen生成
├── api/                                     # API仕様
│   ├── openapi.yaml
│   └── docs/
│       └── swagger-ui/
├── migration/                               # DBマイグレーション
├── scripts/
├── test/
├── di/                                      # 依存性注入
│   ├── wire.go
│   ├── container.go
│   └── provider/
├── build/
├── go.mod
└── go.sum
```

## 1.3. 各層の責務と配置基準

### 1.3.1. Domain層（ドメイン層）

**責務**: ビジネスロジックとビジネスルールの実装

- **Entity**: ビジネスオブジェクトとその振る舞い
- **Value Object**: 値に関するビジネスルールとバリデーション
- **Repository Interface**: データアクセスの抽象化
- **Domain Service**: 複数エンティティにまたがるビジネスロジック
- **Factory**: エンティティの複雑な生成ロジック

**配置基準**:

- ビジネスルールを含むもの
- 外部技術に依存しないもの
- ドメインエキスパートと議論する内容

```go
// ✅ Domain層に配置
type Product struct {
// ビジネスロジック
func (p *Product) ChangePrice(newPrice Price) error
func (p *Product) ApplyDiscount(rate float64) error
}

// ❌ Domain層に配置すべきでない
type ProductHTTPHandler struct {} // → Presentation層
type ProductDBRepository struct {} // → Infrastructure層
```

### 1.3.2. Application層（アプリケーション層）

**責務**: ユースケースの実行とワークフローの調整

- **UseCase**: 具体的なビジネスユースケースの実行
- **DTO**: 層間でのデータ転送
- **Application Service**: 複数ドメインの調整（必要に応じて）

**配置基準**:

- ユーザーのアクションに対応するもの
- 複数のドメインサービスやリポジトリを協調させるもの
- トランザクション境界を定義するもの

```go
// ✅ Application層に配置
type CreateProductUseCase struct {
// ユースケースの実行
func Execute(req CreateProductRequest) (*CreateProductResponse, error)
}

// ❌ Application層に配置すべきでない
type Product struct {} // → Domain層
type ProductController struct {} // → Presentation層
```

### 1.3.3. Infrastructure層（インフラストラクチャ層）

**責務**: 外部技術との連携と技術的実装（**アプリケーション → 外部**の通信）

- **Repository Implementation**: データベースアクセスの実装
- **External Service**: 外部APIとの連携（アプリから外部API呼び出し）
- **Mapper**: ドメインオブジェクトと外部データ形式の変換
- **Configuration**: 技術的設定の管理

**配置基準**:

- アプリケーションから外部への**出力**
- 具体的な技術に依存するもの
- データベース、ファイルシステム、外部APIとの連携
- フレームワークやライブラリの詳細

```go
// ✅ Infrastructure層に配置
type ProductSQLRepository struct {
// データベース実装
func FindByID(id ProductID) (*Product, error)
}

type S3ImageStorage struct {
// AWS S3実装
func Upload(image []byte) (string, error)
}

// ❌ Infrastructure層に配置すべきでない
type CreateProductUseCase struct {} // → Application層
```

### 1.3.4. Presentation層（プレゼンテーション層）

**責務**: 外部インターフェースとの通信（**外部 → アプリケーション**の通信）

- **REST Handler**: HTTP APIの処理（クライアントからのリクエスト受信）
- **gRPC Handler**: gRPC APIの処理（他サービスからのリクエスト受信）
- **GraphQL Resolver**: GraphQL APIの処理（クエリ受信）
- **CLI Command**: コマンドラインインターフェース（ユーザーからのコマンド受信）
- **Presenter**: レスポンス形式の変換

**配置基準**:

- 外部からアプリケーションへの**入力**
- 外部からのリクエストを受け取るもの
- プロトコル固有の処理
- UI/API仕様に依存するもの

```go
// ✅ Presentation層に配置
type ProductRESTHandler struct {
// HTTP API処理
func CreateProduct(ctx echo.Context) error
}

type ProductGRPCHandler struct {
// gRPC API処理
func CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
}

// ❌ Presentation層に配置すべきでない
type CreateProductUseCase struct {} // → Application層
type ProductRepository struct {} // → Infrastructure層
```

## 1.4. 新技術追加時の配置判断フロー

### 1.4.1. gRPC追加の場合

```text
product/
└── presentation/
    ├── rest/          # 既存のREST API
    │   └── handler/
    └── grpc/          # 新規追加
        ├── handler/   # gRPCハンドラー
        ├── interceptor/ # gRPCインターセプター
        └── pb/        # protobuf生成ファイル
```

### 1.4.2. GraphQL追加の場合

```text
product/
└── presentation/
    ├── rest/
    ├── grpc/
    └── graphql/       # 新規追加
        ├── resolver/  # GraphQLリゾルバー
        ├── schema/    # GraphQLスキーマ
        └── model/     # GraphQL専用モデル
```

### 1.4.3. CLI追加の場合

```text
product/
└── presentation/
    ├── rest/
    ├── grpc/
    └── cli/           # 新規追加
        ├── command/   # CLIコマンド
        └── flag/      # CLIフラグ定義
```

### 1.4.4. 新しいデータベース（Redis等）追加の場合

```text
product/
└── infrastructure/
    ├── persistence/   # 既存のRDB
    └── cache/         # 新規追加
        ├── repository/ # Redis実装
        └── config/    # Redis設定
```

### 1.4.5. 外部API連携追加の場合

```text
product/
└── infrastructure/
    ├── persistence/
    └── external/      # 既存
        ├── storage/   # 既存のS3
        └── payment/   # 新規追加（決済API等）
            ├── stripe/
            └── paypal/
```

## 1.5. 判断基準まとめ

| 要素                     | Domain | Application | Infrastructure | Presentation |
| ------------------------ | ------ | ----------- | -------------- | ------------ |
| **ビジネスロジック**     | ✅      | ❌           | ❌              | ❌            |
| **ユースケース実行**     | ❌      | ✅           | ❌              | ❌            |
| **データベースアクセス** | ❌      | ❌           | ✅              | ❌            |
| **HTTP API**             | ❌      | ❌           | ❌              | ✅            |
| **gRPC API**             | ❌      | ❌           | ❌              | ✅            |
| **外部API連携**          | ❌      | ❌           | ✅              | ❌            |
| **設定管理**             | ❌      | ❌           | ✅              | ❌            |
| **バリデーション**       | ✅      | ✅           | ❌              | ✅            |

**バリデーションの配置**:

- **Domain**: ビジネスルールのバリデーション（価格は0以上など）
- **Application**: ユースケース固有のバリデーション（必須項目チェックなど）
- **Presentation**: 入力形式のバリデーション（JSON形式、必須フィールドなど）

## 1.6. 主要な設計方針

### 1.6.1. SQLBoiler生成ファイルの統一配置

```text
internal/shared/infrastructure/model/
├── products.go      # 全テーブル
├── categories.go    # を1箇所に
├── inventory.go     # 統一配置
└── ...
```

各ドメインはこれらをimportして使用：

```go
// internal/product/infrastructure/persistence/mapper/product_mapper.go
import "github.com/project/internal/shared/infrastructure/model"

func (m *ProductMapper) ToEntity(dbModel *model.Product) (*entity.Product, error) {
return &entity.Product{
ID:    ProductID(dbModel.ID),
Name:  dbModel.Name,
Price: NewPrice(dbModel.Price),
// ...
}
}
```

### 1.6.2. Query側のReader + Mapperパターン

複数ドメインをまたぐRead専用処理：

```go
// internal/query/infrastructure/reader/product_catalog_reader.go
func (r *ProductCatalogReader) FindProductsWithDetails(ctx context.Context, params ProductListParams) ([]*model.Product, error) {
query := `
        SELECT
            p.id, p.name, p.price, p.image_url, p.sku,
            c.id as category_id, c.name as category_name, c.slug as category_slug,
            i.quantity, i.id as inventory_id,
            CASE WHEN i.quantity > 0 THEN true ELSE false END as in_stock
        FROM products p
        LEFT JOIN categories c ON p.category_id = c.id
        LEFT JOIN inventory i ON p.id = i.product_id
        WHERE p.deleted_at IS NULL
        ORDER BY p.created_at DESC
        LIMIT ? OFFSET ?
    `

var results []*model.Product
err := r.db.SelectContext(ctx, &results, query, params.Limit, params.Offset)
return results, err
}

// internal/query/infrastructure/mapper/product_catalog_mapper.go
func (m *ProductCatalogMapper) ToProductListResponse(products []*model.Product) *ProductListResponse {
items := make([]ProductItem, len(products))
for i, p := range products {
items[i] = ProductItem{
ID:           p.ID,
Name:         p.Name,
Price:        p.Price,
CategoryName: p.R.Category.Name,
InStock:      len(p.R.Inventories) > 0 && p.R.Inventories[0].Quantity > 0,
}
}
return &ProductListResponse{Items: items}
}
```

### 1.6.3. APIハンドラーでの使い分け

```go
// internal/product/presentation/rest/handler/product_handler.go
type ProductHandler struct {
// Command用（単一ドメイン）
createProductUseCase  product.CreateProductUseCase
updateProductUseCase  product.UpdateProductUseCase
deleteProductUseCase  product.DeleteProductUseCase
}

// 商品作成（単一ドメイン）
func (h *ProductHandler) CreateProduct(ctx echo.Context) error {
err := h.createProductUseCase.Execute(ctx.Request().Context(), ...)
// ...
}

// internal/query/presentation/rest/handler/product_catalog_handler.go
type ProductCatalogHandler struct {
reader ProductCatalogReader
mapper ProductCatalogMapper
}

// 商品一覧取得（複数ドメイン）
func (h *ProductCatalogHandler) ListProducts(ctx echo.Context, params openapi.ListProductsParams) error {
products, err := h.reader.FindProductsWithDetails(ctx.Request().Context(), ...)
if err != nil {
return err
}

response := h.mapper.ToProductListResponse(products)
return ctx.JSON(http.StatusOK, response)
}
```

### 1.6.4. 設定の共有

```go
// internal/shared/infrastructure/config/app_config.go
var (
AppConfig      Config
DatabaseConfig DatabaseConfig
AWSConfig      AWSConfig
)

// 各パッケージから参照
// internal/product/infrastructure/persistence/repository/product_repository_impl.go
import "github.com/project/internal/shared/infrastructure/config"

func NewProductRepository() *ProductRepository {
db := config.DatabaseConfig.GetConnection()
return &ProductRepository{db: db}
}
```

## 1.7. 利点

1. **業務理解の容易性**: フォルダ構成が業務ドメインを反映
2. **パフォーマンス**: Query側のReader + MapperによりN+1問題を回避
3. **保守性**: 設定とプレゼンテーション層の適切な配置
4. **拡張性**: 必要に応じてドメインの追加が容易
5. **現実的妥協**: Query側でのDIP緩和により実装効率向上

## 1.8. Lambdaでの利用について

**この構成はLambdaでも利用可能です。**ただし、以下の点を考慮する必要があります：

### 1.8.1. Lambda用の構成パターン

#### 1.8.1.1. パターンA: 単一Lambda（モノリシック）

```text
backend-api/
├── cmd/
│   └── lambda/
│       └── main.go              # Lambda エントリーポイント
├── internal/
│   ├── product/
│   ├── category/
│   ├── inventory/
│   ├── query/
│   └── shared/
└── lambda/                      # Lambda固有
    ├── handler/
    │   ├── api_gateway_handler.go    # API Gateway → 全ドメイン
    │   ├── s3_event_handler.go       # S3イベント → 画像処理
    │   └── sqs_event_handler.go      # SQSイベント → 非同期処理
    └── adapter/
        ├── event_mapper.go          # AWSイベント → DTOマッピング
        └── response_mapper.go       # DTO → AWSレスポンスマッピング
```

#### 1.8.1.2. パターンB: マイクロLambda（機能分離）

```text
lambda-product/
├── cmd/lambda/main.go
└── internal/
    ├── product/                    # 商品ドメインのみ
    ├── shared/                     # 共通部分（別リポジトリまたは共有ライブラリ）
    └── lambda/
        └── handler/
            └── product_handler.go

lambda-inventory/
├── cmd/lambda/main.go
└── internal/
    ├── inventory/                  # 在庫ドメインのみ
    ├── shared/
    └── lambda/
        └── handler/
            └── inventory_handler.go
```

### 1.8.2. Lambda固有の考慮点

#### 1.8.2.1. Presentation層の扱い

```go
// lambda/handler/api_gateway_handler.go
type APIGatewayHandler struct {
productHandler   *product.ProductHandler    // 既存のhandler再利用
categoryHandler  *category.CategoryHandler
inventoryHandler *inventory.InventoryHandler
}

func (h *APIGatewayHandler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// ルーティング処理
switch {
case strings.HasPrefix(req.Path, "/products"):
return h.handleProductRequest(ctx, req)
case strings.HasPrefix(req.Path, "/categories"):
return h.handleCategoryRequest(ctx, req)
}
}
```

#### 1.8.2.2. イベント駆動処理

```go
// lambda/handler/s3_event_handler.go
type S3EventHandler struct {
uploadImageUseCase product.UploadProductImageUseCase
}

func (h *S3EventHandler) Handle(ctx context.Context, s3Event events.S3Event) error {
// S3イベント → ドメインロジック
for _, record := range s3Event.Records {
// 既存のusecaseを再利用
err := h.uploadImageUseCase.Execute(ctx, dto.UploadImageRequest{
ProductID: extractProductID(record.S3.Object.Key),
ImageURL:  buildImageURL(record.S3.Bucket.Name, record.S3.Object.Key),
})
}
}
```

#### 1.8.2.3. 依存性注入の調整

```go
// di/lambda_container.go
func NewLambdaContainer() *Container {
// Lambda環境用の設定
db := initializeDatabaseConnection()

container := &Container{db: db}

// 必要な場合のみインスタンス化（コールドスタート最適化）
return container
}
```

### 1.8.3. Lambda使用時の利点

1. **ドメインロジックの再利用**: 既存のusecase/entityをそのまま利用
2. **テスタビリティ**: ビジネスロジックはLambda環境に依存しない
3. **段階的移行**: 一部機能のみLambda化が可能
4. **スケーラビリティ**: 負荷に応じたドメイン単位のスケーリング

この構成により、業務理解の容易性、実装の現実性、パフォーマンスをすべて両立できます。
