"use client";

import { useState, useEffect } from "react";
import { Product, Category, SortOption } from "@/services/products/types";
import { filterAndSortProducts } from "@/services/products/api";
import FilterSidebar from "./components/FilterSidebar";
import ProductGrid from "./components/ProductGrid";
import SortSelector from "./components/SortSelector";

// プレゼンテーショナルコンポーネントに渡すプロパティの型定義
interface ProductsClientProps {
  initialProducts: Product[];
  categories: Category[];
}

export default function ProductsClient({ initialProducts, categories }: ProductsClientProps) {
  // クライアントサイドの状態管理
  const [activeCategory, setActiveCategory] = useState<number>(0);
  const [sortOption, setSortOption] = useState<SortOption>("recommended");
  const [selectedPriceRanges, setSelectedPriceRanges] = useState<string[]>([]);
  const [selectedStatuses, setSelectedStatuses] = useState<string[]>([]);
  const [products, setProducts] = useState<Product[]>(initialProducts);
  const [filteredCategories, setFilteredCategories] = useState<Category[]>(categories);

  // 価格フィルターに基づいて商品をフィルタリングする関数
  const applyPriceFilter = (product: Product): boolean => {
    if (selectedPriceRanges.length === 0) return true;
    
    const price = product.salePrice || product.price;
    
    return selectedPriceRanges.some(range => {
      switch (range) {
        case "under-10000":
          return price < 10000;
        case "10000-30000":
          return price >= 10000 && price < 30000;
        case "30000-50000":
          return price >= 30000 && price < 50000;
        case "over-50000":
          return price >= 50000;
        default:
          return true;
      }
    });
  };

  // 状態フィルターに基づいて商品をフィルタリングする関数
  const applyStatusFilter = (product: Product): boolean => {
    if (selectedStatuses.length === 0) return true;
    
    return selectedStatuses.some(status => {
      switch (status) {
        case "new":
          return product.isNew;
        case "sale":
          return product.salePrice !== null;
        default:
          return true;
      }
    });
  };

  // フィルター条件が変更されたときに実行されるエフェクト
  useEffect(() => {
    // 商品のフィルタリング
    let filteredProducts = filterAndSortProducts(initialProducts, activeCategory, sortOption);
    
    // 価格フィルター適用
    filteredProducts = filteredProducts.filter(applyPriceFilter);
    
    // 状態フィルター適用
    filteredProducts = filteredProducts.filter(applyStatusFilter);

    // 商品リストを更新
    setProducts(filteredProducts);

    // カテゴリごとの商品数を再計算
    const updatedCategories = categories.map(category => {
      // 「すべて」カテゴリの場合
      if (category.id === 0) {
        // すべての商品に対してフィルターだけ適用
        const count = initialProducts
          .filter(applyPriceFilter)
          .filter(applyStatusFilter)
          .length;
        return { ...category, productCount: count };
      }
      
      // 特定のカテゴリの場合、そのカテゴリに属する商品のみをカウント
      const count = initialProducts
        .filter(product => product.categoryId === category.id)
        .filter(applyPriceFilter)
        .filter(applyStatusFilter)
        .length;
      
      return { ...category, productCount: count };
    });
    
    // 更新されたカテゴリ情報をセット
    setFilteredCategories(updatedCategories);
  }, [activeCategory, sortOption, selectedPriceRanges, selectedStatuses, initialProducts, categories]);

  // シンプルにした各ハンドラー
  const handleCategoryChange = (categoryId: number) => {
    setActiveCategory(categoryId);
  };

  const handleSortChange = (option: SortOption) => {
    setSortOption(option);
  };

  const handlePriceFilterChange = (priceRanges: string[]) => {
    setSelectedPriceRanges(priceRanges);
  };

  const handleStatusFilterChange = (statuses: string[]) => {
    setSelectedStatuses(statuses);
  };

  return (
    <div className="container mx-auto px-6 py-8">
      <div className="flex flex-col md:flex-row gap-8">
        {/* フィルターサイドバー */}
        <FilterSidebar
          categories={filteredCategories}
          activeCategory={activeCategory}
          onCategoryChange={handleCategoryChange}
          selectedPriceRanges={selectedPriceRanges}
          onPriceFilterChange={handlePriceFilterChange}
          selectedStatuses={selectedStatuses}
          onStatusFilterChange={handleStatusFilterChange}
        />

        {/* 商品一覧 */}
        <div className="flex-1">
          <div className="flex justify-between items-center mb-6 flex-wrap gap-4">
            <h1 className="text-2xl md:text-3xl font-bold">商品一覧</h1>
            <SortSelector
              sortOption={sortOption}
              onSortChange={handleSortChange}
            />
          </div>

          <ProductGrid products={products} />
          
          {/* 商品がない場合のメッセージ */}
          {products.length === 0 && (
            <div className="bg-gray-100 dark:bg-gray-800 p-6 rounded-lg text-center">
              <p className="text-gray-600 dark:text-gray-300">
                条件に一致する商品がありません。フィルターを変更してください。
              </p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
