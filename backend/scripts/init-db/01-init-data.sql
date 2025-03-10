-- カテゴリーの初期データ
INSERT INTO categories (name, description) VALUES
('エレクトロニクス', '最新のガジェットと電子機器'),
('ファッション', '洋服、靴、アクセサリー'),
('ホーム＆キッチン', '家庭用品とキッチン用品');

-- エレクトロニクスのサブカテゴリー
INSERT INTO categories (name, description, parent_id) VALUES
('スマートフォン', 'スマートフォンとアクセサリー', 1),
('ノートパソコン', 'ノートブックとアクセサリー', 1);

-- 商品の初期データ
INSERT INTO products (name, description, price, image_url, category_id) VALUES
('スマートフォンX', '最新のスマートフォン、高性能カメラと大容量バッテリー', 89000, '/images/smartphone-x.jpg', 4),
('ノートパソコンY', 'ビジネス向け高性能ノートパソコン', 120000, '/images/laptop-y.jpg', 5),
('デニムジャケット', 'カジュアルなデニムジャケット、オールシーズン着用可能', 7500, '/images/denim-jacket.jpg', 2),
('キッチンミキサー', 'プロ仕様の高性能キッチンミキサー', 15000, '/images/kitchen-mixer.jpg', 3);

-- 在庫データ
INSERT INTO inventory (product_id, quantity) VALUES
(1, 50),
(2, 30),
(3, 100),
(4, 25);

-- 管理者ユーザー (パスワード: admin123)
INSERT INTO admin_users (username, password_hash, email) VALUES
('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'admin@example.com');
