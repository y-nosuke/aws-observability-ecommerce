"use client";

import { Category, ProductListResponse } from "@/services/products/types";
import Pagination from "./components/Pagination";
import ProductTable from "./components/ProductTable";
import ProductToolbar from "./components/ProductToolbar";

// Presentational コンポーネントの props 型定義
type ProductsClientProps = {
  initialProductsData: ProductListResponse;
  categories: Category[];
};

export default function ProductsClient({
  initialProductsData,
  categories,
}: ProductsClientProps) {
  return (
    <div className="space-y-6">
      {/* ページヘッダー */}
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-900 dark:text-gray-100">
          商品管理
        </h1>
        <div className="text-sm text-gray-600 dark:text-gray-400">
          {initialProductsData.totalCount || 0}件の商品
        </div>
      </div>

      {/* 検索・フィルターツールバー */}
      <ProductToolbar categories={categories} />

      {/* 商品テーブル */}
      <ProductTable initialProductsData={initialProductsData} />

      {/* ページネーション */}
      <Pagination
        currentPage={initialProductsData.currentPage || 1}
        totalPages={initialProductsData.totalPages || 1}
        hasNextPage={initialProductsData.hasNextPage || false}
        hasPreviousPage={initialProductsData.hasPreviousPage || false}
      />
    </div>
  );
}
