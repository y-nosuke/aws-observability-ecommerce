'use client';

import React from 'react';

// Next.js error boundary component
export default function Error({ error, reset }: { error: Error; reset: () => void }) {
  React.useEffect(() => {
    console.error('商品ページの初期データ取得に失敗しました:', error);
  }, [error]);
  return (
    <div className="flex min-h-[400px] items-center justify-center">
      <div className="text-center">
        <h2 className="mb-2 text-lg font-semibold text-red-600">データの読み込みに失敗しました</h2>
        <p className="text-gray-600">しばらく時間をおいてから再度お試しください。</p>
        <button
          className="mt-4 rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700"
          onClick={() => reset()}
        >
          再試行
        </button>
      </div>
    </div>
  );
}
