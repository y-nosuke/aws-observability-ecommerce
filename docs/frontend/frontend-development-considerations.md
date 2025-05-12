# Next.js 15のフロントエンド開発で考慮すべき重要事項

ディレクトリ構造とテスト戦略は基盤として重要ですが、フロントエンド開発の成功には他にも重要な検討事項があります。特にNext.js 15でContainer/Presentationalパターンを採用する場合の核心的な要素を解説します。

## 1. 状態管理戦略

### クライアント状態の扱い方

- **React Hooksの最適な活用**

  ```tsx
  // 単純な状態管理: useState
  const [filter, setFilter] = useState('all');

  // 複雑な状態管理: useReducer
  const [state, dispatch] = useReducer(productsReducer, initialState);

  // 最適化された小規模なコンテキスト
  const CartContext = createContext<CartContextType | null>(null);
  ```

- **状態管理ライブラリの選定基準**

  | ライブラリ    | 用途                     | Container/Presentationalとの親和性              |
  | ------------- | ------------------------ | ----------------------------------------------- |
  | Redux Toolkit | 大規模なグローバル状態   | Container層でストア設定、Presentational層で利用 |
  | Zustand       | シンプルで軽量な状態管理 | Client Componentsでの利用に最適                 |
  | Jotai/Recoil  | 原子的な状態管理         | 細分化されたPresentational間での状態共有        |

### サーバー状態との連携

- **React Query / SWRの活用**

  ```tsx
  'use client'

  import { useQuery } from '@tanstack/react-query';

  // Presentational層でのサーバーデータ活用
  function ProductsClient({ initialProducts }) {
    // 初期データはServer Component (Container)から提供
    // その後の更新はクライアント側でハンドリング
    const { data: products } = useQuery({
      queryKey: ['products'],
      queryFn: fetchProducts,
      initialData: initialProducts,
    });

    return (/* JSX */);
  }
  ```

## 2. スタイリング戦略

### スタイリング手法の選定

- **CSS-in-JSとCSS Modulesの使い分け**

  ```tsx
  // CSS Modules (Next.js標準サポート)
  import styles from './Button.module.css';

  function Button() {
    return <button className={styles.button}>Click me</button>;
  }
  ```

  ```tsx
  // styled-componentsの例
  'use client'

  import styled from 'styled-components';

  const StyledButton = styled.button`
    /* スタイル定義 */
  `;

  // 注意: styled-componentsはClient Componentsでのみ使用可能
  ```

- **コンポーネントライブラリの採用指針**
  - **shadcn/ui**: Server/Client分離を考慮した設計
  - **Tailwind CSS**: Server/Client両方で使用可能、バンドルサイズ最適化
  - **Chakra UI/MUI**: Client Components向け、リッチなコンポーネント

### デザインシステム構築

- **コンポーネント階層の設計**

  ```text
  components/
  ├── ui/               # 原子レベルのコンポーネント (Atoms)
  │   ├── button.tsx
  │   └── input.tsx
  ├── molecules/        # 複合コンポーネント (Molecules)
  │   └── search-bar.tsx
  └── organisms/        # 機能単位のコンポーネント (Organisms)
      └── product-card.tsx
  ```

## 3. パフォーマンス最適化

### Server/Client分割の最適化

- **最適なコンポーネント分割の例**

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

- **JavaScriptバンドルサイズ最適化**
  - Server Componentsを最大限活用して、クライアントバンドルを削減
  - コンポーネントの動的インポートと遅延ロード

### Next.js特有の最適化

- **Image Componentの適用**

  ```tsx
  import Image from 'next/image';

  // 最適化された画像表示
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

- **Route Segmentごとのコード分割**

  ```text
  app/
  ├── layout.tsx          # 共通レイアウト (重要: 最小限に保つ)
  ├── (marketing)/        # マーケティング関連ページグループ
  │   ├── page.tsx
  │   └── layout.tsx      # このグループ専用レイアウト
  └── (shop)/             # ショップ関連ページグループ
      ├── products/
      │   └── page.tsx
      └── layout.tsx      # ショップ専用レイアウト
  ```

## 4. API通信設計

### データフェッチング戦略

- **コンテナ層でのデータフェッチング**

  ```tsx
  // 推奨: Container層(page.tsx)での並列データフェッチング
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

- **効果的なエラーハンドリング**

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

### API層のモジュール化

- **サービス層の設計**

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

## 5. 型定義戦略（TypeScript）

### 厳格な型定義

- **Container/Presentational間の型定義**

  ```tsx
  // types/products.ts
  export interface Product {
    id: string;
    name: string;
    price: number;
    imageUrl: string;
    // その他のプロパティ
  }

  // Container->Presentationalへの型定義
  export interface ProductsClientProps {
    products: Product[];
    totalCount: number;
  }

  // client.tsx
  'use client'

  import type { ProductsClientProps } from '@/types/products';

  export function ProductsClient({ products, totalCount }: ProductsClientProps) {
    // 実装
  }
  ```

- **zodによるランタイムバリデーション**

  ```tsx
  import { z } from 'zod';

  // APIレスポンスのバリデーションスキーマ
  const productSchema = z.object({
    id: z.string(),
    name: z.string(),
    price: z.number().positive(),
    // 他のフィールド
  });

  type Product = z.infer<typeof productSchema>;

  // APIレスポンスのバリデーション
  async function fetchProduct(id: string): Promise<Product> {
    const response = await fetch(`/api/products/${id}`);
    const data = await response.json();

    // ランタイムバリデーション
    return productSchema.parse(data);
  }
  ```

## 6. エラーハンドリング戦略

### エラーUI設計

- **段階的なエラーバウンダリ**

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
      // 例: Sentryにエラーを送信
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

- **機能単位のエラーバウンダリ**

  ```tsx
  // app/features/products/error.tsx
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

## 7. セキュリティ考慮事項

### 認証・認可設計

- **NextAuth.js (Auth.js) の実装**

  ```tsx
  // app/api/auth/[...nextauth]/route.ts
  import NextAuth from 'next-auth';
  import { authOptions } from '@/lib/auth';

  const handler = NextAuth(authOptions);
  export { handler as GET, handler as POST };
  ```

- **Server Actionsでの権限チェック**

  ```tsx
  // actions.ts
  'use server'

  import { getServerSession } from 'next-auth';
  import { authOptions } from '@/lib/auth';
  import { revalidatePath } from 'next/cache';

  export async function createProduct(formData: FormData) {
    // 認証・認可チェック
    const session = await getServerSession(authOptions);
    if (!session || session.user.role !== 'admin') {
      throw new Error('Unauthorized');
    }

    // 処理実行
    const result = await db.products.create({
      // formDataの処理
    });

    // キャッシュ再検証
    revalidatePath('/products');

    return result;
  }
  ```

## 8. アクセシビリティ

### アクセシブルなコンポーネント設計

- **セマンティックHTML**

  ```tsx
  // 悪い例
  function ProductCard({ product }) {
    return (
      <div onClick={handleClick}>
        <div>{product.name}</div>
        <div>{product.price}</div>
      </div>
    );
  }

  // 良い例
  function ProductCard({ product }) {
    return (
      <article>
        <h3>{product.name}</h3>
        <p>
          <span className="sr-only">価格:</span>
          {product.price}円
        </p>
        <button
          onClick={handleClick}
          aria-label={`${product.name}の詳細を見る`}
        >
          詳細
        </button>
      </article>
    );
  }
  ```

- **フォーム要素のアクセシビリティ**

  ```tsx
  function SearchForm() {
    return (
      <form>
        <div>
          <label htmlFor="search">検索:</label>
          <input
            id="search"
            type="search"
            aria-describedby="search-hint"
          />
          <p id="search-hint" className="hint">
            商品名で検索できます
          </p>
        </div>
        <button type="submit">検索</button>
      </form>
    );
  }
  ```

## 9. 拡張性と保守性

### ドキュメント戦略

- **Storybookによるコンポーネントカタログ**

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

- **コードコメント規約**

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

## 10. 国際化(i18n)とローカライゼーション

### 多言語対応

- **ディレクトリベースのローカライゼーション**

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

- **言語切り替え機能**

  ```tsx
  'use client'

  import { usePathname, useRouter } from 'next/navigation';

  export function LanguageSwitcher({ currentLang }) {
    const pathname = usePathname();
    const router = useRouter();

    const switchLanguage = (newLang) => {
      // /ja/products -> /en/products のように言語パスを変更
      const newPath = pathname.replace(`/${currentLang}/`, `/${newLang}/`);
      router.push(newPath);
    };

    return (
      <select
        value={currentLang}
        onChange={(e) => switchLanguage(e.target.value)}
        aria-label="言語を選択"
      >
        <option value="ja">日本語</option>
        <option value="en">English</option>
      </select>
    );
  }
  ```

## まとめ

フロントエンド開発は、単なるコードの書き方だけでなく、上記の様々な要素をバランスよく考慮する必要があります。特にNext.js 15でContainer/Presentationalパターンを採用する場合、Server ComponentsとClient Componentsの責任境界を明確にしつつ、これらの要素を適切に取り入れることが成功の鍵となります。

プロジェクトの初期段階でこれらの事項について意思決定を行い、チーム全体で共有することで、開発の一貫性と効率が大幅に向上します。
