"use client";

import { useProducts } from "@/lib/hooks/useProducts";
import Link from "next/link";

export default function ProductList() {
  const { products, isLoading, isError } = useProducts();

  if (isLoading) {
    return <div className="text-center py-10">読み込み中...</div>;
  }

  if (isError) {
    return (
      <div className="text-center py-10 text-red-500">
        商品データの取得に失敗しました
      </div>
    );
  }

  if (!products || products.length === 0) {
    return <div className="text-center py-10">商品がありません</div>;
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {products.map((product) => (
        <div
          key={product.id}
          className="border rounded-lg overflow-hidden shadow-lg"
        >
          <div className="h-48 bg-gray-200">
            {/* 商品画像（実際の実装では画像を表示） */}
            <div className="h-full flex items-center justify-center text-gray-500">
              商品画像
            </div>
          </div>
          <div className="p-4">
            <h2 className="text-xl font-semibold mb-2">{product.name}</h2>
            <p className="text-gray-600 mb-2 line-clamp-2">
              {product.description}
            </p>
            <div className="flex justify-between items-center mt-4">
              <span className="text-lg font-bold">
                {product.price.toLocaleString()}円
              </span>
              <Link
                href={`/products/${product.id}`}
                className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
              >
                詳細を見る
              </Link>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
