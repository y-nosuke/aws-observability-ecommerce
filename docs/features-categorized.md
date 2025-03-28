# 1. AWSオブザーバビリティ学習用eコマースアプリ - 分類別機能一覧

- [1. AWSオブザーバビリティ学習用eコマースアプリ - 分類別機能一覧](#1-awsオブザーバビリティ学習用eコマースアプリ---分類別機能一覧)
  - [1.1. 顧客向け機能](#11-顧客向け機能)
    - [1.1.1. 商品閲覧・検索](#111-商品閲覧検索)
    - [1.1.2. ショッピングとチェックアウト](#112-ショッピングとチェックアウト)
    - [1.1.3. ユーザー管理](#113-ユーザー管理)
    - [1.1.4. 通知機能](#114-通知機能)
  - [1.2. 管理者向け機能](#12-管理者向け機能)
    - [1.2.1. 商品管理](#121-商品管理)
    - [1.2.2. 在庫管理](#122-在庫管理)
    - [1.2.3. 注文・出荷管理](#123-注文出荷管理)
    - [1.2.4. ユーザー管理](#124-ユーザー管理)
    - [1.2.5. 通知送信](#125-通知送信)
    - [1.2.6. バッチ処理機能](#126-バッチ処理機能)
  - [1.3. オブザーバビリティ機能](#13-オブザーバビリティ機能)
    - [1.3.1. ログ関連](#131-ログ関連)
    - [1.3.2. メトリクス収集](#132-メトリクス収集)
    - [1.3.3. 分散トレース](#133-分散トレース)
    - [1.3.4. アラートと可視化](#134-アラートと可視化)
    - [1.3.5. 相関・連携機能](#135-相関連携機能)
    - [1.3.6. Lambda固有機能](#136-lambda固有機能)
    - [1.3.7. インフラストラクチャ機能](#137-インフラストラクチャ機能)
    - [1.3.8. RUMと合成監視](#138-rumと合成監視)
    - [1.3.9. 耐障害性とカオスエンジニアリング](#139-耐障害性とカオスエンジニアリング)
    - [1.3.10. 比較と最適化](#1310-比較と最適化)
    - [1.3.11. 高度なオブザーバビリティ](#1311-高度なオブザーバビリティ)
  - [1.4. 期待される成果物](#14-期待される成果物)

## 1.1. 顧客向け機能

### 1.1.1. 商品閲覧・検索

| 機能ID          | 機能名             | 概要                                       | 優先度   |
| --------------- | ------------------ | ------------------------------------------ | -------- |
| **C-BROWSE-01** | 商品一覧表示       | ユーザーが閲覧可能な商品リストの表示       | MVP      |
| **C-BROWSE-02** | 商品詳細表示       | 個別商品の詳細情報と画像の表示             | MVP      |
| **C-BROWSE-03** | カテゴリー別表示   | カテゴリーに基づいた商品の分類と表示       | MVP      |
| **C-BROWSE-04** | 商品検索           | キーワードによる商品検索と結果表示         | 拡張機能 |
| **C-BROWSE-05** | 商品レビュー       | ユーザーによる商品評価とレビュー投稿・閲覧 | 拡張機能 |
| **C-BROWSE-06** | 人気商品ランキング | 閲覧数・購入数に基づく人気商品の表示       | 拡張機能 |
| **C-BROWSE-07** | 在庫状況表示       | 商品の在庫状況と在庫不足の表示             | MVP      |

### 1.1.2. ショッピングとチェックアウト

| 機能ID        | 機能名                 | 概要                                       | 優先度   |
| ------------- | ---------------------- | ------------------------------------------ | -------- |
| **C-SHOP-01** | カート機能             | 商品の追加/削除とカート内容表示            | MVP      |
| **C-SHOP-02** | 注文処理               | カートからチェックアウトと注文確定プロセス | MVP      |
| **C-SHOP-03** | 支払い処理（模擬）     | 支払い方法選択と支払いプロセスの模擬実装   | 拡張機能 |
| **C-SHOP-04** | 注文履歴・トラッキング | 過去の注文内容と配送状況の追跡             | 拡張機能 |

### 1.1.3. ユーザー管理

| 機能ID        | 機能名             | 概要                                          | 優先度   |
| ------------- | ------------------ | --------------------------------------------- | -------- |
| **C-USER-01** | 会員登録・ログイン | ユーザーアカウント作成とJWT認証によるログイン | 拡張機能 |
| **C-USER-02** | プロフィール管理   | ユーザープロフィール情報の表示と編集          | 拡張機能 |
| **C-USER-03** | お届け先住所管理   | 配送先住所の登録・編集・選択                  | 拡張機能 |
| **C-USER-04** | 退会処理           | ユーザーアカウントの削除とデータの取り扱い    | 拡張機能 |

### 1.1.4. 通知機能

| 機能ID          | 機能名             | 概要                           | 優先度   |
| --------------- | ------------------ | ------------------------------ | -------- |
| **C-NOTIFY-01** | 通知設定           | メール通知の設定と管理         | 拡張機能 |
| **C-NOTIFY-02** | 注文確認メール受信 | 注文完了時の自動確認メール受信 | 拡張機能 |
| **C-NOTIFY-03** | 出荷通知受信       | 商品出荷時の自動通知メール受信 | 拡張機能 |
| **C-NOTIFY-04** | 在庫状況通知受信   | 在庫切れ商品の入荷時通知受信   | 拡張機能 |

## 1.2. 管理者向け機能

### 1.2.1. 商品管理

| 機能ID        | 機能名           | 概要                                             | 優先度   |
| ------------- | ---------------- | ------------------------------------------------ | -------- |
| **A-PROD-01** | 商品情報管理     | RDS (Aurora) を使用した商品データの基本的な管理  | MVP      |
| **A-PROD-02** | 商品登録機能     | 管理者による新規商品の登録と情報入力             | MVP      |
| **A-PROD-03** | 商品一括管理     | 複数商品の一括登録・更新・削除機能               | 拡張機能 |
| **A-PROD-04** | 商品カテゴリ管理 | 商品カテゴリの作成・編集・階層管理               | 拡張機能 |
| **A-PROD-05** | 商品画像処理     | アップロードされた商品画像のリサイズと最適化処理 | MVP      |

### 1.2.2. 在庫管理

| 機能ID       | 機能名           | 概要                                       | 優先度   |
| ------------ | ---------------- | ------------------------------------------ | -------- |
| **A-INV-01** | 在庫レベル監視   | 全商品の在庫状況の監視とダッシュボード表示 | MVP      |
| **A-INV-02** | 在庫不足アラート | 設定済みしきい値を下回る在庫のアラート通知 | 拡張機能 |
| **A-INV-03** | 在庫更新         | 在庫数の手動・自動更新機能                 | MVP      |

### 1.2.3. 注文・出荷管理

| 機能ID         | 機能名         | 概要                               | 優先度   |
| -------------- | -------------- | ---------------------------------- | -------- |
| **A-ORDER-01** | 注文確認・処理 | 新規注文の確認と処理ステータス管理 | 拡張機能 |
| **A-ORDER-02** | 出荷管理       | 出荷情報の登録と出荷ステータス管理 | 拡張機能 |

### 1.2.4. ユーザー管理

| 機能ID        | 機能名           | 概要                                 | 優先度   |
| ------------- | ---------------- | ------------------------------------ | -------- |
| **A-USER-01** | ユーザー情報管理 | 管理者によるユーザー情報の閲覧・編集 | 拡張機能 |
| **A-USER-02** | 退会ユーザー復帰 | 退会済みユーザーのアカウント復元処理 | 拡張機能 |

### 1.2.5. 通知送信

| 機能ID          | 機能名             | 概要                                                    | 優先度   |
| --------------- | ------------------ | ------------------------------------------------------- | -------- |
| **A-NOTIFY-01** | 非同期通知         | メール通知(SQS + Go Lambda + SES)とイベント駆動アラート | 拡張機能 |
| **A-NOTIFY-02** | 注文確認メール送信 | 注文完了時の自動確認メール送信                          | 拡張機能 |
| **A-NOTIFY-03** | 出荷通知送信       | 商品出荷時の自動通知メール送信                          | 拡張機能 |
| **A-NOTIFY-04** | 在庫状況通知送信   | 在庫切れ商品の入荷時通知送信                            | 拡張機能 |

### 1.2.6. バッチ処理機能

| 機能ID         | 機能名                 | 概要                                               | 優先度   |
| -------------- | ---------------------- | -------------------------------------------------- | -------- |
| **A-BATCH-01** | 在庫レポート生成       | 定期的な在庫状況レポートの自動生成と管理者への配信 | MVP      |
| **A-BATCH-02** | 商品データバックアップ | 商品データの定期的なバックアップ処理               | 拡張機能 |

## 1.3. オブザーバビリティ機能

### 1.3.1. ログ関連

| 機能ID       | 機能名                         | 概要                                           | 対応するAWSサービス      | 優先度 |
| ------------ | ------------------------------ | ---------------------------------------------- | ------------------------ | ------ |
| **O-LOG-01** | 構造化ログ設定                 | JSON形式での統一的なログ形式実装と収集         | CloudWatch Logs          | MVP    |
| **O-LOG-02** | ログレベル管理                 | ERROR/WARN/INFO/DEBUGの適切な使い分け          | CloudWatch Logs          | MVP    |
| **O-LOG-03** | CloudWatch Logsへのログ転送    | LocalStackを使用した開発環境でのログ転送設定   | CloudWatch Logs          | MVP    |
| **O-LOG-04** | コンテキスト情報付与           | リクエストID、ユーザーIDなどの情報をログに追加 | CloudWatch Logs          | MVP    |
| **O-LOG-05** | CloudWatch Logs Insightsクエリ | ログデータの効率的な検索と分析用クエリの作成   | CloudWatch Logs Insights | MVP    |
| **O-LOG-06** | 本番環境ログ設定               | 本番環境でのCloudWatch Logs設定と最適化        | CloudWatch Logs          | MVP    |

### 1.3.2. メトリクス収集

| 機能ID          | 機能名                        | 概要                                                        | 対応するAWSサービス             | 優先度 |
| --------------- | ----------------------------- | ----------------------------------------------------------- | ------------------------------- | ------ |
| **O-METRIC-01** | 基本メトリクス収集            | リクエスト数、レイテンシー、エラー率の測定                  | CloudWatch Metrics              | MVP    |
| **O-METRIC-02** | インフラメトリクス収集        | CPU/メモリ使用率、ディスクI/O、ネットワークスループット計測 | CloudWatch Metrics              | MVP    |
| **O-METRIC-03** | レスポンスタイム分布測定      | p50, p90, p95, p99パーセンタイルの計測                      | CloudWatch Metrics              | MVP    |
| **O-METRIC-04** | ビジネスメトリクス収集        | 閲覧数、購入数など業務関連指標の収集                        | CloudWatch Metrics              | MVP    |
| **O-METRIC-05** | CloudWatch Custom Metrics連携 | カスタムメトリクスのCloudWatchへの送信                      | CloudWatch Metrics              | MVP    |
| **O-METRIC-06** | ビジネスメトリクスの拡張      | 高度なビジネス指標の追加と可視化                            | CloudWatch Metrics, EventBridge | MVP    |
| **O-METRIC-07** | メトリクスサービス実装        | 独立したメトリクス収集・管理サービス                        | CloudWatch Metrics              | MVP    |
| **O-METRIC-08** | フロントエンドユーザー体験    | フロントエンドのユーザー体験に関わるメトリクス              | CloudWatch RUM                  | MVP    |
| **O-METRIC-09** | フロントエンドパフォーマンス  | フロントエンドの応答性能に関わるメトリクス                  | CloudWatch RUM                  | MVP    |
| **O-METRIC-10** | 管理アクションのメトリクス    | 管理操作の頻度や種類の計測                                  | CloudWatch Metrics              | MVP    |
| **O-METRIC-11** | リソース使用状況モニタリング  | サーバーやDBのリソース使用状況の監視                        | CloudWatch Metrics              | MVP    |

### 1.3.3. 分散トレース

| 機能ID         | 機能名                       | 概要                                              | 対応するAWSサービス | 優先度 |
| -------------- | ---------------------------- | ------------------------------------------------- | ------------------- | ------ |
| **O-TRACE-01** | X-Ray基本統合                | AWS X-Rayによる基本的な分散トレース実装           | AWS X-Ray           | MVP    |
| **O-TRACE-02** | Echo/sqlboiler統合           | ウェブフレームワークとORMへのトレース統合         | AWS X-Ray           | MVP    |
| **O-TRACE-03** | トレースフィルタリング       | 特定のリクエストやエラーケースを効率的に検索      | AWS X-Ray           | MVP    |
| **O-TRACE-04** | 管理操作カスタムトレース     | 管理機能に特化したカスタムトレースの実装          | AWS X-Ray           | MVP    |
| **O-TRACE-05** | アップロード進捗トラッキング | ファイルアップロード処理の進捗トラッキング        | AWS X-Ray           | MVP    |
| **O-TRACE-06** | 本番環境X-Ray設定            | 本番環境でのX-Ray設定とサンプリングレートの最適化 | AWS X-Ray           | MVP    |

### 1.3.4. アラートと可視化

| 機能ID         | 機能名                       | 概要                                      | 対応するAWSサービス    | 優先度 |
| -------------- | ---------------------------- | ----------------------------------------- | ---------------------- | ------ |
| **O-ALERT-01** | ヘルスチェックエンドポイント | サービス稼働状況確認用APIエンドポイント   | CloudWatch, Route 53   | MVP    |
| **O-ALERT-02** | 基本アラート設定             | エラー率や異常なレイテンシーの検出と通知  | CloudWatch Alarms      | MVP    |
| **O-ALERT-03** | アラートとSNS通知            | アラート発生時のSNSを使った通知設定       | CloudWatch Alarms, SNS | MVP    |
| **O-ALERT-04** | 障害検出とアラート検証       | 障害シナリオにおけるアラート動作の検証    | CloudWatch Alarms      | MVP    |
| **O-DASH-01**  | X-Rayダッシュボード          | X-Rayの分析データを活用したダッシュボード | AWS X-Ray              | MVP    |
| **O-DASH-02**  | サービスマップ強化           | サービス間依存関係の詳細な可視化          | AWS X-Ray              | MVP    |
| **O-DASH-03**  | REDダッシュボード            | Rate, Error, Duration指標のダッシュボード | CloudWatch Dashboards  | MVP    |
| **O-DASH-04**  | カスタムダッシュボード       | 目的別にカスタマイズしたダッシュボード    | CloudWatch Dashboards  | MVP    |

### 1.3.5. 相関・連携機能

| 機能ID        | 機能名               | 概要                                     | 対応するAWSサービス           | 優先度 |
| ------------- | -------------------- | ---------------------------------------- | ----------------------------- | ------ |
| **O-CORR-01** | 観測データの相関付け | ログ、メトリクス、トレースの相関関係表示 | CloudWatch ServiceLens, X-Ray | MVP    |

### 1.3.6. Lambda固有機能

| 機能ID          | 機能名               | 概要                                 | 対応するAWSサービス         | 優先度 |
| --------------- | -------------------- | ------------------------------------ | --------------------------- | ------ |
| **O-LAMBDA-01** | Lambda実行ログ構造化 | Lambda関数の構造化されたログ出力     | CloudWatch Logs             | MVP    |
| **O-LAMBDA-02** | Lambda固有メトリクス | Lambda関数の実行時間、使用メモリなど | CloudWatch, Lambda Insights | MVP    |

### 1.3.7. インフラストラクチャ機能

| 機能ID         | 機能名                         | 概要                                              | 対応するAWSサービス            | 優先度 |
| -------------- | ------------------------------ | ------------------------------------------------- | ------------------------------ | ------ |
| **O-INFRA-01** | オブザーバビリティインフラ構築 | Terraform定義に組み込まれたオブザーバビリティ設定 | CloudWatch, X-Ray, EventBridge | MVP    |

### 1.3.8. RUMと合成監視

| 機能ID         | 機能名                    | 概要                                     | 対応するAWSサービス   | 優先度 |
| -------------- | ------------------------- | ---------------------------------------- | --------------------- | ------ |
| **O-RUM-01**   | CloudWatch RUM設定        | 実ユーザーモニタリングの設定と分析       | CloudWatch RUM        | MVP    |
| **O-SYNTH-01** | CloudWatch Synthetics設定 | 合成モニタリングによるユーザーフロー検証 | CloudWatch Synthetics | MVP    |

### 1.3.9. 耐障害性とカオスエンジニアリング

| 機能ID         | 機能名                      | 概要                                     | 対応するAWSサービス         | 優先度 |
| -------------- | --------------------------- | ---------------------------------------- | --------------------------- | ------ |
| **O-FAULT-01** | Fault Injection Service設定 | 故障注入サービスの基本設定               | AWS Fault Injection Service | MVP    |
| **O-FAULT-02** | 高度な障害シナリオ          | 複合的な障害シナリオの設計と実行         | AWS Fault Injection Service | MVP    |
| **O-CHAOS-01** | カオスエンジニアリング実験  | 制御された環境でのシステム耐障害性テスト | AWS Fault Injection Service | MVP    |

### 1.3.10. 比較と最適化

| 機能ID        | 機能名                           | 概要                                       | 対応するAWSサービス            | 優先度 |
| ------------- | -------------------------------- | ------------------------------------------ | ------------------------------ | ------ |
| **O-COMP-01** | オブザーバビリティアプローチ比較 | AWS SDK v2とOpenTelemetryの実装比較と評価  | X-Ray, ADOT                    | MVP    |
| **O-COST-01** | コスト最適化                     | オブザーバビリティ関連コストの分析と最適化 | AWS Cost Explorer, AWS Budgets | MVP    |

### 1.3.11. 高度なオブザーバビリティ

| 機能ID       | 機能名                 | 概要                                               | 対応するAWSサービス            | 優先度   |
| ------------ | ---------------------- | -------------------------------------------------- | ------------------------------ | -------- |
| **O-ADV-01** | 異常検出実装           | 動的ベースラインによる異常検出                     | CloudWatch Anomaly Detection   | 拡張機能 |
| **O-ADV-02** | ML駆動問題検出         | 運用メトリクスの自動分析と潜在問題検知             | Amazon DevOps Guru             | 拡張機能 |
| **O-ADV-03** | 根本原因分析自動化     | 関連リソースとイベントの相関分析                   | Amazon DevOps Guru, CloudWatch | 拡張機能 |
| **O-ADV-04** | シナリオベースアラート | 複合条件に基づく高度なアラート設定                 | CloudWatch Composite Alarms    | 拡張機能 |
| **O-ADV-05** | Grafanaダッシュボード  | 統合可視化ダッシュボードの作成と管理               | Amazon Managed Grafana         | 拡張機能 |
| **O-ADV-06** | 大規模ログ分析         | OpenSearch Serviceによる高度なログ分析             | Amazon OpenSearch Service      | 拡張機能 |
| **O-ADV-07** | Web Vitals測定         | FCP、LCP、CLSなどのWeb Vitals指標の測定            | CloudWatch RUM                 | 拡張機能 |
| **O-ADV-08** | カスタム分析レポート   | 大規模データに対するアドホッククエリとレポート生成 | Amazon Athena, QuickSight      | 拡張機能 |

## 1.4. 期待される成果物

1. **フルスタックeコマースアプリケーション**
   - Go/Echo + Next.jsで実装された機能的なMVPアプリケーション
   - AWSマネージドサービスを活用したスケーラブルなバックエンド
   - レスポンシブで使いやすいフロントエンド

2. **包括的なオブザーバビリティ実装の参照実装**
   - ログ、メトリクス、トレースの3本柱を網羅した実装
   - AWS SDK v2とOpenTelemetryの2つのアプローチによる比較環境
   - 統合ダッシュボードによる可視化

3. **マイクロサービスアーキテクチャのベストプラクティス**
   - 疎結合でスケーラブルなサービス設計
   - イベント駆動型処理の効果的な活用
   - ステートレスサービスと適切なDBアクセスパターン

4. **AWSサービス連携パターンの実例**
   - ECS/Fargate、Lambda、RDS、S3/CloudFrontの連携
   - CloudWatch、X-Ray、EventBridgeによるオブザーバビリティ
   - Terraformによるインフラストラクチャのコード化

5. **個人学習用のドキュメントとケーススタディ**
   - 実装フェーズごとの学習ポイントと自己評価フレームワーク
   - 障害シナリオと対応手順の実例集
   - オブザーバビリティアプローチ比較のケーススタディ
