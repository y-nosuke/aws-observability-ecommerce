"use client";

import Image from "next/image";
import Link from "next/link";

interface ProductCardProps {
  id: number;
  name: string;
  description?: string;
  price: number;
  salePrice?: number | null;
  isNew?: boolean;
  imageUrl?: string | null;
}

export default function ProductCard({
  id,
  name,
  description,
  price,
  salePrice = null,
  isNew = false,
  imageUrl = null,
}: ProductCardProps) {
  return (
    <div className="product-card bg-white dark:bg-gray-800 shadow-md">
      {isNew && <div className="sale-badge">新着</div>}
      {salePrice && (
        <div className="sale-badge bg-gradient-to-r from-red-500 to-pink-500">
          SALE
        </div>
      )}
      <div className="image-container h-52 bg-gray-100 dark:bg-gray-700 overflow-hidden">
        {imageUrl ? (
          <Image
            src={imageUrl}
            alt={name}
            width={400}
            height={300}
            priority={true}
            className="w-full h-full object-cover"
          />
        ) : (
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
        )}
      </div>
      <div className="p-4">
        <Link href={`/products/${id}`} className="block">
          <h3 className="font-semibold text-lg mb-1 hover:text-primary transition-colors">
            {name}
          </h3>
        </Link>
        <p className="text-gray-600 dark:text-gray-300 text-sm mb-3 line-clamp-2">
          {description}
        </p>
        <div className="flex justify-between items-center">
          {salePrice ? (
            <div className="price sale">
              <span>¥{salePrice.toLocaleString()}</span>
              <span className="original">¥{price.toLocaleString()}</span>
            </div>
          ) : (
            <div className="price">¥{price.toLocaleString()}</div>
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
  );
}
