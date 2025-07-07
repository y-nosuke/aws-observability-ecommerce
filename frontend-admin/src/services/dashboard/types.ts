// ダッシュボード関連の型定義

// 統計データの型定義
export interface StatItem {
  label: string;
  value: string;
  change: string;
  isPositive: boolean;
}

// 注文データの型定義
export interface OrderItem {
  id: string;
  customer: string;
  amount: string;
  status: string;
  statusColor: string;
}

// 商品データの型定義
export interface ProductItem {
  name: string;
  salesCount: number;
  stockCount: number;
  revenue: string;
}

// 在庫アラートの型定義
export interface InventoryAlertItem {
  name: string;
  currentStock: number;
  minStock: number;
  status: string;
  statusColor: string;
}

// ダッシュボードデータ型定義
export interface DashboardData {
  stats: StatItem[];
  latestOrders: OrderItem[];
  popularProducts: ProductItem[];
  inventoryAlerts: InventoryAlertItem[];
}
