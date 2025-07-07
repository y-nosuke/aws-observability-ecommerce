'use client';

import DashboardCard from '@/components/ui/DashboardCard';
import DashboardStats from '@/components/ui/DashboardStats';
import SystemHealthStatus from '@/components/ui/SystemHealthStatus';

import InventoryAlertsTable from './components/InventoryAlertsTable';
import LatestOrdersTable from './components/LatestOrdersTable';
import PopularProductsTable from './components/PopularProductsTable';
import { InventoryAlertItem, OrderItem, ProductItem, StatItem } from './types';

// Presentational コンポーネントの props 型定義
type HomeClientProps = {
  stats: StatItem[];
  latestOrders: OrderItem[];
  popularProducts: ProductItem[];
  inventoryAlerts: InventoryAlertItem[];
};

export default function HomeClient({
  stats,
  latestOrders,
  popularProducts,
  inventoryAlerts,
}: HomeClientProps) {
  return (
    <div className="space-y-6">
      <h1 className="mb-6 text-2xl font-bold">管理ダッシュボード</h1>

      {/* 統計カード */}
      <DashboardStats stats={stats} />

      <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
        {/* 最新の注文 */}
        <DashboardCard
          title="最新の注文"
          icon={
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="h-5 w-5"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M9 12h3.75M9 15h3.75M9 18h3.75m3 .75H18a2.25 2.25 0 002.25-2.25V6.108c0-1.135-.845-2.098-1.976-2.192a48.424 48.424 0 00-1.123-.08m-5.801 0c-.065.21-.1.433-.1.664 0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75 2.25 2.25 0 00-.1-.664m-5.8 0A2.251 2.251 0 0113.5 2.25H15c1.012 0 1.867.668 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m0 0H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V9.375c0-.621-.504-1.125-1.125-1.125H8.25zM6.75 12h.008v.008H6.75V12zm0 3h.008v.008H6.75V15zm0 3h.008v.008H6.75V18z"
              />
            </svg>
          }
        >
          <LatestOrdersTable orders={latestOrders} />
        </DashboardCard>

        {/* 人気商品 */}
        <DashboardCard
          title="人気商品"
          icon={
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="h-5 w-5"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z"
              />
            </svg>
          }
        >
          <PopularProductsTable products={popularProducts} />
        </DashboardCard>
      </div>

      {/* 在庫アラート */}
      <DashboardCard
        title="在庫アラート"
        icon={
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className="h-5 w-5"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z"
            />
          </svg>
        }
      >
        <InventoryAlertsTable alerts={inventoryAlerts} />
      </DashboardCard>

      {/* システムヘルスステータス */}
      <SystemHealthStatus mode="detailed" autoRefreshMinutes={5} />
    </div>
  );
}
