# AWSオブザーバビリティ学習用eコマースアプリ

AWSオブザーバビリティのパターンを学習するためのeコマースアプリケーションの参照実装です。Go Echo+sqlboiler+Next.jsのスタックと、LocalStackによるAWSサービスのエミュレーションを活用しています。

## 概要

このプロジェクトは、AWSのオブザーバビリティサービス（CloudWatch、X-Ray）を学習するための実践的な環境を提供します。eコマースアプリケーションの基本機能を実装しながら、以下の2つのアプローチでオブザーバビリティを実装・比較します：

1. **AWS SDK v2アプローチ**：AWSのネイティブSDKを使用
2. **OpenTelemetryアプローチ**：ベンダー中立なOTEL SDKを使用

## 前提条件

- Docker と Docker Compose
- Go 1.24以上
- Node.js 23以上
- AWS CLI
- Git
- Terraform
- LocalStack CLI (`pip install localstack`)
- LocalStack Desktop（[ダウンロードページ](https://app.localstack.cloud/resources/desktop)からインストール）

## プロジェクト構成

```text
aws-observability-ecommerce/
├── backend/           # Goバックエンド
├── frontend/          # Next.jsフロントエンド
├── infra/             # インフラストラクチャコード
│   ├── localstack/    # LocalStack設定
│   └── terraform/     # Terraformコード
└── docs/              # ドキュメント
```

## クイックスタート

1. リポジトリをクローン：

  ```bash
  git clone https://github.com/y-nosuke/aws-observability-ecommerce.git
  cd aws-observability-ecommerce
  ```

1. Docker Composeで環境を起動：

```bash
docker-compose up -d
```

1. バックエンドとフロントエンドの状態を確認：

- バックエンド：<http://backend.localhost/api/health>
- フロントエンド（顧客画面）：<http://customer.localhost>
- フロントエンド（管理者画面）：<http://admin.localhost>

## 開発環境

このプロジェクトは、以下の開発環境を設定しています：

- **Go Echo** バックエンドAPI（ホットリロード対応）
- **MySQL** データベース
- **Next.js** フロントエンド（ホットリロード対応）
- **LocalStack** AWSサービスエミュレーター

## オブザーバビリティの実装

### AWS SDK v2アプローチ

- CloudWatch Logsによるログ収集
- X-Ray SDKによる分散トレース
- CloudWatch Metricsによるメトリクス収集

### OpenTelemetryアプローチ

- OpenTelemetry SDK + ADOT Collector
- クラウドベンダーに依存しない計装
- X-Ray、CloudWatch Metricsへのエクスポート

## 機能一覧

MVP（最小実装）には、以下の機能が含まれています：

- 商品カタログの閲覧・検索
- 商品詳細の表示
- カートへの商品追加
- 簡易的な注文処理
- 商品と在庫の管理（管理者機能）

## フェーズ別実装計画

本プロジェクトは以下のフェーズで実装されています：

1. **フェーズ1**: 開発環境のセットアップとプロジェクト骨組み構築
2. **フェーズ2**: バックエンド実装
3. **フェーズ3**: フロントエンド実装
4. **フェーズ4**: AWS環境へのデプロイとオブザーバビリティ強化
5. **フェーズ5**: OpenTelemetryアプローチの実装
6. **フェーズ6**: オブザーバビリティ比較と最終調整

## 貢献方法

貢献に興味がある場合は、[CONTRIBUTING.md](CONTRIBUTING.md)を参照してください。

## ライセンス

MITライセンスの下で公開されています。詳細は[LICENSE](LICENSE)ファイルを参照してください。

## LocalStackを使用したCloudWatch Logsのセットアップ

### セットアップ手順

1. LocalStackの起動

    ```bash
    docker-compose up -d localstack
    ```

2. CloudWatch Logsロググループの作成

    ```bash
    cd infra/terraform
    tflocal init
    tflocal apply
    ```

3. ログストリームの作成:

    ```bash
    awslocal logs create-log-stream \
      --log-group-name /my-app/logs \
      --log-stream-name test-stream
    ```

4. ログの送信テスト

    ```bash
    # JSON形式のログを送信
    awslocal logs put-log-events \
      --log-group-name /my-app/logs \
      --log-stream-name test-stream \
      --log-events '[{"timestamp":'$(date +%s000)',"message":"{\"level\":\"info\",\"message\":\"テストログ\"}"}]'
    ```

5. ログの確認方法

      - LocalStack Desktop UIを使用する場合:
        - LocalStack Desktopアプリケーションを起動
        - LocalStack Desktopのダッシュボードを開く
        - 左側のメニューから「CloudWatch」を選択
        - 「Logs」セクションでロググループとログストリームを確認

      - AWS CLIを使用する場合:

      ```bash
      awslocal logs get-log-events \
        --log-group-name /my-app/logs \
        --log-stream-name test-stream
      ```

### トラブルシューティング

- LocalStackの状態確認:

```bash
awslocal logs describe-log-groups
```
