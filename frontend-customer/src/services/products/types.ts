export interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  salePrice: number | null;
  category: string;
  isNew: boolean;
  imageUrl: string | null;
}

export interface Category {
  id: string;
  name: string;
}

export type SortOption = 'recommended' | 'newest' | 'price-asc' | 'price-desc';
