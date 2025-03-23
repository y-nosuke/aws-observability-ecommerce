import Link from 'next/link';

export default function Home() {
  return (
    <main className="min-h-screen flex flex-col items-center justify-center p-4 bg-gray-50">
      <div className="max-w-4xl w-full bg-white rounded-xl shadow-md overflow-hidden p-8">
        <h1 className="text-4xl font-bold text-center text-indigo-600 mb-6">
          AWS オブザーバビリティ学習用 eコマースサイト
        </h1>

        <div className="text-center text-lg text-gray-700 mb-8">
          <p>
            このサイトはAWSのオブザーバビリティ機能を学習するための実験的なeコマースサイトです。
          </p>
          <p className="mt-2">実際の商品販売は行っていません。</p>
        </div>

        <div className="space-y-4">
          <div className="bg-indigo-50 p-6 rounded-lg">
            <h2 className="text-2xl font-semibold text-indigo-700 mb-4">機能一覧</h2>
            <ul className="space-y-2 text-gray-700">
              <li className="flex items-center">
                <svg
                  className="h-5 w-5 text-indigo-500 mr-2"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 5l7 7-7 7"
                  />
                </svg>
                商品一覧の閲覧
              </li>
              <li className="flex items-center">
                <svg
                  className="h-5 w-5 text-indigo-500 mr-2"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 5l7 7-7 7"
                  />
                </svg>
                商品詳細の確認
              </li>
              <li className="flex items-center">
                <svg
                  className="h-5 w-5 text-indigo-500 mr-2"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 5l7 7-7 7"
                  />
                </svg>
                カテゴリー別商品表示
              </li>
            </ul>
          </div>

          <div className="flex justify-center mt-8">
            <Link
              href="/products"
              className="bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-3 px-8 rounded-lg text-lg transition-colors duration-300"
            >
              商品一覧を見る
            </Link>
          </div>
        </div>
      </div>
    </main>
  );
}
