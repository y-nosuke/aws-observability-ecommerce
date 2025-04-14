# 1. 統一ID体系の提案

ドキュメント全体（ユースケース、機能、画面、サービス、API、データモデル）で利用する、明確で一貫性があり、相互参照可能な命名規則を提案します。

**基本フォーマット:** `[種別]-[コンテキスト]-[コンポーネント/ドメイン]-[識別子]`

## 1.1. **種別プレフィックス:**

- `UC`: ユースケース (Use Case)
- `FEAT`: 機能 (Feature)
- `SCR`: 画面 (Screen)
- `SVC`: バックエンドサービス (Backend Service)
- `API`: APIエンドポイント (API Endpoint)
- `DM`: データモデル / エンティティ (Data Model)

## 1.2. **コンテキストプレフィックス:**

- `CUST`: 顧客向け (Customer-facing)
- `ADMIN`: 管理者向け (Administrator-facing)
- `OBS`: オブザーバビリティプロセス / システム監視 (Observability Process / System-level Monitoring)
- `BI`: ビジネスインテリジェンス (Business Intelligence)
- `SYS`: システムレベル（特定のアクターに限定されない共通コンポーネント、ヘルスチェックなど）(System-level)
- `CORE`: コアビジネスドメイン（基本的な概念を扱うサービス/APIなど）(Core business domain)

## 1.3. **コンポーネント/ドメインプレフィックス:**

- これは種別とコンテキストによって異なります。例：
  - *顧客向け(CUST):* `BROWSE` (閲覧), `SEARCH` (検索), `DETAIL` (詳細), `CART` (カート), `CHECKOUT` (注文手続き), `ORDER` (注文), `AUTH` (認証), `PROFILE` (プロフィール), `HISTORY` (履歴), `NOTIF` (通知)
  - *管理者向け(ADMIN):* `DASH` (ダッシュボード), `PROD` (商品), `INV` (在庫), `ORDER` (注文), `SHIP` (出荷), `RETURN` (返品), `REPORT` (レポート), `BATCH` (バッチ), `AUTH` (認証), `ROLE` (ロール), `USERMGMT` (ユーザー管理), `AUDIT` (監査)
  - *オブザーバビリティ(OBS):* `LOG` (ログ), `METRIC` (メトリクス), `TRACE` (トレース), `ALERT` (アラート), `DASH` (ダッシュボード), `RUM` (RUM), `SYNTH` (Synthetics), `SEC` (セキュリティ), `AUDIT` (監査), `PERF` (パフォーマンス), `CAPACITY` (キャパシティ), `FAULT` (障害注入), `COST` (コスト), `INSIGHT` (インサイト)
  - *ビジネスインテリジェンス(BI):* `SALES` (売上), `TRAFFIC` (トラフィック), `PRODPERF` (商品パフォーマンス), `USERBEHAV` (ユーザー行動), `CAMPAIGN` (キャンペーン)
  - *サービス/API/データモデル:* `PRODUCT` (商品), `CATEGORY` (カテゴリ), `INVENTORY` (在庫), `CART` (カート), `ORDER` (注文), `PAYMENT` (支払い), `USER` (ユーザー), `AUTH` (認証), `ADDRESS` (住所), `NOTIFICATION` (通知), `IMAGE` (画像), `BATCH` (バッチ), `HEALTH` (ヘルスチェック), `METRICS` (メトリクス), `SESSION` (セッション), `ROLE` (ロール), `AUDITLOG` (監査ログ)

## 1.4. **識別子:**

- 短く説明的な名前、または連番を使用します（名前が複雑になったり数が多くなったりする場合は連番を使用）。連番を使用する場合は、一貫したソートのためにゼロパディング（例: `01`, `02`）を行うことを推奨します。

**適用例:**

ドキュメントの例にこの体系を適用してみましょう。

### 1.4.1. **ユースケース (UC):** `UC-[コンテキスト]-[コンポーネント]-[識別子]`

- `UC-B01: カテゴリブラウジング` -> `UC-CUST-BROWSE-01`
- `UC-C01: カートへの商品追加` -> `UC-CUST-CART-01`
- `UC-U01: アカウント登録` -> `UC-CUST-AUTH-01` (UがUser/Authを指すと仮定)
- `UC-H01: 注文履歴確認` -> `UC-CUST-HISTORY-01`
- `UC-I01: 在庫モニタリング` -> `UC-ADMIN-INV-01`
- `UC-O01: 注文確認処理` -> `UC-ADMIN-ORDER-01`
- `UC-T01: バッチ処理管理` -> `UC-ADMIN-BATCH-01`
- `UC-A01: 管理者アカウント管理` -> `UC-ADMIN-AUTH-01`
- `UC-S01: セキュリティイベント監視` -> `UC-OBS-SEC-01`
- `UC-M01: パフォーマンス低下の検知と対応` -> `UC-OBS-PERF-01` (MはMonitoring/Maintenanceのようですが、PERFの方が明確)
- `UC-A01: コンプライアンス監査` (オブザーバビリティ内) -> `UC-OBS-AUDIT-01`
- `UC-P01: フロントエンドパフォーマンス監視` -> `UC-OBS-RUM-01` (PはPerformanceですが、RUMの方が具体的)
- `UC-BI01: 売上・トラフィック分析` -> `UC-BI-SALES-01`

### 1.4.2. **機能 (FEAT):** `FEAT-[コンテキスト]-[コンポーネント]-[識別子]`

- `C-BROWSE-01: 商品一覧表示` -> `FEAT-CUST-BROWSE-01`
- `C-SHOP-01: カート機能` -> `FEAT-CUST-CART-01` (SHOPは広範、CARTは具体的)
- `C-USER-01: 会員登録` -> `FEAT-CUST-AUTH-01`
- `C-NOTIF-01: 注文確認メール` -> `FEAT-CUST-NOTIF-01`
- `A-INV-01: 在庫レベル監視` -> `FEAT-ADMIN-INV-01`
- `A-PROD-01: 商品情報管理` -> `FEAT-ADMIN-PROD-01`
- `A-ORDER-01: 注文確認処理` -> `FEAT-ADMIN-ORDER-01`
- `A-BATCH-01: 在庫レポート生成` -> `FEAT-ADMIN-BATCH-01`
- `A-AUTH-01: 管理者認証` -> `FEAT-ADMIN-AUTH-01`
- `O-LOG-01: 構造化ログ設定` -> `FEAT-OBS-LOG-01`
- `O-METRIC-01: 基本メトリクス収集` -> `FEAT-OBS-METRIC-01`
- `O-TRACE-01: X-Ray基本統合` -> `FEAT-OBS-TRACE-01`
- `O-ALERT-01: ヘルスチェックエンドポイント` -> `FEAT-OBS-ALERT-01`
- `O-RUM-01: CloudWatch RUM設定` -> `FEAT-OBS-RUM-01`
- `O-FAULT-01: Fault Injection Service設定` -> `FEAT-OBS-FAULT-01`
- `O-SEC-01: 認証セキュリティモニタリング` -> `FEAT-OBS-SEC-01`
- `O-AUDIT-01: 監査ログシステム` -> `FEAT-OBS-AUDIT-01`
- `O-ADV-01: 異常検出実装` -> `FEAT-OBS-INSIGHT-01` (ADVはAdvancedの略ですが、INSIGHTやANOMALYの方が説明的)

### 1.4.3. **画面 (SCR):** `SCR-[コンテキスト]-[コンポーネント]-[識別子]`

- `C-SCR-01: トップページ` -> `SCR-CUST-BROWSE-01` または `SCR-CUST-HOME-01`
- `C-SCR-05: カート` -> `SCR-CUST-CART-01`
- `C-SCR-08: アカウント登録` -> `SCR-CUST-AUTH-01`
- `A-SCR-01: 管理ダッシュボード` -> `SCR-ADMIN-DASH-01`
- `O-SCR-01: システム概要ダッシュボード` -> `SCR-OBS-DASH-01`
- `B-SCR-01: 売上・トラフィック分析` -> `SCR-BI-SALES-01`

### 1.4.4. **バックエンドサービス (SVC):** `SVC-[ドメイン]-[識別子]`

- `SVC-01: 商品カタログサービス` -> `SVC-PRODUCT-01`
- `SVC-02: 在庫管理サービス` -> `SVC-INVENTORY-01`
- `SVC-03: カートサービス` -> `SVC-CART-01`
- `SVC-04: 注文処理サービス` -> `SVC-ORDER-01`
- `SVC-05: 認証サービス` -> `SVC-AUTH-01`
- `SVC-06: ヘルスチェックサービス` -> `SVC-SYS-HEALTH-01` (SYSコンテキストを使用)
- `SVC-07: メトリクスサービス` -> `SVC-OBS-METRICS-01` (オブザーバビリティ固有のためOBSコンテキストを使用)
- `SVC-10: 画像処理サービス` -> `SVC-IMAGE-01`
- `SVC-11: バッチ処理サービス` -> `SVC-BATCH-01`
- `SVC-12: イベント処理サービス` -> `SVC-SYS-EVENT-01` (または、より具体的なドメイン名)

### 1.4.5. **APIエンドポイント (API):** `API-[ドメイン]-[リソース]-[アクション/識別子]`

- `API-PRODUCT-01: GET /api/products` -> `API-PRODUCT-PRODUCTS-LIST`
- `API-PRODUCT-02: GET /api/products/{id}` -> `API-PRODUCT-PRODUCTS-GET`
- `API-PRODUCT-05: POST /api/products` -> `API-PRODUCT-PRODUCTS-CREATE`
- `API-INVENTORY-01: GET /api/inventory/{productId}` -> `API-INVENTORY-PRODUCT-GET`
- `API-CART-01: GET /api/carts/{userId}` -> `API-CART-USER-GET`
- `API-ORDER-01: POST /api/orders` -> `API-ORDER-ORDERS-CREATE`
- `API-AUTH-01: POST /api/auth/login` -> `API-AUTH-SESSION-CREATE` または `API-AUTH-LOGIN-01`
- `API-USER-01: GET /api/users/{id}` -> `API-USER-PROFILE-GET`
- `API-HEALTH-01: GET /api/health` -> `API-SYS-HEALTH-GET`
- `API-METRICS-01: GET /api/metrics` -> `API-OBS-METRICS-GET`
- `API-NOTIF-01: POST /api/notifications/send` -> `API-NOTIF-MESSAGES-CREATE`

### 1.4.6. **データモデル (DM):** `DM-[ストア種別]-[エンティティ名]`

- `Product (MySQL)` -> `DM-RDS-PRODUCT`
- `Category (MySQL)` -> `DM-RDS-CATEGORY`
- `Order (MySQL)` -> `DM-RDS-ORDER`
- `Cart (DynamoDB)` -> `DM-DDB-CART`
- `UserSessions (DynamoDB)` -> `DM-DDB-SESSION`
- `AuditLog (MySQL)` -> `DM-RDS-AUDITLOG`

### 1.4.7. **このID体系の利点:**

1. **明確性:** プレフィックスで成果物の種類がすぐにわかります。
2. **コンテキスト:** 2番目の部分でアクターやドメイン（顧客、管理者、オブザーバビリティ、システム、コア）が明確になります。
3. **グルーピング:** 3番目の部分で関連アイテム（例: `BROWSE`関連の全機能、`PRODUCT`関連の全API）がグループ化されます。
4. **一意性:** プロジェクト全体で一意なIDを保証します。
5. **相互参照:** アイテム間のリンク付けが容易になります。例えば、`UC-CUST-BROWSE-01`は`FEAT-CUST-BROWSE-01`と`FEAT-CUST-BROWSE-02`によって実装され、`SCR-CUST-BROWSE-01`に表示され、`SVC-PRODUCT-01`が`API-PRODUCT-PRODUCTS-LIST`経由でサポートし、`DM-RDS-PRODUCT`のデータを使用する、といった関連付けがしやすくなります。
