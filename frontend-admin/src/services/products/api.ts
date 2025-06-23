import { apiClient } from "../api-client";
import {
  Category,
  Product,
  ProductListResponse,
  ProductSearchParams,
} from "./types";

/**
 * 商品一覧を取得
 * @param params 検索パラメータ
 * @returns 商品一覧レスポンス
 */
export async function fetchProducts(
  params: ProductSearchParams = {}
): Promise<ProductListResponse> {
  try {
    const response = await apiClient.get("/products", { params });
    const data = response.data;

    // バックエンドのレスポンスをフロントエンドの Product 型に変換
    const products: Product[] = data.items || [];

    // ページネーション情報を適切に変換
    const totalPages =
      data.total_pages || Math.ceil((data.total || 0) / (params.limit || 10));
    const currentPage = data.page || 1;

    // バックエンドのレスポンス形式をフロントエンドの型に変換
    const result: ProductListResponse = {
      products,
      totalCount: data.total || 0,
      totalPages: totalPages,
      currentPage: currentPage,
      hasNextPage: currentPage < totalPages,
      hasPreviousPage: currentPage > 1,
    };

    return result;
  } catch (error) {
    console.error("商品一覧の取得に失敗しました:", error);
    throw error;
  }
}

/**
 * カテゴリ一覧を取得
 * @returns カテゴリ一覧
 */
export async function fetchCategories(): Promise<Category[]> {
  try {
    const response = await apiClient.get("/categories");
    // バックエンドのレスポンス形式 { items: Category[] } に対応
    return response.data.items || [];
  } catch (error) {
    console.error("カテゴリ一覧の取得に失敗しました:", error);
    throw error;
  }
}

/**
 * 商品を取得
 * @param id 商品ID
 * @returns 商品
 */
export async function fetchProduct(id: string): Promise<Product | null> {
  try {
    const response = await apiClient.get(`/products/${id}`);
    return response.data;
  } catch (error) {
    console.error("商品の取得に失敗しました:", error);
    throw error;
  }
}
