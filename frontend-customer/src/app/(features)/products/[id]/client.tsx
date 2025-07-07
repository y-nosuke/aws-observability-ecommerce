'use client';

import Image from 'next/image';
import Link from 'next/link';
import { useState } from 'react';

import { Product } from '@/services/products/types';

import ImageModal from './components/ImageModal';

interface ProductDetailClientProps {
  product: Product;
}

export default function ProductDetailClient({ product }: ProductDetailClientProps) {
  const [imageModalOpen, setImageModalOpen] = useState(false);
  const [quantity, setQuantity] = useState(1);

  const handleAddToCart = () => {
    // カート機能の実装は今回の範囲外ですが、プレースホルダーとして実装
    console.log(`Added ${quantity} of ${product.name} to cart`);
    alert(`「${product.name}」を${quantity}個カートに追加しました`);
  };

  // 数量変更時のバリデーションをより厳密に
  const handleQuantityChange = (value: number) => {
    if (isNaN(value) || value < 1) {
      setQuantity(1);
      return;
    }
    const maxQuantity = product.stockQuantity || 99;
    setQuantity(Math.min(value, maxQuantity));
  };

  const displayPrice = product.salePrice || product.price;
  const hasDiscount = product.salePrice && product.salePrice < product.price;

  // 商品説明文のXSS対策
  const sanitizedDescription = product.description ? product.description.replace(/[<>]/g, '') : '';

  return (
    <div className="container mx-auto px-4 py-8">
      {/* パンくずナビゲーション */}
      <nav className="mb-8 flex text-sm" aria-label="パンくずリスト">
        <Link href="/" className="text-gray-500 hover:text-gray-700">
          ホーム
        </Link>
        <span className="mx-2 text-gray-500" aria-hidden="true">
          /
        </span>
        <Link href="/products" className="text-gray-500 hover:text-gray-700">
          商品一覧
        </Link>
        <span className="mx-2 text-gray-500" aria-hidden="true">
          /
        </span>
        <span className="text-gray-900">{product.name}</span>
      </nav>

      <div className="grid grid-cols-1 gap-12 lg:grid-cols-2">
        {/* 商品画像 */}
        <div className="space-y-4">
          <div className="relative aspect-square overflow-hidden rounded-lg bg-gray-100">
            {product.imageUrl ? (
              <Image
                src={product.imageUrl}
                alt={product.name}
                fill
                className="cursor-pointer object-cover transition-transform duration-300 hover:scale-105"
                onClick={() => setImageModalOpen(true)}
                sizes="(max-width: 768px) 100vw, 50vw"
                priority
              />
            ) : (
              <div className="flex h-full w-full items-center justify-center bg-gradient-to-br from-gray-200 to-gray-300">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-24 w-24 text-gray-400"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  aria-hidden="true"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={1}
                    d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
                  />
                </svg>
              </div>
            )}

            {/* バッジ */}
            {product.isNew && (
              <div
                className="badge absolute top-4 left-4 bg-blue-500"
                role="status"
                aria-label="新着商品"
              >
                新着
              </div>
            )}
            {hasDiscount && (
              <div
                className="badge absolute top-4 right-4 bg-red-500"
                role="status"
                aria-label="セール中"
              >
                SALE
              </div>
            )}
          </div>

          {/* 画像クリックヒント */}
          {product.imageUrl && (
            <p className="text-center text-sm text-gray-500">画像をクリックして拡大表示</p>
          )}
        </div>

        {/* 商品情報 */}
        <div className="space-y-6">
          <div>
            <h1 className="mb-2 text-3xl font-bold text-gray-900">{product.name}</h1>
            {product.categoryName && (
              <p className="mb-4 text-sm text-gray-500">カテゴリー: {product.categoryName}</p>
            )}
            {product.sku && <p className="mb-4 text-sm text-gray-500">SKU: {product.sku}</p>}
          </div>

          {/* 価格 */}
          <div className="space-y-2">
            <div className="flex items-center gap-4">
              <span className="text-primary text-3xl font-bold">
                ¥{displayPrice.toLocaleString()}
              </span>
              {hasDiscount && (
                <span className="text-xl text-gray-500 line-through">
                  ¥{product.price.toLocaleString()}
                </span>
              )}
            </div>
            {hasDiscount && (
              <p className="text-sm font-medium text-red-600">
                {Math.round(((product.price - product.salePrice!) / product.price) * 100)}% OFF
              </p>
            )}
          </div>

          {/* 商品説明 */}
          {sanitizedDescription && (
            <div>
              <h2 className="mb-2 text-lg font-semibold">商品説明</h2>
              <p className="leading-relaxed whitespace-pre-line text-gray-700">
                {sanitizedDescription}
              </p>
            </div>
          )}

          {/* 在庫情報 */}
          <div className="space-y-2">
            <div className="flex items-center gap-2">
              <span className="text-sm font-medium">在庫状況:</span>
              {product.inStock ? (
                <span className="flex items-center gap-1 font-medium text-green-600">
                  <svg
                    className="h-4 w-4"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                    aria-hidden="true"
                  >
                    <path
                      fillRule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                      clipRule="evenodd"
                    />
                  </svg>
                  在庫あり
                </span>
              ) : (
                <span className="flex items-center gap-1 font-medium text-red-600">
                  <svg
                    className="h-4 w-4"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                    aria-hidden="true"
                  >
                    <path
                      fillRule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                      clipRule="evenodd"
                    />
                  </svg>
                  売り切れ
                </span>
              )}
            </div>
            {product.stockQuantity && product.stockQuantity > 0 && (
              <p className="text-sm text-gray-600">残り{product.stockQuantity}個</p>
            )}
          </div>

          {/* 購入オプション */}
          {product.inStock && (
            <div className="space-y-4 border-t pt-6">
              <div className="flex items-center gap-4">
                <label htmlFor="quantity" className="text-sm font-medium">
                  数量:
                </label>
                <div className="flex items-center rounded-md border">
                  <button
                    onClick={() => handleQuantityChange(quantity - 1)}
                    className="px-3 py-2 text-gray-600 hover:bg-gray-100 hover:text-gray-800"
                    disabled={quantity <= 1}
                    aria-label="数量を減らす"
                  >
                    -
                  </button>
                  <input
                    id="quantity"
                    type="number"
                    min="1"
                    max={product.stockQuantity || 99}
                    value={quantity}
                    onChange={(e) => handleQuantityChange(parseInt(e.target.value) || 1)}
                    className="focus:ring-primary w-16 border-x px-3 py-2 text-center focus:ring-2 focus:outline-none"
                    aria-label="商品の数量"
                  />
                  <button
                    onClick={() => handleQuantityChange(quantity + 1)}
                    className="px-3 py-2 text-gray-600 hover:bg-gray-100 hover:text-gray-800"
                    disabled={quantity >= (product.stockQuantity || 99)}
                    aria-label="数量を増やす"
                  >
                    +
                  </button>
                </div>
              </div>

              <button
                onClick={handleAddToCart}
                className="bg-primary hover:bg-primary-dark focus:ring-primary w-full rounded-md px-6 py-3 font-medium text-white transition-colors focus:ring-2 focus:ring-offset-2 focus:outline-none"
                aria-label={`${product.name}を${quantity}個カートに追加`}
              >
                カートに追加 - ¥{(displayPrice * quantity).toLocaleString()}
              </button>
            </div>
          )}

          {/* 商品情報 */}
          <div className="space-y-2 border-t pt-6 text-sm text-gray-600">
            {product.createdAt && (
              <p>登録日: {new Date(product.createdAt).toLocaleDateString('ja-JP')}</p>
            )}
            {product.updatedAt && (
              <p>更新日: {new Date(product.updatedAt).toLocaleDateString('ja-JP')}</p>
            )}
          </div>
        </div>
      </div>

      {/* 画像モーダル */}
      {product.imageUrl && (
        <ImageModal
          isOpen={imageModalOpen}
          onClose={() => setImageModalOpen(false)}
          imageUrl={product.imageUrl}
          altText={product.name}
        />
      )}
    </div>
  );
}
