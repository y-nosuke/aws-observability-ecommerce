# 1. 機能一覧

## 1.1. 目次

- [1. 機能一覧](#1-機能一覧)
  - [1.1. 目次](#11-目次)
  - [1.2. はじめに](#12-はじめに)
    - [1.2.1. 目的](#121-目的)
    - [1.2.2. 凡例](#122-凡例)
  - [1.3. 機能一覧](#13-機能一覧)

## 1.2. はじめに

### 1.2.1. 目的

この文書は、「AWSオブザーバビリティ学習用eコマースアプリ」で想定される全ての機能をリスト化し、その概要、関連コンポーネント、およびMVPでの実装ステータスを明確にすることを目的とします。

### 1.2.2. 凡例

- **ステータス:**
  - ✅: MVP (リリース1) で実装
  - ⚪️: MVP以降 (リリース2など) で実装検討
  - ❌: 将来の拡張候補 / MVPでは実装しない
- **FE:** フロントエンド, **BE:** バックエンド, **API:** APIエンドポイント, **DB:** データベース, **OBS:** オブザーバビリティ

## 1.3. 機能一覧

| 機能カテゴリ               | 機能ID                   | 機能名                      | 概要                                                        | 主要関連コンポーネント (例)                                                         | ステータス | 関連ユーザーストーリー (例)        |
| :------------------------- | :----------------------- | :-------------------------- | :---------------------------------------------------------- | :---------------------------------------------------------------------------------- | :--------: | :--------------------------------- |
| **商品閲覧 (顧客)**        | FN-CUST-PROD-01          | 商品一覧表示                | カテゴリやキーワードでフィルタリングされた商品リスト表示。  | FE: 商品一覧, API: `/products`, BE: Lambda_Product                                  |     ✅      | US-CUST-BROWSE-01, 02, 03          |
|                            | FN-CUST-PROD-02          | 商品詳細表示                | 商品詳細情報（説明、価格、在庫）表示。                      | FE: 商品詳細, API: `/products/{id}`, BE: Lambda_Product                             |     ✅      | US-CUST-BROWSE-04, 05              |
|                            | FN-CUST-PROD-03          | 関連商品表示                | 商品詳細ページで関連する商品を表示。                        | FE: 商品詳細, API: `/products/{id}/related`, BE: Lambda_Product                     |     ❌      | US-CUST-BROWSE-06                  |
|                            | FN-CUST-PROD-NOTIF-01    | 入荷通知登録                | 在庫切れ商品の入荷通知をメールアドレスで登録。              | FE: 商品詳細, API: `/products/{id}/notify`, BE: Lambda_Product                      |     ⚪️      | US-CUST-BROWSE-NOTIF-01            |
|                            | FN-CUST-PROD-NOTIF-02    | 入荷通知送信                | 商品入荷時に登録ユーザーへメール通知。                      | BE: Fargate_Stock -> SNS -> SES, DB                                                 |     ⚪️      | US-CUST-BROWSE-NOTIF-01            |
| **カート (顧客)**          | FN-CUST-CART-01          | カートへの商品追加          | 選択した商品をカートに追加。                                | FE: 商品詳細, API: `/cart/items` (POST), BE: (Lambda or Fargate)                    |     ✅      | US-CUST-CART-01                    |
|                            | FN-CUST-CART-02          | カート内容表示              | カート内商品リストと合計金額表示。                          | FE: カート, API: `/cart` (GET), BE: (Lambda or Fargate)                             |     ✅      | US-CUST-CART-04                    |
|                            | FN-CUST-CART-03          | カート内商品数量変更        | カート内の商品数量を変更。                                  | FE: カート, API: `/cart/items/{itemId}` (PUT), BE: (Lambda or Fargate)              |     ⚪️      | US-CUST-CART-02                    |
|                            | FN-CUST-CART-04          | カートからの商品削除        | カートから特定の商品を削除。                                | FE: カート, API: `/cart/items/{itemId}` (DELETE), BE: (Lambda or Fargate)           |     ⚪️      | US-CUST-CART-03                    |
| **注文 (顧客)**            | FN-CUST-ORDER-01         | 注文処理 (ゲスト・ログイン) | 配送先・支払い情報を受け取り、注文作成、在庫引き当て。      | FE: チェックアウト, API: `/orders` (POST), BE: Lambda_Order -> SQS -> Fargate_Stock |     ✅      | US-CUST-CHECKOUT-01, 02            |
|                            | FN-CUST-ORDER-02         | 注文完了表示                | 注文完了メッセージ、注文番号表示。                          | FE: 注文完了                                                                        |     ✅      | (Checkout Flow)                    |
|                            | FN-CUST-ORDER-COUPON-01  | クーポン適用                | チェックアウト時にクーポンコードを適用し割引。              | FE: チェックアウト, API: `/orders/apply-coupon` (POST), BE: Lambda_Order            |     ❌      | US-CUST-CHECKOUT-03                |
|                            | FN-CUST-ORDER-NOTIF-01   | 注文確認通知                | 注文完了時に顧客へ確認メール送信。                          | BE: Lambda_Order -> SNS -> SES                                                      |     ⚪️      | US-CUST-CHECKOUT-NOTIF-01          |
| **アカウント (顧客)**      | FN-CUST-AUTH-01          | 顧客アカウント登録          | Eメール、パスワード等で顧客アカウントを作成。               | FE: 登録, API: `/auth/register`, BE: Cognito                                        |     ⚪️      | US-CUST-AUTH-01                    |
|                            | FN-CUST-AUTH-02          | 顧客ログイン                | Eメール/パスワード or ソーシャルログインで認証。            | FE: ログイン, API: `/auth/login`, BE: Cognito                                       |     ⚪️      | US-CUST-AUTH-02                    |
|                            | FN-CUST-AUTH-03          | パスワードリセット          | パスワード忘れ時にリセットフローを提供。                    | FE: PWリセット, API: `/auth/forgot-password`, `/auth/reset-password`, BE: Cognito   |     ⚪️      | US-CUST-AUTH-03                    |
|                            | FN-CUST-PROFILE-01       | プロフィール表示・編集      | 登録された顧客情報（氏名、連絡先、住所等）の表示・編集。    | FE: マイページ/プロフ編集, API: `/account/profile`, BE: Lambda                      |     ❌      | US-CUST-PROFILE-01                 |
|                            | FN-CUST-HISTORY-01       | 注文履歴一覧表示            | ログイン顧客の過去注文リスト表示。                          | FE: マイページ/注文履歴, API: `/account/orders`, BE: Lambda_Order                   |     ❌      | US-CUST-HISTORY-01                 |
|                            | FN-CUST-HISTORY-02       | 注文履歴詳細表示            | 特定の過去注文の詳細表示。                                  | FE: 注文履歴詳細, API: `/account/orders/{orderId}`, BE: Lambda_Order                |     ❌      | US-CUST-HISTORY-01                 |
|                            | FN-CUST-HISTORY-TRACK-01 | 配送状況確認                | 注文の配送ステータス・追跡番号表示。                        | FE: 注文履歴詳細, API: `/account/orders/{orderId}/tracking`, BE: Lambda_Order       |     ❌      | US-CUST-HISTORY-02                 |
|                            | FN-CUST-HISTORY-NOTIF-01 | 注文ステータス更新通知      | 注文ステータス変更（発送等）時に顧客へ通知。                | BE: Fargate_Stock/Lambda_Order -> SNS -> SES                                        |     ❌      | US-CUST-HISTORY-NOTIF-01           |
| **商品管理 (管理者)**      | FN-MGR-PROD-01           | 商品登録                    | 新商品情報（名前、説明、価格、カテゴリ、在庫、画像）登録。  | FE: 商品登録/編集, API: `/admin/products` (POST), BE: Lambda_Product                |     ✅      | US-MGR-PROD-01, 04, INV-02         |
|                            | FN-MGR-PROD-02           | 商品更新                    | 既存の商品情報更新。                                        | FE: 商品登録/編集, API: `/admin/products/{id}` (PUT), BE: Lambda_Product            |     ✅      | US-MGR-PROD-02                     |
|                            | FN-MGR-PROD-03           | 商品一覧取得 (管理用)       | 管理画面用商品リスト取得。                                  | FE: 商品管理一覧, API: `/admin/products` (GET), BE: Lambda_Product                  |     ✅      | (Admin View)                       |
|                            | FN-MGR-PROD-04           | 商品画像アップロード処理    | アップロード画像をS3保存、(リサイズ等)。                    | BE: (Lambda_Product or dedicated Lambda), S3                                        |     ✅      | US-MGR-PROD-04                     |
|                            | FN-MGR-PROD-05           | 商品有効/無効化             | 商品の販売ステータス切替。                                  | FE: 商品管理一覧, API: `/admin/products/{id}/status` (PUT), BE: Lambda              |     ✅      | US-MGR-PROD-03                     |
| **カテゴリ管理(管理者)**   | FN-MGR-CAT-01            | カテゴリ登録                | 新商品カテゴリ登録。                                        | FE: カテゴリ管理, API: `/admin/categories` (POST), BE: Lambda_Product               |     ✅      | US-MGR-CAT-01                      |
|                            | FN-MGR-CAT-02            | カテゴリ一覧取得            | 登録カテゴリリスト取得。                                    | FE: カテゴリ管理, 商品登録/編集, API: `/admin/categories` (GET), BE: Lambda         |     ✅      | (Admin View)                       |
|                            | FN-MGR-CAT-03            | カテゴリ更新・削除          | 既存カテゴリの名称変更や削除。                              | FE: カテゴリ管理, API: `/admin/categories/{id}` (PUT/DELETE), BE: Lambda            |     ❌      | US-MGR-CAT-02                      |
| **在庫管理 (管理者)**      | FN-MGR-INV-01            | 在庫一覧取得 (管理用)       | 商品在庫数一覧表示。                                        | FE: 在庫管理一覧, API: `/admin/inventory` (GET), BE: Lambda/Fargate                 |     ✅      | US-MGR-INV-01                      |
|                            | FN-MGR-INV-02            | 在庫数手動更新              | 特定商品在庫数手動更新。                                    | FE: 商品登録/編集, API: `/admin/products/{id}` (PUT内)                              |     ✅      | US-MGR-INV-02                      |
|                            | FN-MGR-INV-ALERT-01      | 低在庫アラート設定          | 商品ごとに低在庫アラート閾値設定。                          | FE: 在庫アラート設定, API: `/admin/inventory/alerts` (POST/PUT)                     |     ⚪️      | US-MGR-INV-03                      |
|                            | FN-MGR-INV-ALERT-02      | 低在庫アラート通知          | 在庫が閾値を下回った際に管理者に通知。                      | BE: Fargate_Stock -> SNS                                                            |     ⚪️      | US-MGR-INV-03                      |
| **注文管理 (管理者)**      | FN-MGR-ORDER-01          | 注文一覧取得 (管理用)       | 受注注文リスト取得。                                        | FE: 注文管理一覧, API: `/admin/orders` (GET), BE: Lambda_Order                      |     ✅      | US-MGR-ORDER-01                    |
|                            | FN-MGR-ORDER-02          | 注文詳細取得 (管理用)       | 特定注文詳細情報取得。                                      | FE: 注文詳細, API: `/admin/orders/{id}` (GET), BE: Lambda_Order                     |     ✅      | US-MGR-ORDER-03                    |
|                            | FN-MGR-ORDER-03          | 注文ステータス更新          | 注文ステータス変更（処理中、出荷済み等）。                  | FE: 注文詳細, API: `/admin/orders/{id}/status` (PUT), BE: Lambda_Order              |     ⚪️      | US-MGR-ORDER-02                    |
|                            | FN-MGR-ORDER-04          | 出荷情報登録                | 注文に追跡番号などの出荷情報を登録。                        | FE: 注文詳細, API: `/admin/orders/{id}/shipment` (POST/PUT), BE: Lambda             |     ⚪️      | US-MGR-ORDER-04                    |
|                            | FN-MGR-ORDER-RETURN-01   | 返品・交換処理              | 顧客からの返品・交換リクエストを管理。                      | FE: 返品交換管理, API: `/admin/returns`, BE: Lambda_Order                           |     ❌      | US-MGR-ORDER-05                    |
| **レポート (管理者)**      | FN-MGR-REPORT-01         | 基本売上レポート表示        | 日次/月次売上などの基本レポート表示。                       | FE: 基本レポート, API: `/admin/reports/sales-summary`, BE: Lambda                   |     ⚪️      | US-MGR-REPORT-01                   |
| **マーケティング(管理者)** | FN-MGR-PROMO-01          | プロモーション設定・管理    | セールやクーポンなどのプロモーションを作成・管理。          | FE: プロモ管理, API: `/admin/promotions`, BE: Lambda                                |     ❌      | US-MGR-PROMO-01                    |
| **認証 (管理者)**          | FN-ADMIN-AUTH-01         | 管理者ログイン              | ID/パスワードで管理者を認証。                               | FE: 管理者ログイン, BE: Cognito                                                     |     ✅      | US-MGR-AUTH-LOGIN-01               |
| **アカウント管理(管理者)** | FN-SEC-USER-01           | 管理者アカウント作成        | 新規管理者アカウント作成。                                  | FE: アカウント管理, API: `/admin/users` (POST), BE: Cognito                         |     ✅      | US-SEC-AUTH-MANAGE-01              |
|                            | FN-SEC-USER-02           | 管理者アカウント一覧取得    | 登録済管理者リスト取得。                                    | FE: アカウント管理, API: `/admin/users` (GET), BE: Cognito                          |     ✅      | (Admin View)                       |
|                            | FN-SEC-USER-03           | 管理者アカウント更新        | アカウント有効/無効化、ロール割り当て等。                   | FE: アカウント管理, API: `/admin/users/{userId}` (PUT), BE: Cognito                 |     ✅      | US-SEC-AUTH-MANAGE-01              |
| **ロール管理 (管理者)**    | FN-SEC-ROLE-01           | ロール作成                  | 新規管理ロール作成。                                        | FE: ロール管理, API: `/admin/roles` (POST), BE: Cognito (Groups)                    |     ✅      | US-SEC-AUTH-RBAC-01                |
|                            | FN-SEC-ROLE-02           | ロール一覧取得              | 既存ロールリスト取得。                                      | FE: ロール管理, API: `/admin/roles` (GET), BE: Cognito (Groups)                     |     ✅      | (Admin View)                       |
|                            | FN-SEC-ROLE-03           | ロール権限設定              | ロールに紐づく権限設定。                                    | FE: ロール管理, BE: IAM Policy / API GW Authorizer?                                 |     ✅      | US-SEC-AUTH-RBAC-01                |
|                            | FN-SEC-POLICY-AUTH-01    | 認証ポリシー設定            | MFA必須化、パスワードポリシー、セッションポリシー設定。     | FE: 認証ポリシー設定, BE: Cognito                                                   |     ⚪️      | US-SEC-AUTH-MFA-01, POLICY-01      |
| **監査 (管理者)**          | FN-SEC-AUDIT-01          | 認証・認可監査ログ確認      | 管理者のログイン試行や権限変更などのログ確認。              | FE: 監査ログビューア, BE: CloudWatch Logs (from Cognito/IAM)                        |     ⚪️      | US-SEC-AUTH-AUDIT-01               |
|                            | FN-SEC-AUDIT-02          | 操作ログ監査                | 重要なデータ変更などの操作ログ確認。                        | FE: 監査ログビューア, BE: CloudWatch Logs (App Logs)                                |     ⚪️      | US-SEC-OBS-AUDIT-01                |
|                            | FN-SEC-COMPLIANCE-01     | コンプライアンス状況確認    | AWS Config等による設定のコンプライアンスチェック結果確認。  | FE: (AWS Console or custom view), BE: AWS Config                                    |     ❌      | US-SEC-COMPLIANCE-01               |
| **BI分析 (アナリスト)**    | FN-BA-AUTH-01            | 分析ツールログイン          | BIツールやデータウェアハウスへのログイン。                  | FE: (BI Tool), BE: (BI Tool Auth)                                                   |     ⚪️      | US-BA-AUTH-LOGIN-01                |
|                            | FN-BA-ANALYZE-PROD-01    | 商品パフォーマンス分析      | 売上、閲覧数、カート追加率など商品別分析。                  | FE: (BI Tool), BE: DWH/Data Lake, QuickSight                                        |     ❌      | US-BA-BI-PROD-01                   |
|                            | FN-BA-ANALYZE-SALES-01   | 売上・トラフィック分析      | 売上とサイトアクセスデータの相関分析。                      | FE: (BI Tool), BE: DWH/Data Lake, QuickSight                                        |     ❌      | US-BA-BI-SALES-01                  |
| **システム基盤**           | FN-SYS-NOTIF-01          | 基本通知送信                | 注文確認メール等送信。                                      | BE: Lambda_Order, SNS/SES                                                           |     ✅      | US-DEV-NOTIF-IMPL-01               |
|                            | FN-SYS-ASYNC-01          | 基本非同期処理              | SQS経由でのLambda/Fargate連携。                             | BE: Lambda_Order, SQS, Fargate_Stock                                                |     ✅      | US-DEV-ASYNC-IMPL-01               |
|                            | FN-SYS-IMAGE-01          | 基本画像処理                | 画像アップロード時の処理(バリデーション等)。                | BE: (Lambda_Product or dedicated Lambda)                                            |     ✅      | US-DEV-IMAGE-IMPL-01               |
|                            | FN-SYS-CICD-01           | CI/CDパイプライン           | テスト、ビルド、DEV環境への自動デプロイ。                   | GitHub Actions                                                                      |     ✅      | US-DEV-CICD-IMPL-01                |
|                            | FN-SYS-BATCH-01          | データエクスポート          | システムデータの定期的/オンデマンドエクスポート。           | BE: Step Functions / Batch                                                          |     ❌      | US-SRE-BATCH-01                    |
|                            | FN-SYS-BATCH-02          | 定期メンテナンス            | DB最適化などの定期メンテナンスタスク実行。                  | BE: Step Functions / Lambda (Scheduled)                                             |     ❌      | US-SRE-BATCH-02                    |
|                            | FN-SYS-BATCH-03          | 定期レポート生成            | 在庫レポートなどの定期バッチ生成。                          | BE: Step Functions / Lambda (Scheduled)                                             |     ⚪️      | US-SRE-BATCH-03                    |
| **オブザーバビリティ**     | FN-SYS-OBS-LOG-01        | 構造化ログ計装・収集        | 各コンポーネントからの構造化ログ出力・集約。                | BE: 各コンポーネント, OBS: CloudWatch Logs                                          |     ✅      | US-DEV-OBS-LOG-IMPL-01             |
|                            | FN-SYS-OBS-MET-01        | メトリクス計装・収集        | 標準/カスタムメトリクス(EMF)送信。                          | BE: 各コンポーネント, OBS: CloudWatch Metrics                                       |     ✅      | US-DEV-OBS-METRIC-IMPL-01          |
|                            | FN-SYS-OBS-TRC-01        | 分散トレーシング計装・収集  | リクエストパスのトレース情報送信。                          | BE: API GW, Lambda, Fargate, OBS: X-Ray SDK                                         |     ✅      | US-DEV-OBS-TRACE-IMPL-01           |
|                            | FN-SYS-OBS-INFRA-01      | オブザーバビリティインフラ  | ロググループ、アラーム等をIaC構築。                         | Terraform, OBS: CloudWatch, X-Ray                                                   |     ✅      | US-SRE-OBS-IMPL-01                 |
|                            | FN-SRE-OBS-DASH-01       | ヘルスダッシュボード利用    | システム状態をダッシュボードで確認。                        | FE: (CloudWatch Console), OBS: CloudWatch Dashboards                                |     ⚪️      | US-SRE-OBS-DASH-01                 |
|                            | FN-SRE-OBS-ALERT-01      | アラート受信・対応          | システム異常アラートを受信し対応。                          | OBS: CloudWatch Alarms, SNS                                                         |     ⚪️      | US-SRE-OBS-ALERT-01                |
|                            | FN-SRE-OBS-TREND-01      | パフォーマンストレンド分析  | メトリクスの長期トレンドを分析。                            | FE: (CloudWatch Console), OBS: CloudWatch Metrics                                   |     ⚪️      | US-SRE-OBS-TREND-01                |
|                            | FN-SRE-OBS-LAMBDA-01     | Lambda監視・最適化          | Lambda関数の実行状況・性能・コストを分析・最適化。          | FE: (CloudWatch Console), OBS: CloudWatch Logs/Metrics/X-Ray                        |     ⚪️      | US-SRE-LAMBDA-01                   |
|                            | FN-SRE-OBS-COST-01       | コスト監視・最適化          | AWS利用コスト（特にオブザーバビリティ関連）を分析・最適化。 | FE: (AWS Cost Explorer), OBS: Cost Explorer                                         |     ⚪️      | US-SRE-COST-01                     |
|                            | FN-SRE-OBS-RUM-01        | RUMデータ活用・分析         | フロントエンドのリアルユーザー体験データを分析。            | FE: (CloudWatch Console), OBS: CloudWatch RUM                                       |     ⚪️      | US-SRE-RUM-01                      |
|                            | FN-SRE-OBS-INCIDENT-01   | インシデント調査            | ログ・メトリクス・トレースを横断的に調査し根本原因特定。    | FE: (CloudWatch/X-Ray Console), OBS: CW Logs/Metrics, X-Ray                         |     ✅      | US-SRE-OBS-INCIDENT-INVESTIGATE-01 |
|                            | FN-SRE-OBS-VIS-01        | 高度な可視化・分析          | 複数データを組み合わせて複雑な問題を可視化・分析。          | FE: (CloudWatch/X-Ray Console/Grafana?), OBS: CW Logs/Metrics, X-Ray                |     ✅      | US-SRE-VIS-01                      |
|                            | FN-DEV-OBS-OTEL-01       | OpenTelemetry計装           | OTel SDKを用いた計装実装。                                  | BE: 各コンポーネント, OBS: OTel SDK, ADOT Collector                                 |     ⚪️      | US-DEV-OBS-OTEL-IMPL-01            |
|                            | FN-SRE-OBS-AUTO-01       | 自動修復メカニズム          | 既知の問題に対する自動修復処理実装。                        | BE: Lambda/Systems Manager, OBS: CloudWatch Events/Alarms                           |     ❌      | US-SRE-AUTO-IMPL-01                |
|                            | FN-SRE-OBS-CHAOS-01      | カオスエンジニアリング検証  | FIS等を用いた意図的な障害注入と耐障害性検証。               | OBS: AWS FIS, CloudWatch                                                            |     ❌      | US-SRE-CHAOS-01                    |
|                            | FN-SRE-OBS-UX-01         | E2Eユーザー体験監視         | Synthetics等を用いた主要ユーザージャーニーの監視。          | OBS: CloudWatch Synthetics                                                          |     ❌      | US-SRE-UX-01                       |
|                            | FN-SRE-OBS-CAPACITY-01   | キャパシティ分析・計画      | リソース使用率分析に基づくキャパシティプランニング。        | OBS: CloudWatch Metrics                                                             |     ❌      | US-SRE-CAPACITY-01                 |
