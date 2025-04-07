# 1. Week 2: データモデルと基本API - 概要

## 1.1. 目次

- [1. Week 2: データモデルと基本API - 概要](#1-week-2-データモデルと基本api---概要)
  - [1.1. 目次](#11-目次)
  - [1.2. 学習目標](#12-学習目標)
    - [1.2.1. 主要目標](#121-主要目標)
    - [1.2.2. 完了時の成果物](#122-完了時の成果物)
    - [1.2.3. 前提知識](#123-前提知識)
  - [1.3. 週の概要](#13-週の概要)
    - [1.3.1. Day 1: データベーススキーマの設計と実装](#131-day-1-データベーススキーマの設計と実装)
    - [1.3.2. Day 2: sqlboilerによるORM設定](#132-day-2-sqlboilerによるorm設定)
    - [1.3.3. Day 3: OpenAPI仕様の定義](#133-day-3-openapi仕様の定義)
    - [1.3.4. Day 4: サーバーレスアーキテクチャの基本](#134-day-4-サーバーレスアーキテクチャの基本)
    - [1.3.5. Day 5: Reactコンポーネントの基礎](#135-day-5-reactコンポーネントの基礎)
  - [1.4. 技術スタック](#14-技術スタック)
    - [1.4.1. 使用技術](#141-使用技術)
    - [1.4.2. 環境要件](#142-環境要件)
  - [1.5. 事前準備](#15-事前準備)
  - [1.6. 学習のポイント](#16-学習のポイント)
  - [1.7. 関連リソース](#17-関連リソース)

## 1.2. 学習目標

### 1.2.1. 主要目標

1. データベーススキーマの設計原則とMySQLでの実装方法を理解する
2. golang-migrateを使用したデータベースマイグレーション管理を習得する
3. sqlboilerによるORM設定と型安全なデータアクセスの実装方法を学ぶ
4. OpenAPI仕様を使ったAPI定義とogenによるコード生成を理解する
5. サーバーレスアーキテクチャの基本概念と、LocalStackを使用したLambda/S3環境のセットアップを学ぶ
6. データアクセスコードでのトランザクション管理の基本を理解する
7. Reactの基本的な概念とJSXの構文を学ぶ

### 1.2.2. 完了時の成果物

- **データベーススキーマ**: 商品(products)、カテゴリー(categories)、在庫(inventory)テーブルが適切な関係を持って設計・実装されていること
- **マイグレーションファイル**: golang-migrateを使用したマイグレーションスクリプトが実装され、テストデータが投入されていること
- **sqlboilerモデル**: MySQLデータベースから自動生成された型安全なモデルが正しく動作していること
- **OpenAPI仕様書**: 基本的なAPI仕様がYAML形式で定義され、Swagger UIで閲覧可能になっていること
- **ogenによる生成コード**: OpenAPI仕様からサーバー/クライアントコードが生成され、統合されていること
- **LocalStack環境設定**: Lambda関数とS3バケットが正しく設定され、基本操作が可能になっていること
- **基本的なReactコンポーネント**: 単純なReactコンポーネントが作成され、TypeScriptの型定義が適用されていること

### 1.2.3. 前提知識

- Go言語の基本的な構文とコンセプト（構造体、インターフェース、エラーハンドリングなど）
- MySQLを含むリレーショナルデータベースの基礎知識（テーブル、主キー、外部キーなど）
- SQL文の基本（SELECT、INSERT、UPDATE、DELETE）
- REST APIの基本概念（リソース、HTTPメソッド、エンドポイントなど）
- Docker Composeの基本的な使い方
- JSON/YAMLフォーマットの読み書き
- JavaScript/TypeScriptの基本構文

## 1.3. 週の概要

### 1.3.1. Day 1: データベーススキーマの設計と実装

データベース設計の基本原則を学びながら、eコマースアプリケーションに必要なテーブル設計を行います。golang-migrateを使用したマイグレーションファイルを作成し、テーブルの作成とテストデータの投入を実装します。

### 1.3.2. Day 2: sqlboilerによるORM設定

sqlboilerを使ったORMの構成方法を学び、MySQLスキーマからGoのモデルコードを自動生成します。基本的なデータアクセスコードを実装し、テストを通じて生成されたモデルの操作方法を習得します。

### 1.3.3. Day 3: OpenAPI仕様の定義

API設計の原則を学び、OpenAPI仕様(OAS3)を使ってAPIエンドポイントを定義します。ogenツールを使用してGoサーバーとクライアントコードを生成し、API仕様とコードの連携方法を習得します。

### 1.3.4. Day 4: サーバーレスアーキテクチャの基本

AWS Lambdaを中心としたサーバーレスアーキテクチャの基本概念を学びます。LocalStackを使用してローカル環境でLambda関数とS3バケットを設定し、シンプルなファイル処理機能の実装を通じて、イベント駆動型の開発パターンを理解します。

### 1.3.5. Day 5: Reactコンポーネントの基礎

Reactの基本的な概念とJSXの構文を理解し、シンプルなコンポーネントを作成します。TypeScriptを用いた型定義とPropsの受け渡し方法を学び、コンポーネントライフサイクルとイベントハンドリングの基礎を習得します。

## 1.4. 技術スタック

### 1.4.1. 使用技術

- **データベース**: MySQL 8.0
- **マイグレーションツール**: golang-migrate
- **ORM**: sqlboiler (MySQLドライバー)
- **API仕様**: OpenAPI 3.0 (OAS3)
- **コード生成ツール**: ogen
- **APIドキュメンテーション**: Swagger UI
- **AWS エミュレーター**: LocalStack
- **サーバーレスサービス**: Lambda, S3
- **フロントエンド**: React 18, TypeScript 5

### 1.4.2. 環境要件

- Docker Desktop 最新版
- Docker Compose v2.10.0 以上
- Go 1.21 以上
- Node.js v18 以上
- AWS CLI と awslocal (LocalStack用CLIラッパー)
- エディタ: Visual Studio Code (推奨拡張機能: Go, REST Client, SQLTools)
- Git

## 1.5. 事前準備

学習を始める前に、以下の準備を行っておくことをお勧めします：

1. Week 1で構築したDocker Compose環境が正常に動作することを確認する
2. MySQLクライアントツール（CLI、GUI、またはVS Codeの拡張機能）を用意する
3. golang-migrate CLIをインストールする: `go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
4. sqlboilerとsqlboiler-mysqlをインストールする: `go install github.com/volatiletech/sqlboiler/v4@latest && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest`
5. ogenをインストールする: `go install github.com/ogen-go/ogen/cmd/ogen@latest`
6. AWS CLIとawslocal (LocalStack CLI) をインストールする
7. 基本的なSQL文法とリレーショナルデータベース設計の資料に目を通しておく

## 1.6. 学習のポイント

この週で特に注目すべきポイントは以下の通りです：

1. **データベース設計の重要性**: テーブル間のリレーションシップの適切な設計がアプリケーション全体の性能と拡張性に大きく影響することを理解しましょう。一度設計したスキーマを後から変更することは難しいため、最初の設計段階で十分に検討することが重要です。

2. **型安全なデータアクセス**: sqlboilerが提供する型安全なクエリビルダーを活用することで、コンパイル時にSQLの誤りを検出できます。これにより実行時エラーを減少させ、コードの信頼性を向上させることができます。

3. **API駆動開発（API-first）**: OpenAPI仕様を先に定義することで、フロントエンドとバックエンドの開発者が共通の理解を持って並行して作業できるようになります。ドキュメントとしての役割だけでなく、コード生成によって開発効率も向上します。

4. **サーバーレスの考え方**: 従来のサーバーベースのアーキテクチャとサーバーレスアーキテクチャの違いを理解し、それぞれの適切なユースケースを学びましょう。イベント駆動型の設計思想はマイクロサービスアーキテクチャの重要な基盤となります。

5. **宣言的UIの概念**: Reactの宣言的なUI構築アプローチは、命令型のDOMマニピュレーションと比較して、コードの可読性と保守性を高めます。コンポーネントベースの開発がもたらす再利用性とモジュール性を理解しましょう。

## 1.7. 関連リソース

- [MySQL公式ドキュメント](https://dev.mysql.com/doc/)
- [Go Database Tutorial](https://go.dev/doc/tutorial/database-access)
- [golang-migrate ドキュメント](https://github.com/golang-migrate/migrate)
- [sqlboiler ドキュメント](https://github.com/volatiletech/sqlboiler)
- [OpenAPI 3.0 仕様](https://spec.openapis.org/oas/v3.0.3)
- [ogen ドキュメント](https://github.com/ogen-go/ogen)
- [AWS Lambda 開発者ガイド](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
- [LocalStack ドキュメント](https://docs.localstack.cloud/)
- [React 公式ドキュメント](https://reactjs.org/docs/getting-started.html)
- [TypeScript ハンドブック](https://www.typescriptlang.org/docs/handbook/intro.html)
