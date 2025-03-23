import Link from 'next/link';
import { Suspense } from 'react';
import ErrorDisplay from './components/ErrorDisplay';
import LoadingIndicator from './components/LoadingIndicator';
import ProductTable from './components/ProductTable';
import SearchForm from './components/SearchForm';

export default async function AdminProductsPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; category?: string }>;
}) {
  // URLパラメータの処理
  const { page, category } = await searchParams;
  const pageParam = Number(page) || 1;
  const categoryParam = category ? Number(category) : undefined;

  try {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold text-gray-800">商品管理</h1>
          <Link
            href="/products/create"
            className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-300"
          >
            新規商品登録
          </Link>
        </div>

        <Suspense fallback={<div className="animate-pulse h-24 bg-gray-200 rounded-lg mb-6"></div>}>
          <SearchForm />
        </Suspense>

        <Suspense fallback={<LoadingIndicator />}>
          <ProductTable pageParam={pageParam} categoryParam={categoryParam} />
        </Suspense>
      </div>
    );
  } catch (error) {
    console.error('Failed to load products:', error);

    return (
      <div className="container mx-auto px-4 py-8">
        <h1 className="text-3xl font-bold text-gray-800 mb-6">商品管理</h1>
        <ErrorDisplay error={error} />
      </div>
    );
  }
}
