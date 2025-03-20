'use client';

import { Product } from '@/lib/api/products';
import Image from 'next/image';
import Link from 'next/link';
import { useState } from 'react';

export default function ProductCard({ product }: { product: Product }) {
  // 画像の読み込み状態を管理するステート
  const [imageError, setImageError] = useState(false);

  return (
    <div className="bg-white rounded-lg overflow-hidden shadow-lg hover:shadow-xl transition-shadow duration-300">
      <div className="h-48 bg-gray-200 flex items-center justify-center relative">
        {product.image_url && !imageError ? (
          <div className="w-full h-full relative">
            <Image
              src={product.image_url}
              alt={product.name}
              fill
              sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
              className="object-cover"
              onError={() => {
                console.warn('Product image failed to load', {
                  productId: product.id,
                  imageUrl: product.image_url,
                });
                setImageError(true);
              }}
            />
          </div>
        ) : (
          // 画像がない場合または読み込みエラーの場合の代替表示
          <div className="flex flex-col items-center justify-center text-gray-500 p-4">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-12 w-12 mb-2"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={1.5}
                d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
              />
            </svg>
            <p className="text-center">{!product.image_url ? '商品画像なし' : product.name}</p>
          </div>
        )}
      </div>
      {/* 以下のコードは変更なし */}
      <div className="p-4">
        <h2 className="text-xl font-semibold text-gray-800 mb-2">{product.name}</h2>
        <p className="text-gray-600 text-sm mb-4 h-12 overflow-hidden">{product.description}</p>
        <div className="flex justify-between items-center">
          <span className="text-xl font-bold text-indigo-600">
            ¥{product.price.toLocaleString()}
          </span>
          <Link
            href={`/products/${product.id}`}
            className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-300"
          >
            詳細を見る
          </Link>
        </div>
      </div>
    </div>
  );
}
