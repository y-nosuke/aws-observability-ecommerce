# AWSオブザーバビリティ学習用eコマースアプリ - 画面一覧

## 顧客向け画面

| 画面ID | 画面名 | 関連ユースケース | 主な機能 | オブザーバビリティポイント |
|--------|--------|-----------------|---------|------------------------|
| **C-SCR-01** | トップページ | - | ・カテゴリナビゲーション<br>・注目商品表示<br>・検索バー<br>・ログイン/登録リンク | ・ページ読み込み時間<br>・カテゴリクリック率<br>・検索利用率<br>・滞在時間 |
| **C-SCR-02** | カテゴリ商品一覧 | UC-C01 | ・カテゴリ別商品表示<br>・フィルタリング<br>・ソート機能<br>・ページネーション | ・カテゴリ別アクセス率<br>・フィルター使用頻度<br>・ページネーション利用率<br>・商品クリック率 |
| **C-SCR-03** | 検索結果 | UC-C02 | ・検索結果表示<br>・フィルタリング<br>・ソート機能<br>・関連検索キーワード提案 | ・検索クエリパターン<br>・検索結果数<br>・ゼロヒット率<br>・フィルター使用頻度 |
| **C-SCR-04** | 商品詳細 | UC-C03 | ・商品情報表示<br>・商品画像ギャラリー<br>・在庫状況表示<br>・カートに追加ボタン<br>・商品レビュー | ・滞在時間<br>・画像閲覧率<br>・カート追加率<br>・レビュー閲覧数 |
| **C-SCR-05** | カート | UC-C04, UC-C05 | ・カート内商品一覧<br>・数量変更<br>・商品削除<br>・小計/合計金額表示<br>・チェックアウトボタン | ・カート放棄率<br>・商品削除率<br>・数量変更頻度<br>・チェックアウト移行率 |
| **C-SCR-06** | チェックアウト | UC-C06 | ・配送先情報入力<br>・支払い方法選択<br>・注文内容確認<br>・注文確定ボタン | ・ステップ完了率<br>・フォーム入力エラー率<br>・支払い処理時間<br>・完了率 |
| **C-SCR-07** | 注文確認 | UC-C06 | ・注文完了メッセージ<br>・注文番号表示<br>・注文詳細サマリー<br>・ショッピング継続ボタン | ・ページ滞在時間<br>・リダイレクト率<br>・追加購入率 |
| **C-SCR-08** | アカウント登録 | UC-C07 | ・ユーザー登録フォーム<br>・パスワード要件表示<br>・利用規約同意<br>・登録ボタン | ・フォーム完了率<br>・入力エラー頻度<br>・登録完了率<br>・離脱ポイント |
| **C-SCR-09** | ログイン | UC-C08 | ・メールアドレス/パスワード入力<br>・ログインボタン<br>・パスワードリセットリンク<br>・新規登録リンク | ・ログイン成功率<br>・失敗理由分析<br>・リセット利用率<br>・新規登録移行率 |
| **C-SCR-10** | マイページ | UC-C09, UC-C10 | ・アカウント概要<br>・注文履歴へのリンク<br>・プロフィール管理リンク<br>・お届け先住所管理リンク | ・ページ閲覧頻度<br>・セクション利用率<br>・滞在時間 |
| **C-SCR-11** | プロフィール管理 | UC-C09 | ・個人情報表示/編集<br>・パスワード変更<br>・メール設定<br>・保存ボタン | ・編集頻度<br>・フィールド更新率<br>・エラー率<br>・完了率 |
| **C-SCR-12** | お届け先管理 | UC-C09 | ・住所一覧<br>・新規住所追加<br>・住所編集<br>・削除機能<br>・デフォルト設定 | ・住所登録数<br>・編集頻度<br>・デフォルト変更率 |
| **C-SCR-13** | 注文履歴 | UC-C10 | ・注文一覧<br>・ステータスフィルター<br>・注文詳細リンク<br>・再注文ボタン | ・閲覧頻度<br>・フィルター使用率<br>・詳細確認率<br>・再注文率 |
| **C-SCR-14** | 注文詳細 | UC-C10, UC-C11 | ・注文情報詳細<br>・配送状況表示<br>・配送追跡リンク<br>・注文商品一覧 | ・ページ閲覧頻度<br>・追跡リンククリック率<br>・問い合わせ率 |
| **C-SCR-15** | 在庫切れ通知登録 | UC-C12 | ・商品情報表示<br>・在庫切れ状態表示<br>・メールアドレス入力<br>・通知登録ボタン | ・通知登録率<br>・入荷後の購入率<br>・登録商品数/ユーザー |

## 管理者向け画面

| 画面ID | 画面名 | 関連ユースケース | 主な機能 | オブザーバビリティポイント |
|--------|--------|-----------------|---------|------------------------|
| **A-SCR-01** | 管理ダッシュボード | - | ・重要KPI概要<br>・新規注文数<br>・在庫アラート<br>・売上グラフ | ・ダッシュボード読み込み時間<br>・セクション閲覧率<br>・アラート確認率 |
| **A-SCR-02** | 在庫管理 | UC-A01, UC-A02 | ・商品在庫一覧<br>・在庫レベルフィルター<br>・在庫数編集<br>・在庫履歴グラフ | ・在庫確認頻度<br>・フィルター使用率<br>・更新操作数<br>・低在庫商品比率 |
| **A-SCR-03** | 低在庫アラート | UC-A03 | ・低在庫商品リスト<br>・閾値設定<br>・アラート確認機能<br>・発注指示ボタン | ・アラート確認時間<br>・対応完了率<br>・発注処理数<br>・アラート頻度 |
| **A-SCR-04** | 注文管理 | UC-A04 | ・新規注文一覧<br>・ステータスフィルター<br>・注文詳細表示<br>・注文承認/保留ボタン | ・注文確認時間<br>・承認率<br>・保留理由分布<br>・処理時間分布 |
| **A-SCR-05** | 出荷管理 | UC-A05 | ・出荷待ち注文一覧<br>・出荷処理フォーム<br>・追跡番号入力<br>・ステータス更新ボタン | ・出荷リードタイム<br>・処理時間<br>・通知送信率<br>・問い合わせ発生率 |
| **A-SCR-06** | 返品・交換管理 | UC-A06 | ・返品/交換リクエスト一覧<br>・詳細確認<br>・承認/拒否ボタン<br>・返金/交換処理機能 | ・処理完了時間<br>・承認率<br>・返品理由分布<br>・顧客満足度指標 |

## オブザーバビリティ管理画面

| 画面ID | 画面名 | 関連ユースケース | 主な機能 | 主な表示メトリクス |
|--------|--------|-----------------|---------|------------------------|
| **O-SCR-01** | システム概要ダッシュボード | UC-O01, UC-O02 | ・システム全体の健全性表示<br>・主要メトリクスグラフ<br>・アラート概要<br>・サービスマップ | ・サービス稼働率<br>・エラー率<br>・レスポンスタイム<br>・リソース使用率 |
| **O-SCR-02** | パフォーマンスダッシュボード | UC-O01, UC-O04 | ・API/エンドポイント別レスポンスタイム<br>・リソース使用率グラフ<br>・ボトルネック可視化<br>・トレース詳細表示 | ・レスポンスタイム（p50,p90,p99）<br>・処理時間内訳<br>・コンポーネント別タイミング<br>・CPU/メモリ使用率 |
| **O-SCR-03** | エラー分析ダッシュボード | UC-O02 | ・エラー率グラフ<br>・エラー種類分布<br>・エラーログ検索<br>・例外スタックトレース表示 | ・HTTPステータスコード分布<br>・例外タイプ分布<br>・影響ユーザー数<br>・エラークラスタリング |
| **O-SCR-04** | リソース計画ダッシュボード | UC-O03 | ・リソース使用率トレンド<br>・予測分析グラフ<br>・季節変動パターン<br>・スケーリングイベント履歴 | ・日次/週次/月次使用率<br>・自動スケーリングイベント<br>・予測vs実績<br>・コスト最適化指標 |
| **O-SCR-05** | トレース分析 | UC-O01, UC-O04 | ・分散トレース可視化<br>・サービスマップ<br>・トレース詳細検索<br>・レイテンシヒートマップ | ・サービス間依存関係<br>・コンポーネント別レイテンシ<br>・エラー伝播パス<br>・クリティカルパス |
| **O-SCR-06** | ログ分析 | UC-O01, UC-O02 | ・構造化ログ検索<br>・ログパターン分析<br>・時系列ログ表示<br>・コンテキスト相関表示 | ・ログボリューム<br>・エラーログ頻度<br>・パターンマッチング<br>・相関イベント |

## ビジネスインテリジェンス画面

| 画面ID | 画面名 | 関連ユースケース | 主な機能 | 主な表示メトリクス |
|--------|--------|-----------------|---------|------------------------|
| **B-SCR-01** | 売上・トラフィック分析 | UC-B01 | ・売上トレンドグラフ<br>・トラフィック分析<br>・コンバージョンファネル<br>・地域別分析 | ・売上金額/件数<br>・訪問者数/セッション<br>・コンバージョン率<br>・客単価/LTV |
| **B-SCR-02** | 商品パフォーマンス分析 | UC-B02 | ・商品別売上ランキング<br>・カテゴリ分析<br>・低パフォーマンス商品一覧<br>・詳細分析ドリルダウン | ・商品別閲覧/購入率<br>・カート追加率<br>・カテゴリ別パフォーマンス<br>・価格帯分析 |
| **B-SCR-03** | ユーザー行動分析 | UC-B01 | ・ユーザーセグメント分析<br>・行動パターン追跡<br>・リピート率グラフ<br>・離脱ポイント分析 | ・新規/リピート比率<br>・セッション長<br>・ページ遷移パターン<br>・セグメント別LTV |
| **B-SCR-04** | キャンペーン効果分析 | UC-B01 | ・キャンペーン比較<br>・時系列効果分析<br>・目標達成率<br>・ROI計算 | ・キャンペーン売上寄与度<br>・コスト対効果<br>・トラフィック増加率<br>・コンバージョン影響度 |
