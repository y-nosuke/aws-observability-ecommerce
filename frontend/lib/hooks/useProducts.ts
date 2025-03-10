import apiClient from "@/lib/api/client";
import useSWR from "swr";

// 商品データの型定義
export interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  imageUrl: string;
  categoryId: number;
}

// 商品一覧を取得するフェッチャー関数
const productsFetcher = async () => {
  const response = await apiClient.get("/api/products");
  return response.data;
};

// 商品詳細を取得するフェッチャー関数
const productDetailsFetcher = async (id: string) => {
  const response = await apiClient.get(`/api/products/${id}`);
  return response.data;
};

// 商品一覧を取得するカスタムフック
export function useProducts() {
  const { data, error, isLoading, mutate } = useSWR<Product[]>(
    "/api/products",
    productsFetcher
  );

  return {
    products: data,
    isLoading,
    isError: error,
    mutate,
  };
}

// 商品詳細を取得するカスタムフック
export function useProductDetails(id: string) {
  const { data, error, isLoading } = useSWR<Product>(
    id ? `/api/products/${id}` : null,
    () => productDetailsFetcher(id)
  );

  return {
    product: data,
    isLoading,
    isError: error,
  };
}
