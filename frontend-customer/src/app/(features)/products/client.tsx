'use client';

import { useEffect, useState } from 'react';

import { productsApi } from '@/services/products/api';
import { Category, Product, SortOption } from '@/services/products/types';

import FilterSidebar from './components/FilterSidebar';
import ProductGrid from './components/ProductGrid';
import SortSelector from './components/SortSelector';

// プレゼンテーショナルコンポーネントに渡すプロパティの型定義
interface ProductsClientProps {
  initialProducts: Product[];
  categories: Category[];
}

export default function ProductsClient({ initialProducts, categories }: ProductsClientProps) {
  // クライアントサイドの状態管理
  const [activeCategory, setActiveCategory] = useState<number>(0);
  const [sortOption, setSortOption] = useState<SortOption>('recommended');
  const [selectedPriceRanges, setSelectedPriceRanges] = useState<string[]>([]);
  const [selectedStatuses, setSelectedStatuses] = useState<string[]>([]);
  const [products, setProducts] = useState<Product[]>(initialProducts);
  const [filteredCategories, setFilteredCategories] = useState<Category[]>(categories);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  // フィルター条件が変更されたときに実行されるエフェクト
  useEffect(() => {
    const fetchProductsByCategory = async () => {
      setIsLoading(true);
      setError(null);
      try {
        if (activeCategory !== 0) {
          const response = await productsApi.getProductsByCategory(activeCategory);
          setProducts(response.items);
        } else {
          setProducts(initialProducts);
        }
        setFilteredCategories(categories);
      } catch (error) {
        console.error('Failed to fetch products by category:', error);
        setError('商品データの取得に失敗しました。しばらく時間をおいてから再度お試しください。');
      } finally {
        setIsLoading(false);
      }
    };

    fetchProductsByCategory();
  }, [activeCategory, categories, initialProducts]);

  // シンプルにした各ハンドラー
  const handleCategoryChange = (categoryId: number) => {
    setActiveCategory(categoryId);
  };

  const handleSortChange = (option: SortOption) => {
    // 並び替えはサーバーサイドで行うため、ここでは実装しない
    setSortOption(option);
  };

  const handlePriceFilterChange = (priceRanges: string[]) => {
    // 価格フィルターはサーバーサイドで行うため、ここでは実装しない
    setSelectedPriceRanges(priceRanges);
  };

  const handleStatusFilterChange = (statuses: string[]) => {
    // 状態フィルターはサーバーサイドで行うため、ここでは実装しない
    setSelectedStatuses(statuses);
  };

  return (
    <div className="container mx-auto px-6 py-8">
      <div className="flex flex-col gap-8 md:flex-row">
        {/* フィルターサイドバー */}
        <FilterSidebar
          categories={filteredCategories}
          activeCategory={activeCategory}
          onCategoryChange={handleCategoryChange}
          selectedPriceRanges={selectedPriceRanges}
          onPriceFilterChange={handlePriceFilterChange}
          selectedStatuses={selectedStatuses}
          onStatusFilterChange={handleStatusFilterChange}
        />

        {/* 商品一覧 */}
        <div className="flex-1">
          <div className="mb-6 flex flex-wrap items-center justify-between gap-4">
            <h1 className="text-2xl font-bold md:text-3xl">商品一覧</h1>
            <SortSelector sortOption={sortOption} onSortChange={handleSortChange} />
          </div>

          {isLoading ? (
            <div className="flex min-h-[200px] items-center justify-center">
              <div className="h-12 w-12 animate-spin rounded-full border-b-2 border-blue-600"></div>
            </div>
          ) : error ? (
            <div className="rounded-lg bg-red-50 p-6 text-center dark:bg-red-900/20">
              <p className="text-red-600 dark:text-red-400">{error}</p>
              <button
                onClick={() => {
                  setError(null);
                  handleCategoryChange(activeCategory);
                }}
                className="mt-4 rounded-md bg-red-600 px-4 py-2 text-white transition-colors hover:bg-red-700"
              >
                再試行
              </button>
            </div>
          ) : (
            <>
              <ProductGrid products={products} />

              {/* 商品がない場合のメッセージ */}
              {products.length === 0 && (
                <div className="rounded-lg bg-gray-100 p-6 text-center dark:bg-gray-800">
                  <p className="text-gray-600 dark:text-gray-300">
                    条件に一致する商品がありません。フィルターを変更してください。
                  </p>
                </div>
              )}
            </>
          )}
        </div>
      </div>
    </div>
  );
}
