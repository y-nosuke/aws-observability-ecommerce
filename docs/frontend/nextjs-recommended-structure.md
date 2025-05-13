# 1. Next.js 15 Container/Presentationalパターンの推奨ディレクトリ構成

## 1.1. 目次

- [1. Next.js 15 Container/Presentationalパターンの推奨ディレクトリ構成](#1-nextjs-15-containerpresentationalパターンの推奨ディレクトリ構成)
  - [1.1. 目次](#11-目次)
  - [1.2. 推奨ディレクトリ構成](#12-推奨ディレクトリ構成)
  - [1.3. ディレクトリ構造の詳細な説明](#13-ディレクトリ構造の詳細な説明)
    - [1.3.1. グループディレクトリ `(features)`](#131-グループディレクトリ-features)
    - [1.3.2. ルートパス (`/`) の扱い方](#132-ルートパス--の扱い方)
    - [1.3.3. Container/Presentationalパターンの実装](#133-containerpresentationalパターンの実装)
  - [1.4. テストコードの推奨配置場所](#14-テストコードの推奨配置場所)
    - [1.4.1. コンポーネントに隣接したテスト（推奨）](#141-コンポーネントに隣接したテスト推奨)
    - [1.4.2. プロジェクトルートのテスト](#142-プロジェクトルートのテスト)
  - [1.5. テスト戦略](#15-テスト戦略)
    - [1.5.1. Presentationalコンポーネントのテスト（Client Components）](#151-presentationalコンポーネントのテストclient-components)
    - [1.5.2. Presentationalコンポーネントのテスト（Server Components）](#152-presentationalコンポーネントのテストserver-components)
  - [1.6. ベストプラクティス](#16-ベストプラクティス)

## 1.2. 推奨ディレクトリ構成

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
│   ├── layout.tsx                   # ルートレイアウト
│   └── globals.css                  # グローバルスタイル
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
└── tests/                           # テスト関連
    ├── setup.ts                     # テスト設定
    └── mocks/                       # モックデータ
        └── products.ts              # 製品モックデータ
```

## 1.3. ディレクトリ構造の詳細な説明

### 1.3.1. グループディレクトリ `(features)`

- **括弧 `()` の意味**: Next.jsでは、括弧で囲まれたディレクトリ名はルーティングパスに含まれません。これを利用して機能をグループ化します。
- **URLマッピング例**:
  - `/` → `src/app/page.tsx` (ルートパス)
  - `/products` → `src/app/(features)/products/page.tsx`
  - `/products/[id]` → `src/app/(features)/products/[id]/page.tsx`

### 1.3.2. ルートパス (`/`) の扱い方

`src/app/page.tsx` はルートURL (`/`) にアクセスしたときに表示されるページです。
このファイルには、以下のような実装パターンがあります：

1. **ホームページコンポーネントのインポート** (推奨アプローチ)

   ```tsx
   // /src/app/page.tsx
   import HomePage from "./(features)/home/page";

   export default function Home() {
     return <HomePage />;
   }
   ```

   これにより、ホームページの実装が他の機能と同じパターンで `(features)/home` ディレクトリに格納され、一貫性のある構造を維持できます。

2. **リダイレクトパターン** (必要な場合のみ)

   ```tsx
   // 「/」→「/home」へリダイレクトが必要な場合
   import { redirect } from 'next/navigation';
   export default function Home() { redirect('/home'); }
   ```

### 1.3.3. Container/Presentationalパターンの実装

各機能ディレクトリは以下のパターンに従います：

1. **Container (page.tsx)**:
   - サーバーコンポーネント
   - データ取得ロジックを含む
   - Presentationalコンポーネントにデータを渡す

   ```tsx
   // src/app/(features)/products/page.tsx
   export default async function ProductsPage() {
     const products = await fetchProducts();  // データ取得
     return <ProductsClient products={products} />;  // データを渡す
   }
   ```

2. **Presentational (client.tsx)**:
   - クライアントコンポーネント (`"use client"` 宣言)
   - UIとユーザーインタラクションを処理
   - Containerから受け取ったデータを表示

   ```tsx
   // src/app/(features)/products/client.tsx
   "use client";

   export default function ProductsClient({ products }) {
     // クライアントサイドの状態管理、イベントハンドリング
     return (
       // ユーザーインターフェースの実装
     );
   }
   ```

この分離により、サーバーサイドの処理とクライアントサイドの処理を明確に区別でき、パフォーマンスとテスト容易性が向上します。

## 1.4. テストコードの推奨配置場所

次の2つのアプローチがあります：

### 1.4.1. コンポーネントに隣接したテスト（推奨）

```text
src/
├── app/
│   └── (features)/
│       └── products/
│           ├── client.tsx
│           ├── __tests__/                 # ディレクトリ単位のテスト
│           │   ├── client.test.tsx        # Presentationalコンポーネントのテスト
│           │   └── page.test.tsx          # Server Componentのテスト
│           └── components/
│               ├── ProductCard.tsx
│               └── __tests__/             # コンポーネント単位のテスト
│                   └── ProductCard.test.tsx
```

### 1.4.2. プロジェクトルートのテスト

大規模プロジェクトでは、以下のように分離することも有効：

```text
src/
└── tests/
    ├── unit/                         # ユニットテスト
    │   ├── components/               # コンポーネントテスト
    │   │   └── products/
    │   │       ├── client.test.tsx
    │   │       └── ProductCard.test.tsx
    │   └── lib/                      # ユーティリティテスト
    │       └── utils/
    │           └── date.test.ts
    ├── integration/                  # 統合テスト
    │   └── products/
    │       └── api.test.ts
    └── e2e/                          # E2Eテスト
        └── products/
            └── products.spec.ts
```

## 1.5. テスト戦略

### 1.5.1. Presentationalコンポーネントのテスト（Client Components）

```tsx
// client.test.tsx
import { render, screen, fireEvent } from '@testing-library/react';
import ProductsClient from '../client';

// モックデータ
const mockProducts = [
  { id: '1', name: 'Product 1', price: 100 },
  { id: '2', name: 'Product 2', price: 200 },
];

describe('ProductsClient', () => {
  it('商品リストを正しくレンダリングする', () => {
    render(<ProductsClient products={mockProducts} />);

    expect(screen.getByText('Product 1')).toBeInTheDocument();
    expect(screen.getByText('Product 2')).toBeInTheDocument();
  });

  it('ソートボタンをクリックすると商品が並び替えられる', () => {
    render(<ProductsClient products={mockProducts} />);

    // 名前順にソートされた状態を確認
    const items = screen.getAllByRole('listitem');
    expect(items[0]).toHaveTextContent('Product 1');
    expect(items[1]).toHaveTextContent('Product 2');

    // 価格順にソート
    fireEvent.click(screen.getByText('価格でソート'));

    // 価格順に並び替えられたことを確認
    const sortedItems = screen.getAllByRole('listitem');
    expect(sortedItems[0]).toHaveTextContent('Product 1');
    expect(sortedItems[1]).toHaveTextContent('Product 2');
  });
});
```

### 1.5.2. Presentationalコンポーネントのテスト（Server Components）

Server Componentsのテストは少し複雑です。Next.jsの`jest-environment-jsdom`環境では完全なテストは難しいため、以下のようなアプローチを取ります：

1. **ロジック部分のみを分離してテスト**：

    ```tsx
    // src/app/(features)/products/actions.ts
    export async function fetchProducts() {
      // データ取得ロジック
      return [...productsData];
    }

    // src/app/(features)/products/page.tsx
    import { ProductsList } from './client';
    import { fetchProducts } from './actions';

    export default async function ProductsPage() {
      const products = await fetchProducts();
      return <ProductsList products={products} />;
    }
    ```

    ```tsx
    // actions.test.ts
    import { fetchProducts } from '../actions';

    describe('Product Actions', () => {
      it('製品データを正しく取得する', async () => {
        const products = await fetchProducts();
        expect(products.length).toBeGreaterThan(0);
        expect(products[0]).toHaveProperty('id');
        expect(products[0]).toHaveProperty('name');
      });
    });
    ```

2. **統合テストまたはE2Eテスト**：

    Server Componentsの完全なテストはPlaywrightなどのE2Eテストツールで行うのが効果的です：

    ```typescript
    // tests/e2e/products/products.spec.ts
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

## 1.6. ベストプラクティス

1. **テスト優先の構成**:
   - Presentational（client.tsx）コンポーネントは純粋な関数として実装し、テスト容易性を確保
   - データ取得ロジックは独立した関数に分離（actions.ts等）

2. **モックの活用**:
   - 外部APIやデータベースアクセスはモック化
   - `jest.mock()`や`msw`を使用してAPIリクエストをモック

3. **テストカバレッジ優先順位**:
   1. Presentationalコンポーネント（UIとイベントハンドラ）
   2. ユーティリティ関数とカスタムフック
   3. データ取得ロジック（actions, services）
   4. E2Eテストで統合されたアプリケーションフロー

この構成により、テストしやすく保守性の高いContainer/Presentationalパターンの実装が可能になります。
