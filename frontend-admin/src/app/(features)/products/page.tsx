// Container Component (Server Component)
import { fetchCategories, fetchProducts } from '@/services/products/api';

import ProductsClient from './client';

type ProductsPageProps = {
  searchParams: { [key: string]: string | string[] | undefined };
};

export default async function ProductsPage({ searchParams }: ProductsPageProps) {
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
}
