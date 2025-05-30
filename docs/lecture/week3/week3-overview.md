# 1. Week 3: 商品カタログバックエンドの完成

## 1.1. 目次

- [1. Week 3: 商品カタログバックエンドの完成](#1-week-3-商品カタログバックエンドの完成)
  - [1.1. 目次](#11-目次)
  - [1.2. 学習目標](#12-学習目標)
    - [1.2.1. 主要目標](#121-主要目標)
    - [1.2.2. 完了時の成果物](#122-完了時の成果物)
    - [1.2.3. 前提知識](#123-前提知識)
  - [1.3. 各日の概要](#13-各日の概要)
    - [1.3.1. Day 1: 商品一覧APIの実装](#131-day-1-商品一覧apiの実装)
    - [1.3.2. Day 2: 商品詳細・カテゴリAPI実装](#132-day-2-商品詳細カテゴリapi実装)
    - [1.3.3. Day 3: Lambda関数とS3連携の実装](#133-day-3-lambda関数とs3連携の実装)
    - [1.3.4. Day 4: 高度なバリデーション実装とAPI品質向上](#134-day-4-高度なバリデーション実装とapi品質向上)
    - [1.3.5. Day 5: テスト駆動開発と単体テスト](#135-day-5-テスト駆動開発と単体テスト)
  - [1.4. 使用技術と環境要件](#14-使用技術と環境要件)
  - [1.5. 事前準備事項](#15-事前準備事項)
  - [1.6. 学習のポイント](#16-学習のポイント)
  - [1.7. 関連リソース](#17-関連リソース)

## 1.2. 学習目標

### 1.2.1. 主要目標

- RESTful APIの設計と実装パターンをマスターする（ページネーション、フィルタリング、エラーレスポンス）
- バリデーションとエラーハンドリングの体系的な実装方法を理解する
- テスト駆動開発（TDD）の基本原則と実践的な適用方法を習得する
- AWSサーバーレス関数（Lambda）の基本概念と実装パターンを学ぶ
- S3を活用した画像保存・取得のワークフローを実装できるようになる

### 1.2.2. 完了時の成果物

- **商品カタログAPI一式**
  - 商品一覧API（ページネーション対応）
  - 商品詳細API
  - カテゴリー別商品一覧API
  - 適切なレスポンス形式とステータスコード

- **エラーハンドリング実装**
  - グローバルエラーハンドラー
  - 構造化されたエラーレスポンス形式
  - 適切なHTTPステータスコードのマッピング

- **テストコード**
  - ユニットテスト
  - テーブル駆動テストの実装
  - 70%以上のコードカバレッジ

- **サーバーレス実装**
  - 商品画像リサイズLambda関数
  - S3バケット連携機能
  - 画像アップロード・取得フロー

### 1.2.3. 前提知識

- Go言語の基本構文と標準ライブラリの使用経験
- Echo Webフレームワークの基本的な理解（ミドルウェア、ハンドラー）
- sqlboilerによるデータベース操作の基礎知識（Week 2で学習済み）
- RESTful APIの基本概念（リソース、HTTPメソッド、ステータスコード）
- AWS基本概念（IAM、リージョン、サービス連携）

## 1.3. 各日の概要

### 1.3.1. Day 1: 商品一覧APIの実装

ページネーション対応の商品一覧APIを設計・実装します。クエリパラメータによるフィルタリング機能を追加し、レスポンス形式の標準化を行います。sqlboilerを活用したデータアクセスレイヤーの実装とAPI実装パターンを習得します。

### 1.3.2. Day 2: 商品詳細・カテゴリAPI実装

商品詳細取得APIとカテゴリー別商品一覧APIを実装します。URLパラメータを使用したリソース取得、関連データの取得方法、OpenAPI仕様の詳細化を学びます。レスポンス形式の一貫性と統一性を確保する実装パターンを習得します。

### 1.3.3. Day 3: Lambda関数とS3連携の実装

AWSのサーバーレス機能を活用した商品画像リサイズ機能を実装します。AWS Lambda関数の基本構造を学び、S3バケットとの連携方法を実装します。LocalStackを使ってローカル環境でのテストも行います。

### 1.3.4. Day 4: 高度なバリデーション実装とAPI品質向上

oapi-codgenを使ったOpenAPI仕様からのコード生成、go-playground/validatorの導入、Echoフレームワークのカスタムバリデータの実装を学びます。体系的なエラーハンドリングと適切なHTTPステータスコードのマッピング、構造化されたエラーレスポンスの実装を行います。ユーザーフレンドリーなエラーメッセージと開発者向けのデバッグ情報のバランスを考慮した設計を習得します。

### 1.3.5. Day 5: テスト駆動開発と単体テスト

テスト駆動開発の基本概念と実践方法を学びます。ユニットテストとテーブル駆動テストを実装し、APIの動作を検証します。モックを活用したテスト手法も学び、依存関係を適切に分離します。

## 1.4. 使用技術と環境要件

- **バックエンド**
  - Go言語 (1.21以上)
  - Echo Webフレームワーク
  - sqlboiler (ORM)
  - golang-migrate (マイグレーションツール)
  - MySQL (データベース)

- **サーバーレス**
  - AWS Lambda
  - AWS S3
  - LocalStack (AWSエミュレーション)

- **開発環境**
  - Docker & Docker Compose
  - AWS CLI
  - go-task (タスクランナー)
  - air (ホットリロード)

## 1.5. 事前準備事項

- Week 1、2の実装が完了していること
- Docker環境が正常に動作していること
- Go言語の開発環境が整っていること
- LocalStackの設定が完了していること（S3、Lambda関連）
- AWS CLIがインストールされ、設定されていること

## 1.6. 学習のポイント

- **RESTful APIデザインの原則**
  リソース指向の設計、適切なHTTPメソッドとステータスコードの選択、ページネーションとフィルタリングの標準的な実装パターンを理解しましょう。

- **体系的なエラーハンドリング**
  エラーの種類に応じた適切なレスポンス形式と、一貫性のあるエラー処理フローの構築が重要です。ユーザーフレンドリーなエラーメッセージと開発者向けのデバッグ情報のバランスを考慮しましょう。

- **テスト駆動開発の実践**
  テストファーストの開発手法を体験し、コードの品質向上と保守性の高さを実感しましょう。テーブル駆動テストを活用することで、様々な入力パターンを効率的にテストできます。

- **サーバーレスアーキテクチャの理解**
  従来のサーバーベースのアプローチとサーバーレスアプローチの違いを理解し、適材適所での技術選択ができるようになりましょう。イベント駆動型の設計思想も同時に学びます。

## 1.7. 関連リソース

- **公式ドキュメント**
  - [Go言語公式ドキュメント](https://golang.org/doc/)
  - [Echo Webフレームワーク](https://echo.labstack.com/guide)
  - [AWS Lambda開発者ガイド](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
  - [AWS S3開発者ガイド](https://docs.aws.amazon.com/AmazonS3/latest/dev/Welcome.html)
  - [LocalStack公式ドキュメント](https://docs.localstack.cloud/overview/)

- **チュートリアルと参考資料**
  - [RESTful APIベストプラクティス](https://restfulapi.net/)
  - [Go言語でのテスト駆動開発入門](https://quii.gitbook.io/learn-go-with-tests/)
  - [サーバーレスアーキテクチャ入門](https://www.serverless.com/blog/serverless-architecture-code-patterns)
  - [AWS Well-Architected Labs: Serverless](https://www.wellarchitectedlabs.com/serverless/)

- **ブログと記事**
  - [GoでのRESTfulサービス実装パターン](https://eli.thegreenplace.net/2021/rest-servers-in-go-part-1-standard-library/)
  - [エラーハンドリングベストプラクティス](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
  - [AWS Lambdaとは何か](https://aws.amazon.com/jp/lambda/resources/)
