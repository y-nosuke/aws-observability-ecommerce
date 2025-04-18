# 1. AWSオブザーバビリティ学習用eコマースアプリ - データモデル一覧

このドキュメントでは、MVPと完成版で利用されるデータモデルを比較し、各エンティティの目的、主な属性、関連性、およびオブザーバビリティとの関連性を示します。データストアのタイプによって分類して整理しています。

## 1.1. 目次

- [1. AWSオブザーバビリティ学習用eコマースアプリ - データモデル一覧](#1-awsオブザーバビリティ学習用eコマースアプリ---データモデル一覧)
  - [1.1. 目次](#11-目次)
  - [1.2. リレーショナルデータベース(MySQL/RDS)モデル](#12-リレーショナルデータベースmysqlrdsモデル)
    - [1.2.1. 商品関連エンティティ](#121-商品関連エンティティ)
      - [1.2.1.1. Product (商品)](#1211-product-商品)
      - [1.2.1.2. Category (カテゴリー)](#1212-category-カテゴリー)
      - [1.2.1.3. Inventory (在庫)](#1213-inventory-在庫)
      - [1.2.1.4. Brand (ブランド)](#1214-brand-ブランド)
    - [1.2.2. 注文関連エンティティ](#122-注文関連エンティティ)
      - [1.2.2.1. Order (注文)](#1221-order-注文)
      - [1.2.2.2. OrderItem (注文アイテム)](#1222-orderitem-注文アイテム)
      - [1.2.2.3. User (ユーザー)](#1223-user-ユーザー)
      - [1.2.2.4. Address (住所)](#1224-address-住所)
    - [1.2.3. 管理者関連エンティティ](#123-管理者関連エンティティ)
      - [1.2.3.1. AdminUser (管理者ユーザー)](#1231-adminuser-管理者ユーザー)
      - [1.2.3.2. Role (ロール)](#1232-role-ロール)
      - [1.2.3.3. AuditLog (監査ログ)](#1233-auditlog-監査ログ)
  - [1.3. DynamoDB (NoSQL) モデル](#13-dynamodb-nosql-モデル)
    - [1.3.1. Cart (カート)](#131-cart-カート)
    - [1.3.2. UserSessions (ユーザーセッション)](#132-usersessions-ユーザーセッション)
    - [1.3.3. Inventory (動的在庫状態)](#133-inventory-動的在庫状態)
    - [1.3.4. StockAlerts (在庫通知)](#134-stockalerts-在庫通知)
    - [1.3.5. ProductViews (商品閲覧履歴)](#135-productviews-商品閲覧履歴)
    - [1.3.6. OrderStatus (注文状態)](#136-orderstatus-注文状態)
    - [1.3.7. UserActivity (ユーザー活動)](#137-useractivity-ユーザー活動)
  - [1.4. 主要なリレーションシップ](#14-主要なリレーションシップ)
    - [1.4.1. 完成版のリレーションシップ](#141-完成版のリレーションシップ)
  - [1.5. データストア選択の根拠](#15-データストア選択の根拠)
    - [1.5.1. RDSの選択理由（リレーショナルデータベース）](#151-rdsの選択理由リレーショナルデータベース)
    - [1.5.2. DynamoDBの選択理由（NoSQLデータベース）](#152-dynamodbの選択理由nosqlデータベース)
  - [1.6. MVPと完成版のデータモデル比較](#16-mvpと完成版のデータモデル比較)
    - [1.6.1. MVPのデータモデル特性](#161-mvpのデータモデル特性)
    - [1.6.2. 完成版のデータモデル特性](#162-完成版のデータモデル特性)
    - [1.6.3. データモデル進化のポイント](#163-データモデル進化のポイント)
  - [1.7. データアクセスパターンと最適化](#17-データアクセスパターンと最適化)
    - [1.7.1. MVPのデータアクセスパターン](#171-mvpのデータアクセスパターン)
    - [1.7.2. 完成版のデータアクセスパターン](#172-完成版のデータアクセスパターン)

## 1.2. リレーショナルデータベース(MySQL/RDS)モデル

### 1.2.1. 商品関連エンティティ

#### 1.2.1.1. Product (商品)

| 項目                       | MVP                                                                          | 完成版                                                                     |
| -------------------------- | ---------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| **テーブル名**             | products                                                                     | products                                                                   |
| **説明**                   | 商品情報を格納                                                               | 商品情報を格納（拡張属性付き）                                             |
| **主要属性**               | id, name, description, price, image_url, category_id, created_at, updated_at | MVPの属性に加えて: brand_id, tags, attributes (JSON), seo_keywords, status |
| **制約**                   | PK(id), FK(category_id)                                                      | PK(id), FK(category_id, brand_id)                                          |
| **インデックス**           | category_id, name                                                            | category_id, brand_id, name, status, tags                                  |
| **オブザーバビリティ関連** | 商品閲覧数のメトリクス                                                       | 商品閲覧数、コンバージョン率、検索ヒット率のメトリクス                     |

#### 1.2.1.2. Category (カテゴリー)

| 項目                       | MVP                                                      | 完成版                                                              |
| -------------------------- | -------------------------------------------------------- | ------------------------------------------------------------------- |
| **テーブル名**             | categories                                               | categories                                                          |
| **説明**                   | 商品カテゴリー情報を格納                                 | 商品カテゴリー情報を格納（階層構造対応）                            |
| **主要属性**               | id, name, description, parent_id, created_at, updated_at | MVPの属性に加えて: level, path, image_url, is_active, display_order |
| **制約**                   | PK(id), FK(parent_id自己参照・NULL許容)                  | PK(id), FK(parent_id自己参照・NULL許容)                             |
| **インデックス**           | parent_id, name                                          | parent_id, path, name, is_active                                    |
| **オブザーバビリティ関連** | カテゴリー別アクセス数                                   | カテゴリー別アクセス数、コンバージョン率、トラフィックパターン      |

#### 1.2.1.3. Inventory (在庫)

| 項目                       | MVP                                        | 完成版                                                                                        |
| -------------------------- | ------------------------------------------ | --------------------------------------------------------------------------------------------- |
| **テーブル名**             | inventory                                  | inventory                                                                                     |
| **説明**                   | 在庫情報を格納                             | 在庫情報を格納（履歴追跡機能付き）                                                            |
| **主要属性**               | id, product_id, quantity, last_updated     | MVPの属性に加えて: reserved_quantity, reorder_point, reorder_quantity, warehouse_id, location |
| **制約**                   | PK(id), FK(product_id), NOT NULL(quantity) | PK(id), FK(product_id, warehouse_id), NOT NULL(quantity)                                      |
| **インデックス**           | product_id                                 | product_id, warehouse_id, quantity                                                            |
| **オブザーバビリティ関連** | 在庫レベルのメトリクス                     | 在庫レベル、在庫回転率、品切れ時間のメトリクス、在庫アラート                                  |

#### 1.2.1.4. Brand (ブランド)

| 項目                       | MVP | 完成版                                                           |
| -------------------------- | --- | ---------------------------------------------------------------- |
| **テーブル名**             | -   | brands                                                           |
| **説明**                   | -   | 商品ブランド情報を格納                                           |
| **主要属性**               | -   | id, name, description, logo_url, website, created_at, updated_at |
| **制約**                   | -   | PK(id), UNIQUE(name)                                             |
| **インデックス**           | -   | name                                                             |
| **オブザーバビリティ関連** | -   | ブランド別アクセス数、コンバージョン率                           |

### 1.2.2. 注文関連エンティティ

#### 1.2.2.1. Order (注文)

| 項目                       | MVP                                                                             | 完成版                                                                                                                                             |
| -------------------------- | ------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| **テーブル名**             | orders                                                                          | orders                                                                                                                                             |
| **説明**                   | 注文情報を格納                                                                  | 注文情報を格納（拡張機能付き）                                                                                                                     |
| **主要属性**               | id, customer_name, email, address, total_amount, status, created_at, updated_at | MVPの属性に加えて: user_id, shipping_address_id, billing_address_id, payment_method, shipping_method, discount_amount, tax_amount, tracking_number |
| **制約**                   | PK(id), NOT NULL(customer_name, email, total_amount)                            | PK(id), FK(user_id, shipping_address_id, billing_address_id), NOT NULL(status, total_amount)                                                       |
| **インデックス**           | email, created_at                                                               | user_id, status, created_at                                                                                                                        |
| **オブザーバビリティ関連** | 注文数メトリクス                                                                | 注文完了率、処理時間、平均注文額、ステータス別分布のメトリクス                                                                                     |

#### 1.2.2.2. OrderItem (注文アイテム)

| 項目                       | MVP                                                         | 完成版                                                            |
| -------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------------- |
| **テーブル名**             | order_items                                                 | order_items                                                       |
| **説明**                   | 注文内の商品アイテム情報を格納                              | 注文内の商品アイテム情報を格納（拡張属性付き）                    |
| **主要属性**               | id, order_id, product_id, quantity, price, subtotal         | MVPの属性に加えて: sku, item_discount, tax_amount, options (JSON) |
| **制約**                   | PK(id), FK(order_id, product_id), NOT NULL(quantity, price) | PK(id), FK(order_id, product_id), NOT NULL(quantity, price)       |
| **インデックス**           | order_id, product_id                                        | order_id, product_id                                              |
| **オブザーバビリティ関連** | 商品別注文数                                                | 商品別注文数、バンドル率、キャンセル率のメトリクス                |

#### 1.2.2.3. User (ユーザー)

| 項目                       | MVP | 完成版                                                                                             |
| -------------------------- | --- | -------------------------------------------------------------------------------------------------- |
| **テーブル名**             | -   | users                                                                                              |
| **説明**                   | -   | ユーザー情報を格納                                                                                 |
| **主要属性**               | -   | id, email, password_hash, first_name, last_name, phone, status, created_at, updated_at, last_login |
| **制約**                   | -   | PK(id), UNIQUE(email)                                                                              |
| **インデックス**           | -   | email, status, created_at                                                                          |
| **オブザーバビリティ関連** | -   | ログイン成功率、アカウント作成完了率、アクティブユーザー率のメトリクス                             |

#### 1.2.2.4. Address (住所)

| 項目                       | MVP | 完成版                                                                                                               |
| -------------------------- | --- | -------------------------------------------------------------------------------------------------------------------- |
| **テーブル名**             | -   | addresses                                                                                                            |
| **説明**                   | -   | 配送先・請求先住所情報を格納                                                                                         |
| **主要属性**               | -   | id, user_id, address_type, first_name, last_name, line1, line2, city, state, country, postal_code, phone, is_default |
| **制約**                   | -   | PK(id), FK(user_id)                                                                                                  |
| **インデックス**           | -   | user_id, address_type, is_default                                                                                    |
| **オブザーバビリティ関連** | -   | 住所検証エラー率、デフォルト住所変更率のメトリクス                                                                   |

### 1.2.3. 管理者関連エンティティ

#### 1.2.3.1. AdminUser (管理者ユーザー)

| 項目                       | MVP                                                        | 完成版                                                                                       |
| -------------------------- | ---------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| **テーブル名**             | admin_users                                                | admin_users                                                                                  |
| **説明**                   | 管理者ユーザー情報を格納                                   | 管理者ユーザー情報を格納（権限管理機能付き）                                                 |
| **主要属性**               | id, username, password_hash, email, last_login, created_at | MVPの属性に加えて: role_id, status, mfa_enabled, failed_login_attempts, last_password_change |
| **制約**                   | PK(id), UNIQUE(username, email)                            | PK(id), FK(role_id), UNIQUE(username, email)                                                 |
| **インデックス**           | username, email                                            | username, email, role_id, status                                                             |
| **オブザーバビリティ関連** | 基本的なログイン記録                                       | ログイン成功率、ロール別操作分布、MFA使用率、不審アクセスパターンのメトリクス                |

#### 1.2.3.2. Role (ロール)

| 項目                       | MVP | 完成版                                                            |
| -------------------------- | --- | ----------------------------------------------------------------- |
| **テーブル名**             | -   | roles                                                             |
| **説明**                   | -   | 管理者ロール情報を格納                                            |
| **主要属性**               | -   | id, name, description, permissions (JSON), created_at, updated_at |
| **制約**                   | -   | PK(id), UNIQUE(name)                                              |
| **インデックス**           | -   | name                                                              |
| **オブザーバビリティ関連** | -   | ロール別アクセス分布、権限変更頻度のメトリクス、特権操作の監査    |

#### 1.2.3.3. AuditLog (監査ログ)

| 項目                       | MVP | 完成版                                                                                                                      |
| -------------------------- | --- | --------------------------------------------------------------------------------------------------------------------------- |
| **テーブル名**             | -   | audit_logs                                                                                                                  |
| **説明**                   | -   | 管理操作の監査記録を格納                                                                                                    |
| **主要属性**               | -   | id, admin_user_id, action, entity_type, entity_id, old_values (JSON), new_values (JSON), ip_address, user_agent, created_at |
| **制約**                   | -   | PK(id), FK(admin_user_id)                                                                                                   |
| **インデックス**           | -   | admin_user_id, action, entity_type, entity_id, created_at                                                                   |
| **オブザーバビリティ関連** | -   | 操作種類分布、頻度、時間帯、異常パターンの検出、セキュリティメトリクス                                                      |

## 1.3. DynamoDB (NoSQL) モデル

### 1.3.1. Cart (カート)

| 項目                       | MVP                                    | 完成版                                                   |
| -------------------------- | -------------------------------------- | -------------------------------------------------------- |
| **テーブル名**             | -                                      | Cart                                                     |
| **キー構造**               | -                                      | PK: user_id                                              |
| **説明**                   | クライアントサイドのみ（LocalStorage） | ユーザーのショッピングカート情報                         |
| **主要属性**               | -                                      | product_id, quantity, added_at                           |
| **インデックス**           | -                                      | GSI: product_id-index                                    |
| **オブザーバビリティ関連** | -                                      | カート追加率、放棄率、カート内アイテム数、セッション分析 |

### 1.3.2. UserSessions (ユーザーセッション)

| 項目                       | MVP | 完成版                                                               |
| -------------------------- | --- | -------------------------------------------------------------------- |
| **テーブル名**             | -   | UserSessions                                                         |
| **キー構造**               | -   | PK: session_id                                                       |
| **説明**                   | -   | ユーザーのセッション情報                                             |
| **主要属性**               | -   | user_id, expires_at, attributes, device_info                         |
| **インデックス**           | -   | GSI: user_id-index                                                   |
| **オブザーバビリティ関連** | -   | セッション長、デバイス分布、地理的分布、セッション活性度のメトリクス |

### 1.3.3. Inventory (動的在庫状態)

| 項目                       | MVP | 完成版                                                 |
| -------------------------- | --- | ------------------------------------------------------ |
| **テーブル名**             | -   | Inventory                                              |
| **キー構造**               | -   | PK: product_id                                         |
| **説明**                   | -   | 商品の現在の在庫状況                                   |
| **主要属性**               | -   | quantity, reserved, last_updated                       |
| **インデックス**           | -   | -                                                      |
| **オブザーバビリティ関連** | -   | リアルタイム在庫レベル、在庫変動率、予約率のメトリクス |

### 1.3.4. StockAlerts (在庫通知)

| 項目                       | MVP | 完成版                                             |
| -------------------------- | --- | -------------------------------------------------- |
| **テーブル名**             | -   | StockAlerts                                        |
| **キー構造**               | -   | PK: product_id                                     |
| **説明**                   | -   | 在庫切れ商品の入荷通知登録者リスト                 |
| **主要属性**               | -   | users_list, notification_type                      |
| **インデックス**           | -   | -                                                  |
| **オブザーバビリティ関連** | -   | 通知登録率、通知後のコンバージョン率、人気商品指標 |

### 1.3.5. ProductViews (商品閲覧履歴)

| 項目                       | MVP | 完成版                                           |
| -------------------------- | --- | ------------------------------------------------ |
| **テーブル名**             | -   | ProductViews                                     |
| **キー構造**               | -   | PK: user_id, SK: timestamp                       |
| **説明**                   | -   | ユーザーの商品閲覧履歴                           |
| **主要属性**               | -   | product_id, session_id                           |
| **インデックス**           | -   | GSI: product_id-index                            |
| **オブザーバビリティ関連** | -   | 閲覧パターン、再訪問率、推薦効果測定のメトリクス |

### 1.3.6. OrderStatus (注文状態)

| 項目                       | MVP | 完成版                                                       |
| -------------------------- | --- | ------------------------------------------------------------ |
| **テーブル名**             | -   | OrderStatus                                                  |
| **キー構造**               | -   | PK: order_id                                                 |
| **説明**                   | -   | 注文のステータス履歴                                         |
| **主要属性**               | -   | current_status, status_history, updated_at                   |
| **インデックス**           | -   | GSI: status-updated_at-index                                 |
| **オブザーバビリティ関連** | -   | ステータス変更時間、プロセス効率性、異常遅延検出のメトリクス |

### 1.3.7. UserActivity (ユーザー活動)

| 項目                       | MVP | 完成版                                                   |
| -------------------------- | --- | -------------------------------------------------------- |
| **テーブル名**             | -   | UserActivity                                             |
| **キー構造**               | -   | PK: user_id, SK: activity_type#timestamp                 |
| **説明**                   | -   | ユーザーのサイト内活動記録                               |
| **主要属性**               | -   | activity_details, context                                |
| **インデックス**           | -   | GSI: activity_type-index                                 |
| **オブザーバビリティ関連** | -   | ユーザー行動分析、コンバージョンパス分析、セグメント分析 |

## 1.4. 主要なリレーションシップ

### 1.4.1. 完成版のリレーションシップ

| 主エンティティ | 関係 | 従エンティティ | 説明                                                 |
| -------------- | ---- | -------------- | ---------------------------------------------------- |
| Categories     | 1:N  | Products       | 1つのカテゴリーに複数の商品が属する                  |
| Categories     | 1:N  | Categories     | 親カテゴリーに複数の子カテゴリーが属する（自己参照） |
| Brands         | 1:N  | Products       | 1つのブランドに複数の商品が属する                    |
| Products       | 1:1  | Inventory      | 1つの商品に対して1つの在庫情報                       |
| Products       | 1:N  | OrderItems     | 1つの商品は複数の注文アイテムとなりうる              |
| Users          | 1:N  | Orders         | 1人のユーザーが複数の注文を行う                      |
| Users          | 1:N  | Addresses      | 1人のユーザーが複数の住所を持つ                      |
| Users          | 1:1  | Cart           | 1人のユーザーに1つのカート                           |
| Orders         | 1:N  | OrderItems     | 1つの注文に複数の注文アイテムが含まれる              |
| Roles          | 1:N  | AdminUsers     | 1つのロールに複数の管理者ユーザーが属する            |
| AdminUsers     | 1:N  | AuditLogs      | 1人の管理者ユーザーが複数の監査ログを生成            |

## 1.5. データストア選択の根拠

### 1.5.1. RDSの選択理由（リレーショナルデータベース）

以下のデータには、RDS (MySQL/PostgreSQL) を利用することが適しています：

1. **強い一貫性と参照整合性が必要なデータ**
   - 商品、カテゴリー、ブランド間の関連
   - 注文と注文アイテムの関連
   - ユーザーと住所の関連

2. **トランザクション処理が必要なデータ**
   - 注文処理
   - 在庫更新
   - 支払い処理

3. **複雑な結合や集計が必要なデータ**
   - 注文レポート
   - 商品カタログ
   - 管理ダッシュボード向けデータ

4. **スキーマが明確に定義されたデータ**
   - 商品仕様
   - ユーザープロファイル
   - 注文情報

### 1.5.2. DynamoDBの選択理由（NoSQLデータベース）

以下のデータには、DynamoDB を利用することが適しています：

1. **高いスケーラビリティが必要なデータ**
   - ユーザーセッション
   - ショッピングカート
   - 商品閲覧履歴
   - ユーザー活動ログ

2. **読み書きが頻繁で低レイテンシが求められるデータ**
   - リアルタイムの在庫状態
   - ショッピングカート状態
   - 注文ステータス

3. **柔軟なスキーマが適しているデータ**
   - ユーザー活動（様々な種類のイベント）
   - 商品閲覧履歴
   - 注文ステータス履歴

4. **時系列データやイベントログ**
   - ユーザー活動ログ
   - 商品閲覧履歴
   - ステータス変更履歴

## 1.6. MVPと完成版のデータモデル比較

### 1.6.1. MVPのデータモデル特性

1. **シンプルさと基本機能重視**
   - 最小限のエンティティで基本機能をカバー（Product, Category, Inventory, Order, OrderItem, AdminUser）
   - シンプルな属性セットで基本情報のみ格納
   - 基本的な関連性のみ実装（1対多の基本リレーション）
   - MySQL (RDS) のみの単一データストア

2. **限定的なデータアクセスパターン**
   - 基本的なCRUD操作に最適化
   - シンプルなクエリとフィルタリング
   - 基本的なページネーション

3. **オブザーバビリティの基礎**
   - 基本的なログ属性
   - シンプルなメトリクス収集ポイント

### 1.6.2. 完成版のデータモデル特性

1. **包括的なeコマース機能サポート**
   - 多様なエンティティで完全なeコマース機能をカバー
   - 拡張された属性セットで詳細情報を格納
   - 複雑な関連性をサポート
   - 目的に応じたデータストアの使い分け（RDSとDynamoDB）

2. **高度なデータアクセスパターン**
   - 多様なクエリパターンに対応
   - 高度なフィルタリングと集計
   - リアルタイムデータアクセス
   - 履歴追跡と時系列分析

3. **包括的なオブザーバビリティサポート**
   - オブザーバビリティに最適化された属性設計
   - 多様なメトリクス収集ポイント
   - 監査と追跡のための設計
   - 行動分析と予測のためのデータ構造

### 1.6.3. データモデル進化のポイント

1. **ユーザー中心のデータモデルへの変化**
   - MVPでは主に商品と注文に焦点
   - 完成版ではユーザープロフィール、行動履歴、設定など、ユーザー中心のデータ構造を追加

2. **リアルタイム処理への対応**
   - MVPではバッチ処理や同期操作を中心としたデータモデル
   - 完成版ではリアルタイム処理に最適化されたDynamoDBエンティティの追加

3. **オブザーバビリティの深化**
   - MVPでは基本的なログとメトリクス
   - 完成版では詳細な監査、行動追跡、パフォーマンス分析、セキュリティモニタリングをサポートするデータ構造

4. **柔軟性と拡張性の向上**
   - MVPでは固定スキーマのデータモデル
   - 完成版ではJSON属性の活用やNoSQLエンティティの追加による柔軟性の確保

## 1.7. データアクセスパターンと最適化

### 1.7.1. MVPのデータアクセスパターン

1. **読み取りパターン**
   - 商品リスト表示（カテゴリフィルタ、ページネーション）
   - 商品詳細表示
   - 注文履歴表示

2. **書き込みパターン**
   - 注文作成
   - 在庫更新
   - 商品情報更新

3. **最適化手法**
   - 基本的なインデックス（外部キー、検索キー）
   - シンプルな結合クエリ

### 1.7.2. 完成版のデータアクセスパターン

1. **読み取りパターン**
   - 商品検索（複雑なフィルタ、ファセット検索）
   - パーソナライズされた商品推奨
   - リアルタイム在庫確認
   - 詳細な注文履歴と追跡
   - ユーザー行動分析

2. **書き込みパターン**
   - イベント駆動型の在庫更新
   - マルチステップの注文処理
   - ユーザー活動の継続的記録
   - 監査ログの自動生成

3. **最適化手法**
   - 目的別データストアの選択
   - 複雑なインデックス戦略
   - データ非正規化による読み取り最適化
   - キャッシュ層の導入
   - リードレプリカの活用
