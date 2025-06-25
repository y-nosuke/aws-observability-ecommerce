"use client";

import { fetchProducts } from "@/services/products/api";
import { ProductListResponse } from "@/services/products/types";
import Image from "next/image";
import { useSearchParams } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

type ProductTableProps = {
  initialProductsData: ProductListResponse;
};

export default function ProductTable({
  initialProductsData,
}: ProductTableProps) {
  const searchParams = useSearchParams();
  const [productsData, setProductsData] =
    useState<ProductListResponse>(initialProductsData);
  const [isLoading, setIsLoading] = useState(false);

  // URLから検索パラメータを構築
  const buildSearchParams = useCallback(() => {
    const keyword = searchParams.get("keyword");
    const categoryIdParam = searchParams.get("categoryId");
    const page = parseInt(searchParams.get("page") || "1");
    const limit = parseInt(searchParams.get("limit") || "10");

    // categoryIdは数値として扱う必要がある
    const categoryId = categoryIdParam ? parseInt(categoryIdParam) : undefined;

    return {
      keyword: keyword || undefined,
      categoryId,
      page,
      limit,
    };
  }, [searchParams]);

  // 商品データを取得する関数
  const loadProducts = useCallback(async () => {
    setIsLoading(true);
    try {
      const params = buildSearchParams();
      const data = await fetchProducts(params);
      setProductsData(data);
    } catch (error) {
      console.error("商品データの取得に失敗しました:", error);
      // エラーハンドリング - 実際のアプリケーションではトーストやエラー表示を実装
    } finally {
      setIsLoading(false);
    }
  }, [buildSearchParams]);

  // 検索パラメータが変更された時にデータを再取得
  useEffect(() => {
    // 初回表示時は初期データを使用し、2回目以降の検索パラメータ変更時のみAPI呼び出し
    const currentParams = buildSearchParams();
    const isInitialLoad =
      currentParams.page === 1 &&
      !currentParams.keyword &&
      !currentParams.categoryId;

    if (!isInitialLoad) {
      loadProducts();
    }
  }, [searchParams, loadProducts, buildSearchParams]);

  // 価格をフォーマットする関数
  const formatPrice = (price: number) => {
    return new Intl.NumberFormat("ja-JP", {
      style: "currency",
      currency: "JPY",
      minimumFractionDigits: 0,
    }).format(price);
  };

  // 在庫状況に応じたスタイルを返す関数
  const getStockStyle = (stock: number) => {
    if (stock === 0) {
      return "bg-red-100 text-red-800";
    } else if (stock <= 10) {
      return "bg-yellow-100 text-yellow-800";
    }
    return "bg-green-100 text-green-800";
  };

  // 在庫状況のテキストを返す関数
  const getStockText = (stock: number) => {
    if (stock === 0) {
      return "在庫切れ";
    } else if (stock <= 10) {
      return "在庫少";
    }
    return "在庫有";
  };

  // ローディング中の表示
  if (isLoading) {
    return (
      <div className="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 overflow-hidden shadow-sm">
        <div className="px-6 py-4 border-b border-gray-200 dark:border-gray-700">
          <h3 className="text-lg font-medium text-gray-900 dark:text-gray-100">
            商品一覧
          </h3>
        </div>
        <div className="p-6">
          <div className="flex items-center justify-center h-64">
            <div className="flex items-center gap-3">
              <svg
                className="animate-spin h-8 w-8 text-indigo-600"
                viewBox="0 0 24 24"
              >
                <circle
                  className="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  strokeWidth="4"
                  fill="none"
                />
                <path
                  className="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                />
              </svg>
              <span className="text-lg text-gray-600">
                商品データを読み込み中...
              </span>
            </div>
          </div>
        </div>
      </div>
    );
  }

  // 商品がない場合の表示
  if (
    !isLoading &&
    (!productsData.products || productsData.products.length === 0)
  ) {
    return (
      <div className="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 overflow-hidden shadow-sm">
        <div className="px-6 py-4 border-b border-gray-200 dark:border-gray-700">
          <h3 className="text-lg font-medium text-gray-900 dark:text-gray-100">
            商品一覧
          </h3>
        </div>
        <div className="p-6">
          <div className="flex flex-col items-center justify-center h-64 text-center">
            <svg
              className="h-16 w-16 text-gray-400 mb-4"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1}
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z"
              />
            </svg>
            <h3 className="text-lg font-medium text-gray-900 dark:text-gray-100 mb-2">
              商品がありません
            </h3>
            <p className="text-gray-600 dark:text-gray-400">
              検索条件を変更するか、新しい商品を登録してください。
            </p>
          </div>
        </div>
      </div>
    );
  }

  // 商品テーブルの表示
  return (
    <div className="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 overflow-hidden shadow-sm">
      <div className="px-6 py-4 border-b border-gray-200 dark:border-gray-700">
        <h3 className="text-lg font-medium text-gray-900 dark:text-gray-100">
          商品一覧
        </h3>
      </div>
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
          <thead className="bg-gray-50 dark:bg-gray-700">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                商品
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                価格
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                在庫数
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                カテゴリ
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                状態
              </th>
            </tr>
          </thead>
          <tbody className="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
            {(productsData.products || []).map((product) => (
              <tr
                key={product.id}
                className="hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
              >
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="flex items-center">
                    <div className="flex-shrink-0 h-12 w-12">
                      {product.imageUrl ? (
                        <Image
                          src={product.imageUrl}
                          alt={product.name}
                          width={48}
                          height={48}
                          className="h-12 w-12 rounded object-cover bg-gray-100"
                          onError={(e) => {
                            // 画像の読み込みに失敗した場合はプレースホルダーを表示
                            e.currentTarget.style.display = "none";
                          }}
                        />
                      ) : (
                        <div className="h-12 w-12 rounded bg-gray-100 flex items-center justify-center">
                          <svg
                            className="h-6 w-6 text-gray-400"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 24 24"
                            strokeWidth={1.5}
                            stroke="currentColor"
                          >
                            <path
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z"
                            />
                          </svg>
                        </div>
                      )}
                    </div>
                    <div className="ml-4">
                      <div className="text-sm font-medium text-gray-900 dark:text-gray-100">
                        {product.name}
                      </div>
                      {product.description && (
                        <div className="text-sm text-gray-500 dark:text-gray-400 truncate max-w-xs">
                          {product.description}
                        </div>
                      )}
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
                  {formatPrice(product.price)}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
                  <span className="font-medium">{product.stockQuantity}</span>
                  <span className="text-gray-500 dark:text-gray-400 ml-1">
                    個
                  </span>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
                  {product.categoryName}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span
                    className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${getStockStyle(
                      product.stockQuantity
                    )}`}
                  >
                    {getStockText(product.stockQuantity)}
                  </span>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
