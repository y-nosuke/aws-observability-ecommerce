import { Product, Category, SortOption } from './types';

// モック商品データ
const MOCK_PRODUCTS: Product[] = [
  {
    id: 1,
    name: "超高性能ノートPC",
    description: "最新のプロセッサ・16GB RAM・高速 SSD 搭載",
    price: 198000,
    salePrice: null,
    sku: "PC-HIGH-001",
    categoryId: 1,
    categoryName: "コンピューター",
    inStock: true,
    stockQuantity: 10,
    isNew: true,
    isFeatured: true,
    imageUrl: null,
    createdAt: "2023-01-15T09:30:00Z",
    updatedAt: "2023-01-15T09:30:00Z"
  },
  {
    id: 2,
    name: "ワイヤレスイヤホン",
    description: "ノイズキャンセリング機能付き高音質イヤホン",
    price: 19800,
    salePrice: 14800,
    sku: "AUDIO-EAR-002",
    categoryId: 3,
    categoryName: "オーディオ",
    inStock: true,
    stockQuantity: 25,
    isNew: false,
    isFeatured: false,
    imageUrl: null,
    createdAt: "2022-11-05T14:20:00Z",
    updatedAt: "2023-02-01T10:15:00Z"
  },
  {
    id: 3,
    name: "高画質タブレット",
    description: "10.2インチディスプレイ、長持ちバッテリー搭載",
    price: 36800,
    salePrice: null,
    sku: "TAB-PRO-003",
    categoryId: 1,
    categoryName: "コンピューター",
    inStock: true,
    stockQuantity: 8,
    isNew: false,
    isFeatured: true,
    imageUrl: null,
    createdAt: "2022-12-10T11:45:00Z",
    updatedAt: "2022-12-10T11:45:00Z"
  },
  {
    id: 4,
    name: "スマートウォッチ",
    description: "健康管理機能付き、長持ちバッテリー",
    price: 32800,
    salePrice: null,
    sku: "WEAR-WATCH-004",
    categoryId: 4,
    categoryName: "ウェアラブル",
    inStock: true,
    stockQuantity: 15,
    isNew: false,
    isFeatured: false,
    imageUrl: null,
    createdAt: "2022-10-20T08:30:00Z",
    updatedAt: "2022-10-20T08:30:00Z"
  },
  {
    id: 5,
    name: "4Kモニター 27インチ",
    description: "鮮明な映像、広色域ディスプレイ",
    price: 49800,
    salePrice: 44800,
    sku: "MON-4K-005",
    categoryId: 1,
    categoryName: "コンピューター",
    inStock: true,
    stockQuantity: 5,
    isNew: true,
    isFeatured: true,
    imageUrl: null,
    createdAt: "2023-01-25T15:10:00Z",
    updatedAt: "2023-02-05T09:20:00Z"
  },
  {
    id: 6,
    name: "ハイエンドスマートフォン",
    description: "最新チップ、高性能カメラ、大容量バッテリー",
    price: 128000,
    salePrice: null,
    sku: "PHONE-PRO-006",
    categoryId: 2,
    categoryName: "スマートフォン",
    inStock: true,
    stockQuantity: 12,
    isNew: true,
    isFeatured: true,
    imageUrl: null,
    createdAt: "2023-02-01T13:45:00Z",
    updatedAt: "2023-02-01T13:45:00Z"
  },
  {
    id: 7,
    name: "Bluetoothスピーカー",
    description: "防水機能付き、高音質ポータブルスピーカー",
    price: 12800,
    salePrice: 9800,
    sku: "AUDIO-SPK-007",
    categoryId: 3,
    categoryName: "オーディオ",
    inStock: true,
    stockQuantity: 20,
    isNew: false,
    isFeatured: false,
    imageUrl: null,
    createdAt: "2022-09-15T10:30:00Z",
    updatedAt: "2023-01-10T08:45:00Z"
  },
  {
    id: 8,
    name: "ワイヤレスマウス",
    description: "高精度センサー、長持ちバッテリー",
    price: 6800,
    salePrice: null,
    sku: "ACC-MOUSE-008",
    categoryId: 5,
    categoryName: "アクセサリー",
    inStock: true,
    stockQuantity: 30,
    isNew: false,
    isFeatured: false,
    imageUrl: null,
    createdAt: "2022-08-05T16:20:00Z",
    updatedAt: "2022-08-05T16:20:00Z"
  },
];

// カテゴリーデータ
const CATEGORIES: Category[] = [
  { 
    id: 0, 
    name: "すべて", 
    slug: "all", 
    productCount: MOCK_PRODUCTS.length 
  },
  { 
    id: 1, 
    name: "コンピューター", 
    slug: "computers", 
    description: "ノートPC、タブレット、デスクトップなど", 
    imageUrl: "/images/categories/computers.jpg",
    productCount: MOCK_PRODUCTS.filter(p => p.categoryId === 1).length 
  },
  { 
    id: 2, 
    name: "スマートフォン", 
    slug: "smartphones", 
    description: "最新のスマートフォン", 
    imageUrl: "/images/categories/smartphones.jpg",
    productCount: MOCK_PRODUCTS.filter(p => p.categoryId === 2).length
  },
  { 
    id: 3, 
    name: "オーディオ", 
    slug: "audio", 
    description: "イヤホン、スピーカー、オーディオ機器", 
    imageUrl: "/images/categories/audio.jpg",
    productCount: MOCK_PRODUCTS.filter(p => p.categoryId === 3).length
  },
  { 
    id: 4, 
    name: "ウェアラブル", 
    slug: "wearables", 
    description: "スマートウォッチ、フィットネストラッカーなど", 
    imageUrl: "/images/categories/wearables.jpg",
    productCount: MOCK_PRODUCTS.filter(p => p.categoryId === 4).length 
  },
  { 
    id: 5, 
    name: "アクセサリー", 
    slug: "accessories", 
    description: "各種デバイスのアクセサリー", 
    imageUrl: "/images/categories/accessories.jpg",
    productCount: MOCK_PRODUCTS.filter(p => p.categoryId === 5).length 
  },
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
  // フィーチャー商品をフィルタリング
  const featured = MOCK_PRODUCTS.filter(p => p.isFeatured);
  return Promise.resolve(featured.slice(0, limit));
}

// カテゴリー一覧の取得
export async function fetchCategories(): Promise<Category[]> {
  return Promise.resolve([...CATEGORIES]);
}

// 商品をIDで取得
export async function fetchProductById(id: number): Promise<Product | null> {
  const product = MOCK_PRODUCTS.find(p => p.id === id);
  return Promise.resolve(product || null);
}

// 商品のフィルタリングと並び替え
export function filterAndSortProducts(
  products: Product[],
  category: number = 0,
  sortOption: SortOption = 'recommended'
): Product[] {
  // カテゴリーでフィルタリング
  let filtered = products;
  
  // カテゴリーIDを使用してフィルタリング
  filtered = category === 0 
    ? products 
    : products.filter(product => product.categoryId === category);
  
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
      return [...products].sort((a, b) => (a.isFeatured && !b.isFeatured) ? -1 : (!a.isFeatured && b.isFeatured) ? 1 : 0);
  }
}
