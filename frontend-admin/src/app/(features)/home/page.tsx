// Container Component (Server Component)
import { fetchDashboardData } from '@/services/dashboard/api';

import HomeClient from './client';

export default async function HomePage() {
  // ダッシュボードデータをサービスから取得
  const dashboardData = await fetchDashboardData();

  // Presentational コンポーネントにデータを渡す
  return (
    <HomeClient
      stats={dashboardData.stats}
      latestOrders={dashboardData.latestOrders}
      popularProducts={dashboardData.popularProducts}
      inventoryAlerts={dashboardData.inventoryAlerts}
    />
  );
}
