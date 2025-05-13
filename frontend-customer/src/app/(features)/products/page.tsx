import { fetchProducts, fetchCategories } from "@/services/products/api";
import ProductsClient from "./client";

// サーバーコンポーネント：データを取得し、プレゼンテーショナルコンポーネントに渡す
export default async function ProductsPage() {
  // サーバーサイドでデータを取得
  const [products, categories] = await Promise.all([
    fetchProducts(),
    fetchCategories(),
  ]);

  // クライアントコンポーネントにデータを渡す
  return <ProductsClient initialProducts={products} categories={categories} />;
}
