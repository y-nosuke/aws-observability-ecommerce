# 1. Next.js 15におけるContainer/Presentationalパターンガイド

## 1.1. 目次

- [1. Next.js 15におけるContainer/Presentationalパターンガイド](#1-nextjs-15におけるcontainerpresentationalパターンガイド)
  - [1.1. 目次](#11-目次)
  - [1.2. Container/Presentationalパターンの基本](#12-containerpresentationalパターンの基本)
    - [1.2.1. Presentationalコンポーネント（表示担当）](#121-presentationalコンポーネント表示担当)
    - [1.2.2. Containerコンポーネント（ロジック担当）](#122-containerコンポーネントロジック担当)
  - [1.3. Next.js 15とReact Server Components](#13-nextjs-15とreact-server-components)
    - [1.3.1. Server Components = Containerの自然な形](#131-server-components--containerの自然な形)
    - [1.3.2. Client Components = Presentationalの自然な形](#132-client-components--presentationalの自然な形)
  - [1.4. 実装パターン例](#14-実装パターン例)
    - [1.4.1. Containerコンポーネント (page.tsx)](#141-containerコンポーネント-pagetsx)
    - [1.4.2. Presentationalコンポーネント (presentation.tsx)](#142-presentationalコンポーネント-presentationtsx)
  - [1.5. 相性の良いディレクトリ構成](#15-相性の良いディレクトリ構成)
    - [1.5.1. 4.1 基本的なページ単位のパターン](#151-41-基本的なページ単位のパターン)
    - [1.5.2. 4.2 コンポーネント分離型](#152-42-コンポーネント分離型)
    - [1.5.3. 4.3 日本語資料で見られた実装パターン](#153-43-日本語資料で見られた実装パターン)
    - [1.5.4. 4.4 テストを考慮したパターン](#154-44-テストを考慮したパターン)
    - [1.5.5. 4.5 機能モジュール型](#155-45-機能モジュール型)
    - [1.5.6. 4.6 Next.js 15に最適化されたハイブリッドパターン](#156-46-nextjs-15に最適化されたハイブリッドパターン)
  - [1.6. srcディレクトリの使用について](#16-srcディレクトリの使用について)
    - [1.6.1. srcディレクトリを使用するメリット](#161-srcディレクトリを使用するメリット)
    - [1.6.2. srcディレクトリを使用した推奨構成](#162-srcディレクトリを使用した推奨構成)
  - [1.7. Next.js 15でのContainer/Presentationalパターンのメリット](#17-nextjs-15でのcontainerpresentationalパターンのメリット)
  - [1.8. 適用時の注意点](#18-適用時の注意点)

## 1.2. Container/Presentationalパターンの基本

Container/Presentationalパターンは、ReactアプリケーションにおいてロジックとUIを分離するための設計パターンです。このパターンでは、コンポーネントを2つの役割に分けます：

### 1.2.1. Presentationalコンポーネント（表示担当）

- UIの見た目のみに責任を持つ
- データをpropsを通じてのみ受け取る
- 基本的に状態を持たない（UIに関する状態のみ例外的に許容）
- 再利用しやすい純粋なコンポーネント

### 1.2.2. Containerコンポーネント（ロジック担当）

- データの取得やビジネスロジックを担当
- APIや状態管理と連携
- 取得したデータをPresentationalコンポーネントにpropsとして渡す

## 1.3. Next.js 15とReact Server Components

Next.js 15はReact 19に対応し、App Routerを通じてReact Server Components（RSC）を採用しています。React Server Componentsの導入により、Container/Presentationalパターンの適用方法も進化しました。

### 1.3.1. Server Components = Containerの自然な形

- サーバー側で実行され、データベースへの直接アクセスが可能
- APIからのデータ取得を行う
- 大きなライブラリを含められる（クライアントにJavaScriptを送信せずに済む）
- 機密情報を安全に扱える

### 1.3.2. Client Components = Presentationalの自然な形

- クライアント側で実行され、インタラクティブな要素を担当
- ファイルの先頭に`'use client'`ディレクティブを追加
- useStateやuseEffectなどのReact Hooksを利用可能
- UIの表示に専念

## 1.4. 実装パターン例

### 1.4.1. Containerコンポーネント (page.tsx)

```tsx
// Server Componentとしてデータ取得を担当
import { ProductsList } from './presentation';
import { fetchProducts } from '@/lib/api';

export default async function ProductsPage() {
  // サーバーサイドでデータ取得
  const products = await fetchProducts();

  // Presentationalコンポーネントにデータを渡す
  return <ProductsList products={products} />;
}
```

### 1.4.2. Presentationalコンポーネント (presentation.tsx)

```tsx
'use client'; // Client Componentを明示

import { useState } from 'react';

type Product = {
  id: string;
  name: string;
  price: number;
};

type ProductsListProps = {
  products: Product[];
};

export function ProductsList({ products }: ProductsListProps) {
  const [sortBy, setSortBy] = useState<'name' | 'price'>('name');

  // UIのみに関するロジックはここに記述
  const sortedProducts = [...products].sort((a, b) => {
    if (sortBy === 'name') return a.name.localeCompare(b.name);
    return a.price - b.price;
  });

  return (
    <div>
      <div className="controls">
        <button onClick={() => setSortBy('name')}>名前でソート</button>
        <button onClick={() => setSortBy('price')}>価格でソート</button>
      </div>
      <ul>
        {sortedProducts.map(product => (
          <li key={product.id}>
            {product.name} - ¥{product.price}
          </li>
        ))}
      </ul>
    </div>
  );
}
```

## 1.5. 相性の良いディレクトリ構成

### 1.5.1. 4.1 基本的なページ単位のパターン

```text
src/app/
├── features/
│   └── products/
│       ├── page.tsx          # Server Component (Container)
│       ├── presentation.tsx  # Client Component (Presentational)
│       └── loading.tsx       # ローディング状態のUI
```

### 1.5.2. 4.2 コンポーネント分離型

```text
src/app/
├── features/
│   └── products/
│       ├── page.tsx                  # Server Component (Container)
│       ├── components/
│       │   ├── ProductsList.tsx      # Client Component (Presentational)
│       │   ├── ProductsFilter.tsx    # Client Component
│       │   └── ProductCard.tsx       # Client Component
│       └── loading.tsx
```

### 1.5.3. 4.3 日本語資料で見られた実装パターン

```text
src/app/
├── (application)/            # Route Group (URLに影響しない)
│   └── products/
│       ├── page.tsx          # Container (サーバーコンポーネント)
│       ├── presentation.tsx  # Presentation (クライアントコンポーネント)
│       └── loading.tsx
```

### 1.5.4. 4.4 テストを考慮したパターン

```text
src/app/
├── features/
│   └── products/
│       ├── page.tsx          # Server Component (Container)
│       ├── presentation.tsx  # Client Component (Presentational)
│       └── __tests__/        # テストディレクトリ
│           └── presentation.test.tsx  # Presentationalコンポーネントのテスト
```

React Testing LibraryはServer Componentsに対応していないため、テスト容易性を考慮すると、presentation.tsxは純粋なReactコンポーネントとして実装することが重要です。

### 1.5.5. 4.5 機能モジュール型

```text
src/
├── app/
│   └── products/
│       ├── page.tsx        # ルートコンポーネント（Container）
│       └── loading.tsx     # ローディングUI
├── features/
│   └── products/
│       ├── containers/     # Containerコンポーネント
│       │   └── ProductsContainer.tsx
│       ├── components/     # Presentationalコンポーネント
│       │   ├── ProductsList.tsx
│       │   └── ProductCard.tsx
│       ├── hooks/          # カスタムフック
│       │   └── useProducts.ts
│       └── api/            # API関連
│           └── products.ts
└── components/             # 共通コンポーネント
    └── ui/                 # 共通UIコンポーネント
```

### 1.5.6. 4.6 Next.js 15に最適化されたハイブリッドパターン

```text
src/app/
├── api/                  # APIルート
├── (groups)/             # ページをグループ化
│   ├── admin/            # 管理画面関連
│   │   └── products/
│   │       ├── page.tsx               # Server Component (Container)
│   │       ├── client.tsx             # メインのClient Component
│   │       ├── components/            # 小さなPresentational Components
│   │       │   ├── ProductForm.tsx
│   │       │   └── ProductList.tsx
│   │       └── actions.ts             # Server Actions
│   └── shop/             # 一般ユーザー向け画面
├── components/           # 共通コンポーネント
└── lib/                  # ユーティリティ関数やヘルパー
```

## 1.6. srcディレクトリの使用について

Container/Presentationalパターンを適用する場合、srcディレクトリの使用をお勧めします。

### 1.6.1. srcディレクトリを使用するメリット

1. **ソースコードの整理**:
   - ソースコードと設定ファイルを明確に分離
   - 大規模なプロジェクトでルートディレクトリが煩雑になるのを防止

2. **拡張性の向上**:
   - アプリケーションが成長したとき、より構造化された組織化が可能
   - Container/Presentationalパターンの責任分離をサポート

3. **一般的な慣習との整合性**:
   - 多くのReactプロジェクトでsrcディレクトリが標準的
   - チーム間での共通理解が容易

4. **コードベースのクリーンさ**:
   - テストファイル、型定義、ユーティリティなどを整理しやすい

### 1.6.2. srcディレクトリを使用した推奨構成

```text
/
├── src/
│   ├── app/                  # Next.js App Router
│   │   ├── products/
│   │   │   ├── page.tsx      # Server Component (Container)
│   │   │   ├── client.tsx    # Client Component (Presentational)
│   │   │   └── loading.tsx
│   ├── components/           # 共通コンポーネント
│   │   ├── ui/               # 再利用可能なUIコンポーネント
│   │   └── layout/           # レイアウト関連コンポーネント
│   ├── lib/                  # ユーティリティとヘルパー関数
│   ├── types/                # TypeScript型定義
│   ├── hooks/                # カスタムフック
│   └── services/             # APIクライアント、外部サービス連携
├── public/                   # 静的アセット
├── next.config.js            # Next.js設定
├── package.json
├── tsconfig.json
└── ...その他の設定ファイル
```

## 1.7. Next.js 15でのContainer/Presentationalパターンのメリット

1. **サーバーとクライアントの責務明確化**:
   - Server Components（Container）はデータ取得とロジックを担当
   - Client Components（Presentational）はUIとユーザーインタラクションを担当

2. **パフォーマンス向上**:
   - JavaScriptバンドルサイズの削減（Server Componentsのコードはクライアントに送信されない）
   - データ取得の最適化（サーバーサイドで実行）

3. **コード再利用性の向上**:
   - Presentationalコンポーネントは純粋なUIとして再利用可能
   - コンポーネントの責務が明確に分離される

4. **テスト容易性**:
   - UIとビジネスロジックが分離されているため、テストが書きやすい

## 1.8. 適用時の注意点

1. **過剰な分離を避ける**:
   - 小規模なコンポーネントではシンプルな構成が望ましい場合も

2. **命名の一貫性**:
   - `presentation.tsx`、`client.tsx`、`view.tsx`など、Presentationalコンポーネントの命名は統一する

3. **ファイル数のバランス**:
   - 小規模なプロジェクトでは2〜3のファイル構成
   - 大規模プロジェクトでは機能モジュール型の構成が適している

どのディレクトリ構成を選ぶかは、プロジェクトの規模、チームの好み、開発の複雑さによって異なります。最も重要なのは、選んだ構成を一貫して適用することです。
