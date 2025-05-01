# 1. Week 2 - Day 1: データベーススキーマの設計と実装

## 1.1. 目次

- [1. Week 2 - Day 1: データベーススキーマの設計と実装](#1-week-2---day-1-データベーススキーマの設計と実装)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. eコマースデータベーススキーマの設計](#141-eコマースデータベーススキーマの設計)
    - [1.4.2. golang-migrateの設定と初期化](#142-golang-migrateの設定と初期化)
    - [1.4.3. マイグレーションファイルの作成（テーブル作成）](#143-マイグレーションファイルの作成テーブル作成)
    - [1.4.4. マイグレーションファイルの作成（ロールバック）](#144-マイグレーションファイルの作成ロールバック)
    - [1.4.5. テストデータ投入用マイグレーションファイルの作成](#145-テストデータ投入用マイグレーションファイルの作成)
    - [1.4.6. マイグレーションの実行](#146-マイグレーションの実行)
    - [1.4.7. 確認ポイント](#147-確認ポイント)
  - [1.5. 【詳細解説】](#15-詳細解説)
    - [1.5.1. リレーショナルデータベース設計の基本原則](#151-リレーショナルデータベース設計の基本原則)
    - [1.5.2. eコマースアプリケーションのデータモデル](#152-eコマースアプリケーションのデータモデル)
    - [1.5.3. マイグレーションベースの開発手法](#153-マイグレーションベースの開発手法)
  - [1.6. 【補足情報】](#16-補足情報)
    - [1.6.1. データベースの正規化](#161-データベースの正規化)
    - [1.6.2. インデックス設計の考え方](#162-インデックス設計の考え方)
  - [1.7. 【よくある問題と解決法】](#17-よくある問題と解決法)
    - [1.7.1. 問題1: マイグレーションバージョンの競合](#171-問題1-マイグレーションバージョンの競合)
    - [1.7.2. 問題2: 外部キー制約エラー](#172-問題2-外部キー制約エラー)
    - [1.7.3. 問題3: 文字セットと照合順序の問題](#173-問題3-文字セットと照合順序の問題)
  - [1.8. 【今日の重要なポイント】](#18-今日の重要なポイント)
  - [1.9. 【次回の準備】](#19-次回の準備)

## 1.2. 【要点】

- eコマースアプリケーションに必要なデータベーススキーマを設計し実装する
- リレーショナルデータベースの基本原則と正規化について理解する
- golang-migrateを使用したマイグレーションベースのデータベース管理手法を習得する
- テストデータの作成と投入方法を実践する
- マイグレーションの適用とロールバックの仕組みを理解する

## 1.3. 【準備】

データベーススキーマの設計と実装を行うために、以下の環境とツールを準備します。

### 1.3.1. チェックリスト

- [ ] Week 1で構築したDocker Compose環境が正常に起動していることを確認する

  ```bash
  docker compose ps
  ```

- [ ] MySQL CLIが利用可能であることを確認する

  ```bash
  docker compose exec mysql mysql --version
  ```

- [ ] golang-migrate CLIがインストールされていることを確認する

  ```bash
  migrate --version
  ```

  インストールされていない場合は以下を実行：

  ```bash
  go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```

- [ ] プロジェクト構造のmigrationディレクトリが存在することを確認、なければ作成する

  ```bash
  mkdir -p backend/internal/db/migrations
  ```

- [ ] MySQL接続情報が.envrcまたは環境変数で設定されていることを確認する

  ```bash
  # .envrc の例
  export MYSQL_HOST=localhost
  export MYSQL_PORT=3306
  export MYSQL_DATABASE=ecommerce
  export MYSQL_USER=ecommerce_user
  export MYSQL_PASSWORD=ecommerce_password
  ```

## 1.4. 【手順】

### 1.4.1. eコマースデータベーススキーマの設計

eコマースアプリケーションに必要な主要テーブルを設計します。以下の3つの基本テーブルを作成します：

1. **categories**: 商品カテゴリーを管理するテーブル
2. **products**: 商品情報を管理するテーブル
3. **inventory**: 商品の在庫情報を管理するテーブル

各テーブルの構造とリレーションシップを図に表すと以下のようになります:

```text
categories
+-----------+--------------+------+-----+-------------------+----------------+
| Field     | Type         | Null | Key | Default           | Extra          |
+-----------+--------------+------+-----+-------------------+----------------+
| id        | int          | NO   | PRI | NULL              | auto_increment |
| name      | varchar(255) | NO   |     | NULL              |                |
| slug      | varchar(255) | NO   | UNI | NULL              |                |
| parent_id | int          | YES  | MUL | NULL              |                |
| created_at| timestamp    | NO   |     | CURRENT_TIMESTAMP |                |
| updated_at| timestamp    | NO   |     | CURRENT_TIMESTAMP | on update...   |
+-----------+--------------+------+-----+-------------------+----------------+

products
+------------+--------------+------+-----+-------------------+----------------+
| Field      | Type         | Null | Key | Default           | Extra          |
+------------+--------------+------+-----+-------------------+----------------+
| id         | int          | NO   | PRI | NULL              | auto_increment |
| name       | varchar(255) | NO   |     | NULL              |                |
| description| text         | YES  |     | NULL              |                |
| price      | decimal(10,2)| NO   |     | 0.00              |                |
| category_id| int          | NO   | MUL | NULL              |                |
| sku        | varchar(100) | NO   | UNI | NULL              |                |
| image_url  | varchar(2048)| YES  |     | NULL              |                |
| created_at | timestamp    | NO   |     | CURRENT_TIMESTAMP |                |
| updated_at | timestamp    | NO   |     | CURRENT_TIMESTAMP | on update...   |
+------------+--------------+------+-----+-------------------+----------------+

inventory
+------------+--------------+------+-----+-------------------+----------------+
| Field      | Type         | Null | Key | Default           | Extra          |
+------------+--------------+------+-----+-------------------+----------------+
| id         | int          | NO   | PRI | NULL              | auto_increment |
| product_id | int          | NO   | MUL | NULL              |                |
| quantity   | int          | NO   |     | 0                 |                |
| created_at | timestamp    | NO   |     | CURRENT_TIMESTAMP |                |
| updated_at | timestamp    | NO   |     | CURRENT_TIMESTAMP | on update...   |
+------------+--------------+------+-----+-------------------+----------------+
```

### 1.4.2. golang-migrateの設定と初期化

マイグレーションを管理するためのディレクトリ構造を作成します。

```bash
# マイグレーションファイル用のディレクトリを作成
mkdir -p backend/internal/db/migrations
cd backend/internal/db/migrations
```

### 1.4.3. マイグレーションファイルの作成（テーブル作成）

golang-migrateを使って、テーブルを作成するための最初のマイグレーションファイルを作成します。

```bash
# マイグレーションディレクトリに移動
cd backend/internal/db/migrations

# 初期スキーマ作成用のマイグレーションファイルを作成
migrate create -ext sql -dir backend/internal/db/migrations -seq create_initial_schema
```

上記のコマンドにより、`000001_create_initial_schema.up.sql`と`000001_create_initial_schema.down.sql`という2つのファイルが作成されます。

`000001_create_initial_schema.up.sql`を以下の内容で編集します：

```sql
-- 商品カテゴリーテーブルの作成
CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    parent_id INT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_categories_parent_id (parent_id),
    CONSTRAINT fk_categories_parent FOREIGN KEY (parent_id) REFERENCES categories(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 商品テーブルの作成
CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    category_id INT NOT NULL,
    sku VARCHAR(100) NOT NULL UNIQUE,
    image_url VARCHAR(2048),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_products_category (category_id),
    CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 在庫テーブルの作成
CREATE TABLE IF NOT EXISTS inventory (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_inventory_product (product_id),
    CONSTRAINT fk_inventory_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

### 1.4.4. マイグレーションファイルの作成（ロールバック）

次に、マイグレーションのロールバック用のSQLを`000001_create_initial_schema.down.sql`に記述します。

```sql
-- 在庫テーブルの削除
DROP TABLE IF EXISTS inventory;

-- 商品テーブルの削除
DROP TABLE IF EXISTS products;

-- カテゴリーテーブルの削除
DROP TABLE IF EXISTS categories;
```

### 1.4.5. テストデータ投入用マイグレーションファイルの作成

テストデータを投入するための新しいマイグレーションファイルを作成します。

```bash
# テストデータ投入用のマイグレーションファイルを作成
migrate create -ext sql -dir backend/internal/db/migrations -seq insert_test_data
```

`000002_insert_test_data.up.sql`を以下の内容で編集します：

```sql
-- カテゴリーのテストデータ
INSERT INTO categories (name, slug) VALUES
('電子機器', 'electronics'),
('洋服', 'clothing'),
('書籍', 'books'),
('ホーム＆キッチン', 'home-kitchen');

-- 電子機器のサブカテゴリー
INSERT INTO categories (name, slug, parent_id) VALUES
('スマートフォン', 'smartphones', 1),
('ノートパソコン', 'laptops', 1),
('タブレット', 'tablets', 1);

-- 洋服のサブカテゴリー
INSERT INTO categories (name, slug, parent_id) VALUES
('メンズ', 'mens-clothing', 2),
('レディース', 'womens-clothing', 2);

-- 商品のテストデータ
INSERT INTO products (name, description, price, category_id, sku, image_url) VALUES
('Acme スマートフォン', 'Acmeの最新スマートフォン', 89800, 5, 'ACME-SP-001', 'https://example.com/images/acme-smartphone.jpg'),
('Zenith ノートPC', '高性能ノートパソコン', 125000, 6, 'ZNT-NP-001', 'https://example.com/images/zenith-laptop.jpg'),
('Quantum タブレット', '10インチタブレット', 45000, 7, 'QNT-TB-001', 'https://example.com/images/quantum-tablet.jpg'),
('メンズカジュアルシャツ', '綿100%のカジュアルシャツ', 3800, 8, 'MCS-001', 'https://example.com/images/mens-casual-shirt.jpg'),
('レディースブラウス', 'エレガントなデザインのブラウス', 4200, 9, 'LDB-001', 'https://example.com/images/ladies-blouse.jpg'),
('プログラミング入門', 'プログラミングの基礎を学ぶ', 2800, 3, 'BK-PRG-001', 'https://example.com/images/programming-book.jpg'),
('キッチンミキサー', '多機能キッチンミキサー', 6500, 4, 'KM-001', 'https://example.com/images/kitchen-mixer.jpg');

-- 在庫データ
INSERT INTO inventory (product_id, quantity) VALUES
(1, 50),  -- Acme スマートフォン
(2, 25),  -- Zenith ノートPC
(3, 30),  -- Quantum タブレット
(4, 100), -- メンズカジュアルシャツ
(5, 80),  -- レディースブラウス
(6, 60),  -- プログラミング入門
(7, 15);  -- キッチンミキサー
```

`000002_insert_test_data.down.sql`を以下の内容で編集します：

```sql
-- 在庫データの削除
DELETE FROM inventory;

-- 商品データの削除
DELETE FROM products;

-- サブカテゴリーの削除
DELETE FROM categories WHERE parent_id IS NOT NULL;

-- 親カテゴリーの削除
DELETE FROM categories WHERE parent_id IS NULL;
```

### 1.4.6. マイグレーションの実行

作成したマイグレーションファイルを実行してデータベーススキーマを作成し、テストデータを投入します。

```bash
# プロジェクトのルートディレクトリに移動
cd ../../..

# マイグレーションを実行
migrate -path ${MIGRATIONS_PATH} -database "${MYSQL_DSN}" up
```

### 1.4.7. 確認ポイント

- [ ] マイグレーションが正常に実行され、テーブルが作成されていることを確認

  ```bash
  docker compose exec mysql mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} -e "SHOW TABLES FROM ${MYSQL_DATABASE}"
  ```

  期待される出力:

  ```text
  Tables_in_ecommerce
  categories
  inventory
  products
  schema_migrations
  ```

- [ ] カテゴリーデータが正しく投入されていることを確認

  ```bash
  docker compose exec mysql mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} -e "SELECT id, name, slug, parent_id FROM ${MYSQL_DATABASE}.categories"
  ```

- [ ] 商品データが正しく投入されていることを確認

  ```bash
  docker compose exec mysql mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} -e "SELECT id, name, price, category_id FROM ${MYSQL_DATABASE}.products"
  ```

- [ ] 在庫データが正しく投入されていることを確認

  ```bash
  docker compose exec mysql mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} -e "SELECT product_id, quantity FROM ${MYSQL_DATABASE}.inventory"
  ```

- [ ] 外部キー制約が正しく設定されていることを確認

  ```bash
  docker compose exec mysql mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} -e "SELECT TABLE_NAME, COLUMN_NAME, CONSTRAINT_NAME, REFERENCED_TABLE_NAME, REFERENCED_COLUMN_NAME FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE WHERE REFERENCED_TABLE_SCHEMA = '${MYSQL_DATABASE}'"
  ```

## 1.5. 【詳細解説】

### 1.5.1. リレーショナルデータベース設計の基本原則

リレーショナルデータベース設計では以下の原則が重要です：

1. **正規化（Normalization）**: データの重複を減らし、整合性を確保するための設計技法です。今回のスキーマでは、商品情報と在庫情報を別テーブルに分割することで第三正規形（3NF）に近い設計を実現しています。

2. **主キーと外部キー**: 各テーブルには一意の識別子（主キー）を設け、テーブル間の関連はこの主キーを参照する外部キーで表現します。例えば、`products`テーブルの`category_id`は`categories`テーブルの`id`を参照する外部キーです。

3. **参照整合性**: 外部キー制約を通じてデータの整合性を保ちます。例えば、商品が参照するカテゴリーが存在することを保証します。また、カスケード削除やNULL設定などのオプションで、親レコードが削除された場合の子レコードの扱いを定義しています。

4. **インデックス設計**: クエリのパフォーマンスを向上させるため、頻繁に検索や結合に使用されるカラムにはインデックスを設定します。今回は外部キーのカラムと、一意性が必要なカラム（slugやsku）にインデックスを設定しています。

### 1.5.2. eコマースアプリケーションのデータモデル

今回実装したeコマースアプリケーションの基本的なデータモデルには以下の特徴があります：

1. **カテゴリー階層構造**: `categories`テーブルの`parent_id`を使用して、自己参照する階層構造を表現しています。これにより、メインカテゴリーとサブカテゴリーを柔軟に管理できます。

2. **商品と在庫の分離**: 商品の基本情報（名前、説明、価格など）と在庫情報を別テーブルに分離しています。これにより、在庫の動きを追跡しやすくなり、拡張性も高まります。

3. **識別子の工夫**: 自動採番の`id`を使って内部識別するとともに、`sku`（商品管理番号）など業務的な識別子も設けています。また、カテゴリーには`slug`フィールドを設け、URLで使用可能な文字列識別子としています。

4. **監査証跡**: すべてのテーブルに`created_at`と`updated_at`のタイムスタンプを設け、レコードの作成・更新履歴を追跡できるようにしています。

### 1.5.3. マイグレーションベースの開発手法

golang-migrateを使用したマイグレーションベースの開発には以下のメリットがあります：

1. **バージョン管理**: 各マイグレーションにはバージョン番号（シーケンス番号）が付与され、どのマイグレーションが適用されたかを`schema_migrations`テーブルで管理します。これにより、チーム内での整合性を保ちやすくなります。

2. **冪等性（べきとうせい）**: 各マイグレーションは一度だけ適用されるため、同じ変更が繰り返し適用される心配がありません。

3. **ロールバック**: 各マイグレーションには「up」（適用）と「down」（ロールバック）の両方のSQLが定義され、必要に応じて変更を元に戻すことができます。

4. **環境間の一貫性**: 開発環境、テスト環境、本番環境など、異なる環境でも同じマイグレーションを適用することで、スキーマの一貫性を保つことができます。

## 1.6. 【補足情報】

### 1.6.1. データベースの正規化

データベースの正規化は、データの重複を減らし、データ整合性を高めるための設計手法です。代表的な正規化レベルは以下の通りです：

1. **第一正規形（1NF）**: 各カラムが原子的（分割できない）値を持ち、繰り返しグループがない状態。

2. **第二正規形（2NF）**: 1NFを満たし、かつ部分関数従属性がない状態。つまり、非キーカラムが主キー全体に依存している状態。

3. **第三正規形（3NF）**: 2NFを満たし、かつ推移的関数従属性がない状態。つまり、非キーカラムが他の非キーカラムに依存していない状態。

eコマースシステムではしばしば、パフォーマンスとのバランスを考慮して、特定の部分で非正規化を行うこともあります。例えば、商品の合計評価点や評価数など、頻繁に計算が必要な集計値を保存することがあります。

### 1.6.2. インデックス設計の考え方

効果的なインデックス設計は、データベースのパフォーマンスに大きく影響します。主な考慮点は以下の通りです：

1. **検索条件の分析**: クエリの`WHERE`句で頻繁に使用されるカラムにインデックスを設定します。

2. **結合条件の分析**: `JOIN`条件に使用されるカラム（外部キーなど）にインデックスを設定します。

3. **カーディナリティ**: 値の種類が多いカラム（一意性が高い）ほどインデックスの効果が高くなります。

4. **複合インデックス**: 複数のカラムで検索する場合、それらのカラムを組み合わせた複合インデックスを検討します。複合インデックスの順序は重要で、最も絞り込み効果が高いカラムを最初に配置します。

5. **更新頻度**: 更新が頻繁に行われるカラムにインデックスを設定すると、書き込みパフォーマンスが低下する可能性があります。

6. **インデックスのオーバーヘッド**: インデックスはディスク容量と維持コストを消費するため、必要以上に多くのインデックスを設定しないよう注意します。

今回の設計では、外部キーのカラムと一意制約のあるカラムにインデックスを設定しています。プロジェクトの進行に伴ってクエリパターンが明確になったら、必要に応じて追加のインデックスを検討するとよいでしょう。

## 1.7. 【よくある問題と解決法】

### 1.7.1. 問題1: マイグレーションバージョンの競合

**症状**: チームの複数のメンバーが同時にマイグレーションファイルを作成し、同じバージョン番号を持つマイグレーションが発生する。

**解決策**:

1. マイグレーションファイルの作成はチーム内で調整し、順序を揃える。
2. 競合が発生した場合は、以下の手順で解決する：
   - 競合したマイグレーションをマージして内容を統合する
   - マイグレーションの適用状態をリセットする（必要に応じてデータベースを再作成）
   - 統合したマイグレーションを適用する
3. マイグレーションファイル名にタイムスタンプを使用する方式に切り替える：

   ```bash
   migrate create -ext sql -dir . -format "20060102150405" create_new_table
   ```

### 1.7.2. 問題2: 外部キー制約エラー

**症状**: マイグレーションの実行時に「Cannot add or update a child row: a foreign key constraint fails」というエラーが発生する。

**解決策**:

1. マイグレーションの順序を確認し、参照先のテーブルが先に作成されるようにする。
2. テーブル作成とデータ投入を別のマイグレーションに分離する。
3. 一時的に外部キーチェックを無効にしてデータを投入する（本番環境では注意が必要）：

   ```sql
   SET FOREIGN_KEY_CHECKS = 0;
   -- データの投入
   SET FOREIGN_KEY_CHECKS = 1;
   ```

4. マイグレーションをロールバックして順序を修正し、再度適用する：

   ```bash
   migrate -path ${MIGRATIONS_PATH} -database "${MYSQL_DSN}" down
   migrate -path ${MIGRATIONS_PATH} -database "${MYSQL_DSN}" up
   ```

### 1.7.3. 問題3: 文字セットと照合順序の問題

**症状**: 日本語などの非ASCII文字を含むデータで文字化けや比較の問題が発生する。

**解決策**:

1. テーブル作成時に明示的に文字セットと照合順序を指定する：

   ```sql
   CREATE TABLE example (
       ...
   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
   ```

2. データベース接続文字列に文字セットパラメータを追加する：

   ```text
   mysql://user:password@tcp(host:port)/dbname?charset=utf8mb4&collation=utf8mb4_0900_ai_ci
   ```

3. アプリケーションのコードでもUTF-8エンコーディングを使用し、文字セットの一貫性を保つ。

4. 既存のテーブルの文字セットと照合順序を変更する（必要な場合）：

   ```sql
   ALTER TABLE example CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
   ```

## 1.8. 【今日の重要なポイント】

1. **データモデルは全てのスタート**：適切なデータモデルはアプリケーション全体の基盤となります。時間をかけて適切に設計することで、後の開発がスムーズになります。

2. **リレーションシップを意識した設計**：テーブル間の関係を外部キー制約として明示的にモデル化することで、データの整合性を保ちやすくなります。

3. **マイグレーションによるバージョン管理**：スキーマ変更をマイグレーションとして管理することで、変更履歴を追跡し、環境間の一貫性を保つことができます。

4. **ロールバック可能な設計**：マイグレーションの「down」部分を適切に実装することで、必要に応じて変更を安全に元に戻すことができます。

5. **テストデータの重要性**：実際のデータに近いテストデータを用意することで、開発中のユースケースを効果的に検証できます。

## 1.9. 【次回の準備】

次回（Day 2）では、sqlboilerを使用したORMの設定と基本的なデータアクセスコードの実装を行います。以下の点について事前に確認しておくと良いでしょう：

1. **sqlboilerのインストール**：

   ```bash
   go install github.com/volatiletech/sqlboiler/v4@latest
   go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
   ```

2. **Go言語の基本的な構文**：特に構造体、インターフェース、エラーハンドリングについて確認しておくとよいでしょう。

3. **GoのDBアクセスパターン**：`database/sql`パッケージの基本的な使い方について理解しておくと、sqlboilerの使い方を理解しやすくなります。

4. **実装したデータベーススキーマの復習**：テーブル構造と関連性を確認し、どのようなデータアクセスパターンが必要になるか考えておきましょう。

次回も引き続き、Dockerコンテナ内のMySQLデータベースを使用するため、Docker環境が正常に動作していることを確認しておいてください。
