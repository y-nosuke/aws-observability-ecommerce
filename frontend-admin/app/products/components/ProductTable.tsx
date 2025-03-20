import { getCategories, getProducts } from '@/lib/api/products';
import Link from 'next/link';
import DeleteButton from './DeleteButton';
import ProductImage from './ProductImage';
import ProductTablePagination from './ProductTablePagination';

export default async function ProductTable({
  pageParam,
  categoryParam,
}: {
  pageParam: number;
  categoryParam?: number;
}) {
  // 商品データの取得
  const { items: products, total_pages: totalPages } = await getProducts({
    page: pageParam,
    page_size: 10,
    category_id: categoryParam,
  });

  // カテゴリーデータの取得
  const categories = await getCategories();

  return (
    <>
      {/* 商品一覧テーブル */}
      <div className="bg-white rounded-lg shadow overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                ID
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                商品名
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                価格
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                カテゴリー
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                在庫状況
              </th>
              <th className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                操作
              </th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {products.map(product => (
              <tr key={product.id} className="hover:bg-gray-50">
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{product.id}</td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="flex items-center">
                    <div className="h-10 w-10 flex-shrink-0 relative">
                      <ProductImage
                        imageUrl={product.image_url}
                        productName={product.name}
                        productId={product.id}
                      />
                    </div>
                    <div className="ml-4">
                      <div className="text-sm font-medium text-gray-900">{product.name}</div>
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  ¥{product.price.toLocaleString()}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                    {categories.find(c => c.id === product.category_id)?.name || '不明'}
                  </span>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-blue-100 text-blue-800">
                    在庫あり
                  </span>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                  <Link
                    href={`/products/${product.id}`}
                    className="text-emerald-600 hover:text-emerald-900 mr-4"
                  >
                    詳細
                  </Link>
                  <Link
                    href={`/products/${product.id}/edit`}
                    className="text-indigo-600 hover:text-indigo-900 mr-4"
                  >
                    編集
                  </Link>
                  <DeleteButton productId={product.id} />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* ページネーション */}
      <ProductTablePagination
        pageParam={pageParam}
        totalPages={totalPages}
        totalProducts={products.length}
      />
    </>
  );
}
