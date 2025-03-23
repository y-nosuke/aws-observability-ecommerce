// 商品の型定義
export interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  image_url: string;
  category_id: number;
}

// カテゴリーの型定義
export interface Category {
  id: number;
  name: string;
  description: string;
}

// ページネーションレスポンスの型定義
export interface PaginatedResponse<T> {
  items: T[];
  total_items: number;
  page: number;
  page_size: number;
  total_pages: number;
}

// API呼び出しパラメータの型定義
export interface GetProductsParams {
  page?: number;
  page_size?: number;
  category_id?: number;
}

// APIのベースURL
// Next.jsのサーバーコンポーネントの場合、絶対URLが必要
// 開発環境では環境変数からAPIのURLを取得するか、デフォルトで絶対パスを使用
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://backend:8080/api';

/**
 * API呼び出しのエラー処理を行うヘルパー関数
 * @param response fetchのレスポンス
 */
async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const errorText = await response.text();
    console.error('API request failed', {
      status: response.status,
      statusText: response.statusText,
      url: response.url,
      error: errorText,
    });

    throw new Error(`API request failed: ${response.status} ${response.statusText}`);
  }

  return response.json() as Promise<T>;
}

/**
 * 商品一覧を取得する
 * @param params 検索パラメータ
 */
export async function getProducts(
  params: GetProductsParams = {}
): Promise<PaginatedResponse<Product>> {
  const queryParams = new URLSearchParams();

  // パラメータの設定
  if (params.page) queryParams.append('page', params.page.toString());
  if (params.page_size) queryParams.append('page_size', params.page_size.toString());
  if (params.category_id) queryParams.append('category_id', params.category_id.toString());

  const url = `${API_BASE_URL}/products?${queryParams.toString()}`;

  try {
    console.info('Fetching products', { url });

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        Accept: 'application/json',
      },
      // キャッシュ戦略を修正、cache: 'no-store'はNext.js 14でサポート
      cache: 'no-store',
    });

    const data = await handleResponse<PaginatedResponse<Product>>(response);
    console.info('Products fetched successfully', { count: data.items.length });
    return data;
  } catch (error) {
    console.error('Failed to fetch products', {
      error,
      url,
      message: error instanceof Error ? error.message : String(error),
    });
    throw error;
  }
}

/**
 * カテゴリー一覧を取得する
 */
export async function getCategories(): Promise<Category[]> {
  const url = `${API_BASE_URL}/products/categories`;

  try {
    console.info('Fetching categories', { url });

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        Accept: 'application/json',
      },
      // キャッシュ戦略を修正
      cache: 'no-store',
    });

    const data = await handleResponse<Category[]>(response);
    console.info('Categories fetched successfully', { count: data.length });
    return data;
  } catch (error) {
    console.error('Failed to fetch categories', {
      error,
      url,
      message: error instanceof Error ? error.message : String(error),
    });
    throw error;
  }
}

/**
 * 商品の詳細を取得する
 * @param id 商品ID
 */
export async function getProductById(id: number): Promise<Product> {
  const url = `${API_BASE_URL}/products/${id}`;

  try {
    console.info('Fetching product details', { id, url });

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        Accept: 'application/json',
      },
      // キャッシュ戦略を修正
      cache: 'no-store',
    });

    const data = await handleResponse<Product>(response);
    console.info('Product details fetched successfully', { product_id: id });
    return data;
  } catch (error) {
    console.error('Failed to fetch product details', {
      id,
      url,
      error: error instanceof Error ? error.message : String(error),
    });
    throw error;
  }
}

// APIクライアントをエクスポート
export const ProductsApi = {
  getProducts,
  getCategories,
  getProductById,
};

export default ProductsApi;
