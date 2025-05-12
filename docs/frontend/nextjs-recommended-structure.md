# Next.js 15 Container/Presentationalパターンの推奨ディレクトリ構成

## 推奨ディレクトリ構成

```text
src/
├── app/                             # Next.js App Router
│   ├── features/                    # 機能別ディレクトリ
│   │   └── products/                # 製品機能
│   │       ├── page.tsx             # Server Component (Container)
│   │       ├── client.tsx           # メインのPresentational
│   │       ├── loading.tsx          # ローディングUI
│   │       └── components/          # 機能固有の小さなコンポーネント
│   │           ├── ProductCard.tsx
│   │           └── ProductFilter.tsx
│   ├── api/                         # API Routes
│   │   └── products/
│   │       └── route.ts
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
└── tests/                           # テスト関連
    ├── setup.ts                     # テスト設定
    └── mocks/                       # モックデータ
        └── products.ts              # 製品モックデータ
```

## テストコードの推奨配置場所

次の2つのアプローチがあります：

### 1. コンポーネントに隣接したテスト（推奨）

```text
src/
├── app/
│   └── features/
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

### 2. プロジェクトルートのテスト

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

## テスト戦略

### Presentationalコンポーネントのテスト（Client Components）

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

### Presentationalコンポーネントのテスト（Server Components）

Server Componentsのテストは少し複雑です。Next.jsの`jest-environment-jsdom`環境では完全なテストは難しいため、以下のようなアプローチを取ります：

1. **ロジック部分のみを分離してテスト**：

    ```tsx
    // src/app/features/products/actions.ts
    export async function fetchProducts() {
      // データ取得ロジック
      return [...productsData];
    }

    // src/app/features/products/page.tsx
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
      await page.goto('/features/products');

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

## ベストプラクティス

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
