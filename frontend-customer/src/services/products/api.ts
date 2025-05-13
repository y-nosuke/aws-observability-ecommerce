import { Product, Category, SortOption } from './types';

// モック商品データ
const MOCK_PRODUCTS: Product[] = [
  {
    id: "1",
    name: "超高性能ノートPC",
    description: "最新のプロセッサ・16GB RAM・高速 SSD 搭載",
    price: 198000,
    salePrice: null,
    category: "computers",
    isNew: true,
    imageUrl: null,
  },
  {
    id: "2",
    name: "ワイヤレスイヤホン",
    description: "ノイズキャンセリング機能付き高音質イヤホン",
    price: 19800,
    salePrice: 14800,
    category: "audio",
    isNew: false,
    imageUrl: null,
  },
  {
    id: "3",
    name: "高画質タブレット",
    description: "10.2インチディスプレイ、長持ちバッテリー搭載",
    price: 36800,
    salePrice: null,
    category: "computers",
    isNew: false,
    imageUrl: null,
  },
  {
    id: "4",
    name: "スマートウォッチ",
    description: "健康管理機能付き、長持ちバッテリー",
    price: 32800,
    salePrice: null,
    category: "wearables",
    isNew: false,
    imageUrl: null,
  },
  {
    id: "5",
    name: "4Kモニター 27インチ",
    description: "鮮明な映像、広色域ディスプレイ",
    price: 49800,
    salePrice: 44800,
    category: "computers",
    isNew: true,
    imageUrl: null,
  },
  {
    id: "6",
    name: "ハイエンドスマートフォン",
    description: "最新チップ、高性能カメラ、大容量バッテリー",
    price: 128000,
    salePrice: null,
    category: "phones",
    isNew: true,
    imageUrl: null,
  },
  {
    id: "7",
    name: "Bluetoothスピーカー",
    description: "防水機能付き、高音質ポータブルスピーカー",
    price: 12800,
    salePrice: 9800,
    category: "audio",
    isNew: false,
    imageUrl: null,
  },
  {
    id: "8",
    name: "ワイヤレスマウス",
    description: "高精度センサー、長持ちバッテリー",
    price: 6800,
    salePrice: null,
    category: "accessories",
    isNew: false,
    imageUrl: null,
  },
];

// カテゴリーデータ
const CATEGORIES: Category[] = [
  { id: "all", name: "すべて" },
  { id: "computers", name: "コンピューター" },
  { id: "phones", name: "スマートフォン" },
  { id: "audio", name: "オーディオ" },
  { id: "wearables", name: "ウェアラブル" },
  { id: "accessories", name: "アクセサリー" },
];

// 商品一覧の取得
export async function fetchProducts(): Promise<Product[]> {
  // 実際のAPIからデータを取得する場合はここで外部APIを呼び出す
  // 例: const response = await fetch('api/products');
  // return response.json();
  
  // モックデータを返す
  return Promise.resolve([...MOCK_PRODUCTS]);
}

// 人気商品の取得（TOP N件）
export async function fetchPopularProducts(limit: number = 4): Promise<Product[]> {
  // モックデータを返す（実際のアプリではアクセス数や売上数でソートなど）
  return Promise.resolve([...MOCK_PRODUCTS].slice(0, limit));
}

// カテゴリー一覧の取得
export async function fetchCategories(): Promise<Category[]> {
  return Promise.resolve([...CATEGORIES]);
}

// 商品をIDで取得
export async function fetchProductById(id: string): Promise<Product | null> {
  const product = MOCK_PRODUCTS.find(p => p.id === id);
  return Promise.resolve(product || null);
}

// 商品のフィルタリングと並び替え
export function filterAndSortProducts(
  products: Product[],
  category: string = 'all',
  sortOption: SortOption = 'recommended'
): Product[] {
  // カテゴリーでフィルタリング
  const filtered = category === 'all' 
    ? products 
    : products.filter(product => product.category === category);
  
  // 並び替え
  return sortProducts(filtered, sortOption);
}

// 商品の並び替え
export function sortProducts(
  products: Product[],
  sortOption: SortOption = 'recommended'
): Product[] {
  switch (sortOption) {
    case 'price-asc':
      return [...products].sort(
        (a, b) => (a.salePrice || a.price) - (b.salePrice || b.price)
      );
    case 'price-desc':
      return [...products].sort(
        (a, b) => (b.salePrice || b.price) - (a.salePrice || a.price)
      );
    case 'newest':
      return [...products].sort((a, b) => (a.isNew ? -1 : b.isNew ? 1 : 0));
    default: // recommended
      return products;
  }
}
