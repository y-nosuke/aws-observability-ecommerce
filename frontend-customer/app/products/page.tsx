import { Suspense } from 'react';
import CategoryFilter from './components/CategoryFilter';
import LoadingIndicator from './components/LoadingIndicator';
import ProductsList from './ProductsList';

export default async function ProductsPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; category?: string }>;
}) {
  // URLパラメータの処理
  const { page, category } = await searchParams;
  const pageParam = Number(page) || 1;
  const categoryParam = category ? Number(category) : undefined;

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold text-gray-800 mb-6">商品一覧</h1>

      <div className="flex flex-col md:flex-row gap-6">
        <div className="w-full md:w-64 flex-shrink-0">
          <Suspense fallback={<div className="h-24 bg-gray-200 animate-pulse rounded"></div>}>
            <CategoryFilter selectedCategoryId={categoryParam} />
          </Suspense>
        </div>

        <div className="flex-grow">
          <Suspense fallback={<LoadingIndicator />}>
            <ProductsList page={pageParam} pageSize={9} categoryId={categoryParam} />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
