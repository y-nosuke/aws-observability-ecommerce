# 1. AWSオブザーバビリティ学習用eコマースアプリ - バックエンドサービスとAPI一覧

このドキュメントでは、AWSオブザーバビリティ学習用eコマースアプリのMVPと完成版のバックエンドサービスおよびAPIエンドポイントを比較します。各サービスの目的、実装技術、エンドポイント、およびオブザーバビリティ要素と、各APIエンドポイントの詳細を整理して示します。

## 1.1. 目次

- [1. AWSオブザーバビリティ学習用eコマースアプリ - バックエンドサービスとAPI一覧](#1-awsオブザーバビリティ学習用eコマースアプリ---バックエンドサービスとapi一覧)
  - [1.1. 目次](#11-目次)
  - [1.2. バックエンドサービス一覧](#12-バックエンドサービス一覧)
    - [1.2.1. 基本サービス](#121-基本サービス)
    - [1.2.2. 認証・ユーザー管理サービス](#122-認証ユーザー管理サービス)
    - [1.2.3. システム・モニタリングサービス](#123-システムモニタリングサービス)
    - [1.2.4. イベント処理とサーバーレス関数](#124-イベント処理とサーバーレス関数)
  - [1.3. API一覧](#13-api一覧)
    - [1.3.1. 商品カタログAPI](#131-商品カタログapi)
    - [1.3.2. 在庫管理API](#132-在庫管理api)
    - [1.3.3. カートAPI](#133-カートapi)
    - [1.3.4. 注文処理API](#134-注文処理api)
    - [1.3.5. 認証API](#135-認証api)
    - [1.3.6. ユーザー管理API](#136-ユーザー管理api)
    - [1.3.7. ヘルスチェックAPI](#137-ヘルスチェックapi)
    - [1.3.8. メトリクスAPI](#138-メトリクスapi)
    - [1.3.9. 通知API](#139-通知api)
  - [1.4. マイクロサービスアーキテクチャの進化](#14-マイクロサービスアーキテクチャの進化)
    - [1.4.1. MVPのアーキテクチャ特性](#141-mvpのアーキテクチャ特性)
    - [1.4.2. 完成版のアーキテクチャ特性](#142-完成版のアーキテクチャ特性)
    - [1.4.3. サービス間連携の進化](#143-サービス間連携の進化)

## 1.2. バックエンドサービス一覧

### 1.2.1. 基本サービス

| サービスID              | サービス名           | 説明                             | 実装技術(MVP)                          | 実装技術(完成版)                             | エンドポイント                    | 対応機能(MVP)                                               | 対応機能(完成版)                           | オブザーバビリティ要素(MVP)        | オブザーバビリティ要素(完成版)                      |
| ----------------------- | -------------------- | -------------------------------- | -------------------------------------- | -------------------------------------------- | --------------------------------- | ----------------------------------------------------------- | ------------------------------------------ | ---------------------------------- | --------------------------------------------------- |
| SVC-CORE-PRODUCT-01     | 商品カタログサービス | 商品情報の取得と管理を担当       | Go/Echo, MySQL                         | Fargate (ECS), RDS (MySQL/PostgreSQL)        | `/api/products/*`                 | FEAT-CUST-BROWSE-01, FEAT-CUST-BROWSE-02, FEAT-CUST-BROWSE-03, FEAT-ADMIN-PROD-01, FEAT-ADMIN-PROD-02 | すべての商品関連機能                       | 基本的なログ、メトリクス、トレース | 完全な3本柱統合、カスタムメトリクス、異常検出       |
| SVC-CORE-INVENTORY-01   | 在庫管理サービス     | 商品の在庫状況の取得と更新を担当 | Go/Echo, MySQL                         | Fargate (ECS), RDS (MySQL), Lambda, DynamoDB | `/api/inventory/*`                | FEAT-CUST-BROWSE-07, FEAT-ADMIN-INV-01, FEAT-ADMIN-INV-03                             | すべての在庫関連機能                       | 基本的なログ、メトリクス           | 完全な3本柱統合、在庫アラート、予測分析             |
| SVC-CUST-CART-01        | カートサービス       | カート情報の管理                 | クライアントサイドのみ（LocalStorage） | Fargate (ECS), DynamoDB                      | クライアント側のみ／`api/carts/*` | FEAT-CUST-CART-01 (クライアント側実装)                              | FEAT-CUST-CART-01 (サーバー側実装)                 | フロントエンドログのみ             | サーバー側の完全な3本柱統合、カート分析メトリクス   |
| SVC-CORE-ORDER-01       | 注文処理サービス     | 注文の処理と管理を担当           | Go/Echo, MySQL (シンプル実装)          | Fargate (ECS), RDS (MySQL), SQS, SNS         | `/api/orders/*`                   | FEAT-CUST-SHOP-02 (シンプル実装)                                    | FEAT-CUST-SHOP-02, FEAT-CUST-SHOP-03, FEAT-CUST-SHOP-04, FEAT-CUST-SHOP-05 | 基本的なログ                       | 完全な3本柱統合、トランザクション監視、イベント追跡 |

### 1.2.2. 認証・ユーザー管理サービス

| サービスID              | サービス名           | 説明                               | 実装技術(MVP)    | 実装技術(完成版)                   | エンドポイント | 対応機能(MVP)                | 対応機能(完成版)     | オブザーバビリティ要素(MVP) | オブザーバビリティ要素(完成版)                 |
| ----------------------- | -------------------- | ---------------------------------- | ---------------- | ---------------------------------- | -------------- | ---------------------------- | -------------------- | --------------------------- | ---------------------------------------------- |
| SVC-CORE-AUTH-01        | 認証サービス         | 管理者認証を担当（シンプルな実装） | Go/Echo, JWTのみ | Fargate (ECS), Cognito, OAuth/OIDC | `/api/auth/*`  | 管理者ログイン（モック実装） | すべての認証関連機能 | 基本的なログ                | 認証セキュリティモニタリング、不審アクセス検知 |
| SVC-CORE-USER-01        | ユーザー管理サービス | ユーザープロフィールと設定の管理   | -                | Lambda, DynamoDB                   | `/api/users/*` | -                            | FEAT-CUST-USER-03, FEAT-CUST-USER-04 | -                           | 完全な3本柱統合、プロフィール変更追跡          |

### 1.2.3. システム・モニタリングサービス

| サービスID              | サービス名             | 説明                                 | 実装技術(MVP) | 実装技術(完成版)                  | エンドポイント         | 対応機能(MVP)                         | 対応機能(完成版)                               | オブザーバビリティ要素(MVP) | オブザーバビリティ要素(完成版)                     |
| ----------------------- | ---------------------- | ------------------------------------ | ------------- | --------------------------------- | ---------------------- | ------------------------------------- | ---------------------------------------------- | --------------------------- | -------------------------------------------------- |
| SVC-SYS-HEALTH-01       | ヘルスチェックサービス | サービス状態とデータベース接続の確認 | Go/Echo       | Fargate (ECS), Lambda             | `/api/health`          | FEAT-OBS-ALERT-01                            | FEAT-OBS-ALERT-01, FEAT-OBS-ALERT-02, FEAT-OBS-ALERT-03             | 基本的なログとステータス    | 詳細なヘルスチェック、サブシステムごとのステータス |
| SVC-OBS-METRICS-01      | メトリクスサービス     | 基本的なメトリクスの収集と報告       | Go/Echo       | Fargate (ECS), Lambda, CloudWatch | `/api/metrics`         | FEAT-OBS-METRIC-01, FEAT-OBS-METRIC-02, FEAT-OBS-METRIC-03 | すべてのメトリクス関連機能                     | 基本的なメトリクスとログ    | 高度なメトリクス分析、異常検出                     |
| SVC-CORE-NOTIFICATION-01 | 通知サービス           | 各種通知の送信と管理                 | -             | Lambda, SNS, SES                  | `/api/notifications/*` | -                                     | FEAT-CUST-NOTIF-01, FEAT-CUST-NOTIF-02, FEAT-CUST-NOTIF-03, FEAT-CUST-NOTIF-04 | -                           | 完全な3本柱統合、通知成功率モニタリング            |

### 1.2.4. イベント処理とサーバーレス関数

| サービスID             | サービス名           | 説明                         | 実装技術(MVP) | 実装技術(完成版)                    | トリガー(MVP)     | トリガー(完成版)                       | 対応機能(MVP)                | 対応機能(完成版)                   | オブザーバビリティ要素(MVP) | オブザーバビリティ要素(完成版)                           |
| ---------------------- | -------------------- | ---------------------------- | ------------- | ----------------------------------- | ----------------- | -------------------------------------- | ---------------------------- | ---------------------------------- | --------------------------- | -------------------------------------------------------- |
| SVC-CORE-IMAGE-01      | 画像処理サービス     | 商品画像の処理とリサイズ     | Lambda, S3    | Lambda, S3, Step Functions          | S3イベント        | S3イベント、API Gateway                | FEAT-ADMIN-PROD-05 (基本)    | FEAT-ADMIN-PROD-05 (拡張)         | 基本的なログ                | Lambda Insights, トレース, カスタムメトリクス            |
| SVC-CORE-BATCH-01      | バッチ処理サービス   | 在庫レポート生成             | Lambda        | Lambda, Step Functions, EventBridge | CloudWatch Events | EventBridge Rules                      | FEAT-ADMIN-BATCH-01         | FEAT-ADMIN-BATCH-01, FEAT-ADMIN-BATCH-02, FEAT-ADMIN-BATCH-03 | 基本的なログ                | 完全な3本柱統合、実行メトリクス、エラー追跡              |
| SVC-SYS-EVENT-01       | イベント処理サービス | イベント処理とサービス間連携 | -             | Lambda, SNS, SQS, EventBridge       | -                 | 各種イベント（注文確定、在庫更新など） | -                            | すべてのイベント駆動型機能         | -                           | イベントトレース、メッセージ追跡、デッドレターキュー監視 |

## 1.3. API一覧

### 1.3.1. 商品カタログAPI

| API-ID                      | エンドポイント                  | メソッド | 説明                                                | MVP | 完成版 |
| --------------------------- | ------------------------------- | -------- | --------------------------------------------------- | --- | ------ |
| API-CORE-PRODUCT-PRODUCTS-LIST   | `/api/products`                 | GET      | 商品一覧取得（ページネーション/フィルタリング対応） | ✅   | ✅      |
| API-CORE-PRODUCT-PRODUCTS-GET    | `/api/products/{id}`            | GET      | 指定IDの商品詳細取得                                | ✅   | ✅      |
| API-CORE-PRODUCT-CATEGORIES-LIST | `/api/products/categories`      | GET      | 商品カテゴリー一覧取得                              | ✅   | ✅      |
| API-CORE-PRODUCT-CATEGORY-LIST   | `/api/products/categories/{id}` | GET      | 指定カテゴリーの商品一覧取得                        | ✅   | ✅      |
| API-CORE-PRODUCT-PRODUCTS-CREATE | `/api/products`                 | POST     | 新規商品登録（管理者向け）                          | ✅   | ✅      |
| API-CORE-PRODUCT-PRODUCTS-UPDATE | `/api/products/{id}`            | PUT      | 指定IDの商品情報更新（管理者向け）                  | ✅   | ✅      |
| API-CORE-PRODUCT-PRODUCTS-DELETE | `/api/products/{id}`            | DELETE   | 指定IDの商品削除（管理者向け）                      | ✅   | ✅      |
| API-CORE-PRODUCT-SEARCH-GET      | `/api/products/search`          | GET      | 高度な商品検索                                      | ❌   | ✅      |
| API-CORE-PRODUCT-BATCH-CREATE    | `/api/products/batch`           | POST     | 商品の一括操作                                      | ❌   | ✅      |
| API-CORE-PRODUCT-EXPORT-GET      | `/api/products/export`          | GET      | 商品データのエクスポート                            | ❌   | ✅      |
| API-CORE-PRODUCT-IMPORT-CREATE   | `/api/products/import`          | POST     | 商品データのインポート                              | ❌   | ✅      |

### 1.3.2. 在庫管理API

| API-ID                             | エンドポイント                        | メソッド | 説明                               | MVP | 完成版 |
| ---------------------------------- | ------------------------------------- | -------- | ---------------------------------- | --- | ------ |
| API-CORE-INVENTORY-PRODUCT-GET     | `/api/inventory/{productId}`          | GET      | 指定商品の在庫状況取得             | ✅   | ✅      |
| API-CORE-INVENTORY-LIST            | `/api/inventory`                      | GET      | 全商品の在庫状況取得（管理者向け） | ✅   | ✅      |
| API-CORE-INVENTORY-PRODUCT-UPDATE  | `/api/inventory/{productId}`          | PUT      | 指定商品の在庫数更新（管理者向け） | ✅   | ✅      |
| API-CORE-INVENTORY-BATCH-UPDATE    | `/api/inventory/batch`                | PUT      | 在庫の一括更新                     | ❌   | ✅      |
| API-CORE-INVENTORY-HISTORY-GET     | `/api/inventory/{productId}/history`  | GET      | 在庫履歴取得                       | ❌   | ✅      |
| API-CORE-INVENTORY-ALERTS-LIST     | `/api/inventory/alerts`               | GET      | 在庫アラート一覧取得               | ❌   | ✅      |
| API-CORE-INVENTORY-SETTINGS-UPDATE | `/api/inventory/alerts/settings`      | PUT      | アラート設定更新                   | ❌   | ✅      |
| API-CORE-INVENTORY-FORECAST-GET    | `/api/inventory/{productId}/forecast` | GET      | 在庫予測データ取得                 | ❌   | ✅      |

### 1.3.3. カートAPI

| API-ID                   | エンドポイント                       | メソッド | 説明                             | MVP | 完成版 |
| ------------------------ | ------------------------------------ | -------- | -------------------------------- | --- | ------ |
| API-CUST-CART-USER-GET   | `/api/carts/{userId}`                | GET      | ユーザーのカート情報取得         | ❌   | ✅      |
| API-CUST-CART-ITEMS-ADD  | `/api/carts/{userId}/items`          | POST     | カートに商品追加                 | ❌   | ✅      |
| API-CUST-CART-ITEM-UPDATE | `/api/carts/{userId}/items/{itemId}` | PUT      | カート内商品の数量更新           | ❌   | ✅      |
| API-CUST-CART-ITEM-DELETE | `/api/carts/{userId}/items/{itemId}` | DELETE   | カートから商品削除               | ❌   | ✅      |
| API-CUST-CART-MERGE-CREATE | `/api/carts/{userId}/merge`          | POST     | 未認証カートと認証カートのマージ | ❌   | ✅      |

### 1.3.4. 注文処理API

| API-ID                       | エンドポイント               | メソッド | 説明                               | MVP | 完成版 |
| ---------------------------- | ---------------------------- | -------- | ---------------------------------- | --- | ------ |
| API-CORE-ORDER-ORDERS-CREATE | `/api/orders`                | POST     | 新規注文作成                       | ✅   | ✅      |
| API-CORE-ORDER-ORDER-GET     | `/api/orders/{id}`           | GET      | 指定IDの注文詳細取得               | ✅   | ✅      |
| API-CORE-ORDER-ORDERS-LIST   | `/api/orders`                | GET      | 注文一覧取得（フィルタリング付き） | ❌   | ✅      |
| API-CORE-ORDER-STATUS-UPDATE | `/api/orders/{id}/status`    | PUT      | 注文ステータス更新                 | ❌   | ✅      |
| API-CORE-ORDER-PAYMENT-CREATE | `/api/orders/{id}/payment`   | POST     | 支払い処理                         | ❌   | ✅      |
| API-CORE-ORDER-SHIPPING-CREATE | `/api/orders/{id}/shipping`  | POST     | 配送情報登録                       | ❌   | ✅      |
| API-CORE-ORDER-USER-ORDERS-LIST | `/api/users/{userId}/orders` | GET      | ユーザーの注文履歴取得             | ❌   | ✅      |

### 1.3.5. 認証API

| API-ID                      | エンドポイント                         | メソッド | 説明                           | MVP            | 完成版 |
| --------------------------- | -------------------------------------- | -------- | ------------------------------ | -------------- | ------ |
| API-CORE-AUTH-SESSION-CREATE | `/api/auth/login`                      | POST     | ログイン認証                   | ✅ (管理者のみ) | ✅      |
| API-CORE-AUTH-SESSION-DELETE | `/api/auth/logout`                     | POST     | ログアウト                     | ✅ (管理者のみ) | ✅      |
| API-CORE-AUTH-USER-CREATE    | `/api/auth/register`                   | POST     | ユーザー登録                   | ❌              | ✅      |
| API-CORE-AUTH-TOKEN-REFRESH  | `/api/auth/refresh`                    | POST     | トークンリフレッシュ           | ❌              | ✅      |
| API-CORE-AUTH-PASSWORD-RESET | `/api/auth/password/reset`             | POST     | パスワードリセット             | ❌              | ✅      |
| API-CORE-AUTH-SOCIAL-INIT    | `/api/auth/social/{provider}`          | GET      | ソーシャルログイン開始         | ❌              | ✅      |
| API-CORE-AUTH-SOCIAL-CALLBACK | `/api/auth/social/{provider}/callback` | GET      | ソーシャルログインコールバック | ❌              | ✅      |
| API-CORE-AUTH-MFA-SETUP      | `/api/auth/mfa/setup`                  | POST     | 多要素認証設定                 | ❌              | ✅      |
| API-CORE-AUTH-MFA-VERIFY     | `/api/auth/mfa/verify`                 | POST     | 多要素認証検証                 | ❌              | ✅      |

### 1.3.6. ユーザー管理API

| API-ID                        | エンドポイント                          | メソッド | 説明                     | MVP | 完成版 |
| ----------------------------- | --------------------------------------- | -------- | ------------------------ | --- | ------ |
| API-CORE-USER-PROFILE-GET     | `/api/users/{id}`                       | GET      | ユーザープロフィール取得 | ❌   | ✅      |
| API-CORE-USER-PROFILE-UPDATE  | `/api/users/{id}`                       | PUT      | ユーザープロフィール更新 | ❌   | ✅      |
| API-CORE-USER-ADDRESSES-LIST  | `/api/users/{id}/addresses`             | GET      | 配送先住所一覧取得       | ❌   | ✅      |
| API-CORE-USER-ADDRESS-CREATE  | `/api/users/{id}/addresses`             | POST     | 新規配送先住所追加       | ❌   | ✅      |
| API-CORE-USER-ADDRESS-UPDATE  | `/api/users/{id}/addresses/{addressId}` | PUT      | 配送先住所更新           | ❌   | ✅      |
| API-CORE-USER-ADDRESS-DELETE  | `/api/users/{id}/addresses/{addressId}` | DELETE   | 配送先住所削除           | ❌   | ✅      |
| API-CORE-USER-PREFERENCES-GET | `/api/users/{id}/preferences`           | GET      | ユーザー設定取得         | ❌   | ✅      |
| API-CORE-USER-PREFERENCES-UPDATE | `/api/users/{id}/preferences`           | PUT      | ユーザー設定更新         | ❌   | ✅      |

### 1.3.7. ヘルスチェックAPI

| API-ID                       | エンドポイント           | メソッド | 説明                         | MVP | 完成版 |
| ---------------------------- | ------------------------ | -------- | ---------------------------- | --- | ------ |
| API-SYS-HEALTH-GET           | `/api/health`            | GET      | 基本的なヘルスステータス確認 | ✅   | ✅      |
| API-SYS-HEALTH-DETAILS-GET   | `/api/health/details`    | GET      | 詳細なシステム状態確認       | ❌   | ✅      |
| API-SYS-HEALTH-SUBSYSTEMS-GET | `/api/health/subsystems` | GET      | サブシステムごとの状態確認   | ❌   | ✅      |
| API-SYS-HEALTH-DATABASE-GET  | `/api/health/database`   | GET      | データベース接続の詳細確認   | ❌   | ✅      |

### 1.3.8. メトリクスAPI

| API-ID                       | エンドポイント           | メソッド | 説明                                   | MVP | 完成版 |
| ---------------------------- | ------------------------ | -------- | -------------------------------------- | --- | ------ |
| API-OBS-METRICS-GET          | `/api/metrics`           | GET      | 基本的なアプリケーションメトリクス取得 | ✅   | ✅      |
| API-OBS-METRICS-SYSTEM-GET   | `/api/metrics/system`    | GET      | システムリソースメトリクス取得         | ❌   | ✅      |
| API-OBS-METRICS-BUSINESS-GET | `/api/metrics/business`  | GET      | ビジネスメトリクス取得                 | ❌   | ✅      |
| API-OBS-METRICS-ANOMALIES-GET | `/api/metrics/anomalies` | GET      | 異常検出結果取得                       | ❌   | ✅      |

### 1.3.9. 通知API

| API-ID                       | エンドポイント                            | メソッド | 説明                     | MVP | 完成版 |
| ---------------------------- | ----------------------------------------- | -------- | ------------------------ | --- | ------ |
| API-CORE-NOTIF-MESSAGES-CREATE | `/api/notifications/send`                 | POST     | 通知送信                 | ❌   | ✅      |
| API-CORE-NOTIF-TEMPLATES-LIST  | `/api/notifications/templates`            | GET      | 通知テンプレート一覧取得 | ❌   | ✅      |
| API-CORE-NOTIF-HISTORY-LIST    | `/api/notifications/history`              | GET      | 通知履歴取得             | ❌   | ✅      |
| API-CORE-NOTIF-PREFERENCES-GET | `/api/notifications/preferences/{userId}` | GET      | 通知設定取得             | ❌   | ✅      |
| API-CORE-NOTIF-PREFERENCES-UPDATE | `/api/notifications/preferences/{userId}` | PUT      | 通知設定更新             | ❌   | ✅      |

## 1.4. マイクロサービスアーキテクチャの進化

### 1.4.1. MVPのアーキテクチャ特性

MVPアーキテクチャは、シンプルなモノリシックアプローチに近い構成で、基本的な商品カタログとシンプルな注文処理機能に焦点を当てています。

1. **シンプルなサービス構成**:
   - 基本的な商品カタログサービス
   - シンプルな在庫管理
   - 限定的な注文処理
   - 管理者向けの基本認証
   - 基本的なヘルスチェックとメトリクス

2. **技術スタック**:
   - 主にGo/Echo + MySQLの組み合わせ
   - 一部の基本的なLambda機能（画像処理など）
   - クライアントサイドでのカート管理

3. **統合パターン**:
   - 主に同期API呼び出し
   - 限定的なイベント処理

4. **オブザーバビリティ**:
   - 基本的なログ記録
   - シンプルなメトリクス収集
   - 限定的なトレース機能

### 1.4.2. 完成版のアーキテクチャ特性

完成版アーキテクチャは、フルスケールのマイクロサービスアプローチに進化し、高度なイベント駆動型のパターンを採用しています。

1. **包括的なサービス構成**:
   - 商品カタログの拡張機能
   - 高度な在庫管理と予測
   - 完全な注文処理と支払い連携
   - 包括的なユーザー管理と認証
   - 通知とイベント処理
   - 詳細なモニタリングサービス

2. **技術スタック**:
   - Fargate (ECS)とLambdaの適材適所での使用
   - RDS(MySQL)とDynamoDBの併用
   - SNS、SQS、EventBridgeによるメッセージング
   - Cognitoとの認証統合

3. **統合パターン**:
   - イベント駆動型アーキテクチャ
   - 非同期メッセージング
   - APIゲートウェイによる統合

4. **オブザーバビリティ**:
   - 包括的な3本柱（ログ、メトリクス、トレース）統合
   - サービス間のトレース連鎖
   - 詳細なビジネスメトリクス
   - 異常検出と自動アラート
   - イベント追跡

### 1.4.3. サービス間連携の進化

1. **MVPでの連携**:
   - 主に直接のAPI呼び出し
   - データベースを介した間接的な連携
   - シンプルなイベントトリガー（S3イベントなど）

2. **完成版での連携**:
   - EventBridgeによるイベントバス
   - SNSによるパブリッシュ/サブスクライブパターン
   - SQSによるメッセージキュー
   - Step Functionsによるワークフロー調整
   - APIゲートウェイによる統合レイヤー
