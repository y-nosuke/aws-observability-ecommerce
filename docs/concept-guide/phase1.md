# 1. AWS オブザーバビリティ学習用 eコマースアプリ - フェーズ1 全体設計と概念ガイド

## 1.1. 目次

- [1. AWS オブザーバビリティ学習用 eコマースアプリ - フェーズ1 全体設計と概念ガイド](#1-aws-オブザーバビリティ学習用-eコマースアプリ---フェーズ1-全体設計と概念ガイド)
  - [1.1. 目次](#11-目次)
  - [1.2. フェーズ1の概要と目標](#12-フェーズ1の概要と目標)
    - [1.2.1. 主要目標](#121-主要目標)
  - [1.3. アーキテクチャ設計](#13-アーキテクチャ設計)
    - [1.3.1. 全体アーキテクチャ](#131-全体アーキテクチャ)
    - [1.3.2. 技術スタックの概要](#132-技術スタックの概要)
      - [1.3.2.1. バックエンド](#1321-バックエンド)
      - [1.3.2.2. フロントエンド](#1322-フロントエンド)
      - [1.3.2.3. データストア](#1323-データストア)
      - [1.3.2.4. 開発環境](#1324-開発環境)
    - [1.3.3. 環境構成](#133-環境構成)
  - [1.4. 主要技術の概念説明と選定理由](#14-主要技術の概念説明と選定理由)
    - [1.4.1. バックエンド](#141-バックエンド)
      - [1.4.1.1. Go言語](#1411-go言語)
      - [1.4.1.2. Echo Webフレームワーク](#1412-echo-webフレームワーク)
      - [1.4.1.3. sqlboiler](#1413-sqlboiler)
      - [1.4.1.4. OpenAPI/ogen](#1414-openapiogen)
      - [1.4.1.5. slog](#1415-slog)
    - [1.4.2. フロントエンド](#142-フロントエンド)
      - [1.4.2.1. Next.js](#1421-nextjs)
      - [1.4.2.2. TypeScript](#1422-typescript)
      - [1.4.2.3. TailwindCSS](#1423-tailwindcss)
    - [1.4.3. データベース](#143-データベース)
      - [1.4.3.1. MySQL](#1431-mysql)
      - [1.4.3.2. golang-migrate](#1432-golang-migrate)
    - [1.4.4. 開発環境](#144-開発環境)
      - [1.4.4.1. Docker/Docker Compose](#1441-dockerdocker-compose)
      - [1.4.4.2. LocalStack](#1442-localstack)
      - [1.4.4.3. Traefik](#1443-traefik)
  - [1.5. コンポーネント間の関連性とデータフロー](#15-コンポーネント間の関連性とデータフロー)
    - [1.5.1. 主要コンポーネント](#151-主要コンポーネント)
    - [1.5.2. データフロー](#152-データフロー)
    - [1.5.3. API通信](#153-api通信)
  - [1.6. 機能要件と設計方針](#16-機能要件と設計方針)
    - [1.6.1. 商品カタログ機能](#161-商品カタログ機能)
      - [1.6.1.1. 商品一覧表示](#1611-商品一覧表示)
      - [1.6.1.2. 商品詳細表示](#1612-商品詳細表示)
      - [1.6.1.3. カテゴリーナビゲーション](#1613-カテゴリーナビゲーション)
    - [1.6.2. 管理画面の基本機能](#162-管理画面の基本機能)
      - [1.6.2.1. 管理者認証（モック）](#1621-管理者認証モック)
      - [1.6.2.2. 商品情報管理](#1622-商品情報管理)
      - [1.6.2.3. カテゴリー管理](#1623-カテゴリー管理)
    - [1.6.3. 基本的なオブザーバビリティ機能とサーバーレス基盤](#163-基本的なオブザーバビリティ機能とサーバーレス基盤)
      - [1.6.3.1. ヘルスチェック](#1631-ヘルスチェック)
      - [1.6.3.2. サーバーレス基盤](#1632-サーバーレス基盤)
  - [1.7. 各週の作業内容と連携](#17-各週の作業内容と連携)
    - [1.7.1. 週1: プロジェクト基盤構築](#171-週1-プロジェクト基盤構築)
    - [1.7.2. 週2: データモデルと基本API](#172-週2-データモデルと基本api)
    - [1.7.3. 週3: 商品カタログバックエンドの完成](#173-週3-商品カタログバックエンドの完成)
    - [1.7.4. 週4: 顧客向け商品閲覧UI](#174-週4-顧客向け商品閲覧ui)
    - [1.7.5. 週5: 管理画面の基本実装とスタイリング](#175-週5-管理画面の基本実装とスタイリング)
  - [1.8. 完成時の期待される状態](#18-完成時の期待される状態)
    - [1.8.1. システム全体](#181-システム全体)
    - [1.8.2. バックエンド](#182-バックエンド)
    - [1.8.3. フロントエンド](#183-フロントエンド)
    - [1.8.4. データベース](#184-データベース)
  - [1.9. テスト方法](#19-テスト方法)
    - [1.9.1. 機能テスト](#191-機能テスト)
    - [1.9.2. 技術テスト](#192-技術テスト)
    - [1.9.3. 手動テスト手順](#193-手動テスト手順)
    - [1.9.4. 自動テストの実行](#194-自動テストの実行)
  - [1.10. フェーズ2への橋渡し](#110-フェーズ2への橋渡し)
  - [1.11. 付録: ディレクトリ構造とファイル構成](#111-付録-ディレクトリ構造とファイル構成)

## 1.2. フェーズ1の概要と目標

フェーズ1「基盤構築と商品閲覧機能」は、このプロジェクトの土台となる5週間です。このフェーズでは、開発環境の構築から始め、eコマースアプリケーションの最も基本的な機能である「商品閲覧」の実装までを行います。

### 1.2.1. 主要目標

1. **開発環境の構築**: Docker、Go/Echo、Next.js、LocalStackなどの開発環境をセットアップします
2. **バックエンドとフロントエンドの基本構造実装**: 両環境の基本構造を実装し連携させます
3. **データモデルの設計と実装**: 商品カタログに必要なデータベーススキーマを設計・実装します
4. **商品カタログ機能の実装**: 商品一覧・詳細表示などの基本APIとフロントエンドを実装します
5. **管理画面の基本的なUI実装**: 商品管理のための基本的な管理画面を実装します
6. **サーバーレスの基本概念理解と実装**: LocalStackを活用し、シンプルなサーバーレス関数（Lambda）とS3バケットの基本機能を実装します

このフェーズを通じて、MVPレベルのeコマースアプリケーションを構築します。フェーズ1で実装する機能は限定的ですが、以降のフェーズで拡張していくための堅固な基盤を作ることが重要です。

## 1.3. アーキテクチャ設計

### 1.3.1. 全体アーキテクチャ

フェーズ1では、以下の図のようなシンプルな構成でアプリケーションを構築します。顧客向け画面と管理画面を別のプロジェクトとして分離し、ホスト名ベースのルーティングを採用しています。

```text
┌───────────────┐      ┌───────────────┐      ┌───────────────┐
│               │      │               │      │               │
│  Next.js      │      │  Go/Echo      │      │  MySQL        │
│ 顧客向け画面   │ ←──→ │  Backend      │ ←──→ │  Database     │
│(customer.localhost)│      │               │      │               │
└───────────────┘      └───────────────┘      └───────────────┘
                              ↑
┌───────────────┐             │
│               │             │
│  Next.js      │ ←───────────┘
│  管理画面      │
│(admin.localhost)│
└───────────────┘      ┌───────────────┐      ┌───────────────┐
                       │               │      │               │
                       │  LocalStack   │      │  Traefik      │
                       │  AWS エミュレータ │      │  リバースプロキシ │
                       │               │      │(ホスト名ベースルーティング)│
                       └───────────────┘      └───────────────┘
```

このアーキテクチャでは、典型的な3層構造に加えて、フロントエンドを顧客向け画面と管理画面の2つのプロジェクトに分離しています。両方のフロントエンドは同じバックエンドAPIを利用しますが、独立して開発・デプロイが可能です。Traefikリバースプロキシがホスト名に基づいてルーティングを行い、異なるドメイン（shop.localhostとadmin.localhost）として提供します。また、LocalStackを追加して、開発環境でもAWSサービスをエミュレートできるようにしています。

### 1.3.2. 技術スタックの概要

#### 1.3.2.1. バックエンド

- **言語**: Go
- **Webフレームワーク**: Echo
- **ORM**: sqlboiler (MySQLに特化)
- **API定義**: OpenAPI (ogenでコード生成)
- **ログ**: slog (Go標準の構造化ログ)

#### 1.3.2.2. フロントエンド

- **顧客向け画面**:
  - **フレームワーク**: Next.js (Reactベース)
  - **言語**: TypeScript
  - **スタイリング**: TailwindCSS
  - **状態管理**: React Hooks (フェーズ1ではシンプルに)

- **管理画面**:
  - **フレームワーク**: Next.js (Reactベース)
  - **言語**: TypeScript
  - **スタイリング**: TailwindCSS
  - **状態管理**: React Hooks (フェーズ1ではシンプルに)
  - **特徴**: 顧客向け画面とは別プロジェクトとして実装

#### 1.3.2.3. データストア

- **データベース**: MySQL
- **マイグレーション**: golang-migrate

#### 1.3.2.4. 開発環境

- **コンテナ化**: Docker, Docker Compose
- **ロードバランサー**: Traefik
- **AWSエミュレータ**: LocalStack
- **バージョン管理**: Git, GitHub
- **タスクランナー**: go-task

### 1.3.3. 環境構成

フェーズ1では、Docker Composeを使用して以下のサービスを構成します：

1. **mysql**: データベースサーバー
2. **backend**: Go/Echoバックエンドアプリケーション
3. **frontend-customer**: 顧客向けNext.jsフロントエンドアプリケーション
4. **frontend-admin**: 管理者向けNext.jsフロントエンドアプリケーション
5. **traefik**: リバースプロキシとロードバランサー
6. **localstack**: AWSサービスエミュレーター

これらは全て同一のDocker Networkに所属し、コンテナ間の通信が可能です。Traefikは外部からのリクエストを適切なサービスにルーティングする役割を担います。ルーティングはホスト名ベースで行われ、顧客向け画面は`shop.localhost`、管理画面は`admin.localhost`というホスト名でアクセスします。これにより、両フロントエンドは完全に分離された環境で動作することができます。

## 1.4. 主要技術の概念説明と選定理由

### 1.4.1. バックエンド

#### 1.4.1.1. Go言語

**概要**: Googleによって開発された静的型付け、コンパイル言語。シンプルな構文、高いパフォーマンス、優れた並行処理機能が特徴です。

**選定理由**:

- 優れたパフォーマンスとリソース効率（メモリ消費が少ない）
- AWS SDKがGoをファーストクラスサポート
- 静的型付けによるコンパイル時のエラー検出
- シンプルで学びやすい構文
- マイクロサービスやクラウドネイティブアプリケーションでの人気と実績

#### 1.4.1.2. Echo Webフレームワーク

**概要**: Goのミニマリストで高性能なWebフレームワークで、ルーティング、ミドルウェア、HTTPリクエスト/レスポンス操作の機能を提供します。

**選定理由**:

- 高性能でありながらシンプルなAPI
- ミドルウェアのサポートが充実（ログ、認証、CORS等）
- 一般的なWebサービスに必要な機能が揃っている
- コード生成との相性が良い
- 活発なコミュニティとドキュメント

#### 1.4.1.3. sqlboiler

**概要**: データベーススキーマからGoのコードを生成するORM（Object Relational Mapper）ツール。型安全なデータアクセスが可能です。

**選定理由**:

- コード生成によるボイラープレートの削減
- 型安全なクエリビルダー
- RawモードでのSQLクエリの実行もサポート
- MySQLに最適化された機能
- パフォーマンスを重視した設計

#### 1.4.1.4. OpenAPI/ogen

**概要**: OpenAPIはRESTful APIを定義するための標準仕様。ogenはOpenAPI仕様からGoコードを生成するツールです。

**選定理由**:

- API仕様の標準化と一元管理
- クライアント/サーバー両方のコード生成が可能
- 自動ドキュメント生成
- 型安全なAPI定義
- フロントエンドとバックエンドの開発を並行して進められる

#### 1.4.1.5. slog

**概要**: Go 1.21で導入された標準ライブラリの構造化ログパッケージ。JSON形式のログ出力などをサポートしています。

**選定理由**:

- Go標準ライブラリの一部であり、将来的な互換性が保証される
- 構造化ログによるログ解析の容易さ
- 様々なバックエンド（コンソール、ファイル、クラウドロギングサービス）に対応
- パフォーマンスが最適化されている
- 後のフェーズでCloudWatch Logsやオブザーバビリティツールとの統合が容易

### 1.4.2. フロントエンド

#### 1.4.2.1. Next.js

**概要**: Reactベースのフレームワークで、サーバーサイドレンダリング（SSR）、静的サイト生成（SSG）、APIルートなどの機能を提供します。

**選定理由**:

- サーバーサイドレンダリングによるSEO対策とパフォーマンス向上
- ファイルベースのルーティングによる直感的な開発体験
- APIルートによるバックエンド機能の実装
- React Serverコンポーネントのサポート
- Vercelによる本番環境へのデプロイが容易

#### 1.4.2.2. TypeScript

**概要**: JavaScriptのスーパーセットとして、静的型付けを追加した言語です。

**選定理由**:

- 静的型システムによるエラー発見の早期化
- IDEによるコード補完とドキュメンテーションの向上
- リファクタリングの安全性向上
- チーム開発における品質とメンテナンス性の向上
- JavaScript互換性による学習曲線の緩和

#### 1.4.2.3. TailwindCSS

**概要**: ユーティリティファーストのCSSフレームワークで、HTMLにクラスを直接追加してスタイリングを行います。

**選定理由**:

- ユーティリティクラスによる迅速な開発
- カスタマイズ性の高さ
- コンポーネント間のスタイルの一貫性維持が容易
- CSSファイルの管理が不要
- レスポンシブデザインの実装が容易

### 1.4.3. データベース

#### 1.4.3.1. MySQL

**概要**: オープンソースのリレーショナルデータベース管理システム。

**選定理由**:

- 信頼性と安定性の高さ
- 広く使われているため情報やリソースが豊富
- eコマースアプリケーション用途に十分な機能
- AWSのRDSとの互換性
- sqlboilerとの相性の良さ

#### 1.4.3.2. golang-migrate

**概要**: データベースマイグレーション管理ツール。スキーマの変更を追跡し、バージョン管理します。

**選定理由**:

- Go言語からの操作が容易
- シンプルで直感的なAPIとコマンドライン
- SQLベースのマイグレーションで自由度が高い
- バージョン管理との統合が容易
- ロールバック機能による安全性

### 1.4.4. 開発環境

#### 1.4.4.1. Docker/Docker Compose

**概要**: コンテナ化技術とその管理ツールで、アプリケーションと依存関係を一貫した環境で実行できます。

**選定理由**:

- 環境の一貫性と再現性
- 複数サービスの統合的な管理
- 開発環境と本番環境の違いを最小化
- チーム内での環境共有の容易さ
- 「動作環境」の配布が容易

#### 1.4.4.2. LocalStack

**概要**: AWSサービスをローカル環境でエミュレートするツール。

**選定理由**:

- AWS環境の無料エミュレーション
- オブザーバビリティサービス（CloudWatch, X-Ray等）のローカルテスト
- サーバーレスサービス（Lambda, S3等）のローカル開発とテスト
- 開発とテストのサイクルを高速化
- オフライン開発が可能
- 将来的なAWS環境への移行が容易

#### 1.4.4.3. Traefik

**概要**: モダンなHTTPリバースプロキシとロードバランサー。

**選定理由**:

- 自動設定検出機能によるシームレスな統合
- 動的設定変更のサポート
- SSL/TLS対応の容易さ
- HTTPSリダイレクト、基本認証など、一般的な機能のサポート
- Docker Composeとの優れた統合性

## 1.5. コンポーネント間の関連性とデータフロー

### 1.5.1. 主要コンポーネント

フェーズ1では以下の主要コンポーネントを実装します：

1. **データベースレイヤー**
   - 商品、カテゴリー、在庫情報のテーブル
   - sqlboilerによる生成されたモデル
   - マイグレーションスクリプト

2. **バックエンドAPI**
   - OpenAPI定義に基づいたエンドポイント
   - リポジトリパターンによるデータアクセス
   - 商品カタログ操作のビジネスロジック
   - 基本的なログ出力

3. **フロントエンドUI**
   - **顧客向け商品閲覧画面**:
     - 商品一覧・詳細表示画面
     - カテゴリーナビゲーション
     - TailwindCSSによるスタイリング
     - APIクライアント

   - **管理者向け基本操作画面** (別プロジェクト):
     - 管理者認証
     - 商品管理画面
     - カテゴリー管理
     - TailwindCSSによるスタイリング
     - APIクライアント

4. **開発環境コンポーネント**
   - Docker Composeによるサービス定義
   - LocalStackによるAWSエミュレーション
   - Traefikによるリバースプロキシ

### 1.5.2. データフロー

基本的なデータフローは以下のようになります：

1. **顧客による商品閲覧**:

   ```text
   ユーザー → 顧客向けNext.js → API Client → Go/Echo バックエンド → sqlboiler → MySQL →
   sqlboiler → Go/Echo バックエンド → API Client → 顧客向けNext.js → ユーザー
   ```

2. **管理者による商品管理**:

   ```text
   管理者 → 管理画面Next.js → API Client → Go/Echo バックエンド → sqlboiler → MySQL →
   sqlboiler → Go/Echo バックエンド → API Client → 管理画面Next.js → 管理者
   ```

3. **ログ出力**:

   ```text
   アプリケーションイベント → slog → コンソール出力
   ```

### 1.5.3. API通信

フェーズ1では、以下のようなAPI通信が実装されます：

1. **商品閲覧API**:
   - GET /api/products - 商品一覧取得（ページネーション、フィルタリング）
   - GET /api/products/{id} - 商品詳細取得
   - GET /api/categories - カテゴリー一覧取得
   - GET /api/categories/{id}/products - カテゴリー別商品一覧

2. **管理API** (モック認証付き):
   - GET /api/admin/products - 管理用商品一覧
   - POST /api/admin/products - 商品作成
   - PUT /api/admin/products/{id} - 商品更新
   - DELETE /api/admin/products/{id} - 商品削除

3. **システムAPI**:
   - GET /api/health - ヘルスチェック

Next.jsのフロントエンドは、これらのAPIを呼び出してデータのやり取りを行います。OpenAPIの定義に基づいてコード生成されたクライアントを使用することで、型安全なAPI呼び出しが可能になります。

## 1.6. 機能要件と設計方針

### 1.6.1. 商品カタログ機能

#### 1.6.1.1. 商品一覧表示

**要件**:

- 商品のページネーションによる一覧表示
- カテゴリーによるフィルタリング
- 基本的な商品情報（名前、価格、画像など）の表示
- ページネーション制御

**設計方針**:

- ページネーションはオフセットベースで実装（シンプルさ優先）
- 一度に表示する商品数を制限（デフォルト20件）
- カテゴリーIDによるフィルタリングをクエリパラメータで実装
- レスポンシブデザインによるモバイル対応

#### 1.6.1.2. 商品詳細表示

**要件**:

- 商品の詳細情報（説明、仕様など）の表示
- 在庫状況の表示
- 関連商品の基本表示

**設計方針**:

- 商品IDによるルーティング
- 在庫情報は簡易表示（在庫あり/なし）
- 画像は最適化されたNext.jsの`Image`コンポーネントで表示
- 関連商品は同カテゴリー商品から抽出

#### 1.6.1.3. カテゴリーナビゲーション

**要件**:

- カテゴリー一覧の表示
- カテゴリーによる商品フィルタリング
- パンくずナビゲーション

**設計方針**:

- ヘッダーとサイドバーにカテゴリーナビゲーションを配置
- 階層構造はシンプルに1段階（サブカテゴリーはフェーズ1では未実装）
- カテゴリー一覧はキャッシュして頻繁な再読み込みを避ける

### 1.6.2. 管理画面の基本機能

#### 1.6.2.1. 管理者認証（モック）

**要件**:

- 管理画面へのアクセス制限
- ログイン機能
- ログアウト機能

**設計方針**:

- フェーズ1では簡易的なモック認証を実装
- JWTを使用したシンプルな認証
- ログイン状態の保持にはブラウザのlocalStorageを使用
- 保護されたルートの設定（認証なしではアクセス不可）

#### 1.6.2.2. 商品情報管理

**要件**:

- 商品一覧の表示
- 商品情報の閲覧
- 基本的な商品情報の編集

**設計方針**:

- 商品一覧はテーブル形式で表示
- 基本的なフィルタリングと検索機能
- インラインでの簡易編集機能
- 詳細情報は別画面で表示

#### 1.6.2.3. カテゴリー管理

**要件**:

- カテゴリー一覧表示
- カテゴリーの基本操作（閲覧のみ）

**設計方針**:

- リスト形式でのカテゴリー表示
- 階層構造の視覚的表現
- フェーズ1では編集機能は最小限に抑える

### 1.6.3. 基本的なオブザーバビリティ機能とサーバーレス基盤

#### 1.6.3.1. ヘルスチェック

**要件**:

- システムの基本的な健全性確認
- APIの動作確認

**設計方針**:

- シンプルなヘルスチェックエンドポイント
- データベース接続の確認
- 基本的なシステムメトリクスの収集（フェーズ1ではローカルのみ）

#### 1.6.3.2. サーバーレス基盤

**要件**:

- 商品画像リサイズなどの基本的な処理機能
- S3バケットへのファイル保存と取得
- トリガーとイベントの基本的な理解

**設計方針**:

- LocalStackを活用したLambda環境の構築
- S3バケットの基本設定と操作
- シンプルな画像処理Lambda関数の実装
- 基本的なS3トリガーパターンの理解
- サーバーレス関数と通常バックエンドの連携基盤

## 1.7. 各週の作業内容と連携

### 1.7.1. 週1: プロジェクト基盤構築

**主要タスク**:

- Docker Compose環境の構築
- Go/Echo開発環境の準備
- Next.js/TypeScript環境の準備
- ヘルスチェックAPIの実装
- GitHubリポジトリのセットアップ

**作業内容の連携**:

- この週は後続の全ての作業の基盤となります
- Docker ComposeはLocalStackを含めた環境を整備し、フェーズ2以降のAWS統合の準備になります
- バックエンドとフロントエンドの基本構造が、以降の機能実装の土台になります

**成果物**:

- 動作する開発環境（Docker Compose）
- 基本的なヘルスチェックAPI
- プロジェクト構造とリポジトリの初期設定
- Traefikのホスト名ベースルーティング設定

### 1.7.2. 週2: データモデルと基本API

**主要タスク**:

- データベーススキーマの設計と実装
- sqlboilerの設定とモデル生成
- OpenAPI仕様の初期定義
- Reactの基本的な概念の理解
- サーバーレスアーキテクチャの基本概念学習
- LocalStackでのLambdaとS3の基本設定

**作業内容の連携**:

- データモデルは商品カタログの中核となり、週3以降の全ての機能実装の基礎になります
- OpenAPI仕様は、バックエンドとフロントエンドの契約として機能し、並行開発を可能にします
- LocalStackでのLambdaとS3の基本設定は、フェーズ2以降のサーバーレス機能拡張の準備になります
- サーバーレスアーキテクチャの基本概念を学ぶことで、イベント駆動型設計への理解を深めますての機能実装の基礎になります
- OpenAPI仕様は、バックエンドとフロントエンドの契約として機能し、並行開発を可能にします
- 基本的なログ出力はフェーズ2のオブザーバビリティ強化の準備になります

**成果物**:

- 商品、カテゴリー、在庫のデータモデル
- sqlboilerによる生成モデル
- 基本的なAPIエンドポイント定義

### 1.7.3. 週3: 商品カタログバックエンドの完成

**主要タスク**:

- 商品一覧APIの実装
- 商品詳細APIの実装
- カテゴリー別商品一覧APIの実装
- バリデーションとエラーハンドリング
- バックエンドのテスト実装
- シンプルな画像処理Lambda関数の実装
- S3を使った商品画像の保存と取得機能の実装

**作業内容の連携**:

- 商品カタログAPIは週4のフロントエンドUI実装の基盤になります
- エラーハンドリングは全体のエラー管理とオブザーバビリティの基礎になります
- テスト実装は品質保証の基礎として、以降の全ての開発に適用されます
- Lambda関数実装はサーバーレスパターンの基本を学ぶ機会となります
- S3との連携実装は、クラウドストレージの基本的な操作パターンを確立します

**成果物**:

- 完全な商品カタログAPI
- エラーハンドリングメカニズム
- テストコード
- シンプルな商品画像処理Lambda関数
- S3バケット連携機能

### 1.7.4. 週4: 顧客向け商品閲覧UI

**主要タスク**:

- 顧客向け商品一覧ページの実装
- 顧客向け商品詳細ページの実装
- 顧客向けカテゴリーナビゲーションの実装
- 顧客向けフロントエンドのテスト実装
- 顧客向けNext.jsの基本ルーティング設定

**作業内容の連携**:

- 顧客向け商品閲覧UIは顧客体験の基本であり、フェーズ4のカート/注文機能の基盤になります
- コンポーネント設計は再利用可能なUIライブラリの基礎となります
- API連携パターンは、フロントエンド-バックエンド通信の雛形となります
- 管理画面とは分離されているため、顧客向け機能に集中して開発できます

**成果物**:

- 顧客向け商品一覧ページ
- 顧客向け商品詳細ページ
- 顧客向けカテゴリーナビゲーション
- 再利用可能なUIコンポーネント

### 1.7.5. 週5: 管理画面の基本実装とスタイリング

**主要タスク**:

- 管理画面（別プロジェクト）の基本設定
- 管理画面のTailwindCSS設定と活用
- 管理画面レイアウトの実装
- 商品管理の基本画面実装
- モック認証システムの実装
- レスポンシブデザインの実装

**作業内容の連携**:

- 管理画面は完全に分離したプロジェクトとして実装し、フェーズ5の管理機能拡張の基盤になります
- 顧客向け画面とは独立してデプロイ・スケーリングが可能な設計とします
- TailwindCSSのスタイリングパターンは、両方のフロントエンドプロジェクトに適用されます
- モック認証はフェーズ4以降の本格的な認証システムの準備になります

**成果物**:

- スタイリングガイドラインとパターン
- 管理画面の基本UI（独立したNext.jsプロジェクト）
- 商品管理の基本機能
- モック認証機能
- 両フロントエンドのデプロイ構成

## 1.8. 完成時の期待される状態

フェーズ1完了時には、以下の状態になっていることが期待されます：

### 1.8.1. システム全体

- Docker Compose環境が正常に稼働し、全てのサービスが連携している
- ローカル環境での開発が容易に行える状態
- 基本的なログ出力とエラーハンドリングが整備されている
- GitHubリポジトリが整備され、コミット履歴が整理されている

### 1.8.2. バックエンド

- 全ての商品カタログ関連APIが実装され、正常に動作する
- データモデルが適切に設計され、マイグレーション管理されている
- エラーハンドリングが統一的に実装されている
- ユニットテストとインテグレーションテストが実装されている
- OpenAPI仕様が正確に定義され、コード生成に使用されている

### 1.8.3. フロントエンド

**顧客向け画面**:

- 商品一覧と詳細表示が実装されている
- カテゴリーナビゲーションが機能している
- TailwindCSSによる適切なスタイリングが適用されている
- レスポンシブデザインが実装されている

**管理画面**:

- 基本レイアウトと商品管理機能が実装されている
- モック認証機能が実装されている
- TailwindCSSによる適切なスタイリングが適用されている
- レスポンシブデザインが実装されている
- 適切なURL構造（/admin/以下）でルーティングされている

### 1.8.4. データベース

- 商品、カテゴリー、在庫のテーブルが作成されている
- テストデータが投入されている
- マイグレーションスクリプトが整備されている

## 1.9. テスト方法

フェーズ1の成果物をテストするための方法を以下に示します：

### 1.9.1. 機能テスト

1. **商品一覧機能テスト**
   - トップページにアクセスし、商品が表示されることを確認
   - ページネーションが機能し、前後のページに移動できることを確認
   - カテゴリーフィルタリングが機能することを確認
   - 商品カードから詳細ページに遷移できることを確認

2. **商品詳細機能テスト**
   - 商品詳細ページにアクセスし、正しい商品情報が表示されることを確認
   - 画像が適切に表示されることを確認
   - 在庫状況が表示されることを確認
   - 関連商品が表示されることを確認

3. **管理画面機能テスト**
   - モックログイン機能を使って管理画面にアクセスできることを確認
   - 商品一覧が表示されることを確認
   - 基本的なフィルタリングと検索が機能することを確認
   - 簡易的な商品情報編集が機能することを確認

4. **レスポンシブデザインテスト**
   - 異なる画面サイズ（デスクトップ、タブレット、モバイル）でUIが適切に表示されることを確認
   - モバイル表示でメニュー操作が可能なことを確認

### 1.9.2. 技術テスト

1. **バックエンドテスト**
   - APIエンドポイントが仕様通りに動作することを確認
   - エラーレスポンスが正しく返されることを確認
   - バリデーションが機能することを確認
   - 自動テストが実行され、パスすることを確認

2. **フロントエンドテスト**
   - 顧客向け画面と管理画面の両方でルーティングが正しく機能することを確認
   - それぞれのNext.jsアプリケーションがAPIクライアントを通じて正しくデータを取得できることを確認
   - コンポーネントが想定通りにレンダリングされることを確認
   - エラー状態やローディング状態が適切に処理されることを確認

3. **開発環境テスト**
   - Docker Composeでのサービス起動が正常に行われることを確認
   - LocalStackでのAWSエミュレーションが機能することを確認
   - Traefikのホスト名ベースルーティングが適切に機能することを確認
   - ホットリロードが機能することを確認

### 1.9.3. 手動テスト手順

1. **開発環境の起動**

   ```bash
   docker-compose up -d
   ```

2. **ホスト名の設定**
   - `/etc/hosts`ファイル（Windowsの場合は`C:\Windows\System32\drivers\etc\hosts`）に以下を追加:

   ```text
   127.0.0.1  shop.localhost
   127.0.0.1  admin.localhost
   ```

3. **顧客向け画面テスト**
   - `http://shop.localhost` にアクセス
   - 商品リストが表示されることを確認
   - カテゴリーフィルターを操作
   - 商品カードをクリックして詳細ページを確認

4. **管理機能テスト**
   - `http://admin.localhost` にアクセス
   - 提供されたテスト認証情報でログイン
   - 商品管理インターフェースで操作を確認

5. **API直接テスト**
   - Postman/curlなどで各APIエンドポイントに直接リクエスト
   - レスポンスを確認

### 1.9.4. 自動テストの実行

```bash
# バックエンドテスト実行
cd backend
go test ./...

# フロントエンドテスト実行
cd frontend
npm test
```

## 1.10. フェーズ2への橋渡し

フェーズ1で構築した基盤は、フェーズ2「統合オブザーバビリティとサーバーレス」へのスムーズな移行を可能にします。フェーズ2に向けて、以下の点を意識しておくことが重要です：

1. **ログ基盤の構築**
   - フェーズ2ではslogを活用した構造化ログシステムとCloudWatch Logs統合が実装されます
   - フェーズ1で構築したLocalStack環境がCloudWatch Logsエミュレーションの基盤になります

2. **メトリクス収集の準備**
   - フェーズ2ではメトリクス収集が実装されるため、計測ポイントとなる箇所（リクエスト処理、DBアクセスなど）を意識しておきましょう
   - パフォーマンス計測の基本的な考え方を理解しておくと有益です

3. **LocalStackの活用**
   - フェーズ1で設定したLocalStackは、フェーズ2でのAWSサービス（CloudWatch, X-Ray等）エミュレーションの基盤になります
   - 基本的な操作方法に慣れておくことで、スムーズな移行が可能になります

4. **フロントエンドデータフェッチの拡張**
   - フェーズ1で実装したAPIクライアントは、フェーズ2でSWRを用いたデータ管理に拡張されます
   - 基本的なデータフェッチパターンを理解しておくことで、拡張がスムーズになります

5. **イベント処理の準備**
   - フェーズ2では基本的なサーバーレス機能が導入されるため、イベント駆動型設計の基本概念を理解しておくと有益です

フェーズ1完了時に、これらの点を意識した復習を行うことで、フェーズ2へのスムーズな移行が可能になります。

## 1.11. 付録: ディレクトリ構造とファイル構成

フェーズ1で構築されるプロジェクトの基本的なディレクトリ構造とファイル構成を示します。これは実装の指針となるものです。

```text
/
├── .github/                     # GitHub関連設定
│   └── workflows/               # GitHub Actions設定
│       └── ci.yml               # CI設定
│
├── backend-api/                 # バックエンドAPIアプリケーション
│   ├── internal/                # 非公開パッケージ
│   │   ├── api/                 # API定義と実装
│   │   │   ├── handlers/        # APIハンドラー
│   │   │   ├── models/          # APIモデル
│   │   │   └── openapi/         # OpenAPI生成コード
│   │   │
│   │   ├── config/              # 設定関連
│   │   │   └── config.go        # 設定管理
│   │   │
│   │   ├── db/                  # データベース関連
│   │   │   ├── models/          # sqlboiler生成モデル
│   │   │   └── migrations/      # マイグレーションファイル
│   │   │
│   │   ├── repository/          # リポジトリパターン実装
│   │   │   ├── product/         # 商品リポジトリ
│   │   │   └── category/        # カテゴリーリポジトリ
│   │   │
│   │   └── service/             # サービスレイヤー
│   │
│   ├── static/                  # 静的ファイル
│   │   └── swagger-ui/          # Swagger UI
│   │
│   ├── tmp/                     # 一時ファイル
│   │
│   ├── main.go                  # メインエントリポイント
│   ├── openapi.yaml             # OpenAPI仕様ファイル
│   ├── oapi-codegen-config.yaml # OpenAPI生成設定
│   ├── Dockerfile               # バックエンドDockerfile
│   ├── go.mod                   # Goモジュール定義
│   ├── go.sum                   # Goモジュールロック
│   └── sqlboiler.toml           # sqlboiler設定
│
├── backend-image-processor/     # 画像処理サービス
│   └── main.go                  # メインエントリポイント
│
├── frontend-customer/           # 顧客向けフロントエンドアプリケーション
│   ├── app/                     # Next.js App Router
│   │   └── products/            # 商品ページ
│   │
│   ├── public/                  # 静的ファイル
│   │
│   ├── src/                     # ソースコード
│   │   ├── components/          # Reactコンポーネント
│   │   │   ├── layout/          # レイアウトコンポーネント
│   │   │   └── ui/              # UIコンポーネント
│   │   │
│   │   ├── lib/                 # ユーティリティ
│   │   │   └── api/             # APIクライアント
│   │   │
│   │   └── types/               # 型定義
│   │
│   ├── node_modules/            # npmパッケージ（git管理外）
│   └── package.json             # npmパッケージ定義
│
├── frontend-admin/              # 管理画面フロントエンドアプリケーション
│   ├── app/                     # Next.js App Router
│   │
│   ├── public/                  # 静的ファイル
│   │
│   ├── src/                     # ソースコード
│   │   ├── components/          # Reactコンポーネント
│   │   │   ├── layout/          # レイアウトコンポーネント
│   │   │   └── ui/              # UIコンポーネント
│   │   │
│   │   ├── lib/                 # ユーティリティ
│   │   │   ├── api/             # APIクライアント
│   │   │   └── auth/            # 認証関連
│   │   │
│   │   └── types/               # 型定義
│   │
│   ├── node_modules/            # npmパッケージ（git管理外）
│   └── package.json             # npmパッケージ定義
│
├── docs/                        # ドキュメント
│   ├── concept-guide/           # 概念ガイド
│   ├── design/                  # 設計ドキュメント
│   ├── lecture/                 # 講義資料
│   ├── lecture-request/         # 講義リクエスト
│   ├── notion/                  # Notion関連
│   └── template/                # テンプレート
│
├── infra/                       # インフラ関連
│   ├── localstack/              # LocalStack設定
│   │   └── init-scripts/        # 初期化スクリプト
│   │
│   ├── mysql/                   # MySQL設定
│   │   ├── conf.d/              # MySQL設定ファイル
│   │   └── initdb.d/            # 初期化スクリプト
│   │
│   ├── scripts/                 # スクリプト
│   │   └── aws/                 # AWS関連スクリプト
│   │
│   └── traefik/                 # Traefik設定
│       ├── config/              # 静的設定
│       └── dynamic/             # 動的設定
│
├── logs/                        # ログファイル
│   └── traefik/                 # Traefikログ
│
├── prh-rules/                   # 文章校正ルール
│
├── compose.yml                  # Docker Compose設定
├── Taskfile.yml                 # go-task設定
├── .gitignore                   # Git除外設定
└── README.md                    # プロジェクト説明
```

この構造は、関心の分離と層の明確化を意識しています。バックエンド、フロントエンド、インフラがそれぞれ独立したディレクトリに分かれており、マイクロサービスアーキテクチャの考え方を反映しています。
