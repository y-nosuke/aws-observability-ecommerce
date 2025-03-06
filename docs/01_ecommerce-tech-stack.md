# AWSオブザーバビリティ学習用eコマースアプリ - 技術要素

## インフラストラクチャ

- **フロントエンド**: Next.js、S3、CloudFront
- **API管理**: API Gateway、Application Load Balancer (ALB)
- **コンピューティング**: Lambda、Fargate (ECS)
- **データストア**: DynamoDB、RDS (PostgreSQL/MySQL)
- **ストレージ**: S3
- **メッセージング**: SQS、SNS、EventBridge
- **通知**: SES
- **オブザーバビリティ**:
  - **モニタリング基盤**: CloudWatch (Logs, Metrics, Alarms, Dashboards, Synthetics), X-Ray
  - **高度な分析**: Amazon DevOps Guru, Amazon Managed Service for Prometheus (AMP), Amazon Managed Grafana (AMG)
  - **標準化**: AWS Distro for OpenTelemetry (ADOT)
  - **監査・コンプライアンス**: AWS CloudTrail, AWS Config
  - **データ分析**: Amazon OpenSearch Service, Amazon Athena
  - **ネットワーク監視**: VPC Flow Logs
  - **テスト・実験**: AWS Fault Injection Service
  - **フロントエンド監視**: Amazon CloudWatch RUM (Real User Monitoring)
  - **サービスヘルス**: AWS Health Dashboard

## 開発言語・フレームワーク

- **フロントエンド**: TypeScript、Next.js、React
- **バックエンド**: Go（Fargate）、Node.js（Lambda）

## AWSオブザーバビリティサービスの概要と用途

| サービス | 主な用途 |
|---------|---------|
| **CloudWatch Logs** | アプリケーションログの集約、保存、検索 |
| **CloudWatch Metrics** | 基本的なシステムとアプリケーションメトリクスの収集 |
| **CloudWatch Alarms** | 基本的なリソース使用率とエラー率の監視 |
| **AWS X-Ray** | 分散トレーシング、サービス間の依存関係の可視化 |
| **CloudWatch Dashboards** | カスタムダッシュボードの作成、主要指標の可視化 |
| **AWS Distro for OpenTelemetry (ADOT)** | 標準化されたテレメトリデータの収集 |
| **VPC Flow Logs** | ネットワークトラフィックの可視化、セキュリティ分析 |
| **CloudWatch Synthetics** | エンドツーエンドのユーザーフロー監視、カナリアテスト |
| **CloudWatch RUM** | 実際のユーザー体験とパフォーマンスの測定 |
| **Amazon Managed Service for Prometheus (AMP)** | 高度なメトリクス収集とクエリ |
| **Amazon Managed Grafana (AMG)** | 高度なダッシュボードと可視化 |
| **Amazon DevOps Guru** | ML駆動の異常検出と問題の自動診断 |
| **AWS CloudTrail** | API呼び出しの監査とコンプライアンス |
| **AWS Config** | インフラストラクチャの構成変更の追跡 |
| **Amazon OpenSearch Service** | 大規模なログデータの検索と分析 |
| **Amazon Athena** | ログデータへのSQLクエリとアドホック分析 |
| **AWS Fault Injection Service** | カオスエンジニアリングとレジリエンステスト |
| **AWS Health Dashboard** | AWSサービスのヘルスステータス監視 |
| **EventBridge** | イベント駆動型のオブザーバビリティとアラート連携 |