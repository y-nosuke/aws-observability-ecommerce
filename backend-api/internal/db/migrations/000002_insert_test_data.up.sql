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
INSERT INTO products (name, description, price, sale_price, category_id, sku, image_url, is_new, is_featured) VALUES
-- スマートフォン (9)
('Acme スマートフォン', 'Acmeの最新スマートフォン', 89800, 79800, 9, 'ACME-SP-001', 'https://picsum.photos/400/300?random=43', true, true),
('Ultra スマートフォン', '最新のAI機能搭載スマートフォン', 99800, NULL, 9, 'ULT-SP-002', 'https://picsum.photos/400/300?random=44', true, true),
('Lite スマートフォン', 'エントリーモデルスマートフォン', 49800, 39800, 9, 'LIT-SP-003', 'https://picsum.photos/400/300?random=45', false, false),
('Pro スマートフォン', 'プロフェッショナル向けスマートフォン', 129800, NULL, 9, 'PRO-SP-004', 'https://picsum.photos/400/300?random=46', false, true),
('Mini スマートフォン', 'コンパクトサイズスマートフォン', 39800, 29800, 9, 'MIN-SP-005', 'https://picsum.photos/400/300?random=47', false, false),
('Gaming スマートフォン', 'ゲーミング特化型スマートフォン', 149800, NULL, 9, 'GAM-SP-006', 'https://picsum.photos/400/300?random=48', true, true),
('Camera スマートフォン', 'カメラ機能特化型スマートフォン', 119800, 99800, 9, 'CAM-SP-007', 'https://picsum.photos/400/300?random=49', false, false),
('Business スマートフォン', 'ビジネス向けスマートフォン', 109800, NULL, 9, 'BUS-SP-008', 'https://picsum.photos/400/300?random=50', false, false),
('Eco スマートフォン', '環境配慮型スマートフォン', 79800, 59800, 9, 'ECO-SP-009', 'https://picsum.photos/400/300?random=51', true, false),
('Fold スマートフォン', '折りたたみ式スマートフォン', 159800, NULL, 9, 'FLD-SP-010', 'https://picsum.photos/400/300?random=52', true, true),

-- ノートパソコン (10)
('Zenith ノートPC', '高性能ノートパソコン', 125000, NULL, 10, 'ZNT-NP-001', 'https://picsum.photos/400/300?random=53', false, true),
('ProBook ノートPC', 'クリエイター向け高性能ノートPC', 158000, 138000, 10, 'PRB-NP-002', 'https://picsum.photos/400/300?random=54', true, true),
('Student ノートPC', '学生向けコストパフォーマンス重視のノートPC', 78000, 58000, 10, 'STU-NP-003', 'https://picsum.photos/400/300?random=55', false, false),
('Gaming ノートPC', 'ゲーミング特化型ノートPC', 198000, NULL, 10, 'GAM-NP-004', 'https://picsum.photos/400/300?random=56', true, true),
('UltraBook', '超薄型ノートPC', 138000, 118000, 10, 'ULT-NP-005', 'https://picsum.photos/400/300?random=57', false, false),
('WorkStation', 'ワークステーション向けノートPC', 228000, NULL, 10, 'WRK-NP-006', 'https://picsum.photos/400/300?random=58', false, true),
('Convertible', '2in1コンバーチブルノートPC', 148000, 128000, 10, 'CON-NP-007', 'https://picsum.photos/400/300?random=59', true, false),
('Business ノートPC', 'ビジネス向けノートPC', 128000, NULL, 10, 'BUS-NP-008', 'https://picsum.photos/400/300?random=60', false, false),
('Budget ノートPC', 'エントリーモデルノートPC', 58000, 48000, 10, 'BGT-NP-009', 'https://picsum.photos/400/300?random=61', false, false),
('Premium ノートPC', 'プレミアムモデルノートPC', 258000, NULL, 10, 'PRM-NP-010', 'https://picsum.photos/400/300?random=62', true, true),

-- タブレット (11)
('Quantum タブレット', '10インチタブレット', 45000, 35000, 11, 'QNT-TB-001', 'https://picsum.photos/400/300?random=63', false, false),
('Mini タブレット', '8インチコンパクトタブレット', 35000, NULL, 11, 'MIN-TB-002', 'https://picsum.photos/400/300?random=64', true, false),
('Pro タブレット', '12インチプロフェッショナルタブレット', 68000, 58000, 11, 'PRO-TB-003', 'https://picsum.photos/400/300?random=65', false, true),
('Kids タブレット', '子供向けタブレット', 25000, 19800, 11, 'KID-TB-004', 'https://picsum.photos/400/300?random=66', false, false),
('Drawing タブレット', 'ペンタブレット機能付きタブレット', 78000, NULL, 11, 'DRW-TB-005', 'https://picsum.photos/400/300?random=67', true, true),
('Reading タブレット', '読書特化型タブレット', 42000, 32000, 11, 'RED-TB-006', 'https://picsum.photos/400/300?random=68', false, false),
('Business タブレット', 'ビジネス向けタブレット', 58000, NULL, 11, 'BUS-TB-007', 'https://picsum.photos/400/300?random=69', false, false),
('Gaming タブレット', 'ゲーミング特化型タブレット', 88000, 78000, 11, 'GAM-TB-008', 'https://picsum.photos/400/300?random=70', true, false),
('Budget タブレット', 'エントリーモデルタブレット', 28000, 19800, 11, 'BGT-TB-009', 'https://picsum.photos/400/300?random=71', false, false),
('Premium タブレット', 'プレミアムモデルタブレット', 98000, NULL, 11, 'PRM-TB-010', 'https://picsum.photos/400/300?random=72', true, true),

-- デスクトップPC (12)
('Gaming Desktop', 'ハイエンドゲーミングPC', 298000, NULL, 12, 'GAM-DT-001', 'https://picsum.photos/400/300?random=73', true, true),
('WorkStation Desktop', 'ワークステーション向けPC', 398000, 358000, 12, 'WRK-DT-002', 'https://picsum.photos/400/300?random=74', false, true),
('Budget Desktop', 'エントリーモデルPC', 78000, 58000, 12, 'BGT-DT-003', 'https://picsum.photos/400/300?random=75', false, false),
('All-in-One Desktop', 'オールインワンPC', 158000, NULL, 12, 'AIO-DT-004', 'https://picsum.photos/400/300?random=76', true, false),
('Mini Desktop', 'コンパクトデスクトップPC', 98000, 78000, 12, 'MIN-DT-005', 'https://picsum.photos/400/300?random=77', false, false),
('Business Desktop', 'ビジネス向けPC', 128000, NULL, 12, 'BUS-DT-006', 'https://picsum.photos/400/300?random=78', false, false),
('Creative Desktop', 'クリエイター向けPC', 228000, 198000, 12, 'CRE-DT-007', 'https://picsum.photos/400/300?random=79', true, true),
('Server Desktop', 'サーバー向けPC', 498000, NULL, 12, 'SRV-DT-008', 'https://picsum.photos/400/300?random=80', false, true),
('Home Desktop', 'ホームユース向けPC', 98000, 78000, 12, 'HOM-DT-009', 'https://picsum.photos/400/300?random=81', false, false),
('Premium Desktop', 'プレミアムモデルPC', 598000, NULL, 12, 'PRM-DT-010', 'https://picsum.photos/400/300?random=82', true, true),

-- カメラ (13)
('DSLR Camera', 'プロフェッショナル一眼レフカメラ', 198000, 178000, 13, 'DSL-CM-001', 'https://picsum.photos/400/300?random=83', false, true),
('Mirrorless Camera', 'ミラーレス一眼カメラ', 158000, NULL, 13, 'MRL-CM-002', 'https://picsum.photos/400/300?random=84', true, true),
('Action Camera', 'アクションカメラ', 48000, 38000, 13, 'ACT-CM-003', 'https://picsum.photos/400/300?random=85', false, false),
('Instant Camera', 'インスタントカメラ', 9800, 7800, 13, 'INS-CM-004', 'https://picsum.photos/400/300?random=86', false, false),
('Video Camera', 'ビデオカメラ', 128000, NULL, 13, 'VID-CM-005', 'https://picsum.photos/400/300?random=87', true, false),
('Compact Camera', 'コンパクトデジタルカメラ', 28000, 19800, 13, 'CMP-CM-006', 'https://picsum.photos/400/300?random=88', false, false),
('Drone Camera', 'ドローンカメラ', 98000, 88000, 13, 'DRN-CM-007', 'https://picsum.photos/400/300?random=89', true, true),
('Security Camera', 'セキュリティカメラ', 15800, 12800, 13, 'SEC-CM-008', 'https://picsum.photos/400/300?random=90', false, false),
('360 Camera', '360度カメラ', 68000, NULL, 13, '360-CM-009', 'https://picsum.photos/400/300?random=91', true, false),
('Film Camera', 'フィルムカメラ', 48000, 38000, 13, 'FLM-CM-010', 'https://picsum.photos/400/300?random=92', false, false),

-- オーディオ (14)
('Wireless Earbuds', 'ワイヤレスイヤホン', 19800, 15800, 14, 'WLS-EB-001', 'https://picsum.photos/400/300?random=93', true, true),
('Noise Cancelling Headphones', 'ノイズキャンセリングヘッドホン', 29800, NULL, 14, 'NCH-HP-001', 'https://picsum.photos/400/300?random=94', false, true),
('Bluetooth Speaker', 'Bluetoothスピーカー', 12800, 9800, 14, 'BTS-SP-001', 'https://picsum.photos/400/300?random=95', false, false),
('Sound Bar', 'サウンドバー', 45800, 39800, 14, 'SND-SB-001', 'https://picsum.photos/400/300?random=96', true, false),
('Studio Monitor', 'スタジオモニター', 98000, NULL, 14, 'STM-MN-001', 'https://picsum.photos/400/300?random=97', false, true),
('Portable Speaker', 'ポータブルスピーカー', 15800, 12800, 14, 'PRT-SP-001', 'https://picsum.photos/400/300?random=98', false, false),
('Gaming Headset', 'ゲーミングヘッドセット', 24800, 19800, 14, 'GAM-HS-001', 'https://picsum.photos/400/300?random=99', true, false),
('Wireless Headphones', 'ワイヤレスヘッドホン', 32800, NULL, 14, 'WLS-HP-001', 'https://picsum.photos/400/300?random=100', false, false),
('Earphones', 'イヤホン', 9800, 7800, 14, 'EAR-EP-001', 'https://picsum.photos/400/300?random=101', false, false),
('Home Theater System', 'ホームシアターシステム', 128000, 108000, 14, 'HTS-SY-001', 'https://picsum.photos/400/300?random=102', true, true),

-- メンズ (15)
('メンズカジュアルシャツ', '綿100%のカジュアルシャツ', 3800, 2980, 15, 'MCS-001', 'https://picsum.photos/400/300?random=103', false, false),
('メンズデニムパンツ', 'スリムフィットデニム', 5800, NULL, 15, 'MDP-002', 'https://picsum.photos/400/300?random=104', true, true),
('メンズジャケット', 'カジュアルジャケット', 12800, 9800, 15, 'MJK-003', 'https://picsum.photos/400/300?random=105', false, false),
('メンズTシャツ', 'カジュアルTシャツ', 2800, 1980, 15, 'MTS-004', 'https://picsum.photos/400/300?random=106', false, false),
('メンズセーター', 'ウールセーター', 7800, NULL, 15, 'MSW-005', 'https://picsum.photos/400/300?random=107', true, false),
('メンズコート', 'ウールコート', 25800, 22800, 15, 'MCT-006', 'https://picsum.photos/400/300?random=108', false, true),
('メンズスニーカー', 'カジュアルスニーカー', 9800, 7800, 15, 'MSN-007', 'https://picsum.photos/400/300?random=109', false, false),
('メンズブーツ', 'レザーブーツ', 15800, NULL, 15, 'MBT-008', 'https://picsum.photos/400/300?random=110', true, false),
('メンズベルト', 'レザーベルト', 3800, 2980, 15, 'MBL-009', 'https://picsum.photos/400/300?random=111', false, false),
('メンズバッグ', 'ビジネスバッグ', 12800, 9800, 15, 'MBG-010', 'https://picsum.photos/400/300?random=112', false, true),

-- レディース (16)
('レディースブラウス', 'エレガントなデザインのブラウス', 4200, 3280, 16, 'LDB-001', 'https://picsum.photos/400/300?random=113', false, false),
('レディースワンピース', 'シンプルでエレガントなワンピース', 7800, NULL, 16, 'LDW-002', 'https://picsum.photos/400/300?random=114', true, true),
('レディーススカート', 'オフィスカジュアルスカート', 5800, 4580, 16, 'LDS-003', 'https://picsum.photos/400/300?random=115', false, false),
('レディースパンツ', 'スリムフィットパンツ', 4800, NULL, 16, 'LDP-004', 'https://picsum.photos/400/300?random=116', true, false),
('レディースセーター', 'カシミアセーター', 9800, 7800, 16, 'LDS-005', 'https://picsum.photos/400/300?random=117', false, false),
('レディースコート', 'ウールコート', 22800, NULL, 16, 'LDC-006', 'https://picsum.photos/400/300?random=118', false, true),
('レディースブーツ', 'レザーブーツ', 14800, 12800, 16, 'LDB-007', 'https://picsum.photos/400/300?random=119', false, false),
('レディースバッグ', 'トートバッグ', 15800, NULL, 16, 'LDB-008', 'https://picsum.photos/400/300?random=120', true, false),
('レディースアクセサリー', 'ネックレス', 3800, 2980, 16, 'LDA-009', 'https://picsum.photos/400/300?random=121', false, false),
('レディースシューズ', 'パンプス', 9800, 7800, 16, 'LDS-010', 'https://picsum.photos/400/300?random=122', false, true),

-- 書籍 (19-22)
('プログラミング入門', 'プログラミングの基礎を学ぶ', 2800, 2280, 19, 'BK-PRG-001', 'https://picsum.photos/400/300?random=123', false, false),
('データベース設計入門', '実践的なデータベース設計の基礎', 3200, NULL, 19, 'BK-DB-002', 'https://picsum.photos/400/300?random=124', true, true),
('クラウドコンピューティング入門', 'クラウドサービスの基礎知識', 3500, 2980, 19, 'BK-CC-003', 'https://picsum.photos/400/300?random=125', false, false),
('AIと機械学習', '人工知能と機械学習の基礎', 3800, NULL, 19, 'BK-AI-004', 'https://picsum.photos/400/300?random=126', true, false),
('Webデザイン入門', 'モダンなWebデザインの基礎', 2800, 2280, 19, 'BK-WD-005', 'https://picsum.photos/400/300?random=127', false, false),
('マーケティング戦略', 'デジタルマーケティングの実践', 3200, NULL, 19, 'BK-MK-006', 'https://picsum.photos/400/300?random=128', false, true),
('リーダーシップ論', '効果的なリーダーシップの実践', 2800, 2280, 19, 'BK-LD-007', 'https://picsum.photos/400/300?random=129', false, false),
('財務分析入門', '企業財務の基礎知識', 3500, NULL, 19, 'BK-FN-008', 'https://picsum.photos/400/300?random=130', true, false),
('プロジェクトマネジメント', 'プロジェクト管理の実践', 3200, 2680, 19, 'BK-PM-009', 'https://picsum.photos/400/300?random=131', false, false),
('起業家精神', 'スタートアップの成功法則', 2800, NULL, 19, 'BK-EN-010', 'https://picsum.photos/400/300?random=132', false, true),

-- ホーム＆キッチン (23-26)
('キッチンミキサー', '多機能キッチンミキサー', 6500, 4980, 23, 'KM-001', 'https://picsum.photos/400/300?random=133', false, false),
('電子レンジ', 'コンパクト電子レンジ', 12800, NULL, 23, 'KM-002', 'https://picsum.photos/400/300?random=134', true, true),
('コーヒーメーカー', 'ドリップコーヒーメーカー', 9800, 7800, 23, 'KM-003', 'https://picsum.photos/400/300?random=135', false, false),
('トースター', '2枚焼きトースター', 4800, NULL, 23, 'KM-004', 'https://picsum.photos/400/300?random=136', true, false),
('フードプロセッサー', '多機能フードプロセッサー', 7800, 5800, 23, 'KM-005', 'https://picsum.photos/400/300?random=137', false, false),
('ダイニングテーブル', '4人用ダイニングテーブル', 45800, NULL, 24, 'FRN-DT-001', 'https://picsum.photos/400/300?random=138', false, true),
('ソファ', '3人掛けソファ', 89800, 79800, 24, 'FRN-SF-001', 'https://picsum.photos/400/300?random=139', true, false),
('ベッド', 'シングルベッド', 45800, NULL, 24, 'FRN-BD-001', 'https://picsum.photos/400/300?random=140', false, false),
('本棚', '5段本棚', 15800, 12800, 24, 'FRN-BS-001', 'https://picsum.photos/400/300?random=141', false, false),
('テレビ台', 'モダンテレビ台', 25800, NULL, 24, 'FRN-TV-001', 'https://picsum.photos/400/300?random=142', true, true);

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
