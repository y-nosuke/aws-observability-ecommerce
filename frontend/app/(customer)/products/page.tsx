import ProductList from "@/components/product/ProductList";

export const metadata = {
  title: "商品一覧 | ECアプリ",
  description: "商品一覧ページです。",
};

export default function ProductsPage() {
  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">商品一覧</h1>
      <ProductList />
    </div>
  );
}
