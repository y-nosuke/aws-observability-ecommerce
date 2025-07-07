'use client';

import Link from 'next/link';
import { useEffect } from 'react';

interface ErrorProps {
  error: Error & { digest?: string };
  reset: () => void;
}

export default function Error({ error, reset }: ErrorProps) {
  useEffect(() => {
    // エラーをログに記録（本番環境では適切なログサービスに送信）
    console.error('Product detail page error:', error);
  }, [error]);

  return (
    <div className="container mx-auto px-4 py-16 text-center">
      <div className="mx-auto max-w-md">
        <div className="mb-8">
          <svg
            className="mx-auto h-24 w-24 text-red-400"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={1}
              d="M12 9V6.414a1 1 0 00-1.707-.707L7.586 8.414A2 2 0 006.414 9H5a1 1 0 100 2h1.414a2 2 0 001.172.293l2.707 2.707A1 1 0 0012 13.586V11a1 1 0 100-2zm0 0V9"
            />
          </svg>
        </div>

        <h1 className="mb-4 text-3xl font-bold text-gray-900">エラーが発生しました</h1>

        <p className="mb-8 leading-relaxed text-gray-600">
          商品詳細の読み込み中にエラーが発生しました。
          <br />
          再度お試しいただくか、しばらく時間をおいてからアクセスしてください。
        </p>

        <div className="space-y-3">
          <button
            onClick={reset}
            className="bg-primary hover:bg-primary-dark focus:ring-primary inline-block w-full rounded-md px-6 py-3 font-medium text-white transition-colors focus:ring-2 focus:ring-offset-2 focus:outline-none"
          >
            再試行
          </button>

          <Link
            href="/products"
            className="inline-block w-full rounded-md bg-gray-100 px-6 py-3 font-medium text-gray-700 transition-colors hover:bg-gray-200 focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 focus:outline-none"
          >
            商品一覧に戻る
          </Link>

          <Link
            href="/"
            className="inline-block w-full rounded-md border border-gray-300 bg-white px-6 py-3 font-medium text-gray-700 transition-colors hover:bg-gray-50 focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 focus:outline-none"
          >
            ホームに戻る
          </Link>
        </div>

        {process.env.NODE_ENV === 'development' && (
          <details className="mt-8 text-left">
            <summary className="cursor-pointer text-sm text-gray-500 hover:text-gray-700">
              エラー詳細 (開発環境)
            </summary>
            <pre className="mt-4 overflow-auto rounded-md bg-gray-100 p-4 text-xs text-gray-700">
              {error.message}
            </pre>
          </details>
        )}
      </div>
    </div>
  );
}
