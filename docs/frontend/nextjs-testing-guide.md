# Next.js フロントエンドテスト環境構築ガイド

このガイドでは、AWS オブザーバビリティ学習用 eコマースアプリのフロントエンドテスト環境を構築する方法について説明します。

## 1. テスト環境のセットアップ

### 1.1. 必要なパッケージのインストール

プロジェクトのルートディレクトリで以下のコマンドを実行し、必要なパッケージをインストールします：

```bash
# テスト関連の依存パッケージをインストール
npm install --save-dev @testing-library/react @testing-library/jest-dom jest jest-environment-jsdom @types/jest identity-obj-proxy
```

### 1.2. Jest設定ファイルの作成

プロジェクトのルートディレクトリに `jest.config.js` ファイルを作成します。このファイルは Jest テストランナーの設定を行います。

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
    // SVGやCSSモジュールなどのアセットをモック
    '\\.(css|less|sass|scss)$': 'identity-obj-proxy',
    '\\.(jpg|jpeg|png|gif|svg)$': '<rootDir>/__mocks__/fileMock.js',
  },
  // カバレッジ設定
  collectCoverageFrom: [
    'app/**/*.{js,jsx,ts,tsx}',
    'components/**/*.{js,jsx,ts,tsx}',
    'lib/**/*.{js,jsx,ts,tsx}',
    '!**/*.d.ts',
    '!**/node_modules/**',
  ],
  // テスト実行時に無視するディレクトリ
  testPathIgnorePatterns: ['<rootDir>/node_modules/', '<rootDir>/.next/'],
};

// createJestConfigは非同期でNext.jsの設定を読み込む
module.exports = createJestConfig(customJestConfig);
```

### 1.3. Jest セットアップファイルの作成

プロジェクトのルートディレクトリに `jest.setup.js` ファイルを作成します。このファイルは各テストの前に実行され、グローバルな設定やモックを提供します。

```js
// jest-dom拡張をインポート
import '@testing-library/jest-dom';

// Next.jsの画像コンポーネントのモック
jest.mock('next/image', () => ({
  __esModule: true,
  default: (props) => {
    // eslint-disable-next-line jsx-a11y/alt-text
    return <img {...props} data-testid="next-image" />;
  },
}));

// Next.jsのリンクコンポーネントのモック
jest.mock('next/link', () => ({
  __esModule: true,
  default: ({ href, children, ...props }) => {
    return (
      <a href={href} {...props} data-testid="next-link">
        {children}
      </a>
    );
  },
}));

// Next.jsのRouterをモック
jest.mock('next/navigation', () => ({
  useRouter() {
    return {
      push: jest.fn(),
      back: jest.fn(),
      forward: jest.fn(),
      refresh: jest.fn(),
      replace: jest.fn(),
      prefetch: jest.fn(),
      events: {
        on: jest.fn(),
        off: jest.fn(),
        emit: jest.fn(),
      },
    };
  },
  usePathname() {
    return '/';
  },
  useSearchParams() {
    return new URLSearchParams();
  },
  useParams() {
    return {};
  },
}));

// matchMediaのモック（レスポンシブデザインのテスト用）
Object.defineProperty(window, 'matchMedia', {
  writable: true,
  value: jest.fn().mockImplementation(query => ({
    matches: false,
    media: query,
    onchange: null,
    addListener: jest.fn(),
    removeListener: jest.fn(),
    addEventListener: jest.fn(),
    removeEventListener: jest.fn(),
    dispatchEvent: jest.fn(),
  })),
});
```

### 1.4. ファイルモックの作成

`__mocks__` ディレクトリを作成し、その中に `fileMock.js` を追加します：

```bash
mkdir -p __mocks__
touch __mocks__/fileMock.js
```

`fileMock.js` の内容：

```js
// ファイルインポートのモック
module.exports = 'test-file-stub';
```

### 1.5. package.json にテスト用スクリプトを追加

`package.json` ファイルのスクリプトセクションに以下を追加します：

```json
"scripts": {
  "dev": "next dev",
  "build": "next build",
  "start": "next start",
  "lint": "next lint",
  "test": "jest",
  "test:watch": "jest --watch",
  "test:coverage": "jest --coverage"
}
```

## 2. テストディレクトリ構造

テストを整理するために、以下のような構造でテストファイルを配置します：

```
__tests__/
├── components/         # UIコンポーネントのテスト
├── pages/              # ページコンポーネントのテスト
└── lib/                # ユーティリティとAPIのテスト
    └── api/            # APIクライアントのテスト
```

これらのディレクトリを作成します：

```bash
mkdir -p __tests__/components __tests__/pages __tests__/lib/api
```

## 3. APIクライアントのテスト例

API呼び出しをテストするサンプルとして、`__tests__/lib/api/products.test.ts` を作成します：

```typescript
import { getProducts, getCategories } from '@/lib/api/products';

// globals.fetchのモック
global.fetch = jest.fn();

describe('ProductsAPI', () => {
  beforeEach(() => {
    // テスト前にモックをリセット
    jest.resetAllMocks();
  });

  describe('getProducts', () => {
    it('正しいURLとパラメータでリクエストを送信する', async () => {
      // モックレスポンスを設定
      const mockProductsResponse = {
        ok: true,
        json: jest.fn().mockResolvedValue({
          items: [],
          total_items: 0,
          page: 1,
          page_size: 10,
          total_pages: 0
        })
      };
      (global.fetch as jest.Mock).mockResolvedValue(mockProductsResponse);

      // 関数を実行
      await getProducts({ page: 2, page_size: 20, category_id: 1 });

      // アサーション
      expect(global.fetch).toHaveBeenCalledWith(
        expect.stringMatching(/\/products\?page=2&page_size=20&category_id=1/),
        expect.any(Object)
      );
    });

    // 他のテストケース...
  });
});
```

## 4. テストの実行

テストを実行するには、以下のコマンドを使用します：

```bash
# すべてのテストを実行
npm test

# ファイル名に基づいてテストを絞り込む
npm test -- -t "ProductsAPI"

# 特定のファイルのみテスト
npm test -- __tests__/lib/api/products.test.ts

# ウォッチモードでテストを実行（変更を監視して自動再実行）
npm run test:watch

# テストカバレッジレポートを生成
npm run test:coverage
```

Docker環境でテストを実行する場合：

```bash
docker compose exec frontend-customer npm test
```

## 5. App Router (React Server Components) でのテスト注意点

Next.js 13以降のApp Routerを使用している場合、React Server Components (RSC) のテストには制限があります：

1. **サーバーコンポーネントは直接テストできない**：
   - サーバーコンポーネントは、クライアントコンポーネントに分割してテストを行う
   - データ取得ロジックとUIを分離する

2. **'use client' ディレクティブ**：
   - テスト対象のコンポーネントファイルの先頭に `'use client'` を追加してクライアントコンポーネントとして明示する

3. **モック戦略**：
   - サーバーアクションやデータ取得関数は適切にモック化する
   - `jest.mock()` を使用して外部依存を置き換える

4. **統合テスト**：
   - 複雑なサーバーコンポーネントロジックは、エンドツーエンドテスト（Playwright, Cypressなど）で検証することを検討する

---

この環境設定ガイドは、プロジェクトのUIが安定した後でテストを追加する準備として使用できます。UIが頻繁に変更される初期段階では、まずはAPI呼び出しなどの比較的安定した部分からテストを始めることをお勧めします。
