import { getProducts } from '@/lib/api/products';
import ErrorDisplay from './components/ErrorDisplay';
import Pagination from './components/Pagination';
import ProductCard from './components/ProductCard';

export default async function ProductsList({
  page,
  pageSize,
  categoryId,
}: {
  page: number;
  pageSize: number;
  categoryId?: number;
}) {
  try {
    const { items: products, total_pages: totalPages } = await getProducts({
      page,
      page_size: pageSize,
      category_id: categoryId,
    });

    if (products.length === 0) {
      return <div className="text-center py-12 text-gray-500">商品が見つかりませんでした</div>;
    }

    return (
      <>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {products.map(product => (
            <ProductCard key={product.id} product={product} />
          ))}
        </div>

        <Pagination currentPage={page} totalPages={totalPages} />
      </>
    );
  } catch (error) {
    console.error('Failed to load products:', error);

    return (
      <ErrorDisplay
        message={error instanceof Error ? error.message : '不明なエラーが発生しました'}
      />
    );
  }
}
