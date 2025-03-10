import Link from "next/link";

export default function Home() {
  return (
    <div className="py-12">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <h1 className="text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight lg:text-6xl">
            AWSオブザーバビリティ学習用eコマースアプリ
          </h1>
          <p className="mt-5 max-w-xl mx-auto text-xl text-gray-500">
            このアプリはAWSオブザーバビリティパターンを学習するための参照実装です。
          </p>
          <div className="mt-8 flex justify-center">
            <div className="inline-flex rounded-md shadow">
              <Link
                href="/products"
                className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
              >
                商品を見る
              </Link>
            </div>
            <div className="ml-3 inline-flex">
              <Link
                href="/admin/dashboard"
                className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-blue-700 bg-blue-100 hover:bg-blue-200"
              >
                管理者ページ
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
