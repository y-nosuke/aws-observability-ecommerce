-- カテゴリーのテストデータ
INSERT INTO categories (name, slug, description, image_url) VALUES
('電子機器', 'electronics', '最新の電子機器製品', 'https://picsum.photos/400/300?random=1'),
('洋服', 'clothing', 'トレンドの洋服コレクション', 'https://picsum.photos/400/300?random=2'),
('書籍', 'books', '書籍、電子書籍など', 'https://picsum.photos/400/300?random=3'),
('ホーム＆キッチン', 'home-kitchen', '家庭用品とキッチン用品', 'https://picsum.photos/400/300?random=4'),
('スポーツ', 'sports', 'スポーツ用品、フィットネス機器', 'https://picsum.photos/400/300?random=5'),
('ビューティー', 'beauty', 'コスメ、スキンケア、ヘアケア', 'https://picsum.photos/400/300?random=6'),
('おもちゃ', 'toys', '子供向けおもちゃ、ゲーム', 'https://picsum.photos/400/300?random=7'),
('ペット用品', 'pet-supplies', 'ペットフード、おもちゃ、ケア用品', 'https://picsum.photos/400/300?random=8');

-- 電子機器のサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('スマートフォン', 'smartphones', '最新のスマートフォン', 'https://picsum.photos/400/300?random=9', 1),
('ノートパソコン', 'laptops', '高性能ノートPC', 'https://picsum.photos/400/300?random=10', 1),
('タブレット', 'tablets', 'タブレット端末', 'https://picsum.photos/400/300?random=11', 1),
('デスクトップPC', 'desktop-pcs', '高性能デスクトップPC', 'https://picsum.photos/400/300?random=12', 1),
('カメラ', 'cameras', 'デジタルカメラ、ビデオカメラ', 'https://picsum.photos/400/300?random=13', 1),
('オーディオ', 'audio', 'スピーカー、イヤホン、ヘッドホン', 'https://picsum.photos/400/300?random=14', 1);

-- 洋服のサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('メンズ', 'mens-clothing', 'メンズファッション', 'https://picsum.photos/400/300?random=15', 2),
('レディース', 'womens-clothing', 'レディースファッション', 'https://picsum.photos/400/300?random=16', 2),
('キッズ', 'kids-clothing', '子供服', 'https://picsum.photos/400/300?random=17', 2),
('スポーツウェア', 'sports-wear', 'スポーツ用衣類', 'https://picsum.photos/400/300?random=18', 2);

-- 書籍のサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('ビジネス', 'business-books', 'ビジネス、経営、経済', 'https://picsum.photos/400/300?random=19', 3),
('コンピュータ', 'computer-books', 'プログラミング、IT', 'https://picsum.photos/400/300?random=20', 3),
('小説', 'novels', '小説、文学', 'https://picsum.photos/400/300?random=21', 3),
('趣味', 'hobby-books', '趣味、実用書', 'https://picsum.photos/400/300?random=22', 3);

-- ホーム＆キッチンのサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('キッチン家電', 'kitchen-appliances', '調理家電、キッチン用品', 'https://picsum.photos/400/300?random=23', 4),
('家具', 'furniture', 'テーブル、椅子、収納家具', 'https://picsum.photos/400/300?random=24', 4),
('インテリア', 'interior', '照明、カーテン、装飾品', 'https://picsum.photos/400/300?random=25', 4),
('生活用品', 'daily-necessities', '日用品、消耗品', 'https://picsum.photos/400/300?random=26', 4);

-- スポーツのサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('フィットネス', 'fitness', 'トレーニング機器、ウェア', 'https://picsum.photos/400/300?random=27', 5),
('アウトドア', 'outdoor', 'キャンプ、登山用品', 'https://picsum.photos/400/300?random=28', 5),
('球技', 'ball-sports', 'サッカー、バスケットボール用品', 'https://picsum.photos/400/300?random=29', 5),
('水泳', 'swimming', '水着、アクセサリー', 'https://picsum.photos/400/300?random=30', 5);

-- ビューティーのサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('スキンケア', 'skin-care', '化粧水、クリーム、美容液', 'https://picsum.photos/400/300?random=31', 6),
('メイクアップ', 'makeup', 'ファンデーション、リップ、アイメイク', 'https://picsum.photos/400/300?random=32', 6),
('ヘアケア', 'hair-care', 'シャンプー、トリートメント', 'https://picsum.photos/400/300?random=33', 6),
('フレグランス', 'fragrance', '香水、ボディミスト', 'https://picsum.photos/400/300?random=34', 6);

-- おもちゃのサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('知育玩具', 'educational-toys', '学習、知育玩具', 'https://picsum.photos/400/300?random=35', 7),
('アクションフィギュア', 'action-figures', 'フィギュア、人形', 'https://picsum.photos/400/300?random=36', 7),
('ボードゲーム', 'board-games', 'カードゲーム、ボードゲーム', 'https://picsum.photos/400/300?random=37', 7),
('ぬいぐるみ', 'plush-toys', 'ぬいぐるみ、ソフトトイ', 'https://picsum.photos/400/300?random=38', 7);

-- ペット用品のサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('ドッグ', 'dog-supplies', '犬用品', 'https://picsum.photos/400/300?random=39', 8),
('キャット', 'cat-supplies', '猫用品', 'https://picsum.photos/400/300?random=40', 8),
('小動物', 'small-pets', '小動物用品', 'https://picsum.photos/400/300?random=41', 8),
('アクアリウム', 'aquarium', '熱帯魚、水槽用品', 'https://picsum.photos/400/300?random=42', 8);

-- 商品のテストデータ
INSERT INTO products (name, description, price, category_id, sku, image_url) VALUES
-- スマートフォン (9)
('Acme スマートフォン', 'Acmeの最新スマートフォン', 89800, 9, 'ACME-SP-001', 'https://picsum.photos/400/300?random=43'),
('Ultra スマートフォン', '最新のAI機能搭載スマートフォン', 99800, 9, 'ULT-SP-002', 'https://picsum.photos/400/300?random=44'),
('Lite スマートフォン', 'エントリーモデルスマートフォン', 49800, 9, 'LIT-SP-003', 'https://picsum.photos/400/300?random=45'),
('Pro スマートフォン', 'プロフェッショナル向けスマートフォン', 129800, 9, 'PRO-SP-004', 'https://picsum.photos/400/300?random=46'),
('Mini スマートフォン', 'コンパクトサイズスマートフォン', 39800, 9, 'MIN-SP-005', 'https://picsum.photos/400/300?random=47'),
('Gaming スマートフォン', 'ゲーミング特化型スマートフォン', 149800, 9, 'GAM-SP-006', 'https://picsum.photos/400/300?random=48'),
('Camera スマートフォン', 'カメラ機能特化型スマートフォン', 119800, 9, 'CAM-SP-007', 'https://picsum.photos/400/300?random=49'),
('Business スマートフォン', 'ビジネス向けスマートフォン', 109800, 9, 'BUS-SP-008', 'https://picsum.photos/400/300?random=50'),
('Eco スマートフォン', '環境配慮型スマートフォン', 79800, 9, 'ECO-SP-009', 'https://picsum.photos/400/300?random=51'),
('Fold スマートフォン', '折りたたみ式スマートフォン', 159800, 9, 'FLD-SP-010', 'https://picsum.photos/400/300?random=52'),

-- ノートパソコン (10)
('Zenith ノートPC', '高性能ノートパソコン', 125000, 10, 'ZNT-NP-001', 'https://picsum.photos/400/300?random=53'),
('ProBook ノートPC', 'クリエイター向け高性能ノートPC', 158000, 10, 'PRB-NP-002', 'https://picsum.photos/400/300?random=54'),
('Student ノートPC', '学生向けコストパフォーマンス重視のノートPC', 78000, 10, 'STU-NP-003', 'https://picsum.photos/400/300?random=55'),
('Gaming ノートPC', 'ゲーミング特化型ノートPC', 198000, 10, 'GAM-NP-004', 'https://picsum.photos/400/300?random=56'),
('UltraBook', '超薄型ノートPC', 138000, 10, 'ULT-NP-005', 'https://picsum.photos/400/300?random=57'),
('WorkStation', 'ワークステーション向けノートPC', 228000, 10, 'WRK-NP-006', 'https://picsum.photos/400/300?random=58'),
('Convertible', '2in1コンバーチブルノートPC', 148000, 10, 'CON-NP-007', 'https://picsum.photos/400/300?random=59'),
('Business ノートPC', 'ビジネス向けノートPC', 128000, 10, 'BUS-NP-008', 'https://picsum.photos/400/300?random=60'),
('Budget ノートPC', 'エントリーモデルノートPC', 58000, 10, 'BGT-NP-009', 'https://picsum.photos/400/300?random=61'),
('Premium ノートPC', 'プレミアムモデルノートPC', 258000, 10, 'PRM-NP-010', 'https://picsum.photos/400/300?random=62'),

-- タブレット (11)
('Quantum タブレット', '10インチタブレット', 45000, 11, 'QNT-TB-001', 'https://picsum.photos/400/300?random=63'),
('Mini タブレット', '8インチコンパクトタブレット', 35000, 11, 'MIN-TB-002', 'https://picsum.photos/400/300?random=64'),
('Pro タブレット', '12インチプロフェッショナルタブレット', 68000, 11, 'PRO-TB-003', 'https://picsum.photos/400/300?random=65'),
('Kids タブレット', '子供向けタブレット', 25000, 11, 'KID-TB-004', 'https://picsum.photos/400/300?random=66'),
('Drawing タブレット', 'ペンタブレット機能付きタブレット', 78000, 11, 'DRW-TB-005', 'https://picsum.photos/400/300?random=67'),
('Reading タブレット', '読書特化型タブレット', 42000, 11, 'RED-TB-006', 'https://picsum.photos/400/300?random=68'),
('Business タブレット', 'ビジネス向けタブレット', 58000, 11, 'BUS-TB-007', 'https://picsum.photos/400/300?random=69'),
('Gaming タブレット', 'ゲーミング特化型タブレット', 88000, 11, 'GAM-TB-008', 'https://picsum.photos/400/300?random=70'),
('Budget タブレット', 'エントリーモデルタブレット', 28000, 11, 'BGT-TB-009', 'https://picsum.photos/400/300?random=71'),
('Premium タブレット', 'プレミアムモデルタブレット', 98000, 11, 'PRM-TB-010', 'https://picsum.photos/400/300?random=72'),

-- デスクトップPC (12)
('Gaming Desktop', 'ハイエンドゲーミングPC', 298000, 12, 'GAM-DT-001', 'https://picsum.photos/400/300?random=73'),
('WorkStation Desktop', 'ワークステーション向けPC', 398000, 12, 'WRK-DT-002', 'https://picsum.photos/400/300?random=74'),
('Budget Desktop', 'エントリーモデルPC', 78000, 12, 'BGT-DT-003', 'https://picsum.photos/400/300?random=75'),
('All-in-One Desktop', 'オールインワンPC', 158000, 12, 'AIO-DT-004', 'https://picsum.photos/400/300?random=76'),
('Mini Desktop', 'コンパクトデスクトップPC', 98000, 12, 'MIN-DT-005', 'https://picsum.photos/400/300?random=77'),
('Business Desktop', 'ビジネス向けPC', 128000, 12, 'BUS-DT-006', 'https://picsum.photos/400/300?random=78'),
('Creative Desktop', 'クリエイター向けPC', 228000, 12, 'CRE-DT-007', 'https://picsum.photos/400/300?random=79'),
('Server Desktop', 'サーバー向けPC', 498000, 12, 'SRV-DT-008', 'https://picsum.photos/400/300?random=80'),
('Home Desktop', 'ホームユース向けPC', 98000, 12, 'HOM-DT-009', 'https://picsum.photos/400/300?random=81'),
('Premium Desktop', 'プレミアムモデルPC', 598000, 12, 'PRM-DT-010', 'https://picsum.photos/400/300?random=82'),

-- カメラ (13)
('DSLR Camera', 'プロフェッショナル一眼レフカメラ', 198000, 13, 'DSL-CM-001', 'https://picsum.photos/400/300?random=83'),
('Mirrorless Camera', 'ミラーレス一眼カメラ', 158000, 13, 'MRL-CM-002', 'https://picsum.photos/400/300?random=84'),
('Action Camera', 'アクションカメラ', 48000, 13, 'ACT-CM-003', 'https://picsum.photos/400/300?random=85'),
('Instant Camera', 'インスタントカメラ', 9800, 13, 'INS-CM-004', 'https://picsum.photos/400/300?random=86'),
('Video Camera', 'ビデオカメラ', 128000, 13, 'VID-CM-005', 'https://picsum.photos/400/300?random=87'),
('Compact Camera', 'コンパクトデジタルカメラ', 28000, 13, 'CMP-CM-006', 'https://picsum.photos/400/300?random=88'),
('Drone Camera', 'ドローンカメラ', 98000, 13, 'DRN-CM-007', 'https://picsum.photos/400/300?random=89'),
('Security Camera', 'セキュリティカメラ', 15800, 13, 'SEC-CM-008', 'https://picsum.photos/400/300?random=90'),
('360 Camera', '360度カメラ', 68000, 13, '360-CM-009', 'https://picsum.photos/400/300?random=91'),
('Film Camera', 'フィルムカメラ', 48000, 13, 'FLM-CM-010', 'https://picsum.photos/400/300?random=92'),

-- オーディオ (14)
('Wireless Earbuds', 'ワイヤレスイヤホン', 19800, 14, 'WLS-EB-001', 'https://picsum.photos/400/300?random=93'),
('Noise Cancelling Headphones', 'ノイズキャンセリングヘッドホン', 29800, 14, 'NCH-HP-001', 'https://picsum.photos/400/300?random=94'),
('Bluetooth Speaker', 'Bluetoothスピーカー', 12800, 14, 'BTS-SP-001', 'https://picsum.photos/400/300?random=95'),
('Sound Bar', 'サウンドバー', 45800, 14, 'SND-SB-001', 'https://picsum.photos/400/300?random=96'),
('Studio Monitor', 'スタジオモニター', 98000, 14, 'STM-MN-001', 'https://picsum.photos/400/300?random=97'),
('Portable Speaker', 'ポータブルスピーカー', 15800, 14, 'PRT-SP-001', 'https://picsum.photos/400/300?random=98'),
('Gaming Headset', 'ゲーミングヘッドセット', 24800, 14, 'GAM-HS-001', 'https://picsum.photos/400/300?random=99'),
('Wireless Headphones', 'ワイヤレスヘッドホン', 32800, 14, 'WLS-HP-001', 'https://picsum.photos/400/300?random=100'),
('Earphones', 'イヤホン', 9800, 14, 'EAR-EP-001', 'https://picsum.photos/400/300?random=101'),
('Home Theater System', 'ホームシアターシステム', 128000, 14, 'HTS-SY-001', 'https://picsum.photos/400/300?random=102'),

-- メンズ (15)
('メンズカジュアルシャツ', '綿100%のカジュアルシャツ', 3800, 15, 'MCS-001', 'https://picsum.photos/400/300?random=103'),
('メンズデニムパンツ', 'スリムフィットデニム', 5800, 15, 'MDP-002', 'https://picsum.photos/400/300?random=104'),
('メンズジャケット', 'カジュアルジャケット', 12800, 15, 'MJK-003', 'https://picsum.photos/400/300?random=105'),
('メンズTシャツ', 'カジュアルTシャツ', 2800, 15, 'MTS-004', 'https://picsum.photos/400/300?random=106'),
('メンズセーター', 'ウールセーター', 7800, 15, 'MSW-005', 'https://picsum.photos/400/300?random=107'),
('メンズコート', 'ウールコート', 25800, 15, 'MCT-006', 'https://picsum.photos/400/300?random=108'),
('メンズスニーカー', 'カジュアルスニーカー', 9800, 15, 'MSN-007', 'https://picsum.photos/400/300?random=109'),
('メンズブーツ', 'レザーブーツ', 15800, 15, 'MBT-008', 'https://picsum.photos/400/300?random=110'),
('メンズベルト', 'レザーベルト', 3800, 15, 'MBL-009', 'https://picsum.photos/400/300?random=111'),
('メンズバッグ', 'ビジネスバッグ', 12800, 15, 'MBG-010', 'https://picsum.photos/400/300?random=112'),

-- レディース (16)
('レディースブラウス', 'エレガントなデザインのブラウス', 4200, 16, 'LDB-001', 'https://picsum.photos/400/300?random=113'),
('レディースワンピース', 'シンプルでエレガントなワンピース', 7800, 16, 'LDW-002', 'https://picsum.photos/400/300?random=114'),
('レディーススカート', 'オフィスカジュアルスカート', 5800, 16, 'LDS-003', 'https://picsum.photos/400/300?random=115'),
('レディースパンツ', 'スリムフィットパンツ', 4800, 16, 'LDP-004', 'https://picsum.photos/400/300?random=116'),
('レディースセーター', 'カシミアセーター', 9800, 16, 'LDS-005', 'https://picsum.photos/400/300?random=117'),
('レディースコート', 'ウールコート', 22800, 16, 'LDC-006', 'https://picsum.photos/400/300?random=118'),
('レディースブーツ', 'レザーブーツ', 14800, 16, 'LDB-007', 'https://picsum.photos/400/300?random=119'),
('レディースバッグ', 'トートバッグ', 15800, 16, 'LDB-008', 'https://picsum.photos/400/300?random=120'),
('レディースアクセサリー', 'ネックレス', 3800, 16, 'LDA-009', 'https://picsum.photos/400/300?random=121'),
('レディースシューズ', 'パンプス', 9800, 16, 'LDS-010', 'https://picsum.photos/400/300?random=122'),

-- 書籍 (19-22)
('プログラミング入門', 'プログラミングの基礎を学ぶ', 2800, 19, 'BK-PRG-001', 'https://picsum.photos/400/300?random=123'),
('データベース設計入門', '実践的なデータベース設計の基礎', 3200, 19, 'BK-DB-002', 'https://picsum.photos/400/300?random=124'),
('クラウドコンピューティング入門', 'クラウドサービスの基礎知識', 3500, 19, 'BK-CC-003', 'https://picsum.photos/400/300?random=125'),
('AIと機械学習', '人工知能と機械学習の基礎', 3800, 19, 'BK-AI-004', 'https://picsum.photos/400/300?random=126'),
('Webデザイン入門', 'モダンなWebデザインの基礎', 2800, 19, 'BK-WD-005', 'https://picsum.photos/400/300?random=127'),
('マーケティング戦略', 'デジタルマーケティングの実践', 3200, 19, 'BK-MK-006', 'https://picsum.photos/400/300?random=128'),
('リーダーシップ論', '効果的なリーダーシップの実践', 2800, 19, 'BK-LD-007', 'https://picsum.photos/400/300?random=129'),
('財務分析入門', '企業財務の基礎知識', 3500, 19, 'BK-FN-008', 'https://picsum.photos/400/300?random=130'),
('プロジェクトマネジメント', 'プロジェクト管理の実践', 3200, 19, 'BK-PM-009', 'https://picsum.photos/400/300?random=131'),
('起業家精神', 'スタートアップの成功法則', 2800, 19, 'BK-EN-010', 'https://picsum.photos/400/300?random=132'),

-- ホーム＆キッチン (23-26)
('キッチンミキサー', '多機能キッチンミキサー', 6500, 23, 'KM-001', 'https://picsum.photos/400/300?random=133'),
('電子レンジ', 'コンパクト電子レンジ', 12800, 23, 'KM-002', 'https://picsum.photos/400/300?random=134'),
('コーヒーメーカー', 'ドリップコーヒーメーカー', 9800, 23, 'KM-003', 'https://picsum.photos/400/300?random=135'),
('トースター', '2枚焼きトースター', 4800, 23, 'KM-004', 'https://picsum.photos/400/300?random=136'),
('フードプロセッサー', '多機能フードプロセッサー', 7800, 23, 'KM-005', 'https://picsum.photos/400/300?random=137'),
('ダイニングテーブル', '4人用ダイニングテーブル', 45800, 24, 'FRN-DT-001', 'https://picsum.photos/400/300?random=138'),
('ソファ', '3人掛けソファ', 89800, 24, 'FRN-SF-001', 'https://picsum.photos/400/300?random=139'),
('ベッド', 'シングルベッド', 45800, 24, 'FRN-BD-001', 'https://picsum.photos/400/300?random=140'),
('本棚', '5段本棚', 15800, 24, 'FRN-BS-001', 'https://picsum.photos/400/300?random=141'),
('テレビ台', 'モダンテレビ台', 25800, 24, 'FRN-TV-001', 'https://picsum.photos/400/300?random=142');

-- 在庫データ
INSERT INTO inventory (product_id, quantity) VALUES
(1, 50),  -- Acme スマートフォン
(2, 35),  -- Ultra スマートフォン
(3, 60),  -- Lite スマートフォン
(4, 25),  -- Pro スマートフォン
(5, 40),  -- Mini スマートフォン
(6, 20),  -- Gaming スマートフォン
(7, 30),  -- Camera スマートフォン
(8, 45),  -- Business スマートフォン
(9, 55),  -- Eco スマートフォン
(10, 25), -- Fold スマートフォン
(11, 25), -- Zenith ノートPC
(12, 15), -- ProBook ノートPC
(13, 30), -- Student ノートPC
(14, 20), -- Gaming ノートPC
(15, 25), -- UltraBook
(16, 10), -- WorkStation
(17, 20), -- Convertible
(18, 35), -- Business ノートPC
(19, 40), -- Budget ノートPC
(20, 15), -- Premium ノートPC
(21, 30), -- Quantum タブレット
(22, 40), -- Mini タブレット
(23, 25), -- Pro タブレット
(24, 35), -- Kids タブレット
(25, 20), -- Drawing タブレット
(26, 30), -- Reading タブレット
(27, 25), -- Business タブレット
(28, 15), -- Gaming タブレット
(29, 40), -- Budget タブレット
(30, 20), -- Premium タブレット
(31, 15), -- Gaming Desktop
(32, 10), -- WorkStation Desktop
(33, 30), -- Budget Desktop
(34, 20), -- All-in-One Desktop
(35, 25), -- Mini Desktop
(36, 30), -- Business Desktop
(37, 15), -- Creative Desktop
(38, 10), -- Server Desktop
(39, 25), -- Home Desktop
(40, 8),  -- Premium Desktop
(41, 20), -- DSLR Camera
(42, 25), -- Mirrorless Camera
(43, 35), -- Action Camera
(44, 50), -- Instant Camera
(45, 20), -- Video Camera
(46, 30), -- Compact Camera
(47, 25), -- Drone Camera
(48, 40), -- Security Camera
(49, 20), -- 360 Camera
(50, 15), -- Film Camera
(51, 45), -- Wireless Earbuds
(52, 30), -- Noise Cancelling Headphones
(53, 40), -- Bluetooth Speaker
(54, 25), -- Sound Bar
(55, 15), -- Studio Monitor
(56, 35), -- Portable Speaker
(57, 30), -- Gaming Headset
(58, 25), -- Wireless Headphones
(59, 50), -- Earphones
(60, 10), -- Home Theater System
(61, 100),-- メンズカジュアルシャツ
(62, 120),-- メンズデニムパンツ
(63, 70), -- メンズジャケット
(64, 150),-- メンズTシャツ
(65, 80), -- メンズセーター
(66, 50), -- メンズコート
(67, 60), -- メンズスニーカー
(68, 40), -- メンズブーツ
(69, 90), -- メンズベルト
(70, 55), -- メンズバッグ
(71, 90), -- レディースブラウス
(72, 85), -- レディースワンピース
(73, 95), -- レディーススカート
(74, 110),-- レディースパンツ
(75, 75), -- レディースセーター
(76, 45), -- レディースコート
(77, 50), -- レディースブーツ
(78, 65), -- レディースバッグ
(79, 120),-- レディースアクセサリー
(80, 80), -- レディースシューズ
(81, 60), -- プログラミング入門
(82, 45), -- データベース設計入門
(83, 50), -- クラウドコンピューティング入門
(84, 40), -- AIと機械学習
(85, 55), -- Webデザイン入門
(86, 45), -- マーケティング戦略
(87, 50), -- リーダーシップ論
(88, 40), -- 財務分析入門
(89, 45), -- プロジェクトマネジメント
(90, 50), -- 起業家精神
(91, 15), -- キッチンミキサー
(92, 20), -- 電子レンジ
(93, 25), -- コーヒーメーカー
(94, 30), -- トースター
(95, 20), -- フードプロセッサー
(96, 10), -- ダイニングテーブル
(97, 8),  -- ソファ
(98, 12), -- ベッド
(99, 20), -- 本棚
(100, 15); -- テレビ台
