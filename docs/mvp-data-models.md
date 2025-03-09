# 1. MVPデータモデル一覧

## 1.1. リレーショナルデータベース(MySQL)スキーマ

| テーブル名  | 説明                                 | 主なカラム                                                                      |
| ----------- | ------------------------------------ | ------------------------------------------------------------------------------- |
| products    | 商品情報を格納                       | id, name, description, price, image_url, category_id, created_at, updated_at    |
| categories  | 商品カテゴリー情報を格納             | id, name, description, parent_id, created_at, updated_at                        |
| inventory   | 商品の在庫情報を格納                 | id, product_id, quantity, last_updated                                          |
| orders      | 注文情報を格納                       | id, customer_name, email, address, total_amount, status, created_at, updated_at |
| order_items | 注文に含まれる商品アイテム情報を格納 | id, order_id, product_id, quantity, price, subtotal                             |
| admin_users | 管理者ユーザー情報を格納             | id, username, password_hash, email, last_login                                  |

## 1.2. 詳細データモデル

### 1.2.1. Product (商品)

| フィールド  | タイプ        | 説明          | 制約                                |
| ----------- | ------------- | ------------- | ----------------------------------- |
| id          | int           | 商品ID        | PK, AUTO_INCREMENT                  |
| name        | varchar(255)  | 商品名        | NOT NULL                            |
| description | text          | 商品説明      | -                                   |
| price       | decimal(10,2) | 価格          | NOT NULL                            |
| image_url   | varchar(255)  | 商品画像のURL | -                                   |
| category_id | int           | カテゴリーID  | FK(categories.id)                   |
| created_at  | timestamp     | 作成日時      | DEFAULT CURRENT_TIMESTAMP           |
| updated_at  | timestamp     | 更新日時      | DEFAULT CURRENT_TIMESTAMP ON UPDATE |

### 1.2.2. Category (カテゴリー)

| フィールド  | タイプ       | 説明           | 制約                                |
| ----------- | ------------ | -------------- | ----------------------------------- |
| id          | int          | カテゴリーID   | PK, AUTO_INCREMENT                  |
| name        | varchar(100) | カテゴリー名   | NOT NULL, UNIQUE                    |
| description | varchar(255) | カテゴリー説明 | -                                   |
| parent_id   | int          | 親カテゴリーID | FK(categories.id), NULL許容         |
| created_at  | timestamp    | 作成日時       | DEFAULT CURRENT_TIMESTAMP           |
| updated_at  | timestamp    | 更新日時       | DEFAULT CURRENT_TIMESTAMP ON UPDATE |

### 1.2.3. Inventory (在庫)

| フィールド   | タイプ    | 説明         | 制約                                |
| ------------ | --------- | ------------ | ----------------------------------- |
| id           | int       | 在庫ID       | PK, AUTO_INCREMENT                  |
| product_id   | int       | 商品ID       | FK(products.id), NOT NULL           |
| quantity     | int       | 在庫数量     | NOT NULL, DEFAULT 0                 |
| last_updated | timestamp | 最終更新日時 | DEFAULT CURRENT_TIMESTAMP ON UPDATE |

### 1.2.4. Order (注文)

| フィールド    | タイプ        | 説明                                  | 制約                                |
| ------------- | ------------- | ------------------------------------- | ----------------------------------- |
| id            | int           | 注文ID                                | PK, AUTO_INCREMENT                  |
| customer_name | varchar(255)  | 顧客名                                | NOT NULL                            |
| email         | varchar(255)  | 顧客メールアドレス                    | NOT NULL                            |
| address       | text          | 配送先住所                            | NOT NULL                            |
| total_amount  | decimal(10,2) | 合計金額                              | NOT NULL                            |
| status        | enum          | 注文状態(新規,処理中,完了,キャンセル) | NOT NULL, DEFAULT '新規'            |
| created_at    | timestamp     | 作成日時                              | DEFAULT CURRENT_TIMESTAMP           |
| updated_at    | timestamp     | 更新日時                              | DEFAULT CURRENT_TIMESTAMP ON UPDATE |

### 1.2.5. OrderItem (注文アイテム)

| フィールド | タイプ        | 説明           | 制約                      |
| ---------- | ------------- | -------------- | ------------------------- |
| id         | int           | 注文アイテムID | PK, AUTO_INCREMENT        |
| order_id   | int           | 注文ID         | FK(orders.id), NOT NULL   |
| product_id | int           | 商品ID         | FK(products.id), NOT NULL |
| quantity   | int           | 数量           | NOT NULL                  |
| price      | decimal(10,2) | 購入時の価格   | NOT NULL                  |
| subtotal   | decimal(10,2) | 小計           | NOT NULL                  |

### 1.2.6. AdminUser (管理者ユーザー)

| フィールド    | タイプ       | 説明               | 制約                      |
| ------------- | ------------ | ------------------ | ------------------------- |
| id            | int          | 管理者ID           | PK, AUTO_INCREMENT        |
| username      | varchar(50)  | ユーザー名         | NOT NULL, UNIQUE          |
| password_hash | varchar(255) | パスワードハッシュ | NOT NULL                  |
| email         | varchar(255) | メールアドレス     | NOT NULL, UNIQUE          |
| last_login    | timestamp    | 最終ログイン日時   | NULL許容                  |
| created_at    | timestamp    | 作成日時           | DEFAULT CURRENT_TIMESTAMP |

## 1.3. ERD (Entity Relationship Diagram) の関連性

- Product (1) -- (0..*) Inventory
- Category (1) -- (0..*) Product
- Category (0..1) -- (0..*) Category (親子関係)
- Order (1) -- (1..*) OrderItem
- Product (1) -- (0..*) OrderItem
