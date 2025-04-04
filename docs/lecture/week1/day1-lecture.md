# 1. Week 1 - Day 1: 開発環境とGitHubリポジトリの設定

## 1.1. 目次

- [1. Week 1 - Day 1: 開発環境とGitHubリポジトリの設定](#1-week-1---day-1-開発環境とgithubリポジトリの設定)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. GitHubリポジトリのセットアップ](#141-githubリポジトリのセットアップ)
    - [1.4.2. ブランチ戦略とprotected branchの設定](#142-ブランチ戦略とprotected-branchの設定)
    - [1.4.3. プロジェクト構造の作成](#143-プロジェクト構造の作成)
    - [1.4.4. Docker Compose 設定ファイルの作成](#144-docker-compose-設定ファイルの作成)
    - [1.4.5. MySQL の設定](#145-mysql-の設定)
      - [1.4.5.1. 設定ファイル権限の設定方法](#1451-設定ファイル権限の設定方法)
        - [1.4.5.1.1. Linux/macOSの場合](#14511-linuxmacosの場合)
        - [1.4.5.1.2. Windowsの場合（WSL使用時も含む）](#14512-windowsの場合wsl使用時も含む)
    - [1.4.6. Docker Compose 環境の起動と検証](#146-docker-compose-環境の起動と検証)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. Docker Compose の役割と利点](#161-docker-compose-の役割と利点)
    - [1.6.2. GitHubのブランチ保護と開発ワークフロー](#162-githubのブランチ保護と開発ワークフロー)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. Docker ネットワークについて](#171-docker-ネットワークについて)
    - [1.7.2. ボリュームによるデータ永続化](#172-ボリュームによるデータ永続化)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: ポートの競合](#181-問題1-ポートの競合)
    - [1.8.2. 問題2: MySQLの接続エラー](#182-問題2-mysqlの接続エラー)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)
  - [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- GitHubリポジトリを作成し、適切なブランチ戦略とprotected branchを設定する
- Docker Compose を使用してMySQLデータベースを含む開発環境を構築する
- プロジェクトの基本的なディレクトリ構造を設定する
- データベースの永続化とネットワーク設定を適切に構成する
- コード品質と開発フローを維持するためのGitHub設定を行う

## 1.3. 【準備】

このプロジェクトを始めるにあたり、以下のツールとソフトウェアが必要です。実装を始める前に、すべてのツールが正しくインストールされていることを確認してください。

### 1.3.1. チェックリスト

- [x] Git (バージョン管理)

  ```bash
  git --version
  # git version 2.34.1 以上が望ましい
  ```

- [x] GitHub アカウント

  ```bash
  # GitHubアカウントを持っていることを確認
  # https://github.com/
  ```

- [x] Docker Engine

  ```bash
  docker --version
  # Docker version 20.10.0 以上が望ましい
  ```

- [x] Docker Compose

  ```bash
  docker compose version
  # Docker Compose version v2.10.0 以上が望ましい
  ```

- [x] テキストエディタまたはIDE (Visual Studio Code推奨)
- [x] ターミナル (Linuxベースであればどれでも可)
- [x] curl または wget (動作確認用)

  ```bash
  curl --version
  # curl 7.68.0 以上が望ましい
  ```

- [x] direnv (オプション、環境変数管理用)

  ```bash
  direnv --version
  # direnv v2.32.0 以上が望ましい
  ```

## 1.4. 【手順】

### 1.4.1. GitHubリポジトリのセットアップ

まず、プロジェクト用のGitHubリポジトリを作成します。

1. GitHubにログインし、右上の「+」アイコンから「New repository」を選択します。

2. リポジトリ設定を行います：
   - Repository name: `aws-observability-ecommerce`
   - Description: `AWSオブザーバビリティ学習用のeコマースアプリケーション`
   - Visibility: `Private` または `Public` (学習用途に応じて選択)
   - README初期化: チェックを入れる
   - .gitignore: `Go` を選択
   - License: `MIT License` を選択

3. 「Create repository」ボタンをクリックして、リポジトリを作成します。

4. ローカル環境にリポジトリをクローンします：

    ```bash
    git clone https://github.com/あなたのユーザー名/aws-observability-ecommerce.git
    cd aws-observability-ecommerce
    ```

5. 追加の.gitignore項目を設定します：

    ```bash
    # .gitignoreファイルを編集
    cat << EOF >> .gitignore
    # 環境変数
    .env
    .envrc

    # エディタの設定
    .vscode/
    .idea/

    # ビルド成果物
    bin/
    build/
    dist/
    tmp/

    # データベースファイル
    *.db
    *.sqlite

    # ログファイル
    *.log
    logs/

    # システム固有のファイル
    .DS_Store
    Thumbs.db
    EOF
    ```

### 1.4.2. ブランチ戦略とprotected branchの設定

GitHubでブランチ保護設定を行い、安全な開発ワークフローを確立します。

1. GitHubリポジトリページで「Settings」タブを開きます。

2. 左側のメニューから「Branches」を選択します。

3. 「Branch protection rules」セクションで「Add rule」ボタンをクリックします。

4. 以下の設定を行います：
   - Branch name pattern: `main`
   - Require pull request reviews before merging: チェックを入れる
     - Required approving reviews: `1`
   - Require status checks to pass before merging: チェックを入れる
   - Require branches to be up to date before merging: チェックを入れる
   - Include administrators: オプションでチェックを入れる
   - Allow force pushes: チェックを外す
   - Allow deletions: チェックを外す

5. 「Create」ボタンをクリックして、保護ルールを作成します。

6. 開発用の`develop`ブランチを作成します：

    ```bash
    git checkout -b develop
    git push -u origin develop
    ```

7. 必要に応じて、`develop`ブランチにも同様の保護ルールを設定します。

### 1.4.3. プロジェクト構造の作成

プロジェクト用の基本的なディレクトリ構造をセットアップします。

```bash
# 既にクローンしたリポジトリ内で作業
cd aws-observability-ecommerce

# Dockerおよび環境関連のディレクトリを作成
mkdir -p infra/mysql/{initdb.d,conf.d}

# バックエンドとフロントエンド用のディレクトリを作成
mkdir -p backend/{cmd,internal,pkg,api}
mkdir -p frontend-customer
mkdir -p frontend-admin

# 各種設定ファイル用のディレクトリを作成
mkdir -p .github/workflows
```

### 1.4.4. Docker Compose 設定ファイルの作成

プロジェクトのルートディレクトリに `compose.yml` ファイルを作成し、MySQL サービスを定義します。

```bash
# Docker Compose 設定ファイルを作成
touch compose.yml
```

`compose.yml` に以下の内容を記述します：

```yaml
services:
  mysql:
    image: mysql:latest
    container_name: mysql
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD:-rootpassword}
      MYSQL_DATABASE: ${MYSQL_DATABASE:-ecommerce}
      MYSQL_USER: ${MYSQL_USER:-ecommerce_user}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD:-ecommerce_password}
    ports:
      - "3306:3306"
    volumes:
      - ./infra/mysql/initdb.d:/docker-entrypoint-initdb.d:ro
      - ./infra/mysql/conf.d:/etc/mysql/conf.d:ro
      - mysql_data:/var/lib/mysql
    healthcheck:
      test:
        [
          "CMD",
          "mysqladmin",
          "ping",
          "-h",
          "localhost",
          "-u",
          "root",
          "-p${MYSQL_ROOT_PASSWORD:-rootpassword}",
        ]
      interval: 5s
      timeout: 5s
      retries: 5
      start_period: 10s
    networks:
      - ecommerce-network
    deploy:
      resources:
        limits:
          memory: 512M

  phpmyadmin:
    image: phpmyadmin/phpmyadmin
    container_name: phpmyadmin
    restart: always
    ports:
      - "8080:80"
    environment:
      - PMA_HOST=mysql
      - PMA_USER=ecommerce_user
      - PMA_PASSWORD=ecommerce_password
      - UPLOAD_LIMIT=300M
    depends_on:
      - mysql
    networks:
      - ecommerce-network

volumes:
  mysql_data:
    driver: local

networks:
  ecommerce-network:
    driver: bridge
    name: ecommerce-network
```

### 1.4.5. MySQL の設定

MySQL の初期化スクリプトと設定ファイルを作成します。

```bash
# MySQL 初期化スクリプトと設定ファイルを作成
touch infra/mysql/initdb.d/01_init.sql
touch infra/mysql/conf.d/my.cnf
```

`infra/mysql/initdb.d/01_init.sql` に以下の内容を記述します：

```sql
-- 基本的な初期化スクリプト
-- 詳細なテーブル定義は Week 2 で実装します

-- データベースが存在しない場合は作成
CREATE DATABASE IF NOT EXISTS `ecommerce`;

-- 権限の設定
GRANT ALL PRIVILEGES ON `ecommerce`.* TO 'ecommerce_user'@'%';
FLUSH PRIVILEGES;

-- ecommerceデータベースを選択
USE `ecommerce`;

-- 基本的なテストテーブル作成
CREATE TABLE IF NOT EXISTS `test` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- テストデータ挿入
INSERT INTO `test` (`name`) VALUES ('This is a test');
```

`infra/mysql/conf.d/my.cnf` に以下の内容を記述します：

```ini
[mysqld]
character-set-server=utf8mb4
collation-server=utf8mb4_0900_ai_ci
default-time-zone='+09:00'

[mysql]
default-character-set=utf8mb4

[client]
default-character-set=utf8mb4
```

#### 1.4.5.1. 設定ファイル権限の設定方法

MySQLの設定ファイルを作成したら、適切な権限を設定する必要があります。環境によって設定方法が異なります。

##### 1.4.5.1.1. Linux/macOSの場合

```bash
# 設定ファイルの権限を変更（所有者のみ書き込み可能、他は読み取り専用）
chmod 644 infra/mysql/conf.d/my.cnf
```

##### 1.4.5.1.2. Windowsの場合（WSL使用時も含む）

Windows環境またはWSL環境でWindowsファイルシステム上のファイルを操作する場合：

1. エクスプローラーでファイルを右クリック
2. 「プロパティ」を選択
3. 「読み取り専用」にチェックを入れて「OK」をクリック

これは、MySQLがセキュリティ上の理由から誰でも書き込み可能な設定ファイルを無視するため、重要なステップです。権限が適切に設定されていないと、以下のような警告が表示されることがあります：

```text
mysqld: [Warning] World-writable config file '/etc/mysql/conf.d/my.cnf' is ignored.
```

### 1.4.6. Docker Compose 環境の起動と検証

作成した Docker Compose 環境を起動し、正常に動作することを確認します。

```bash
# Docker Compose 環境の起動
docker compose up -d

# サービスの状態確認
docker compose ps
```

[1.5. 【確認ポイント】](#15-確認ポイント)を実施します。

使用が終わったら、サービスを停止します：

```bash
# サービスの停止
docker compose down
```

ボリュームも含めて完全にクリーンアップする場合は次のコマンドを使用します：

```bash
# ボリュームも含めて削除
docker compose down -v
```

## 1.5. 【確認ポイント】

Docker Compose 環境とGitHubリポジトリが正しくセットアップされたことを確認するためのチェックリストです：

- [x] GitHubリポジトリが正常に作成され、ローカルにクローンされている

  ```bash
  # リモートリポジトリの確認
  git remote -v
  # origin  https://github.com/あなたのユーザー名/aws-observability-ecommerce.git (fetch)
  # origin  https://github.com/あなたのユーザー名/aws-observability-ecommerce.git (push)
  ```

- [x] ブランチ保護ルールが正しく設定されている

  ```bash
  # GitHubのSettings > Branchesページで確認
  # mainとdevelopブランチに保護ルールが適用されていることを確認
  ```

- [x] Docker Composeのコンテナが起動し、ステータスが「Up」になっている

  ```bash
  $ docker compose ps
  NAME         IMAGE                       COMMAND                  SERVICE      CREATED         STATUS                    PORTS
  mysql        mysql:latest                "docker-entrypoint.s…"   mysql        9 minutes ago   Up 9 minutes (healthy)   0.0.0.0:3306->3306/tcp, 33060/tcp
  phpmyadmin   phpmyadmin/phpmyadmin       "/docker-entrypoint.…"   phpmyadmin   9 minutes ago   Up 9 minutes             0.0.0.0:8080->80/tcp
  ```

- [x] MySQL コンテナに接続できる

  ```bash
  $ docker compose exec mysql mysql -uecommerce_user -pecommerce_password -e "SELECT * FROM ecommerce.test;"
  mysql: [Warning] Using a password on the command line interface can be insecure.
  +----+----------------+---------------------+
  | id | name           | created_at          |
  +----+----------------+---------------------+
  |  1 | This is a test | 2025-03-29 14:30:17 |
  +----+----------------+---------------------+
  # 「This is a test」というデータが表示されればOK
  ```

- [x] phpMyAdminにアクセスできる

  ```bash
  # ブラウザで以下のURLにアクセス
  # http://localhost:8080

  # ログイン情報
  # サーバー: mysql
  # ユーザー名: ecommerce_user
  # パスワード: ecommerce_password
  ```

- [x] プロジェクトのディレクトリ構造が正しく作成されている

  ```bash
  # ディレクトリ構造の確認
  ls -la
  # .github/、backend/、frontend-*/、infra/などのディレクトリが存在することを確認
  ```

## 1.6. 【詳細解説】

### 1.6.1. Docker Compose の役割と利点

Docker Compose は、複数のコンテナを定義し実行するためのツールです。本プロジェクトでは、以下の利点を活かして開発環境を構築しています：

1. **依存関係の明確化**: 各サービス間の依存関係を`compose.yml`ファイル内で明示的に定義できます。今回のプロジェクトでは、phpMyAdminがMySQLに依存するという関係を表現しています。

2. **環境の一貫性**: 開発チームの全員が同じ環境で作業できるようになります。「私の環境では動くのに」という問題を避けることができます。

3. **簡単な起動と停止**: `docker compose up` と `docker compose down` コマンドだけで、すべてのサービスをまとめて起動・停止できます。

4. **環境変数の統合**: `.env` ファイルや環境変数を利用して、設定値を柔軟に変更できます。セキュリティ上重要な情報（パスワードなど）を設定ファイルから分離できます。

5. **ボリュームによるデータ永続化**: `volumes` セクションで定義したように、コンテナが削除されてもデータを保持できます。

Docker Compose は、本番環境での使用は想定されていませんが、開発やテスト環境としては非常に有用です。本プロジェクトでは、開発中の効率化とフェーズ6での本番デプロイに向けた準備として使用しています。

### 1.6.2. GitHubのブランチ保護と開発ワークフロー

GitHubのブランチ保護機能を活用することで、チーム開発における品質とセキュリティを確保できます。本プロジェクトでは、以下の開発ワークフローとブランチ戦略を採用しています：

1. **Gitflow ワークフロー**: 主に2つの主要ブランチ（`main`と`develop`）を使用します。
   - `main`: 本番環境に対応する安定したコード
   - `develop`: 開発中の最新コード
   - 機能開発は`feature/機能名`ブランチで行い、完了後に`develop`へマージ
   - リリース準備は`release/バージョン`ブランチで行い、完了後に`main`と`develop`の両方へマージ
   - 緊急修正は`hotfix/修正内容`ブランチで行い、完了後に`main`と`develop`の両方へマージ

2. **Protected Branch（保護されたブランチ）**: `main`と`develop`ブランチを保護することで、以下の利点があります。
   - 直接のコミットを禁止し、Pull Requestを通じた変更のみを許可
   - コードレビューの強制
   - CI/CDテストの成功を必須とする
   - マージ前の最新状態への更新を強制
   - 履歴の書き換え（force push）を防止

3. **Pull Request (PR) プロセス**:
   - 新機能や修正は、適切な命名規則に従ったブランチで開発
   - 開発完了後、`develop`ブランチへのPRを作成
   - コードレビューとCIテストのパス
   - 承認後、PRをマージ

4. **レビュープロセス**:
   - PRのレビューでは、コード品質、テストカバレッジ、ドキュメンテーションなどを確認
   - コメントやフィードバックを通じて改善点を指摘
   - 必要な修正が完了し、レビュアーが承認すると、マージ可能に

5. **自動化とCI/CD**:
   - GitHub Actionsを使用して、PR時に自動テストやリンターを実行
   - テストが失敗した場合、PRのマージがブロックされる
   - コード品質の基準を満たすことを保証

この開発ワークフローを採用することで、以下のメリットがあります：

- コードの品質維持
- 履歴の一貫性と追跡可能性
- チームコラボレーションの促進
- 本番環境のコードの安定性確保

後のフェーズでは、このワークフローにコード静的解析ツール、セキュリティチェック、自動デプロイなどを追加して、より堅牢な開発プロセスを構築していきます。

## 1.7. 【補足情報】

### 1.7.1. Docker ネットワークについて

Docker Compose 設定では、サービスが `ecommerce-network` という名前のカスタムブリッジネットワークに接続されています。これにより、以下のメリットがあります：

1. **サービス名による名前解決**: 同じネットワーク内のサービスは、サービス名でお互いを参照できます。例えば、phpMyAdminからMySQLに接続する場合、ホスト名として `mysql` を使用できます。

2. **ネットワークの分離**: カスタムネットワークを使用することで、プロジェクト外の他のDockerコンテナと分離できます。

3. **セキュリティの向上**: 公開する必要のないポートを外部に公開せず、同じネットワーク内のサービスだけがアクセスできるようにできます。

Docker ネットワークの詳細情報を確認するには、以下のコマンドを使用できます：

```bash
# ネットワーク一覧を表示
docker network ls

# ecommerce-networkの詳細情報を表示
docker network inspect ecommerce-network
```

将来、バックエンドとフロントエンドのサービスを追加する際には、この同じネットワークに接続することで、サービス間の通信が容易になります。

### 1.7.2. ボリュームによるデータ永続化

Docker コンテナ自体は一時的なものであり、コンテナが削除されると内部のデータも失われます。これを防ぐために、Docker Compose 設定では名前付きボリューム `mysql_data` を使用しています。

これによる主なメリットは以下の通りです：

1. **データの永続化**: コンテナを再作成しても、MySQLのデータは失われません。
2. **パフォーマンス**: 名前付きボリュームは、バインドマウントよりも一般的にパフォーマンスが良いです。
3. **バックアップの容易さ**: ボリュームのデータを簡単にバックアップできます。

ボリュームの詳細情報を確認するには、以下のコマンドを使用できます：

```bash
# ボリューム一覧を表示
docker volume ls

# mysql_dataボリュームの詳細情報を表示
docker volume inspect mysql_data
```

ボリュームをバックアップするには、データをコンテナ外に取り出す必要があります：

```bash
# MySQLデータのバックアップ例
docker run --rm -v aws-observability-ecommerce_mysql_data:/source -v $(pwd)/backup:/backup alpine tar -czvf /backup/mysql_data_backup.tar.gz -C /source .
```

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: ポートの競合

**症状**: Docker Compose の起動時に `Error starting userland proxy: listen tcp 0.0.0.0:3306: bind: address already in use` のようなエラーが表示される。

**解決策**:

1. 競合しているポートを使用しているプロセスを特定します：

   ```bash
   # Linuxの場合
   sudo lsof -i :3306

   # macOSの場合
   sudo lsof -i :3306

   # Windowsの場合
   netstat -aon | findstr :3306
   ```

2. 競合しているプロセスを停止するか、Docker Compose の設定で使用するポートを変更します：

   ```yaml
   # compose.ymlの該当部分を変更
   ports:
     - "3307:3306"  # ローカルの3307ポートをコンテナの3306ポートにマッピング
   ```

3. 変更後、Docker Compose を再起動します：

   ```bash
   docker compose down
   docker compose up -d
   ```

### 1.8.2. 問題2: MySQLの接続エラー

**症状**: MySQLに接続できない、または `Access denied for user` というエラーが発生する。

**解決策**:

1. 環境変数が正しく設定されているか確認します：

   ```bash
   # .envファイルを確認
   cat .env
   # または環境変数を確認
   echo $MYSQL_USER
   echo $MYSQL_PASSWORD
   ```

2. MySQLコンテナが正常に起動しているか確認します：

   ```bash
   docker compose ps
   docker compose logs mysql
   ```

3. 認証情報を明示的に指定して接続を試みます：

   ```bash
   docker compose exec mysql mysql -u ecommerce_user -pecommerce_password ecommerce
   ```

4. 必要に応じてMySQLコンテナを再作成します：

   ```bash
   docker compose down -v  # データも削除したい場合
   docker compose up -d
   ```

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **GitHubリポジトリとブランチ保護**: 安全で効率的な開発ワークフローの基盤を構築しました。これにより、コード品質の維持とチームでの協業が容易になります。

2. **基本的なプロジェクト構造**: フロントエンド、バックエンド、インフラの明確な分離を持つディレクトリ構造を設定しました。これはプロジェクトの拡張性と保守性に貢献します。

3. **Docker Compose による環境の統合**: MySQLなどの開発依存サービスを一元管理し、開発環境を簡単に再現できるようになりました。これにより、チーム全体で一貫した開発が可能になります。

4. **持続可能な開発環境**: ボリュームによるデータ永続化とサービス間のネットワーク接続を設定することで、長期的な開発に適した環境を構築しました。

これらのポイントは、次回以降の実装においても基盤となる重要な概念です。特に、GitHubワークフローとDocker環境の理解はプロジェクト全体を通じて必要となります。

## 1.10. 【次回の準備】

次回（Day 2）では、バックエンドの基本構造を実装します。以下の点について事前に確認しておくと良いでしょう：

1. **Go言語の基本**: Go言語の基本的な構文や概念を確認しておくと、スムーズに進められます。

2. **Echoフレームワーク**: Go言語のWebフレームワークであるEchoの基本概念や使い方を確認しておくと良いでしょう。

3. **Docker環境の動作確認**: 今回構築したDocker環境が正常に動作していることを確認しておきましょう。次回もこの環境をベースに開発を進めます。

4. **依存関係管理ツール**: Goのモジュール管理（go.modとgo.sum）について理解しておくと良いでしょう。

次回はこれらの知識をベースに、バックエンドの基本構造を実装していきます。

## 1.11. 【.envrc サンプル】

以下は本日の実装で使用する.envrcのサンプルです。ご自身の環境に合わせて修正して使用してください。このファイルはgitにコミットしないようにしてください。

```bash
# .envrc サンプル
# direnvがインストールされている場合、このディレクトリに入ると自動的に環境変数が設定されます
# このファイルはgitにコミットしないでください

# MySQL設定
export MYSQL_ROOT_PASSWORD=rootpassword
export MYSQL_DATABASE=ecommerce
export MYSQL_USER=ecommerce_user
export MYSQL_PASSWORD=ecommerce_password

# 開発環境設定
export ENVIRONMENT=development
```

.gitignoreファイルに.envrcを追加して、誤ってコミットしないようにしましょう：

```bash
echo ".envrc" >> .gitignore
```
