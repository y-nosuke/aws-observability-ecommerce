# 1. 要求定義から設計準備までのプロセスと考え方：eコマースアプリ開発事例

## 1.1. 目次

- [1. 要求定義から設計準備までのプロセスと考え方：eコマースアプリ開発事例](#1-要求定義から設計準備までのプロセスと考え方eコマースアプリ開発事例)
  - [1.1. 目次](#11-目次)
  - [1.2. はじめに](#12-はじめに)
    - [1.2.1. この資料の目的](#121-この資料の目的)
    - [1.2.2. 対象読者](#122-対象読者)
  - [1.3. 要求定義・設計準備要素の概要](#13-要求定義設計準備要素の概要)
  - [1.4. ID命名規則](#14-id命名規則)
  - [1.5. 定義・整理・計画のプロセスとポイント](#15-定義整理計画のプロセスとポイント)
    - [1.5.1. ペルソナ (Persona)](#151-ペルソナ-persona)
    - [1.5.2. エピック (Epic)](#152-エピック-epic)
    - [1.5.3. ユーザーストーリー (User Story)](#153-ユーザーストーリー-user-story)
    - [1.5.4. ユースケース (Use Case)](#154-ユースケース-use-case)
    - [1.5.5. ユーザーストーリーマップの活用](#155-ユーザーストーリーマップの活用)
    - [1.5.6. 要求定義から設計準備への移行](#156-要求定義から設計準備への移行)
    - [1.5.7. MVP定義の重要性](#157-mvp定義の重要性)
    - [1.5.8. 設計ドキュメント（概要レベル）の作成](#158-設計ドキュメント概要レベルの作成)
    - [1.5.9. 環境定義の早期検討](#159-環境定義の早期検討)
    - [1.5.10. 継続的な見直しと対話の重要性](#1510-継続的な見直しと対話の重要性)
  - [1.6. 成果物一覧](#16-成果物一覧)
  - [1.7. まとめ](#17-まとめ)

## 1.2. はじめに

### 1.2.1. この資料の目的

この資料は、「AWSオブザーバビリティ学習用eコマースアプリ」プロジェクトにおいて、要求を定義し、それを基に**開発着手前の設計準備（アーキテクチャ概要、技術スタック選定、MVP定義など）** を行うまでの一連のプロセスと考え方をまとめたものです。エピック、ペルソナ、ユーザーストーリー、ユースケース、非機能要件、ユーザーストーリーマップといった要素をどのように定義・活用し、**議論を通じて洗練させ、開発可能な計画に落とし込んでいったか**を記録しています。

将来、同様のアプリケーション開発プロジェクトを開始する際に、要求定義から設計準備までの進め方やドキュメント構成の参考となることを目的としています。

### 1.2.2. 対象読者

- プロダクトマネージャー、プロダクトオーナー
- ビジネスアナリスト
- ソフトウェア開発者、エンジニアリングマネージャー
- SRE、運用担当者
- セキュリティ担当者
- QAエンジニア
- その他、ソフトウェア開発の要求定義・設計準備プロセスに関わる方

## 1.3. 要求定義・設計準備要素の概要

本プロジェクトでは、以下の要素を用いて要求を段階的に詳細化し、開発計画を策定しました。

- **ペルソナ (Persona):** ユーザー像を具体的に定義し、要求の主体とニーズを明確化します。（→ `01_user-personas.md`）
- **エピック (Epic):** 大きなビジネス価値の塊や戦略目標を示します。（→ `02-1_epic.md`）
- **ユーザーストーリー (User Story):** 特定ペルソナにとって価値ある機能要求を、開発可能な単位で記述します。（→ `03-1_user-stories.md`）
- **ユースケース (Use Case):** システムがアクターとどのように相互作用するか、具体的な振る舞いを記述します。（→ `04_use-cases.md`）
- **非機能要件 (Non-Functional Requirement):** 品質特性（性能、セキュリティ、オブザーバビリティ等）に関する要求を定義します。特に本プロジェクトではオブザーバビリティ学習目標を反映しました。（→ `Non-Functional Requirements`）
- **ユーザーストーリーマップ (User Story Map):** ユーザーストーリーをユーザー活動の時系列に沿ってマッピングし、全体像把握、優先順位付け、MVP特定に活用します。（→ `03-2-2_story-map.md` 修正版）
- **アーキテクチャ設計（概要）:** システム全体の構成、主要コンポーネント、連携方式を定義します。（→ `Architecture Design & Tech Stack (MVP)`）
- **技術スタック (Technology Stack):** 採用する言語、フレームワーク、AWSサービス等を明確にします。（→ `Architecture Design & Tech Stack (MVP)`）
- **機能一覧 (Function List):** ユーザーストーリーから具体的なシステム機能をリスト化します。（→ `Function List & Data Model (MVP)`）
- **データモデル（概要）:** 主要なデータエンティティ、属性、関連性を定義します。（→ `Function List & Data Model (MVP)`）
- **画面一覧 (Screen List) / ワイヤーフレーム:** ユーザーインターフェースの構成要素をリスト化し、必要に応じて画面構成案（ワイヤーフレーム）を作成します。（→ `Screen List (Full)`, Bolt.new依頼文）
- **MVP定義 (Minimum Viable Product Definition):** 最初のリリースで提供する最小限の価値とスコープを明確にします。（→ `MVP Definition`）
- **環境定義 (Environment Definition):** 開発から本番までの各環境構成と管理方針を定義します。（→ `Environment Definition`）
- **API設計（概要）:** バックエンドAPIのエンドポイント、リクエスト/レスポンス形式を定義します。（→ `Backend API Design (Full/MVP)`）

エピック、ユーザーストーリー、ユースケースの関係性については、`epic-user-story-usecase-relation.md` で詳細に解説しています。

## 1.4. ID命名規則

一貫性と追跡可能性を高めるため、以下の命名規則を導入しました。

- **基本形式:** `[種別]-[ペルソナ略称]-[機能エリア略称]-[連番]`
- **種別:** `EP`, `US`, `UC`, `NFR`, `SCR`, `FN` など
- **ペルソナ略称:** `CUST`, `MGR`, `DEV`, `SRE`, `SEC`, `BA`, `SYS`, `ADMIN` など
- **機能エリア略称:** `SHOP`, `AUTH`, `OBS`, `PROD`, `IMPL` など多数定義
- **横断的エピック:** `EP-CROSSCUTTING-{機能エリア}` (例: `EP-CROSSCUTTING-AUTH`)

詳細は `id-convention.md` を参照してください。

## 1.5. 定義・整理・計画のプロセスとポイント

以下に、各要素を定義・整理し、開発計画に繋げていく過程で議論となった点や、採用した考え方をまとめます。

### 1.5.1. ペルソナ (Persona)

- **初期定義:** 当初は「顧客」「管理者」「開発/運用担当者」の3つのシンプルなペルソナでした。
- **課題:** ユーザーストーリーやユースケースを具体化する中で、「管理者」や「開発/運用」の役割が多岐にわたり、1つのペルソナでは要求を正確に表現できないことが判明しました。特にオブザーバビリティやセキュリティに関する要求は、専門的な視点が必要でした。
- **改善:**
  - 「管理者」を「ECストアマネージャー (`MGR`)」とし、店舗運営業務にフォーカス。
  - 「開発/運用担当者」を「開発者 (`DEV`)」と「SRE/運用担当者 (`SRE`)」に分離。
  - 新たに「セキュリティ担当者 (`SEC`)」と「ビジネスアナリスト (`BA`)」を追加。
- **ポイント:** **詳細なペルソナ定義は、要求の解像度を高める上で非常に重要です。** 誰のための機能なのか、その人は何を課題とし、何を求めているのかを具体化することで、ユーザーストーリーやユースケースの質が向上します。

### 1.5.2. エピック (Epic)

- **役割:** 大きなビジネス価値や機能領域を示すものとして定義しました。
- **横断的関心事:** 認証 (`AUTH`)、通知 (`NOTIF`)、オブザーバビリティ (`OBS`) など、複数のビジネス機能から利用される、あるいはシステム全体の品質に関わる要素については、当初 `EP-AUTHENTICATION` のように定義していましたが、その横断的な性質を明確にするため、IDを `EP-CROSSCUTTING-AUTH` のように変更しました。
- **ポイント:** エピックレベルでビジネス機能と横断的基盤を区別することで、アーキテクチャ上の考慮点を早期に意識できます。

### 1.5.3. ユーザーストーリー (User Story)

- **分類軸:** **ペルソナ**を第一の分類軸としました。各ペルソナのセクション内に、関連する機能エリアごとのストーリーを配置する構成を採用しました。これにより、「誰が何をしたいのか」が明確になります。
- **ID体系:** `US-[ペルソナ略称]-[機能エリア略称]-[連番]` を採用し、IDから主体と内容を推測しやすくしました。
- **実装ストーリー vs 利用ストーリー:**
  - 特に技術基盤（認証、通知、CI/CD等）やオブザーバビリティ計装など、**それ自体の導入・設定が将来の効率性、品質向上、非機能要件達成に繋がる**ものについては、「実装ストーリー (`US-[ペルソナ]-...-IMPL-01`)」として独立させました。価値の主体は主に開発者(`DEV`)やSRE(`SRE`)自身です。`(実装)` と明記しました。
  - **実装ストーリーの価値記述:** 目的(`Why`)には、将来の効率性、保守性、非機能要件への貢献、技術的負債の防止などを記述します。例: `US-DEV-CICD-IMPL-01` (CI/CDパイプライン構築)。
  - **乖離リスクと対策:** 実装ストーリーを独立させる場合、**利用側のニーズから乖離し、過剰または過少な実装となるリスク**があります。この対策として、以下の点を意識しました。
    - **目的の明確化:** 利用ストーリーとの関連（例：「`US-DEV-OBS-DEBUG-01` のために」）を明記。
    - **依存関係:** 関連する利用ストーリーとの依存関係を意識。
    - **受入基準:** 利用シーンを反映した基準を設定。
    - **段階的実装:** MVPアプローチで最小限から始め、早期フィードバックを得る。
    - **チーム連携:** 開発者と利用者（SRE等）間の密なコミュニケーション。
  - 機能を利用するストーリーは「利用ストーリー」として定義し、その機能がもたらす直接的な価値を記述します。`(利用)` と明記（または省略）。
- **横断的機能（通知など）の配置:**
  - 当初、通知関連のユーザーストーリー (`US-CUST-NOTIF-*`) を `EP-CROSSCUTTING-NOTIF` に配置する案もありましたが、**顧客が通知を受け取る文脈（注文完了時、在庫入荷時など）**が重要であるため、関連するビジネスエピック (`EP-SHOPPING`, `EP-CUSTOMER-ENGAGEMENT`) に属するストーリーとして、該当セクションに配置し直しました。IDも `US-CUST-CHECKOUT-NOTIF-01` のように文脈が分かるように変更しました。
  - 一方で、通知基盤の**実装**そのものは `EP-CROSSCUTTING-NOTIF` に属する `US-DEV-NOTIF-IMPL-01` として定義しました。
- **ポイント:** ユーザーストーリーはペルソナの価値を中心に記述しつつ、実装と利用を意識的に区別すること、横断的機能は利用文脈と実装基盤を分けて考えること、実装ストーリーのリスクを認識し対策を講じることが、要求の明確化と適切な開発スコープ設定に繋がります。

### 1.5.4. ユースケース (Use Case)

- **役割:** ユーザーストーリーで定義された要求に対し、**システムがどのように振る舞うか**を具体的に記述する目的で利用します。アクターとシステムの相互作用、正常系・代替・例外フローを記述します。
- **ID体系:** ユーザーストーリーに合わせて `UC-[ペルソナ略称]-[機能エリア略称]-[連番]` を採用しました。
- **関連付け:** 各ユースケース記述の中に、関連するユーザーストーリーIDを明記することで、トレーサビリティを確保します。
- **詳細化レベルと Just Enough:**
  - 本プロジェクトでは、特にオブザーバビリティ学習の目的から、オブザーバビリティポイントを洗い出すために多くのユースケースを詳細化しました。
  - しかし、**全てのユーザーストーリーに詳細なユースケースが必要なわけではありません。** アジャイル開発の原則である **「Just Enough (ちょうど十分なだけ)」** に基づき、機能の複雑性、重要度、チームの共通理解度に応じて詳細化のレベルを判断することが重要です。シンプルな機能はユーザーストーリーと受け入れ基準で十分な場合もあります。ドキュメント作成が目的化しないよう注意が必要です。
- **ポイント:** ユースケースはシステムの振る舞いを具体化し、認識齟齬を防ぐ有効なツールですが、その作成コストと効果を考慮し、必要な箇所に重点的に適用することが推奨されます。ペルソナ別のユースケース定義は、多様な利用者の要求を正確に捉える上で役立ちます。

### 1.5.5. ユーザーストーリーマップの活用

- **目的:** ユーザーストーリー一覧だけでは見えにくい**プロダクト全体の流れ（ユーザー体験）**を可視化し、チーム全体の共通認識を形成するために作成しました。
- **構成:** 横軸に主要なユーザー活動（バックボーン）、縦軸に活動内のタスク（ステップ）、その下に個々のユーザーストーリーを配置しました。
- **効果:**
  - **全体像把握:** 機能の配置やユーザーフローの妥当性を俯瞰的に確認できました。
  - **抜け漏れ発見:** マップ上でストーリーが少ない箇所などから、要求の抜け漏れを議論するきっかけになりました。
  - **優先順位付けとMVP特定:** マップ上で「どこまでを最初のリリース（MVP）とするか」の境界線を引くことで、**視覚的にスコープを決定**しやすくなりました。どのストーリーを実装すれば、最低限のエンドツーエンドの価値を提供できるかを議論する上で不可欠でした。
- **ポイント:** ユーザーストーリーマップは、要求定義とリリース計画をつなぐ重要なツールです。特にMVP（Minimum Viable Product）を定義する際には強力な効果を発揮します。

### 1.5.6. 要求定義から設計準備への移行

- **プロセスフロー:** ユーザーストーリーマップでMVPの候補範囲が見えた後、以下のステップを並行または反復的に進めました。
    1. **非機能要件の明確化:** 特にオブザーバビリティ学習目標を具体化し、性能、可用性などの品質要件を定義しました。これが技術選定やアーキテクチャの制約条件となりました。
    2. **アーキテクチャ設計（概要）:** MVPスコープの機能と非機能要件を満たすための主要コンポーネント（AWSサービス）とその連携方式を図示しました。LambdaとFargateを組み合わせるなど、学習目標を達成できる構成を意識しました。
    3. **技術スタックの選定:** アーキテクチャに基づき、使用する言語、フレームワーク、ライブラリ、主要AWSサービスを具体的に決定しました。Bobへの変更など、最新状況も反映しました。
    4. **MVP定義の確定:** 上記の検討結果（技術的実現性、工数見積もり等）をユーザーストーリーマップにフィードバックし、最終的なMVPスコープを確定しました。
    5. **設計ドキュメント（概要レベル）作成:** MVPスコープに基づき、機能一覧、データモデル、画面一覧、API設計を作成し、開発のインプットとしました。
    6. **環境定義:** 開発、テスト、デプロイ、学習活動を円滑に進めるための環境構成（ローカル～本番）と管理方針を定義しました。
- **相互依存性:** これらのステップは相互に影響し合います。例えば、非機能要件（高可用性）がアーキテクチャ（Multi-AZ構成）に影響を与え、選択した技術スタック（Go+Bob）がデータモデルやAPI設計に影響を与えます。そのため、**行ったり来たりしながら整合性を取っていく**必要がありました。
- **ポイント:** 要求定義からスムーズに開発に移るためには、アーキテクチャ、技術スタック、MVPスコープ、そしてそれらを実現するための具体的な機能、データ、画面、APIといった要素を、**概要レベルで定義し、関係者間で合意形成**しておくことが重要です。

### 1.5.7. MVP定義の重要性

- **目的:** 「Minimum Viable Product = 実用最小限の製品」を定義することで、以下のメリットを享受しました。
  - **早期価値提供:** 学習に必要なコア機能を早期に利用可能にする。
  - **早期フィードバック:** 実際に動くものを通じて、より具体的な改善点や次のステップを特定する。
  - **リスク低減:** 小さく始めることで、技術的・計画的な不確実性を低減する。
  - **スコープ管理:** プロジェクトの初期段階で「やること」と「やらないこと」を明確にする。
- **プロセス:** ユーザーストーリーマップ上で、提供したいコアバリュー（基本的なECフロー + 基本的なオブザーバビリティ体験）を実現できる最小限のストーリー群を選択しました。技術的な実現性や想定工数も考慮して調整しました。
- **ポイント:** MVPは単に機能を削ることではなく、**学習やフィードバック獲得という目的に対して、最小限で最大の効果を発揮する範囲**を見極める戦略的な活動です。

### 1.5.8. 設計ドキュメント（概要レベル）の作成

- **対象:** 機能一覧、データモデル、画面一覧、API設計など。
- **目的:**
  - **共通認識の形成:** チーム内で「何を作るのか」についての具体的なイメージを共有する。
  - **実装のインプット:** 開発者が詳細設計やコーディングを進める上での基本的な指針とする。
  - **スコープの明確化:** MVPで実装する機能、データ、画面、APIの範囲を具体的に示す。
- **詳細度 (Just Enough):** 詳細設計レベルまで作り込まず、**概要レベル**に留めることを意識しました。ER図は主要エンティティとリレーションのみ、API設計は主要なエンドポイントとリクエスト/レスポンスの骨子のみ、といった具合です。これは、アジャイル開発において**詳細設計は実装直前や実装中に行われる**ことが多いこと、そしてドキュメントのメンテナンスコストを考慮したためです。
- **ポイント:** 設計ドキュメントは開発を円滑に進めるためのツールであり、過剰な詳細化は避けるべきです。**「ちょうど十分な (Just Enough)」**レベルで情報を共有し、詳細は実装前のタスクブレイクダウンや実装中のコミュニケーションで補完するのが効率的です。

### 1.5.9. 環境定義の早期検討

- **目的:** ローカルでの開発効率、CI/CDによる自動化、各段階でのテスト、本番（学習用）環境での安定した学習体験を実現するために、早期に環境構成と管理方針を定義しました。
- **考慮事項:** 環境分離、IaCによるインフラ管理、CI/CDパイプライン、設定管理、データ管理、オブザーバビリティ設定の環境差などを検討しました。`awslocal`や`tflocal`の利用方針も定めました。
- **ポイント:** 環境定義は後回しにされがちですが、**開発プロセス全体の効率性と品質**に大きく影響します。特にCI/CDやIaCを導入する場合、初期段階での方針決定が重要です。

### 1.5.10. 継続的な見直しと対話の重要性

- 本プロジェクトの要求定義から設計準備に至るプロセスは、**一方通行ではなく、対話とフィードバックによる反復的な改善**の連続でした。
  - ペルソナ定義 → ユーザーストーリー洗い出し → マップ作成 → MVP議論 → 非機能要件定義 → アーキテクチャ検討 → 技術選定見直し → MVP再定義 → 設計概要作成 → 環境定義 → ...
- **ポイント:** 最初から完璧な定義・設計を目指すのではなく、各要素を作成・議論する中で出てきた疑問や矛盾点、新たな発見をもとに、**柔軟に定義や計画を見直していく**ことが、最終的により整合性の取れた、実現可能な計画に繋がります。ドキュメントはあくまで**共通理解を形成・記録するためのツール**であり、**関係者間の継続的な対話**こそが、要求定義から設計、開発に至るまで、プロジェクトの成功に最も重要な要素となります。

## 1.6. 成果物一覧

このプロセスを経て、以下のドキュメントが作成・更新されました。

- `00-0_requirements-process-summary.md`: このドキュメント
- `00-1_id-convention.md`: ID命名規則
- `01_user-personas.md`: ユーザーペルソナ定義 (修正版)
- `02-1_epic.md`: エピック一覧 (修正版)
- `03-1_user-stories.md`: ユーザーストーリー一覧 (修正版)
- `03-2-2_story-map.md`: ユーザーストーリーマップ (新規作成・修正版)
- `04_use-cases.md`: ユースケース一覧 (修正版)
- `Non-Functional Requirements`: 非機能要件定義 (新規作成・ID修正版)
- `Architecture Design & Tech Stack (MVP)`: アーキテクチャ設計と技術スタック (新規作成・修正版)
- `Function List & Data Model (MVP)`: 機能一覧とデータモデル (新規作成・修正版)
- `Screen List (Full)`: 画面一覧 (新規作成・修正版)
- `Backend API Design (Full/MVP)`: バックエンドAPI設計 (新規作成・修正版)
- `MVP Definition`: MVP定義 (新規作成)
- `Environment Definition`: 環境定義 (新規作成)
- `epic-user-story-usecase-relation.md`: エピック・ユーザーストーリー・ユースケースの関係 (修正版)

## 1.7. まとめ

本プロジェクトでは、ペルソナを詳細化し、ペルソナ中心の視点でユーザーストーリーとユースケースを整理・分類することから始め、ユーザーストーリーマップを活用して全体像を把握し、MVPスコープを決定しました。その後、非機能要件（特にオブザーバビリティ学習目標）を明確にし、それを満たすアーキテクチャ概要と技術スタックを選定しました。さらに、MVP開発に必要な機能一覧、データモデル、画面一覧、API設計を「概要レベル」で定義し、開発環境の方針も定めました。

この一連のプロセスを通じて、**要求定義から具体的な開発計画へと段階的に落とし込む**ことの重要性を再確認しました。特に、ユーザーストーリーマップによる**全体像の可視化と優先順位付け**、非機能要件とアーキテクチャの**早期連携**、MVP定義による**スコープ管理**、そして設計準備ドキュメントにおける **「Just Enough」の原則**が、効率的かつ目的指向の開発を進める上で有効でした。

最終的に、これらのドキュメントは完成形ではなく、**継続的な対話とフィードバックを通じて進化させていくべきもの**です。この記録が、今後のプロジェクトにおける要求定義・設計準備プロセスの改善に繋がることを期待します。
