-- カテゴリーのテストデータ
INSERT INTO categories (name, slug, description, image_url) VALUES
('電子機器', 'electronics', '最新の電子機器製品', 'https://example.com/images/electronics.jpg'),
('洋服', 'clothing', 'トレンドの洋服コレクション', 'https://example.com/images/clothing.jpg'),
('書籍', 'books', '書籍、電子書籍など', 'https://example.com/images/books.jpg'),
('ホーム＆キッチン', 'home-kitchen', '家庭用品とキッチン用品', 'https://example.com/images/home-kitchen.jpg');

-- 電子機器のサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('スマートフォン', 'smartphones', '最新のスマートフォン', 'https://example.com/images/smartphones.jpg', 1),
('ノートパソコン', 'laptops', '高性能ノートPC', 'https://example.com/images/laptops.jpg', 1),
('タブレット', 'tablets', 'タブレット端末', 'https://example.com/images/tablets.jpg', 1);

-- 洋服のサブカテゴリー
INSERT INTO categories (name, slug, description, image_url, parent_id) VALUES
('メンズ', 'mens-clothing', 'メンズファッション', 'https://example.com/images/mens-clothing.jpg', 2),
('レディース', 'womens-clothing', 'レディースファッション', 'https://example.com/images/womens-clothing.jpg', 2);

-- 商品のテストデータ
INSERT INTO products (name, description, price, category_id, sku, image_url) VALUES
('Acme スマートフォン', 'Acmeの最新スマートフォン', 89800, 5, 'ACME-SP-001', 'https://example.com/images/acme-smartphone.jpg'),
('Zenith ノートPC', '高性能ノートパソコン', 125000, 6, 'ZNT-NP-001', 'https://example.com/images/zenith-laptop.jpg'),
('Quantum タブレット', '10インチタブレット', 45000, 7, 'QNT-TB-001', 'https://example.com/images/quantum-tablet.jpg'),
('メンズカジュアルシャツ', '綿100%のカジュアルシャツ', 3800, 8, 'MCS-001', 'https://example.com/images/mens-casual-shirt.jpg'),
('レディースブラウス', 'エレガントなデザインのブラウス', 4200, 9, 'LDB-001', 'https://example.com/images/ladies-blouse.jpg'),
('プログラミング入門', 'プログラミングの基礎を学ぶ', 2800, 3, 'BK-PRG-001', 'https://example.com/images/programming-book.jpg'),
('キッチンミキサー', '多機能キッチンミキサー', 6500, 4, 'KM-001', 'https://example.com/images/kitchen-mixer.jpg');

-- 在庫データ
INSERT INTO inventory (product_id, quantity) VALUES
(1, 50),  -- Acme スマートフォン
(2, 25),  -- Zenith ノートPC
(3, 30),  -- Quantum タブレット
(4, 100), -- メンズカジュアルシャツ
(5, 80),  -- レディースブラウス
(6, 60),  -- プログラミング入門
(7, 15);  -- キッチンミキサー
