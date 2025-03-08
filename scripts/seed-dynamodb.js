// DynamoDBにサンプルデータを投入するスクリプト（AWS SDK v3版）
const { DynamoDBClient } = require("@aws-sdk/client-dynamodb");
const { DynamoDBDocumentClient, PutCommand } = require("@aws-sdk/lib-dynamodb");

// リージョン設定（環境に合わせて変更）
const region = process.env.AWS_REGION || "ap-northeast-1";

// DynamoDB クライアントの作成
const client = new DynamoDBClient({ region });
// DynamoDBドキュメントクライアントの作成（高レベルのドキュメント操作用）
const docClient = DynamoDBDocumentClient.from(client);

// プロジェクト名（環境に合わせて変更）
const PROJECT_PREFIX = process.env.PROJECT_PREFIX || "ecommerce-observability";

// テーブル名
const PRODUCTS_TABLE = `${PROJECT_PREFIX}-products`;
const CATEGORIES_TABLE = `${PROJECT_PREFIX}-categories`;

// サンプルカテゴリデータ
const categories = [
  {
    id: "cat-001",
    name: "家電製品",
    description: "最新の家電製品を揃えています。",
  },
  {
    id: "cat-002",
    name: "ファッション",
    description: "流行のファッションアイテムを販売しています。",
  },
  {
    id: "cat-003",
    name: "食品・飲料",
    description: "新鮮な食材と美味しい飲み物をお届けします。",
  },
  {
    id: "cat-004",
    name: "本・雑誌",
    description: "様々なジャンルの書籍を取り揃えています。",
  },
  {
    id: "cat-005",
    name: "スポーツ用品",
    description: "アウトドアやスポーツに必要なアイテムが豊富です。",
  },
];

// サンプル商品データ
const products = [
  {
    id: "prod-001",
    name: "4Kテレビ 55インチ",
    description:
      "高精細な4K解像度で、映像をより鮮明に楽しめる55インチの大型テレビです。",
    price: 89800,
    category_id: "cat-001",
    image_url: "https://example.com/images/tv.jpg",
  },
  {
    id: "prod-002",
    name: "ノートパソコン 15.6インチ",
    description:
      "最新のCPUとグラフィックスカードを搭載したハイスペックノートパソコンです。",
    price: 128000,
    category_id: "cat-001",
    image_url: "https://example.com/images/laptop.jpg",
  },
  {
    id: "prod-003",
    name: "ワイヤレスイヤホン",
    description: "ノイズキャンセリング機能付きの高音質ワイヤレスイヤホンです。",
    price: 19800,
    category_id: "cat-001",
    image_url: "https://example.com/images/earphone.jpg",
  },
  {
    id: "prod-004",
    name: "スマートウォッチ",
    description: "心拍数や睡眠の質を測定できる多機能スマートウォッチです。",
    price: 32800,
    category_id: "cat-001",
    image_url: "https://example.com/images/smartwatch.jpg",
  },
  {
    id: "prod-005",
    name: "デニムジャケット",
    description: "カジュアルなスタイルにぴったりの定番デニムジャケットです。",
    price: 8900,
    category_id: "cat-002",
    image_url: "https://example.com/images/denim.jpg",
  },
  {
    id: "prod-006",
    name: "レザースニーカー",
    description: "上質な本革を使用した、履き心地の良いスニーカーです。",
    price: 12800,
    category_id: "cat-002",
    image_url: "https://example.com/images/sneaker.jpg",
  },
  {
    id: "prod-007",
    name: "カシミヤセーター",
    description: "柔らかく暖かいカシミヤ100%のセーターです。",
    price: 24800,
    category_id: "cat-002",
    image_url: "https://example.com/images/sweater.jpg",
  },
  {
    id: "prod-008",
    name: "有機野菜セット",
    description: "農薬を使用せずに栽培された新鮮な有機野菜の詰め合わせです。",
    price: 3980,
    category_id: "cat-003",
    image_url: "https://example.com/images/vegetables.jpg",
  },
  {
    id: "prod-009",
    name: "プレミアムコーヒー豆",
    description: "厳選された産地から直輸入した、香り高い上質なコーヒー豆です。",
    price: 2480,
    category_id: "cat-003",
    image_url: "https://example.com/images/coffee.jpg",
  },
  {
    id: "prod-010",
    name: "高級ワインセット",
    description: "フランス産の厳選された赤ワイン5本セットです。",
    price: 32800,
    category_id: "cat-003",
    image_url: "https://example.com/images/wine.jpg",
  },
  {
    id: "prod-011",
    name: "ベストセラー小説",
    description: "話題のミステリー小説の最新作です。",
    price: 1580,
    category_id: "cat-004",
    image_url: "https://example.com/images/novel.jpg",
  },
  {
    id: "prod-012",
    name: "ビジネス書",
    description:
      "成功するビジネスパーソンのための実践的なノウハウが詰まった一冊です。",
    price: 1980,
    category_id: "cat-004",
    image_url: "https://example.com/images/business.jpg",
  },
  {
    id: "prod-013",
    name: "料理本",
    description: "初心者でも作れる簡単レシピが100種類以上掲載されています。",
    price: 2480,
    category_id: "cat-004",
    image_url: "https://example.com/images/cookbook.jpg",
  },
  {
    id: "prod-014",
    name: "ヨガマット",
    description: "滑りにくく、クッション性に優れたプロ仕様のヨガマットです。",
    price: 3980,
    category_id: "cat-005",
    image_url: "https://example.com/images/yogamat.jpg",
  },
  {
    id: "prod-015",
    name: "ランニングシューズ",
    description: "衝撃吸収性に優れた、長距離ランにも最適なシューズです。",
    price: 12800,
    category_id: "cat-005",
    image_url: "https://example.com/images/shoes.jpg",
  },
];

// DynamoDBにデータを投入する関数
const seedData = async () => {
  console.log("データ投入を開始します...");

  // カテゴリデータの投入
  console.log(`カテゴリデータを ${CATEGORIES_TABLE} テーブルに投入します...`);
  for (const category of categories) {
    const params = {
      TableName: CATEGORIES_TABLE,
      Item: category,
    };

    try {
      await docClient.send(new PutCommand(params));
      console.log(`カテゴリ "${category.name}" を追加しました`);
    } catch (error) {
      console.error(`カテゴリ "${category.name}" の追加に失敗しました:`, error);
    }
  }

  // 商品データの投入
  console.log(`商品データを ${PRODUCTS_TABLE} テーブルに投入します...`);
  for (const product of products) {
    const params = {
      TableName: PRODUCTS_TABLE,
      Item: product,
    };

    try {
      await docClient.send(new PutCommand(params));
      console.log(`商品 "${product.name}" を追加しました`);
    } catch (error) {
      console.error(`商品 "${product.name}" の追加に失敗しました:`, error);
    }
  }

  console.log("データ投入が完了しました");
};

// スクリプトの実行
seedData().catch((error) => {
  console.error("データ投入中にエラーが発生しました:", error);
  process.exit(1);
});
