# Next.jsフロントエンド開発入門 - AWSオブザーバビリティeコマースアプリ

## 目次

1. [Next.jsの基本概念](#1-nextjsの基本概念)
2. [プロジェクト構造の解説](#2-プロジェクト構造の解説)
3. [ルーティングシステム](#3-ルーティングシステム)
4. [コンポーネント設計](#4-コンポーネント設計)
5. [TailwindCSSによるスタイリング](#5-tailwindcssによるスタイリング)
6. [クライアントサイドとサーバーサイド](#6-クライアントサイドとサーバーサイド)
7. [今後の学習ステップ](#7-今後の学習ステップ)

## 1. Next.jsの基本概念

### Next.jsとは

Next.jsは、Reactをベースとした強力なフロントエンドフレームワークです。Reactの機能に加え、以下のような機能を提供します：

- **ファイルベースのルーティング**: ファイル構造がそのままURLルートになるシンプルな仕組み
- **サーバーサイドレンダリング(SSR)**: 初回読み込み時にサーバー側でHTMLを生成
- **静的サイト生成(SSG)**: ビルド時にHTMLを生成し、高速なページ表示を実現
- **APIルート**: サーバーサイド処理を同一プロジェクト内に統合
- **自動コード分割**: パフォーマンス最適化のため必要なJavaScriptのみ読み込み
- **画像最適化**: サイズ、フォーマット、品質を自動調整

これらの機能により、Next.jsはパフォーマンスに優れた、SEOフレンドリーなアプリケーションの開発に適しています。

### App Router vs Pages Router

Next.jsには2つのルーティングシステムがあります：

1. **App Router**: Next.js 13以降の新しいルーティングシステム（プロジェクトで採用）
   - `/app` ディレクトリを使用
   - Reactサーバーコンポーネントをサポート
   - レイアウト、エラーハンドリング、ローディングUIの改善
   - より宣言的なルーティングアプローチ

2. **Pages Router**: 従来のルーティングシステム
   - `/pages` ディレクトリを使用
   - より単純な構造だが、新機能のサポートは限定的
   - 使いやすいが拡張性が若干低い

このプロジェクトでは、より新しく柔軟性の高いApp Routerを採用しています。

## 2. プロジェクト構造の解説

プロジェクトは2つの独立したNext.jsアプリケーション（顧客向けと管理者向け）で構成されています。これにより、それぞれのインターフェースを独立して開発・デプロイすることができます。

### 主要なファイルとディレクトリ

```text
frontend-customer/  (または frontend-admin/)
├── app/                  # App Routerのルートディレクトリ
│   ├── layout.tsx        # 全ページ共通のレイアウト
│   ├── page.tsx          # ホームページ
│   ├── products/         # 商品関連ページ
│   │   └── page.tsx      # 商品一覧ページ
│   └── globals.css       # グローバルスタイル
├── src/
│   ├── components/       # 再利用可能なコンポーネント
│   │   ├── layout/       # レイアウト関連コンポーネント
│   │   └── ui/           # UI要素のコンポーネント
│   ├── lib/              # ユーティリティ関数やヘルパー
│   └── types/            # TypeScript型定義
├── public/               # 静的ファイル（画像など）
└── ...設定ファイル       # next.config.js, package.json など
```

この構造は、**関心の分離**と**責務の明確化**を重視しています。各ディレクトリが特定の役割を持ち、コードの整理と保守が容易になっています。

### 各ディレクトリの役割

- **app/**: Next.jsのルーティングと各ページの実装
- **src/components/**: 再利用可能なUIコンポーネント
  - **layout/**: ヘッダー、フッター、レイアウトなどの構造的コンポーネント
  - **ui/**: ボタン、カード、フォーム要素などの基本的なUI要素
- **src/lib/**: データフェッチ、ユーティリティ関数、ヘルパーなど
- **src/types/**: TypeScriptの型定義ファイル
- **public/**: 画像やフォントなどの静的ファイル

## 3. ルーティングシステム

### ファイルベースのルーティング

Next.jsのApp Routerでは、`app`ディレクトリ内のフォルダ構造がそのままURLになります：

| ファイルパス                 | URL                  |
| ---------------------------- | -------------------- |
| `app/page.tsx`               | `/`                  |
| `app/products/page.tsx`      | `/products`          |
| `app/products/[id]/page.tsx` | `/products/123` など |

各ディレクトリには`page.tsx`ファイルを配置することで、そのルートに対応するページコンポーネントを定義します。

### レイアウトの仕組み

`layout.tsx`ファイルは、そのディレクトリ以下の全てのページで共有されるレイアウトを定義します。ネストされたレイアウトも可能で、ルートレイアウト（`app/layout.tsx`）は必須です。

例えば、プロジェクトのルートレイアウト：

```tsx
// app/layout.tsx
import MainLayout from "@/components/layout/MainLayout";
import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Observability Shop | 高品質な商品をお求めやすい価格で",
  description: "オブザーバビリティショップでは、高品質な電子機器や最新ガジェットを、お求めやすい価格で提供しています。",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
        <MainLayout>{children}</MainLayout>
      </body>
    </html>
  );
}
```

この例では：

- フォントの設定（Geist, Geist Mono）
- メタデータの設定（タイトル、説明）
- 全ページ共通のMainLayoutコンポーネントの適用

### Linkコンポーネントによるナビゲーション

Next.jsでは、`<Link>`コンポーネントを使用してページ間のナビゲーションを行います：

```tsx
import Link from "next/link";

// 使用例
<Link href="/products" className="hover:text-indigo-200 transition-colors">
  商品一覧
</Link>
```

`<Link>`コンポーネントは内部的にクライアントサイドナビゲーションを実現し、ページ全体をリロードせずに遷移するため、スムーズなユーザー体験を提供します。

## 4. コンポーネント設計

### コンポーネントの分類

プロジェクトでは、コンポーネントを以下のように分類しています：

1. **レイアウトコンポーネント** (`components/layout/`)
   - `MainLayout.tsx`: 全体のレイアウト構造
   - `Header.tsx`: ヘッダー（ナビゲーション）
   - `Footer.tsx`: フッター

2. **UI要素コンポーネント** (`components/ui/`)
   - `ProductCard.tsx`: 商品表示カード
   - `AnimateInView.tsx`: 要素が画面に表示されたときのアニメーション効果

### コンポーネントの責務分離

良いコンポーネント設計の原則として、各コンポーネントは単一の責務を持つようにします：

- **レイアウトコンポーネント**: ページの構造と配置を担当
- **UI要素コンポーネント**: 個々のインターフェース要素を担当
- **ページコンポーネント**: 特定のURLルートに対応するコンテンツを担当

### クライアントコンポーネントとサーバーコンポーネント

Next.jsのApp Routerでは、デフォルトですべてのコンポーネントが**サーバーコンポーネント**になります。インタラクティブな機能（フックやイベントハンドラー）が必要な場合は、明示的に**クライアントコンポーネント**として宣言する必要があります：

```tsx
"use client"; // この指示子でクライアントコンポーネントとして宣言

import { useState } from "react";

export default function InteractiveComponent() {
  const [count, setCount] = useState(0);

  return (
    <button onClick={() => setCount(count + 1)}>
      クリック回数: {count}
    </button>
  );
}
```

例えば、プロジェクト内の`Header.tsx`は`"use client"`ディレクティブを使用しています（スクロール状態を監視するため）。

### コンポーネントの再利用性

`ProductCard.tsx`は再利用可能なコンポーネントの良い例です：

```tsx
// 商品カードコンポーネントの型定義
interface ProductCardProps {
  id: string;
  name: string;
  description: string;
  price: number;
  salePrice?: number | null;
  isNew?: boolean;
  imageUrl?: string | null;
}

// 使用例
<ProductCard
  id="1"
  name="超高性能ノートPC"
  description="最新のプロセッサ・16GB RAM・高速 SSD 搭載"
  price={198000}
  isNew={true}
/>
```

TypeScriptの型定義により、コンポーネントのAPIが明確になり、誤った使用を防止することができます。

## 5. TailwindCSSによるスタイリング

### TailwindCSSとは

TailwindCSSは、ユーティリティファーストのCSSフレームワークで、クラス名を直接HTMLに適用することでスタイリングを行います。このプロジェクトではTailwindCSSを採用しています。

### ユーティリティクラスの基本

TailwindCSSでは、以下のような形でスタイルを適用します：

```html
<div class="bg-white dark:bg-gray-800 shadow-md p-4 rounded-lg">
  <h3 class="text-xl font-bold mb-2 text-gray-800 dark:text-white">タイトル</h3>
  <p class="text-gray-600 dark:text-gray-300">説明文</p>
</div>
```

各クラスは特定のCSS機能に対応しています：

- `bg-white`: 背景色を白に
- `dark:bg-gray-800`: ダークモード時は背景色をgray-800に
- `p-4`: パディングを1rem (16px)に
- `rounded-lg`: 大きめの角丸に
- `text-xl`: フォントサイズを大きく
- `font-bold`: 太字に
- `mb-2`: 下マージンを0.5remに

### レスポンシブデザイン

TailwindCSSでは、接頭辞を使用してブレークポイントを指定できます：

```html
<div class="block md:flex lg:grid">
  <!--
    - スマホ（デフォルト）: block表示
    - タブレット (md:): flex表示
    - デスクトップ (lg:): grid表示
  -->
</div>
```

プロジェクトのヘッダーコンポーネントは、デスクトップとモバイルでの表示を切り替えています：

```tsx
{/* デスクトップメニュー */}
<nav className="hidden md:block">
  {/* メニュー内容 */}
</nav>

{/* モバイルメニューボタン */}
<button className="md:hidden">
  {/* ハンバーガーアイコン */}
</button>
```

### カスタムスタイルの定義

`globals.css`では、TailwindCSSの拡張やカスタム変数、独自のクラスなどを定義しています：

```css
:root {
  --background: #ffffff;
  --foreground: #171717;
  --primary: #4f46e5;
  --primary-light: #6366f1;
  --primary-dark: #4338ca;
  /* ... */
}

/* カードスタイル改善 */
.card-hover {
  transition: all 0.3s ease;
}

.card-hover:hover {
  transform: translateY(-6px);
  box-shadow: 0 12px 20px -8px rgba(var(--primary-dark), 0.3);
}

/* ボタン効果 */
.btn-primary {
  background: linear-gradient(to right, var(--primary), var(--primary-light));
  transition: all 0.3s ease;
}
```

## 6. クライアントサイドとサーバーサイド

### コンポーネントの実行環境

Next.jsのApp Routerでは、コンポーネントは2つの環境で実行されます：

1. **サーバーコンポーネント**（デフォルト）
   - サーバー上でのみ実行
   - データベースやファイルシステムに直接アクセス可能
   - APIキーなどの機密情報を安全に使用可能
   - クライアントサイドの状態や副作用（useStateなど）は使用不可

2. **クライアントコンポーネント**（"use client"ディレクティブを使用）
   - ブラウザで実行（ハイドレーション後）
   - React Hooksが使用可能（useState, useEffect等）
   - イベントハンドラ（onClick等）が利用可能
   - ブラウザAPIにアクセス可能（localStorage等）

### データフェッチの基本

Next.jsでは、サーバーコンポーネント内で直接データをフェッチすることができます：

```tsx
// app/products/page.tsx (サーバーコンポーネント)
async function getProducts() {
  const res = await fetch('https://api.example.com/products');
  if (!res.ok) throw new Error('Failed to fetch products');
  return res.json();
}

export default async function ProductsPage() {
  const products = await getProducts();

  return (
    <div>
      <h1>商品一覧</h1>
      <div className="grid grid-cols-4 gap-4">
        {products.map(product => (
          <ProductCard key={product.id} {...product} />
        ))}
      </div>
    </div>
  );
}
```

### クライアントサイドのインタラクション

ユーザーのインタラクションを処理するためには、クライアントコンポーネントを使用します：

```tsx
"use client";

import { useState } from 'react';

export default function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>カウント: {count}</p>
      <button onClick={() => setCount(count + 1)}>増加</button>
    </div>
  );
}
```

プロジェクト内の`AnimateInView.tsx`コンポーネントは、Intersection Observerを使用して要素が画面に表示されたときのアニメーションを制御するクライアントコンポーネントの例です。

## 7. 今後の学習ステップ

フロントエンド学習計画に基づき、次のステップとして学ぶべきトピックは以下の通りです：

### 短期的な学習目標（フェーズ1-2）

1. **データフェッチと非同期処理**
   - `fetch` APIやaxiosを使ったデータ取得
   - async/awaitパターン
   - ローディング状態とエラー処理

2. **SWRによるデータ管理**
   - データキャッシュと再検証
   - グローバル状態管理
   - リアルタイム更新

3. **エラーハンドリングとフォールバックUI**
   - エラー境界（Error Boundaries）
   - ユーザーフレンドリーなエラー表示
   - 機能の段階的劣化

4. **状態管理の基礎**
   - useState、useReducerの使いわけ
   - React Contextによる状態共有
   - 不変性（Immutability）の原則

### 中長期的な学習ロードマップ（フェーズ3-6）

1. **高度なUI実装とモニタリング**
   - クライアントサイドのログ収集
   - パフォーマンス計測
   - Web Vitalsとユーザー体験の最適化
   - スケルトンUIとローディング最適化

2. **フォームと注文処理UI**
   - フォーム実装の基礎
   - React Hook Formとバリデーション
   - 複数ステップフォームの設計
   - ユニットテスト

3. **管理機能UIと高度なパターン**
   - データテーブルと管理UI
   - モーダルとポップオーバー
   - ドラッグ＆ドロップ機能

4. **最適化とデプロイ**
   - Next.jsのビルド最適化
   - 本番環境デプロイとVercel活用

### 学習リソース

以下のリソースを活用して学習を進めることをお勧めします：

- [Next.js公式ドキュメント](https://nextjs.org/docs)
- [React公式ドキュメント](https://reactjs.org/docs/getting-started.html)
- [TailwindCSS公式ドキュメント](https://tailwindcss.com/docs)
- [TypeScript公式ドキュメント](https://www.typescriptlang.org/docs/)
- [SWR公式ドキュメント](https://swr.vercel.app/)
- [Next.js Learn](https://nextjs.org/learn) - 公式チュートリアル

これらの知識を順に習得していくことで、Next.jsを使った効果的なフロントエンド開発のスキルを身につけることができます。
