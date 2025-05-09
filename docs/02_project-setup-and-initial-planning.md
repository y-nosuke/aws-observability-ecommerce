# プロジェクトの準備と初期計画 (開発開始前)

アジャイルな開発スプリントを開始する前に、プロダクトの全体像を理解し、開発の方向性を定めるための準備と初期計画を行います。これらの活動は、プロジェクトの初期段階で集中的に行われ、その後の開発の基盤となります。

## 1. 要求定義 (プロダクトの全体像と価値の明確化)

このフェーズでは、「何を作るべきか」「ユーザーにとっての価値は何か」を明らかにします。成果物は主にドキュメントとして管理されます。本ドキュメントで使用される用語の詳細は `docs/01_agile-glossary-and-ids.md` も参照してください。

### 1.1. アクティビティの定義

- **目的:** ユーザーがシステムで行う主要な行動や体験、またはシステムが提供する主要な機能領域（アクティビティ）を特定し、定義します。これはプロダクトの全体的なスコープを捉えるための最初のステップです。
- **進め方:** プロジェクトのペルソナ定義や全体目標に基づき、主要なアクティビティを洗い出します。
- **成果物:** `docs/02_activities.md` に、各アクティビティのID (`ACT-[機能エリア略称]`)、名称、説明を記述します。
  - 例: `ACT-SHOPPING`, `ACT-OBSERVABILITY`

### 1.2. ユーザーストーリーマッピングの実施

- **目的:** 特定のアクティビティにおけるユーザージャーニーを視覚化し、ユーザーが価値を得るための一連のステップと、各ステップで必要となる機能要求（マッピングレベルのユーザーストーリー）を明らかにします。プロダクトの全体像を把握し、リリースの優先順位付けの基礎となります。
- **進め方 (概要):**
    1. 対象とするアクティビティを選択します (例: `ACT-SHOPPING`)。
    2. そのアクティビティにおけるユーザーの主要な行動ステップ（バックボーン）を時系列に並べます。
    3. 各行動ステップの下に、ユーザーがそのステップで達成したいことや、システムが提供すべき機能をマッピングレベルのユーザーストーリーとして記述していきます (`docs/03-1_user-stories.md` の内容を参照・拡充)。
    4. 必要に応じて、異なるユーザータイプや代替パスも考慮します。
    5. マッピングされたユーザーストーリーを、価値や依存関係、リスクなどを考慮して優先順位付け（例: MVPに必要なもの、その後のリリースで追加するものなど）を行います。
- **成果物:** 視覚的なユーザーストーリーマップ（Miroなどのツールや、Markdownでの簡易的な表現）、およびそれを構成するマッピングレベルのユーザーストーリーリスト (`docs/03-1_user-stories.md`)。

### 1.3. ユーザーストーリー (マッピングレベル) の作成とカタログ化

- **目的:** ユーザーストーリーマッピングで洗い出された個々の機能要求を、標準的なユーザーストーリー形式で具体的に記述し、カタログとして管理します。これは、後のPBI作成の元ネタとなります。
- **進め方:** マッピングで特定された各機能要求について、「〇〇（役割）として、△△（目的）のために、□□（機能）がしたい」という形式で記述します。
- **成果物:** `docs/03-1_user-stories.md` に、各マッピングレベルのユーザーストーリーのID (`US-[ペルソナ略称]-[機能エリア略称]-[識別子]`)、説明、関連アクティビティ、主な価値、関連ペルソナなどを記述します。このドキュメントはPBI候補のカタログとして機能します。
  - 例: `US-CUST-BROWSE-LIST-ALL`, `US-SRE-LOG-E2E-DEBUG`

## 2. プロダクトバックログの初期構築

要求定義で得られたマッピングレベルのユーザーストーリーを基に、開発可能なアイテムの初期リスト（プロダクトバックログのタネ）を作成します。このプロダクトバックログは、最終的にGitHub Projects上でIssueとして管理されていきます。

### 2.1. マッピングレベルのユーザーストーリーから初期PBI候補リストを作成

- `docs/03-1_user-stories.md` に記載されたマッピングレベルのユーザーストーリーをレビューし、スプリントで開発するPBIの候補としてリストアップします。
- 各PBI候補には、開発管理用のPBI ID (`PBI-[ペルソナ略称]-[機能エリア略称]-[US識別子]-[連番]`) を付与する準備をします。
  - 例: `US-CUST-BROWSE-LIST-ALL` から → `PBI-CUST-BROWSE-LIST-ALL-01`
  - 例: `US-SRE-LOG-E2E-DEBUG` を分割する場合 → `PBI-SRE-LOG-E2E-DEBUG-01`, `PBI-SRE-LOG-E2E-DEBUG-02`
- 最初は全てのPBI候補をGitHub Issue化せず、優先度の高いものから順次Issue化していくアプローチも可能です。

### 2.2. 各PBI候補の概要記述の確認

- 各PBI候補が、「〇〇として、△△のために、□□したい」というユーザーストーリー形式で明確に記述されているかを確認・洗練させます。
- スプリント計画時のリファインメントで詳細化されますが、この段階で大まかな受け入れ条件のイメージを持っておくと良いでしょう。

### 2.3. リファクタリングタスクの初期洗い出し

- 現時点で想定される技術的負債や、将来の拡張性・保守性のために必要な内部的な改善タスク（リファクタリング）があれば、それらもPBI候補としてリストアップし、必要に応じてPBI IDを付与します。

## 3. 開発管理ツールのセットアップ (GitHub Projects)

スプリントベースの開発を効率的に管理するため、GitHub Projectsの初期セットアップを行います。

### 3.1. プロジェクト作成

1. GitHubリポジトリの「Projects」タブを開きます。
2. 「New project」をクリックし、「Board」テンプレートなどを選択します。
3. プロジェクト名を入力（例: `[プロジェクト名] 開発ボード`)。
4. 「Create」をクリックします。

### 3.2. カスタムフィールド設定

以下のカスタムフィールドをプロジェクトに追加します。これらはGitHub Issue（PBIやエピックなど）の分類や管理に使用します。

| フィールド名         | タイプ   | 選択肢/フォーマット                                                 |
| -------------------- | -------- | ------------------------------------------------------------------- |
| **Type**             | 単一選択 | `Epic`, `User Story`, `Refactoring`, `Sprint Plan`, `Retrospective` |
| **Sprint**           | Number   | 1, 2, 3...                                                          |
| **Story Points**     | 単一選択 | `S`, `M`, `L`, `XXL` (目安: S=1, M=3, L=5)                          |
| **Status**           | 単一選択 | `Todo`, `In Progress`, `Done` (GitHub Projects標準)                 |
| (任意) **Epic Link** | Text     | (例: `EP-SHOPPING-01` のIssue番号やタイトル)                        |

### 3.3. Issueテンプレートの準備

リポジトリの `.github/ISSUE_TEMPLATE/` ディレクトリに、以下の目的のIssueテンプレートファイル (`.md`) を配置します。

- **Epic (`epic.md`):** 開発管理上のエピックを作成するためのテンプレート。
- **User Story (`user-story.md`):** PBI（ユーザーストーリー形式）を作成するためのテンプレート。
- **Refactoring (`refactoring.md`):** PBI（リファクタリングタスク）を作成するためのテンプレート。
- **Sprint Planning (`sprint-planning.md`):** スプリント計画を記録するためのテンプレート。
- **Retrospective (`retrospective.md`):** スプリントの振り返りを記録するためのテンプレート。

### 3.4. 基本的なビューの設定

GitHub Projects内で、作業状況を把握しやすくするために、以下のようなビューを作成・設定します。

- **プロダクトバックログビュー (Table Layout):** `Type` is `User Story` or `Refactoring`, `Status` is `Todo`。優先度順。
- **現在のスプリントビュー (Board Layout):** `Sprint` is `[現在のスプリント番号]`。列は `Status`。
- **エピック一覧ビュー (Table or Board Layout):** `Type` is `Epic`。

## 4. リリース計画 (中期的な目標設定)

プロダクトの初期リリースや、その後の大規模な機能追加など、中期的な目標と進め方の大枠を計画します。これは、スプリント計画よりも長期的な視点で行われます。

### 4.1. 開発管理上のエピック (Epic) の定義と初期計画

- **目的:** プロジェクトの大きな目標（例: MVPリリース、特定の機能群の完成）を達成するための開発の塊（エピック）を定義します。各エピックは、設計上の特定のアクティビティ (`ACT-...`) の一部または全体を実現することを目的とします。
- **進め方:**
    1. 初期リリースで達成したい主要なアクティビティやマッピングレベルのユーザーストーリー群を特定します。
    2. それらをいくつかのエピック (`EP-[機能エリア略称]-[連番]`) にまとめ、各エピックの具体的な目標、スコープ（含める主要なPBI候補）、完了条件（成功指標）を設定します。
        - 例: `EP-SHOPPING-01`, `EP-OBSERVABILITY-01`
    3. 各エピックのIssueをGitHub Projectsに作成し、`Type` を `Epic` に設定します。Issue本文には、関連するアクティビティIDやエピックの目標、主要PBI候補などを記述します。
- **成果物:** GitHub Issueとして作成されたエピック群。

### 4.2. 大まかなリリースロードマップの検討 (任意)

- **目的:** 各エピックがいつ頃のリリースを目指すのか、大まかなタイムライン（ロードマップ）を視覚化します。
- **進め方:** 各エピックの想定される規模と、チーム（今回はご自身）の想定ベロシティを考慮し、どのエピックが何スプリント程度かかりそうかを見積もります。
- **成果物:** リリースロードマップ図（簡易的なもの）。
