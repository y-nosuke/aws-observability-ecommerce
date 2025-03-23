'use client';

import Image from 'next/image';
import { useState } from 'react';

interface ProductImageProps {
  imageUrl: string;
  productName: string;
  productId: number;
}

/**
 * 商品画像を表示するクライアントコンポーネント
 * 画像の読み込みエラーを処理し、エラー時には代替表示を行う
 */
export default function ProductImage({ imageUrl, productName, productId }: ProductImageProps) {
  // 画像の読み込み状態を管理するステート
  const [imageError, setImageError] = useState(false);

  // 画像エラー発生時のハンドラ
  const handleImageError = () => {
    console.warn('Product image failed to load', {
      productId,
      imageUrl,
    });
    setImageError(true);
  };

  // 画像がないか、読み込みエラーが発生した場合
  if (!imageUrl || imageError) {
    return (
      <div className="h-10 w-10 rounded-full bg-gray-200 flex items-center justify-center text-gray-500 text-xs">
        {productName?.charAt(0)?.toUpperCase() || 'N'}
      </div>
    );
  }

  // 画像を表示
  return (
    <Image
      className="h-10 w-10 rounded-full object-cover"
      src={imageUrl}
      alt={productName}
      width={40}
      height={40}
      onError={handleImageError}
    />
  );
}
