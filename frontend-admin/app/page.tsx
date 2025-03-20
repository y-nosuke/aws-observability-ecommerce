import Link from 'next/link';

export default function Home() {
  return (
    <main className="min-h-screen flex flex-col items-center justify-center p-4 bg-gray-50">
      <div className="max-w-4xl w-full bg-white rounded-xl shadow-md overflow-hidden p-8">
        <h1 className="text-4xl font-bold text-center text-emerald-600 mb-6">
          管理者ダッシュボード
        </h1>

        <div className="text-center text-lg text-gray-700 mb-8">
          <p>AWSオブザーバビリティ学習用eコマースサイトの管理者ダッシュボードへようこそ。</p>
          <p className="mt-2">このダッシュボードでは商品管理や在庫管理などの操作が行えます。</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          <div className="bg-emerald-50 p-6 rounded-lg">
            <h2 className="text-xl font-semibold text-emerald-700 mb-4">商品管理</h2>
            <p className="text-gray-600 mb-4">商品の追加、編集、削除を行います。</p>
            <Link href="/products" className="text-emerald-600 hover:text-emerald-800 font-medium">
              商品一覧へ →
            </Link>
          </div>

          <div className="bg-emerald-50 p-6 rounded-lg">
            <h2 className="text-xl font-semibold text-emerald-700 mb-4">在庫管理</h2>
            <p className="text-gray-600 mb-4">商品の在庫状況の確認と更新を行います。</p>
            <span className="text-gray-400">準備中...</span>
          </div>

          <div className="bg-emerald-50 p-6 rounded-lg">
            <h2 className="text-xl font-semibold text-emerald-700 mb-4">注文管理</h2>
            <p className="text-gray-600 mb-4">注文の確認と処理を行います。</p>
            <span className="text-gray-400">準備中...</span>
          </div>

          <div className="bg-emerald-50 p-6 rounded-lg">
            <h2 className="text-xl font-semibold text-emerald-700 mb-4">分析レポート</h2>
            <p className="text-gray-600 mb-4">販売データの分析と統計情報を確認します。</p>
            <span className="text-gray-400">準備中...</span>
          </div>
        </div>

        <div className="flex justify-center">
          <Link
            href="/products"
            className="bg-emerald-600 hover:bg-emerald-700 text-white font-bold py-3 px-8 rounded-lg text-lg transition-colors duration-300"
          >
            商品管理を開始する
          </Link>
        </div>
      </div>
    </main>
  );
}
