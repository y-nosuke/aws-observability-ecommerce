export default function Loading() {
  return (
    <div className="container mx-auto px-4 py-8">
      {/* パンくずナビゲーション スケルトン */}
      <div className="mb-8 flex animate-pulse">
        <div className="h-4 w-12 rounded bg-gray-200"></div>
        <div className="mx-2 text-gray-300">/</div>
        <div className="h-4 w-16 rounded bg-gray-200"></div>
        <div className="mx-2 text-gray-300">/</div>
        <div className="h-4 w-24 rounded bg-gray-200"></div>
      </div>

      <div className="grid grid-cols-1 gap-12 lg:grid-cols-2">
        {/* 商品画像 スケルトン */}
        <div className="animate-pulse space-y-4">
          <div className="aspect-square rounded-lg bg-gray-200"></div>
          <div className="mx-auto h-4 w-32 rounded bg-gray-200"></div>
        </div>

        {/* 商品情報 スケルトン */}
        <div className="animate-pulse space-y-6">
          <div>
            <div className="mb-2 h-8 w-3/4 rounded bg-gray-200"></div>
            <div className="mb-4 h-4 w-32 rounded bg-gray-200"></div>
            <div className="h-4 w-28 rounded bg-gray-200"></div>
          </div>

          {/* 価格 スケルトン */}
          <div className="space-y-2">
            <div className="flex items-center gap-4">
              <div className="h-8 w-32 rounded bg-gray-200"></div>
              <div className="h-6 w-24 rounded bg-gray-200"></div>
            </div>
            <div className="h-4 w-16 rounded bg-gray-200"></div>
          </div>

          {/* 商品説明 スケルトン */}
          <div>
            <div className="mb-2 h-6 w-24 rounded bg-gray-200"></div>
            <div className="space-y-2">
              <div className="h-4 w-full rounded bg-gray-200"></div>
              <div className="h-4 w-5/6 rounded bg-gray-200"></div>
              <div className="h-4 w-4/5 rounded bg-gray-200"></div>
            </div>
          </div>

          {/* 在庫情報 スケルトン */}
          <div className="space-y-2">
            <div className="flex items-center gap-2">
              <div className="h-4 w-16 rounded bg-gray-200"></div>
              <div className="h-4 w-20 rounded bg-gray-200"></div>
            </div>
            <div className="h-4 w-16 rounded bg-gray-200"></div>
          </div>

          {/* 購入オプション スケルトン */}
          <div className="space-y-4 border-t pt-6">
            <div className="flex items-center gap-4">
              <div className="h-4 w-8 rounded bg-gray-200"></div>
              <div className="h-10 w-32 rounded bg-gray-200"></div>
            </div>
            <div className="h-12 w-full rounded bg-gray-200"></div>
          </div>

          {/* 商品情報 スケルトン */}
          <div className="space-y-2 border-t pt-6">
            <div className="h-4 w-32 rounded bg-gray-200"></div>
            <div className="h-4 w-28 rounded bg-gray-200"></div>
          </div>
        </div>
      </div>
    </div>
  );
}
