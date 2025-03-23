# 1. 週2: バックエンド・フロントエンドのログ基盤実装

## 1.1. 目次

- [1. 週2: バックエンド・フロントエンドのログ基盤実装](#1-週2-バックエンドフロントエンドのログ基盤実装)
  - [1.1. 目次](#11-目次)
  - [1.2. はじめに](#12-はじめに)
    - [1.2.1. 前回の復習](#121-前回の復習)
    - [1.2.2. 今週の学習目標](#122-今週の学習目標)
    - [1.2.3. オブザーバビリティとは何か](#123-オブザーバビリティとは何か)
    - [1.2.4. ログの役割と重要性](#124-ログの役割と重要性)
  - [1.3. LocalStack環境とawslocal・tflocal設定](#13-localstack環境とawslocaltflocal設定)
    - [1.3.1. LocalStackのセットアップと構成](#131-localstackのセットアップと構成)
    - [1.3.2. awslocalのセットアップと基本的なコマンド](#132-awslocalのセットアップと基本的なコマンド)
    - [1.3.3. tflocalの基本設定](#133-tflocalの基本設定)
    - [1.3.4. CloudWatch Logs用のTerraformモジュール作成](#134-cloudwatch-logs用のterraformモジュール作成)
  - [1.4. CloudWatch Logsの基本概念](#14-cloudwatch-logsの基本概念)
    - [1.4.1. CloudWatch Logsの概要とアーキテクチャ](#141-cloudwatch-logsの概要とアーキテクチャ)
      - [1.4.1.1. CloudWatch Logsのアーキテクチャ](#1411-cloudwatch-logsのアーキテクチャ)
    - [1.4.2. ロググループとログストリームの理解](#142-ロググループとログストリームの理解)
      - [1.4.2.1. ロググループ（Log Group）](#1421-ロググループlog-group)
      - [1.4.2.2. ログストリーム（Log Stream）](#1422-ログストリームlog-stream)
      - [1.4.2.3. ロググループとログストリームの実際の選択方法](#1423-ロググループとログストリームの実際の選択方法)
    - [1.4.3. リテンション設定と料金構造](#143-リテンション設定と料金構造)
      - [1.4.3.1. リテンション設定](#1431-リテンション設定)
      - [1.4.3.2. CloudWatch Logsの料金構造](#1432-cloudwatch-logsの料金構造)
    - [1.4.4. アクセス制御とセキュリティ](#144-アクセス制御とセキュリティ)
      - [1.4.4.1. IAMポリシーによるアクセス制御](#1441-iamポリシーによるアクセス制御)
        - [1.4.4.1.1. アプリケーション用のIAMポリシー例](#14411-アプリケーション用のiamポリシー例)
        - [1.4.4.1.2. 開発者用のIAMポリシー例](#14412-開発者用のiamポリシー例)
      - [1.4.4.2. 暗号化とセキュリティのベストプラクティス](#1442-暗号化とセキュリティのベストプラクティス)
    - [1.4.5. LocalStackでのCloudWatch Logsの動作](#145-localstackでのcloudwatch-logsの動作)
      - [1.4.5.1. LocalStackにおけるCloudWatch Logsの基本機能](#1451-localstackにおけるcloudwatch-logsの基本機能)
      - [1.4.5.2. 実際のAWSとLocalStackの相違点](#1452-実際のawsとlocalstackの相違点)
      - [1.4.5.3. LocalStackでCloudWatch Logsを効果的に使用するためのヒント](#1453-localstackでcloudwatch-logsを効果的に使用するためのヒント)
      - [1.4.5.4. LocalStackを使用した実践的な例](#1454-localstackを使用した実践的な例)
  - [1.5. バックエンドの構造化ログ実装](#15-バックエンドの構造化ログ実装)
    - [1.5.1. 構造化ログの概念と利点](#151-構造化ログの概念と利点)
      - [1.5.1.1. 従来のテキストベースログと構造化ログの比較](#1511-従来のテキストベースログと構造化ログの比較)
      - [1.5.1.2. 構造化ログの主な利点](#1512-構造化ログの主な利点)
      - [1.5.1.3. JSONフォーマットを使用する利点](#1513-jsonフォーマットを使用する利点)
    - [1.5.2. Goの標準ロギングライブラリslogの紹介](#152-goの標準ロギングライブラリslogの紹介)
      - [1.5.2.1. slogの基本概念](#1521-slogの基本概念)
      - [1.5.2.2. slogの主な特徴](#1522-slogの主な特徴)
      - [1.5.2.3. slogの基本的な使用例](#1523-slogの基本的な使用例)
    - [1.5.3. slogを使用した構造化ログの設計](#153-slogを使用した構造化ログの設計)
      - [1.5.3.1. ロガーの初期化とカスタマイズ](#1531-ロガーの初期化とカスタマイズ)
      - [1.5.3.2. アプリケーションでのロガーの使用](#1532-アプリケーションでのロガーの使用)
      - [1.5.3.3. 構造化された属性（Attr）の使用](#1533-構造化された属性attrの使用)
    - [1.5.4. ログレベル管理（ERROR/WARN/INFO/DEBUG）の実装](#154-ログレベル管理errorwarninfodebugの実装)
      - [1.5.4.1. slogの標準ログレベル](#1541-slogの標準ログレベル)
      - [1.5.4.2. 環境に応じたログレベルの設定](#1542-環境に応じたログレベルの設定)
        - [1.5.4.2.1. `internal/config/config.go`](#15421-internalconfigconfiggo)
        - [1.5.4.2.2. `.env` ファイルの例](#15422-env-ファイルの例)
      - [1.5.4.3. ログレベルに適したメッセージの例](#1543-ログレベルに適したメッセージの例)
        - [1.5.4.3.1. ERRORレベル](#15431-errorレベル)
        - [1.5.4.3.2. WARNレベル](#15432-warnレベル)
        - [1.5.4.3.3. INFOレベル](#15433-infoレベル)
        - [1.5.4.3.4. DEBUGレベル](#15434-debugレベル)
    - [1.5.5. コンテキスト情報の付与（リクエストID、ユーザーIDなど）](#155-コンテキスト情報の付与リクエストidユーザーidなど)
      - [1.5.5.1. コンテキスト対応のロガー作成](#1551-コンテキスト対応のロガー作成)
      - [1.5.5.2. リクエストIDの生成と追加](#1552-リクエストidの生成と追加)
      - [1.5.5.3. ユーザー認証情報の追加](#1553-ユーザー認証情報の追加)
      - [1.5.5.4. 構造化された例外情報の記録](#1554-構造化された例外情報の記録)
    - [1.5.6. ミドルウェアを使用したリクエスト/レスポンスのログ記録](#156-ミドルウェアを使用したリクエストレスポンスのログ記録)
      - [1.5.6.1. ミドルウェアの設定](#1561-ミドルウェアの設定)
    - [1.5.7. Echo統合とハンドラーへのログ組み込み](#157-echo統合とハンドラーへのログ組み込み)
      - [1.5.7.1. 商品ハンドラーとヘルスチェックハンドラーでのログ使用例](#1571-商品ハンドラーとヘルスチェックハンドラーでのログ使用例)
      - [1.5.7.2. エラーハンドリングとログ記録](#1572-エラーハンドリングとログ記録)
      - [1.5.7.3. ロガーのパフォーマンスに関する考慮事項](#1573-ロガーのパフォーマンスに関する考慮事項)
    - [1.5.8. 構造化ログ実装の動作確認](#158-構造化ログ実装の動作確認)
      - [1.5.8.1. アプリケーションの起動と基本ログの確認](#1581-アプリケーションの起動と基本ログの確認)
      - [1.5.8.2. APIリクエストによるログの生成と確認](#1582-apiリクエストによるログの生成と確認)
        - [1.5.8.2.1. ヘルスチェックAPIの動作確認](#15821-ヘルスチェックapiの動作確認)
        - [1.5.8.2.2. 商品一覧APIの動作確認](#15822-商品一覧apiの動作確認)
      - [1.5.8.3. 異なるログレベルの確認](#1583-異なるログレベルの確認)
      - [1.5.8.4. エラーシナリオのテスト](#1584-エラーシナリオのテスト)
      - [1.5.8.5. リクエストとレスポンスの本文ログの確認](#1585-リクエストとレスポンスの本文ログの確認)
      - [1.5.8.6. トラブルシューティング](#1586-トラブルシューティング)
        - [1.5.8.6.1. ログが出力されない](#15861-ログが出力されない)
        - [1.5.8.6.2. JSON以外の形式でログが出力される](#15862-json以外の形式でログが出力される)
        - [1.5.8.6.3. リクエストIDが含まれない](#15863-リクエストidが含まれない)
        - [1.5.8.6.4. エラー情報が不十分](#15864-エラー情報が不十分)
      - [1.5.8.7. jqを使ったログ分析](#1587-jqを使ったログ分析)
      - [1.5.8.8. 本番環境を想定した設定のテスト](#1588-本番環境を想定した設定のテスト)
      - [1.5.8.9. まとめ](#1589-まとめ)
  - [1.6. バックエンドログのCloudWatch Logs連携](#16-バックエンドログのcloudwatch-logs連携)
    - [1.6.1. AWS SDK for Go v2の設定](#161-aws-sdk-for-go-v2の設定)
      - [1.6.1.1. AWS SDK for Go v2の主な特徴](#1611-aws-sdk-for-go-v2の主な特徴)
      - [1.6.1.2. インストール手順](#1612-インストール手順)
      - [1.6.1.3. SDK設定の基本](#1613-sdk設定の基本)
    - [1.6.2. CloudWatch Logs用のslogハンドラーの実装](#162-cloudwatch-logs用のslogハンドラーの実装)
      - [1.6.2.1. CloudWatch Logs用のslogハンドラー実装](#1621-cloudwatch-logs用のslogハンドラー実装)
      - [1.6.2.2. ロガーの初期化処理実装](#1622-ロガーの初期化処理実装)
    - [1.6.3. バッチ処理とエラーハンドリング](#163-バッチ処理とエラーハンドリング)
      - [1.6.3.1. バッチ処理の仕組み](#1631-バッチ処理の仕組み)
      - [1.6.3.2. エラーハンドリングの方法](#1632-エラーハンドリングの方法)
    - [1.6.4. LocalStackへのログ転送設定と検証](#164-localstackへのログ転送設定と検証)
      - [1.6.4.1. メインアプリケーションへのLogger統合](#1641-メインアプリケーションへのlogger統合)
      - [1.6.4.2. 環境変数設定](#1642-環境変数設定)
      - [1.6.4.3. ログ転送の検証方法](#1643-ログ転送の検証方法)
    - [1.6.5. 非同期ログ転送と性能最適化](#165-非同期ログ転送と性能最適化)
      - [1.6.5.1. 非同期処理の重要性](#1651-非同期処理の重要性)
      - [1.6.5.2. さらなる最適化戦略](#1652-さらなる最適化戦略)
        - [1.6.5.2.1. ロググループとログストリームの効率的な管理](#16521-ロググループとログストリームの効率的な管理)
        - [1.6.5.2.2. ログの重要度に基づいたサンプリング](#16522-ログの重要度に基づいたサンプリング)
        - [1.6.5.2.3. ネットワーク再試行とバックオフ戦略](#16523-ネットワーク再試行とバックオフ戦略)
      - [1.6.5.3. パフォーマンスベンチマークの実施](#1653-パフォーマンスベンチマークの実施)
      - [1.6.5.4. ロギング実装のテスト](#1654-ロギング実装のテスト)

## 1.2. はじめに

### 1.2.1. 前回の復習

前週では、ECサイトアプリケーションの基本構造とアーキテクチャの設計について学びました。具体的には以下の内容を扱いました：

- マイクロサービスアーキテクチャの基本概念と設計原則
- GoとEchoフレームワークを使用したバックエンドAPIの基本実装
- Next.jsを用いたフロントエンド（顧客向けサイトと管理者向けサイト）の基本構造
- Docker Composeを使用した開発環境のセットアップと構成

これらの知識をベースに、今週はアプリケーションの「可観測性（オブザーバビリティ）」に焦点を当て、ログ基盤の実装に取り組みます。ソフトウェアの開発と運用において、適切なログ記録は問題の早期発見、デバッグ、パフォーマンス分析に不可欠です。

### 1.2.2. 今週の学習目標

今週の講義を通じて、以下の学習目標を達成することを目指します：

1. **オブザーバビリティの基本概念を理解する**
   - ログ、メトリクス、トレースの違いと役割を理解する
   - 効果的なログ戦略の設計方法を学ぶ

2. **LocalStackを使用したAWSサービスのローカル環境構築方法を習得する**
   - LocalStackの基本的なセットアップと使用方法
   - AWS CLIとTerraformを用いたインフラ構成の管理

3. **CloudWatch Logsの基本概念と使用方法を理解する**
   - ロググループとログストリームの概念
   - CloudWatch Logsの基本的な設定とアクセス制御

4. **Goのslogを使用した構造化ログの実装方法を習得する**
   - 構造化ログの利点と実装パターン
   - ログレベル管理とコンテキスト情報の付与方法

5. **Next.jsアプリケーションでのログ実装と収集の方法を学ぶ**
   - サーバーサイドとクライアントサイドのログ戦略
   - フロントエンドエラーの効果的な捕捉と記録

6. **ログデータの分析と活用方法を理解する**
   - CloudWatch Logs Insightsの基本的な使用方法
   - ログに基づくアラートとモニタリングの設定

これらの目標を達成することで、本番環境での運用に耐えうる堅牢なログ基盤の設計・実装能力を身につけることができます。

### 1.2.3. オブザーバビリティとは何か

オブザーバビリティ（Observability、可観測性）は、システムの内部状態を外部から観測・理解する能力を指します。この概念は制御理論から派生し、現代のクラウドネイティブな分散システムにおいて重要な役割を果たしています。

**オブザーバビリティの3つの柱**:

1. **ログ（Logs）**: システム内で発生したイベントの時系列記録
2. **メトリクス（Metrics）**: システムの状態を数値で表した時系列データ
3. **トレース（Traces）**: 分散システム内でのリクエストの流れを追跡するデータ

従来の「モニタリング」が既知の問題を検出することに重点を置いていたのに対し、オブザーバビリティは未知の問題を含めたシステムの振る舞いを理解することに重点を置いています。

**オブザーバビリティが重要な理由**:

- **複雑性の増大**: マイクロサービスアーキテクチャの採用により、システムの複雑性は飛躍的に高まっています
- **分散システムの課題**: 複数のサービスにまたがる問題の特定と解決が困難になっています
- **迅速な問題解決**: ユーザー体験に影響する問題を素早く特定し解決する必要があります
- **予防的対応**: 問題が大きくなる前に早期に検出し対処することが可能になります

今週の講義では、オブザーバビリティの柱の一つである「ログ」に焦点を当て、次週以降でメトリクスとトレースについて学んでいきます。

### 1.2.4. ログの役割と重要性

ログは、アプリケーションの動作状況を記録した時系列データであり、システムで発生したイベントに関する情報を提供します。適切に設計・実装されたログシステムは、開発・運用の様々な局面で重要な役割を果たします。

**ログの主な役割**:

1. **デバッグと問題解決**
   - エラーの詳細な情報とコンテキストを提供
   - 問題発生時の状況を再現するための情報を記録
   - 根本原因分析（RCA）のための証拠を提供

2. **セキュリティとコンプライアンス**
   - 不審なアクティビティの検出と分析
   - セキュリティ監査のための証跡の提供
   - 規制要件を満たすための活動記録

3. **ユーザー行動とビジネスインサイト**
   - ユーザーの行動パターンの分析
   - 機能の利用状況と有効性の評価
   - ビジネス指標の導出と意思決定支援

4. **システムのパフォーマンス分析**
   - ボトルネックの特定とパフォーマンス最適化
   - リソース使用状況の監視と容量計画
   - スケーリングの必要性の予測

**効果的なログに求められる特性**:

- **構造化**: 機械処理が容易なJSON等の形式
- **コンテキスト豊富**: リクエストID、ユーザーID等の関連情報を含む
- **適切な詳細度**: 必要十分な情報を含み、ノイズを最小限に
- **一貫性**: サービス間で統一されたフォーマットと命名規則
- **タイムスタンプの正確性**: 正確な時刻情報の記録
- **分類可能**: ログレベル（ERROR/WARN/INFO/DEBUG）による重要度の区別

本講義では、これらの特性を持つ効果的なログシステムの構築方法を学びます。特にAWSのCloudWatch Logsを活用し、フロントエンドとバックエンドの両方でのログ実装を行います。また、LocalStackを使用してローカル開発環境でもクラウドサービスと同様の機能を利用できるようにします。

次のセクションでは、LocalStack環境とAWS CLI設定について学んでいきましょう。

## 1.3. LocalStack環境とawslocal・tflocal設定

クラウドサービスを活用した開発において、本番環境と同等の機能をローカル開発環境で利用できることは非常に重要です。LocalStackはAWSサービスをローカル環境でエミュレートするツールであり、開発とテストのコストと時間を大幅に削減することができます。このセクションでは、LocalStackのセットアップと設定、LocalStack専用のラッパーツールである「awslocal」と「tflocal」の基本的な使用方法について学んでいきます。

### 1.3.1. LocalStackのセットアップと構成

LocalStackは、AWSサービスのローカルエミュレーションを提供するコンテナベースのツールです。これを使用することで、実際のAWSアカウントを使わずに開発やテストを行うことができます。

**LocalStackの主な利点**:

- 本番環境に近い環境をローカルで再現できる
- AWSの利用料金が発生しない
- インターネット接続なしで開発が可能
- 開発サイクルの高速化（デプロイ時間の短縮）
- テスト自動化の容易さ

今回のプロジェクトでは、Docker Composeを使用してLocalStackをセットアップします。以下に、`docker-compose.yml`ファイルへの追加内容を示します：

```yaml
services:
  # 既存のサービス定義...

  localstack:
    image: localstack/localstack:latest
    container_name: localstack
    ports:
      - "4566:4566"            # LocalStackのメインエンドポイント
      - "4510-4559:4510-4559"  # 内部サービス用ポート範囲
    environment:
      - DEBUG=1
      - DOCKER_HOST=unix:///var/run/docker.sock
      - HOSTNAME_EXTERNAL=localstack
      - SERVICES=cloudwatch,logs,s3
      - DEFAULT_REGION=ap-northeast-1
      - AWS_DEFAULT_REGION=ap-northeast-1
      - AWS_ACCESS_KEY_ID=localstack
      - AWS_SECRET_ACCESS_KEY=localstack
    volumes:
      - "${LOCALSTACK_VOLUME_DIR:-./localstack-data}:/var/lib/localstack"
      - "/var/run/docker.sock:/var/run/docker.sock"
    networks:
      - aws-observability-network
```

この設定では、CloudWatch LogsとS3のエミュレーションを有効にしています。`ap-northeast-1`をデフォルトリージョンとして設定し、アクセスキーとシークレットキーも`localstack`という簡易的な値で設定しています。

**LocalStackの起動と確認**:

Docker Composeを使用してLocalStackを起動します：

```bash
docker-compose up -d localstack
```

LocalStackの状態を確認するには、次のコマンドを実行します：

```bash
docker logs localstack
```

正常に起動していれば、以下のような出力が表示されます：

```text
...
2023-XX-XX:XX:XX:XX INFO --- Ready.
...
```

### 1.3.2. awslocalのセットアップと基本的なコマンド

通常、AWSサービスを操作するにはAWS CLIを使用しますが、LocalStackでは専用のラッパーツールである「awslocal」を使用することで、エンドポイントURLなどのパラメータを毎回指定する手間を省くことができます。awslocalは、内部的にはAWS CLIを使用しながら、LocalStackに適したパラメータを自動的に設定してくれるツールです。

**Pythonとpipのインストール**:

awslocalはPythonパッケージとして提供されているため、まずPythonとpipがインストールされているか確認します。インストールされていない場合は、以下のコマンドでインストールします：

```bash
# Ubuntu/Debianの場合
sudo apt update
sudo apt install -y python3 python3-pip

# CentOS/RHEL/Fedoraの場合
sudo dnf install -y python3 python3-pip

# macOSの場合（Homebrewを使用）
brew install python

# pipxを使用する
# まずpipxをインストール
brew install pipx

# Windowsの場合
# python.orgからインストーラーをダウンロードしてインストール
# または、Chocolateyを使用：
# choco install -y python
```

Pythonとpipがインストールされていることを確認します：

```bash
python3 --version
pip3 --version
```

**awslocalのインストール**:

pipを使ってawslocalをインストールします：

```bash
pip3 install awscli-local
# pipxを使ってawscli-localをインストール
pipx install awscli-local

```

**注意**: awslocalを使用するには、AWS CLIがインストールされている必要があります。まだインストールしていない場合は、以下のコマンドでインストールします：

```bash
# Ubuntu/Debianの場合
pip3 install awscli

# macOSの場合
brew install awscli

# Windowsの場合（Chocolateyを使用）
choco install awscli
```

**awslocalの基本的な使用方法**:

awslocalを使用すると、AWS CLIのコマンドをLocalStack環境に対して簡単に実行できます。基本的なコマンド構文はAWS CLIと同じですが、`aws`の代わりに`awslocal`を使用します。

```bash
# CloudWatch Logsのロググループを一覧表示
awslocal logs describe-log-groups

# ロググループを作成
awslocal logs create-log-group --log-group-name /aws-observability-ecommerce/backend

# 作成したロググループを確認
awslocal logs describe-log-groups

# ログストリームを作成
awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/backend --log-stream-name api-server

# S3バケットを作成
awslocal s3 mb s3://aws-observability-ecommerce-logs

# S3バケットを一覧表示
awslocal s3 ls
```

awslocalは内部的に`--endpoint-url=http://localhost:4566`パラメータを自動的に追加してLocalStackに接続するため、コマンドがシンプルになります。

### 1.3.3. tflocalの基本設定

Terraformを使ってLocalStackのリソースを管理する場合も、専用のラッパーツール「tflocal」を使用すると便利です。tflocalは、Terraformコマンドを実行する際にLocalStackのエンドポイントに自動的に接続するツールです。

**tflocalのインストール**:

tflocalもPythonパッケージとして提供されています：

```bash
pip3 install terraform-local

# pipxを使ってterraform-localをインストール
pipx install terraform-local

```

**注意**: tflocalを使用するには、Terraformがインストールされている必要があります。まだインストールしていない場合は、以下のコマンドでインストールします：

```bash
# Ubuntu/Debianの場合
wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform

# macOSの場合（Homebrewを使用）
brew install terraform

# Windowsの場合（Chocolateyを使用）
choco install terraform
```

**プロジェクト構造の作成**:

以下のディレクトリ構造を作成します。必要なディレクトリとファイルを作成するコマンドも記載します：

```bash
# プロジェクトのルートディレクトリから実行
mkdir -p infra/terraform/modules/cloudwatch

# メインのTerraformファイルを作成
touch infra/terraform/{main.tf,variables.tf,outputs.tf,providers.tf}

# CloudWatchモジュール用のファイルを作成
touch infra/terraform/modules/cloudwatch/{main.tf,variables.tf,outputs.tf}
```

作成されるディレクトリ構造は以下のようになります：

```text
infra/terraform/
├── main.tf        # メインの設定ファイル
├── variables.tf   # 変数定義
├── outputs.tf     # 出力定義
├── providers.tf   # プロバイダー設定
└── modules/       # 再利用可能なモジュール
    └── cloudwatch/ # CloudWatch関連のモジュール
        ├── main.tf
        ├── variables.tf
        └── outputs.tf
```

**providers.tfの設定**:

LocalStackと連携するためのAWSプロバイダー設定を行います。tflocalを使用する場合でも、通常のTerraformファイルと同じ内容を記述します：

```hcl
provider "aws" {
  region                      = "ap-northeast-1"
  access_key                  = "localstack"
  secret_key                  = "localstack"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  # tflocalを使用する場合、以下の設定は自動的に行われるため、コメントアウトしておきます
  # endpoints {
  #   cloudwatch     = "http://localhost:4566"
  #   logs           = "http://localhost:4566"
  #   s3             = "http://localhost:4566"
  # }
}
```

**tflocalの基本的なコマンド**:

tflocalを使用すると、通常のTerraformコマンドをLocalStack環境に対して実行できます。基本的なコマンド構文はTerraformと同じですが、`terraform`の代わりに`tflocal`を使用します：

```bash
# ディレクトリに移動
cd infra/terraform

# 初期化
tflocal init

# 計画の作成
tflocal plan

# 適用
tflocal apply

# 状態の確認
tflocal state list

# 破棄
tflocal destroy
```

tflocalは内部的にエンドポイントURLなどのパラメータを自動的に設定してくれるため、LocalStack環境へのデプロイが簡単になります。

### 1.3.4. CloudWatch Logs用のTerraformモジュール作成

次に、CloudWatch Logsリソースを管理するためのTerraformモジュールを作成します。このモジュールでは、アプリケーションで使用するロググループとログストリームを定義します。

**modules/cloudwatch/variables.tf**:

```hcl
variable "app_name" {
  description = "アプリケーション名"
  type        = string
  default     = "aws-observability-ecommerce"
}

variable "env" {
  description = "環境名（dev, staging, prod等）"
  type        = string
  default     = "dev"
}

variable "log_groups" {
  description = "作成するロググループの設定"
  type = list(object({
    name       = string
    retention  = number
    streams    = list(string)
  }))
  default = []
}
```

**modules/cloudwatch/main.tf**:

```hcl
resource "aws_cloudwatch_log_group" "this" {
  for_each = { for lg in var.log_groups : lg.name => lg }

  name              = "/${var.app_name}/${var.env}/${each.value.name}"
  retention_in_days = each.value.retention

  tags = {
    Environment = var.env
    Application = var.app_name
  }
}

resource "aws_cloudwatch_log_stream" "this" {
  for_each = {
    for pair in flatten([
      for lg in var.log_groups : [
        for stream in lg.streams : {
          log_group = lg.name
          stream    = stream
        }
      ]
    ]) : "${pair.log_group}-${pair.stream}" => pair
  }

  name           = each.value.stream
  log_group_name = "/${var.app_name}/${var.env}/${each.value.log_group}"

  depends_on = [aws_cloudwatch_log_group.this]
}
```

**modules/cloudwatch/outputs.tf**:

```hcl
output "log_group_arns" {
  description = "作成されたロググループのARN"
  value       = { for k, v in aws_cloudwatch_log_group.this : k => v.arn }
}

output "log_group_names" {
  description = "作成されたロググループの名前"
  value       = { for k, v in aws_cloudwatch_log_group.this : k => v.name }
}
```

**メインのTerraformファイル（infra/terraform/main.tf）**:

```hcl
module "cloudwatch_logs" {
  source = "./modules/cloudwatch"

  app_name = "aws-observability-ecommerce"
  env      = "dev"

  log_groups = [
    {
      name      = "backend"
      retention = 30  # 30日間保持
      streams   = ["api-server", "error"]
    },
    {
      name      = "frontend-customer"
      retention = 14  # 14日間保持
      streams   = ["app", "error", "access"]
    },
    {
      name      = "frontend-admin"
      retention = 14  # 14日間保持
      streams   = ["app", "error", "access"]
    }
  ]
}
```

**outputs.tf**:

```hcl
output "cloudwatch_log_groups" {
  value = module.cloudwatch_logs.log_group_names
}
```

**variables.tf**:

```hcl
# グローバル変数定義（必要に応じて追加）
```

**Terraformモジュールの適用**:

作成したTerraformモジュールをLocalStack環境に適用するには、以下のコマンドを実行します：

```bash
cd infra/terraform
tflocal init
tflocal plan
tflocal apply -auto-approve
```

適用が成功すると、指定したロググループとログストリームがLocalStackに作成されます。これを確認するには、awslocalコマンドを使用します：

```bash
# ロググループの一覧を表示
awslocal logs describe-log-groups

# 特定のロググループのログストリームを表示
awslocal logs describe-log-streams --log-group-name /aws-observability-ecommerce/dev/backend
```

これでLocalStack環境にCloudWatch Logsのリソースが作成され、アプリケーションからログを送信する準備が整いました。次のセクションでは、CloudWatch Logsの基本概念について学んでいきます。

## 1.4. CloudWatch Logsの基本概念

CloudWatch Logsは、AWSのマネージドログサービスであり、アプリケーションやAWSリソースからログを収集、監視、保存、分析するための機能を提供します。このセクションでは、CloudWatch Logsの基本的な概念や機能、そして実際にどのように活用できるかを学びます。

### 1.4.1. CloudWatch Logsの概要とアーキテクチャ

CloudWatch Logsは、分散システムにおけるログ管理の複雑さを解決するために設計されたマネージドサービスです。以下にその主な特徴を示します：

- **一元管理**: 複数のサービスやインスタンスからのログを一箇所に集約
- **リアルタイム監視**: ログデータのリアルタイムモニタリングとアラート機能
- **長期保存**: ログの長期的な保存と保持ポリシーの設定
- **検索と分析**: 強力な検索機能とログ分析ツール（CloudWatch Logs Insights）
- **セキュリティ**: IAMと統合されたアクセス制御
- **統合**: 他のAWSサービスとの緊密な連携

#### 1.4.1.1. CloudWatch Logsのアーキテクチャ

CloudWatch Logsのアーキテクチャは以下のコンポーネントで構成されています：

1. **ログイベント**: 最小単位のログデータ。タイムスタンプとメッセージを持ちます。
2. **ログストリーム**: 同一のソースからのログイベントのシーケンス。例えば特定のアプリケーションインスタンスのログなど。
3. **ロググループ**: 関連するログストリームの集合。通常、同じタイプのログを含みます。
4. **メトリクスフィルター**: ログデータからメトリクスを抽出するためのフィルター。
5. **サブスクリプションフィルター**: ログデータを他のサービス（Lambda、Kinesisなど）にリアルタイムに送信するためのフィルター。

以下は、CloudWatch Logsのアーキテクチャを表す簡易的な図です：

```text
アプリケーション/AWS サービス
        ↓
    ログイベント
        ↓
    ログストリーム
        ↓
    ロググループ
        ↓
メトリクスフィルター / サブスクリプションフィルター
        ↓
アラーム / 分析 / 外部サービス
```

### 1.4.2. ロググループとログストリームの理解

CloudWatch Logsの階層構造を理解することは、効果的なログ管理の基本です。

#### 1.4.2.1. ロググループ（Log Group）

ロググループは、CloudWatch Logsの最上位の組織単位で、以下の特徴があります：

- **名前空間**: 通常、アプリケーション名やサービス名に基づいて命名します
- **設定の共有**: リテンション設定、アクセス制御などがロググループ単位で適用されます
- **命名規則**: AWS公式ドキュメントでは `/[application]/[environment]/[component]` のような階層的な命名が推奨されています
  - 例: `/aws-observability-ecommerce/dev/backend`

ロググループの作成と管理：

```bash
# AWS CLIでロググループを作成する
aws logs create-log-group --log-group-name /aws-observability-ecommerce/dev/backend

# ロググループのリテンション期間を設定（30日間）
aws logs put-retention-policy --log-group-name /aws-observability-ecommerce/dev/backend --retention-in-days 30

# ロググループの一覧を表示
aws logs describe-log-groups

# ロググループを削除
aws logs delete-log-group --log-group-name /aws-observability-ecommerce/dev/backend
```

LocalStackでは、`aws` の代わりに `awslocal` コマンドを使用します：

```bash
# LocalStackでロググループを作成する
awslocal logs create-log-group --log-group-name /aws-observability-ecommerce/dev/backend

# LocalStackでロググループのリテンション期間を設定（30日間）
awslocal logs put-retention-policy --log-group-name /aws-observability-ecommerce/dev/backend --retention-in-days 30
```

#### 1.4.2.2. ログストリーム（Log Stream）

ログストリームは、同じソースから発生するログイベントのシーケンスです：

- **単一ソース**: 通常、特定のインスタンス、コンテナ、プロセスからのログを表します
- **時系列データ**: ログイベントはタイムスタンプ順に保存されます
- **命名規則**: インスタンスID、コンテナID、またはコンポーネント名などが一般的です
  - 例: `api-server-instance-1`, `payment-processor`, `error`

ログストリームの作成と管理：

```bash
# ログストリームを作成
aws logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/backend --log-stream-name api-server

# 特定のロググループのログストリーム一覧を表示
aws logs describe-log-streams --log-group-name /aws-observability-ecommerce/dev/backend

# ログストリームの内容を表示
aws logs get-log-events --log-group-name /aws-observability-ecommerce/dev/backend --log-stream-name api-server
```

LocalStackでの同等のコマンド：

```bash
# LocalStackでログストリームを作成
awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/backend --log-stream-name api-server

# LocalStackでログストリームの一覧を表示
awslocal logs describe-log-streams --log-group-name /aws-observability-ecommerce/dev/backend
```

#### 1.4.2.3. ロググループとログストリームの実際の選択方法

プロジェクトでロググループとログストリームをどのように組織化するかは、アプリケーションのアーキテクチャ、規模、監視ニーズに依存します。一般的なパターンをいくつか紹介します：

1. **サービスベース**: 各マイクロサービスに1つのロググループを割り当て、インスタンスごとにログストリームを作成

   ```text
   /aws-observability-ecommerce/dev/backend
     - instance-1
     - instance-2
   /aws-observability-ecommerce/dev/frontend-customer
     - instance-1
     - instance-2
   ```

2. **機能ベース**: 機能やコンポーネントごとにロググループを分け、レベルごとにログストリームを分ける

   ```text
   /aws-observability-ecommerce/dev/api
     - info
     - error
     - debug
   /aws-observability-ecommerce/dev/database
     - info
     - error
     - slow-queries
   ```

3. **環境ベース**: 環境（開発、ステージング、本番）ごとにロググループを分け、サービスごとにログストリームを分ける

   ```text
   /aws-observability-ecommerce/dev
     - backend
     - frontend-customer
     - frontend-admin
   /aws-observability-ecommerce/prod
     - backend
     - frontend-customer
     - frontend-admin
   ```

私たちのプロジェクトでは、以下のような構成を採用します：

```text
/aws-observability-ecommerce/dev/backend
  - api-server
  - error
/aws-observability-ecommerce/dev/frontend-customer
  - app
  - error
  - access
/aws-observability-ecommerce/dev/frontend-admin
  - app
  - error
  - access
```

この構成は、サービスごとにロググループを分け、ログの種類ごとにログストリームを分けるハイブリッドなアプローチです。これにより、サービスごとの監視と問題の種類ごとの分析が容易になります。

### 1.4.3. リテンション設定と料金構造

CloudWatch Logsはログデータを長期間保存できますが、保存期間とコストのバランスを取る必要があります。

#### 1.4.3.1. リテンション設定

ロググループごとにリテンション期間（保持期間）を設定できます：

- **デフォルト**: リテンション期間を設定しない場合、ログは無期限に保持されます
- **カスタム期間**: 1日、3日、5日、1週間、2週間、1ヶ月、2ヶ月、3ヶ月、4ヶ月、5ヶ月、6ヶ月、1年、13ヶ月、18ヶ月、2年、5年、10年、または無期限から選択できます
- **自動削除**: 設定したリテンション期間を過ぎたログデータは自動的に削除されます

リテンション設定の適用：

```bash
# リテンション期間を30日に設定
aws logs put-retention-policy --log-group-name /aws-observability-ecommerce/dev/backend --retention-in-days 30

# リテンション期間を確認
aws logs describe-log-groups --log-group-name-prefix /aws-observability-ecommerce/dev/backend

# リテンション設定を削除（無期限に戻す）
aws logs delete-retention-policy --log-group-name /aws-observability-ecommerce/dev/backend
```

LocalStackでの設定：

```bash
# LocalStackでリテンション期間を設定
awslocal logs put-retention-policy --log-group-name /aws-observability-ecommerce/dev/backend --retention-in-days 30
```

#### 1.4.3.2. CloudWatch Logsの料金構造

CloudWatch Logsの料金は主に以下の要素から構成されています：

1. **取り込み (Ingestion)**: ログデータの収集にかかるコスト
2. **保存 (Storage)**: ログデータの保存にかかるコスト
3. **分析 (Analysis)**: CloudWatch Logs Insightsを使用したログ分析にかかるコスト

一般的な料金（2023年時点のアジアパシフィック（東京）リージョンの例）：

- 取り込み: 5.42 USD / GB
- 保存: 0.0342 USD / GB-月
- 分析: 0.00619 USD / GB（スキャンしたデータ量）

コスト最適化のためのベストプラクティス：

1. **適切なリテンション期間の設定**
   - 必要な期間のみログを保持し、古いログは自動削除
   - 重要度の低いログはより短い期間に設定

2. **ログの選別**
   - すべてのログを送信するのではなく、必要なログのみを選択
   - デバッグログは開発環境のみに制限

3. **ログの圧縮**
   - データ量を減らすためにログを圧縮してから送信
   - 特に繰り返しの多いログパターン

4. **CloudWatch Logs Insightsのクエリ最適化**
   - 必要なデータのみをスキャンするようにクエリを最適化
   - 必要に応じてログを前処理して検索を効率化

LocalStackを使用する開発環境では料金は発生しませんが、本番環境に移行する前にこれらのコスト要因を考慮することが重要です。

### 1.4.4. アクセス制御とセキュリティ

CloudWatch Logsへのアクセスと操作は、AWS Identity and Access Management (IAM) によって制御されます。適切なアクセス制御を設定することで、ログデータのセキュリティを確保できます。

#### 1.4.4.1. IAMポリシーによるアクセス制御

CloudWatch Logsに関連する主なIAMアクションは以下の通りです：

- **logs:CreateLogGroup**: ロググループの作成
- **logs:CreateLogStream**: ログストリームの作成
- **logs:PutLogEvents**: ログイベントの書き込み
- **logs:DescribeLogGroups**: ロググループの一覧表示
- **logs:DescribeLogStreams**: ログストリームの一覧表示
- **logs:GetLogEvents**: ログイベントの読み取り
- **logs:DeleteLogGroup**: ロググループの削除
- **logs:DeleteLogStream**: ログストリームの削除

##### 1.4.4.1.1. アプリケーション用のIAMポリシー例

以下は、アプリケーションがログを書き込むための最小権限ポリシーの例です：

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:ap-northeast-1:123456789012:log-group:/aws-observability-ecommerce/dev/*:*"
    }
  ]
}
```

##### 1.4.4.1.2. 開発者用のIAMポリシー例

開発者向けのより広範なアクセス権限を持つポリシー例：

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:DescribeLogGroups",
        "logs:DescribeLogStreams",
        "logs:PutLogEvents",
        "logs:GetLogEvents",
        "logs:FilterLogEvents"
      ],
      "Resource": "arn:aws:logs:ap-northeast-1:123456789012:log-group:/aws-observability-ecommerce/dev/*:*"
    }
  ]
}
```

#### 1.4.4.2. 暗号化とセキュリティのベストプラクティス

CloudWatch Logsでのセキュリティを強化するためのベストプラクティス：

1. **保管時の暗号化**
   - AWS KMSを使用してログデータを暗号化

   ```bash
   # KMSキーを使用してロググループを暗号化
   aws logs create-log-group --log-group-name /aws-observability-ecommerce/prod/backend --kms-key-id arn:aws:kms:ap-northeast-1:123456789012:key/abcd1234-1234-1234-1234-123456abcdef
   ```

2. **最小権限の原則**
   - 必要な権限のみを付与する
   - ロググループごとにきめ細かなアクセス制御を設定

3. **センシティブデータの処理**
   - ログにパスワードやクレジットカード番号などの機密情報を含めない
   - 必要に応じてマスキングや匿名化を実施

4. **リソースポリシー**
   - 特定のAWSアカウントやサービスからのアクセスを制限

   ```bash
   # ロググループにリソースポリシーを設定
   aws logs put-resource-policy --policy-name AllowCloudTrail --policy-document '{"Version":"2012-10-17","Statement":[{"Sid":"","Effect":"Allow","Principal":{"Service":"cloudtrail.amazonaws.com"},"Action":"logs:PutLogEvents","Resource":"arn:aws:logs:ap-northeast-1:123456789012:log-group:/aws-observability-ecommerce/prod/cloudtrail:*"}]}'
   ```

5. **VPCエンドポイント**
   - プライベートVPC内からCloudWatch Logsへの安全なアクセスを確保

   ```bash
   # CloudWatch Logs用のVPCエンドポイントを作成
   aws ec2 create-vpc-endpoint --vpc-id vpc-12345678 --service-name com.amazonaws.ap-northeast-1.logs --route-table-ids rtb-12345678
   ```

LocalStack環境でこれらの機能を使用する場合、実装は簡略化される場合があります。本番環境では、これらのセキュリティベストプラクティスを適切に適用することが重要です。

### 1.4.5. LocalStackでのCloudWatch Logsの動作

LocalStackは、AWSサービスをローカル環境でエミュレートするための強力なツールですが、すべての機能が実際のAWS環境と完全に一致するわけではありません。ここでは、LocalStackでのCloudWatch Logsの動作特性とその違いについて説明します。

#### 1.4.5.1. LocalStackにおけるCloudWatch Logsの基本機能

LocalStackでサポートされている主なCloudWatch Logs機能：

- ロググループの作成と管理
- ログストリームの作成と管理
- ログイベントの書き込みと読み取り
- 基本的なログクエリ機能

#### 1.4.5.2. 実際のAWSとLocalStackの相違点

1. **パフォーマンスの違い**
   - LocalStackはローカル環境で動作するため、実際のAWSより高速
   - リソース制約により、大量のログ処理は実際のAWSより低速になる場合あり

2. **機能の制限**
   - LocalStackはすべてのCloudWatch Logs機能を完全にサポートしているわけではない
   - 特に高度な機能（メトリクスフィルター、サブスクリプションフィルターなど）は制限あり
   - CloudWatch Logs Insightsのサポート範囲も限定的

3. **永続性の違い**
   - 特別な設定をしない限り、LocalStack再起動時にデータは失われる
   - ボリュームマウントを使用して永続化することが可能

4. **認証と認可**
   - LocalStackでは認証情報は簡易的（"localstack"/"localstack"など）
   - IAMポリシーの完全な適用はされない場合がある

#### 1.4.5.3. LocalStackでCloudWatch Logsを効果的に使用するためのヒント

1. **データの永続化**

   ```yaml
   # docker-compose.ymlでの永続化設定
   services:
     localstack:
       # ...他の設定...
       volumes:
         - "${LOCALSTACK_VOLUME_DIR:-./localstack-data}:/var/lib/localstack"
         - "/var/run/docker.sock:/var/run/docker.sock"
   ```

2. **初期化スクリプトの活用**
   - コンテナ起動時に自動的にリソースを作成するスクリプトを用意

   ```bash
   # infra/localstack/init.sh（コンテナ起動時に実行）
   #!/bin/bash

   # ロググループの作成
   awslocal logs create-log-group --log-group-name /aws-observability-ecommerce/dev/backend
   awslocal logs create-log-group --log-group-name /aws-observability-ecommerce/dev/frontend-customer
   awslocal logs create-log-group --log-group-name /aws-observability-ecommerce/dev/frontend-admin

   # ログストリームの作成
   awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/backend --log-stream-name api-server
   awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/backend --log-stream-name error

   # フロントエンド顧客用
   awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/frontend-customer --log-stream-name app
   awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/frontend-customer --log-stream-name error
   awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/frontend-customer --log-stream-name access

   # フロントエンド管理者用
   awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/frontend-admin --log-stream-name app
   awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/frontend-admin --log-stream-name error
   awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/frontend-admin --log-stream-name access
   ```

3. **リソース確認用のユーティリティ関数**
   - `task` コマンドを使った便利なショートカットの作成

   ```yaml
   # Taskfile.yml
   version: '3'

   tasks:
     logs:check:
       desc: CloudWatch Logsのロググループを一覧表示
       cmds:
         - awslocal logs describe-log-groups

     logs:streams:
       desc: 特定のロググループのログストリームを一覧表示
       cmds:
         - awslocal logs describe-log-streams --log-group-name /aws-observability-ecommerce/dev/{{.CLI_ARGS}}
   ```

4. **テストデータの投入**
   - テスト用のログデータを作成する関数

   ```bash
   # テスト用ログデータの投入
   function put_test_logs() {
     local log_group=$1
     local log_stream=$2
     local message=$3
     local timestamp=$(date +%s)000

     awslocal logs put-log-events \
       --log-group-name "$log_group" \
       --log-stream-name "$log_stream" \
       --log-events timestamp=$timestamp,message="$message"
   }

   # 使用例
   put_test_logs "/aws-observability-ecommerce/dev/backend" "api-server" '{"level":"info","message":"Test log message","requestId":"12345"}'
   ```

5. **開発環境と本番環境の違いを考慮**
   - 設定ファイルを環境ごとに分離
   - 環境変数を使った条件分岐

   ```go
   // Go言語での環境ごとの条件分岐例
   var logClient *cloudwatchlogs.Client

   if os.Getenv("APP_ENV") == "local" {
       // LocalStack用の設定
       customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
           return aws.Endpoint{
               URL:           "http://localstack:4566",
               SigningRegion: region,
           }, nil
       })

       cfg, err := config.LoadDefaultConfig(context.Background(),
           config.WithRegion("ap-northeast-1"),
           config.WithEndpointResolverWithOptions(customResolver),
           config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("localstack", "localstack", "dummy")),
       )
       if err != nil {
           // エラーハンドリング
       }

       logClient = cloudwatchlogs.NewFromConfig(cfg)
   } else {
       // 本番環境用の設定
       cfg, err := config.LoadDefaultConfig(context.Background())
       if err != nil {
           // エラーハンドリング
       }

       logClient = cloudwatchlogs.NewFromConfig(cfg)
   }
   ```

#### 1.4.5.4. LocalStackを使用した実践的な例

以下は、LocalStackでCloudWatch Logsを使用した実際の例です：

1. **ロググループとログストリームの作成**

    ```bash
    # ロググループの作成
    awslocal logs create-log-group --log-group-name /aws-observability-ecommerce/dev/backend

    # ログストリームの作成
    awslocal logs create-log-stream --log-group-name /aws-observability-ecommerce/dev/backend --log-stream-name api-server
    ```

2. **ログイベントの書き込み**

    ```bash
    # ログイベントの書き込み
    awslocal logs put-log-events \
      --log-group-name "/aws-observability-ecommerce/dev/backend" \
      --log-stream-name "api-server" \
      --log-events \
      timestamp=$(date +%s)000,message='{"level":"info","message":"API request received","path":"/api/products","method":"GET","requestId":"req-123"}' \
      timestamp=$(($(date +%s) + 1))000,message='{"level":"info","message":"API request completed","path":"/api/products","method":"GET","requestId":"req-123","statusCode":200,"responseTime":42}'
    ```

3. **ログイベントの読み取り**

    ```bash
    # ログイベントの読み取り
    awslocal logs get-log-events \
      --log-group-name "/aws-observability-ecommerce/dev/backend" \
      --log-stream-name "api-server"
    ```

4. **ログフィルタリング（基本的な機能のみ）**

    ```bash
    # ログのフィルタリング
    awslocal logs filter-log-events \
      --log-group-name "/aws-observability-ecommerce/dev/backend" \
      --filter-pattern '{ $.level = "error" }'
    ```

LocalStackを使用することで、本番環境と同様のAWSサービスをローカルで利用でき、コストや複雑さを気にせずに開発やテストを行うことができます。ただし、実際のAWS環境との違いを理解し、必要に応じて対応することが重要です。

次のセクションでは、これらの概念を踏まえた上で、Goのslogを使用した構造化ログの実装方法について学んでいきます。

## 1.5. バックエンドの構造化ログ実装

効果的なログ記録はアプリケーションの観測性を高める上で不可欠です。特に分散システムやクラウド環境では、構造化されたログは問題のトラブルシューティングや分析を大幅に改善します。このセクションでは、Go言語のアプリケーションにおける構造化ログの実装方法と、Echo Webフレームワークとの統合について学びます。

まず、必要なディレクトリとファイルを作成しましょう：

```bash
## ロガー関連のディレクトリとファイルを作成
mkdir -p backend/internal/logger
touch backend/internal/logger/{logger.go,error.go}

## APIミドルウェア関連のディレクトリとファイルを作成
mkdir -p backend/internal/api/middleware
touch backend/internal/api/middleware/{logger.go,request_id.go,auth_logger.go,log_middleware.go,setup.go,error_handler.go}

## 既存のコンフィグファイルを更新
## touch backend/internal/config/config.go (すでに存在)

## ハンドラー関連ファイルも更新
## touch backend/internal/api/handlers/health_handler.go (すでに存在)
## touch backend/internal/api/handlers/product_handler.go (すでに存在)
```

### 1.5.1. 構造化ログの概念と利点

構造化ログとは、ログメッセージを単なるテキスト文字列ではなく、構造化されたデータ形式（通常はJSON）で記録する方法です。これにより、ログデータの処理、検索、分析が容易になります。

#### 1.5.1.1. 従来のテキストベースログと構造化ログの比較

**従来のテキストベースログ**:

```text
[2023-10-21 15:04:23] INFO: User 12345 logged in successfully from 192.168.1.1
```

**構造化ログ（JSON形式）**:

```json
{
  "timestamp": "2023-10-21T15:04:23Z",
  "level": "info",
  "message": "User logged in successfully",
  "user_id": 12345,
  "ip_address": "192.168.1.1"
}
```

#### 1.5.1.2. 構造化ログの主な利点

1. **検索性の向上**
   - 特定のフィールドに基づいて簡単に検索可能
   - 例：特定のユーザーIDに関連するすべてのログを検索

2. **分析のしやすさ**
   - ログデータをより簡単に集計、フィルタリング、分析できる
   - クエリ言語（CloudWatch Logs Insightsなど）を使用した高度な分析

3. **機械処理の効率化**
   - ログデータを自動的に処理するシステムにとって理想的
   - モニタリングシステムとの統合が容易

4. **コンテキスト情報の豊富さ**
   - より多くの構造化されたメタデータを含めることが可能
   - トラブルシューティングに必要な情報をすべて一箇所に記録

5. **形式の一貫性**
   - すべてのログエントリが一貫した形式で記録される
   - アプリケーション全体で統一されたログ形式

#### 1.5.1.3. JSONフォーマットを使用する利点

JSONは構造化ログの標準的なフォーマットとして広く採用されています。その主な利点は：

1. **広範なサポート**
   - ほぼすべてのプログラミング言語とツールでサポート
   - CloudWatch Logsを含む多くのログサービスがネイティブにJSONをサポート

2. **人間にも機械にも読みやすい**
   - 開発者にとって読みやすく、デバッグが容易
   - 同時にコンピュータでの処理も効率的

3. **柔軟なスキーマ**
   - 厳密なスキーマ定義が不要
   - 必要に応じて新しいフィールドを追加可能

4. **階層的なデータ表現**
   - ネストされたデータ構造を表現可能
   - 複雑なコンテキスト情報も適切に表現できる

### 1.5.2. Goの標準ロギングライブラリslogの紹介

Go 1.21から、標準ライブラリに構造化ログのためのパッケージ `log/slog` が追加されました。slogは、Goの従来のログライブラリの制限を克服し、構造化ログ記録のための強力な機能を提供します。

#### 1.5.2.1. slogの基本概念

slogは以下の主要なコンポーネントで構成されています：

1. **Logger**: ログエントリを記録するためのメインインターフェース
2. **Handler**: ログレコードの処理とフォーマット方法を定義
3. **Record**: 個々のログエントリを表す構造体
4. **Attr**: 名前と値のペアで表されるログ属性（キー・バリューペア）

#### 1.5.2.2. slogの主な特徴

- **構造化ログのネイティブサポート**
- **複数のログレベル** (ERROR, WARN, INFO, DEBUG)
- **コンテキスト伝播のサポート**
- **複数の出力フォーマット** (JSONやテキスト)
- **カスタムハンドラーによる拡張性**
- **既存のログライブラリとの互換性**

#### 1.5.2.3. slogの基本的な使用例

```go
package main

import (
    "log/slog"
    "os"
)

func main() {
    // JSONハンドラーを使用したロガーの作成
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    // デフォルトロガーとして設定
    slog.SetDefault(logger)

    // 基本的なログ記録
    slog.Info("Application started", "version", "1.0.0")

    // 構造化された追加情報を含むログ
    slog.Info("User logged in",
        "user_id", 12345,
        "ip_address", "192.168.1.1",
        "login_time", "2023-10-21T15:04:23Z",
    )

    // エラーログの記録
    err := someFunction()
    if err != nil {
        slog.Error("Operation failed",
            "error", err,
            "operation", "data_processing",
        )
    }
}
```

出力例：

```json
{"time":"2023-10-21T15:04:23Z","level":"INFO","msg":"Application started","version":"1.0.0"}
{"time":"2023-10-21T15:04:23Z","level":"INFO","msg":"User logged in","user_id":12345,"ip_address":"192.168.1.1","login_time":"2023-10-21T15:04:23Z"}
{"time":"2023-10-21T15:04:24Z","level":"ERROR","msg":"Operation failed","error":"resource not found","operation":"data_processing"}
```

### 1.5.3. slogを使用した構造化ログの設計

eコマースアプリケーションにslogを効果的に導入するためには、適切な設計とカスタマイズが必要です。ここでは、アプリケーション全体で一貫したログ記録を実現するための設計パターンを示します。

#### 1.5.3.1. ロガーの初期化とカスタマイズ

まず、`internal/logger/logger.go` ファイルを作成して、アプリケーション全体で使用するロガーを初期化します：

```go
package logger

import (
    "context"
    "io"
    "log/slog"
    "os"
    "time"
)

// Config はロガーの設定を表す構造体
type Config struct {
    Environment string
    LogLevel    string
    ServiceName string
    Version     string
}

// ロガーのインスタンスを格納するグローバル変数
var defaultLogger *slog.Logger

// Init はロガーを初期化する
func Init(cfg Config) *slog.Logger {
    // ログレベルの設定
    var level slog.Level
    switch cfg.LogLevel {
    case "debug":
        level = slog.LevelDebug
    case "info":
        level = slog.LevelInfo
    case "warn":
        level = slog.LevelWarn
    case "error":
        level = slog.LevelError
    default:
        level = slog.LevelInfo
    }

    // 出力先の設定（本番環境では別の書き込み先を設定することもある）
    var w io.Writer = os.Stdout

    // JSONハンドラーのオプション設定
    opts := &slog.HandlerOptions{
        Level: level,
        ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
            // タイムスタンプのフォーマット変更
            if a.Key == "time" {
                if t, ok := a.Value.Any().(time.Time); ok {
                    a.Value = slog.StringValue(t.Format(time.RFC3339))
                }
            }
            return a
        },
    }

    // JSONハンドラーの作成
    var handler slog.Handler
    handler = slog.NewJSONHandler(w, opts)

    // アプリケーション全体の共通属性を持つハンドラーをラップ
    handler = NewContextHandler(handler, map[string]interface{}{
      "service":     cfg.ServiceName,
      "environment": cfg.Environment,
      "version":     cfg.Version,
    })

    // ロガーの作成と設定
    logger := slog.New(handler)
    slog.SetDefault(logger)
    defaultLogger = logger

    return logger
}

// Logger は現在のコンテキストに基づいてロガーを返す
func Logger(ctx context.Context) *slog.Logger {
    if ctx == nil {
        return defaultLogger
    }

    // コンテキストからロガーを取得
    if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
        return logger
    }

    return defaultLogger
}

// WithLogger は指定されたロガーをコンテキストに追加する
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
    return context.WithValue(ctx, loggerKey{}, logger)
}

// コンテキストキー
type loggerKey struct{}

// カスタムコンテキストハンドラー
type contextHandler struct {
    handler slog.Handler
    attrs   []slog.Attr
}

// NewContextHandler は共通属性を持つハンドラーを作成する
func NewContextHandler(handler slog.Handler, attrs map[string]interface{}) slog.Handler {
    slogAttrs := make([]slog.Attr, 0, len(attrs))
    for k, v := range attrs {
        slogAttrs = append(slogAttrs, slog.Any(k, v))
    }
    return &contextHandler{
        handler: handler,
        attrs:   slogAttrs,
    }
}

// Enabled はハンドラーのEnabled関数を呼び出す
func (h *contextHandler) Enabled(ctx context.Context, level slog.Level) bool {
    return h.handler.Enabled(ctx, level)
}

// Handle はすべてのログレコードに共通属性を追加する
func (h *contextHandler) Handle(ctx context.Context, r slog.Record) error {
    // 共通属性をレコードに追加
    for _, attr := range h.attrs {
        r.AddAttrs(attr)
    }
    return h.handler.Handle(ctx, r)
}

// WithAttrs は新しい属性を持つハンドラーを返す
func (h *contextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
    return &contextHandler{
        handler: h.handler.WithAttrs(attrs),
        attrs:   h.attrs,
    }
}

// WithGroup はグループを持つハンドラーを返す
func (h *contextHandler) WithGroup(name string) slog.Handler {
    return &contextHandler{
        handler: h.handler.WithGroup(name),
        attrs:   h.attrs,
    }
}
```

#### 1.5.3.2. アプリケーションでのロガーの使用

次に、`cmd/api/main.go` でロガーを初期化し、アプリケーション全体で使用できるようにします：

```go
package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/labstack/echo/v4"

    "github.com/y-nosuke/aws-observability-ecommerce/internal/api/handlers"
    "github.com/y-nosuke/aws-observability-ecommerce/internal/api/middleware"
    "github.com/y-nosuke/aws-observability-ecommerce/internal/api/router"
    "github.com/y-nosuke/aws-observability-ecommerce/internal/config"
    "github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

func main() {
    // コンテキストの初期化（シグナルハンドリング）
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    // ロガーの初期化
    log := logger.Init(logger.Config{
        Environment: config.Config.App.Environment,
        LogLevel:    config.Config.Log.Level,
        ServiceName: config.Config.App.Name,
        Version:     config.Config.App.Version,
    })

    // アプリケーションの起動をログに記録
    log.Info("Starting application",
        "version", config.Config.App.Version,
        "environment", config.Config.App.Environment)

    // Echoインスタンスの作成
    e := echo.New()
    e.HideBanner = true
    e.HidePort = true

    // カスタムエラーハンドラーの設定
    e.HTTPErrorHandler = middleware.ErrorHandler

    // ミドルウェアの設定（ロガーも含む）
    middleware.SetupMiddleware(e)

    // ハンドラーの作成
    healthHandler := handlers.NewHealthHandler()
    productHandler := handlers.NewProductHandler()

    // ルーターの設定
    router.SetupRoutes(e, healthHandler, productHandler)

    // サーバーの起動（非同期）
    go func() {
        log.Info("Server starting", "port", config.Config.Server.Port)
        if err := e.Start(":" + config.Config.Server.Port); err != nil {
            log.Error("Server shutdown", "error", err)
        }
    }()

    // シグナルを待機
    <-ctx.Done()
    log.Info("Shutdown signal received, gracefully shutting down...")

    // グレースフルシャットダウン
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := e.Shutdown(shutdownCtx); err != nil {
        log.Error("Server shutdown failed", "error", err)
    }

    log.Info("Server has been shutdown gracefully")
}
```

#### 1.5.3.3. 構造化された属性（Attr）の使用

`slog`では、`Attr`型を使用して構造化された属性をログに追加できます。これにより、必要なコンテキスト情報を一貫した形式で記録できます：

```go
// 基本的な属性の追加
slog.Info("Product viewed",
    "product_id", productID,
    "category", category,
    "user_id", userID)

// グループ化された属性
slog.Info("Order processed",
    slog.Group("order",
        "id", orderID,
        "amount", amount,
        "items_count", itemsCount,
    ),
    slog.Group("customer",
        "id", customerID,
        "type", customerType,
    ))

// 時間関連の属性
start := time.Now()
// ... 何らかの処理 ...
elapsed := time.Since(start)
slog.Info("Operation completed",
    "operation", "data_export",
    "duration_ms", elapsed.Milliseconds())
```

### 1.5.4. ログレベル管理（ERROR/WARN/INFO/DEBUG）の実装

ログレベルは、ログメッセージの重要度を示すために使用されます。slogには標準的なログレベルが定義されていますが、アプリケーションの要件に合わせてカスタマイズすることもできます。

#### 1.5.4.1. slogの標準ログレベル

slogには以下の標準ログレベルが定義されています：

- **ERROR**: エラーや例外など、アプリケーションの動作に影響を与える重要な問題
- **WARN**: 潜在的な問題や将来的に問題になる可能性のある警告
- **INFO**: アプリケーションの通常の動作に関する情報
- **DEBUG**: デバッグに役立つ詳細な情報（開発環境でのみ使用）

#### 1.5.4.2. 環境に応じたログレベルの設定

アプリケーションの環境（開発、ステージング、本番など）に応じて適切なログレベルを設定することが重要です。これを設定ファイルで管理する例を示します：

##### 1.5.4.2.1. `internal/config/config.go`

```go
package config

import (
    "github.com/spf13/viper"
)

// Config はアプリケーション設定を表す構造体
type AppConfig struct {
    App struct {
        Name        string
        Version     string
        Environment string
    }

    Log struct {
        Level  string
        Format string
    }

    // ... 他の設定 ...
}

var Config *AppConfig

// ... 他の設定 ...

// LoadConfig は環境変数と設定ファイルから設定をロードします
func LoadConfig() (*AppConfig, error) {
    // ... 環境変数のデフォルト値の設定 ...

    // ログ関連のデフォルト設定
    viper.SetDefault("log.level", "info")  // デフォルトはinfo
    viper.SetDefault("log.format", "json") // デフォルトはJSON

    // 開発環境ではデバッグログを有効化
    if viper.GetString("app.environment") == "development" {
        viper.SetDefault("log.level", "debug")
    }

    // 本番環境では警告以上のみを記録
    if viper.GetString("app.environment") == "production" {
        viper.SetDefault("log.level", "warn")
    }

    // ... 設定の読み込み ...

    if err := viper.BindEnv("log.level", "LOG_LEVEL"); err != nil {
      log.Fatalf("Failed to bind env var: %v", err)
    }
    if err := viper.BindEnv("log.format", "LOG_FORMAT"); err != nil {
      log.Fatalf("Failed to bind env var: %v", err)
    }

    // 設定を構造体にマッピング
    if err := viper.Unmarshal(&Config); err != nil {
        panic(err)
    }
}
```

##### 1.5.4.2.2. `.env` ファイルの例

環境変数を使用して様々な環境でログレベルを設定するために、以下のような`.envrc`ファイルを作成します：

```bash
## 開発環境設定
export APP_NAME=aws-observability-ecommerce
export APP_VERSION=1.0.0
export APP_ENV=development
export PORT=8080
export LOG_LEVEL=debug
export LOG_FORMAT=json
```

本番環境では、以下のように設定することが想定されます：

```bash
## 本番環境設定
export APP_NAME=aws-observability-ecommerce
export APP_VERSION=1.0.0
export APP_ENV=production
export PORT=8080
export LOG_LEVEL=info  # または必要に応じて warn
export LOG_FORMAT=json
```

#### 1.5.4.3. ログレベルに適したメッセージの例

各ログレベルで記録すべき内容の例を示します：

##### 1.5.4.3.1. ERRORレベル

```go
// データベース接続エラー
if err := db.Connect(); err != nil {
    slog.Error("Failed to connect to database",
        "error", err,
        "database", config.Config.Database.Name,
        "retry_count", retryCount)
}

// APIリクエストの処理に失敗
slog.Error("Failed to process API request",
    "error", err,
    "method", c.Request().Method,
    "path", c.Path(),
    "client_ip", c.RealIP(),
    "status_code", statusCode)
```

##### 1.5.4.3.2. WARNレベル

```go
// データベース接続が遅い
if connectionTime > slowThreshold {
    slog.Warn("Database connection is slow",
        "connection_time_ms", connectionTime.Milliseconds(),
        "threshold_ms", slowThreshold.Milliseconds())
}

// 在庫が少ない商品
if product.StockLevel < lowStockThreshold {
    slog.Warn("Product stock is low",
        "product_id", product.ID,
        "current_stock", product.StockLevel,
        "threshold", lowStockThreshold)
}
```

##### 1.5.4.3.3. INFOレベル

```go
// アプリケーション起動情報
slog.Info("Application started",
    "version", config.Config.App.Version,
    "environment", config.Config.App.Environment)

// 注文処理完了
slog.Info("Order processed successfully",
    "order_id", order.ID,
    "customer_id", order.CustomerID,
    "total_amount", order.TotalAmount)
```

##### 1.5.4.3.4. DEBUGレベル

```go
// リクエストの詳細情報
slog.Debug("Received API request",
    "method", c.Request().Method,
    "path", c.Path(),
    "query_params", c.QueryParams(),
    "headers", headers)

// データベースクエリの実行詳細
slog.Debug("Executing database query",
    "query", query,
    "params", params,
    "timeout", timeout)
```

### 1.5.5. コンテキスト情報の付与（リクエストID、ユーザーIDなど）

マイクロサービス環境やクラウドベースのアプリケーションでは、リクエストの追跡とコンテキスト情報の伝播が重要です。slogとEchoフレームワークを組み合わせて、リクエストの追跡とコンテキスト情報を付与する方法を見ていきましょう。

#### 1.5.5.1. コンテキスト対応のロガー作成

まず、Echoのコンテキストからロガーを取得・設定するためのヘルパー関数を追加します：

```go
// internal/api/middleware/logger.go
package middleware

import (
    "log/slog"

    "github.com/labstack/echo/v4"

    "github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

// コンテキストキー
const (
    RequestIDKey = "request_id"
    LoggerKey    = "logger"
)

// GetRequestID はEchoコンテキストからリクエストIDを取得する
func GetRequestID(c echo.Context) string {
    if requestID, ok := c.Get(RequestIDKey).(string); ok {
        return requestID
    }
    return ""
}

// GetLogger はEchoコンテキストからロガーを取得する
func GetLogger(c echo.Context) *slog.Logger {
    if l, ok := c.Get(LoggerKey).(*slog.Logger); ok {
        return l
    }
    return logger.Logger(c.Request().Context()) // デフォルトロガーを返す
}

// SetLogger はEchoコンテキストにロガーを設定する
func SetLogger(c echo.Context, l *slog.Logger) {
    c.Set(LoggerKey, l)
}
```

#### 1.5.5.2. リクエストIDの生成と追加

次に、各リクエストに一意のIDを割り当てるミドルウェアを実装します：

```go
// internal/api/middleware/request_id.go
package middleware

import (
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "log/slog"

    "github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

// RequestIDMiddleware は各リクエストに一意のIDを割り当てるミドルウェア
func RequestIDMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // リクエストIDがヘッダーに含まれているか確認
            requestID := c.Request().Header.Get("X-Request-ID")
            if requestID == "" {
                // 含まれていない場合は新しいIDを生成
                requestID = uuid.New().String()
            }

            // コンテキストにリクエストIDを設定
            c.Set(RequestIDKey, requestID)
            // レスポンスヘッダーにもリクエストIDを設定
            c.Response().Header().Set("X-Request-ID", requestID)

            // リクエストIDを含むロガーを作成してコンテキストに設定
            log := logger.Logger(c.Request().Context()).With("request_id", requestID)
            SetLogger(c, log)

            // 次のハンドラーを呼び出す
            return next(c)
        }
    }
}
```

#### 1.5.5.3. ユーザー認証情報の追加

認証されたユーザーの情報をログに追加するミドルウェアも実装できます：

```go
// internal/api/middleware/auth_logger.go
package middleware

import (
    "github.com/labstack/echo/v4"
    "log/slog"
)

// AuthLoggerMiddleware は認証情報をロガーに追加するミドルウェア
func AuthLoggerMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // 現在のロガーを取得
            log := GetLogger(c)

            // 認証されたユーザーがいれば情報を追加
            // 注意: 実際の認証の実装に合わせて調整する必要があります
            if userID := getUserIDFromContext(c); userID != "" {
                log = log.With("user_id", userID)
                SetLogger(c, log)
            }

            // 管理者ユーザーの場合は追加情報
            if isAdmin := isAdminUser(c); isAdmin {
                log = log.With("user_role", "admin")
                SetLogger(c, log)
            }

            return next(c)
        }
    }
}

// ダミー実装 - 実際の認証システムに合わせて実装する必要があります
func getUserIDFromContext(c echo.Context) string {
    if userID, ok := c.Get("user_id").(string); ok {
        return userID
    }
    return ""
}

// ダミー実装 - 実際の認証システムに合わせて実装する必要があります
func isAdminUser(c echo.Context) bool {
    if role, ok := c.Get("user_role").(string); ok {
        return role == "admin"
    }
    return false
}
```

#### 1.5.5.4. 構造化された例外情報の記録

Goのエラー処理と構造化ログを組み合わせるためのヘルパー関数を作成します：

```go
// internal/logger/error.go
package logger

import (
    "errors"
    "fmt"
    "log/slog"
    "runtime"
    "strings"
)

// ErrorAttr はエラーを構造化された属性に変換する
func ErrorAttr(err error) slog.Attr {
    return ErrorAttrWithKey("error", err)
}

// ErrorAttrWithKey は指定されたキーでエラーを構造化された属性に変換する
func ErrorAttrWithKey(key string, err error) slog.Attr {
    if err == nil {
        return slog.String(key, "")
    }

    // エラー情報をマップに変換
    errorInfo := map[string]interface{}{
        "message": err.Error(),
    }

    // スタックトレース情報の追加（開発環境のみ）
    var programCounter [50]uintptr
    n := runtime.Callers(2, programCounter[:])
    frames := runtime.CallersFrames(programCounter[:n])

    stackTrace := make([]string, 0, n)
    for {
        frame, more := frames.Next()
        if !strings.Contains(frame.File, "runtime/") {
            stackTrace = append(stackTrace, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
        }
        if !more {
            break
        }
        if len(stackTrace) >= 5 {
            // スタックトレースは5フレームまで
            break
        }
    }

    if len(stackTrace) > 0 {
        errorInfo["stack_trace"] = stackTrace
    }

    // エラーがラップされている場合は展開
    var unwrapped error
    if errors.Unwrap(err) != nil {
        unwrapped = errors.Unwrap(err)
        errorInfo["cause"] = unwrapped.Error()
    }

    return slog.Any(key, errorInfo)
}
```

このヘルパー関数を使用して、エラーを構造化されたログに記録できます：

```go
// ハンドラー内でのエラーログ
if err := service.ProcessOrder(order); err != nil {
    log := GetLogger(c)
    log.Error("Failed to process order",
        logger.ErrorAttr(err),
        "order_id", order.ID,
        "customer_id", order.CustomerID)
    return echo.NewHTTPError(http.StatusInternalServerError, "Order processing failed")
}
```

### 1.5.6. ミドルウェアを使用したリクエスト/レスポンスのログ記録

Echoフレームワークのミドルウェア機能を利用して、すべてのHTTPリクエストとレスポンスを自動的にログに記録するミドルウェアを実装します。

```go
// internal/api/middleware/log_middleware.go
package middleware

import (
    "bytes"
    "io"
    "net/http"
    "time"

    "slices"

    "github.com/labstack/echo/v4"
)

// LoggerConfig はロギングミドルウェアの設定
type LoggerConfig struct {
    // ログに含めないURLパス（ヘルスチェックなど）
    SkipPaths []string
    // リクエスト本文をログに含めるかどうか
    LogRequestBody bool
    // レスポンス本文をログに含めるかどうか
    LogResponseBody bool
    // 最大本文サイズ（バイト単位）
    MaxBodySize int
}

// DefaultLoggerConfig はLoggerConfigのデフォルト値
var DefaultLoggerConfig = LoggerConfig{
    SkipPaths:       []string{"/api/health", "/api/metrics"},
    LogRequestBody:  false,
    LogResponseBody: false,
    MaxBodySize:     1024, // 1KB
}

// LoggerMiddleware はリクエストとレスポンスをログに記録するミドルウェア
func LoggerMiddleware(config LoggerConfig) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // ヘルスチェックなど特定のパスはスキップ
            path := c.Request().URL.Path
            if slices.Contains(config.SkipPaths, path) {
                return next(c)
            }

            start := time.Now()

            // リクエスト情報を記録
            log := GetLogger(c)
            reqBody := ""
            if config.LogRequestBody && c.Request().Body != nil {
                // リクエスト本文を読み取り、バッファに保存
                buf, err := io.ReadAll(io.LimitReader(c.Request().Body, int64(config.MaxBodySize)))
                if err != nil {
                    return err
                }
                reqBody = string(buf)
                // 本文を復元
                c.Request().Body = io.NopCloser(bytes.NewBuffer(buf))
            }

            // リクエスト情報のログ記録
            log.Info("API request received",
                "method", c.Request().Method,
                "path", c.Request().URL.Path,
                "query", c.Request().URL.RawQuery,
                "remote_ip", c.RealIP(),
                "user_agent", c.Request().UserAgent())

            if config.LogRequestBody && reqBody != "" {
                log.Debug("Request body", "body", reqBody)
            }

            // レスポンスをキャプチャするためのレスポンスライター
            resBody := new(bytes.Buffer)
            mw := io.MultiWriter(c.Response().Writer, resBody)
            writer := &bodyDumpResponseWriter{
                ResponseWriter: c.Response().Writer,
                Writer:         mw,
            }
            c.Response().Writer = writer

            // 次のハンドラーを呼び出し
            err := next(c)

            // レスポンス情報を記録
            elapsed := time.Since(start)
            statusCode := c.Response().Status
            responseSize := c.Response().Size

            // ログレベルをステータスコードに基づいて決定
            var logFunc func(msg string, args ...interface{})
            if statusCode >= 500 {
                logFunc = log.Error
            } else if statusCode >= 400 {
                logFunc = log.Warn
            } else {
                logFunc = log.Info
            }

            logFunc("API request completed",
                "method", c.Request().Method,
                "path", c.Request().URL.Path,
                "status", statusCode,
                "elapsed_ms", elapsed.Milliseconds(),
                "size", responseSize)

            // レスポンス本文のログ記録（必要な場合）
            if config.LogResponseBody && resBody.Len() > 0 && resBody.Len() <= config.MaxBodySize {
                log.Debug("Response body", "body", resBody.String())
            }

            return err
        }
    }
}

// bodyDumpResponseWriter はレスポンス本文をキャプチャするためのレスポンスライター
type bodyDumpResponseWriter struct {
    http.ResponseWriter
    Writer io.Writer
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
    return w.Writer.Write(b)
}
```

#### 1.5.6.1. ミドルウェアの設定

作成したミドルウェアをEchoインスタンスに設定します：

```go
// internal/api/middleware/setup.go
package middleware

import (
    "github.com/labstack/echo/v4"
    echomw "github.com/labstack/echo/v4/middleware"

    "github.com/y-nosuke/aws-observability-ecommerce/internal/config"
)

// SetupMiddleware はすべてのミドルウェアを設定する
func SetupMiddleware(e *echo.Echo) {
    // 基本的なミドルウェア
    e.Use(echomw.Recover()) // パニック回復
    e.Use(echomw.CORS())    // CORS対応

    // カスタムミドルウェア
    e.Use(RequestIDMiddleware()) // リクエストID生成

    // リクエスト/レスポンスのログ記録
    loggerConfig := DefaultLoggerConfig
    // 開発環境では本文も記録
    if config.Config.App.Environment == "development" {
        loggerConfig.LogRequestBody = true
        loggerConfig.LogResponseBody = true
    }
    e.Use(LoggerMiddleware(loggerConfig))

    // 認証情報のログ付与（認証後に実行）
    e.Use(AuthLoggerMiddleware())
}
```

### 1.5.7. Echo統合とハンドラーへのログ組み込み

最後に、これまでに実装したログ機能をEchoのハンドラーに統合します。ハンドラー内で適切にログを記録するパターンを示します。

#### 1.5.7.1. 商品ハンドラーとヘルスチェックハンドラーでのログ使用例

```go
// internal/api/handlers/product_handler.go
package handlers

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"

    "github.com/y-nosuke/aws-observability-ecommerce/internal/api/middleware"
    "github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

// ProductHandler は商品関連のハンドラーを表す構造体
type ProductHandler struct {
    // ... 前回と同じ定義 ...
}

// HandleGetProducts は商品一覧を取得するハンドラー関数
func (h *ProductHandler) HandleGetProducts(c echo.Context) error {
    // コンテキストからロガーを取得
    log := middleware.GetLogger(c)

    // リクエストパラメータを取得してログに記録
    page := 1
    pageStr := c.QueryParam("page")
    if pageStr != "" {
        parsedPage, err := strconv.Atoi(pageStr)
        if err != nil {
            log.Warn("Invalid page parameter", "page", pageStr, "error", err.Error())
        } else if parsedPage > 0 {
            page = parsedPage
        }
    }

    pageSize := 10 // デフォルトのページサイズ
    pageSizeStr := c.QueryParam("page_size")
    if pageSizeStr != "" {
        parsedPageSize, err := strconv.Atoi(pageSizeStr)
        if err != nil {
            log.Warn("Invalid page_size parameter", "page_size", pageSizeStr, "error", err.Error())
        } else if parsedPageSize > 0 && parsedPageSize <= 50 {
            pageSize = parsedPageSize
        }
    }

    // カテゴリーIDによるフィルタリング
    categoryID := 0
    categoryIDStr := c.QueryParam("category_id")
    if categoryIDStr != "" {
        parsedCategoryID, err := strconv.Atoi(categoryIDStr)
        if err != nil {
            log.Warn("Invalid category_id parameter", "category_id", categoryIDStr, "error", err.Error())
        } else if parsedCategoryID > 0 {
            categoryID = parsedCategoryID
        }
    }

    // リクエストパラメータの詳細をデバッグレベルでログに記録
    log.Debug("Product list request parameters",
        "page", page,
        "page_size", pageSize,
        "category_id", categoryID)

    var filteredProducts []Product
    if categoryID > 0 {
        // カテゴリーでフィルタリング
        for _, p := range h.products {
            if p.CategoryID == categoryID {
                filteredProducts = append(filteredProducts, p)
            }
        }

        log.Info("Products filtered by category",
            "category_id", categoryID,
            "filtered_count", len(filteredProducts),
            "total_count", len(h.products))
    } else {
        // フィルタリングなし
        filteredProducts = h.products

        log.Info("All products requested",
            "total_count", len(filteredProducts))
    }

    // 製品の総数
    totalItems := len(filteredProducts)

    // 総ページ数を計算
    totalPages := (totalItems + pageSize - 1) / pageSize

    // 現在のページの開始と終了インデックスを計算
    startIndex := (page - 1) * pageSize
    endIndex := startIndex + pageSize
    if endIndex > totalItems {
        endIndex = totalItems
    }

    // ページに表示する製品を取得
    var pageProducts []Product
    if startIndex < totalItems {
        pageProducts = filteredProducts[startIndex:endIndex]
    } else {
        pageProducts = []Product{}
        log.Warn("Requested page exceeds available products",
            "page", page,
            "total_pages", totalPages)
    }

    // レスポンスを構築
    response := PaginatedResponse{
        Items:      pageProducts,
        TotalItems: totalItems,
        Page:       page,
        PageSize:   pageSize,
        TotalPages: totalPages,
    }

    // レスポンスの送信をログに記録
    log.Info("Products list response generated",
        "page", page,
        "page_size", pageSize,
        "total_items", totalItems,
        "items_returned", len(pageProducts))

    return c.JSON(http.StatusOK, response)
}

// ・・・

// HandleGetCategories はカテゴリー一覧を取得するハンドラー関数
func (h *ProductHandler) HandleGetCategories(c echo.Context) error {
    // コンテキストからロガーを取得
    log := middleware.GetLogger(c)

    // リクエストの処理開始をログに記録
    log.Info("Categories list requested",
        "remote_ip", c.RealIP(),
        "request_id", middleware.GetRequestID(c))

    categories := h.GetCategories()

    // レスポンスの送信をログに記録
    log.Info("Categories list response generated",
        "categories_count", len(categories))

    return c.JSON(http.StatusOK, categories)
}
```

ヘルスチェックハンドラーも同様に更新します：

```go
// internal/api/handlers/health_handler.go
package handlers

import (
    "net/http"
    "runtime"
    "time"

    "github.com/labstack/echo/v4"

    "github.com/y-nosuke/aws-observability-ecommerce/internal/api/middleware"
    "github.com/y-nosuke/aws-observability-ecommerce/internal/config"
)

// HealthResponse はヘルスチェックの応答を表す構造体
type HealthResponse struct {
    Status    string                 `json:"status"`
    Timestamp string                 `json:"timestamp"`
    Version   string                 `json:"version"`
    Uptime    int64                  `json:"uptime"`
    Resources map[string]interface{} `json:"resources"`
    Services  map[string]interface{} `json:"services"`
}

// HealthHandler はヘルスチェックのハンドラーを表す構造体
type HealthHandler struct {
    startTime time.Time
    version   string
}

// NewHealthHandler は新しいヘルスハンドラーを作成します
func NewHealthHandler() *HealthHandler {
    return &HealthHandler{
        startTime: time.Now(),
        version:   config.Config.App.Version, // アプリケーションバージョン
    }
}

// HandleHealthCheck はヘルスチェックエンドポイントのハンドラー関数
func (h *HealthHandler) HandleHealthCheck(c echo.Context) error {
    // コンテキストからロガーを取得
    log := middleware.GetLogger(c)

    // リクエストの処理開始をログに記録
    log.Debug("Health check request received",
        "method", c.Request().Method,
        "path", c.Path(),
        "remote_ip", c.RealIP(),
        "request_id", middleware.GetRequestID(c),
    )

    // サービスの状態をチェック（ここでは簡易的にすべて稼働中とする）
    services := map[string]interface{}{
        "api": map[string]string{
            "status": "up",
        },
        // 実際のアプリケーションでは、データベース接続などをチェックする
        // "database": checkDatabaseConnection(),
    }

    // システムリソースの状態を取得
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)

    resources := map[string]interface{}{
        "memory": map[string]interface{}{
            "allocated": memStats.Alloc,
            "total":     memStats.TotalAlloc,
            "system":    memStats.Sys,
        },
        "goroutines": runtime.NumGoroutine(),
    }

    // レスポンスを構築
    response := &HealthResponse{
        Status:    "ok",
        Timestamp: time.Now().Format(time.RFC3339),
        Version:   h.version,
        Uptime:    time.Since(h.startTime).Milliseconds(),
        Resources: resources,
        Services:  services,
    }

    // レスポンスの送信をログに記録
    log.Info("Health check completed",
        "status", response.Status,
        "uptime_ms", response.Uptime,
        "goroutines", resources["goroutines"],
    )

    return c.JSON(http.StatusOK, response)
}
```

#### 1.5.7.2. エラーハンドリングとログ記録

さらに、グローバルなエラーハンドラーを設定して、すべてのエラーを適切にログに記録できます：

```go
// internal/api/middleware/error_handler.go
package middleware

import (
    "net/http"

    "github.com/labstack/echo/v4"

    "github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

// ErrorHandler はグローバルなエラーハンドラー
func ErrorHandler(err error, c echo.Context) {
    // コンテキストからロガーを取得
    log := GetLogger(c)

    code := http.StatusInternalServerError
    message := "Internal server error"

    // Echoの標準エラーの場合
    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
        if message, ok = he.Message.(string); ok {
            message = "unknown"
        }
    }

    // エラーの重大度に応じてログレベルを変える
    if code >= 500 {
        log.Error("Request error",
            logger.ErrorAttr(err),
            "status_code", code,
            "path", c.Request().URL.Path,
            "method", c.Request().Method)
    } else if code >= 400 {
        log.Warn("Request warning",
            "error", err.Error(),
            "status_code", code,
            "path", c.Request().URL.Path,
            "method", c.Request().Method)
    }

    // エラーレスポンスを返す
    if !c.Response().Committed {
        if c.Request().Method == http.MethodHead {
            err = c.NoContent(code)
        } else {
            err = c.JSON(code, map[string]interface{}{
                "error": message,
            })
        }
        if err != nil {
            log.Error("Failed to send error response", logger.ErrorAttr(err))
        }
    }
}
```

エラーハンドラーの設定は、`main.go`に追加します：

```go
// cmd/api/main.go
func main() {
    // ... 前述のコード ...

    // Echoインスタンスの作成
    e := echo.New()
    e.HideBanner = true
    e.HidePort = true

    // カスタムエラーハンドラーの設定
    e.HTTPErrorHandler = middleware.ErrorHandler

    // ... 残りのコード ...
}
```

#### 1.5.7.3. ロガーのパフォーマンスに関する考慮事項

構造化ログは非常に有用ですが、過度のロギングはパフォーマンスに影響を与える可能性があります。以下のポイントに注意してください：

1. **適切なログレベルの使用**
   - 本番環境では、デバッグログを無効にする
   - 頻繁に呼び出されるコードパスでは詳細なログを避ける

2. **条件付きロギング**
   - 高負荷なログ処理は実際にロギングされる場合にのみ実行する

   ```go
   if log.Enabled(context.Background(), slog.LevelDebug) {
       log.Debug("Detailed information",
           "data", generateExpensiveDebugData())
   }
   ```

3. **バッファリングとバッチ処理**
   - 高頻度のログイベントはバッファリングしてバッチ処理する
   - CloudWatch Logsへの送信は効率的なバッチで行う

4. **サンプリング**
   - 高ボリュームのログでは、サンプリングを検討する
   - 例えば、10件に1件のみログを記録するなど

5. **ログローテーションと圧縮**
   - ファイルへのログ出力時はログローテーションと圧縮を設定する
   - CloudWatch Logsでは適切なリテンション期間を設定する

以上の実装により、構造化ログ記録の基盤が整いました。次のセクションでは、これらのログをCloudWatch Logsに送信する方法について学びます。

### 1.5.8. 構造化ログ実装の動作確認

構造化ログを実装した後は、実際に期待通りに動作しているかを確認する必要があります。このセクションでは、実装した構造化ログの動作確認方法とログ出力の分析方法について説明します。

#### 1.5.8.1. アプリケーションの起動と基本ログの確認

まず、Docker Composeでアプリケーションを起動し、基本的なログ出力を確認します。

```bash
### アプリケーションを起動
task start

### バックエンドのログを表示
task logs:backend
```

起動時のログが構造化JSON形式で出力されていることを確認してください。以下のようなログが表示されるはずです：

```json
{"time":"2023-10-21T15:04:23Z","level":"INFO","msg":"Starting application","service":"aws-observability-ecommerce","environment":"development","version":"1.0.0"}
{"time":"2023-10-21T15:04:23Z","level":"INFO","msg":"Server starting","service":"aws-observability-ecommerce","environment":"development","version":"1.0.0","port":"8080"}
```

各ログエントリに、以下の情報が含まれていることを確認します：

- timestamp（`time`）
- ログレベル（`level`）
- メッセージ（`msg`）
- サービス名（`service`）
- 環境（`environment`）
- バージョン（`version`）

#### 1.5.8.2. APIリクエストによるログの生成と確認

次に、実際にAPIリクエストを送信して、リクエスト処理に関連するログが適切に記録されるかを確認します。

##### 1.5.8.2.1. ヘルスチェックAPIの動作確認

```bash
### ヘルスチェックAPIを呼び出す
curl -v http://api.localhost/api/health
```

ターミナルで`task logs:backend`を実行して出力されるログを確認します。以下のようなログが表示されるはずです：

```json
{"time":"2023-10-21T15:05:00Z","level":"DEBUG","msg":"Health check request processing","service":"aws-observability-ecommerce","environment":"development","version":"1.0.0","request_id":"550e8400-e29b-41d4-a716-446655440000","method":"GET","path":"/api/health","remote_ip":"172.17.0.1"}
{"time":"2023-10-21T15:05:00Z","level":"INFO","msg":"Health check completed","service":"aws-observability-ecommerce","environment":"development","version":"1.0.0","request_id":"550e8400-e29b-41d4-a716-446655440000","status":"ok","uptime":37000,"goroutines":10}
```

これらのログを確認し、以下のポイントが満たされているか検証します：

1. リクエストIDが自動生成され、すべてのログエントリに含まれているか
2. DEBUGレベルとINFOレベルのログが適切に出力されているか
3. リクエスト情報（メソッド、パス、IPアドレス）が含まれているか
4. 応答情報（ステータス、実行時間）が記録されているか

##### 1.5.8.2.2. 商品一覧APIの動作確認

```bash
### すべての商品を取得
curl -v http://api.localhost/api/products

### カテゴリーで絞り込み
curl -v http://api.localhost/api/products?category_id=1

### ページネーションの確認
curl -v http://api.localhost/api/products?page=2&page_size=3
```

ログを確認し、以下のポイントを検証します：

1. 各リクエストに固有のリクエストIDが割り当てられているか
2. リクエストパラメータがログに記録されているか
3. フィルタリングやページネーションの情報が適切にログに記録されているか
4. 処理結果の情報（アイテム数、ページ情報など）が含まれているか

#### 1.5.8.3. 異なるログレベルの確認

ログレベルが適切に機能しているかを確認します。`.env`ファイルまたは`.envrc`ファイルでログレベルを変更して、動作を確認します。

```bash
### .env または .envrc ファイルの例
LOG_LEVEL=debug  # すべてのログレベルを表示
### LOG_LEVEL=info   # INFO以上のログを表示
### LOG_LEVEL=warn   # WARNとERRORのみ表示
### LOG_LEVEL=error  # ERRORのみ表示
```

設定を変更した後、アプリケーションを再起動してログを確認します：

```bash
### アプリケーションを再起動
task restart:backend

### ログを確認
task logs:backend
```

各ログレベルの設定で、以下のようにログ出力が変化するか確認します：

- `debug`: すべてのログレベル（DEBUG, INFO, WARN, ERROR）が表示される
- `info`: INFO, WARN, ERRORレベルのログのみが表示される
- `warn`: WARN, ERRORレベルのログのみが表示される
- `error`: ERRORレベルのログのみが表示される

#### 1.5.8.4. エラーシナリオのテスト

エラー時のログ記録を確認するため、エラーを発生させるリクエストを送信します。

```bash
### 存在しないエンドポイントにアクセス
curl -v http://api.localhost/api/nonexistent

### 不正なパラメータでリクエスト
curl -v http://api.localhost/api/products?page=invalid
```

ログを確認し、以下のポイントを検証します：

1. WARNレベルまたはERRORレベルのログが出力されているか
2. エラーメッセージが含まれているか
3. ステータスコードが記録されているか
4. エラー発生場所やリクエスト情報が含まれているか

#### 1.5.8.5. リクエストとレスポンスの本文ログの確認

開発環境では、リクエスト本文とレスポンス本文のログ記録が有効になっているはずです。POSTリクエストを送信して確認します。

```bash
### POSTリクエストの送信
curl -v -X POST -H "Content-Type: application/json" -d '{"name":"テスト商品","price":1000}' http://api.localhost/api/products
```

ログには、以下のような情報が含まれているか確認します：

```json
{"time":"2023-10-21T15:06:00Z","level":"DEBUG","msg":"Request body","service":"aws-observability-ecommerce","environment":"development","version":"1.0.0","request_id":"7e9d5eb7-8b4f-4a0f-9c0e-7c9e9b6ac1f2","body":"{\"name\":\"テスト商品\",\"price\":1000}"}
```

リクエスト本文とレスポンス本文の両方が記録されているか、また、本文が適切に表示されているかを確認します。

#### 1.5.8.6. トラブルシューティング

構造化ログの実装中に発生する可能性のある問題と、その解決方法を紹介します。

##### 1.5.8.6.1. ログが出力されない

1. **ログレベルの確認**: 設定されているログレベルが適切かを確認します。たとえば `LOG_LEVEL=error` に設定されている場合、INFO や DEBUG レベルのログは表示されません。
2. **環境変数の読み込み確認**: 環境変数が正しく読み込まれているか確認します。

   ```bash
   # コンテナ内で環境変数を確認
   docker exec -it backend env | grep LOG
   ```

3. **標準出力の確認**: ログは標準出力に書き込まれるため、コンテナの標準出力にアクセスできることを確認します。

##### 1.5.8.6.2. JSON以外の形式でログが出力される

1. **LogFormatの確認**: `LOG_FORMAT` が `json` に設定されているか確認します。
2. **slogハンドラーの確認**: ロガーの初期化コードで `slog.NewJSONHandler` が使用されているか確認します。

##### 1.5.8.6.3. リクエストIDが含まれない

1. **ミドルウェアの有効化確認**: `RequestIDMiddleware` が適切に設定されているか確認します。
2. **ミドルウェアの順序**: ミドルウェアの適用順序が適切かを確認します。通常、RequestIDミドルウェアは他のログ関連ミドルウェアより先に適用する必要があります。

##### 1.5.8.6.4. エラー情報が不十分

1. **エラーハンドラーの実装確認**: グローバルエラーハンドラーが適切に実装されているか確認します。
2. **エラー処理方法**: 各ハンドラーでのエラー処理方法が適切かを確認します。エラーはそのまま返すのではなく、`echo.NewHTTPError` でラップするか、グローバルエラーハンドラーに処理を委任します。

#### 1.5.8.7. jqを使ったログ分析

`jq` コマンドラインツールを使うと、JSON形式のログを効果的に分析できます。以下に例を示します：

```bash
### すべてのログをJSON形式で整形して表示
task logs:backend | jq .

### ERRORレベルのログのみ表示
task logs:backend | jq 'select(.level=="ERROR")'

### 特定のリクエストIDのログのみ表示
task logs:backend | jq 'select(.request_id=="550e8400-e29b-41d4-a716-446655440000")'

### レスポンスタイムが100ms以上のリクエストのみ表示
task logs:backend | jq 'select(.elapsed_ms > 100)'

### 特定のパスへのリクエストのみ表示
task logs:backend | jq 'select(.path=="/api/products")'
```

#### 1.5.8.8. 本番環境を想定した設定のテスト

本番環境を想定したログ設定をテストするには、以下の手順を実施します：

1. 環境変数を本番設定に変更：

   ```bash
   # .env または .envrc ファイルで
   APP_ENV=production
   LOG_LEVEL=info  # 本番では通常 info または warn
   ```

2. アプリケーションを再起動して動作確認：

   ```bash
   task stop
   task start
   ```

3. 本番環境では以下の特性があるか確認：
   - DEBUGレベルのログが表示されないこと
   - リクエスト/レスポンスの本文が記録されないこと
   - エラー情報が適切に記録されること

#### 1.5.8.9. まとめ

構造化ログの実装後の動作確認では、以下のポイントが重要です：

1. **ログの構造**: ログがJSON形式で適切な構造になっているか
2. **ログレベル**: 設定したログレベルに応じて適切にログが出力されるか
3. **コンテキスト情報**: リクエストIDやユーザー情報などのコンテキスト情報が含まれているか
4. **エラー処理**: エラーが適切にログに記録されるか
5. **環境別設定**: 開発環境と本番環境で適切なログ設定になっているか

これらを確認することで、構造化ログが期待通りに機能していることを検証できます。次のセクションでは、これらのログをCloudWatch Logsに送信する方法について学びます。

## 1.6. バックエンドログのCloudWatch Logs連携

このセクションでは、Go言語のバックエンドアプリケーションのログをAWS CloudWatch Logsに送信する方法を学びます。LocalStackを使用して開発環境内でCloudWatch Logsをエミュレートし、本番環境と同様のログ管理機能を実現します。

### 1.6.1. AWS SDK for Go v2の設定

AWS SDK for Go v2は、AWSサービスにアクセスするためのモダンな公式クライアントライブラリです。AWS SDK for Go v1と比較して、パフォーマンスと使いやすさが向上しています。

#### 1.6.1.1. AWS SDK for Go v2の主な特徴

- **モジュール化**: 必要なサービスだけを依存関係に追加できる
- **コンテキストサポート**: すべての操作でコンテキストを使用可能
- **パガー**: ページ分割されたAPIレスポンスを簡単に処理できる
- **プレステータスAPIレスポンス**: APIレスポンスのHTTPステータスコードを簡単に確認できる
- **改善されたエラーハンドリング**: エラーチェックとリトライがより簡単

#### 1.6.1.2. インストール手順

AWS SDK for Go v2とCloudWatch Logsモジュールをインストールします：

```bash
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs
```

#### 1.6.1.3. SDK設定の基本

AWS SDK for Go v2をバックエンドアプリケーションに統合するための基本設定を行います。まず、`internal/aws`ディレクトリを作成し、AWS設定を管理するパッケージを実装します：

```bash
mkdir -p internal/aws
touch internal/aws/config.go
```

`internal/aws/config.go`ファイルを実装します：

```go
package aws

import (
 "context"
 "os"

 "github.com/aws/aws-sdk-go-v2/aws"
 "github.com/aws/aws-sdk-go-v2/config"
 "github.com/aws/aws-sdk-go-v2/credentials"
 "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

// AWSConfig はAWS設定を保持する構造体
type AWSConfig struct {
 Config         *aws.Config
 CloudWatchLogs *cloudwatchlogs.Client
}

// NewAWSConfig は新しいAWS設定を作成します
func NewAWSConfig(ctx context.Context) (*AWSConfig, error) {
 // LocalStackを使用するかどうかを環境変数から取得
 useLocalStack := os.Getenv("USE_LOCALSTACK") == "true"

 var cfg aws.Config
 var err error

 if useLocalStack {
  // LocalStack用の設定
  customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
   return aws.Endpoint{
    URL:           "http://localstack:4566", // Docker Compose内でのサービス名
    SigningRegion: region,
   }, nil
  })

  // LocalStack用の認証情報と設定
  cfg, err = config.LoadDefaultConfig(ctx,
   config.WithRegion("ap-northeast-1"),
   config.WithEndpointResolverWithOptions(customResolver),
   config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "test")),
  )
 } else {
  // 本番環境用の標準設定
  cfg, err = config.LoadDefaultConfig(ctx)
 }

 if err != nil {
  return nil, err
 }

 // CloudWatch Logsクライアントの作成
 cwLogsClient := cloudwatchlogs.NewFromConfig(cfg)

 return &AWSConfig{
  Config:         &cfg,
  CloudWatchLogs: cwLogsClient,
 }, nil
}
```

この設定クラスは、アプリケーションの実行環境に応じて適切なAWS設定を提供します：

- 開発環境では `USE_LOCALSTACK=true` 環境変数を設定し、LocalStackに接続
- 本番環境では通常のAWS認証フローを使用

### 1.6.2. CloudWatch Logs用のslogハンドラーの実装

先週実装した構造化ログ（slog）をCloudWatch Logsに送信するためのカスタムハンドラーを作成します。このハンドラーはslogのインターフェースを実装し、ログイベントをCloudWatch Logsに転送します。

#### 1.6.2.1. CloudWatch Logs用のslogハンドラー実装

まず、CloudWatch Logs用のslogハンドラーを実装するファイルを作成します：

```bash
mkdir -p internal/logger
touch internal/logger/cloudwatch_handler.go
```

`internal/logger/cloudwatch_handler.go`を以下のように実装します：

```go
package logger

import (
 "context"
 "encoding/json"
 "fmt"
 "log/slog"
 "sync"
 "time"

 "github.com/aws/aws-sdk-go-v2/aws"
 "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
 "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
 "github.com/google/uuid"
)

// CloudWatchHandler はCloudWatch Logsへのログ転送を行うslogハンドラー
type CloudWatchHandler struct {
 client        *cloudwatchlogs.Client
 logGroupName  string
 logStreamName string
 sequenceToken *string
 attrs         []slog.Attr
 level         slog.Level
 buffer        []types.InputLogEvent
 mu            sync.Mutex
 flushInterval time.Duration
 batchSize     int
 quit          chan struct{}
}

// CloudWatchHandlerOption はCloudWatchHandlerのオプション設定用関数型
type CloudWatchHandlerOption func(*CloudWatchHandler)

// WithLevel はハンドラーのログレベルを設定するオプション
func WithLevel(level slog.Level) CloudWatchHandlerOption {
 return func(h *CloudWatchHandler) {
  h.level = level
 }
}

// WithFlushInterval はログのフラッシュ間隔を設定するオプション
func WithFlushInterval(interval time.Duration) CloudWatchHandlerOption {
 return func(h *CloudWatchHandler) {
  h.flushInterval = interval
 }
}

// WithBatchSize はログのバッチサイズを設定するオプション
func WithBatchSize(size int) CloudWatchHandlerOption {
 return func(h *CloudWatchHandler) {
  h.batchSize = size
 }
}

// NewCloudWatchHandler は新しいCloudWatchHandlerを作成します
func NewCloudWatchHandler(client *cloudwatchlogs.Client, logGroupName string, opts ...CloudWatchHandlerOption) (*CloudWatchHandler, error) {
 // 一意のログストリーム名を生成
 logStreamName := fmt.Sprintf("app-%s", uuid.New().String())

 // デフォルト設定でハンドラーを初期化
 h := &CloudWatchHandler{
  client:        client,
  logGroupName:  logGroupName,
  logStreamName: logStreamName,
  level:         slog.LevelInfo, // デフォルトはINFOレベル
  buffer:        make([]types.InputLogEvent, 0, 100),
  flushInterval: 5 * time.Second, // デフォルトは5秒間隔
  batchSize:     100,             // デフォルトは最大100件
  quit:          make(chan struct{}),
 }

 // オプションを適用
 for _, opt := range opts {
  opt(h)
 }

 // ロググループの存在確認
 _, err := client.DescribeLogGroups(context.Background(), &cloudwatchlogs.DescribeLogGroupsInput{
  LogGroupNamePrefix: aws.String(logGroupName),
 })
 if err != nil {
  // ロググループが存在しない場合は作成
  _, err = client.CreateLogGroup(context.Background(), &cloudwatchlogs.CreateLogGroupInput{
   LogGroupName: aws.String(logGroupName),
  })
  if err != nil {
   return nil, fmt.Errorf("failed to create log group: %w", err)
  }
 }

 // ログストリームの作成
 _, err = client.CreateLogStream(context.Background(), &cloudwatchlogs.CreateLogStreamInput{
  LogGroupName:  aws.String(logGroupName),
  LogStreamName: aws.String(logStreamName),
 })
 if err != nil {
  return nil, fmt.Errorf("failed to create log stream: %w", err)
 }

 // 定期的なログフラッシュを行うゴルーチンを開始
 go h.flushPeriodically()

 return h, nil
}

// flushPeriodically は定期的にバッファされたログをCloudWatch Logsに送信します
func (h *CloudWatchHandler) flushPeriodically() {
 ticker := time.NewTicker(h.flushInterval)
 defer ticker.Stop()

 for {
  select {
  case <-ticker.C:
   if err := h.Flush(); err != nil {
    // エラーは標準エラー出力に記録
    fmt.Printf("Error flushing logs to CloudWatch: %v\n", err)
   }
  case <-h.quit:
   return
  }
 }
}

// Close はハンドラーを閉じ、残りのログをフラッシュします
func (h *CloudWatchHandler) Close() error {
 // 終了シグナルを送信
 close(h.quit)

 // 残りのログをフラッシュ
 return h.Flush()
}

// Flush はバッファされたログをCloudWatch Logsに送信します
func (h *CloudWatchHandler) Flush() error {
 h.mu.Lock()
 defer h.mu.Unlock()

 if len(h.buffer) == 0 {
  return nil
 }

 // CloudWatch Logsへの送信準備
 input := &cloudwatchlogs.PutLogEventsInput{
  LogGroupName:  aws.String(h.logGroupName),
  LogStreamName: aws.String(h.logStreamName),
  LogEvents:     h.buffer,
 }

 // シーケンストークンがある場合は設定
 if h.sequenceToken != nil {
  input.SequenceToken = h.sequenceToken
 }

 // ログを送信
 resp, err := h.client.PutLogEvents(context.Background(), input)
 if err != nil {
  return fmt.Errorf("failed to put log events: %w", err)
 }

 // 次回のリクエスト用にシーケンストークンを更新
 h.sequenceToken = resp.NextSequenceToken

 // バッファをクリア
 h.buffer = h.buffer[:0]

 return nil
}

// Enabled はログレベルが設定されたレベル以上かどうかを返します
func (h *CloudWatchHandler) Enabled(ctx context.Context, level slog.Level) bool {
 return level >= h.level
}

// Handle はログレコードを処理します
func (h *CloudWatchHandler) Handle(ctx context.Context, record slog.Record) error {
 // ログレベルが設定レベル未満の場合は何もしない
 if !h.Enabled(ctx, record.Level) {
  return nil
 }

 // ログレコードを構造体に変換
 logMap := make(map[string]interface{})
 logMap["time"] = record.Time.Format(time.RFC3339)
 logMap["level"] = record.Level.String()
 logMap["message"] = record.Message

 // 属性を追加
 record.Attrs(func(attr slog.Attr) bool {
  addAttr(logMap, "", attr)
  return true
 })

 // グローバル属性を追加
 for _, attr := range h.attrs {
  addAttr(logMap, "", attr)
 }

 // ログをJSON形式にシリアライズ
 jsonData, err := json.Marshal(logMap)
 if err != nil {
  return err
 }

 // ログイベントを作成
 logEvent := types.InputLogEvent{
  Message:   aws.String(string(jsonData)),
  Timestamp: aws.Int64(record.Time.UnixMilli()),
 }

 // ログイベントをバッファに追加
 h.mu.Lock()
 h.buffer = append(h.buffer, logEvent)
 bufferLen := len(h.buffer)
 h.mu.Unlock()

 // バッファサイズが閾値を超えたらフラッシュ
 if bufferLen >= h.batchSize {
  return h.Flush()
 }

 return nil
}

// WithAttrs は属性を持つ新しいハンドラーを返します
func (h *CloudWatchHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
 h2 := *h
 h2.attrs = append(h2.attrs[:], attrs...)
 return &h2
}

// WithGroup はグループ化された属性を持つ新しいハンドラーを返します
func (h *CloudWatchHandler) WithGroup(name string) slog.Handler {
 h2 := *h
 if name == "" {
  return &h2
 }
 // 属性をグループ化
 attrs := make([]slog.Attr, 0, len(h.attrs))
 for _, attr := range h.attrs {
  if attr.Key != "" {
   attr.Key = name + "." + attr.Key
  }
  attrs = append(attrs, attr)
 }
 h2.attrs = attrs
 return &h2
}

// addAttr はログマップに属性を追加するヘルパー関数
func addAttr(m map[string]interface{}, prefix string, attr slog.Attr) {
 key := attr.Key
 if prefix != "" {
  key = prefix + "." + key
 }

 switch attr.Value.Kind() {
 case slog.KindBool, slog.KindInt64, slog.KindUint64, slog.KindFloat64, slog.KindString:
  m[key] = attr.Value.Any()
 case slog.KindTime:
  m[key] = attr.Value.Time().Format(time.RFC3339)
 case slog.KindDuration:
  m[key] = attr.Value.Duration().String()
 case slog.KindGroup:
  for _, a := range attr.Value.Group() {
   addAttr(m, key, a)
  }
 case slog.KindLogValuer:
  addAttr(m, prefix, slog.Attr{
   Key:   attr.Key,
   Value: attr.Value.LogValuer().LogValue(),
  })
 default:
  m[key] = fmt.Sprintf("%v", attr.Value)
 }
}
```

#### 1.6.2.2. ロガーの初期化処理実装

次に、アプリケーション起動時にCloudWatch Logsハンドラーを初期化し、標準ロガーとして設定する処理を実装します。`internal/logger/logger.go`を以下のように実装します：

```go
package logger

import (
 "context"
 "io"
 "log/slog"
 "os"
 "time"

 "github.com/y-nosuke/aws-observability-ecommerce/internal/aws"
 appconfig "github.com/y-nosuke/aws-observability-ecommerce/internal/config"
)

// Logger はアプリケーション全体で使用するロガー構造体
type Logger struct {
 slogger        *slog.Logger
 cloudWatchHandler *CloudWatchHandler
}

// Config はロガーの設定
type Config struct {
 AppName     string
 Environment string
 LogLevel    slog.Level
 UseConsole  bool
 UseFile     bool
 FilePath    string
 UseCloudWatch bool
 LogGroupName  string
}

// DefaultConfig はデフォルトのロガー設定を返します
func DefaultConfig() Config {
 return Config{
  AppName:     appconfig.Config.App.Name,
  Environment: appconfig.Config.App.Environment,
  LogLevel:    slog.LevelInfo,
  UseConsole:  true,
  UseFile:     false,
  FilePath:    "logs/app.log",
  UseCloudWatch: true,
  LogGroupName:  "/aws-observability-ecommerce/dev/backend",
 }
}

// New は新しいロガーインスタンスを作成します
func New(ctx context.Context, config Config) (*Logger, error) {
 var handlers []slog.Handler
 var closers []io.Closer

 // コンソール出力ハンドラー
 if config.UseConsole {
  textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
   Level: config.LogLevel,
  })
  handlers = append(handlers, textHandler)
 }

 // ファイル出力ハンドラー
 if config.UseFile {
  file, err := os.OpenFile(config.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
   return nil, err
  }
  jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
   Level: config.LogLevel,
  })
  handlers = append(handlers, jsonHandler)
  closers = append(closers, file)
 }

 // CloudWatch Logsハンドラー
 var cloudWatchHandler *CloudWatchHandler
 if config.UseCloudWatch {
  // AWS設定の取得
  awsConfig, err := aws.NewAWSConfig(ctx)
  if err != nil {
   return nil, err
  }

  // CloudWatch Logsハンドラーの作成
  cwHandler, err := NewCloudWatchHandler(
   awsConfig.CloudWatchLogs,
   config.LogGroupName,
   WithLevel(config.LogLevel),
   WithFlushInterval(5*time.Second),
   WithBatchSize(100),
  )
  if err != nil {
   return nil, err
  }

  handlers = append(handlers, cwHandler)
  closers = append(closers, cwHandler)
  cloudWatchHandler = cwHandler
 }

 // マルチハンドラーの作成
 var mainHandler slog.Handler
 if len(handlers) == 1 {
  mainHandler = handlers[0]
 } else {
  mainHandler = newMultiHandler(handlers...)
 }

 // 共通属性の追加
 mainHandler = mainHandler.WithAttrs([]slog.Attr{
  slog.String("app", config.AppName),
  slog.String("env", config.Environment),
 })

 // ロガーの作成
 slogger := slog.New(mainHandler)

 // グローバルロガーとして設定
 slog.SetDefault(slogger)

 return &Logger{
  slogger:        slogger,
  cloudWatchHandler: cloudWatchHandler,
 }, nil
}

// Logger はslog.Loggerインスタンスを返します
func (l *Logger) Logger() *slog.Logger {
 return l.slogger
}

// Close はロガーリソースを解放します
func (l *Logger) Close() error {
 if l.cloudWatchHandler != nil {
  return l.cloudWatchHandler.Close()
 }
 return nil
}

// multiHandler は複数のハンドラーに同時にログを送信するslogハンドラー
type multiHandler struct {
 handlers []slog.Handler
}

// newMultiHandler は新しいマルチハンドラーを作成します
func newMultiHandler(handlers ...slog.Handler) *multiHandler {
 return &multiHandler{handlers: handlers}
}

// Enabled は少なくとも1つのハンドラーが有効であるかどうかを返します
func (h *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
 for _, handler := range h.handlers {
  if handler.Enabled(ctx, level) {
   return true
  }
 }
 return false
}

// Handle は全てのハンドラーでログレコードを処理します
func (h *multiHandler) Handle(ctx context.Context, record slog.Record) error {
 for _, handler := range h.handlers {
  if err := handler.Handle(ctx, record); err != nil {
   return err
  }
 }
 return nil
}

// WithAttrs は属性を持つ新しいマルチハンドラーを返します
func (h *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
 handlers := make([]slog.Handler, len(h.handlers))
 for i, handler := range h.handlers {
  handlers[i] = handler.WithAttrs(attrs)
 }
 return newMultiHandler(handlers...)
}

// WithGroup はグループ化された属性を持つ新しいマルチハンドラーを返します
func (h *multiHandler) WithGroup(name string) slog.Handler {
 handlers := make([]slog.Handler, len(h.handlers))
 for i, handler := range h.handlers {
  handlers[i] = handler.WithGroup(name)
 }
 return newMultiHandler(handlers...)
}
```

### 1.6.3. バッチ処理とエラーハンドリング

CloudWatch Logsへのログ送信は、パフォーマンスとコスト効率の観点から最適化が必要です。前節で実装したCloudWatch Logsハンドラーには、以下のようなバッチ処理とエラーハンドリングの機能が組み込まれています：

#### 1.6.3.1. バッチ処理の仕組み

1. **ログバッファリング**: ログイベントはメモリ内のバッファに蓄積されます。
2. **定期的なフラッシュ**: 設定された間隔（デフォルト5秒）でバッファのログを送信します。
3. **サイズベースのフラッシュ**: バッファサイズが閾値（デフォルト100件）を超えた場合も送信します。
4. **シーケンストークン管理**: CloudWatch Logs API要求には正しいシーケンストークンが必要です。

#### 1.6.3.2. エラーハンドリングの方法

CloudWatch Logsへのログ送信中に発生する可能性のあるエラーを適切に処理することが重要です：

1. **接続エラー**: ネットワーク接続の問題やタイムアウトが発生した場合
2. **認証エラー**: AWS認証情報が無効または期限切れの場合
3. **リソース制限**: APIリクエスト制限やクォータ超過の場合
4. **シーケンストークンエラー**: 無効なシーケンストークンでリクエストした場合

これらのエラーに対処するために、以下の戦略を実装しています：

1. **エラーのログ記録**: ログ送信エラーは標準エラー出力にフォールバックして記録
2. **グレースフルシャットダウン**: アプリケーション終了時に残りのログをフラッシュ
3. **バッファ管理**: エラー発生時にもバッファを適切に管理して、メモリリークを防止

エラー発生時のフォールバック処理を追加するには、`cloudwatch_handler.go`の`flushPeriodically`メソッドを以下のように拡張できます：

```go
// flushPeriodically は定期的にバッファされたログをCloudWatch Logsに送信します
func (h *CloudWatchHandler) flushPeriodically() {
 ticker := time.NewTicker(h.flushInterval)
 defer ticker.Stop()

 retryCount := 0
 maxRetries := 3

 for {
  select {
  case <-ticker.C:
   err := h.Flush()
   if err != nil {
    // エラーは標準エラー出力に記録
    fmt.Printf("Error flushing logs to CloudWatch: %v\n", err)

    // リトライ処理
    if retryCount < maxRetries {
     retryCount++
     // 指数バックオフでリトライ
     retryTime := time.Duration(1<<retryCount) * time.Second
     fmt.Printf("Will retry in %v (attempt %d/%d)\n", retryTime, retryCount, maxRetries)
     time.Sleep(retryTime)
     continue
    }

    // リトライ失敗後はローカルファイルにフォールバック
    h.fallbackToFile(err)
    retryCount = 0
   } else {
    retryCount = 0
   }
  case <-h.quit:
   return
  }
 }
}

// fallbackToFile はCloudWatch Logsへの送信に失敗した場合にログをファイルに保存します
func (h *CloudWatchHandler) fallbackToFile(originalErr error) {
 h.mu.Lock()
 defer h.mu.Unlock()

 if len(h.buffer) == 0 {
  return
 }

 // フォールバック用のファイル名
 fallbackFile := fmt.Sprintf("logs/cloudwatch_fallback_%s.log", time.Now().Format("20060102_150405"))

 // ディレクトリが存在することを確認
 os.MkdirAll("logs", 0755)

 // ファイルを開く
 file, err := os.OpenFile(fallbackFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
 if err != nil {
  fmt.Printf("Failed to open fallback log file: %v\n", err)
  return
 }
 defer file.Close()

 // ヘッダー情報を書き込み
 file.WriteString(fmt.Sprintf("# CloudWatch Logs fallback - %s\n", time.Now().Format(time.RFC3339)))
 file.WriteString(fmt.Sprintf("# Original error: %v\n", originalErr))
 file.WriteString(fmt.Sprintf("# Log group: %s, Log stream: %s\n", h.logGroupName, h.logStreamName))
 file.WriteString("---\n")

 // バッファ内のログイベントをファイルに書き込み
 for _, event := range h.buffer {
  timestamp := time.UnixMilli(*event.Timestamp).Format(time.RFC3339)
  file.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, *event.Message))
 }

 fmt.Printf("Logs saved to fallback file: %s\n", fallbackFile)

 // バッファをクリア
 h.buffer = h.buffer[:0]
}
```

### 1.6.4. LocalStackへのログ転送設定と検証

LocalStackを使用してCloudWatch Logsのエミュレーションを行い、開発環境でログ転送の検証を行います。

#### 1.6.4.1. メインアプリケーションへのLogger統合

アプリケーションのメインファイル（`cmd/api/main.go`）を更新して、ロガーの初期化と利用を組み込みます：

```go
package main

import (
 "context"
 "log"
 "os"
 "os/signal"
 "syscall"
 "time"

 "github.com/labstack/echo/v4"
 "github.com/labstack/echo/v4/middleware"

 "github.com/y-nosuke/aws-observability-ecommerce/internal/api/handlers"
 "github.com/y-nosuke/aws-observability-ecommerce/internal/api/router"
 "github.com/y-nosuke/aws-observability-ecommerce/internal/config"
 "github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

func main() {
 // コンテキストの初期化（シグナルハンドリング）
 ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
 defer stop()

 // ロガーの初期化
 loggerConfig := logger.DefaultConfig()
 appLogger, err := logger.New(ctx, loggerConfig)
 if err != nil {
  log.Fatalf("Failed to initialize logger: %v", err)
 }
 defer appLogger.Close()

 slogger := appLogger.Logger()

 // アプリケーションの起動をログに記録
 slogger.Info("Starting application",
  "version", config.Config.App.Version,
  "environment", config.Config.App.Environment)

 // Echoインスタンスの作成
 e := echo.New()
 e.HideBanner = true
 e.HidePort = true

 // ミドルウェアの設定
 e.Use(middleware.Recover())
 e.Use(middleware.CORS())

 // カスタムロギングミドルウェアの設定
 e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
   req := c.Request()
   res := c.Response()
   start := time.Now()

   // リクエスト情報をログに記録
   slogger.Info("Request started",
    "method", req.Method,
    "path", req.URL.Path,
    "remote_ip", c.RealIP(),
    "user_agent", req.UserAgent())

   // 次のハンドラーを実行
   err := next(c)

   // レスポンス情報をログに記録
   latency := time.Since(start)
   fields := []interface{}{
    "method", req.Method,
    "path", req.URL.Path,
    "status", res.Status,
    "latency_ms", latency.Milliseconds(),
    "bytes_out", res.Size,
   }

   if err != nil {
    fields = append(fields, "error", err.Error())
    slogger.Error("Request failed", fields...)
   } else {
    slogger.Info("Request completed", fields...)
   }

   return err
  }
 })

 // ハンドラーの作成
 healthHandler := handlers.NewHealthHandler()
 productHandler := handlers.NewProductHandler()

 // ルーターの設定
 router.SetupRoutes(e, healthHandler, productHandler)

 // サーバーの起動（非同期）
 go func() {
  address := ":" + config.Config.Server.Port
  slogger.Info("Server starting", "address", address)
  if err := e.Start(address); err != nil {
   slogger.Error("Server shutdown unexpectedly", "error", err)
  }
 }()

 // シグナルを待機
 <-ctx.Done()
 slogger.Info("Shutdown signal received, gracefully shutting down...")

 // グレースフルシャットダウン
 shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 defer cancel()

 if err := e.Shutdown(shutdownCtx); err != nil {
  slogger.Error("Server shutdown failed", "error", err)
 }

 slogger.Info("Server has been shutdown gracefully")
}
```

#### 1.6.4.2. 環境変数設定

アプリケーションの環境変数に、LocalStackとCloudWatch Logsの設定を追加します。
`backend/.envrc`ファイルを更新します：

```bash
## アプリケーション設定
export APP_NAME=aws-observability-ecommerce
export APP_VERSION=1.0.0
export APP_ENV=development
export PORT=8080

## AWS設定
export USE_LOCALSTACK=true
export AWS_REGION=ap-northeast-1
export AWS_ENDPOINT=http://localstack:4566
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test

## ロギング設定
export LOG_LEVEL=debug
export LOG_FORMAT=json
export USE_CLOUDWATCH=true
export LOG_GROUP_NAME=aws-observability-ecommerce
```

Docker Composeファイル（`compose.yml`）のバックエンドサービス設定にも環境変数を追加します：

```yaml
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile.dev
    restart: always
    volumes:
      - ./backend:/app
    expose:
      - "8080"
    environment:
      - APP_NAME=aws-observability-ecommerce
      - APP_VERSION=1.0.0
      - APP_ENV=development
      - PORT=8080
      - USE_LOCALSTACK=true
      - AWS_REGION=ap-northeast-1
      - AWS_ACCESS_KEY_ID=test
      - AWS_SECRET_ACCESS_KEY=test
      - LOG_LEVEL=debug
      - USE_CLOUDWATCH=true
      - LOG_GROUP_NAME=/aws-observability-ecommerce/dev/backend
    depends_on:
      mysql:
        condition: service_healthy
      localstack:
        condition: service_healthy
    # ... 以下は変更なし
```

#### 1.6.4.3. ログ転送の検証方法

アプリケーションとLocalStackを起動し、ログ転送が正常に機能しているか検証します：

1. Docker Compose環境を起動します：

    ```bash
    task start
    ```

2. APIリクエストを送信してログを生成します：

    ```bash
    curl http://api.localhost/api/health
    curl http://api.localhost/api/products
    ```

3. LocalStackのCloudWatch Logsにログが転送されているか確認します：

    ```bash
    ## ロググループの一覧を表示
    task exec:localstack -- awslocal logs describe-log-groups

    ## ログストリームの一覧を表示
    task exec:localstack -- awslocal logs describe-log-streams --log-group-name /aws-observability-ecommerce/dev/backend

    ## ログイベントを取得（ログストリーム名は実際の値に置き換え）
    task exec:localstack -- awslocal logs get-log-events --log-group-name /aws-observability-ecommerce/dev/backend --log-stream-name app-xxxx-xxxx-xxxx-xxxx
    ```

4. ログの内容を確認し、フォーマットと属性が正しいことを確認します。

### 1.6.5. 非同期ログ転送と性能最適化

CloudWatch Logsへのログ転送はアプリケーションのパフォーマンスに影響を与える可能性があります。ここでは、パフォーマンスを最適化するための追加の戦略について説明します。

#### 1.6.5.1. 非同期処理の重要性

CloudWatch Logsへのログ送信は非同期で行われるべきです。先の実装では、以下の2つの非同期メカニズムが導入されています：

1. **バッファリング**: ログイベントをメモリ内に一時的に蓄積
2. **バックグラウンド送信**: 定期的なフラッシュとバッチ処理

#### 1.6.5.2. さらなる最適化戦略

高負荷環境でのパフォーマンスをさらに向上させるための追加の最適化戦略を紹介します：

##### 1.6.5.2.1. ロググループとログストリームの効率的な管理

ロググループとログストリームの設計は、CloudWatch Logsのパフォーマンスとコストに大きく影響します：

```go
// CreateLogGroups はロググループが存在しない場合に作成します
func CreateLogGroups(ctx context.Context, client *cloudwatchlogs.Client, logGroupNames []string) error {
 for _, name := range logGroupNames {
  // ロググループの存在確認
  groups, err := client.DescribeLogGroups(ctx, &cloudwatchlogs.DescribeLogGroupsInput{
   LogGroupNamePrefix: aws.String(name),
  })
  if err != nil {
   return fmt.Errorf("failed to describe log groups: %w", err)
  }

  // 同名のロググループが存在するかチェック
  exists := false
  for _, group := range groups.LogGroups {
   if *group.LogGroupName == name {
    exists = true
    break
   }
  }

  // 存在しない場合は作成
  if !exists {
   _, err = client.CreateLogGroup(ctx, &cloudwatchlogs.CreateLogGroupInput{
    LogGroupName: aws.String(name),
   })
   if err != nil {
    return fmt.Errorf("failed to create log group %s: %w", name, err)
   }

   // 保持期間を設定（30日）
   _, err = client.PutRetentionPolicy(ctx, &cloudwatchlogs.PutRetentionPolicyInput{
    LogGroupName:    aws.String(name),
    RetentionInDays: aws.Int32(30),
   })
   if err != nil {
    return fmt.Errorf("failed to set retention policy for %s: %w", name, err)
   }
  }
 }
 return nil
}
```

##### 1.6.5.2.2. ログの重要度に基づいたサンプリング

全てのログを送信するのではなく、重要度や特定条件に基づいてサンプリングすることで、コストと性能を最適化できます：

```go
// SamplingHandler はログレベルに基づいてサンプリングするハンドラー
type SamplingHandler struct {
 handler      slog.Handler
 infoSampleRate  float64 // INFOレベルのサンプリング率（0.0～1.0）
 debugSampleRate float64 // DEBUGレベルのサンプリング率（0.0～1.0）
}

// NewSamplingHandler は新しいサンプリングハンドラーを作成します
func NewSamplingHandler(handler slog.Handler, infoRate, debugRate float64) *SamplingHandler {
 return &SamplingHandler{
  handler:      handler,
  infoSampleRate:  infoRate,
  debugSampleRate: debugRate,
 }
}

// Enabled はログレベルとサンプリング確率に基づいて有効かどうかを返します
func (h *SamplingHandler) Enabled(ctx context.Context, level slog.Level) bool {
 // ERRORとFATALは常に記録
 if level >= slog.LevelError {
  return true
 }

 // WARNは常に記録
 if level >= slog.LevelWarn {
  return true
 }

 // INFOはサンプリング率に基づいて記録
 if level >= slog.LevelInfo {
  return rand.Float64() < h.infoSampleRate
 }

 // DEBUGはサンプリング率に基づいて記録
 return rand.Float64() < h.debugSampleRate
}

// Handle はサンプリングに基づいてログレコードを処理します
func (h *SamplingHandler) Handle(ctx context.Context, record slog.Record) error {
 if !h.Enabled(ctx, record.Level) {
  return nil
 }
 return h.handler.Handle(ctx, record)
}

// WithAttrs は属性を持つ新しいハンドラーを返します
func (h *SamplingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
 return &SamplingHandler{
  handler:      h.handler.WithAttrs(attrs),
  infoSampleRate:  h.infoSampleRate,
  debugSampleRate: h.debugSampleRate,
 }
}

// WithGroup はグループ化された属性を持つ新しいハンドラーを返します
func (h *SamplingHandler) WithGroup(name string) slog.Handler {
 return &SamplingHandler{
  handler:      h.handler.WithGroup(name),
  infoSampleRate:  h.infoSampleRate,
  debugSampleRate: h.debugSampleRate,
 }
}
```

##### 1.6.5.2.3. ネットワーク再試行とバックオフ戦略

CloudWatch Logsへのログ送信は、ネットワークの問題やAPI制限により失敗することがあります。以下のような再試行ロジックを実装することで信頼性を向上できます：

```go
// retryWithBackoff は指定された関数を再試行します
func retryWithBackoff(ctx context.Context, operation func() error, maxRetries int) error {
 var err error
 base := 100 * time.Millisecond
 cap := 5 * time.Second

 for i := 0; i < maxRetries; i++ {
  err = operation()
  if err == nil {
   return nil
  }

  // 再試行すべきエラーかどうかを確認
  if !isRetryableError(err) {
   return err
  }

  // 次の再試行までの待機時間を計算（指数バックオフ）
  waitTime := exponentialBackoff(i, base, cap)

  // コンテキストのキャンセル確認またはタイムアウト
  select {
  case <-ctx.Done():
   return ctx.Err()
  case <-time.After(waitTime):
   // 次の再試行へ進む
  }
 }

 return fmt.Errorf("operation failed after %d retries: %w", maxRetries, err)
}

// exponentialBackoff は指数バックオフ時間を計算します
func exponentialBackoff(attempt int, base, cap time.Duration) time.Duration {
 // 2^attempt * base + jitter
 jitter := time.Duration(rand.Int63n(int64(base)))
 waitTime := (1 << uint(attempt)) * base + jitter

 if waitTime > cap {
  return cap
 }
 return waitTime
}

// isRetryableError はエラーが再試行可能かどうかを判断します
func isRetryableError(err error) bool {
 // AWS SDKのエラータイプをチェック
 var apiErr smithy.APIError
 if errors.As(err, &apiErr) {
  // スロットリングエラーや一時的なエラーは再試行
  switch apiErr.ErrorCode() {
  case "ThrottlingException", "ServiceUnavailable", "InternalFailure":
   return true
  }
 }

 // ネットワーク関連のエラーは再試行
 var opErr *net.OpError
 var connErr *ConnectError
 return errors.As(err, &opErr) || errors.As(err, &connErr)
}
```

#### 1.6.5.3. パフォーマンスベンチマークの実施

ロギング実装のパフォーマンスを評価するためのベンチマークテストを作成します：

```go
// LoggingBenchmark はロギング実装のパフォーマンスをベンチマークします
func LoggingBenchmark(b *testing.B, handler slog.Handler) {
 logger := slog.New(handler)

 b.ResetTimer()
 for i := 0; i < b.N; i++ {
  logger.Info("This is a benchmark log message",
   "iteration", i,
   "timestamp", time.Now(),
   "benchmark", "logging_performance",
  )
 }
}

// ConsoleLoggerBenchmark はコンソールロガーのベンチマーク
func BenchmarkConsoleLogger(b *testing.B) {
 handler := slog.NewTextHandler(io.Discard, nil)
 LoggingBenchmark(b, handler)
}

// CloudWatchLoggerBenchmark はCloudWatch Logsロガーのベンチマーク
func BenchmarkCloudWatchLogger(b *testing.B) {
 // ベンチマーク用にモックハンドラーを使用
 mockHandler := NewMockCloudWatchHandler()
 LoggingBenchmark(b, mockHandler)
}

// AsyncCloudWatchLoggerBenchmark は非同期CloudWatch Logsロガーのベンチマーク
func BenchmarkAsyncCloudWatchLogger(b *testing.B) {
 // ベンチマーク用にモックハンドラーを使用
 mockHandler := NewMockAsyncCloudWatchHandler()
 LoggingBenchmark(b, mockHandler)
}
```

ベンチマーク結果を分析し、アプリケーションの要件に最適なロギング設定を特定します。本番環境では、以下の指標を監視することをお勧めします：

- **ログ処理のレイテンシ**: ログ記録による処理遅延
- **メモリ使用量**: ログバッファによるメモリ消費
- **CloudWatch Logsのコスト**: 転送されたログデータ量と料金の関係
- **APIリクエスト数**: CloudWatch Logs APIの呼び出し頻度

これらの指標に基づいて、バッファサイズ、フラッシュ間隔、サンプリング率などのパラメータを調整します。

#### 1.6.5.4. ロギング実装のテスト

最後に、CloudWatch Logs連携の実装を単体テストで検証することが重要です。テストでは実際のAWSリソースではなく、モックを使用します：

```go
func TestCloudWatchHandler(t *testing.T) {
 // モックのCloudWatch Logsクライアントを作成
 mockClient := NewMockCloudWatchLogsClient()

 // テスト用のハンドラーを作成
 handler, err := NewCloudWatchHandler(
  mockClient,
  "/test/log-group",
  WithLevel(slog.LevelInfo),
  WithFlushInterval(100*time.Millisecond),
  WithBatchSize(5),
 )
 require.NoError(t, err)
 defer handler.Close()

 // テストロガーの作成
 logger := slog.New(handler)

 // ログメッセージの送信
 logger.Info("Test log message", "test", true)
 logger.Warn("Test warning message", "test", true)
 logger.Error("Test error message", "test", true, "error", fmt.Errorf("test error"))

 // ログがフラッシュされるのを待つ
 time.Sleep(200 * time.Millisecond)

 // モッククライアントが期待通りのログイベントを受信したことを検証
 events := mockClient.GetReceivedEvents()
 assert.Equal(t, 3, len(events))

 // ログメッセージの内容を検証
 assert.Contains(t, *events[0].Message, "Test log message")
 assert.Contains(t, *events[1].Message, "Test warning message")
 assert.Contains(t, *events[2].Message, "Test error message")
}
```

まとめると、CloudWatch Logsを使用した効率的なログ転送の実装には、以下の要素が重要です：

1. **非同期処理**: バックグラウンドでのバッファリングとフラッシュ
2. **エラーハンドリング**: 接続問題やAPIエラーの適切な処理
3. **最適化**: サンプリング、バッチサイズ、フラッシュ間隔の調整
4. **フォールバック**: 障害時のローカルファイルへのログ保存
5. **テスト**: モックを使用した動作の検証

これらの原則に従うことで、アプリケーションのパフォーマンスを維持しながら、信頼性の高いロギングシステムを構築できます。

次のセクションでは、フロントエンドアプリケーションのログ収集とその設計について学びます。
