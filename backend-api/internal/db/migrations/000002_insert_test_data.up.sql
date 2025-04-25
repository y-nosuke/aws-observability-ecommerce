-- カテゴリーのテストデータ
INSERT INTO categories (name, slug) VALUES
('電子機器', 'electronics'),
('洋服', 'clothing'),
('書籍', 'books'),
('ホーム＆キッチン', 'home-kitchen');

-- 電子機器のサブカテゴリー
INSERT INTO categories (name, slug, parent_id) VALUES
('スマートフォン', 'smartphones', 1),
('ノートパソコン', 'laptops', 1),
('タブレット', 'tablets', 1);

-- 洋服のサブカテゴリー
INSERT INTO categories (name, slug, parent_id) VALUES
('メンズ', 'mens-clothing', 2),
('レディース', 'womens-clothing', 2);

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
