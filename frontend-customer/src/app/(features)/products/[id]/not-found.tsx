import Link from 'next/link';

export default function NotFound() {
  return (
    <div className="container mx-auto px-4 py-16 text-center">
      <div className="mx-auto max-w-md">
        <div className="mb-8">
          <svg
            className="mx-auto h-24 w-24 text-gray-400"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={1}
              d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.34 0-4.291-1.004-5.824-2.314M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
            />
          </svg>
        </div>

        <h1 className="mb-4 text-3xl font-bold text-gray-900">商品が見つかりません</h1>

        <p className="mb-8 leading-relaxed text-gray-600">
          お探しの商品は見つかりませんでした。
          <br />
          商品が削除されているか、URLが間違っている可能性があります。
        </p>

        <div className="space-y-3">
          <Link
            href="/products"
            className="bg-primary hover:bg-primary-dark focus:ring-primary inline-block w-full rounded-md px-6 py-3 font-medium text-white transition-colors focus:ring-2 focus:ring-offset-2 focus:outline-none"
          >
            商品一覧に戻る
          </Link>

          <Link
            href="/"
            className="inline-block w-full rounded-md bg-gray-100 px-6 py-3 font-medium text-gray-700 transition-colors hover:bg-gray-200 focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 focus:outline-none"
          >
            ホームに戻る
          </Link>
        </div>
      </div>
    </div>
  );
}
