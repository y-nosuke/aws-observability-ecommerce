# 1. GitHub Projects初期セットアップ

## 1.1. プロジェクト作成手順

1. GitHubリポジトリにアクセス
2. 「Projects」タブを選択
3. 「New project」をクリック
4. テンプレート「Board」を選択
5. プロジェクト名：`AWS オブザーバビリティeコマースアプリ 開発計画`を入力
6. 「Create」をクリック

## 1.2. カスタムフィールドの設定

以下のカスタムフィールドを追加:

| フィールド名 | タイプ   | 選択肢/フォーマット                             |
| ------------ | -------- | ----------------------------------------------- |
| Type         | 単一選択 | Epic, User Story, Refactoring                   |
| Sprint       | Number   | 1, 2, 3...                                      |
| Story Points | 単一選択 | S (1日以内), M (2-3日), L (3-5日), XXL (要分割) |

- **Epic (親issue)**：大きな機能群を表す
- **User Story (子issue)**：Epicに紐づく小さな機能単位
- **リンク方法**：User Story作成時に「Linked issues」セクションで親Epicを指定

## 1.3. ステータスの設定

GitHub Projectsのデフォルトのステータスを以下のように設定:

- Todo (未着手)
- In Progress (進行中)
- Done (完了)

## 1.4. GitHub Issueテンプレートの設定

`.github/ISSUE_TEMPLATE/` ディレクトリに以下のテンプレートファイルを作成します。

### 1.4.1. epic.md

```markdown
---
name: Epic
about: 大きな機能群（エピック）を作成
title: 'EPIC-XX: '
labels: 'epic'
assignees: ''
---

# Epic: [タイトル]

## 目的と価値

[このエピックがユーザーや事業にもたらす価値]

## ストーリーマップ

| ユーザー行動1 | ユーザー行動2 | ユーザー行動3 |
| ------------- | ------------- | ------------- |
| [ストーリー1] | [ストーリー3] | [ストーリー5] |
| [ストーリー2] | [ストーリー4] | [ストーリー6] |
| [ストーリー7] | [ストーリー8] | [ストーリー9] |

## 成功指標

- [このエピックの完了を測る指標]

```

### 1.4.2. user-story.md

```markdown
---
name: User Story
about: 新しいユーザーストーリーを作成
title: 'US-XX: '
labels: 'user-story'
assignees: ''
---

# User Story: [タイトル]

## エピック

[該当するエピック名]

## 説明

[ユーザーストーリーの説明]

## 受け入れ基準

- [ ] [条件1]
- [ ] [条件2]
- [ ] [条件3]

## タスク

### 1. 要件定義・設計

- [ ] 要件確認
- [ ] 基本設計
- [ ] 詳細設計

### 2. バックエンド実装

- [ ] API設計
- [ ] API実装
- [ ] 単体テスト

### 3. フロントエンド実装

- [ ] UI設計
- [ ] UI実装
- [ ] 単体テスト

### 4. テスト・レビュー

- [ ] 結合テスト
- [ ] レビュー
- [ ] 修正対応

### 5. リリース準備

- [ ] ドキュメント作成
- [ ] デプロイ準備
```

### 1.4.3. refactoring.md

```markdown
---
name: Refactoring
about: リファクタリング課題を作成
title: 'RF-XX: '
labels: 'refactoring'
assignees: ''
---

# Refactoring: [タイトル]

## 目的

[リファクタリングの目的と背景]

## 改善目標

- [ ] [目標1]
- [ ] [目標2]

## タスク

### 1. 現状分析

- [ ] 問題点特定
- [ ] 改善計画
- [ ] テスト計画

### 2. 実装改善

- [ ] バックエンド改善
- [ ] フロントエンド改善

### 3. テスト検証

- [ ] 機能テスト
- [ ] 性能検証

### 4. 完了処理

- [ ] ドキュメント更新
- [ ] マージ作業
```

## 1.5. ビューの設定

### 1.5.1. スプリント別ビュー

- グループ化: Sprint
- レイアウト: Board
- 表示項目: Title, Type, Story Points, Status

### 1.5.2. ストーリーポイント別ビュー

- グループ化: Story Points
- レイアウト: Board
- 表示項目: Title, Type, Sprint, Status

### 1.5.3. タイプ別ビュー

- グループ化: Type
- レイアウト: Board
- 表示項目: Title, Story Points, Sprint, Status

### 1.5.4. エピック進捗ビュー

- **並び順**：手動で調整（優先順に）
- **レイアウト**：Board
- **Filter by**：Type is Epic
- **表示項目**：Title, Type, Status, Story Points

### 1.5.5. ストーリービュー

- **グループ化**：Parent Issue
- **レイアウト**：Board
- **表示項目**：Title, Type, Status, Story Points

この方法では優先度を定義せず、エピックの並び順のみで優先順位を表現でき、GitHub標準機能のSub-issuesを活用した階層構造が実現できます。
