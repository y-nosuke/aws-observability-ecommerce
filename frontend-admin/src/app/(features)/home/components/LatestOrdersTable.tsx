'use client';

import { OrderItem } from '../types';

type LatestOrdersTableProps = {
  orders: OrderItem[];
};

export default function LatestOrdersTable({ orders }: LatestOrdersTableProps) {
  return (
    <div className="space-y-4">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
          <thead>
            <tr>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                注文番号
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                顧客
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                金額
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                状態
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-800">
            {orders.map((order) => (
              <tr key={order.id}>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {order.id}
                </td>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {order.customer}
                </td>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {order.amount}
                </td>
                <td className="px-4 py-3 whitespace-nowrap">
                  <span
                    className={`inline-flex rounded-full px-2 text-xs leading-5 font-semibold bg-${order.statusColor}-100 text-${order.statusColor}-800`}
                  >
                    {order.status}
                  </span>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <div className="text-right">
        <a
          href="/orders"
          className="text-sm font-medium text-indigo-600 hover:text-indigo-800 dark:text-indigo-400 dark:hover:text-indigo-300"
        >
          すべての注文を表示 →
        </a>
      </div>
    </div>
  );
}
