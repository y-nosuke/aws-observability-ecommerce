"use client";

import { InventoryAlertItem } from "../types";

type InventoryAlertsTableProps = {
  alerts: InventoryAlertItem[];
};

export default function InventoryAlertsTable({
  alerts,
}: InventoryAlertsTableProps) {
  return (
    <div className="space-y-4">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
          <thead>
            <tr>
              <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                商品名
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                現在の在庫
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                最小在庫数
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                状態
              </th>
            </tr>
          </thead>
          <tbody className="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
            {alerts.map((alert) => (
              <tr key={alert.name}>
                <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-300">
                  {alert.name}
                </td>
                <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-300">
                  {alert.currentStock}
                </td>
                <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-300">
                  {alert.minStock}
                </td>
                <td className="px-4 py-3 whitespace-nowrap">
                  <span
                    className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-${alert.statusColor}-100 text-${alert.statusColor}-800`}
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
          className="text-indigo-600 hover:text-indigo-800 dark:text-indigo-400 dark:hover:text-indigo-300 text-sm font-medium"
        >
          在庫管理へ →
        </a>
      </div>
    </div>
  );
}
