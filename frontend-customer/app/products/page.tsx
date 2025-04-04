"use client";

import Link from "next/link";
import { useState } from "react";

interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  salePrice: number | null;
  category: string;
  isNew: boolean;
  imageUrl: string | null;
}

export default function ProductsPage() {
  const [activeCategory, setActiveCategory] = useState("all");
  const [sortOption, setSortOption] = useState("recommended");

  // カテゴリーリスト
  const categories = [
    { id: "all", name: "すべて" },
    { id: "computers", name: "コンピューター" },
    { id: "phones", name: "スマートフォン" },
    { id: "audio", name: "オーディオ" },
    { id: "wearables", name: "ウェアラブル" },
    { id: "accessories", name: "アクセサリー" },
  ];

  // 商品データ (実際のシステムではAPIから取得)
  const products = [
    {
      id: "1",
      name: "超高性能ノートPC",
      description: "最新のプロセッサ・16GB RAM・高速 SSD 搭載",
      price: 198000,
      salePrice: null,
      category: "computers",
      isNew: true,
      imageUrl: null,
    },
    {
      id: "2",
      name: "ワイヤレスイヤホン",
      description: "ノイズキャンセリング機能付き高音質イヤホン",
      price: 19800,
      salePrice: 14800,
      category: "audio",
      isNew: false,
      imageUrl: null,
    },
    {
      id: "3",
      name: "高画質タブレット",
      description: "10.2インチディスプレイ、長持ちバッテリー搭載",
      price: 36800,
      salePrice: null,
      category: "computers",
      isNew: false,
      imageUrl: null,
    },
    {
      id: "4",
      name: "スマートウォッチ",
      description: "健康管理機能付き、長持ちバッテリー",
      price: 32800,
      salePrice: null,
      category: "wearables",
      isNew: false,
      imageUrl: null,
    },
    {
      id: "5",
      name: "4Kモニター 27インチ",
      description: "鮮明な映像、広色域ディスプレイ",
      price: 49800,
      salePrice: 44800,
      category: "computers",
      isNew: true,
      imageUrl: null,
    },
    {
      id: "6",
      name: "ハイエンドスマートフォン",
      description: "最新チップ、高性能カメラ、大容量バッテリー",
      price: 128000,
      salePrice: null,
      category: "phones",
      isNew: true,
      imageUrl: null,
    },
    {
      id: "7",
      name: "Bluetoothスピーカー",
      description: "防水機能付き、高音質ポータブルスピーカー",
      price: 12800,
      salePrice: 9800,
      category: "audio",
      isNew: false,
      imageUrl: null,
    },
    {
      id: "8",
      name: "ワイヤレスマウス",
      description: "高精度センサー、長持ちバッテリー",
      price: 6800,
      salePrice: null,
      category: "accessories",
      isNew: false,
      imageUrl: null,
    },
  ];

  // 表示する商品をフィルタリング
  const filteredProducts = products.filter(
    (product) => activeCategory === "all" || product.category === activeCategory
  );

  // 商品を並び替え
  const sortProducts = (products: Product[]) => {
    switch (sortOption) {
      case "price-asc":
        return [...products].sort(
          (a, b) => (a.salePrice || a.price) - (b.salePrice || b.price)
        );
      case "price-desc":
        return [...products].sort(
          (a, b) => (b.salePrice || b.price) - (a.salePrice || a.price)
        );
      case "newest":
        return [...products].sort((a, b) => (a.isNew ? -1 : b.isNew ? 1 : 0));
      default: // recommended
        return products;
    }
  };

  const sortedProducts = sortProducts(filteredProducts);

  return (
    <div className="container mx-auto px-6 py-8">
      <div className="flex flex-col md:flex-row gap-8">
        {/* サイドバー */}
        <div className="w-full md:w-64 shrink-0">
          <div className="sticky top-28">
            <h2 className="text-xl font-bold mb-4">カテゴリー</h2>
            <ul className="space-y-2 mb-8">
              {categories.map((category) => (
                <li key={category.id}>
                  <button
                    onClick={() => setActiveCategory(category.id)}
                    className={`w-full text-left py-2 px-3 rounded-lg transition-colors ${
                      activeCategory === category.id
                        ? "bg-primary text-white font-medium"
                        : "hover:bg-gray-100 dark:hover:bg-gray-800"
                    }`}
                  >
                    {category.name}
                  </button>
                </li>
              ))}
            </ul>

            <h2 className="text-xl font-bold mb-4">価格</h2>
            <div className="space-y-2 mb-8">
              <div className="flex items-center">
                <input type="checkbox" id="price-1" className="mr-3" />
                <label htmlFor="price-1">¥10,000以下</label>
              </div>
              <div className="flex items-center">
                <input type="checkbox" id="price-2" className="mr-3" />
                <label htmlFor="price-2">¥10,000〜¥30,000</label>
              </div>
              <div className="flex items-center">
                <input type="checkbox" id="price-3" className="mr-3" />
                <label htmlFor="price-3">¥30,000〜¥50,000</label>
              </div>
              <div className="flex items-center">
                <input type="checkbox" id="price-4" className="mr-3" />
                <label htmlFor="price-4">¥50,000以上</label>
              </div>
            </div>

            <h2 className="text-xl font-bold mb-4">状態</h2>
            <div className="space-y-2">
              <div className="flex items-center">
                <input type="checkbox" id="condition-new" className="mr-3" />
                <label htmlFor="condition-new">新着商品</label>
              </div>
              <div className="flex items-center">
                <input type="checkbox" id="condition-sale" className="mr-3" />
                <label htmlFor="condition-sale">セール中</label>
              </div>
            </div>
          </div>
        </div>

        {/* 商品一覧 */}
        <div className="flex-1">
          <div className="flex justify-between items-center mb-6 flex-wrap gap-4">
            <h1 className="text-2xl md:text-3xl font-bold">商品一覧</h1>

            <div className="flex items-center">
              <label htmlFor="sort" className="mr-2 text-sm">
                並び替え:
              </label>
              <select
                id="sort"
                value={sortOption}
                onChange={(e) => setSortOption(e.target.value)}
                className="bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-700 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"
              >
                <option value="recommended">おすすめ順</option>
                <option value="newest">新着順</option>
                <option value="price-asc">価格: 安い順</option>
                <option value="price-desc">価格: 高い順</option>
              </select>
            </div>
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {sortedProducts.map((product) => (
              <div
                key={product.id}
                className="product-card bg-white dark:bg-gray-800 shadow-md"
              >
                {product.isNew && <div className="sale-badge">新着</div>}
                {product.salePrice && (
                  <div className="sale-badge bg-gradient-to-r from-red-500 to-pink-500">
                    SALE
                  </div>
                )}
                <div className="image-container h-52 bg-gray-100 dark:bg-gray-700 overflow-hidden">
                  <div className="bg-gradient-to-br from-gray-200 to-gray-300 dark:from-gray-600 dark:to-gray-700 h-full w-full flex items-center justify-center">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-16 w-16 text-gray-400"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth={1}
                        d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                      />
                    </svg>
                  </div>
                </div>
                <div className="p-4">
                  <Link href={`/products/${product.id}`} className="block">
                    <h3 className="font-semibold text-lg mb-1 hover:text-primary transition-colors">
                      {product.name}
                    </h3>
                  </Link>
                  <p className="text-gray-600 dark:text-gray-300 text-sm mb-3 line-clamp-2">
                    {product.description}
                  </p>
                  <div className="flex justify-between items-center">
                    {product.salePrice ? (
                      <div className="price sale">
                        <span>¥{product.salePrice.toLocaleString()}</span>
                        <span className="original">
                          ¥{product.price.toLocaleString()}
                        </span>
                      </div>
                    ) : (
                      <div className="price">
                        ¥{product.price.toLocaleString()}
                      </div>
                    )}
                    <button className="p-2 rounded-full bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600 transition-colors">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className="h-5 w-5"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
                        />
                      </svg>
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
