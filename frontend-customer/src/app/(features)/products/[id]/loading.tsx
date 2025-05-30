export default function Loading() {
  return (
    <div className="container mx-auto px-4 py-8">
      {/* パンくずナビゲーション スケルトン */}
      <div className="flex mb-8 animate-pulse">
        <div className="h-4 bg-gray-200 rounded w-12"></div>
        <div className="mx-2 text-gray-300">/</div>
        <div className="h-4 bg-gray-200 rounded w-16"></div>
        <div className="mx-2 text-gray-300">/</div>
        <div className="h-4 bg-gray-200 rounded w-24"></div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
        {/* 商品画像 スケルトン */}
        <div className="space-y-4 animate-pulse">
          <div className="aspect-square bg-gray-200 rounded-lg"></div>
          <div className="h-4 bg-gray-200 rounded w-32 mx-auto"></div>
        </div>

        {/* 商品情報 スケルトン */}
        <div className="space-y-6 animate-pulse">
          <div>
            <div className="h-8 bg-gray-200 rounded w-3/4 mb-2"></div>
            <div className="h-4 bg-gray-200 rounded w-32 mb-4"></div>
            <div className="h-4 bg-gray-200 rounded w-28"></div>
          </div>

          {/* 価格 スケルトン */}
          <div className="space-y-2">
            <div className="flex items-center gap-4">
              <div className="h-8 bg-gray-200 rounded w-32"></div>
              <div className="h-6 bg-gray-200 rounded w-24"></div>
            </div>
            <div className="h-4 bg-gray-200 rounded w-16"></div>
          </div>

          {/* 商品説明 スケルトン */}
          <div>
            <div className="h-6 bg-gray-200 rounded w-24 mb-2"></div>
            <div className="space-y-2">
              <div className="h-4 bg-gray-200 rounded w-full"></div>
              <div className="h-4 bg-gray-200 rounded w-5/6"></div>
              <div className="h-4 bg-gray-200 rounded w-4/5"></div>
            </div>
          </div>

          {/* 在庫情報 スケルトン */}
          <div className="space-y-2">
            <div className="flex items-center gap-2">
              <div className="h-4 bg-gray-200 rounded w-16"></div>
              <div className="h-4 bg-gray-200 rounded w-20"></div>
            </div>
            <div className="h-4 bg-gray-200 rounded w-16"></div>
          </div>

          {/* 購入オプション スケルトン */}
          <div className="space-y-4 border-t pt-6">
            <div className="flex items-center gap-4">
              <div className="h-4 bg-gray-200 rounded w-8"></div>
              <div className="h-10 bg-gray-200 rounded w-32"></div>
            </div>
            <div className="h-12 bg-gray-200 rounded w-full"></div>
          </div>

          {/* 商品情報 スケルトン */}
          <div className="border-t pt-6 space-y-2">
            <div className="h-4 bg-gray-200 rounded w-32"></div>
            <div className="h-4 bg-gray-200 rounded w-28"></div>
          </div>
        </div>
      </div>
    </div>
  );
}
