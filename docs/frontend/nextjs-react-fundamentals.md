# 1. Next.js & React フロントエンド基本ガイド

## 1.1. 目次

- [1. Next.js \& React フロントエンド基本ガイド](#1-nextjs--react-フロントエンド基本ガイド)
  - [1.1. 目次](#11-目次)
  - [1.2. フロントエンド開発の基礎](#12-フロントエンド開発の基礎)
    - [1.2.1. 主要技術](#121-主要技術)
    - [1.2.2. モダンフロントエンド開発の特徴](#122-モダンフロントエンド開発の特徴)
  - [1.3. Reactの基本概念](#13-reactの基本概念)
    - [1.3.1. コンポーネント](#131-コンポーネント)
    - [1.3.2. JSX](#132-jsx)
    - [1.3.3. Props](#133-props)
    - [1.3.4. Hooks](#134-hooks)
  - [1.4. Next.jsフレームワークの概要](#14-nextjsフレームワークの概要)
    - [1.4.1. 主要機能](#141-主要機能)
    - [1.4.2. App Router](#142-app-router)
  - [1.5. 基本的なプロジェクト構造](#15-基本的なプロジェクト構造)
    - [1.5.1. 主要ディレクトリとファイル](#151-主要ディレクトリとファイル)
    - [1.5.2. 設定ファイル](#152-設定ファイル)
  - [1.6. データ取得と状態管理](#16-データ取得と状態管理)
    - [1.6.1. サーバーサイドでのデータフェッチ](#161-サーバーサイドでのデータフェッチ)
    - [1.6.2. クライアントサイドの状態管理](#162-クライアントサイドの状態管理)
  - [1.7. スタイリングの基本](#17-スタイリングの基本)
    - [1.7.1. TailwindCSSの基本](#171-tailwindcssの基本)
    - [1.7.2. レスポンシブデザイン](#172-レスポンシブデザイン)
  - [1.8. API通信の基本](#18-api通信の基本)
    - [1.8.1. 基本的なデータ取得](#181-基本的なデータ取得)
    - [1.8.2. Next.jsでのデータ取得パターン](#182-nextjsでのデータ取得パターン)
  - [1.9. 基本的なエラーハンドリング](#19-基本的なエラーハンドリング)
    - [1.9.1. try-catchを使用したエラーハンドリング](#191-try-catchを使用したエラーハンドリング)
    - [1.9.2. コンポーネントでのエラー表示](#192-コンポーネントでのエラー表示)
    - [1.9.3. Next.jsのエラーページ](#193-nextjsのエラーページ)
  - [1.10. 基本的なログ実装](#110-基本的なログ実装)
    - [1.10.1. コンソールログの適切な使用](#1101-コンソールログの適切な使用)
    - [1.10.2. 開発環境と本番環境の分離](#1102-開発環境と本番環境の分離)
    - [1.10.3. パフォーマンスのログ記録](#1103-パフォーマンスのログ記録)
  - [1.11. テストの基本](#111-テストの基本)
    - [1.11.1. Jestのセットアップ](#1111-jestのセットアップ)
    - [1.11.2. コンポーネントのテスト](#1112-コンポーネントのテスト)
    - [1.11.3. APIのモックとテスト](#1113-apiのモックとテスト)
  - [1.12. 推奨学習リソース](#112-推奨学習リソース)

## 1.2. フロントエンド開発の基礎

フロントエンド開発とは、ウェブサイトやウェブアプリケーションのユーザーインターフェース（UI）を構築するプロセスのことです。ユーザーが直接操作する部分を担当し、見た目と操作性に関わるすべての要素を含みます。

### 1.2.1. 主要技術

- **HTML**: コンテンツ構造を定義
- **CSS**: 見た目やレイアウトを制御
- **JavaScript**: インタラクティブな動作を実装

### 1.2.2. モダンフロントエンド開発の特徴

今回のプロジェクトで使用されているのは、モダンなフロントエンド開発のアプローチです：

1. **コンポーネントベース開発**: UI を再利用可能な独立したパーツに分割
2. **宣言的なプログラミング**: 「何を」表示したいかを定義し、「どのように」表示するかはフレームワークに任せる
3. **仮想 DOM**: 実際の DOM 操作を最小限に抑えるために、仮想的な DOM を使用してパフォーマンスを向上
4. **単一方向データフロー**: データは親から子へと一方向に流れる設計

## 1.3. Reactの基本概念

React は Facebook（現 Meta）によって開発された JavaScript ライブラリで、ユーザーインターフェースを構築するために使用されます。

### 1.3.1. コンポーネント

React の最も基本的な概念は「コンポーネント」です。コンポーネントは独立した再利用可能な UI の部品です。

例: `ProductCard.tsx` は商品カードを表示するコンポーネントです：

```jsx
export default function ProductCard({ product }: { product: Product }) {
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

### 1.3.2. JSX

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

### 1.3.3. Props

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

### 1.3.4. Hooks

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

## 1.4. Next.jsフレームワークの概要

Next.js は React をベースにした、フルスタックの Web アプリケーションフレームワークです。以下の特徴があります：

### 1.4.1. 主要機能

1. **サーバーサイドレンダリング (SSR)**: ページをサーバー側でレンダリングし、完全な HTML として送信
2. **静的サイト生成 (SSG)**: ビルド時にページを事前生成して静的ファイルとして配信
3. **ファイルベースのルーティング**: ディレクトリ構造に基づいた自動的なルーティング
4. **API ルート**: サーバーサイドのエンドポイントを簡単に作成可能
5. **画像最適化**: 自動的な画像最適化機能

### 1.4.2. App Router

このプロジェクトでは Next.js の最新機能である App Router を使用しています。これは `app` ディレクトリに基づいたルーティングシステムで、以下の特徴があります：

- **ネストされたレイアウト**: 共通のレイアウトを親コンポーネントとして定義
- **サーバーコンポーネント**: デフォルトでサーバーサイドレンダリングを行うコンポーネント
- **クライアントコンポーネント**: 'use client' ディレクティブを使用してクライアントサイドで実行されるコンポーネント
- **ローディング UI**: 非同期コンポーネントのローディング状態を表示

例: App Router での基本的なファイル構造

```text
app/
├── layout.tsx          # 全ページで共有されるレイアウト
├── page.tsx            # ルートページ (/)
└── products/           # /products ルート
    ├── page.tsx        # 商品一覧ページ
    └── [id]/           # 動的ルート
        └── page.tsx    # 商品詳細ページ
```

## 1.5. 基本的なプロジェクト構造

Next.jsプロジェクトの基本的な構造について見ていきましょう。

### 1.5.1. 主要ディレクトリとファイル

- **app/**: Next.js の App Router で使用されるルートディレクトリ
  - **layout.tsx**: 全ページで共有されるレイアウト
  - **page.tsx**: ホームページ
  - **products/**: 商品関連のページ
    - **page.tsx**: 商品一覧ページ
    - **[id]/page.tsx**: 商品詳細ページ (動的ルート)
- **components/**: 再利用可能な UI コンポーネント
  - **ui/**: 基本的なUIコンポーネント（ボタン、カードなど）
  - **layout/**: レイアウト関連のコンポーネント（ヘッダー、フッターなど）
- **lib/**: ユーティリティ関数や API クライアント
  - **api/**: バックエンド API と通信するための関数
  - **utils/**: ヘルパー関数
- **public/**: 静的ファイル（画像、フォントなど）

### 1.5.2. 設定ファイル

- **package.json**: プロジェクトの依存関係とスクリプトの定義
- **tsconfig.json**: TypeScript の設定
- **next.config.js**: Next.js の設定

## 1.6. データ取得と状態管理

### 1.6.1. サーバーサイドでのデータフェッチ

Next.js のサーバーコンポーネントを使用すると、サーバー側でデータを取得できます：

```jsx
// app/products/page.tsx
export default async function ProductsPage() {
  // サーバーサイドでデータ取得
  const products = await fetch('https://api.example.com/products').then(res => res.json());

  return (
    <div>
      <h1>商品一覧</h1>
      <div className="products-grid">
        {products.map(product => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>
    </div>
  );
}
```

### 1.6.2. クライアントサイドの状態管理

クライアントコンポーネントでは、React の Hooks を使用して状態を管理できます：

```jsx
'use client';

import { useState } from 'react';

export default function FilterablProductList({ products }) {
  const [filter, setFilter] = useState('');

  // フィルター適用
  const filteredProducts = products.filter(product =>
    product.name.toLowerCase().includes(filter.toLowerCase())
  );

  return (
    <div>
      <input
        type="text"
        placeholder="商品を検索..."
        value={filter}
        onChange={(e) => setFilter(e.target.value)}
      />
      <div className="products-grid">
        {filteredProducts.map(product => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>
    </div>
  );
}
```

## 1.7. スタイリングの基本

Next.js ではいくつかのスタイリング方法をサポートしています。本プロジェクトでは TailwindCSS を使用します。

### 1.7.1. TailwindCSSの基本

TailwindCSSは、ユーティリティファーストのCSSフレームワークです。HTMLクラス名を通じてスタイリングを行います。

```jsx
<div className="bg-white rounded-lg shadow-md p-4 hover:shadow-lg transition-shadow">
  <h2 className="text-xl font-bold text-gray-800 mb-2">{product.name}</h2>
  <p className="text-gray-600 mb-4">{product.description}</p>
  <span className="text-blue-600 font-semibold">¥{product.price}</span>
</div>
```

上記の例では：

- `bg-white`: 背景色を白に設定
- `rounded-lg`: 角を丸くする
- `shadow-md`: 中程度の影を追加
- `p-4`: パディングを追加
- `hover:shadow-lg`: ホバー時の影を大きくする
- `text-xl`: テキストサイズを大きくする
- `font-bold`: フォントを太字にする
- `text-gray-800`: テキスト色を暗いグレーに設定

### 1.7.2. レスポンシブデザイン

Tailwindでは、ブレークポイント接頭辞を使用してレスポンシブデザインを実装できます：

```jsx
<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
  {/* 商品カード */}
</div>
```

この例では：

- 小さい画面では 1カラム
- 中サイズ（md）以上では 2カラム
- 大サイズ（lg）以上では 3カラム

## 1.8. API通信の基本

フロントエンドからバックエンドへの通信は、Web アプリケーションの重要な部分です。

### 1.8.1. 基本的なデータ取得

```javascript
// lib/api/products.js
export async function getProducts(params = {}) {
  // クエリパラメータの構築
  const queryParams = new URLSearchParams();
  if (params.page) queryParams.append("page", params.page.toString());
  if (params.pageSize) queryParams.append("page_size", params.pageSize.toString());

  const url = `${process.env.NEXT_PUBLIC_API_URL}/products?${queryParams.toString()}`;

  try {
    const response = await fetch(url, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`API error: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Failed to fetch products:", error);
    throw error;
  }
}
```

### 1.8.2. Next.jsでのデータ取得パターン

1. **サーバーコンポーネントでのデータ取得**:

    ```jsx
    // app/products/page.jsx
    import { getProducts } from '@/lib/api/products';

    export default async function ProductsPage() {
      // サーバーサイドでデータ取得
      const products = await getProducts();
      return (
        <div>
          {products.map(product => (
            <ProductCard key={product.id} product={product} />
          ))}
        </div>
      );
    }
    ```

2. **クライアントコンポーネントでのデータ取得**:

    ```jsx
    'use client';

    import { useState, useEffect } from 'react';
    import { getProducts } from '@/lib/api/products';

    export function ProductList() {
      const [products, setProducts] = useState([]);
      const [loading, setLoading] = useState(true);
      const [error, setError] = useState(null);

      useEffect(() => {
        async function fetchData() {
          try {
            setLoading(true);
            const data = await getProducts();
            setProducts(data.items);
          } catch (err) {
            setError(err.message);
          } finally {
            setLoading(false);
          }
        }

        fetchData();
      }, []);

      if (loading) return <div>読み込み中...</div>;
      if (error) return <div>エラー: {error}</div>;

      return <div>{products.map(product => <div key={product.id}>{product.name}</div>)}</div>;
    }
    ```

## 1.9. 基本的なエラーハンドリング

エラーハンドリングは、堅牢なアプリケーションを構築するために重要です。

### 1.9.1. try-catchを使用したエラーハンドリング

```javascript
async function fetchData() {
  try {
    const response = await fetch('/api/data');
    if (!response.ok) {
      throw new Error(`サーバーエラー: ${response.status}`);
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('データ取得エラー:', error);
    throw error; // 呼び出し元でも処理できるように再スロー
  }
}
```

### 1.9.2. コンポーネントでのエラー表示

```jsx
'use client';

import { useState, useEffect } from 'react';

export default function ProductPage({ productId }) {
  const [product, setProduct] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    async function loadProduct() {
      try {
        setLoading(true);
        setError(null);
        const data = await fetchProduct(productId);
        setProduct(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    }

    loadProduct();
  }, [productId]);

  if (loading) return <div className="loading">読み込み中...</div>;
  if (error) return <div className="error-message">エラーが発生しました: {error}</div>;
  if (!product) return <div className="not-found">商品が見つかりませんでした</div>;

  return (
    <div className="product-detail">
      <h1>{product.name}</h1>
      {/* 商品詳細 */}
    </div>
  );
}
```

### 1.9.3. Next.jsのエラーページ

アプリケーション全体で使用できるエラーページを作成できます：

```jsx
// app/error.jsx
'use client';

import { useEffect } from 'react';

export default function Error({ error, reset }) {
  useEffect(() => {
    // エラーをログに記録
    console.error('ページエラー:', error);
  }, [error]);

  return (
    <div className="error-container">
      <h2>エラーが発生しました</h2>
      <button onClick={() => reset()}>再試行</button>
    </div>
  );
}
```

## 1.10. 基本的なログ実装

ログ記録は、アプリケーションの動作を理解し、問題を診断するのに役立ちます。

### 1.10.1. コンソールログの適切な使用

```javascript
// 情報ログ
console.log('商品が読み込まれました', { count: products.length });

// 警告ログ
console.warn('非推奨のAPIが使用されています');

// エラーログ
console.error('APIリクエスト失敗', error);

// グループ化したログ
console.group('商品処理');
console.log('フィルタリング開始');
// 処理
console.log('フィルタリング完了');
console.groupEnd();
```

### 1.10.2. 開発環境と本番環境の分離

```javascript
// lib/logging.js
const isDevelopment = process.env.NODE_ENV === 'development';

export const logger = {
  log: (message, data) => {
    if (isDevelopment) {
      console.log(message, data);
    }
  },
  warn: (message, data) => {
    if (isDevelopment) {
      console.warn(message, data);
    }
  },
  error: (message, error) => {
    // エラーは本番環境でも記録
    console.error(message, error);

    // 本番環境では、エラー監視サービスに送信することも検討
    if (!isDevelopment) {
      // 例: エラー監視サービスに送信するコード
    }
  }
};
```

### 1.10.3. パフォーマンスのログ記録

```javascript
// パフォーマンス計測
function measurePerformance(operationName, operation) {
  const start = performance.now();
  const result = operation();
  const end = performance.now();

  logger.log(`${operationName} 所要時間: ${end - start}ms`);

  return result;
}

// 使用例
const filteredProducts = measurePerformance('商品フィルタリング', () => {
  return products.filter(product => product.price > 1000);
});
```

## 1.11. テストの基本

テストは、アプリケーションの信頼性を確保するために重要です。

### 1.11.1. Jestのセットアップ

```javascript
// jest.config.js
const nextJest = require('next/jest');

const createJestConfig = nextJest({
  dir: './',
});

const customJestConfig = {
  testEnvironment: 'jest-environment-jsdom',
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  moduleNameMapper: {
    '^@/(.*)$': '<rootDir>/$1',
  },
};

module.exports = createJestConfig(customJestConfig);
```

```javascript
// jest.setup.js
import '@testing-library/jest-dom';

// Next.jsのコンポーネントモック
jest.mock('next/image', () => ({
  __esModule: true,
  default: (props) => <img {...props} />,
}));

jest.mock('next/navigation', () => ({
  useRouter() {
    return {
      push: jest.fn(),
      back: jest.fn(),
    };
  },
}));
```

### 1.11.2. コンポーネントのテスト

```javascript
// components/ProductCard.test.js
import { render, screen } from '@testing-library/react';
import ProductCard from './ProductCard';

describe('ProductCard', () => {
  const mockProduct = {
    id: '1',
    name: 'テスト商品',
    description: '商品の説明',
    price: 1000,
  };

  it('商品名が表示されること', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText('テスト商品')).toBeInTheDocument();
  });

  it('商品の説明が表示されること', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText('商品の説明')).toBeInTheDocument();
  });

  it('価格が正しくフォーマットされること', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText('¥1,000')).toBeInTheDocument();
  });
});
```

### 1.11.3. APIのモックとテスト

```javascript
// lib/api/products.test.js
import { getProducts } from './products';

// fetchのモック
global.fetch = jest.fn();

describe('Products API', () => {
  beforeEach(() => {
    jest.resetAllMocks();
  });

  it('getProducts が正しくデータを取得すること', async () => {
    // モックレスポンスの設定
    const mockResponse = {
      ok: true,
      json: jest.fn().mockResolvedValue({
        items: [{ id: '1', name: 'テスト商品' }],
        total: 1
      })
    };

    global.fetch.mockResolvedValue(mockResponse);

    // 関数を実行
    const result = await getProducts();

    // 検証
    expect(global.fetch).toHaveBeenCalledWith(
      expect.stringContaining('/products'),
      expect.any(Object)
    );

    expect(result.items).toHaveLength(1);
    expect(result.items[0].name).toBe('テスト商品');
  });

  it('APIエラー時に例外をスローすること', async () => {
    // エラーレスポンスをモック
    const mockErrorResponse = {
      ok: false,
      status: 500
    };

    global.fetch.mockResolvedValue(mockErrorResponse);

    // 関数の実行と例外の検証
    await expect(getProducts()).rejects.toThrow('API error: 500');
  });
});
```

## 1.12. 推奨学習リソース

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

この基本ガイドを通じて、Next.jsとReactを使用したフロントエンド開発の基礎を理解できるでしょう。実際のプロジェクトでは、これらの基本的な概念を組み合わせて、より複雑で機能的なアプリケーションを構築していきます。
