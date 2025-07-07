'use client';

import { ProductItem } from '../types';

type PopularProductsTableProps = {
  products: ProductItem[];
};

export default function PopularProductsTable({ products }: PopularProductsTableProps) {
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
                売上数
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                在庫数
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium tracking-wider text-gray-500 uppercase dark:text-gray-400">
                売上高
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-800">
            {products.map((product) => (
              <tr key={product.name}>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {product.name}
                </td>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {product.salesCount}
                </td>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {product.stockCount}
                </td>
                <td className="px-4 py-3 text-sm whitespace-nowrap text-gray-900 dark:text-gray-300">
                  {product.revenue}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <div className="text-right">
        <a
          href="/products"
          className="text-sm font-medium text-indigo-600 hover:text-indigo-800 dark:text-indigo-400 dark:hover:text-indigo-300"
        >
          すべての商品を表示 →
        </a>
      </div>
    </div>
  );
}
