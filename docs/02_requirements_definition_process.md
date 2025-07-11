# 要件定義プロセスガイド

## 1. はじめに

### 1.1. 要件定義プロセスの全体像

このドキュメントは、「AWSオブザーバビリティ学習用eコマースアプリ」プロジェクトにおいて、プロダクトの要求を定義し、それを基に開発着手前の設計準備（アーキテクチャ概要、技術スタック選定、MVP定義など）を行うまでの一連のプロセスと考え方をまとめたものです。
ペルソナ、アクティビティ、ユーザーストーリーマッピング、ユースケース、MVP定義といった要素をどのように定義・活用し、開発可能な計画に落とし込んでいくかを説明します。

### 1.2. 本ドキュメントの目的と対象読者

- **目的:** 将来、同様のアプリケーション開発プロジェクトを開始する際に、要求定義から設計準備までの進め方やドキュメント構成の参考となること。
- **対象読者:** プロダクトマネージャー、プロダクトオーナー、ビジネスアナリスト、ソフトウェア開発者、SRE、その他ソフトウェア開発の要求定義・設計準備プロセスに関わる方。
- **関連ドキュメント:** 本書を読む前に「[01_concepts_and_terminology.md](01_concepts_and_terminology.md)」で主要な概念と用語を理解することを推奨します。

## 2. 要件定義の基本プロセス

このフェーズでは、「何を作るべきか」「ユーザーにとっての価値は何か」を明らかにします。

### 2.1. プロジェクトの全体像と価値の明確化

- プロジェクトの目的、解決したい課題、ターゲットユーザー、提供する主要な価値を明確にします。
- オブザーバビリティ学習という本プロジェクト特有の目標も考慮に入れます。

### 2.2. アクティビティの定義

- **目的:** ユーザーがシステムで行う主要な行動や体験、またはシステムが提供する主要な機能領域（アクティビティ）を特定し、定義します。これはプロダクトの全体的なスコープを捉えるための最初のステップです。
  - 用語詳細は「[01_concepts_and_terminology.md](01_concepts_and_terminology.md)」参照。
- **進め方:** プロジェクトのペルソナ定義や全体目標に基づき、主要なアクティビティを洗い出します。
- **成果物:** 各アクティビティのID (`ACT-[機能エリア略称]`)、名称、説明を記述したリスト（例: `docs/02_activities.md`）。
  - 例: `ACT-SHOPPING`, `ACT-OBSERVABILITY`

### 2.3. ペルソナの詳細化

- **目的:** ユーザー像を具体的に定義し、要求の主体とニーズを明確化します。
  - 用語詳細は「[01_concepts_and_terminology.md](01_concepts_and_terminology.md)」参照。
- **初期定義からの改善:** 当初「顧客」「管理者」「開発/運用」といったシンプルなペルソナから、役割の多岐性を考慮し、「ECストアマネージャー(`MGR`)」「開発者(`DEV`)」「SRE/運用担当者(`SRE`)」「セキュリティ担当者(`SEC`)」「ビジネスアナリスト(`BA`)」などへ詳細化しました。
- **ポイント:** 詳細なペルソナ定義は、要求の解像度を高め、ユーザーストーリーやユースケースの質を向上させます。
- **成果物:** 各ペルソナの詳細なプロフィール、目標、課題などを記述したドキュメント（例: `docs/01_user-personas.md`）。

## 3. ユーザーストーリーマッピングの実施

### 3.1. 目的と効果

- **目的:** 特定のアクティビティにおけるユーザージャーニーを視覚化し、ユーザーが価値を得るための一連のステップ（バックボーン）と、各ステップで必要となる機能要求（マッピングレベルのユーザーストーリー）を明らかにします。プロダクトの全体像を把握し、リリースの優先順位付けの基礎となります。
- **効果:**
  - チーム全体の共通認識形成。
  - 機能の配置やユーザーフローの妥当性を俯瞰的に確認。
  - 要求の抜け漏れ発見のきっかけ。
  - MVP特定とスコープ決定の視覚的支援。

### 3.2. ユーザーストーリーマップの作成手順 (概要)

1. 対象とするアクティビティを選択します (例: `ACT-SHOPPING`)。
2. そのアクティビティにおけるユーザーの主要な行動ステップ（バックボーン）を時系列に並べます。
3. 各行動ステップの下に、ユーザーがそのステップで達成したいことや、システムが提供すべき機能をマッピングレベルのユーザーストーリーとして記述していきます。
4. 必要に応じて、異なるユーザータイプや代替パスも考慮します。
5. マッピングされたユーザーストーリーを、価値や依存関係、リスクなどを考慮して優先順位付け（例: MVPに必要なもの、その後のリリースで追加するものなど）を行います。これをリリース計画のための優先順位ラインで示します。

### 3.3. 全体像の可視化と抜け漏れ発見

- マップ上でストーリーが少ない箇所などから、要求の抜け漏れを議論するきっかけになります。

### 3.4. 優先順位付けの基礎

- ユーザーストーリーマップは、要求定義とリリース計画をつなぐ重要なツールです。

## 4. ユーザーストーリーとユースケースの定義

### 4.1. マッピングレベルのユーザーストーリー作成

- **目的:** ユーザーストーリーマッピングで洗い出された個々の機能要求を、標準的なユーザーストーリー形式で具体的に記述し、カタログとして管理します。これは、後のPBI作成の元ネタとなります。
  - 用語詳細は「[01_concepts_and_terminology.md](01_concepts_and_terminology.md)」参照。
- **進め方:** マッピングで特定された各機能要求について、「〇〇（役割）として、△△（目的）のために、□□（機能）がしたい」という形式で記述します。
- **分類軸:** ペルソナを第一の分類軸とし、「誰が何をしたいのか」を明確にします。
- **実装ストーリー vs 利用ストーリー:** 技術基盤やオブザーバビリティ計装など、導入自体が価値となる「実装ストーリー」と、機能を利用する「利用ストーリー」を区別します（詳細は「[01_concepts_and_terminology.md](01_concepts_and_terminology.md)」参照）。
- **成果物:** 各マッピングレベルのユーザーストーリーのID (`US-[ペルソナ略称]-[機能エリア略称]-[識別子]`)、説明、関連アクティビティ、主な価値、関連ペルソナなどを記述したリスト（例: `docs/03-1_user-stories.md`）。

### 4.2. ユースケースによる振る舞いの具体化

- **目的:** ユーザーストーリーで定義された要求に対し、システムがどのように振る舞うかを具体的に記述します。アクターとシステムの相互作用、正常系・代替・例外フローを記述します。
  - 用語詳細は「[01_concepts_and_terminology.md](01_concepts_and_terminology.md)」参照。
- **関連付け:** 各ユースケース記述の中に、関連するユーザーストーリーIDを明記し、トレーサビリティを確保します。
- **成果物:** 各ユースケースのID (`UC-[ペルソナ略称]-[機能エリア略称]-[連番]`)、アクター、事前条件、基本フロー、代替フロー、事後条件などを記述したリスト（例: `docs/04_use-cases.md`）。

### 4.3. 詳細化レベルとJust Enoughの考え方

- 全てのユーザーストーリーに詳細なユースケースが必要なわけではありません。機能の複雑性、重要度、チームの共通理解度に応じて詳細化のレベルを判断します（Just Enough）。
- シンプルな機能はユーザーストーリーと受け入れ基準で十分な場合もあります。

## 5. MVP（Minimum Viable Product）の定義

### 5.1. MVPの目的と重要性

- 「Minimum Viable Product = 実用最小限の製品」を定義することで、早期価値提供、早期フィードバック、リスク低減、スコープ管理を実現します。
- MVPは単に機能を削ることではなく、学習やフィードバック獲得という目的に対して、最小限で最大の効果を発揮する範囲を見極める戦略的な活動です。

### 5.2. ユーザーストーリーマップを活用したMVPスコープの決定

- ユーザーストーリーマップ上で、提供したいコアバリューを実現できる最小限のストーリー群を選択し、MVPの境界線を引きます。
- 視覚的にスコープを決定しやすく、どのストーリーを実装すれば最低限のエンドツーエンドの価値を提供できるかを議論する上で不可欠です。

### 5.3. MVPにおける優先順位付けの基準

- ビジネス価値、ユーザー価値、緊急度、技術的実現可能性、依存関係、リスク、学習効果などを総合的に考慮します。

### 5.4. 技術的な実現性とリスク評価

- MVP候補のユーザーストーリー群について、技術的な実現可能性を評価し、潜在的なリスクを洗い出します。
- 必要であれば、技術調査（スパイク）を実施し、リスクの高い要素からプロトタイプを作成することも検討します。

## 6. 要件定義書の作成

MVPスコープとそれ以降のリリース候補に基づき、各種要件を文書化します。

### 6.1. 機能要件のとりまとめ

- ユーザーストーリーやユースケースを基に、システムが提供すべき具体的な機能の一覧を作成します。
- **成果物:** 機能一覧ドキュメント（例: `Function List (MVP)`）

### 6.2. 非機能要件の定義と文書化

- 性能、可用性、セキュリティ、保守性、そして本プロジェクトでは特に**オブザーバビリティ**に関する品質特性の要求を定義します。
- **成果物:** 非機能要件一覧ドキュメント（例: `Non-Functional Requirements`）

### 6.3. 画面一覧・ワイヤーフレーム

- ユーザーインターフェースの構成要素をリスト化し、必要に応じて主要画面の構成案（ワイヤーフレーム）を作成します。
- **成果物:** 画面一覧ドキュメント、ワイヤーフレーム（例: `Screen List (Full)`, `Bolt.new依頼文`）

### 6.4. データ要件（概要）

- システムで扱われる主要なデータエンティティ、その属性、エンティティ間の関連性を概要レベルで定義します。
- **成果物:** データモデル概要図（例: `Function List & Data Model (MVP)` 内のデータモデル部分）

## 7. 設計ドキュメント（概要レベル）の作成

要件定義で得られた情報を基に、開発のインプットとなる設計概要を作成します。

### 7.1. アーキテクチャ設計（概要レベル）

- MVPスコープの機能と非機能要件を満たすためのシステム全体の構成、主要コンポーネント（AWSサービス等）、それらの連携方式を図示します。
- **成果物:** アーキテクチャ概要図、説明ドキュメント（例: `Architecture Design & Tech Stack (MVP)`）

### 7.2. 技術スタックの選定と文書化

- アーキテクチャに基づき、採用するプログラミング言語、フレームワーク、ライブラリ、主要なクラウドサービス等を具体的に決定します。
- **成果物:** 技術スタック一覧（例: `Architecture Design & Tech Stack (MVP)` 内）

### 7.3. データモデル（詳細）

- データ要件（概要）を基に、より詳細なデータエンティティ、属性、型、リレーションシップを定義します。ER図などを用います。
- 詳細度合いは「Just Enough」を意識し、初期は概要、実装直前に詳細化でも可。
- **成果物:** 詳細データモデル図、定義書（例: `Function List & Data Model (MVP)` を拡張）

### 7.4. API設計（概要）

- バックエンドAPIの主要なエンドポイント、リクエスト/レスポンス形式、認証方式などを概要レベルで定義します。
- **成果物:** API設計書（概要）（例: `Backend API Design (Full/MVP)`）

### 7.5. 環境定義

- 開発から本番（学習用）までの各環境構成（ローカル、開発、ステージング、本番等）と管理方針を定義します。CI/CD、IaCの利用方針も検討します。
- **成果物:** 環境定義書（例: `Environment Definition`）

## 8. 継続的な見直しと対話の重要性

- 要求定義から設計準備に至るプロセスは一方通行ではなく、対話とフィードバックによる反復的な改善の連続です。
- 最初から完璧な定義・設計を目指すのではなく、各要素を作成・議論する中で出てきた疑問や矛盾点、新たな発見をもとに、柔軟に定義や計画を見直していくことが重要です。
- ドキュメントはあくまで共通理解を形成・記録するためのツールであり、関係者間の継続的な対話こそがプロジェクト成功の鍵です。
