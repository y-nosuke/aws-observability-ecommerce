# Google Wireによる依存性注入（DI）ガイドライン

このドキュメントは、Google Wireを用いた依存性注入（DI）の設計・実装における考え方・ベストプラクティスをまとめたものです。

## 1. ファイル構成と責務

### ディレクトリ構成

```text
backend-api/di/
├── wire.go           # Wireの指示書（injector関数定義）
├── container.go      # アプリケーションコンテナ
├── wire_gen.go       # Wireが自動生成（git管理対象）
└── provider/         # Wire専用のProvider関数
    ├── shared_provider.go    # 共通インフラ関連
    ├── product_provider.go   # 商品ドメイン関連
    ├── query_provider.go     # クエリ（Read）関連
    └── system_provider.go    # システム関連
```

### 各ファイルの責務

#### 1.1. wire.go - Wireの指示書

```go
//go:build wireinject
// +build wireinject

package di

import (
    "database/sql"
    "github.com/google/wire"
    // ...imports
)

// InitializeAppContainer は全体のアプリケーションコンテナを初期化
func InitializeAppContainer(
    db *sql.DB,
    awsServiceRegistry *aws.ServiceRegistry,
    logger logging.Logger,
) (*AppContainer, error) {
    wire.Build(
        // Provider sets
        SharedProviderSet,
        ProductProviderSet,
        QueryProviderSet,
        SystemProviderSet,

        // Container構築
        NewAppContainer,
    )
    return &AppContainer{}, nil
}
```

#### 1.2. container.go - コンテナ構造体

```go
package di

// AppContainer はアプリケーション全体の依存関係を管理
type AppContainer struct {
    DB                    *sql.DB
    AWSServiceRegistry    *aws.ServiceRegistry
    Logger                logging.Logger
    ProductHandler        *handler.ProductHandler
    // ...その他のハンドラー
}

// NewAppContainer は新しいAppContainerを作成
func NewAppContainer(
    db *sql.DB,
    awsServiceRegistry *aws.ServiceRegistry,
    logger logging.Logger,
    productHandler *handler.ProductHandler,
    // ...その他のパラメータ
) *AppContainer {
    return &AppContainer{
        DB:                 db,
        AWSServiceRegistry: awsServiceRegistry,
        Logger:             logger,
        ProductHandler:     productHandler,
        // ...
    }
}
```

#### 1.3. provider/ - Wire専用のProvider関数

```go
// di/provider/shared_provider.go
package provider

import "github.com/google/wire"

// SharedProviderSet は共通インフラのProvider Set
var SharedProviderSet = wire.NewSet(
    ProvideLogger,
    ProvideS3ClientWrapper,
    ProvideServiceRegistry,
)

// ProvideLogger は設定からLoggerを作成
func ProvideLogger(appConfig *config.AppConfig) logging.Logger {
    logConfig := logging.LogConfig{
        Level:      appConfig.Log.Level,
        Format:     appConfig.Log.Format,
        OutputPath: appConfig.Log.OutputPath,
    }
    return logging.NewLogger(logConfig) // 各パッケージのコンストラクタを使用
}
```

## 2. Provider関数の実装場所判断基準

### 2.1. 各パッケージに実装（推奨）

以下の場合は、型定義と同じパッケージにコンストラクタを実装：

```go
// ✅ pkg/logger/logger.go
func NewLogger(config LogConfig) Logger {
    return &LoggerImpl{config: config}
}

// ✅ internal/product/application/usecase/upload_product_image.go
func NewUploadProductImageUseCase(
    imageStorage domain.ImageStorage,
    bucketPrefix string,
) *UploadProductImageUseCase {
    return &UploadProductImageUseCase{
        imageStorage: imageStorage,
        bucketPrefix: bucketPrefix,
    }
}
```

**判断基準:**

- シンプルなコンストラクタ（引数をそのまま代入）
- 設定値の変換が不要
- ビジネスロジックを含まない

### 2.2. di/provider/に実装

以下の場合は、`di/provider/`にProvider関数を実装：

```go
// ✅ di/provider/product_provider.go
func ProvideUploadProductImageUseCase(
    imageStorage domain.ImageStorage,
    appConfig *config.AppConfig,
) *usecase.UploadProductImageUseCase {
    // 設定値から必要な値を抽出・変換
    bucketPrefix := appConfig.AWS.S3.ProductImagePrefix
    return usecase.NewUploadProductImageUseCase(imageStorage, bucketPrefix)
}

func ProvideS3ImageStorage(
    s3Wrapper *aws.S3ClientWrapper,
    appConfig *config.AppConfig,
) domain.ImageStorage {
    return storage.NewS3ImageStorageImpl(
        s3Wrapper,
        appConfig.AWS.S3.BucketName, // 設定値を注入
    )
}
```

**判断基準:**

- 複数の設定値から必要な値を抽出・変換
- 環境ごとの切り替えロジック
- 複数の依存関係の組み合わせ
- インターフェースと実装の結び付け

### 2.3. 外部パッケージのラッパー

```go
// ✅ di/provider/shared_provider.go
func ProvideDatabaseConnection(config *config.DatabaseConfig) (*sql.DB, error) {
    // 外部パッケージのsql.Openをラップ
    return sql.Open(config.Driver, config.DSN)
}

func ProvideS3ClientWrapper(awsConfig *config.AWSConfig) *aws.S3ClientWrapper {
    // AWS SDKをラップして設定を注入
    cfg := aws.Config{
        Region: awsConfig.Region,
        // ...その他の設定
    }
    return aws.NewS3ClientWrapper(cfg)
}
```

## 3. Provider Setの定義と管理

### 3.1. ドメインごとのProvider Set

```go
// di/provider/product_provider.go
var ProductProviderSet = wire.NewSet(
    // UseCase層
    usecase.NewUploadProductImageUseCase,
    usecase.NewGetProductImageUseCase,

    // Infrastructure層
    ProvideS3ImageStorage, // 設定値注入のためProvider関数使用

    // Presentation層
    handler.NewProductHandler,

    // インターフェースバインディング
    wire.Bind(new(domain.ImageStorage), new(*storage.S3ImageStorageImpl)),
)
```

### 3.2. 共通インフラのProvider Set

```go
// di/provider/shared_provider.go
var SharedProviderSet = wire.NewSet(
    ProvideLogger,           // 設定値注入
    ProvideS3ClientWrapper,  // AWS設定注入
    ProvideServiceRegistry,  // 複数の依存関係組み合わせ
)
```

## 4. ベストプラクティス

### 4.1. インターフェースバインディング

```go
// インターフェースと実装の結び付け
wire.Bind(new(domain.ImageStorage), new(*storage.S3ImageStorageImpl))
```

### 4.2. 設定値の注入パターン

```go
// ❌ 各パッケージで設定を直接参照
func NewSomeService() *SomeService {
    config := config.GetAppConfig() // グローバル変数への依存
    return &SomeService{endpoint: config.ExternalAPI.Endpoint}
}

// ✅ DIで設定値を注入
func ProvideSomeService(appConfig *config.AppConfig) *SomeService {
    return NewSomeService(appConfig.ExternalAPI.Endpoint)
}

func NewSomeService(endpoint string) *SomeService {
    return &SomeService{endpoint: endpoint}
}
```

### 4.3. テスタビリティの確保

```go
// 各パッケージのコンストラクタはDIに依存しない
func NewProductHandler(
    uploadUseCase *usecase.UploadProductImageUseCase,
    getUseCase *usecase.GetProductImageUseCase,
    logger logging.Logger,
) *ProductHandler {
    return &ProductHandler{
        uploadUseCase: uploadUseCase,
        getUseCase:    getUseCase,
        logger:        logger,
    }
}
```

## 5. wire_gen.goの生成

### 5.1. 生成コマンド

```bash
cd backend-api/di
go generate
```

### 5.2. go generateディレクティブ

```go
// wire.go の先頭に追加
//go:generate go run -mod=mod github.com/google/wire/cmd/wire
```

## 6. 利点とメリット

### 6.1. 保守性

- **通常のGo慣習を維持**: 基本的なコンストラクタは各パッケージに
- **責務の明確化**: DI固有のロジックは`di/provider/`に分離

### 6.2. テスタビリティ

- **DIなしでもテスト可能**: 各パッケージのコンストラクタは独立
- **モックの作成が容易**: インターフェースを適切に活用

### 6.3. 柔軟性

- **設定値の注入**: 環境ごとの切り替えが容易
- **複雑な依存関係**: 段階的な組み立てが可能

## 7. 注意点

### 7.1. wire_gen.goの管理

- **Gitに含める**: 自動生成ファイルだがコミット対象
- **レビュー対象**: 依存関係の変更を追跡可能

### 7.2.循環依存の回避

- **レイヤー間の依存方向を守る**: Domain ← Application ← Infrastructure ← Presentation
- **インターフェースを活用**: 依存関係逆転の原則

### 7.3. 過度な抽象化を避ける

- **必要最小限のProvider関数**: シンプルなコンストラクタは各パッケージに
- **実用性を重視**: 理論よりも実装の容易さを優先
