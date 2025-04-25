-- 在庫データの削除
DELETE FROM inventory;

-- 商品データの削除
DELETE FROM products;

-- サブカテゴリーの削除
DELETE FROM categories WHERE parent_id IS NOT NULL;

-- 親カテゴリーの削除
DELETE FROM categories WHERE parent_id IS NULL;
