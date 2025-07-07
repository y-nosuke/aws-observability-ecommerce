'use client';

import { useEffect } from 'react';

export default function ProductsError({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    // エラーログを記録
    console.error('Products page error:', error);
  }, [error]);

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-50">
      <div className="w-full max-w-md rounded-lg bg-white p-6 shadow-lg">
        <div className="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-full bg-red-100">
          <svg
            className="h-6 w-6 text-red-600"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"
            />
          </svg>
        </div>

        <h2 className="mb-2 text-center text-xl font-semibold text-gray-900">
          商品データの取得中にエラーが発生しました
        </h2>

        <p className="mb-6 text-center text-gray-600">
          申し訳ございませんが、商品情報を読み込むことができませんでした。
          しばらく時間をおいてから再度お試しください。
        </p>

        {process.env.NODE_ENV === 'development' && (
          <div className="mb-4 rounded bg-gray-100 p-3 text-sm text-gray-700">
            <strong>開発者向け情報:</strong>
            <br />
            {error.message}
          </div>
        )}

        <div className="flex flex-col gap-3 sm:flex-row">
          <button
            onClick={reset}
            className="flex-1 rounded-md bg-blue-600 px-4 py-2 text-white transition-colors hover:bg-blue-700"
          >
            再試行
          </button>

          <button
            onClick={() => (window.location.href = '/')}
            className="flex-1 rounded-md bg-gray-200 px-4 py-2 text-gray-800 transition-colors hover:bg-gray-300"
          >
            ホームに戻る
          </button>
        </div>
      </div>
    </div>
  );
}
