// Container Component (Server Component)
import { fetchCategories, fetchProducts } from '@/services/products/api';

import ProductsClient from './client';

type ProductsPageProps = {
  searchParams: { [key: string]: string | string[] | undefined };
};

export default async function ProductsPage({ searchParams }: ProductsPageProps) {
  try {
    // Next.js 15では searchParams を await する必要がある
    const params = await searchParams;

    // URL検索パラメータを解析
    const keyword = typeof params.keyword === 'string' ? params.keyword : undefined;
    const categoryIdParam = typeof params.categoryId === 'string' ? params.categoryId : undefined;
    const page = typeof params.page === 'string' ? parseInt(params.page) : 1;
    const limit = typeof params.limit === 'string' ? parseInt(params.limit) : 10;

    // categoryIdは数値として処理
    const categoryId = categoryIdParam ? parseInt(categoryIdParam) : undefined;

    // 初期データを並行して取得
    const [initialProductsData, categories] = await Promise.all([
      fetchProducts({
        page,
        limit,
        keyword,
        categoryId,
      }),
      fetchCategories(),
    ]);

    // Presentational コンポーネントにデータを渡す
    return <ProductsClient initialProductsData={initialProductsData} categories={categories} />;
  } catch (error) {
    console.error('商品ページの初期データ取得に失敗しました:', error);

    // エラー状態を表示
    return (
      <div className="flex min-h-[400px] items-center justify-center">
        <div className="text-center">
          <h2 className="mb-2 text-lg font-semibold text-red-600">
            データの読み込みに失敗しました
          </h2>
          <p className="text-gray-600">しばらく時間をおいてから再度お試しください。</p>
        </div>
      </div>
    );
  }
}
