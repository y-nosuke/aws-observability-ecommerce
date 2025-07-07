'use client';

import { InventoryAlertItem } from '../types';

type InventoryAlertsTableProps = {
  alerts: InventoryAlertItem[];
};

export default function InventoryAlertsTable({ alerts }: InventoryAlertsTableProps) {
  return (
    <div className="space-y-4">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
          <thead>
            <tr>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                商品名
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                現在の在庫
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                最小在庫数
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                状態
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-800">
            {alerts.map((alert) => (
              <tr key={alert.name}>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {alert.name}
                </td>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {alert.currentStock}
                </td>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {alert.minStock}
                </td>
                <td className="px-4 py-3 whitespace-nowrap">
                  <span
                    className={`inline-flex rounded-full px-2 text-xs leading-5 font-semibold bg-${alert.statusColor}-100 text-${alert.statusColor}-800`}
                  >
                    {alert.status}
                  </span>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <div className="text-right">
        <a
          href="/inventory"
          className="text-sm font-medium text-indigo-600 hover:text-indigo-800 dark:text-indigo-400 dark:hover:text-indigo-300"
        >
          在庫管理へ →
        </a>
      </div>
    </div>
  );
}
