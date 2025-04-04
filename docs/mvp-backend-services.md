# 1. MVPバックエンドサービス一覧

| サービスID | サービス名             | 説明                                               | エンドポイント     | 対応機能                                                    |
| ---------- | ---------------------- | -------------------------------------------------- | ------------------ | ----------------------------------------------------------- |
| SVC-01     | 商品カタログサービス   | 商品情報の取得と管理を担当                         | `/api/products/*`  | C-BROWSE-01, C-BROWSE-02, C-BROWSE-03, A-PROD-01, A-PROD-02 |
| SVC-02     | 在庫管理サービス       | 商品の在庫状況の取得と更新を担当                   | `/api/inventory/*` | C-BROWSE-07, A-INV-01, A-INV-03                             |
| SVC-03     | カートサービス         | カート情報の管理（一時的なクライアントサイド実装） | クライアント側のみ | C-SHOP-01                                                   |
| SVC-04     | 注文処理サービス       | 注文の処理と管理を担当                             | `/api/orders/*`    | C-SHOP-02                                                   |
| SVC-05     | 認証サービス           | 管理者認証を担当（シンプルな実装）                 | `/api/auth/*`      | -                                                           |
| SVC-06     | ヘルスチェックサービス | サービス状態とデータベース接続の確認               | `/api/health`      | O-ALERT-01                                                  |
| SVC-07     | メトリクスサービス     | 基本的なメトリクスの収集と報告                     | `/api/metrics`     | O-METRIC-01, O-METRIC-02, O-METRIC-03                       |

## 1.1. API エンドポイント詳細

### 1.1.1. 商品カタログサービス

| API-ID         | エンドポイント                  | メソッド | 説明                                     |
| -------------- | ------------------------------- | -------- | ---------------------------------------- |
| API-PRODUCT-01 | `/api/products`                 | GET      | 全商品または絞り込まれた商品の一覧を取得 |
| API-PRODUCT-02 | `/api/products/{id}`            | GET      | 指定IDの商品詳細を取得                   |
| API-PRODUCT-03 | `/api/products/categories`      | GET      | 商品カテゴリー一覧を取得                 |
| API-PRODUCT-04 | `/api/products/categories/{id}` | GET      | 指定カテゴリーの商品一覧を取得           |
| API-PRODUCT-05 | `/api/products`                 | POST     | 新規商品を登録（管理者向け）             |
| API-PRODUCT-06 | `/api/products/{id}`            | PUT      | 指定IDの商品情報を更新（管理者向け）     |
| API-PRODUCT-07 | `/api/products/{id}`            | DELETE   | 指定IDの商品を削除（管理者向け）         |

### 1.1.2. 在庫管理サービス

| API-ID           | エンドポイント               | メソッド | 説明                                 |
| ---------------- | ---------------------------- | -------- | ------------------------------------ |
| API-INVENTORY-01 | `/api/inventory/{productId}` | GET      | 指定商品の在庫状況を取得             |
| API-INVENTORY-02 | `/api/inventory`             | GET      | 全商品の在庫状況を取得（管理者向け） |
| API-INVENTORY-03 | `/api/inventory/{productId}` | PUT      | 指定商品の在庫数を更新（管理者向け） |

### 1.1.3. 注文処理サービス

| API-ID       | エンドポイント     | メソッド | 説明                             |
| ------------ | ------------------ | -------- | -------------------------------- |
| API-ORDER-01 | `/api/orders`      | POST     | 新規注文を作成                   |
| API-ORDER-02 | `/api/orders/{id}` | GET      | 指定IDの注文詳細を取得           |
| API-ORDER-03 | `/api/orders`      | GET      | 全注文の一覧を取得（管理者向け） |

### 1.1.4. 認証サービス

| API-ID      | エンドポイント     | メソッド | 説明               |
| ----------- | ------------------ | -------- | ------------------ |
| API-AUTH-01 | `/api/auth/login`  | POST     | 管理者ログイン認証 |
| API-AUTH-02 | `/api/auth/logout` | POST     | 管理者ログアウト   |

### 1.1.5. ヘルスチェックサービス

| API-ID        | エンドポイント | メソッド | 説明                             |
| ------------- | -------------- | -------- | -------------------------------- |
| API-HEALTH-01 | `/api/health`  | GET      | サービス状態とDB接続の状態を確認 |

### 1.1.6. メトリクスサービス

| API-ID         | エンドポイント | メソッド | 説明                             |
| -------------- | -------------- | -------- | -------------------------------- |
| API-METRICS-01 | `/api/metrics` | GET      | アプリケーションのメトリクス取得 |
