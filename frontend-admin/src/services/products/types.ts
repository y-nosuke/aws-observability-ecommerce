// 商品関連の型定義

export interface Product {
  id: string;
  name: string;
  price: number;
  stockQuantity: number;
  categoryId: number;
  categoryName: string;
  imageUrl?: string;
  description?: string;
  createdAt: string;
  updatedAt: string;
  sku?: string;
  salePrice?: number;
  inStock?: boolean;
  isNew?: boolean;
  isFeatured?: boolean;
}

export interface Category {
  id: number;
  name: string;
  description?: string;
}

export interface ProductSearchParams {
  keyword?: string;
  categoryId?: number;
  page?: number;
  limit?: number;
}

export interface ProductListResponse {
  products: Product[];
  totalCount: number;
  totalPages: number;
  currentPage: number;
  hasNextPage: boolean;
  hasPreviousPage: boolean;
}

export interface ApiError {
  message: string;
  status: number;
}
