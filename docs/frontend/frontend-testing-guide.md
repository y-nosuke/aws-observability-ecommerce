# Next.jsフロントエンドのテスト作成ガイド

このガイドではAWSオブザーバビリティ学習用eコマースアプリケーションのNext.jsフロントエンドのテスト方法について説明します。

## 1. テスト環境のセットアップ

### 1.1. 必要なパッケージのインストール

Next.jsプロジェクトでテストを実行するために、以下のパッケージをインストールします：

```bash
# テスト関連の依存パッケージをインストール
npm install --save-dev @testing-library/react @testing-library/jest-dom jest jest-environment-jsdom
```

### 1.2. Jest設定ファイルの作成

プロジェクトのルートディレクトリに`jest.config.js`ファイルを作成します：

```js
const nextJest = require('next/jest');

const createJestConfig = nextJest({
  // テスト環境のNext.jsアプリへのパスを指定
  dir: './',
});

// Jestに渡すカスタム設定
const customJestConfig = {
  // jsdom環境をテスト環境に追加
  testEnvironment: 'jest-environment-jsdom',
  // テスト対象ディレクトリを指定
  testMatch: ['**/__tests__/**/*.test.{js,jsx,ts,tsx}'],
  // セットアップファイル
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  // モジュールのモック
  moduleNameMapper: {
    // aliasのマッピング
    '^@/(.*)$': '<rootDir>/$1',
  },
};

// createJestConfigは非同期でNext.jsの設定を読み込む
module.exports = createJestConfig(customJestConfig);
```

### 1.3. Jest setup ファイルの作成

プロジェクトのルートディレクトリに`jest.setup.js`ファイルを作成します：

```js
// jest-dom拡張をインポート
import '@testing-library/jest-dom';

// Next.jsの画像コンポーネントのモック
jest.mock('next/image', () => ({
  __esModule: true,
  default: (props) => {
    // eslint-disable-next-line jsx-a11y/alt-text
    return <img {...props} />;
  },
}));

// Next.jsのRouterをモック
jest.mock('next/navigation', () => ({
  useRouter() {
    return {
      push: jest.fn(),
      back: jest.fn(),
      events: {
        on: jest.fn(),
        off: jest.fn(),
      },
    };
  },
  usePathname() {
    return '/';
  },
  useSearchParams() {
    return new URLSearchParams();
  },
}));
```

### 1.4. package.jsonにテストスクリプトを追加

```json
"scripts": {
  "test": "jest",
  "test:watch": "jest --watch",
  "test:coverage": "jest --coverage"
}
```

## 2. テスト構造

テストは`__tests__`ディレクトリに整理します。以下のような構造を推奨します：

```
__tests__/
├── components/           # コンポーネントのテスト
│   ├── ProductCard.test.tsx
│   ├── Pagination.test.tsx
│   └── ...
├── pages/                # ページコンポーネントのテスト
│   ├── ProductsPage.test.tsx
│   └── ...
└── lib/                  # ユーティリティ関数のテスト
    └── api/
        └── products.test.ts
```

## 3. テスト作成のガイドライン

### 3.1. コンポーネントテスト

コンポーネントテストでは、以下の内容を確認します：

1. コンポーネントが正しくレンダリングされるか
2. プロップスが正しく表示されるか
3. ユーザーインタラクション（クリックなど）が正しく機能するか
4. 条件付きレンダリングが正しく機能するか

### 3.2. ページテスト

ページコンポーネントのテストでは以下を確認します：

1. ページが正しくレンダリングされるか
2. サブコンポーネントが正しく統合されているか
3. データ取得ロジックが正しく機能するか（モックを使用）

### 3.3. APIクライアントテスト

APIクライアントのテストでは以下を確認します：

1. 正しいURLでAPIが呼び出されるか
2. パラメータが正しく処理されるか
3. レスポンスが正しく処理されるか
4. エラーハンドリングが適切か

## 4. モックの活用

テストではモックを活用して外部依存を置き換えます：

- `fetch`のモックで実際のAPIリクエストを回避
- 画像やルーターなどのNext.js特有の機能をモック
- 外部サービスやライブラリをモック

### fetch APIのモック例

```javascript
// fetch のグローバルモック
global.fetch = jest.fn();

// レスポンスのモック
const mockResponse = {
  json: jest.fn().mockResolvedValue({ /* モックデータ */ }),
  ok: true,
};

// fetch実装のモック
global.fetch.mockResolvedValue(mockResponse);
```

## 5. テスト実行

テストを実行するには、以下のコマンドを使用します：

```bash
# すべてのテストを実行
npm test

# テストをウォッチモードで実行（変更を監視して自動再実行）
npm run test:watch

# カバレッジレポートを生成
npm run test:coverage
```

テストはDockerコンテナ内でも実行できます：

```bash
docker compose exec frontend-customer npm test
```

## 6. テストケースの例

### ProductCard.test.tsx の例

```tsx
import { render, screen } from '@testing-library/react';
import ProductCard from '@/components/ProductCard';

describe('ProductCard', () => {
  const mockProduct = {
    id: 1,
    name: 'テスト商品',
    description: '商品の説明文',
    price: 1000,
    image_url: 'https://example.com/image.jpg',
    category_id: 1
  };

  it('商品名が正しく表示される', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText('テスト商品')).toBeInTheDocument();
  });

  it('商品価格が正しくフォーマットされる', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText('¥1,000')).toBeInTheDocument();
  });
});
```

### API Client Test の例

```tsx
import { getProducts } from '@/lib/api/products';

// fetchのモック
global.fetch = jest.fn();

describe('products API', () => {
  beforeEach(() => {
    jest.resetAllMocks();
  });

  it('getProducts が正しいURLとパラメータでリクエストを送信する', async () => {
    // モックレスポンスの設定
    const mockResponse = {
      ok: true,
      json: jest.fn().mockResolvedValue({
        items: [],
        total_items: 0,
        page: 1,
        page_size: 10,
        total_pages: 0
      })
    };
    
    (global.fetch as jest.Mock).mockResolvedValue(mockResponse);

    // 関数の実行
    await getProducts({ page: 2, page_size: 20, category_id: 1 });

    // Assertioaｐn
    expect(global.fetch).toHaveBeenCalledWith(
      expect.stringContaining('/products?page=2&page_size=20&category_id=1'),
      expect.any(Object)
    );
  });
});
```

## 7. テストにおける注意点

1. **環境変数の扱い**: テスト実行時に適切な環境変数が設定されているか確認する
2. **Next.jsの特殊コンポーネント**: `Link`や`Image`などのNext.js特有のコンポーネントを適切にモック化する
3. **サーバーコンポーネント**: Next.js 13/14のApp Routerを使用している場合、サーバーコンポーネントのテストには制限があるため、テスト可能なクライアントコンポーネントに分割することを検討する
4. **ビルド時エラー**: テスト環境でビルドエラーが発生する場合は、適切なモックとPolyfillを提供する

これらのガイドラインに従って、信頼性の高いテストを実装してください。
