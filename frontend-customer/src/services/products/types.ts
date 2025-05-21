export interface Product {
  id: number;
  name: string;
  description?: string;
  sku?: string;
  price: number;
  salePrice?: number | null;
  imageUrl?: string | null;
  inStock: boolean;
  stockQuantity?: number;
  categoryId: number;
  categoryName?: string;
  isNew: boolean;
  isFeatured?: boolean;
  createdAt?: string;
  updatedAt?: string;
}

export interface Category {
  id: number;
  name: string;
  slug: string;
  description?: string | null;
  imageUrl?: string | null;
  parentId?: number | null;
  productCount?: number;
}

export type SortOption = 'recommended' | 'newest' | 'price-asc' | 'price-desc';
