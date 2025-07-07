import { fetchCategories, fetchProducts } from '@/services/products/api';

import ProductsClient from './client';

// 動的レンダリングを強制（ビルド時の静的生成を無効化）
export const dynamic = 'force-dynamic';

// サーバーコンポーネント：データを取得し、プレゼンテーショナルコンポーネントに渡す
export default async function ProductsPage() {
  try {
    // サーバーサイドでデータを取得
    const [products, categories] = await Promise.all([fetchProducts(), fetchCategories()]);

    // クライアントコンポーネントにデータを渡す
    return <ProductsClient initialProducts={products} categories={categories} />;
  } catch (error) {
    // エラーが発生した場合はNext.jsのエラーバウンダリに委ねる
    throw error;
  }
}
