-- 基本的な初期化スクリプト

-- データベースが存在しない場合は作成
CREATE DATABASE IF NOT EXISTS `ecommerce`;

-- 権限の設定
GRANT ALL PRIVILEGES ON `ecommerce`.* TO 'ecommerce_user'@'%';
FLUSH PRIVILEGES;

-- ecommerceデータベースを選択
USE `ecommerce`;

-- 基本的なテストテーブル作成
CREATE TABLE IF NOT EXISTS `test` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- テストデータ挿入
INSERT INTO `test` (`name`) VALUES ('This is a test');
