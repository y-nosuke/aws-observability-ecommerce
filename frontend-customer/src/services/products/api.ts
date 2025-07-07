import { apiClient } from '../api-client';
import { Category, Product, ProductListResponse } from './types';

// 商品関連のAPI関数
export const productsApi = {
  // 商品一覧取得
  async getProducts(params?: {
    page?: number;
    pageSize?: number;
    categoryId?: number;
    keyword?: string;
  }): Promise<ProductListResponse> {
    const response = await apiClient.get('/products', { params });
    return response.data;
  },

  // 商品詳細取得
  async getProduct(id: number): Promise<Product> {
    const response = await apiClient.get(`/products/${id}`);
    return response.data;
  },

  // カテゴリー別商品取得
  async getProductsByCategory(
    categoryId: number,
    params?: {
      page?: number;
      pageSize?: number;
      keyword?: string;
    },
  ): Promise<ProductListResponse> {
    const response = await apiClient.get(`/categories/${categoryId}/products`, {
      params,
    });
    return response.data;
  },

  // カテゴリー一覧取得
  async getCategories(): Promise<{ items: Category[] }> {
    const response = await apiClient.get('/categories');
    return response.data;
  },
};

// Server Components用のfetch関数
export async function fetchProducts(params?: {
  page?: number;
  pageSize?: number;
  categoryId?: number;
  keyword?: string;
}): Promise<Product[]> {
  try {
    const result = await productsApi.getProducts(params);
    return result.items;
  } catch (error) {
    console.error('Failed to fetch products:', error);
    return [];
  }
}

// 人気商品の取得（フィーチャー商品を取得）
export async function fetchPopularProducts(limit: number = 4): Promise<Product[]> {
  try {
    const result = await productsApi.getProducts({ pageSize: 100 }); // 多めに取得
    const featured = result.items.filter((p) => p.isFeatured);
    return featured.slice(0, limit);
  } catch (error) {
    console.error('Failed to fetch popular products:', error);
    return [];
  }
}

// カテゴリー一覧の取得
export async function fetchCategories(): Promise<Category[]> {
  try {
    const result = await productsApi.getCategories();
    return result.items;
  } catch (error) {
    console.error('Failed to fetch categories:', error);
    return [];
  }
}

// 商品をIDで取得
export async function fetchProductById(id: number): Promise<Product | null> {
  try {
    return await productsApi.getProduct(id);
  } catch (error) {
    console.error('Failed to fetch product:', error);
    return null;
  }
}

// 商品のフィルタリングと並び替え（クライアントサイド用）
export function filterAndSortProducts(
  products: Product[],
  category: number = 0,
  sortOption: 'recommended' | 'price-asc' | 'price-desc' | 'newest' = 'recommended',
): Product[] {
  // カテゴリーでフィルタリング
  const filtered =
    category === 0 ? products : products.filter((product) => product.categoryId === category);

  // 並び替え
  return sortProducts(filtered, sortOption);
}

// 商品の並び替え（クライアントサイド用）
export function sortProducts(
  products: Product[],
  sortOption: 'recommended' | 'price-asc' | 'price-desc' | 'newest' = 'recommended',
): Product[] {
  switch (sortOption) {
    case 'price-asc':
      return [...products].sort((a, b) => (a.salePrice || a.price) - (b.salePrice || b.price));
    case 'price-desc':
      return [...products].sort((a, b) => (b.salePrice || b.price) - (a.salePrice || a.price));
    case 'newest':
      return [...products].sort((a, b) => (a.isNew ? -1 : b.isNew ? 1 : 0));
    default: // recommended
      return [...products].sort((a, b) =>
        a.isFeatured && !b.isFeatured ? -1 : !a.isFeatured && b.isFeatured ? 1 : 0,
      );
  }
}
