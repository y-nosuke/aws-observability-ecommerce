import { DashboardData, OrderItem, ProductItem, InventoryAlertItem, StatItem } from './types';

// モックのダッシュボード統計データ
const MOCK_STATS: StatItem[] = [
  {
    label: "今日の売上",
    value: "¥128,430",
    change: "12.3%",
    isPositive: true,
  },
  {
    label: "新規注文",
    value: "24",
    change: "5.4%",
    isPositive: true,
  },
  {
    label: "新規顧客",
    value: "12",
    change: "2.1%",
    isPositive: true,
  },
  {
    label: "在庫切れ商品",
    value: "3",
    change: "50%",
    isPositive: false,
  },
];

// モックの最新注文データ
const MOCK_LATEST_ORDERS: OrderItem[] = [
  {
    id: "#1234",
    customer: "佐藤太郎",
    amount: "¥12,800",
    status: "完了",
    statusColor: "green",
  },
  {
    id: "#1233",
    customer: "田中花子",
    amount: "¥8,500",
    status: "発送中",
    statusColor: "yellow",
  },
  {
    id: "#1232",
    customer: "山田次郎",
    amount: "¥23,400",
    status: "処理中",
    statusColor: "blue",
  },
];

// モックの人気商品データ
const MOCK_POPULAR_PRODUCTS: ProductItem[] = [
  {
    name: "高性能ノートパソコン",
    salesCount: 24,
    stockCount: 18,
    revenue: "¥1,920,000",
  },
  {
    name: "ワイヤレスイヤホン",
    salesCount: 42,
    stockCount: 35,
    revenue: "¥756,000",
  },
  {
    name: "スマートウォッチ",
    salesCount: 38,
    stockCount: 5,
    revenue: "¥684,000",
  },
];

// モックの在庫アラートデータ
const MOCK_INVENTORY_ALERTS: InventoryAlertItem[] = [
  {
    name: "スマートウォッチ",
    currentStock: 5,
    minStock: 10,
    status: "要発注",
    statusColor: "red",
  },
  {
    name: "4Kモニター",
    currentStock: 8,
    minStock: 15,
    status: "要発注",
    statusColor: "red",
  },
  {
    name: "ゲーミングキーボード",
    currentStock: 0,
    minStock: 5,
    status: "在庫切れ",
    statusColor: "red",
  },
];

/**
 * ダッシュボードデータ全体を取得
 * @returns ダッシュボードデータ
 */
export async function fetchDashboardData(): Promise<DashboardData> {
  // 本番環境では実際のAPIエンドポイントからデータを取得する
  // const response = await fetch('api/dashboard/stats');
  // return response.json();
  
  // モックデータを返す
  return Promise.resolve({
    stats: MOCK_STATS,
    latestOrders: MOCK_LATEST_ORDERS,
    popularProducts: MOCK_POPULAR_PRODUCTS,
    inventoryAlerts: MOCK_INVENTORY_ALERTS,
  });
}

/**
 * ダッシュボードの統計データを取得
 * @returns 統計データ
 */
export async function fetchDashboardStats(): Promise<StatItem[]> {
  return Promise.resolve([...MOCK_STATS]);
}

/**
 * 最新の注文データを取得
 * @param limit 取得する注文数
 * @returns 注文データ
 */
export async function fetchLatestOrders(limit: number = 3): Promise<OrderItem[]> {
  return Promise.resolve([...MOCK_LATEST_ORDERS].slice(0, limit));
}

/**
 * 人気商品データを取得
 * @param limit 取得する商品数
 * @returns 商品データ
 */
export async function fetchPopularProducts(limit: number = 3): Promise<ProductItem[]> {
  return Promise.resolve([...MOCK_POPULAR_PRODUCTS].slice(0, limit));
}

/**
 * 在庫アラートデータを取得
 * @param limit 取得するアラート数
 * @returns 在庫アラートデータ
 */
export async function fetchInventoryAlerts(limit: number = 3): Promise<InventoryAlertItem[]> {
  return Promise.resolve([...MOCK_INVENTORY_ALERTS].slice(0, limit));
}
