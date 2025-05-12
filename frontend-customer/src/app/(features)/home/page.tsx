import { fetchPopularProducts } from "@/services/products/api";
import HomeClient from "./client";

export default async function HomePage() {
  // 人気商品のデータを取得
  const popularProducts = await fetchPopularProducts(4);
  
  // クライアントコンポーネントにデータを渡す
  return <HomeClient popularProducts={popularProducts} />;
}
