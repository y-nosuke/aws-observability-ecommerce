# 1. Week 1 - Day 2: バックエンド基本構造の実装

## 1.1. 目次

- [1. Week 1 - Day 2: バックエンド基本構造の実装](#1-week-1---day-2-バックエンド基本構造の実装)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. ディレクトリ構造の作成](#141-ディレクトリ構造の作成)
    - [1.4.2. Go Modulesの初期化](#142-go-modulesの初期化)
    - [1.4.3. Echo Webフレームワークのセットアップ](#143-echo-webフレームワークのセットアップ)
    - [1.4.4. 設定管理モジュールの実装](#144-設定管理モジュールの実装)
    - [1.4.5. ヘルスチェックAPIの実装](#145-ヘルスチェックapiの実装)
    - [1.4.6. Dockerfileの作成](#146-dockerfileの作成)
    - [1.4.7. アプリケーションの起動確認](#147-アプリケーションの起動確認)
    - [1.4.8. 開発を助ける追加ツールの設定](#148-開発を助ける追加ツールの設定)
      - [1.4.8.1. Airの設定（ホットリロード）](#1481-airの設定ホットリロード)
      - [1.4.8.2. golangci-lintの設定](#1482-golangci-lintの設定)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. Echoフレームワークの特徴と利点](#161-echoフレームワークの特徴と利点)
    - [1.6.2. クリーンアーキテクチャと層の分離](#162-クリーンアーキテクチャと層の分離)
    - [1.6.3. ミドルウェアの役割と設計](#163-ミドルウェアの役割と設計)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. Go言語のプロジェクト構成ベストプラクティス](#171-go言語のプロジェクト構成ベストプラクティス)
    - [1.7.2. Viperを使用した設定管理のメリット](#172-viperを使用した設定管理のメリット)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: Goのパッケージ依存関係の解決エラー](#181-問題1-goのパッケージ依存関係の解決エラー)
    - [1.8.2. 問題2: サーバーが起動しない](#182-問題2-サーバーが起動しない)
    - [1.8.3. 問題3: Dockerコンテナが正常に起動しない](#183-問題3-dockerコンテナが正常に起動しない)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)
  - [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- バックエンドアプリケーションの層を分離したクリーンなディレクトリ構造の設計と実装
- Echo Webフレームワークを使用した基本的なWEBサーバー構築
- ロギングやエラーハンドリングなどの共通ミドルウェアの設定
- 環境変数を利用した設定管理モジュールの実装
- シンプルなヘルスチェックAPIの実装によるサーバー動作確認

## 1.3. 【準備】

本日はバックエンドアプリケーションの基本構造を実装します。Go言語を使用して、クリーンアーキテクチャに基づいたディレクトリ構造を作成し、Echo Webフレームワークで基本的なAPIサーバーを構築します。

### 1.3.1. チェックリスト

- [ ] Go言語（バージョン1.18以上）がインストールされていること
- [ ] Docker環境が正常に動作していること
- [ ] Docker Composeがインストールされていること
- [ ] VSCodeなどのコードエディタがセットアップされていること
- [ ] Gitがインストールされていること
- [ ] プロジェクトのルートディレクトリがセットアップされていること
- [ ] GitHub リポジトリが作成されていること（Week 1 - Day 1で作成済みであること）

## 1.4. 【手順】

### 1.4.1. ディレクトリ構造の作成

まず、バックエンドアプリケーションのディレクトリ構造を作成します。クリーンアーキテクチャのプラクティスに従い、関心事の分離を実現する構造にします。

```bash
# バックエンド用のディレクトリ構造を作成
mkdir -p backend/cmd/server
mkdir -p backend/internal/api/handlers
mkdir -p backend/internal/api/middleware
mkdir -p backend/internal/config
mkdir -p backend/pkg/logger
```

この構造は以下の原則に基づいています：

- `cmd/` - アプリケーションのエントリーポイントを格納
- `internal/` - プロジェクト内でのみ使用されるパッケージ
- `internal/api/` - APIとWebサーバー関連のコード
- `internal/config/` - アプリケーション設定関連
- `pkg/` - 再利用可能なパッケージ

### 1.4.2. Go Modulesの初期化

次に、Goのモジュール管理を初期化します。

```bash
# backendディレクトリに移動
cd backend

# Go Modulesの初期化
go mod init github.com/y-nosuke/aws-observability-ecommerce/backend
```

`go.mod`ファイルが作成されたことを確認してください。必要な依存関係を追加します。

```bash
# 必要な依存関係をインストール

# Echo Webフレームワーク
go get -u github.com/labstack/echo/v4
go get -u github.com/labstack/echo/v4/middleware

# 設定管理
go get -u github.com/spf13/viper
```

### 1.4.3. Echo Webフレームワークのセットアップ

まず、メインのエントリーポイントファイルを作成します。

```bash
touch cmd/server/main.go
```

`main.go`の内容を以下のように記述します：

```go
package main

import (
 "context"
 "errors"
 "fmt"
 "log"
 "net/http"
 "os"
 "os/signal"
 "time"

 "github.com/labstack/echo/v4"
 "github.com/labstack/echo/v4/middleware"

 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/api/handlers"
 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/config"
)

func main() {
 // 設定をロード
 if err := config.Load(); err != nil {
  log.Printf("Failed to load configuration: %v\n", err)
  os.Exit(1)
 }

 // Echoインスタンスを作成
 e := echo.New()
 e.HideBanner = true
 e.HidePort = true

 // ミドルウェアの設定
 e.Use(middleware.Recover())
 e.Use(middleware.RequestID())
 e.Use(middleware.Logger()) // 標準のロガーミドルウェアを使用
 e.Use(middleware.CORS())

 // ヘルスチェックエンドポイント
 e.GET("/api/health", func(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]any{
      "status": "ok"
    })
  })

 // コンテキストの初期化（シグナルハンドリング）
 ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
 defer stop()

 // サーバーを起動
 go func() {
  address := fmt.Sprintf(":%d", config.Server.Port)
  if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
   log.Printf("Failed to start server: %v\n", err)
   os.Exit(1)
  }
 }()

 // シグナルを待機
 <-ctx.Done()
 log.Println("Shutdown signal received, gracefully shutting down...")

 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 defer cancel()

 if err := e.Shutdown(ctx); err != nil {
  log.Printf("Failed to shutdown server gracefully: %v\n", err)
 } else {
  log.Printf("Server shutdown gracefully")
 }
}
```

### 1.4.4. 設定管理モジュールの実装

次に、設定管理モジュールを実装します。

```bash
touch internal/config/config.go
```

`config.go`の内容を以下のように記述します：

```go
package config

import (
 "errors"

 "github.com/spf13/viper"
)

// Config はアプリケーション全体の設定を格納する構造体
type Config struct {
 App struct {
  Name        string
  Version     string
  Environment string
 }
 Server struct {
  Port int
 }
}

// アプリケーション設定インスタンス
var (
 config Config
 App    = &config.App
 Server = &config.Server
)

// Load は環境変数と設定ファイルから設定をロードします
func Load() error {
 // 環境変数のデフォルト値の設定
 viper.SetDefault("app.name", "aws-observability-ecommerce")
 viper.SetDefault("app.version", "1.0.0")
 viper.SetDefault("app.environment", "development")
 viper.SetDefault("server.port", "8000")

 // 環境変数のバインド
 if err := viper.BindEnv("app.name", "APP_NAME"); err != nil {
  return err
 }
 if err := viper.BindEnv("app.version", "APP_VERSION"); err != nil {
  return err
 }
 if err := viper.BindEnv("app.environment", "APP_ENV"); err != nil {
  return err
 }
 if err := viper.BindEnv("server.port", "PORT"); err != nil {
  return err
 }
 // 設定ファイルの読み込み（存在する場合）
 viper.SetConfigName("config")
 viper.SetConfigType("yaml")
 viper.AddConfigPath(".")
 viper.AddConfigPath("./config")

 // 設定ファイルの読み込み（存在しなくてもエラーとしない）
 if err := viper.ReadInConfig(); err != nil {
  var configFileNotFoundError viper.ConfigFileNotFoundError
  if !errors.As(err, &configFileNotFoundError) {
   return err
  }
 }

 // viper.Unmarshalを使って設定を一括で読み込む
 if err := viper.Unmarshal(&config); err != nil {
  return err
 }

 return nil
}
```

### 1.4.5. ヘルスチェックAPIの実装

ヘルスチェックハンドラーを専用のファイルに移動しましょう。

```bash
touch internal/api/handlers/health.go
```

`health.go`の内容を以下のように記述します：

```go
package handlers

import (
 "log"
 "net/http"
 "runtime"
 "time"

 "github.com/labstack/echo/v4"

 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/config"
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
  version:   config.App.Version, // アプリケーションバージョン
 }
}

// HandleHealthCheck はヘルスチェックエンドポイントのハンドラー関数
func (h *HealthHandler) HandleHealthCheck(c echo.Context) error {
 // リクエストの処理開始をログに記録
 log.Println("Health check request received",
  "method", c.Request().Method,
  "path", c.Path(),
  "remote_ip", c.RealIP(),
 )

 // サービスの状態をチェック（ここでは簡易的にすべて稼働中とする）
 services := map[string]interface{}{
  "api": map[string]string{
   "name":   config.App.Name,
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
 log.Println("Health check completed",
  "status", response.Status,
  "uptime", response.Uptime,
 )

 return c.JSON(http.StatusOK, response)
}
```

次に、`main.go`を更新して、ハンドラーを使用するように変更します：

```bash
# main.goを編集
```

`main.go`の変更部分（ルートの設定部分を次のように更新）：

```go
// main.go内のルート設定部分を置き換え
import (
 // 既存のインポート
 "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/api/handlers"
)

// ルートの設定部分
// e.GET("/api/health", func(c echo.Context) error { ... }) の代わりに以下を使用
// APIグループ
api := e.Group("/api")

// ハンドラーの作成
healthHandler := handlers.NewHealthHandler()

// ヘルスチェックエンドポイント
api.GET("/health", healthHandler.HandleHealthCheck)
```

### 1.4.6. Dockerfileの作成

バックエンド用のDockerfileを作成します。

```bash
# Dockerfileを作成
touch Dockerfile
```

`Dockerfile`の内容を以下のように記述します：

```Dockerfile
FROM golang:1.24-alpine

WORKDIR /app

# 開発に必要なツールのインストール
RUN apk add --no-cache git curl bash

# airのインストール（ホットリロード用）
RUN go install github.com/air-verse/air@latest

# デバッグツールのインストール
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# モックジェネレーターのインストール
RUN go install go.uber.org/mock/mockgen@latest

# goimportsのインストール
RUN go install golang.org/x/tools/cmd/goimports@latest

# golangci-lintのインストール
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# AWS CLI v2とawscli-localのインストール（LocalStackとの連携用）
RUN apk add --no-cache python3 py3-pip unzip
# 仮想環境を作成してawscli-localをインストール
RUN python3 -m venv /opt/venv && \
    . /opt/venv/bin/activate && \
    pip3 install awscli-local && \
    ln -s /opt/venv/bin/aws /usr/local/bin/aws && \
    ln -s /opt/venv/bin/awslocal /usr/local/bin/awslocal

# タイムゾーンの設定
RUN apk add --no-cache tzdata
ENV TZ=Asia/Tokyo

# アプリケーションの依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# airの設定
COPY .air.toml .

# ポートを公開
EXPOSE 8000

# ホットリロードでアプリケーションを起動
CMD ["air", "-c", ".air.toml"]
```

次に、プロジェクトのルートディレクトリにあるcompose.ymlファイルにbackendサービスの設定を追加します。

`compose.yml`に以下の内容を追加・更新します：

```yaml
services:
  traefik:

・・・

  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "8000:8000"
    environment:
      - APP_NAME=aws-observability-ecommerce
      - APP_VERSION=1.0.0
      - APP_ENV=development
      - PORT=8000
    volumes:
      - ./backend:/app # ホットリロード用
    depends_on:
      mysql:
        condition: service_healthy
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/api/health"]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 10s
    networks:
      - ecommerce-network

・・・
```

この設定では、Dockerfileを使用してバックエンドサービスをビルドし、必要な環境変数を設定しています。また、MySQLサービスに依存関係を設定し、ヘルスチェックを構成しています。

### 1.4.7. アプリケーションの起動確認

Docker Composeを使用してアプリケーションを起動します。プロジェクトのルートディレクトリで以下のコマンドを実行します：

```bash
# Docker Composeでサービスをビルドし起動
docker-compose build backend
docker-compose up -d
```

サービスの起動状況を確認します：

```bash
# サービスの状態確認
docker-compose ps
```

すべてのサービスが起動していることを確認したら、バックエンドのログを確認します：

```bash
# バックエンドのログを表示
docker-compose logs backend
```

ヘルスチェックAPIが正常に動作しているか確認します：

```bash
# ヘルスチェックAPIにリクエストを送信
curl http://localhost:8000/api/health
```

以下のようなレスポンスが返ってくれば成功です：

```json
{
    "status": "ok",
    "timestamp": "2025-03-29T22:00:29+09:00",
    "version": "1.0.0",
    "uptime": 14931,
    "resources": {
        "goroutines": 6,
        "memory": {
            "allocated": 558864,
            "system": 6904848,
            "total": 558864
        }
    },
    "services": {
        "api": {
            "name": "aws-observability-ecommerce",
            "status": "up"
        }
    }
}
```

### 1.4.8. 開発を助ける追加ツールの設定

ここでは、インストールした開発ツールを設定して、開発プロセスをより効率的にします。

#### 1.4.8.1. Airの設定（ホットリロード）

```bash
go install github.com/air-verse/air@latest       # ホットリロードツール
```

Airはコード変更を監視し、変更があると自動的にアプリケーションを再起動するホットリロードツールです。

```bash
# Airの設定ファイルを作成
touch .air.toml
```

`.air.toml`の内容を以下のように設定します：

```toml
# .air.toml
root = "."
tmp_dir = "tmp"

[build]
  bin = "./tmp/main"
  cmd = "goimports -w . && go build -o ./tmp/main ./cmd/server"
  delay = 1000
  exclude_dir = ["tmp", "vendor", "tests"]
  exclude_file = []
  exclude_regex = []
  exclude_unchanged = true
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  kill_delay = "0s"
  log = "build-errors.log"
  send_interrupt = false
  stop_on_error = true

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  time = false

[misc]
  clean_on_exit = false
```

これで、次のコマンドでホットリロード開発が可能になります：

```bash
# Airでホットリロード開発を開始
air
```

#### 1.4.8.2. golangci-lintの設定

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest  # リンター
```

コード品質を維持するためのgolangci-lintの設定ファイルを作成します。

```bash
# golangci-lintの設定ファイルを作成
touch .golangci.yml
```

`.golangci.yml`の内容を以下のように設定します：

```yaml
run:
  # タイムアウト時間（デフォルト: 1m）
  timeout: 5m

# 有効にするリンターの設定
linters:
  disable-all: true
  enable:
    - errcheck # エラーハンドリングの検証
    - gosimple # コードシンプル化の提案
    - govet # Goの怪しい構造を検出
    - ineffassign # 使用されない代入の検出
    - staticcheck # Go用の静的解析ツール
    - unused # 未使用コードの検出
    - gofmt # 標準フォーマッタによるフォーマット
    - goimports # インポート文の整理
    - misspell # よくあるスペルミスの検出
    - unconvert # 不要な型変換の削除
    - gosec # セキュリティの問題を検出
    - bodyclose # レスポンスボディを適切に閉じているか確認

# 特定のリンターの設定
linters-settings:
  errcheck:
    # io/ioutil.WriteFile, os.Create などの特定のエラーを無視する
    check-type-assertions: true
    check-blank: true

  govet:
    # 最新バージョンではcheck-shadowingはサポートされていないため削除
    enable:
      - shadow # シャドウ変数を検出する

  goimports:
    # インポートグループの順序を指定
    local-prefixes: github.com/y-nosuke/aws-observability-ecommerce

  gosec:
    # セキュリティの問題を検出する重大度
    severity: medium
    confidence: medium

# 発見した問題を除外する設定
issues:
  # 最大表示問題数（デフォルト: 50）
  max-issues-per-linter: 0
  max-same-issues: 0

  # 特定のパターンで問題を除外
  exclude-rules:
    - path: _test\.go
      linters:
        - errcheck # テストファイルではerrcheckを無効化

    - text: "G404: Use of weak random number generator"
      linters:
        - gosec # テスト用途などでの弱い乱数生成は許容

  # 解析から除外するディレクトリやファイル
  exclude-files:
    - ".*_test.go$"
  exclude-dirs-use-default: true
```

lintを実行するには、次のコマンドを実行します：

```bash
# lintを実行
golangci-lint run ./...

# 修正
golangci-lint run --fix ./...
```

## 1.5. 【確認ポイント】

実装が正しく完了したことを確認するためのチェックリストです：

- [ ] バックエンドのディレクトリ構造が正しく作成されている
- [ ] Go Modulesが初期化され、必要な依存関係がインストールされている
- [ ] Echo Webフレームワークが正しく設定されている
- [ ] 設定管理モジュールが正しく実装されている
- [ ] ヘルスチェックAPIが`/api/health`エンドポイントで応答する
- [ ] ログが構造化JSONフォーマットで出力される
- [ ] Dockerfileが正しく作成されている
- [ ] Docker Compose環境でアプリケーションが正常に起動し、ヘルスチェックAPIが応答する
- [ ] 開発ツール（Airホットリロード、golangci-lintなど）が正しく設定されている

## 1.6. 【詳細解説】

### 1.6.1. Echoフレームワークの特徴と利点

[Echo](https://echo.labstack.com/)は、Goで書かれた高性能でミニマリストなWebフレームワークです。以下のような特徴と利点があります：

1. **データバインディング**：リクエストボディ（JSON、XML、フォームなど）を構造体に自動的にバインドする機能を提供します。

2. **パスパラメータとクエリパラメータの処理**：URLパスパラメータやクエリパラメータを簡単に取得するためのAPIを提供します。

3. **コンテキスト管理**：リクエスト処理中にデータを渡すためのコンテキストシステムを提供します。

本実装では、Echo を使用して基本的なAPIサーバーを構築し、ミドルウェアを活用してロギングやエラーハンドリングを実装しています。また、グレースフルシャットダウンの実装により、実行中のリクエストが正常に完了してからサーバーを停止できるようにしています。

### 1.6.2. クリーンアーキテクチャと層の分離

今回実装したディレクトリ構造は、クリーンアーキテクチャの原則に従っています。クリーンアーキテクチャの主な目的は、ソフトウェアの関心事を分離し、依存関係を制御することです。

ディレクトリ構造の各部分の役割と目的：

1. **cmd/server**: アプリケーションのエントリーポイントを含みます。ここでは、設定の読み込み、依存関係の注入、サーバーの起動などの初期化処理を行います。

2. **internal**: プロジェクト内でのみ使用されるパッケージを含みます。Goのビルドシステムは`internal`ディレクトリ内のコードが外部から直接インポートされることを防ぎます。

3. **internal/api**: API関連のコードを格納します。`handlers`サブディレクトリにはHTTPハンドラー、`middleware`サブディレクトリにはカスタムミドルウェアを配置します。

4. **internal/config**: アプリケーション設定の読み込みと管理を担当します。

5. **pkg**: 再利用可能な、他のプロジェクトからもインポートできるパッケージを含みます。例えば、ロガーのような共通ユーティリティはここに配置します。

この構造により、テスト容易性、保守性、拡張性に優れたコードベースを構築できます。また、ビジネスロジックをインフラストラクチャの詳細（データベース、外部API、UIなど）から分離できます。

### 1.6.3. ミドルウェアの役割と設計

Echoのミドルウェアは、HTTPリクエスト/レスポンスのパイプラインに挿入される関数です。リクエストの処理前、処理中、処理後に特定の動作を実行できます。

本実装では以下のミドルウェアを使用しています：

1. **Recover**: パニックをキャッチし、500エラーレスポンスを返します。これにより、サーバーがクラッシュするのを防ぎます。

2. **RequestID**: 各リクエストに一意のID（UUID）を割り当てます。これは、分散システムでのトレーサビリティに重要です。

3. **Logger**: リクエストとレスポンスの情報をログに記録します。HTTPメソッド、URL、ステータスコード、レスポンスタイムなどが記録されます。

4. **CORS**: Cross-Origin Resource Sharing (CORS) ヘッダーを設定します。これにより、異なるオリジンからのAPIリクエストを許可できます。

これらのミドルウェアは、アプリケーションの横断的関心事を処理します。例えば、ロギングはすべてのリクエストに適用すべき横断的関心事です。ミドルウェアを使用することで、これらの関心事をハンドラーロジックから分離し、コードの重複を減らすことができます。

## 1.7. 【補足情報】

### 1.7.1. Go言語のプロジェクト構成ベストプラクティス

Goプロジェクトの構成に関しては、いくつかのベストプラクティスがあります：

1. **標準レイアウト**: [Standard Go Project Layout](https://github.com/golang-standards/project-layout) は、多くのGoプロジェクトで採用されている標準的なレイアウトです。本実装もこれに基づいています。

2. **ドメイン駆動設計 (DDD)**: 大規模なプロジェクトでは、ドメイン駆動設計に基づいてパッケージを構成することも効果的です。例えば、`internal/domain/product`、`internal/domain/order`のようにビジネスドメインごとにパッケージを分けます。

3. **インターフェース指向**: Goはインターフェースを中心とした設計が得意です。具体的な実装よりもインターフェースに依存するようにコードを設計すると、テスト容易性や拡張性が高まります。

4. **小さいパッケージ**: 大きな単一のパッケージよりも、小さく焦点を絞ったパッケージに分割する方が好ましいです。1つのパッケージの責任範囲を明確に限定することで、コードの理解や保守が容易になります。

5. **テスト配置**: テストファイルは、テスト対象のコードと同じパッケージに配置します。例えば、`handlers/health.go`のテストは`handlers/health_test.go`になります。

これらのプラクティスは、プロジェクトの要件や規模によって適応させる必要があります。小規模なプロジェクトでは、過度に複雑な構造を避け、シンプルさを保つことも重要です。

### 1.7.2. Viperを使用した設定管理のメリット

今回の実装では、Viperを使用して設定を管理しています。この方法には以下のメリットがあります：

1. **複数の設定ソースのサポート**: Viperは環境変数だけでなく、YAML、JSON、TOML、HCLなど様々な形式の設定ファイルもサポートしています。これにより、異なる環境や状況に応じて最適な設定方法を選択できます。

2. **12 Factor App の原則への準拠**: [The Twelve-Factor App](https://12factor.net/config) の原則では、設定を環境変数として管理することを推奨していますが、Viperはこれをさらに拡張し、より柔軟な設定管理を可能にします。

3. **設定の階層化**: Viperは階層的な設定構造をサポートしており、複雑な設定でも整理された形で管理できます。

4. **設定の動的な再読み込み**: Viperは設定ファイルの変更を監視し、アプリケーションの再起動なしに設定を再読み込みする機能を提供します。

5. **デフォルト値の設定**: 設定値が見つからない場合のデフォルト値を明示的に設定できます。

6. **複数の環境への対応**: 開発、テスト、本番など異なる環境に対して、同じコードベースで異なる設定を適用することが容易です。

7. **機密情報の保護**: APIキーやパスワードなどの機密情報をコードリポジトリに保存せず、設定ファイルや環境変数として管理できます。

この設計により、アプリケーションの設定管理がより柔軟かつ強力になり、メンテナンス性と拡張性が向上します。また、将来的に設定ソースや形式を変更する必要が生じても、アプリケーションコードを変更せずに対応できます。

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: Goのパッケージ依存関係の解決エラー

**症状**: `go build`や`go run`を実行した際に「package xxx is not in GOROOT」などのエラーが発生する。

**解決策**:

1. `go mod tidy`を実行して、依存関係を整理します。

    ```bash
    go mod tidy
    ```

2. インポートパスが正しいか確認します。特に、プロジェクト内のパッケージをインポートする際に、モジュール名が正しいか確認します。

    ```go
    // 正しいインポート例
    import "github.com/y-nosuke/aws-observability-ecommerce/backend/internal/config"
    ```

3. GOPATHやGOROOTの設定が適切か確認します。

    ```bash
    go env
    ```

### 1.8.2. 問題2: サーバーが起動しない

**症状**: サーバーを起動しようとすると「address already in use」などのエラーが発生する。

**解決策**:

1. 指定したポートが既に使用されていないか確認します。

    ```bash
    # Linuxの場合
    netstat -tuln | grep 8000

    # Windowsの場合
    netstat -ano | findstr 8000
    ```

2. 既に使用されている場合は、そのプロセスを終了するか、別のポートを使用します。

    ```bash
    # ポートを使用しているプロセスを終了（Linux）
    kill -9 $(lsof -t -i:8000)

    # 環境変数で別のポートを指定
    export PORT=8081
    ```

3. サーバーの設定やルーティングに問題がないか確認します。

### 1.8.3. 問題3: Dockerコンテナが正常に起動しない

**症状**: `docker-compose up`を実行してもバックエンドコンテナが起動せず、ログにエラーが表示される。

**解決策**:

1. Dockerfileが正しいか確認します。特に、ビルドコマンドやエントリーポイントの設定を確認します。

2. compose.ymlファイルのサービス定義を確認します。環境変数やポートマッピングが適切に設定されているか確認します。

3. Dockerイメージを再ビルドします。

    ```bash
    docker-compose build --no-cache backend
    ```

4. Dockerのログを確認して、具体的なエラー内容を特定します。

    ```bash
    docker-compose logs backend
    ```

5. コンテナ内で直接コマンドを実行して、問題を診断します。

    ```bash
    docker-compose exec backend sh
    ```

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **クリーンなアーキテクチャ設計**: バックエンドのディレクトリ構造を、関心事の分離と依存関係の制御を意識して設計しました。これにより、コードの保守性とテスト容易性が向上します。

2. **ミドルウェアによる横断的関心事の分離**: ロギング、エラーハンドリング、CORSなどの横断的関心事をミドルウェアとして実装することで、ハンドラーコードをシンプルに保ちながら、共通の機能を提供できます。

3. **環境変数を使用した設定管理**: 設定を環境変数として外部化することで、異なる環境でのデプロイが容易になり、コードとインフラストラクチャの結合度を低くできます。

4. **グレースフルシャットダウンの実装**: サーバーの停止時に進行中のリクエストを正常に完了させる仕組みを実装することで、サービスの信頼性が向上します。

5. **構造化ログの導入**: slogを使用した構造化ログを導入することで、ログの解析と可視化が容易になります。これは、後のフェーズでのオブザーバビリティ強化の基盤となります。

これらのポイントは、高品質なバックエンドアプリケーションを構築するための基本です。次回以降の実装でも、これらの原則を継続的に適用していきます。

## 1.10. 【次回の準備】

次回（Day 3）では、データベーススキーマの設計と実装に取り組みます。以下の点について事前に確認しておくと良いでしょう：

1. SQLの基本知識と構文（CREATE TABLE, ALTER TABLE, DROP TABLEなど）
2. データベースのリレーションシップ（One-to-Many, Many-to-Manyなど）
3. インデックスの概念と使用するタイミング
4. MySQLの基本的な動作とデータタイプ
5. golang-migrateの基本的な使用方法
6. sqlboilerの基本的な概念とコード生成の仕組み

また、以下のドキュメントを事前に読んでおくことをお勧めします：

- [MySQL Documentation](https://dev.mysql.com/doc/)
- [golang-migrate GitHub](https://github.com/golang-migrate/migrate)
- [sqlboiler Documentation](https://github.com/volatiletech/sqlboiler)

## 1.11. 【.envrc サンプル】

以下は本日の実装で使用する.envrcのサンプルです。ご自身の環境に合わせて修正して使用してください。このファイルはgitにコミットしないようにしてください。

```bash
# .envrc サンプル - このファイルをバックエンドディレクトリに.envrcとして保存しgitにコミットしないでください
export APP_NAME=development
export APP_VERSION=development
export APP_ENV=development
export PORT=8000

# 後で使用するデータベース関連の環境変数
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=password
export DB_NAME=ecommerce

# direnvを使用している場合は、以下のコマンドで有効化してください
# direnv allow .
```
