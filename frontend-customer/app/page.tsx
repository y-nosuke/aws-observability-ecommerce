import AnimateInView from "@/components/ui/AnimateInView";
import ProductCard from "@/components/ui/ProductCard";

export default function Home() {
  return (
    <>
      {/* ヒーローセクション */}
      <AnimateInView>
        <div className="hero-gradient text-white py-20 md:py-32 mb-12 relative overflow-hidden">
          <div className="absolute top-0 left-0 w-full h-full bg-black opacity-40"></div>
          <div className="container mx-auto px-6 relative z-10">
            <div className="max-w-3xl mx-auto text-center">
              <AnimateInView direction="down" delay={200}>
                <h1 className="text-4xl md:text-5xl font-bold mb-6 leading-tight">
                  <span className="block mb-2">高品質な商品を</span>
                  <span className="bg-clip-text text-transparent bg-gradient-to-r from-white to-indigo-200">
                    お求めやすい価格で
                  </span>
                </h1>
              </AnimateInView>

              <AnimateInView direction="up" delay={400}>
                <p className="text-xl md:text-2xl mb-10 text-gray-100">
                  オブザーバビリティショップで最高のショッピング体験をお楽しみください
                </p>

                <div className="flex flex-col sm:flex-row justify-center gap-4">
                  <a
                    href="/products"
                    className="btn-primary text-white py-3 px-8 rounded-full font-medium text-lg shadow-lg transition-all hover:shadow-xl"
                  >
                    商品を見る
                  </a>
                  <a
                    href="/about"
                    className="bg-white/20 backdrop-blur-sm text-white py-3 px-8 rounded-full font-medium text-lg hover:bg-white/30 transition-colors"
                  >
                    詳細を見る
                  </a>
                </div>
              </AnimateInView>
            </div>
          </div>

          {/* 装飾要素 */}
          <div className="absolute -bottom-16 -left-16 w-64 h-64 bg-indigo-600/20 rounded-full blur-3xl"></div>
          <div className="absolute -top-20 -right-20 w-72 h-72 bg-purple-600/20 rounded-full blur-3xl"></div>
        </div>
      </AnimateInView>

      {/* メインコンテンツ */}
      <div className="container mx-auto px-6">
        <section className="mb-16">
          <AnimateInView direction="up" delay={100}>
            <div className="flex justify-between items-center mb-8">
              <h2 className="text-2xl font-bold">人気の商品</h2>
              <a
                href="/products"
                className="text-primary hover:text-primary-dark font-medium flex items-center transition-colors"
              >
                すべて見る
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 ml-1"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fillRule="evenodd"
                    d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                    clipRule="evenodd"
                  />
                </svg>
              </a>
            </div>
          </AnimateInView>

          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
            <AnimateInView direction="up" delay={200}>
              <ProductCard
                id="1"
                name="超高性能ノートPC"
                description="最新のプロセッサ・16GB RAM・高速 SSD 搭載"
                price={198000}
                isNew={true}
              />
            </AnimateInView>

            <AnimateInView direction="up" delay={300}>
              <ProductCard
                id="2"
                name="ワイヤレスイヤホン"
                description="ノイズキャンセリング機能付き高音質イヤホン"
                price={19800}
                salePrice={14800}
              />
            </AnimateInView>

            <AnimateInView direction="up" delay={400}>
              <ProductCard
                id="3"
                name="高画質タブレット"
                description="10.2インチディスプレイ、長持ちバッテリー搭載"
                price={36800}
              />
            </AnimateInView>

            <AnimateInView direction="up" delay={500}>
              <ProductCard
                id="4"
                name="スマートウォッチ"
                description="健康管理機能付き、長持ちバッテリー"
                price={32800}
              />
            </AnimateInView>
          </div>
        </section>
      </div>

      <section className="mb-16">
        <AnimateInView direction="up" delay={200}>
          <h2 className="text-2xl font-bold mb-8">特集・キャンペーン</h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* キャンペーン1 */}
            <div className="bg-gradient-to-r from-indigo-500 to-purple-600 rounded-2xl shadow-lg overflow-hidden card-hover">
              <div className="p-6 md:p-8 flex flex-col h-full">
                <div className="badge mb-4 self-start">SALE</div>
                <h3 className="text-xl md:text-2xl font-bold text-white mb-2">
                  夏のビッグセール
                </h3>
                <p className="text-indigo-100 mb-6">
                  期間限定で全商品20%オフ。今すぐチェック！
                </p>
                <a
                  href="/sale"
                  className="mt-auto inline-block bg-white text-indigo-600 px-6 py-2 rounded-full font-medium hover:bg-opacity-90 transition-colors self-start"
                >
                  詳細を見る
                </a>
              </div>
            </div>

            {/* キャンペーン2 */}
            <div className="bg-gradient-to-r from-rose-400 to-red-500 rounded-2xl shadow-lg overflow-hidden card-hover">
              <div className="p-6 md:p-8 flex flex-col h-full">
                <div className="badge mb-4 self-start bg-gradient-to-r from-yellow-400 to-orange-500">
                  限定商品
                </div>
                <h3 className="text-xl md:text-2xl font-bold text-white mb-2">
                  新商品登場
                </h3>
                <p className="text-rose-100 mb-6">
                  最新モデルが登場。先行予約で特典プレゼント！
                </p>
                <a
                  href="/new-arrivals"
                  className="mt-auto inline-block bg-white text-rose-600 px-6 py-2 rounded-full font-medium hover:bg-opacity-90 transition-colors self-start"
                >
                  チェックする
                </a>
              </div>
            </div>
          </div>
        </AnimateInView>
      </section>

      {/* お得な情報セクション */}
      <section className="mb-16">
        <AnimateInView direction="up" delay={300}>
          <div className="bg-gray-100 dark:bg-gray-800 rounded-2xl p-6 md:p-8 shadow-md">
            <div className="flex flex-col md:flex-row items-center">
              <div className="mb-6 md:mb-0 md:mr-8">
                <div className="bg-primary/10 dark:bg-primary/20 p-3 rounded-full w-16 h-16 flex items-center justify-center mb-4 mx-auto md:mx-0">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-8 w-8 text-primary"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                    />
                  </svg>
                </div>
              </div>
              <div className="text-center md:text-left md:flex-1">
                <h3 className="text-xl font-bold mb-2">
                  お得な情報をお届けします
                </h3>
                <p className="text-gray-600 dark:text-gray-300 mb-6">
                  メールマガジンに登録すると、新商品情報や限定セールのお知らせをいち早くお届けします。
                </p>
                <div className="flex flex-col sm:flex-row gap-3">
                  <input
                    type="email"
                    placeholder="メールアドレス"
                    className="px-4 py-3 rounded-lg border-0 shadow-sm flex-1"
                  />
                  <button className="btn-primary text-white py-3 px-6 rounded-lg shadow-sm font-medium">
                    登録する
                  </button>
                </div>
              </div>
            </div>
          </div>
        </AnimateInView>
      </section>
    </>
  );
}
