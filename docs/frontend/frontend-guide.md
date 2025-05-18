# Next.js と React フロントエンド開発ガイド

## 目次

1. [フロントエンド開発の基礎](#1-フロントエンド開発の基礎)
2. [React の基本概念](#2-react-の基本概念)
3. [Next.js フレームワークの概要](#3-nextjs-フレームワークの概要)
4. [プロジェクト構造の理解](#4-プロジェクト構造の理解)
5. [主要コンポーネントとその役割](#5-主要コンポーネントとその役割)
6. [クライアントサイドとサーバーサイドのレンダリング](#6-クライアントサイドとサーバーサイドのレンダリング)
7. [TailwindCSS によるスタイリング](#7-tailwindcss-によるスタイリング)
8. [React でのステート管理](#8-react-でのステート管理)
9. [Next.js でのナビゲーションとルーティング](#9-nextjs-でのナビゲーションとルーティング)
10. [API リクエストの処理](#10-api-リクエストの処理)
11. [フロントエンドコンポーネントのテスト](#11-フロントエンドコンポーネントのテスト)
12. [顧客向けと管理者向け画面の違い](#12-顧客向けと管理者向け画面の違い)
13. [次のステップと学習リソース](#13-次のステップと学習リソース)

## 1. フロントエンド開発の基礎

フロントエンド開発とは、ウェブサイトやウェブアプリケーションのユーザーインターフェース（UI）を構築するプロセスのことです。ユーザーが直接操作する部分を担当し、見た目と操作性に関わるすべての要素を含みます。

### 主要技術

- **HTML**: コンテンツ構造を定義
- **CSS**: 見た目やレイアウトを制御
- **JavaScript**: インタラクティブな動作を実装

### モダンフロントエンド開発の特徴

今回のプロジェクトで使用されているのは、モダンなフロントエンド開発のアプローチです：

1. **コンポーネントベース開発**: UI を再利用可能な独立したパーツに分割
2. **宣言的なプログラミング**: 「何を」表示したいかを定義し、「どのように」表示するかはフレームワークに任せる
3. **仮想 DOM**: 実際の DOM 操作を最小限に抑えるために、仮想的な DOM を使用してパフォーマンスを向上
4. **単一方向データフロー**: データは親から子へと一方向に流れる設計

## 2. React の基本概念

React は Facebook（現 Meta）によって開発された JavaScript ライブラリで、ユーザーインターフェースを構築するために使用されます。

### コンポーネント

React の最も基本的な概念は「コンポーネント」です。コンポーネントは独立した再利用可能な UI の部品です。

例: `ProductCard.tsx` は商品カードを表示するコンポーネントです：

```jsx
export default function ProductCard({ product }: { product: Product }) {
  // ...
  return (
    <div className="bg-white rounded-lg overflow-hidden shadow-lg hover:shadow-xl transition-shadow duration-300">
      {/* 商品画像 */}
      <div className="h-48 bg-gray-200 flex items-center justify-center relative">
        {/* ... */}
      </div>
      {/* 商品情報 */}
      <div className="p-4">
        <h2 className="text-xl font-semibold text-gray-800 mb-2">{product.name}</h2>
        <p className="text-gray-600 text-sm mb-4 h-12 overflow-hidden">{product.description}</p>
        {/* ... */}
      </div>
    </div>
  );
}
```

### JSX

React では JSX という JavaScript の拡張構文を使用します。HTML のように見えますが、実際には JavaScript の中で HTML タグのような構文を使えるようにしたものです。

例：
```jsx
// これは JSX です
const element = <h1>Hello, world!</h1>;

// コンポーネント内では、JSX を使って UI を記述します
return (
  <div className="container">
    <h1 className="title">{title}</h1>
    <p>{description}</p>
  </div>
);
```

### Props

Props（プロパティの略）は、親コンポーネントから子コンポーネントにデータを渡す方法です。

例: `ProductCard` コンポーネントは `product` という props を受け取ります：

```jsx
<ProductCard product={productData} />
```

```jsx
// ProductCard.tsx 内で props を受け取る
export default function ProductCard({ product }: { product: Product }) {
  // product オブジェクトのプロパティを使用
  return (
    <div>
      <h2>{product.name}</h2>
      <p>{product.description}</p>
    </div>
  );
}
```

### Hooks

Hooks は React の関数コンポーネントで状態管理やライフサイクル機能を使用するための API です。

例: `useState` は状態を管理するための Hook です：

```jsx
import { useState } from 'react';

export default function ProductCard({ product }) {
  // 画像の読み込み状態を管理するステート
  const [imageError, setImageError] = useState(false);
  
  // ...
  
  return (
    // imageError の値に基づいて UI を変更
    {imageError ? <div>画像読み込みエラー</div> : <img src={product.image_url} />}
  );
}
```

## 3. Next.js フレームワークの概要

Next.js は React をベースにした、フルスタックの Web アプリケーションフレームワークです。以下の特徴があります：

### 主要機能

1. **サーバーサイドレンダリング (SSR)**: ページをサーバー側でレンダリングし、完全な HTML として送信
2. **静的サイト生成 (SSG)**: ビルド時にページを事前生成して静的ファイルとして配信
3. **ファイルベースのルーティング**: ディレクトリ構造に基づいた自動的なルーティング
4. **API ルート**: サーバーサイドのエンドポイントを簡単に作成可能
5. **画像最適化**: 自動的な画像最適化機能

### App Router

このプロジェクトでは Next.js の最新機能である App Router を使用しています。これは `app` ディレクトリに基づいたルーティングシステムで、以下の特徴があります：

- **ネストされたレイアウト**: 共通のレイアウトを親コンポーネントとして定義
- **サーバーコンポーネント**: デフォルトでサーバーサイドレンダリングを行うコンポーネント
- **クライアントコンポーネント**: 'use client' ディレクティブを使用してクライアントサイドで実行されるコンポーネント
- **ローディング UI**: 非同期コンポーネントのローディング状態を表示

例: App Router での基本的なファイル構造

```
app/
├── layout.tsx          # 全ページで共有されるレイアウト
├── page.tsx            # ルートページ (/)
└── products/           # /products ルート
    ├── page.tsx        # 商品一覧ページ
    └── [id]/           # 動的ルート
        └── page.tsx    # 商品詳細ページ
```

## 4. プロジェクト構造の理解

添付されたコードのプロジェクト構造を見ていきましょう。

### 主要ディレクトリとファイル

- **app/**: Next.js の App Router で使用されるルートディレクトリ
  - **layout.tsx**: 全ページで共有されるレイアウト
  - **page.tsx**: ホームページ
  - **products/**: 商品関連のページ
    - **page.tsx**: 商品一覧ページ
    - **ProductsList.tsx**: 商品一覧を表示するコンポーネント
    - **components/**: 商品一覧ページで使用される小さなコンポーネント
- **components/**: 再利用可能な UI コンポーネント
- **lib/**: ユーティリティ関数や API クライアント
  - **api/**: バックエンド API と通信するための関数
  - **logger/**: ログ処理のためのユーティリティ

### 設定ファイル

- **package.json**: プロジェクトの依存関係とスクリプトの定義
- **tsconfig.json**: TypeScript の設定
- **next.config.ts**: Next.js の設定

## 5. 主要コンポーネントとその役割

添付されたコードには、いくつかの主要なコンポーネントがあります。それぞれの役割と実装を見ていきましょう。

### layout.tsx

このファイルは全ページで共有されるレイアウトを定義しています。ヘッダーとフッターを含み、`children` prop を通じてページ固有のコンテンツを表示します。

重要な部分：
```jsx
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body className={inter.className}>
        <header>...</header>
        <main>{children}</main>
        <footer>...</footer>
      </body>
    </html>
  );
}
```

### page.tsx (ホームページ)

ウェブサイトのトップページを定義しています。シンプルな紹介文と商品一覧へのリンクを表示しています。

### products/page.tsx

商品一覧ページを定義しています。URL クエリパラメータを解析し、カテゴリーフィルターと商品リストを表示します。

重要な部分：
```jsx
export default async function ProductsPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; category?: string }>;
}) {
  // URLパラメータの処理
  const { page, category } = await searchParams;
  const pageParam = Number(page) || 1;
  const categoryParam = category ? Number(category) : undefined;

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold text-gray-800 mb-6">商品一覧</h1>

      <div className="flex flex-col md:flex-row gap-6">
        {/* カテゴリーフィルター */}
        <div className="w-full md:w-64 flex-shrink-0">
          <Suspense fallback={...}>
            <CategoryFilter selectedCategoryId={categoryParam} />
          </Suspense>
        </div>

        {/* 商品リスト */}
        <div className="flex-grow">
          <Suspense fallback={...}>
            <ProductsList page={pageParam} pageSize={9} categoryId={categoryParam} />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
```

### ProductCard.tsx

個々の商品カードを表示するコンポーネントです。商品の画像、名前、説明、価格を表示し、詳細ページへのリンクを提供します。

特徴：
- `'use client'` ディレクティブによるクライアントコンポーネントの宣言
- `useState` Hook を使用した画像エラー状態の管理
- 商品データの表示とスタイリング

### CategoryFilter.tsx

カテゴリーによる商品フィルタリングを提供するコンポーネントです。選択されたカテゴリーに基づいて UI を変更します。

### Pagination.tsx

ページネーションを提供するコンポーネントです。前後のページへのナビゲーションを可能にします。

## 6. クライアントサイドとサーバーサイドのレンダリング

Next.js では、コンポーネントをサーバーサイドとクライアントサイドのどちらで実行するかを選択できます。

### サーバーコンポーネント

App Router では、デフォルトですべてのコンポーネントがサーバーコンポーネントとして扱われます。

特徴：
- サーバー上でレンダリングされ、HTML としてクライアントに送信される
- JavaScript バンドルサイズを削減できる
- データベースやファイルシステムに直接アクセスできる
- ユーザーのブラウザで実行されないため、セキュリティが向上する

例：`ProductsList.tsx` はサーバーコンポーネントで、サーバーサイドでデータを取得しています：

```jsx
export default async function ProductsList({
  page,
  pageSize,
  categoryId,
}: {
  page: number;
  pageSize: number;
  categoryId?: number;
}) {
  // サーバーサイドでデータを取得
  try {
    const { items: products, total_pages: totalPages } = await getProducts({
      page,
      page_size: pageSize,
      category_id: categoryId,
    });
    
    // ...
  }
}
```

### クライアントコンポーネント

インタラクティブな機能が必要なコンポーネントは、ファイルの先頭に `'use client'` ディレクティブを追加することでクライアントコンポーネントとして宣言できます。

特徴：
- ブラウザ上で実行される
- イベントリスナーやステート、エフェクトなどの React の機能を使用できる
- ユーザーのインタラクションに応答できる

例：`ProductCard.tsx` はクライアントコンポーネントで、ユーザーインタラクションを処理します：

```jsx
'use client';

import { useState } from 'react';

export default function ProductCard({ product }: { product: Product }) {
  // クライアントサイドでの状態管理
  const [imageError, setImageError] = useState(false);
  
  // ...
  
  return (
    // イベントハンドラーを使用
    <Image
      src={product.image_url}
      alt={product.name}
      onError={() => {
        console.warn('Product image failed to load');
        setImageError(true);
      }}
    />
  );
}
```

### Suspense

Next.js は React の Suspense を使用して、非同期コンポーネントのローディング状態を管理します。

例：
```jsx
<Suspense fallback={<LoadingIndicator />}>
  <ProductsList page={pageParam} pageSize={9} categoryId={categoryParam} />
</Suspense>
```

この例では、`ProductsList` がデータをロードしている間、`LoadingIndicator` が表示されます。

## 7. TailwindCSS によるスタイリング

TailwindCSS は、ユーティリティファーストの CSS フレームワークです。HTML 要素に直接クラス名を適用することでスタイリングを行います。

### 基本的な使い方

```jsx
<div className="bg-white rounded-lg p-4 shadow-md">
  <h2 className="text-xl font-bold text-blue-600">タイトル</h2>
  <p className="text-gray-700 mt-2">説明文</p>
</div>
```

この例では：
- `bg-white`: 背景色を白に設定
- `rounded-lg`: 角を丸くする
- `p-4`: パディングを設定
- `shadow-md`: 中程度の影を追加
- `text-xl`: テキストサイズを大きく設定
- `font-bold`: フォントを太字に設定
- `text-blue-600`: テキスト色を青に設定
- `mt-2`: 上マージンを設定

### レスポンシブデザイン

Tailwind では、ブレークポイント接頭辞を使用してレスポンシブデザインを実装できます：

```jsx
<div className="flex flex-col md:flex-row gap-6">
  <div className="w-full md:w-64 flex-shrink-0">
    {/* ... */}
  </div>
</div>
```

この例では：
- モバイル画面では `flex-col` で縦並びになる
- 中サイズ（`md:`）以上の画面では `flex-row` で横並びになる
- 小さい画面では幅が 100%（`w-full`）
- 中サイズ以上の画面では幅が固定（`md:w-64`）

### コンポーネントの例

`ProductCard.tsx` の例：

```jsx
<div className="bg-white rounded-lg overflow-hidden shadow-lg hover:shadow-xl transition-shadow duration-300">
  <div className="h-48 bg-gray-200 flex items-center justify-center relative">
    {/* 画像コンテンツ */}
  </div>
  <div className="p-4">
    <h2 className="text-xl font-semibold text-gray-800 mb-2">{product.name}</h2>
    <p className="text-gray-600 text-sm mb-4 h-12 overflow-hidden">{product.description}</p>
    <div className="flex justify-between items-center">
      <span className="text-xl font-bold text-indigo-600">
        ¥{product.price.toLocaleString()}
      </span>
      <Link
        href={`/products/${product.id}`}
        className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-300"
      >
        詳細を見る
      </Link>
    </div>
  </div>
</div>
```

このコンポーネントでは、Tailwind のクラスを使用して：
1. カードのレイアウトとスタイリング
2. 画像コンテナのサイズと位置
3. テキストのスタイリングとレイアウト
4. ボタンのスタイリングとホバーエフェクト

を実装しています。

## 8. React でのステート管理

React でのステート（状態）管理は、UI の動的な部分を制御するために重要です。

### useState フック

最も基本的なステート管理方法は `useState` フックを使用することです：

```jsx
import { useState } from 'react';

function Counter() {
  // [現在の状態, 状態を更新する関数] = useState(初期値)
  const [count, setCount] = useState(0);
  
  return (
    <div>
      <p>カウント: {count}</p>
      <button onClick={() => setCount(count + 1)}>増加</button>
      <button onClick={() => setCount(count - 1)}>減少</button>
    </div>
  );
}
```

### ProductCard での例

`ProductCard.tsx` では、画像の読み込みエラーを管理するためにステートを使用しています：

```jsx
'use client';

import { useState } from 'react';

export default function ProductCard({ product }: { product: Product }) {
  // 画像の読み込み状態を管理するステート
  const [imageError, setImageError] = useState(false);
  
  return (
    <div>
      {product.image_url && !imageError ? (
        <Image
          src={product.image_url}
          alt={product.name}
          onError={() => {
            console.warn('Product image failed to load');
            setImageError(true);
          }}
        />
      ) : (
        // エラー時の代替表示
        <div className="flex flex-col items-center justify-center text-gray-500">
          <svg>...</svg>
          <p>商品画像なし</p>
        </div>
      )}
    </div>
  );
}
```

このコードでは：
1. `imageError` ステートの初期値を `false` に設定
2. 画像の読み込みエラーが発生すると `onError` イベントハンドラーが呼び出される
3. `setImageError(true)` でステートを更新
4. ステートの変更によって、コンポーネントが再レンダリングされ、エラー表示に切り替わる

### より複雑なステート管理

大規模なアプリケーションでは、より高度なステート管理ソリューションが必要になることがあります：

- **Context API**: コンポーネントツリー全体でデータを共有
- **Reducers**: 複雑なステートロジックを管理
- **Redux**: 大規模アプリケーション向けの状態管理ライブラリ
- **Zustand**: シンプルで軽量な状態管理ライブラリ

このプロジェクトでは、コンポーネントごとに閉じたシンプルなステート管理が行われていますが、機能が拡張されるにつれて、グローバルなステート管理（例えばショッピングカートの状態）が必要になるかもしれません。

## 9. Next.js でのナビゲーションとルーティング

Next.js のルーティングシステムは、ファイルシステムに基づいています。App Router では、`app` ディレクトリの構造がそのままルートになります。

### 基本的なルーティング

- `app/page.tsx` → `/` (ホームページ)
- `app/products/page.tsx` → `/products` (商品一覧ページ)
- `app/products/[id]/page.tsx` → `/products/1`, `/products/2` など (動的ルート)

### Link コンポーネント

Next.js は、クライアントサイドのナビゲーションのための `Link` コンポーネントを提供します：

```jsx
import Link from 'next/link';

// 基本的な使い方
<Link href="/products">商品一覧</Link>

// 動的なルートへのリンク
<Link href={`/products/${product.id}`}>詳細を見る</Link>

// クエリパラメータを含むリンク
<Link href="/products?category=1">カテゴリー1</Link>
```

`Link` コンポーネントは、通常の `<a>` タグを拡張し、クライアントサイドのナビゲーションを提供します。これにより、ページ全体をリロードせずに高速なナビゲーションが可能になります。

### 動的ルートのパラメータ取得

動的ルート（`app/products/[id]/page.tsx` など）では、ルートパラメータを取得できます：

```jsx
// app/products/[id]/page.tsx
export default function ProductDetail({ params }: { params: { id: string } }) {
  // params.id でルートパラメータを取得
  return <div>商品ID: {params.id}</div>;
}
```

### クエリパラメータの取得

`searchParams` プロパティを使用して、URL のクエリパラメータにアクセスできます：

```jsx
// app/products/page.tsx
export default function ProductsPage({
  searchParams,
}: {
  searchParams: { page?: string; category?: string };
}) {
  // URLパラメータの処理
  const pageParam = searchParams.page ? parseInt(searchParams.page) : 1;
  const categoryParam = searchParams.category
    ? parseInt(searchParams.category)
    : undefined;
  
  // ...
}
```

### ナビゲーションの例

プロジェクトでのナビゲーション例：

1. ヘッダーでの基本的なナビゲーション（`layout.tsx`）：
```jsx
<nav className="ml-6 flex items-center space-x-4">
  <Link
    href="/products"
    className="text-gray-700 hover:text-indigo-600 px-3 py-2 rounded-md text-sm font-medium"
  >
    商品一覧
  </Link>
  <Link
    href="/products?category=1"
    className="text-gray-700 hover:text-indigo-600 px-3 py-2 rounded-md text-sm font-medium"
  >
    カテゴリー
  </Link>
</nav>
```

2. 商品カードから詳細ページへ（`ProductCard.tsx`）：
```jsx
<Link
  href={`/products/${product.id}`}
  className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-300"
>
  詳細を見る
</Link>
```

3. ページネーション（`Pagination.tsx`）：
```jsx
<Link
  href={`/products?page=${Math.max(1, currentPage - 1)}`}
  className="px-4 py-2 rounded-lg bg-gray-200 text-gray-700 hover:bg-gray-300"
>
  前へ
</Link>
```

## 10. API リクエストの処理

フロントエンドからバックエンドへのデータ取得は、Web アプリケーションの重要な部分です。プロジェクトでは、`lib/api/products.ts` ファイルで API リクエストを処理しています。

### データ取得の基本

```typescript
export async function getProducts(
  params: GetProductsParams = {}
): Promise<PaginatedResponse<Product>> {
  const startTime = Date.now();
  const queryParams = new URLSearchParams();

  // パラメータの設定
  if (params.page) queryParams.append("page", params.page.toString());
  if (params.page_size) queryParams.append("page_size", params.page_size.toString());
  if (params.category_id) queryParams.append("category_id", params.category_id.toString());

  const url = `${API_BASE_URL}/products?${queryParams.toString()}`;

  try {
    logger.info("Fetching products", {
      params,
      url,
    });

    const response = await fetch(url, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    });

    const data = await handleResponse<PaginatedResponse<Product>>(response);

    const duration = Date.now() - startTime;
    logger.info("Products fetched successfully", {
      duration_ms: duration,
      count: data.items.length,
      total: data.total_items,
      page: data.page,
    });

    return data;
  } catch (error) {
    // エラーハンドリング
    // ...
    throw error;
  }
}
```

このコードでは：
1. URL クエリパラメータを構築
2. `fetch` API を使用してデータを取得
3. レスポンスを JSON に変換
4. エラーハンドリングとロギング
5. タイプセーフな結果を返す

### Next.js でのデータ取得パターン

Next.js はいくつかのデータ取得パターンをサポートしています：

1. **サーバーコンポーネントでのデータ取得**：
```jsx
// サーバーコンポーネント内で直接データを取得
export default async function ProductsList() {
  const products = await getProducts();
  return <div>{products.map(product => <ProductCard product={product} />)}</div>;
}
```

2. **API ルートを使用したデータ取得**（このプロジェクトでは使用されていません）：
```jsx
// app/api/products/route.ts
import { NextResponse } from 'next/server';

export async function GET(request: Request) {
  // バックエンドからデータを取得
  const products = await fetchProductsFromDatabase();
  return NextResponse.json(products);
}
```

3. **クライアントコンポーネントでのデータ取得**（`useEffect` や SWR/React Query などのライブラリを使用）：
```jsx
'use client';

import { useEffect, useState } from 'react';

export function ProductList() {
  const [products, setProducts] = useState([]);
  
  useEffect(() => {
    async function fetchData() {
      const data = await getProducts();
      setProducts(data.items);
    }
    
    fetchData();
  }, []);
  
  return <div>{products.map(product => <div>{product.name}</div>)}</div>;
}
```

このプロジェクトでは、主にサーバーコンポーネントでデータを取得し、必要な場合にのみクライアントコンポーネントを使用しています。これはパフォーマンスとユーザーエクスペリエンスのバランスを取るための最適なアプローチです。

## 11. フロントエンドコンポーネントのテスト

フロントエンドコンポーネントのテストは、アプリケーションの品質を保証するために重要です。プロジェクトでは Jest と React Testing Library を使用しています。

### Jest と React Testing Library

- **Jest**: JavaScript のテストフレームワーク
- **React Testing Library**: React コンポーネントをテストするためのユーティリティ

### テスト例: ProductCard.test.tsx

```jsx
import { Product } from "@/lib/api/products";
import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";

// ProductCardコンポーネントをモック化
jest.mock("next/link", () => {
  return ({ children, href }: { children: React.ReactNode; href: string }) => {
    return <a href={href}>{children}</a>;
  };
});

// ProductCardコンポーネントをインポート
import ProductCard from "@/components/products/ProductCard";

describe("ProductCard", () => {
  const mockProduct: Product = {
    id: 1,
    name: "テスト商品",
    description: "これはテスト用の商品です",
    price: 1000,
    image_url: "https://example.com/test.jpg",
    category_id: 1,
  };

  it("商品名が正しく表示される", () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText("テスト商品")).toBeInTheDocument();
  });

  it("商品説明が正しく表示される", () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText("これはテスト用の商品です")).toBeInTheDocument();
  });

  it("商品価格が正しくフォーマットされて表示される", () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText("¥1,000")).toBeInTheDocument();
  });

  it("商品詳細へのリンクが正しく設定される", () => {
    render(<ProductCard product={mockProduct} />);
    const link = screen.getByText("詳細を見る");
    expect(link).toBeInTheDocument();
    expect(link.closest("a")).toHaveAttribute("href", "/products/1");
  });

  it("画像がない場合は代替テキストが表示される", () => {
    const productWithoutImage = { ...mockProduct, image_url: "" };
    render(<ProductCard product={productWithoutImage} />);
    expect(screen.getByText("商品画像なし")).toBeInTheDocument();
  });
});
```

このテストでは：
1. モックデータを作成
2. コンポーネントをレンダリング
3. 表示される要素を検証
4. さまざまなシナリオをテスト（例: 画像がない場合）

### テスト実行

`package.json` の scripts セクションにテスト用のコマンドが定義されています：

```json
"scripts": {
  "test": "jest",
  "test:watch": "jest --watch"
}
```

テストを実行するには：
```bash
npm test
```

または、watch モードで実行（ファイル変更時に自動的にテストを再実行）：
```bash
npm run test:watch
```

### テストの利点

- **バグの早期発見**: 変更によって既存の機能が壊れていないことを確認
- **リファクタリングの安全性**: コードを改善しても機能が維持されていることを確認
- **コードの品質向上**: テスト可能なコードは一般的に品質が高い
- **開発者の信頼性向上**: コードが意図したとおりに動作することを確認

## 12. 顧客向けと管理者向け画面の違い

このプロジェクトでは、顧客向け (`frontend-customer`) と管理者向け (`frontend-admin`) の2つのフロントエンドが存在します。それぞれの違いを見ていきましょう。

### 顧客向け (frontend-customer)

顧客向けフロントエンドは、エンドユーザー（消費者）が商品を閲覧して購入するために使用します。

特徴：
- シンプルで使いやすいインターフェース
- 商品の閲覧と検索に焦点
- 商品詳細の表示
- カート機能とチェックアウトプロセス

例（商品カード）：
```jsx
<div className="p-4">
  <h2 className="text-xl font-semibold text-gray-800 mb-2">{product.name}</h2>
  <p className="text-gray-600 text-sm mb-4 h-12 overflow-hidden">{product.description}</p>
  <div className="flex justify-between items-center">
    <span className="text-xl font-bold text-indigo-600">
      ¥{product.price.toLocaleString()}
    </span>
    <Link
      href={`/products/${product.id}`}
      className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-300"
    >
      詳細を見る
    </Link>
  </div>
</div>
```

### 管理者向け (frontend-admin)

管理者向けフロントエンドは、ストア管理者が商品、在庫、注文などを管理するために使用します。

特徴：
- より詳細な情報の表示
- 商品の追加、編集、削除機能
- 在庫管理
- 注文の確認と処理
- 分析およびレポート機能

例（管理者用商品カード）：
```jsx
<div className="p-4">
  <h2 className="text-xl font-semibold text-gray-800 mb-2">{product.name}</h2>
  <p className="text-gray-600 text-sm mb-2 h-8 overflow-hidden">{product.description}</p>
  <div className="flex flex-col space-y-2 mb-2">
    <div className="flex justify-between items-center">
      <span className="text-lg font-bold text-indigo-600">
        ¥{product.price.toLocaleString()}
      </span>
      <span className="text-sm bg-blue-100 text-blue-800 px-2 py-1 rounded">
        カテゴリー: {product.category_id}
      </span>
    </div>
    <div className="flex justify-between items-center">
      <span className="text-sm bg-green-100 text-green-800 px-2 py-1 rounded">
        在庫状況: 在庫あり
      </span>
      <span className="text-sm bg-purple-100 text-purple-800 px-2 py-1 rounded">
        ID: {product.id}
      </span>
    </div>
  </div>
  <div className="flex justify-between mt-4">
    <Link
      href={`/products/${product.id}/edit`}
      className="bg-yellow-500 hover:bg-yellow-600 text-white px-3 py-1.5 rounded text-sm font-medium transition-colors duration-300"
    >
      編集
    </Link>
    <Link
      href={`/products/${product.id}`}
      className="bg-indigo-600 hover:bg-indigo-700 text-white px-3 py-1.5 rounded text-sm font-medium transition-colors duration-300"
    >
      詳細
    </Link>
    <button
      className="bg-red-500 hover:bg-red-600 text-white px-3 py-1.5 rounded text-sm font-medium transition-colors duration-300"
      onClick={() => alert(`商品ID ${product.id} を削除しますか？`)}
    >
      削除
    </button>
  </div>
</div>
```

### 主な違い

1. **表示情報**:
   - 顧客向け: 消費者に必要な情報（商品名、説明、価格）
   - 管理者向け: より詳細な情報（ID、在庫状況、カテゴリーID）

2. **操作**:
   - 顧客向け: 閲覧、カートに追加
   - 管理者向け: 追加、編集、削除

3. **UI デザイン**:
   - 顧客向け: 見栄えや使いやすさを重視
   - 管理者向け: 機能性と効率性を重視

4. **ナビゲーション**:
   - 顧客向け: 商品カテゴリーや特集
   - 管理者向け: 管理メニュー（商品管理、注文管理、在庫管理など）

5. **認証・認可**:
   - 顧客向け: ユーザー登録とログイン（オプション）
   - 管理者向け: 管理者認証と権限管理

### 共通点

両方のフロントエンドは同じバックエンド API を使用し、同じデータベースにアクセスします。また、技術的には同じフレームワーク（Next.js）と構成を使用していますが、目的とユーザーエクスペリエンスが異なります。

## 13. 次のステップと学習リソース

Next.js と React の学習を続けるためのリソースとステップを紹介します。

### 学習のステップ

1. **React の基礎を固める**
   - コンポーネント、props、state の概念を理解する
   - イベント処理とライフサイクルメソッドを学ぶ
   - Hooks の使い方を理解する

2. **Next.js を深く学ぶ**
   - ルーティングシステムを理解する
   - データ取得と SSR/SSG の違いを理解する
   - API ルートの作成と使用方法を学ぶ

3. **実践的なプロジェクトに取り組む**
   - 小さな機能を追加して実装してみる
   - エラー処理や最適化を含む本格的な機能を構築する
   - オブザーバビリティとパフォーマンスの改善に取り組む

### 推奨学習リソース

1. **公式ドキュメント**
   - [React 公式ドキュメント](https://ja.react.dev/)
   - [Next.js 公式ドキュメント](https://nextjs.org/docs)
   - [TypeScript 公式ドキュメント](https://www.typescriptlang.org/docs/)
   - [TailwindCSS 公式ドキュメント](https://tailwindcss.com/docs)

2. **チュートリアルと学習サイト**
   - [MDN Web Docs](https://developer.mozilla.org/ja/docs/Learn)
   - [React Tutorial](https://react-tutorial.app/)
   - [Next.js 入門](https://nextjs.org/learn)
   - [TypeScript Deep Dive](https://typescript-jp.gitbook.io/deep-dive/)

3. **書籍**
   - 「React ハンズオンラーニング」
   - 「実践 TypeScript」
   - 「Next.js によるフルスタックアプリケーション開発」

4. **動画講座**
   - [Udemy の React/Next.js コース](https://www.udemy.com/topic/react/)
   - [YouTube の Next.js チュートリアル](https://www.youtube.com/results?search_query=next.js+tutorial)

### 実践的な学習ステップ

1. **コードの読解を深める**
   - 既存のコンポーネントの動作を理解する
   - コンポーネント間のデータフローを追跡する
   - ファイル間の依存関係を把握する

2. **小さな機能追加から始める**
   - 商品の並び替え機能を追加する
   - 商品検索ボックスを実装する
   - お気に入り機能を追加する

3. **応用的な機能に挑戦**
   - ユーザー認証の実装
   - ショッピングカートの状態管理
   - 在庫リアルタイム更新の実装

4. **テストと最適化**
   - コンポーネントのテストを書く
   - パフォーマンス最適化を行う
   - アクセシビリティを改善する

フロントエンド開発はスタックも進化も早いですが、基本概念はそれほど変わりません。これらの基礎を固めることで、新しいフレームワークや技術にも適応しやすくなります。
