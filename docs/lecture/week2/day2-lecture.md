# 1. Week 2 - Day 2: sqlboilerによるORM設定

## 1.1. 目次

- [1. Week 2 - Day 2: sqlboilerによるORM設定](#1-week-2---day-2-sqlboilerによるorm設定)
	- [1.1. 目次](#11-目次)
	- [1.2. 【要点】](#12-要点)
	- [1.3. 【準備】](#13-準備)
		- [1.3.1. チェックリスト](#131-チェックリスト)
	- [1.4. 【手順】](#14-手順)
		- [1.4.1. sqlboilerの設定ファイル作成](#141-sqlboilerの設定ファイル作成)
		- [1.4.2. sqlboilerによるモデル生成](#142-sqlboilerによるモデル生成)
		- [1.4.3. リポジトリパターンの基本構造実装](#143-リポジトリパターンの基本構造実装)
		- [1.4.4. 基本的なデータアクセス実装](#144-基本的なデータアクセス実装)
		- [1.4.5. トランザクション管理の実装](#145-トランザクション管理の実装)
		- [1.4.6. 簡単なテストコードの作成](#146-簡単なテストコードの作成)
	- [1.5. 【確認ポイント】](#15-確認ポイント)
	- [1.6. 【詳細解説】](#16-詳細解説)
		- [1.6.1. sqlboilerの仕組みと特徴](#161-sqlboilerの仕組みと特徴)
		- [1.6.2. リポジトリパターンとその利点](#162-リポジトリパターンとその利点)
		- [1.6.3. トランザクション管理の重要性](#163-トランザクション管理の重要性)
	- [1.7. 【補足情報】](#17-補足情報)
		- [1.7.1. sqlboilerとその他のORMの比較](#171-sqlboilerとその他のormの比較)
		- [1.7.2. クエリビルダーの応用テクニック](#172-クエリビルダーの応用テクニック)
	- [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
		- [1.8.1. 問題1: モデル生成時のデータベース接続エラー](#181-問題1-モデル生成時のデータベース接続エラー)
		- [1.8.2. 問題2: 生成されたモデルでNULL値の取り扱いエラー](#182-問題2-生成されたモデルでnull値の取り扱いエラー)
	- [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
	- [1.10. 【次回の準備】](#110-次回の準備)
	- [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- sqlboilerはデータベーススキーマから型安全なGoコードを生成するORMツール
- リポジトリパターンを用いたデータアクセスレイヤーの実装方法の習得
- データベーストランザクション管理の基本設計と実装方法の理解
- クエリビルダーを使った型安全なデータベース操作の実装方法の習得
- データアクセスコードのテスト方法の基本を理解

## 1.3. 【準備】

このレッスンを始める前に、以下の環境とツールが準備されていることを確認してください。

### 1.3.1. チェックリスト

- [ ] Day 1で作成したデータベーススキーマ（マイグレーション適用済み）が存在する
- [ ] sqlboilerとsqlboiler-mysqlがインストールされている
- [ ] MySQLが稼働中でアクセス可能である
- [ ] Docker ComposeのMySQLサービスが起動している
- [ ] Go言語の開発環境が構築済みである
- [ ] プロジェクトのディレクトリ構造が作成済みである

sqlboilerをまだインストールしていない場合は、以下のコマンドでインストールしてください：

```bash
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
```

## 1.4. 【手順】

### 1.4.1. sqlboilerの設定ファイル作成

プロジェクトのバックエンドディレクトリに移動し、sqlboilerの設定ファイルを作成します。

```bash
cd backend
touch sqlboiler.toml
```

sqlboiler.tomlファイルに以下の内容を記述します：

```toml
output = "internal/db/models"
no-tests = true
wipe = true
add-global-variants = true
add-panic-variants = true
no-context = false
no-hooks = false
no-auto-timestamps = false
tag-ignore = ["created_at", "updated_at", "deleted_at"]

[mysql]
  dbname = "ecommerce"
  host = "localhost"
  port = 3306
  user = "ecommerce_user"
  pass = "ecommerce_password"
  sslmode = "false"
  blacklist = ["schema_migrations"]

[[types]]
  [types.match]
    type = "types.Decimal"
  [types.replace]
    type = "decimal.Decimal"
  [types.imports]
    third_party = ['"github.com/shopspring/decimal"']
```

この設定ファイルでは：

- `output`: 生成されるモデルの出力先ディレクトリを指定
- `no-tests`: テストコードの生成を無効化（必要に応じて`false`に変更可能）
- `add-global-variants`: グローバル関数バリアントの生成を有効化
- `add-panic-variants`: パニックする関数バリアントの生成を有効化
- `mysql`: MySQL接続情報の設定
- `blacklist`: 生成対象から除外するテーブル名のリスト

### 1.4.2. sqlboilerによるモデル生成

まず、モデルが格納されるディレクトリを作成します：

```bash
mkdir -p internal/db/models
```

sqlboilerコマンドを実行してモデルを生成します：

```bash
sqlboiler mysql
```

生成されたモデルディレクトリの内容を確認してみましょう：

```bash
ls -la internal/db/models
```

出力例：

```bash
-rw-r--r-- 1 user user 12345 Jan 1 10:00 boil_queries.go
-rw-r--r-- 1 user user  5678 Jan 1 10:00 boil_table_names.go
-rw-r--r-- 1 user user  9876 Jan 1 10:00 boil_types.go
-rw-r--r-- 1 user user  3456 Jan 1 10:00 categories.go
-rw-r--r-- 1 user user  7890 Jan 1 10:00 products.go
-rw-r--r-- 1 user user  2345 Jan 1 10:00 inventory.go
```

### 1.4.3. リポジトリパターンの基本構造実装

データアクセスのためのリポジトリパターンを実装します。まずは基本構造を作成します：

```bash
mkdir -p internal/repository
touch internal/repository/repository.go
mkdir -p internal/repository/product
touch internal/repository/product/product_repository.go
```

repository.goファイルには、リポジトリの共通インターフェースとトランザクション管理を実装します：

```go
package repository

import (
 "context"
 "database/sql"

 "github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository は全てのリポジトリの基本インターフェース
type Repository interface {
 WithTx(*sql.Tx) Repository
 DB() boil.ContextExecutor
}

// RepositoryBase は基本リポジトリの共通実装
type RepositoryBase struct {
 executor boil.ContextExecutor
}

// NewRepositoryBase は新しいRepositoryBaseを作成します
func NewRepositoryBase(executor boil.ContextExecutor) RepositoryBase {
 return RepositoryBase{
  executor: executor,
 }
}

// WithTx はトランザクションを設定したリポジトリを返します
func (r RepositoryBase) WithTx(tx *sql.Tx) RepositoryBase {
 r.executor = tx
 return r
}

// DB はデータベースエグゼキュータを返します
func (r RepositoryBase) DB() boil.ContextExecutor {
 return r.executor
}

// RunInTransaction はトランザクション内で関数を実行します
func RunInTransaction(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
 tx, err := db.BeginTx(ctx, nil)
 if err != nil {
  return err
 }

 defer func() {
  if p := recover(); p != nil {
   if err := tx.Rollback(); err != nil {
    panic(err)
   }
   panic(p)
  }
 }()

 if err := fn(tx); err != nil {
  if rollbackErr := tx.Rollback(); rollbackErr != nil {
   return rollbackErr
  }
  return err
 }

 return tx.Commit()
}
```

### 1.4.4. 基本的なデータアクセス実装

product_repository.goファイルに商品リポジトリの実装を追加します：

```go
package product

import (
 "context"
 "database/sql"
 "fmt"

 "github.com/volatiletech/sqlboiler/v4/boil"
 "github.com/volatiletech/sqlboiler/v4/queries/qm"

 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/db/models"
 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/repository"
)

// Repository は商品リポジトリのインターフェース
type Repository interface {
 repository.Repository
 FindByID(ctx context.Context, id int) (*models.Product, error)
 FindAll(ctx context.Context, limit, offset int) ([]*models.Product, error)
 FindByCategory(ctx context.Context, categoryID int, limit, offset int) ([]*models.Product, error)
 Create(ctx context.Context, product *models.Product) error
 Update(ctx context.Context, product *models.Product) error
 Delete(ctx context.Context, id int) error
 Count(ctx context.Context) (int64, error)
 CountByCategory(ctx context.Context, categoryID int) (int64, error)
}

// ProductRepository は商品リポジトリの実装
type ProductRepository struct {
 repository.RepositoryBase
}

// New は新しい商品リポジトリを作成します
func New(executor boil.ContextExecutor) Repository {
 return &ProductRepository{
  RepositoryBase: repository.NewRepositoryBase(executor),
 }
}

// WithTx はトランザクションを設定した商品リポジトリを返します
func (r *ProductRepository) WithTx(tx *sql.Tx) repository.Repository {
 return &ProductRepository{
  RepositoryBase: r.RepositoryBase.WithTx(tx),
 }
}

// FindByID は指定IDの商品を取得します
func (r *ProductRepository) FindByID(ctx context.Context, id int) (*models.Product, error) {
 return models.FindProduct(ctx, r.DB(), id)
}

// FindAll は商品一覧を取得します
func (r *ProductRepository) FindAll(ctx context.Context, limit, offset int) ([]*models.Product, error) {
 mods := []qm.QueryMod{
  qm.Limit(limit),
  qm.Offset(offset),
  qm.OrderBy("created_at DESC"),
 }
 return models.Products(mods...).All(ctx, r.DB())
}

// FindByCategory は指定カテゴリーの商品一覧を取得します
func (r *ProductRepository) FindByCategory(ctx context.Context, categoryID int, limit, offset int) ([]*models.Product, error) {
 mods := []qm.QueryMod{
  qm.Where("category_id = ?", categoryID),
  qm.Limit(limit),
  qm.Offset(offset),
  qm.OrderBy("created_at DESC"),
 }
 return models.Products(mods...).All(ctx, r.DB())
}

// Create は新しい商品を作成します
func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
 return product.Insert(ctx, r.DB(), boil.Infer())
}

// Update は商品情報を更新します
func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
 _, err := product.Update(ctx, r.DB(), boil.Infer())
 return err
}

// Delete は商品を削除します
func (r *ProductRepository) Delete(ctx context.Context, id int) error {
 product, err := r.FindByID(ctx, id)
 if err != nil {
  return err
 }
 if product == nil {
  return fmt.Errorf("product not found: %d", id)
 }
 _, err = product.Delete(ctx, r.DB())
 return err
}

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

同様に、カテゴリーリポジトリも作成します：

```bash
mkdir -p internal/repository/category
touch internal/repository/category/category_repository.go
```

category_repository.goファイルの内容：

```go
package category

import (
 "context"
 "database/sql"

 "github.com/volatiletech/sqlboiler/v4/boil"
 "github.com/volatiletech/sqlboiler/v4/queries/qm"

 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/db/models"
 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/repository"
)

// Repository はカテゴリーリポジトリのインターフェース
type Repository interface {
 repository.Repository
 FindByID(ctx context.Context, id int) (*models.Category, error)
 FindAll(ctx context.Context) ([]*models.Category, error)
 Create(ctx context.Context, category *models.Category) error
 Update(ctx context.Context, category *models.Category) error
 Delete(ctx context.Context, id int) error
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
 mods := []qm.QueryMod{
  qm.OrderBy("name"),
 }
 return models.Categories(mods...).All(ctx, r.DB())
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
 _, err = category.Delete(ctx, r.DB())
 return err
}
```

### 1.4.5. トランザクション管理の実装

リポジトリ層でのトランザクション管理を示す簡単な使用例を作成します：

```bash
mkdir -p internal/service
touch internal/service/product_service.go
```

product_service.goファイルの内容：

```go
package service

import (
 "context"
 "database/sql"
 "fmt"

 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/db/models"
 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/repository"
 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/repository/product"
)

// ProductService は商品サービスの実装
type ProductService struct {
 db          *sql.DB
 productRepo product.Repository
}

// NewProductService は新しい商品サービスを作成します
func NewProductService(db *sql.DB, productRepo product.Repository) *ProductService {
 return &ProductService{
  db:          db,
  productRepo: productRepo,
 }
}

// CreateProduct は商品を作成します
func (s *ProductService) CreateProduct(ctx context.Context, p *models.Product) error {
 return s.productRepo.Create(ctx, p)
}

// TransferProductCategory は商品のカテゴリーを変更するトランザクション処理の例です
func (s *ProductService) TransferProductCategory(ctx context.Context, productID int, newCategoryID int) error {
 return repository.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {
  // トランザクション内でリポジトリを使用
  txProductRepo, ok := s.productRepo.WithTx(tx).(product.Repository)
  if !ok {
   return fmt.Errorf("failed to convert to product.Repository")
  }

  // 商品を取得
  p, err := txProductRepo.FindByID(ctx, productID)
  if err != nil {
   return err
  }

  // カテゴリーを更新
  p.CategoryID = newCategoryID

  // 商品を更新
  return txProductRepo.Update(ctx, p)
 })
}
```

### 1.4.6. 簡単なテストコードの作成

リポジトリ層のテストを実装します：

```bash
touch internal/repository/product/product_repository_test.go
```

product_repository_test.goファイルの内容：

```go
package product

import (
 "context"
 "database/sql"
 "os"
 "testing"

 "github.com/shopspring/decimal"
 "github.com/stretchr/testify/assert"
 "github.com/volatiletech/null/v8"
 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/db/models"

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
```

## 1.5. 【確認ポイント】

このDayの作業が正しく完了したことを確認するためのチェックリストです：

- [ ] sqlboiler.tomlが正しく設定されている
- [ ] sqlboilerコマンドが正常に実行され、モデルが生成されている
- [ ] 生成されたモデルに以下のファイルが含まれている
  - boil_queries.go
  - boil_table_names.go
  - categories.go
  - products.go
  - inventory.go
- [ ] リポジトリパターンの基本構造が実装されている
- [ ] プロダクトリポジトリが実装されている
- [ ] カテゴリーリポジトリが実装されている
- [ ] トランザクション管理のコードが実装されている
- [ ] テストコードが作成されている
- [ ] テストを実行して基本的なCRUD操作が動作することを確認できる

## 1.6. 【詳細解説】

### 1.6.1. sqlboilerの仕組みと特徴

sqlboilerは他のORMとは異なるアプローチを取っています。通常のORMはコードファーストであるのに対し、sqlboilerはデータベースファーストです。つまり、既存のデータベーススキーマからコードを生成します。

**主な特徴:**

1. **型安全性:** 生成されたコードは完全に型付けされており、コンパイル時にエラーを検出できます。これはRuntimeエラーを減らし、安全性を高めます。

2. **パフォーマンス:** 自動生成されたコードはパフォーマンスを最優先に設計されています。余分な変換やリフレクションを避け、効率的なSQLクエリを生成します。

3. **柔軟性:** 生成されたモデルはシンプルな構造体であり、あらゆるGoコードと容易に統合できます。また、rawクエリも実行できるためORMの制約に縛られません。

4. **完全なコード生成:** モデル、クエリビルダー、リレーション処理などのコードが全て自動生成されるため、ボイラープレートコードが削減されます。

**sqlboiler生成コードの構成:**

- **モデルファイル (e.g., products.go):** 各テーブルに対応するモデル構造体と関連メソッド
- **boil_queries.go:** クエリ実行のための共通ヘルパー関数
- **boil_table_names.go:** テーブル名の定数と関数
- **boil_types.go:** 共通の型定義

### 1.6.2. リポジトリパターンとその利点

リポジトリパターンは、データアクセスロジックをビジネスロジックから分離するために使用されるデザインパターンです。

**リポジトリパターンの利点:**

1. **関心の分離:** データアクセスロジックはリポジトリクラスにカプセル化され、ビジネスロジックから分離されます。

2. **テスト容易性:** リポジトリをモック化することで、ビジネスロジックを簡単にテストできます。

3. **コードの再利用:** 共通のデータアクセスパターンはリポジトリ間で再利用できます。

4. **メンテナンス性:** データアクセス方法が変わっても、リポジトリの実装を変更するだけでアプリケーション全体への影響を最小限に抑えることができます。

**今回の実装アプローチ:**

- インターフェースを定義して実装を切り替え可能にする
- 基本リポジトリ（RepositoryBase）を作成して共通機能を集約
- トランザクション管理を統一的に扱う仕組みを提供
- 具体的なエンティティ（商品、カテゴリーなど）ごとにリポジトリを実装

### 1.6.3. トランザクション管理の重要性

データベーストランザクションは、複数のデータベース操作を一つの論理的な単位として扱うメカニズムです。

**トランザクションの重要性:**

1. **データ整合性:** 複数の操作がすべて成功するか、すべて失敗するかのいずれかになることを保証します。中途半端な状態を防ぎます。

2. **並行制御:** 複数のユーザーが同時にデータを操作する場合に、データの整合性を保護します。

3. **障害回復:** システム障害が発生した場合にも、データベースの整合性を維持できます。

**トランザクション管理のパターン:**

今回実装した`RunInTransaction`関数は、以下のようなトランザクション管理パターンを実現しています：

1. トランザクションを開始
2. パニックが発生した場合はロールバック
3. 関数内でエラーが発生した場合はロールバック
4. 成功した場合はコミット

このパターンにより、アプリケーションコードでは次のようにシンプルにトランザクションを扱うことができます：

```go
err := repository.RunInTransaction(ctx, db, func(tx *sql.Tx) error {
    // トランザクション内の処理
    // エラーを返すとロールバックされる
    // nilを返すとコミットされる
    return nil
})
```

## 1.7. 【補足情報】

### 1.7.1. sqlboilerとその他のORMの比較

**sqlboiler vs GORM:**

| 特性             | sqlboiler                                 | GORM                          |
| ---------------- | ----------------------------------------- | ----------------------------- |
| アプローチ       | DB-first (コード生成)                     | Code-first (リフレクション)   |
| 型安全性         | 非常に高い (コンパイル時チェック)         | 中程度 (一部ランタイムエラー) |
| パフォーマンス   | 高い (生成コード、最小限のリフレクション) | 中程度 (リフレクションを多用) |
| 学習曲線         | 中程度 (生成コードの理解)                 | 緩やか (直感的API)            |
| マイグレーション | 外部ツールが必要                          | 組み込みサポート              |
| コード量         | 多い (生成コード)                         | 少ない (マジックが多い)       |

**sqlboiler vs ent:**

| 特性           | sqlboiler            | ent                       |
| -------------- | -------------------- | ------------------------- |
| アプローチ     | DB-first             | Schema-first (コード生成) |
| 型安全性       | 非常に高い           | 非常に高い                |
| クエリビルダー | シンプル             | 強力で表現力豊か          |
| スキーマ進化   | 外部ツールが必要     | 組み込みサポート          |
| グラフクエリ   | 基本的なリレーション | 高度なグラフトラバーサル  |
| コード生成量   | 中程度               | 多い                      |

### 1.7.2. クエリビルダーの応用テクニック

sqlboilerのクエリビルダーは非常に強力で、複雑なクエリも型安全に構築できます。

**応用テクニック例:**

1. **複雑な条件付きクエリ:**

    ```go
    products, err := models.Products(
        qm.Where("price > ?", 1000),
        qm.Where("stock > ?", 0),
        qm.Or("is_featured = ?", true),
        qm.OrderBy("created_at DESC"),
        qm.Limit(10),
    ).All(ctx, db)
    ```

2. **リレーションの読み込み:**

    ```go
    products, err := models.Products(
        qm.Load("Category"),
        qm.Load("Inventory"),
    ).All(ctx, db)
    ```

3. **カスタム結合:**

    ```go
    products, err := models.Products(
        qm.InnerJoin("categories c on c.id = products.category_id"),
        qm.Where("c.active = ?", true),
    ).All(ctx, db)
    ```

4. **集計関数の使用:**

    ```go
    var result struct {
        AvgPrice float64 `boil:"avg_price"`
    }

    err := models.Products(
        qm.Select("AVG(price) as avg_price"),
        qm.Where("category_id = ?", categoryID),
    ).QueryRow(ctx, db).Scan(&result.AvgPrice)
    ```

5. **GROUP BY句の使用:**

    ```go
    type CategoryCount struct {
        CategoryID int    `boil:"category_id"`
        Name       string `boil:"name"`
        Count      int    `boil:"count"`
    }

    var results []CategoryCount

    err := models.Products(
        qm.Select("category_id, categories.name, COUNT(*) as count"),
        qm.InnerJoin("categories on categories.id = products.category_id"),
        qm.GroupBy("category_id"),
        qm.OrderBy("count DESC"),
    ).Bind(ctx, db, &results)
    ```

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: モデル生成時のデータベース接続エラー

**症状**: sqlboilerコマンド実行時に以下のようなエラーが発生する

```bash
Error: unable to initialize tables: unable to fetch table data: Error 1045: Access denied for user 'root'@'172.17.0.1' (using password: YES)
```

**解決策**:

1. sqlboiler.tomlの接続情報が正しいか確認する

    ```toml
    [mysql]
      dbname = "ecommerce"
      host = "mysql"  # Docker Compose内のサービス名
      port = 3306
      user = "root"
      pass = "password"
      sslmode = "false"
    ```

2. Dockerネットワーク内でMySQLサービスにアクセスできるか確認する

    ```bash
    docker-compose exec backend sh -c "ping -c 1 mysql"
    ```

3. MySQLへの直接接続をテストする

    ```bash
    docker-compose exec backend sh -c "mysql -h mysql -u root -ppassword -e 'SHOW DATABASES;'"
    ```

4. sqlboilerコマンドをDocker内で実行することを検討する

    ```bash
    docker-compose exec backend sqlboiler mysql
    ```

### 1.8.2. 問題2: 生成されたモデルでNULL値の取り扱いエラー

**症状**: NULLが許容されるカラムで値を設定せずにInsertするとエラーが発生する

```bash
Error: pq: null value in column "description" violates not-null constraint
```

**解決策**:

1. sqlboilerはNULL許容カラム向けに専用の型を使用している。null.String, null.Int64などを利用する

    ```go
    import "github.com/volatiletech/null/v8"

    product := &models.Product{
        Name: "Test Product",
        Description: null.StringFrom("A description"), // 値を設定
        // または
        Description: null.String{}, // NULL値
    }
    ```

2. NULL許容カラムのデフォルト値をデータベース側で設定する

3. マイグレーションでNOT NULL制約を修正する（必要な場合）

4. モデル生成前に全てのカラムの制約を確認する

    ```sql
    SHOW COLUMNS FROM products;
    ```

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **データベースファーストアプローチ**: sqlboilerは既存のデータベーススキーマからコードを生成するDBファーストアプローチを採用しています。これにより、データモデルとデータベーススキーマの一貫性が保証されます。

2. **型安全なデータアクセス**: sqlboilerが生成するコードは完全に型付けされており、コンパイル時にエラーを捕捉できます。これにより、ランタイムエラーを減らしコードの安全性を高めることができます。

3. **リポジトリパターン**: データアクセスロジックをビジネスロジックから分離するリポジトリパターンを導入することで、コードの保守性とテスト容易性が向上します。

4. **トランザクション管理**: 複数のデータベース操作を一つの論理的な単位として扱うトランザクション管理を実装することで、データの整合性を確保できます。

5. **クエリビルダーの活用**: sqlboilerのクエリビルダーを活用することで、型安全かつ柔軟なデータベースクエリを構築できます。

これらのポイントは次回以降の実装でも活用されますので、よく理解しておきましょう。

## 1.10. 【次回の準備】

次回（Day 3）では、OpenAPI仕様の初期定義とogenによるコード生成を行います。以下の点について事前に確認しておくと良いでしょう：

1. OpenAPI 3.0の基本的な仕様と記法について理解する
2. ogenをインストールする（まだの場合）：

   ```bash
   go install github.com/ogen-go/ogen/cmd/ogen@latest
   ```

3. 作成したsqlboilerモデルの構造と、それをAPIで公開するためのDTO（データ転送オブジェクト）の関係について考えておく
4. REST APIの基本設計原則とリソース命名規則について復習しておく
5. OpenAPI定義からGoコードを生成する際の注意点を調査する

また、今日実装したリポジトリ層とOpenAPIで定義するAPIの関係性についても考えておくとよいでしょう。

## 1.11. 【.envrc サンプル】

以下は本日の実装で使用する.envrcのサンプルです。ご自身の環境に合わせて修正して使用してください。このファイルはgitにコミットしないようにしてください。

```bash
# .envrc サンプル
export DB_USER="root"
export DB_PASSWORD="password"
export DB_HOST="mysql"
export DB_PORT="3306"
export DB_NAME="ecommerce"
export DB_SSLMODE="false"
```

direnvがインストールされている場合は、上記の内容を`.envrc`ファイルに保存し、以下のコマンドで有効化します：

```bash
direnv allow
```

これにより、sqlboilerの設定ファイルを環境変数から動的に生成することもできます：

```toml
output = "internal/db/models"
no-tests = true
wipe = true
add-global-variants = true
add-panic-variants = true
no-context = false
no-hooks = false
no-auto-timestamps = false
tag-ignore = ["created_at", "updated_at", "deleted_at"]

[mysql]
  dbname = "${DB_NAME}"
  host = "${DB_HOST}"
  port = ${DB_PORT}
  user = "${DB_USER}"
  pass = "${DB_PASSWORD}"
  sslmode = "${DB_SSLMODE}"
  blacklist = ["schema_migrations"]
```
