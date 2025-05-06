# 1. データモデル設計

## 1.1. 目次

- [1. データモデル設計](#1-データモデル設計)
  - [1.1. 目次](#11-目次)
  - [1.2. はじめに](#12-はじめに)
    - [1.2.1. 目的](#121-目的)
    - [1.2.2. 凡例](#122-凡例)
  - [1.3. エンティティ定義 (フルスコープ考慮)](#13-エンティティ定義-フルスコープ考慮)
  - [1.4. ER図 (Mermaid - フルスコープ考慮・MVP明示)](#14-er図-mermaid---フルスコープ考慮mvp明示)

## 1.2. はじめに

### 1.2.1. 目的

この文書は、「AWSオブザーバビリティ学習用eコマースアプリ」で想定されるデータエンティティ、その属性、およびエンティティ間の関連性を定義します。MVPスコープと将来的な拡張可能性の両方を示します。

### 1.2.2. 凡例

- 実線: MVPで実装するエンティティとリレーション
- 点線: MVPスコープ外だが将来的に追加される可能性のあるエンティティやリレーション
- 属性名の横に `(MVP ✅)`: MVPで実装する属性
- 属性名の横に `(Future ⚪️)`: MVP以降に追加・変更される可能性のある属性

## 1.3. エンティティ定義 (フルスコープ考慮)

- **users (管理者):** システム管理者。MVPで必要。
- **customers (顧客):** ログインする顧客情報。MVPスコープ外。
- **categories:** 商品カテゴリ。MVPで必要。
- **products:** 商品情報。MVPで必要。
- **product_reviews:** 商品レビュー。MVPスコープ外。
- **carts:** 顧客のショッピングカート情報（ヘッダー）。MVPスコープ外（MVPでは非永続化）。
- **cart_items:** カート内の商品明細。MVPスコープ外。
- **orders:** 注文ヘッダー。MVPで必要（ゲスト注文中心）。
- **order_items:** 注文明細。MVPで必要。
- **promotions:** プロモーション情報。MVPスコープ外。
- **applied_promotions:** 注文に適用されたプロモーション。MVPスコープ外。

## 1.4. ER図 (Mermaid - フルスコープ考慮・MVP明示)

```mermaid
erDiagram
    USERS ||--o{ ORDERS : "places (future)"
    CUSTOMERS ||--o{ ORDERS : "places"
    CUSTOMERS ||--o{ CARTS : "has"
    CUSTOMERS ||--o{ PRODUCT_REVIEWS : "writes"
    CATEGORIES ||--|{ PRODUCTS : "contains"
    PRODUCTS ||--|{ ORDER_ITEMS : "included in"
    PRODUCTS ||--|{ CART_ITEMS : "added to"
    PRODUCTS ||--o{ PRODUCT_REVIEWS : "has"
    ORDERS ||--|{ ORDER_ITEMS : "consists of"
    ORDERS ||--o{ APPLIED_PROMOTIONS : "has"
    CARTS ||--|{ CART_ITEMS : "contains"
    PROMOTIONS ||--o{ APPLIED_PROMOTIONS : "applied as"

    USERS {
        VARCHAR(255) id PK "(MVP ✅) UUID or Cognito Sub"
        VARCHAR(255) email UK "(MVP ✅)"
        VARCHAR(255) password_hash "(MVP ✅) Hashed (if not Cognito)"
        VARCHAR(50) role "(MVP ✅) e.g., admin, manager"
        DATETIME created_at "(MVP ✅)"
        DATETIME updated_at "(MVP ✅)"
    }

    CUSTOMERS {
        VARCHAR(255) id PK "(Future ⚪️) UUID or Cognito Sub"
        VARCHAR(255) email UK "(Future ⚪️)"
        VARCHAR(255) name "(Future ⚪️)"
        TEXT default_shipping_address "(Future ⚪️)"
        DATETIME created_at "(Future ⚪️)"
        DATETIME updated_at "(Future ⚪️)"
    }

    CATEGORIES {
        INT id PK "(MVP ✅)"
        VARCHAR(255) name UK "(MVP ✅)"
        DATETIME created_at "(MVP ✅)"
        DATETIME updated_at "(MVP ✅)"
    }

    PRODUCTS {
        INT id PK "(MVP ✅)"
        VARCHAR(255) name "(MVP ✅)"
        TEXT description "(MVP ✅)"
        DECIMAL(10,2) price "(MVP ✅)"
        INT stock_quantity "(MVP ✅)"
        INT category_id FK "(MVP ✅)"
        VARCHAR(255) image_url "(MVP ✅)"
        BOOLEAN is_active "(MVP ✅)"
        FLOAT average_rating "(Future ⚪️)"
        DATETIME created_at "(MVP ✅)"
        DATETIME updated_at "(MVP ✅)"
    }

    PRODUCT_REVIEWS {
        INT id PK "(Future ⚪️)"
        INT product_id FK "(Future ⚪️)"
        VARCHAR(255) customer_id FK "(Future ⚪️)"
        INT rating "(Future ⚪️) 1-5"
        TEXT comment "(Future ⚪️)"
        DATETIME created_at "(Future ⚪️)"
    }

    CARTS {
        VARCHAR(255) id PK "(Future ⚪️) Session ID or Customer ID"
        VARCHAR(255) customer_id FK "(Future ⚪️) Nullable for guest"
        DATETIME created_at "(Future ⚪️)"
        DATETIME updated_at "(Future ⚪️)"
    }

    CART_ITEMS {
        INT id PK "(Future ⚪️)"
        VARCHAR(255) cart_id FK "(Future ⚪️)"
        INT product_id FK "(Future ⚪️)"
        INT quantity "(Future ⚪️)"
        DATETIME created_at "(Future ⚪️)"
    }

    ORDERS {
        INT id PK "(MVP ✅)"
        VARCHAR(255) customer_id FK "(Future ⚪️) Nullable for guest MVP"
        VARCHAR(255) guest_email "(MVP ✅) Required if customer_id is null"
        TEXT shipping_address "(MVP ✅)"
        DECIMAL(10,2) subtotal_amount "(Future ⚪️)"
        DECIMAL(10,2) discount_amount "(Future ⚪️)"
        DECIMAL(10,2) total_amount "(MVP ✅)"
        VARCHAR(50) status "(MVP ✅) e.g., pending, processing, shipped"
        VARCHAR(255) tracking_number "(Future ⚪️)"
        DATETIME created_at "(MVP ✅)"
        DATETIME updated_at "(MVP ✅)"
        VARCHAR(255) user_id FK "(Future ⚪️) For future admin orders"
    }

    ORDER_ITEMS {
        INT id PK "(MVP ✅)"
        INT order_id FK "(MVP ✅)"
        INT product_id FK "(MVP ✅)"
        INT quantity "(MVP ✅)"
        DECIMAL(10,2) price_at_purchase "(MVP ✅)"
        DATETIME created_at "(MVP ✅)"
        DATETIME updated_at "(MVP ✅)"
    }

    PROMOTIONS {
        INT id PK "(Future ⚪️)"
        VARCHAR(255) code UK "(Future ⚪️) e.g., SUMMER20"
        TEXT description "(Future ⚪️)"
        VARCHAR(50) type "(Future ⚪️) e.g., PERCENTAGE, FIXED_AMOUNT"
        DECIMAL(10,2) value "(Future ⚪️)"
        DATETIME start_date "(Future ⚪️)"
        DATETIME end_date "(Future ⚪️)"
        BOOLEAN is_active "(Future ⚪️)"
    }

    APPLIED_PROMOTIONS {
        INT id PK "(Future ⚪️)"
        INT order_id FK "(Future ⚪️)"
        INT promotion_id FK "(Future ⚪️)"
        DECIMAL(10,2) discount_applied "(Future ⚪️)"
    }

```

**データモデル注記 (MVP):**

- MVPでは `users`, `categories`, `products`, `orders`, `order_items` テーブルを実装します。
- `orders.customer_id` は将来の顧客ログイン機能のために用意しますが、MVPでは NULL 許容とし、代わりに `guest_email`, `shipping_address` にゲスト情報を格納します。
- カート関連テーブル (`carts`, `cart_items`) はMVPでは作成せず、カート情報はクライアントサイドまたは一時的なサーバーサイドストレージで管理します。
- レビュー (`product_reviews`)、プロモーション (`promotions`, `applied_promotions`) 関連テーブルはMVPスコープ外です。
