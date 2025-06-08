# 1. Next.js & React フロントエンドベストプラクティス

## 1.1. 目次

- [1. Next.js \& React フロントエンドベストプラクティス](#1-nextjs--react-フロントエンドベストプラクティス)
  - [1.1. 目次](#11-目次)
  - [1.2. Container/Presentationalパターン](#12-containerpresentationalパターン)
    - [1.2.1. パターンの基本概念](#121-パターンの基本概念)
    - [1.2.2. Next.js 15での最適な実装方法](#122-nextjs-15での最適な実装方法)
    - [1.2.3. パターン適用のメリット](#123-パターン適用のメリット)
  - [1.3. 最適なプロジェクト構造](#13-最適なプロジェクト構造)
    - [1.3.1. スケーラブルなディレクトリ構成](#131-スケーラブルなディレクトリ構成)
    - [1.3.2. 機能モジュール型構成](#132-機能モジュール型構成)
    - [1.3.3. ルートパスの推奨扱い方](#133-ルートパスの推奨扱い方)
  - [1.4. 高度なスタイリング戦略とデザインシステム](#14-高度なスタイリング戦略とデザインシステム)
    - [1.4.1. スタイリング手法の選定と比較](#141-スタイリング手法の選定と比較)
    - [1.4.2. デザインシステムの構築とコンポーネント階層](#142-デザインシステムの構築とコンポーネント階層)
  - [1.5. 高度な状態管理戦略](#15-高度な状態管理戦略)
    - [1.5.1. クライアント状態の最適な扱い方](#151-クライアント状態の最適な扱い方)
    - [1.5.2. サーバー状態との連携](#152-サーバー状態との連携)
  - [1.6. 高度なAPI通信設計](#16-高度なapi通信設計)
    - [1.6.1. API層のモジュール化](#161-api層のモジュール化)
    - [1.6.2. データフェッチング戦略](#162-データフェッチング戦略)
      - [1.6.2.1. コンテナ層での並列データフェッチング](#1621-コンテナ層での並列データフェッチング)
      - [1.6.2.2. 効果的なエラーハンドリング](#1622-効果的なエラーハンドリング)
    - [1.5.3. APIレスポンスのランタイムバリデーション (zod活用)](#153-apiレスポンスのランタイムバリデーション-zod活用)
  - [1.7. 包括的なエラーハンドリング](#17-包括的なエラーハンドリング)
    - [1.7.1. 段階的なエラーバウンダリ](#171-段階的なエラーバウンダリ)
    - [1.7.2. 機能単位のエラーバウンダリ](#172-機能単位のエラーバウンダリ)
    - [1.7.3. コンポーネントレベルのErrorBoundary](#173-コンポーネントレベルのerrorboundary)
  - [1.8. 実務的なログ収集](#18-実務的なログ収集)
    - [1.8.1. 構造化ログの設計](#181-構造化ログの設計)
    - [1.8.2. クライアントサイドのログ収集](#182-クライアントサイドのログ収集)
    - [1.8.3. ユーザーインタラクションのログ記録](#183-ユーザーインタラクションのログ記録)
    - [1.8.4. パフォーマンス情報の収集](#184-パフォーマンス情報の収集)
  - [1.9. 高度なテスト戦略](#19-高度なテスト戦略)
    - [1.9.1. Presentationalコンポーネントのテスト](#191-presentationalコンポーネントのテスト)
    - [1.9.2. Server Componentsのテスト戦略](#192-server-componentsのテスト戦略)
    - [1.9.3. テストの種類と目的](#193-テストの種類と目的)
  - [1.10. パフォーマンス最適化](#110-パフォーマンス最適化)
    - [1.10.1. Server/Client分割の最適化](#1101-serverclient分割の最適化)
    - [1.10.2. 画像最適化](#1102-画像最適化)
    - [1.10.3. コンポーネントの遅延ロード](#1103-コンポーネントの遅延ロード)
    - [1.10.4. Web Vitalsの最適化](#1104-web-vitalsの最適化)
  - [1.11. セキュリティとアクセシビリティ](#111-セキュリティとアクセシビリティ)
    - [1.11.1. 認証・認可設計](#1111-認証認可設計)
    - [1.11.2. アクセシブルなコンポーネント設計](#1112-アクセシブルなコンポーネント設計)
    - [1.11.3. アクセシビリティのベストプラクティス](#1113-アクセシビリティのベストプラクティス)
  - [1.12. 国際化と拡張性](#112-国際化と拡張性)
    - [1.12.1. 多言語対応](#1121-多言語対応)
    - [1.12.2. 拡張性の高いコード設計](#1122-拡張性の高いコード設計)
    - [1.12.3. ドキュメント戦略](#1123-ドキュメント戦略)

## 1.2. Container/Presentationalパターン

Next.js 15のReact Server Componentsと相性の良いContainer/Presentationalパターンを紹介します。

### 1.2.1. パターンの基本概念

Container/Presentationalパターンは、アプリケーションの関心事を分離するためのデザインパターンです：

- **Presentationalコンポーネント**（表示担当）
  - UIの見た目のみに責任を持つ
  - データをpropsを通じてのみ受け取る
  - 基本的に状態を持たない（UIに関する状態のみ例外的に許容）
  - 再利用しやすい純粋なコンポーネント

- **Containerコンポーネント**（ロジック担当）
  - データの取得やビジネスロジックを担当
  - APIや状態管理と連携
  - 取得したデータをPresentationalコンポーネントにpropsとして渡す

### 1.2.2. Next.js 15での最適な実装方法

Next.js 15のApp Routerでは、このパターンが自然に適合します：

1. **Server Components = Container**
   - サーバー側でデータ取得とロジック処理
   - `page.tsx` をContainerとして使用

    ```tsx
    // app/products/page.tsx (Container / Server Component)
    import { ProductsClient } from './client';
    import { fetchProducts } from '@/services/products';

    export default async function ProductsPage() {
      // サーバーサイドでデータ取得
      const products = await fetchProducts();

      // 重い処理もサーバーで実行可能
      const processedData = heavyDataProcessing(products);

      // 必要なデータのみをClient Componentに渡す
      return <ProductsClient products={processedData} />;
    }
    ```

2. **Client Components = Presentational**
   - クライアント側でUIとユーザーインタラクションを担当
   - `client.tsx` をPresentationalとして使用

    ```tsx
    // app/products/client.tsx (Presentational / Client Component)
    'use client';

    import { useState } from 'react';
    import { ProductCard } from '@/components/ProductCard';

    export function ProductsClient({ products }) {
      // UIに関する状態のみを管理
      const [sortOrder, setSortOrder] = useState('name');

      // クライアント側のUIロジック
      const sortedProducts = [...products].sort((a, b) => {
        return sortOrder === 'name'
          ? a.name.localeCompare(b.name)
          : a.price - b.price;
      });

      return (
        <div>
          <div className="controls">
            <button onClick={() => setSortOrder('name')}>名前でソート</button>
            <button onClick={() => setSortOrder('price')}>価格でソート</button>
          </div>
          <div className="product-grid">
            {sortedProducts.map(product => (
              <ProductCard key={product.id} product={product} />
            ))}
          </div>
        </div>
      );
    }
    ```

### 1.2.3. パターン適用のメリット

- **明確な責任分離**: データ取得とUI表示の関心事が分離される
- **パフォーマンス向上**: Server Componentsがデータ取得を担うことでFCP (First Contentful Paint) が改善
- **コード再利用性の向上**: Presentationalコンポーネントが純粋なUIとして再利用可能
- **テスト容易性**: UIとビジネスロジックが分離されていることでテストが書きやすい

## 1.3. 最適なプロジェクト構造

### 1.3.1. スケーラブルなディレクトリ構成

Next.js 15のプロジェクトでは、以下のような構造が推奨されます：

```text
src/
├── app/                             # Next.js App Router
│   ├── (features)/                  # 機能別ディレクトリグループ (URLパスには含まれない)
│   │   ├── products/                # 商品関連機能
│   │   │   ├── page.tsx             # Server Component (Container)
│   │   │   ├── client.tsx           # Presentational Component
│   │   │   ├── loading.tsx          # ローディングUI
│   │   │   └── components/          # 機能固有の小さなコンポーネント
│   │   │       ├── ProductCard.tsx
│   │   │       └── ProductFilter.tsx
│   │   └── home/                    # ホームページ機能
│   │       ├── page.tsx             # Server Component (Container)
│   │       └── client.tsx           # Presentational Component
│   ├── api/                         # API Routes
│   │   └── products/
│   │       └── route.ts
│   ├── page.tsx                     # ルートページ (/) のContainer
│   └── layout.tsx                   # ルートレイアウト
├── components/                      # グローバル共通コンポーネント
│   ├── ui/                          # 汎用UIコンポーネント
│   │   ├── button.tsx
│   │   ├── card.tsx
│   │   ├── input.tsx
│   │   └── modal.tsx
│   └── layout/                      # レイアウトコンポーネント
│       ├── header.tsx
│       ├── footer.tsx
│       └── sidebar.tsx
├── lib/                             # ユーティリティ関数
│   ├── utils/                       # 汎用ユーティリティ
│   │   ├── date.ts                  # 日付操作関数
│   │   └── format.ts                # フォーマット関数
│   └── hooks/                       # カスタムフック
│       ├── useForm.ts
│       └── useClickOutside.ts
├── services/                        # APIクライアント
│   ├── api-client.ts                # ベースAPI設定
│   └── products/                    # 製品関連API
│       ├── types.ts                 # 型定義
│       └── api.ts                   # API関数
└── types/                           # 型定義
```

### 1.3.2. 機能モジュール型構成

大規模なプロジェクトでは、機能ごとに完結したモジュール構造にすることで保守性が向上します：

```text
src/
├── app/                      # ルーティング定義
│   └── (features)/           # 機能別ルート
├── features/                 # 機能モジュール
│   ├── products/             # 商品機能
│   │   ├── components/       # 商品関連UI
│   │   ├── hooks/            # 商品関連ロジック
│   │   ├── services/         # 商品API
│   │   ├── types/            # 商品関連型定義
│   │   └── utils/            # 商品関連ユーティリティ
│   └── cart/                 # カート機能
├── shared/                   # 共有リソース
    ├── components/           # 共通コンポーネント
    ├── hooks/                # 共通フック
    └── utils/                # 共通ユーティリティ
```

### 1.3.3. ルートパスの推奨扱い方

```tsx
// src/app/page.tsx
import HomePage from "./(features)/home/page";

export default function Home() {
  return <HomePage />;
}
```

このアプローチにより、ホームページの実装も他の機能と一貫性のあるパターンで管理できます。

## 1.4. 高度なスタイリング戦略とデザインシステム

Next.jsアプリケーションの規模が大きくなるにつれて、一貫性があり、保守性の高いスタイリング戦略と、それを支えるデザインシステムの構築が重要になります。

### 1.4.1. スタイリング手法の選定と比較

プロジェクトの要件やチームのスキルセットに応じて、最適なスタイリング手法を選定します。

- **Tailwind CSS (推奨)**:
  - **特徴**: ユーティリティファーストのCSSフレームワーク。Server ComponentsとClient Componentsの両方でシームレスに利用可能。ビルド時に未使用のスタイルを削除するため、バンドルサイズが最適化される。
  - **利点**: 迅速なプロトタイピング、高いカスタマイズ性、HTML内でスタイルが完結することによる見通しの良さ。
  - **考慮点**: クラス名が多くなることによるHTMLの肥大化、学習コスト。`@apply`の使いすぎに注意。

- **CSS Modules**:
  - **特徴**: ファイルスコープのCSS。Next.jsで標準サポート。Server Componentsでも利用可能。
  - **利点**: グローバルな名前空間の衝突を回避。Reactコンポーネントとの親和性が高い。
  - **例**:

      ```jsx
      // Button.module.css
      .button { background-color: blue; color: white; }

      // Button.tsx
      import styles from './Button.module.css';
      function Button() { return <button className={styles.button}>Click me</button>; }
      ```

- **CSS-in-JS (例: styled-components, Emotion)**:
  - **特徴**: JavaScript内でCSSを記述。動的なスタイリングやテーマ設定に強み。
  - **利点**: コンポーネントとスタイルが密接に結合、Propsに基づいた動的スタイル変更が容易。
  - **考慮点**: `'use client'`ディレクティブが必要なため、主にClient Componentsで使用。ランタイムのオーバーヘッドやバンドルサイズへの影響を考慮。

      ```tsx
      // StyledButton.tsx (Client Component)
      'use client';
      import styled from 'styled-components';
      const StyledButton = styled.button`
        background-color: ${props => props.primary ? 'blue' : 'gray'};
        color: white;
      `;
      ```

- **コンポーネントライブラリの採用**:
  - **shadcn/ui**: Tailwind CSSベース。Server/Client Componentsフレンドリー。コピー＆ペーストでコンポーネントをプロジェクトに追加する形式で、カスタマイズ性が高い。
  - **Material UI (MUI), Chakra UI**: Client Components向け。リッチなUIコンポーネント群とデザインシステムを提供。

### 1.4.2. デザインシステムの構築とコンポーネント階層

一貫したUI/UXを提供し、開発効率を向上させるためにデザインシステムを構築します。Atomic Designの考え方を参考に、コンポーネントを階層化します。

- **Atoms (原子)**: UIの最小単位。ボタン、インプットフィールド、ラベルなど。
  - 例: `src/components/ui/button.tsx`, `src/components/ui/input.tsx`
- **Molecules (分子)**: 複数のAtomsを組み合わせた小さな機能単位。検索フォーム（ラベル、インプット、ボタン）、ナビゲーションアイテムなど。
  - 例: `src/components/molecules/SearchForm.tsx` (プロジェクトによっては `src/components/ui/` や機能別ディレクトリ内に配置)
- **Organisms (有機体)**: 複数のMoleculesやAtomsを組み合わせた、より複雑なUIセクション。ヘッダー、商品カード、商品リストなど。
  - 例: `src/components/layout/Header.tsx`, `src/app/(features)/products/components/ProductCard.tsx`
- **Templates (テンプレート)**: ページのレイアウト構造を示すワイヤーフレーム。具体的なコンテンツは含まない。
  - 例: `src/components/templates/TwoColumnLayout.tsx`
- **Pages (ページ)**: Templatesに具体的なコンテンツを流し込んだ最終的なページ。Next.jsでは `app/.../page.tsx` がこれに該当。

**ツールとプラクティス**:

- **Storybook**: UIコンポーネントを個別に開発・テスト・ドキュメント化するためのツール。デザインシステムの一覧性を高めます。
- **共通のデザイントークン**: 色、フォントサイズ、スペーシングなどの値を一元管理し、コンポーネント間で共有します。Tailwind CSSの場合は `tailwind.config.js` で設定。
- **アクセシビリティガイドライン**: デザインシステムの初期段階からアクセシビリティを考慮します。

## 1.5. 高度な状態管理戦略

### 1.5.1. クライアント状態の最適な扱い方

複雑な状態管理には、状況に応じて適切なツールを選択します：

1. **React Hooksの最適な活用**

    ```tsx
    // 単純な状態管理: useState
    const [filter, setFilter] = useState('all');

    // 複雑な状態管理: useReducer
    const [state, dispatch] = useReducer(productsReducer, initialState);

    // 最適化されたコンテキスト
    const CartContext = createContext<CartContextType | null>(null);
    ```

2. **状態管理ライブラリの選定基準**

    | ライブラリ    | 用途                     | Container/Presentationalとの親和性              |
    | ------------- | ------------------------ | ----------------------------------------------- |
    | Redux Toolkit | 大規模なグローバル状態   | Container層でストア設定、Presentational層で利用 |
    | Zustand       | シンプルで軽量な状態管理 | Client Componentsでの利用に最適                 |
    | Jotai/Recoil  | 原子的な状態管理         | 細分化されたPresentational間での状態共有        |

### 1.5.2. サーバー状態との連携

React QueryやSWRを使用して、サーバーデータとクライアントの状態を効率的に同期します：

```tsx
'use client'

import { useQuery } from '@tanstack/react-query';

function ProductsClient({ initialProducts }) {
  // 初期データはServer Component (Container)から提供
  // その後の更新はクライアント側でハンドリング
  const { data: products } = useQuery({
    queryKey: ['products'],
    queryFn: fetchProducts,
    initialData: initialProducts,
  });

  return /* JSX */;
}
```

このアプローチは、次の利点があります：

- サーバーからの初期データをすぐに表示（高速な初期ロード）
- クライアントサイドでの自動再取得（最新データの維持）
- キャッシュと再検証機能（効率的なデータ管理）

## 1.6. 高度なAPI通信設計

### 1.6.1. API層のモジュール化

API層を適切にモジュール化することで、メンテナンスとテストが容易になります：

```tsx
// services/api-client.ts
import axios from 'axios';

export const apiClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// services/products/api.ts
import { apiClient } from '../api-client';
import type { Product, ProductsResponse } from './types';

export async function fetchProducts(): Promise<ProductsResponse> {
  const response = await apiClient.get<ProductsResponse>('/products');
  return response.data;
}

export async function fetchProductById(id: string): Promise<Product> {
  const response = await apiClient.get<Product>(`/products/${id}`);
  return response.data;
}
```

### 1.6.2. データフェッチング戦略

#### 1.6.2.1. コンテナ層での並列データフェッチング

```tsx
// ページレベルでの並列データフェッチング
export default async function ProductPage({ params }) {
  // 並列データ取得
  const [product, recommendations] = await Promise.all([
    getProduct(params.id),
    getRecommendations(params.id)
  ]);

  return (
    <div>
      <ProductDetails product={product} />
      <RecommendationList items={recommendations} />
    </div>
  );
}
```

#### 1.6.2.2. 効果的なエラーハンドリング

```tsx
// エラーハンドリングとフォールバック戦略
export default async function ProductsPage() {
  try {
    const products = await fetchProducts();
    return <ProductsList products={products} />;
  } catch (error) {
    // 専用エラーUIを表示
    return <ProductsErrorFallback error={error} />;
  }
}
```

### 1.5.3. APIレスポンスのランタイムバリデーション (zod活用)

TypeScriptの型定義は開発時の静的チェックには非常に有効ですが、外部APIからのレスポンスなど、実行時に実際に受け取るデータが期待通りであるかを保証するものではありません。ここで、`zod`のようなスキーマバリデーションライブラリを使用することで、ランタイムでのデータ検証と型安全性を強化できます。

**zodの導入**:

```bash
npm install zod
#または
yarn add zod
```

**APIクライアントでの適用例**:

```typescript
// services/products/types.ts
import { z } from 'zod';

// Productのスキーマ定義
export const productSchema = z.object({
  id: z.string().uuid(),
  name: z.string().min(1, "Name is required"),
  description: z.string().optional(),
  price: z.number().positive("Price must be positive"),
  imageUrl: z.string().url().optional(),
  // ...その他のプロパティ
});

// Productの型をスキーマから推論
export type Product = z.infer<typeof productSchema>;

// 商品一覧APIレスポンスのスキーマ
export const productsResponseSchema = z.object({
  items: z.array(productSchema),
  total: z.number().int().nonnegative(),
  page: z.number().int().positive(),
  pageSize: z.number().int().positive(),
});

export type ProductsResponse = z.infer<typeof productsResponseSchema>;


// services/products/api.ts (既存のファイルに追記・修正)
import { apiClient } from '../api-client';
import { logger } from '@/lib/logging'; // ロガーをインポート
import {
  type Product,
  productSchema,
  type ProductsResponse,
  productsResponseSchema
} from './types'; // zodスキーマと型をインポート

export async function fetchProducts(params?: { page?: number; pageSize?: number }): Promise<ProductsResponse> {
  const response = await apiClient.get('/products', { params });
  // zodでレスポンスデータをバリデーションし、型安全なデータを返す
  try {
    return productsResponseSchema.parse(response.data);
  } catch (error) {
    logger.error({
      message: "API response validation failed for fetchProducts",
      error: error.errors, // zodのエラー詳細
      responseData: response.data, // 実際のレスポンスデータ（デバッグ用）
    });
    // プロジェクトのエラーハンドリング戦略に応じてエラーをスローまたは適切なフォールバック値を返す
    throw new Error("Invalid API response structure for products list.");
  }
}

export async function fetchProductById(id: string): Promise<Product> {
  const response = await apiClient.get<Product>(`/products/${id}`);
  try {
    return productSchema.parse(response.data);
  } catch (error) {
    logger.error({
      message: `API response validation failed for fetchProductById (id: ${id})`,
      error: error.errors,
      responseData: response.data,
    });
    throw new Error(`Invalid API response structure for product (id: ${id}).`);
  }
}
```

**zod活用のメリット**:

- **ランタイムの型安全性**: APIの仕様変更や予期せぬデータ構造に対して早期に問題を検知。
- **明確なエラーメッセージ**: バリデーションエラー時に、どのフィールドが期待と異なるかを具体的に把握可能。
- **型定義の自動推論**: スキーマからTypeScriptの型を生成できるため、型定義とバリデーションルールを一元管理。
- **複雑なバリデーション**: 文字列のフォーマット、数値の範囲、必須項目など、詳細なバリデーションルールを定義可能。

**Server Actionsでの活用**:
Server Actionsでクライアントからの入力を受け取る際にも、zodを使用してバリデーションを行うことで、セキュリティとデータ整合性を向上させることができます。

```typescript
// app/(features)/products/actions.ts
'use server';
import { z } from 'zod';
import { revalidatePath } from 'next/cache';
import { logger } from '@/lib/logging'; // ロガーをインポート
// ... (他のimport)

const createProductSchema = z.object({
  name: z.string().min(3, "Product name must be at least 3 characters"),
  price: z.coerce.number().positive("Price must be a positive number"), // 文字列で来ても数値に変換
  // ...
});

export async function createProductAction(prevState: any, formData: FormData) {
  const rawFormData = Object.fromEntries(formData.entries());

  const validationResult = createProductSchema.safeParse(rawFormData);

  if (!validationResult.success) {
    logger.warn({
        message: "Product creation validation failed",
        errors: validationResult.error.flatten().fieldErrors,
        formData: rawFormData,
    });
    return {
      errors: validationResult.error.flatten().fieldErrors,
      message: "Failed to create product due to validation errors."
    };
  }

  const { name, price } = validationResult.data;

  try {
    // DBへの保存処理など
    // await db.product.create({ data: { name, price }});
    logger.info({ message: "Product creation attempt", data: { name, price } });
    revalidatePath('/products');
    return { message: "Product created successfully.", errors: null };
  } catch (e) {
    logger.error({ message: "Failed to create product in DB", error: e, data: { name, price } });
    return { message: "Failed to create product.", errors: null };
  }
}
```

この`createProductAction`の例は、React Hook Formと連携してエラーメッセージを表示する際などに役立ちます。

## 1.7. 包括的なエラーハンドリング

### 1.7.1. 段階的なエラーバウンダリ

複数レベルでエラーをキャプチャすることで、ユーザー体験を維持します：

```tsx
// app/error.tsx (グローバルエラーバウンダリ)
'use client'

import { useEffect } from 'react';
import { ErrorDisplay } from '@/components/ui/error-display';

export default function GlobalError({
  error,
  reset
}: {
  error: Error & { digest?: string }
  reset: () => void
}) {
  useEffect(() => {
    // エラーログ送信
    console.error(error);
  }, [error]);

  return (
    <html>
      <body>
        <ErrorDisplay
          title="予期せぬエラーが発生しました"
          message={process.env.NODE_ENV === 'development' ? error.message : null}
          actionLabel="再試行"
          actionFn={reset}
        />
      </body>
    </html>
  );
}
```

### 1.7.2. 機能単位のエラーバウンダリ

```tsx
// app/(features)/products/error.tsx
'use client'

export default function ProductsError({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  return (
    <div className="error-container">
      <h2>商品データの取得中にエラーが発生しました</h2>
      <button onClick={reset}>再試行</button>
    </div>
  );
}
```

### 1.7.3. コンポーネントレベルのErrorBoundary

特定のコンポーネントのエラーを隔離します：

```tsx
// components/ErrorBoundary.tsx
import React from 'react';

interface ErrorBoundaryProps {
  fallback?: React.ReactNode;
  children: React.ReactNode;
  componentName: string;
}

export class ErrorBoundary extends React.Component<
  ErrorBoundaryProps,
  { hasError: boolean; error?: Error }
> {
  constructor(props: ErrorBoundaryProps) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error: Error) {
    return { hasError: true, error };
  }

  componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
    console.error(
      `Error in component: ${this.props.componentName}`,
      error,
      errorInfo
    );
  }

  render() {
    if (this.state.hasError) {
      return this.props.fallback || (
        <div className="error-container">
          <h2>コンポーネントでエラーが発生しました</h2>
        </div>
      );
    }

    return this.props.children;
  }
}
```

使用例:

```tsx
<ErrorBoundary componentName="ProductList">
  <ProductList products={products} />
</ErrorBoundary>
```

## 1.8. 実務的なログ収集

### 1.8.1. 構造化ログの設計

効率的なログ分析のために構造化ログを実装します：

```typescript
// lib/logging.ts
import pino from 'pino';

// 開発環境ではフォーマットされた出力、本番ではJSON形式
const isDev = process.env.NODE_ENV !== 'production';

export const logger = pino({
  level: process.env.LOG_LEVEL || 'info',
  transport: isDev
    ? {
        target: 'pino-pretty',
        options: {
          colorize: true,
          translateTime: 'SYS:standard',
        },
      }
    : undefined,
  // 共通フィールドの追加
  base: {
    env: process.env.NODE_ENV,
    appName: 'ecommerce-platform',
  },
});

// リクエストIDなどのコンテキスト情報を含めたロガーを生成
export function getContextLogger(context: Record<string, any>) {
  return logger.child(context);
}
```

### 1.8.2. クライアントサイドのログ収集

ブラウザからサーバーにログを効率的に送信する仕組み：

```typescript
// lib/client-logging.ts
type LogLevel = 'error' | 'warn' | 'info' | 'debug';

interface LogEntry {
  level: LogLevel;
  message: string;
  timestamp: string;
  data?: any;
  context?: Record<string, any>;
}

// ログバッファ
const logBuffer: LogEntry[] = [];
let isFlushScheduled = false;

// ログエントリの作成
function createLogEntry(level: LogLevel, message: string, data?: any): LogEntry {
  return {
    level,
    message,
    timestamp: new Date().toISOString(),
    data,
    context: {
      url: window.location.href,
      userAgent: navigator.userAgent,
    },
  };
}

// APIを使ってサーバーにログを送信
async function flushLogs() {
  if (logBuffer.length === 0) {
    isFlushScheduled = false;
    return;
  }

  try {
    const logs = [...logBuffer];
    logBuffer.length = 0;

    await fetch('/api/logs', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ logs }),
    });
  } catch (error) {
    console.error('Failed to send logs to server:', error);
    // 失敗したログを再度バッファに追加（バッファオーバーフロー防止の制限付き）
    if (logBuffer.length < 100) {
      logBuffer.unshift(...logBuffer);
    }
  }

  isFlushScheduled = false;
}

// ログをバッファに追加し、必要に応じてフラッシュをスケジュール
function logToServer(level: LogLevel, message: string, data?: any) {
  const entry = createLogEntry(level, message, data);
  logBuffer.push(entry);

  // バッファがある程度たまったら送信、または定期的に送信
  if (logBuffer.length >= 10 && !isFlushScheduled) {
    isFlushScheduled = true;
    setTimeout(flushLogs, 1000);
  } else if (!isFlushScheduled) {
    isFlushScheduled = true;
    setTimeout(flushLogs, 5000); // 最大5秒後にフラッシュ
  }
}

// 公開API
export const clientLogger = {
  error: (message: string, data?: any) => logToServer('error', message, data),
  warn: (message: string, data?: any) => logToServer('warn', message, data),
  info: (message: string, data?: any) => logToServer('info', message, data),
  debug: (message: string, data?: any) => logToServer('debug', message, data),
};

// ページ離脱時に残りのログを送信
if (typeof window !== 'undefined') {
  window.addEventListener('beforeunload', () => {
    if (logBuffer.length > 0) {
      navigator.sendBeacon('/api/logs', JSON.stringify({ logs: logBuffer }));
    }
  });
}
```

### 1.8.3. ユーザーインタラクションのログ記録

ユーザー行動をトラッキングして分析します：

```typescript
// lib/event-tracking.ts
import { clientLogger } from './client-logging';

type EventCategory = 'navigation' | 'interaction' | 'error' | 'performance';

interface TrackEventOptions {
  category: EventCategory;
  action: string;
  label?: string;
  value?: number;
  metadata?: Record<string, any>;
}

export function trackEvent({
  category,
  action,
  label,
  value,
  metadata = {},
}: TrackEventOptions) {
  // 個人情報をフィルタリング
  const sanitizedMetadata = { ...metadata };

  // JSONに変換できるデータのみを許可
  try {
    JSON.stringify(sanitizedMetadata);
  } catch (error) {
    clientLogger.warn('Invalid metadata in trackEvent', {
      error: error.message,
    });
    return;
  }

  clientLogger.info(`Event: ${action}`, {
    category,
    action,
    label,
    value,
    ...sanitizedMetadata,
  });
}
```

使用例：

```tsx
import { trackEvent } from '@/lib/event-tracking';

function ProductCard({ id, name, price }) {
  const handleAddToCart = () => {
    // カートに追加するロジック

    // イベント記録
    trackEvent({
      category: 'interaction',
      action: 'add_to_cart',
      label: name,
      value: price,
      metadata: { productId: id }
    });
  };

  return (
    <div className="product-card">
      <h3>{name}</h3>
      <button onClick={handleAddToCart}>カートに追加</button>
    </div>
  );
}
```

### 1.8.4. パフォーマンス情報の収集

Web Vitalsなどのパフォーマンスメトリクスを計測します：

```typescript
// lib/performance-tracking.ts
import { clientLogger } from './client-logging';

interface WebVitalMetric {
  id: string;
  name: string;
  value: number;
  rating: 'good' | 'needs-improvement' | 'poor';
  delta: number;
}

// 結果を評価範囲に基づいて分類
function getRating(name: string, value: number): 'good' | 'needs-improvement' | 'poor' {
  switch (name) {
    case 'CLS':
      return value <= 0.1 ? 'good' : value <= 0.25 ? 'needs-improvement' : 'poor';
    case 'FID':
    case 'TTFB':
      return value <= 100 ? 'good' : value <= 300 ? 'needs-improvement' : 'poor';
    case 'LCP':
      return value <= 2500 ? 'good' : value <= 4000 ? 'needs-improvement' : 'poor';
    default:
      return 'good';
  }
}

export function reportWebVital(metric: WebVitalMetric) {
  const rating = getRating(metric.name, metric.value);

  clientLogger.info(`Web Vital: ${metric.name}`, {
    category: 'performance',
    metricName: metric.name,
    value: Math.round(metric.value),
    rating,
    delta: Math.round(metric.delta),
    id: metric.id,
  });
}
```

Next.jsへの統合：

```tsx
// app/layout.tsx
import { reportWebVital } from '@/lib/performance-tracking';

export function reportWebVitals(metric) {
  reportWebVital(metric);
}
```

## 1.9. 高度なテスト戦略

### 1.9.1. Presentationalコンポーネントのテスト

Presentationalコンポーネント（Client Components）のテスト例：

```tsx
// app/(features)/products/__tests__/client.test.tsx
import { render, screen, fireEvent } from '@testing-library/react';
import { ProductsClient } from '../client';

// モックデータ
const mockProducts = [
  { id: '1', name: 'Product A', price: 100 },
  { id: '2', name: 'Product B', price: 200 },
];

describe('ProductsClient', () => {
  it('商品リストを正しくレンダリングする', () => {
    render(<ProductsClient products={mockProducts} />);

    expect(screen.getByText('Product A')).toBeInTheDocument();
    expect(screen.getByText('Product B')).toBeInTheDocument();
  });

  it('価格順でソートされる', () => {
    render(<ProductsClient products={mockProducts} />);

    // 最初は名前順
    const items = screen.getAllByRole('listitem');
    expect(items[0]).toHaveTextContent('Product A');

    // 価格順に切り替え
    fireEvent.click(screen.getByText('価格でソート'));

    // 価格順になっていることを確認
    const sortedItems = screen.getAllByRole('listitem');
    expect(sortedItems[0]).toHaveTextContent('Product A');
    expect(sortedItems[1]).toHaveTextContent('Product B');
  });
});
```

### 1.9.2. Server Componentsのテスト戦略

Server Componentsは直接テストすることが難しいため、以下のアプローチを使用します：

1. **ロジック部分を分離してテスト**:

    ```tsx
    // app/(features)/products/actions.ts
    export async function fetchProducts() {
      // データ取得ロジック
      return await fetch('https://api.example.com/products').then(res => res.json());
    }

    // app/(features)/products/page.tsx
    import { ProductsClient } from './client';
    import { fetchProducts } from './actions';

    export default async function ProductsPage() {
      const products = await fetchProducts();
      return <ProductsClient products={products} />;
    }
    ```

    ```tsx
    // actions.test.ts
    import { fetchProducts } from '../actions';

    // グローバルFetchのモック
    global.fetch = jest.fn();

    describe('Product Actions', () => {
      it('製品データを正しく取得する', async () => {
        // fetchのモック実装
        (global.fetch as jest.Mock).mockResolvedValueOnce({
          ok: true,
          json: async () => ([
            { id: '1', name: 'Product A', price: 100 },
            { id: '2', name: 'Product B', price: 200 },
          ])
        });

        const products = await fetchProducts();

        expect(products).toHaveLength(2);
        expect(products[0].name).toBe('Product A');
        expect(global.fetch).toHaveBeenCalledWith(
          'https://api.example.com/products'
        );
      });
    });
    ```

2. **E2Eテストでの統合テスト**:

    ```typescript
    // e2e/products.spec.ts (Playwright)
    import { test, expect } from '@playwright/test';

    test('製品ページが正しくレンダリングされる', async ({ page }) => {
      await page.goto('/products');

      // ページが正しくロードされたことを確認
      await expect(page.getByRole('heading', { name: '製品リスト' })).toBeVisible();

      // 製品リストが表示されていることを確認
      const productItems = page.getByRole('listitem');
      await expect(productItems).toHaveCount(await productItems.count());

      // ソート機能が動作することを確認
      await page.getByRole('button', { name: '価格でソート' }).click();
      // ソート後の状態を確認するロジック
    });
    ```

### 1.9.3. テストの種類と目的

包括的なテスト戦略には、次の種類のテストを含めます：

1. **ユニットテスト**: 小さな関数やコンポーネントの個別テスト
2. **統合テスト**: 複数のコンポーネントが連携して動作することを確認
3. **E2Eテスト**: ユーザーの視点からアプリケーション全体をテスト

**テスト優先順位**:

1. Presentationalコンポーネント（UIとイベントハンドラ）
2. ユーティリティ関数とカスタムフック
3. データ取得ロジック（actions, services）
4. E2Eテストで統合されたアプリケーションフロー

## 1.10. パフォーマンス最適化

### 1.10.1. Server/Client分割の最適化

最適なコンポーネント分割により、パフォーマンスを向上させます：

```tsx
// page.tsx (Server Component / Container)
import { ProductsClient } from './client';
import { fetchProducts } from '@/services/products';

export default async function ProductsPage() {
  const products = await fetchProducts();

  // 重いライブラリ処理はServer Componentで実行
  const processedData = heavyDataProcessing(products);

  return (
    <main>
      <h1>Products</h1>
      {/* 必要なデータのみをClient Componentに渡す */}
      <ProductsClient
        products={processedData.items}
        totalCount={processedData.total}
      />
    </main>
  );
}
```

### 1.10.2. 画像最適化

Next.jsの画像コンポーネントを使用してパフォーマンスを向上：

```tsx
import Image from 'next/image';

function ProductImage({ product }) {
  return (
    <Image
      src={product.imageUrl}
      alt={product.name}
      width={300}
      height={200}
      placeholder="blur"
      blurDataURL={product.blurUrl}
      priority={product.featured}
    />
  );
}
```

### 1.10.3. コンポーネントの遅延ロード

必要なタイミングでコンポーネントを読み込むことでバンドルサイズを削減：

```tsx
import dynamic from 'next/dynamic';

// 重いコンポーネントを動的にインポート
const HeavyChart = dynamic(() => import('@/components/HeavyChart'), {
  loading: () => <p>Loading chart...</p>,
  ssr: false, // クライアントサイドのみでレンダリング
});

export default function DashboardPage() {
  return (
    <div>
      <h1>Dashboard</h1>
      <HeavyChart data={chartData} />
    </div>
  );
}
```

### 1.10.4. Web Vitalsの最適化

CoreWeb Vitalsを改善するためのテクニック：

1. **LCP (Largest Contentful Paint)の最適化**:
   - プリオリティの高い画像に `priority` 属性を使用
   - クリティカルパスのCSSを最小限に抑える
   - Suspenseを使用して重要なコンテンツから表示

2. **CLS (Cumulative Layout Shift)の防止**:
   - 画像やメディアに明示的なサイズを設定
   - プレースホルダーを使用して領域を確保
   - 動的コンテンツのレイアウトシフトを防止

3. **FID (First Input Delay)の改善**:
   - JavaScriptの分割とプライオリタイズ
   - ヘビーな処理をWebワーカーに移動
   - サードパーティスクリプトの読み込みを遅延

## 1.11. セキュリティとアクセシビリティ

### 1.11.1. 認証・認可設計

Next.jsでのセキュリティ実装：

```tsx
// lib/auth.ts
import { NextAuthOptions } from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';

export const authOptions: NextAuthOptions = {
  providers: [
    CredentialsProvider({
      // 認証プロバイダー実装
    }),
  ],
  callbacks: {
    // カスタムコールバック実装
    async jwt({ token, user }) {
      if (user) {
        token.role = user.role;
      }
      return token;
    },
    async session({ session, token }) {
      if (token) {
        session.user.role = token.role;
      }
      return session;
    },
  },
  pages: {
    signIn: '/auth/signin',
    error: '/auth/error',
  },
};
```

```tsx
// Server Actionsでの認証チェック
'use server'

import { getServerSession } from 'next-auth';
import { authOptions } from '@/lib/auth';

export async function createProduct(formData: FormData) {
  // 認証・認可チェック
  const session = await getServerSession(authOptions);
  if (!session || session.user.role !== 'admin') {
    throw new Error('Unauthorized');
  }

  // 処理実行
  // ...
}
```

### 1.11.2. アクセシブルなコンポーネント設計

```tsx
// components/Button.tsx
import { forwardRef } from 'react';

type ButtonProps = {
  variant?: 'primary' | 'secondary' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  isLoading?: boolean;
} & React.ButtonHTMLAttributes<HTMLButtonElement>;

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ variant = 'primary', size = 'md', isLoading, children, ...props }, ref) => {
    // スタイルクラスの構築
    const baseClasses = 'rounded focus:outline-none focus:ring-2 focus:ring-offset-2';
    const variantClasses = {
      primary: 'bg-blue-600 text-white hover:bg-blue-700 focus:ring-blue-500',
      secondary: 'bg-gray-200 text-gray-900 hover:bg-gray-300 focus:ring-gray-500',
      danger: 'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500',
    };
    const sizeClasses = {
      sm: 'px-2 py-1 text-sm',
      md: 'px-4 py-2',
      lg: 'px-6 py-3 text-lg',
    };

    const classes = `${baseClasses} ${variantClasses[variant]} ${sizeClasses[size]} ${props.className || ''}`;

    return (
      <button
        ref={ref}
        className={classes}
        disabled={isLoading || props.disabled}
        aria-disabled={isLoading || props.disabled}
        {...props}
      >
        {isLoading ? (
          <>
            <span className="sr-only">読み込み中</span>
            <svg className="animate-spin h-5 w-5 mr-2" viewBox="0 0 24 24">
              {/* スピナーのSVG */}
            </svg>
          </>
        ) : null}
        {children}
      </button>
    );
  }
);

Button.displayName = 'Button';
```

### 1.11.3. アクセシビリティのベストプラクティス

1. **セマンティックHTML**:
   - 正しい要素（`button`, `a`, `nav`, `article` など）を使用
   - 適切な見出しレベル（`h1` ~ `h6`）を使用

2. **フォーカス管理**:
   - すべてのインタラクティブ要素をキーボードでアクセス可能に
   - カスタムUIコンポーネントでのフォーカストラップ実装

3. **ARIAの適切な使用**:
   - `aria-label`, `aria-describedby`, `aria-expanded` などの属性を適切に使用
   - 動的コンテンツの変更を `aria-live` で通知

4. **色コントラスト**:
   - テキストと背景のコントラスト比を十分に確保
   - 色だけに依存しない情報伝達

## 1.12. 国際化と拡張性

### 1.12.1. 多言語対応

Next.jsでの国際化実装：

```text
app/
├── [lang]/           # 言語パラメータ
│   ├── dictionaries/ # 言語辞書
│   │   ├── en.json
│   │   └── ja.json
│   ├── layout.tsx    # 言語設定を適用
│   └── page.tsx      # 多言語対応ページ
└── middleware.ts     # 言語リダイレクト処理
```

```tsx
// app/[lang]/layout.tsx
import { getDictionary } from './dictionaries';

export default async function Layout({
  children,
  params: { lang },
}) {
  const dict = await getDictionary(lang);

  return (
    <html lang={lang}>
      <body>
        {/* dictをクライアントコンポーネントに渡す */}
        <ClientComponent dict={dict} />
        {children}
      </body>
    </html>
  );
}
```

### 1.12.2. 拡張性の高いコード設計

拡張性を考慮したモジュール設計：

1. **抽象化と依存性の注入**:

    ```tsx
    // services/api/interface.ts
    export interface ApiClient {
      get<T>(url: string, params?: Record<string, any>): Promise<T>;
      post<T>(url: string, data: any): Promise<T>;
      // その他のメソッド
    }

    // services/api/http-client.ts
    import { ApiClient } from './interface';

    export class HttpApiClient implements ApiClient {
      constructor(private baseUrl: string) {}

      async get<T>(url: string, params?: Record<string, any>): Promise<T> {
        // 実装
      }

      async post<T>(url: string, data: any): Promise<T> {
        // 実装
      }
    }

    // services/products/products-service.ts
    import { ApiClient } from '../api/interface';

    export class ProductsService {
      constructor(private apiClient: ApiClient) {}

      async getProducts() {
        return this.apiClient.get('/products');
      }
    }
    ```

2. **機能フラグとプラグイン設計**:

    ```tsx
    // lib/features.ts
    export const FEATURES = {
      NEW_CHECKOUT: process.env.FEATURE_NEW_CHECKOUT === 'true',
      WISHLIST: process.env.FEATURE_WISHLIST === 'true',
    };

    // components/Checkout.tsx
    import { FEATURES } from '@/lib/features';

    export function Checkout() {
      return FEATURES.NEW_CHECKOUT ? <NewCheckout /> : <LegacyCheckout />;
    }
    ```

### 1.12.3. ドキュメント戦略

コードの保守性と拡張性を向上させるドキュメント戦略：

```tsx
/**
 * 製品カードコンポーネント
 *
 * @example
 * <ProductCard product={product} onAddToCart={handleAddToCart} />
 *
 * @param props - コンポーネントプロパティ
 * @param props.product - 製品データ
 * @param props.onAddToCart - カートに追加時のコールバック
 */
function ProductCard({ product, onAddToCart }: ProductCardProps) {
  // 実装
}
```

Storybookによるコンポーネントカタログ：

```tsx
// Button.stories.tsx
import type { Meta, StoryObj } from '@storybook/react';
import { Button } from './Button';

const meta: Meta<typeof Button> = {
  component: Button,
  title: 'UI/Button',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof Button>;

export const Primary: Story = {
  args: {
    variant: 'primary',
    children: 'ボタン',
  },
};

export const Secondary: Story = {
  args: {
    variant: 'secondary',
    children: 'セカンダリ',
  },
};
```

このベストプラクティスガイドを活用することで、堅牢で保守性の高いNext.jsアプリケーションを構築できます。コードの品質、パフォーマンス、アクセシビリティ、セキュリティに配慮した実装を心がけましょう。
