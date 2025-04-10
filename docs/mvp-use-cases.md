# 1. MVPユースケース一覧

## 1.1. 顧客向けユースケース

| ID     | ユースケース名     | 説明                                         | 優先度 | 関連機能    | 完了基準                                         |
| ------ | ------------------ | -------------------------------------------- | ------ | ----------- | ------------------------------------------------ |
| CUS-01 | 基本的な商品閲覧   | ユーザーは商品リストを閲覧できる             | 高     | C-BROWSE-01 | 商品一覧ページで実際の商品データが表示される     |
| CUS-02 | 商品詳細確認       | ユーザーは商品の詳細情報を確認できる         | 高     | C-BROWSE-02 | 商品詳細ページで画像、説明、価格などが表示される |
| CUS-03 | カテゴリー別閲覧   | ユーザーはカテゴリーで商品を絞り込める       | 中     | C-BROWSE-03 | カテゴリー選択による商品フィルタリングが機能する |
| CUS-04 | 在庫確認           | ユーザーは商品の在庫状況を確認できる         | 中     | C-BROWSE-07 | 商品詳細ページで在庫状況が表示される             |
| CUS-05 | カートへの商品追加 | ユーザーは商品をカートに追加できる           | 高     | C-SHOP-01   | カートに商品を追加でき、カート内に表示される     |
| CUS-06 | カート管理         | ユーザーはカート内の商品を管理できる         | 高     | C-SHOP-01   | カート内の商品数量変更・削除が機能する           |
| CUS-07 | 注文手続き         | ユーザーは配送先情報を入力し注文できる       | 高     | C-SHOP-02   | 配送先入力から注文確認までのフローが機能する     |
| CUS-08 | 注文完了           | ユーザーは注文を完了し、注文番号を受け取れる | 高     | C-SHOP-02   | 注文処理が完了し、注文番号が発行される           |

## 1.2. 管理者向けユースケース

| ID     | ユースケース名 | 説明                             | 優先度 | 関連機能  | 完了基準                                                     |
| ------ | -------------- | -------------------------------- | ------ | --------- | ------------------------------------------------------------ |
| ADM-01 | 管理者認証     | 管理者はシステムにログインできる | 中     | -         | 有効な認証情報でログインでき、無効な場合はエラーが表示される |
| ADM-02 | 商品情報管理   | 管理者は商品情報を管理できる     | 高     | A-PROD-01 | 商品の一覧表示、検索、編集、削除が機能する                   |
| ADM-03 | 新規商品登録   | 管理者は新しい商品を登録できる   | 高     | A-PROD-02 | 新規商品情報の入力と登録が完了する                           |
| ADM-04 | 在庫レベル監視 | 管理者は在庫レベルを監視できる   | 中     | A-INV-01  | 在庫レベルのダッシュボード表示が機能する                     |
| ADM-05 | 在庫更新       | 管理者は在庫数を更新できる       | 中     | A-INV-03  | 入荷登録や在庫調整が機能する                                 |

## 1.3. オブザーバビリティユースケース

| ID     | ユースケース名         | 説明                                                                 | 優先度 | 関連機能       | 完了基準                                                |
| ------ | ---------------------- | -------------------------------------------------------------------- | ------ | -------------- | ------------------------------------------------------- |
| OBS-01 | 構造化ログ収集         | システムログが構造化形式で収集される                                 | 高     | O-LOG-01       | JSONフォーマットのログがCloudWatch Logsに格納される     |
| OBS-02 | ログレベル管理         | 適切なログレベルでログが記録される                                   | 高     | O-LOG-02       | ERROR/WARN/INFO/DEBUGの各レベルのログが適切に記録される |
| OBS-03 | リクエストコンテキスト | リクエストにコンテキスト情報が付与される                             | 中     | O-LOG-04       | ログにリクエストID、ユーザーIDなどが含まれる            |
| OBS-04 | 基本メトリクス収集     | リクエスト数、レイテンシー、エラー率などの基本メトリクスが収集される | 高     | O-METRIC-01/02 | CloudWatch Metricsで基本メトリクスが確認できる          |
| OBS-05 | 分散トレース           | リクエストの処理フローがトレースされる                               | 高     | O-TRACE-01/02  | X-Rayでリクエストの処理フローが可視化される             |
| OBS-06 | サービスマップ可視化   | サービス間の依存関係が可視化される                                   | 中     | O-DASH-02      | X-Rayサービスマップでサービス間の関係が表示される       |
| OBS-07 | ヘルスチェック         | システムのヘルス状態が監視される                                     | 高     | O-ALERT-01     | ヘルスチェックエンドポイントが機能し、状態が報告される  |
| OBS-08 | アラート通知           | 異常時にアラートが発報される                                         | 中     | O-ALERT-02/03  | 設定した閾値を超えるとSNS通知が送信される               |
| OBS-09 | OpenTelemetry統合      | OTELでデータ収集と可視化ができる                                     | 高     | O-COMP-01      | ADOTでデータ収集し、X-Rayで表示できる                   |
