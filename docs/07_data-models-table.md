# AWSオブザーバビリティ学習用eコマースアプリ - データモデル一覧

## RDSで実装するエンティティ

| エンティティ名 | 日本語名 | 主要属性 | 関連エンティティ | 説明 |
|--------------|---------|---------|----------------|------|
| Users | ユーザー | user_id (PK)<br>email<br>password_hash<br>name<br>created_at | Addresses, Orders, Reviews, Cart | システムのユーザー情報。認証情報や基本プロフィールを管理。 |
| Addresses | 住所 | address_id (PK)<br>user_id (FK)<br>address_line1<br>address_line2<br>city, state, zip, country | Users, Orders | ユーザーの配送先住所情報。複数登録可能。 |
| Products | 商品 | product_id (PK)<br>category_id (FK)<br>name<br>description<br>price<br>image_url | Categories, OrderItems, Inventory, Reviews | 販売商品の基本情報。 |
| Categories | カテゴリー | category_id (PK)<br>name<br>description | Products | 商品分類のためのカテゴリー情報。 |
| Orders | 注文 | order_id (PK)<br>user_id (FK)<br>address_id (FK)<br>status<br>total_amount<br>created_at | Users, Addresses, OrderItems | 注文の基本情報。ステータス管理を含む。 |
| OrderItems | 注文明細 | order_item_id (PK)<br>order_id (FK)<br>product_id (FK)<br>quantity<br>price | Orders, Products | 注文に含まれる個々の商品と数量、購入時価格。 |
| Reviews | レビュー | review_id (PK)<br>product_id (FK)<br>user_id (FK)<br>rating<br>comment | Products, Users | 商品に対するユーザーレビュー。 |
| Inventory (マスター) | 在庫マスター | inventory_id (PK)<br>product_id (FK)<br>SKU<br>reorder_point | Products | 在庫の基本情報。発注点など静的な在庫管理データ。 |

## DynamoDBで実装するエンティティ

| エンティティ名 | 日本語名 | キー構造 | 主要属性 | 説明 |
|--------------|---------|---------|---------|------|
| Cart | カート | PK: user_id | product_id<br>quantity<br>added_at | ユーザーのショッピングカート情報。リアルタイムの更新が頻繁。 |
| UserSessions | ユーザーセッション | PK: session_id | user_id<br>expires_at<br>attributes | ユーザーのセッション情報。有効期限付きの一時データ。 |
| Inventory (動的) | 在庫状態 | PK: product_id | quantity<br>reserved<br>last_updated | 商品の現在の在庫状況。頻繁に更新される動的データ。 |
| StockAlerts | 在庫通知 | PK: product_id | users_list<br>notification_type | 在庫切れ商品の入荷通知登録者リスト。 |
| ProductViews | 商品閲覧履歴 | PK: user_id<br>SK: timestamp | product_id<br>session_id | ユーザーの商品閲覧履歴。時系列データ。 |
| OrderStatus | 注文状態 | PK: order_id | current_status<br>status_history<br>updated_at | 注文のステータス履歴。頻繁に変更される可能性がある。 |
| SearchIndex | 検索インデックス | PK: keyword | product_ids<br>search_count | 検索キーワードと関連商品のマッピング。検索効率化用。 |
| UserActivity | ユーザー活動 | PK: user_id<br>SK: activity_type#timestamp | activity_details<br>context | ユーザーのサイト内活動記録。分析用途。 |

## 主要なリレーションシップ

| 主エンティティ | 関係 | 従エンティティ | 説明 |
|--------------|------|--------------|------|
| Users | 1:N | Addresses | 1人のユーザーが複数の住所を持つ |
| Users | 1:N | Orders | 1人のユーザーが複数の注文を行う |
| Orders | 1:N | OrderItems | 1つの注文に複数の商品項目が含まれる |
| Categories | 1:N | Products | 1つのカテゴリーに複数の商品が属する |
| Products | 1:1 | Inventory | 1つの商品に対して1つの在庫情報 |
| Products | 1:N | Reviews | 1つの商品に複数のレビューが付く |
| Users | 1:N | Reviews | 1人のユーザーが複数のレビューを投稿 |
| Users | 1:1 | Cart | 1人のユーザーに1つのカート（カート内に複数商品） |

## データストア選択の根拠

### RDSの選択理由

- 強い一貫性と参照整合性が必要なマスターデータ
- トランザクション処理が必要なデータ（注文処理など）
- 複雑な結合や集計が必要なデータ
- 定義されたスキーマが必要なデータ

### DynamoDBの選択理由

- 高いスケーラビリティが必要なデータ
- 読み書きが頻繁で低レイテンシが求められるデータ
- 柔軟なスキーマが適しているデータ
- 時系列データやイベントログのような追加型データ
- セッションデータなど短期間の一時データ
