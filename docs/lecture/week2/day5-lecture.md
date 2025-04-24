# Week 2 - Day 5: Reactの基本概念とTypeScriptの型定義

## 1.1. 目次

- [Week 2 - Day 5: Reactの基本概念とTypeScriptの型定義](#week-2---day-5-reactの基本概念とtypescriptの型定義)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. JSXの基本構文](#141-jsxの基本構文)
    - [1.4.2. コンポーネントの作成](#142-コンポーネントの作成)
    - [1.4.3. TypeScriptによる型定義](#143-typescriptによる型定義)
    - [1.4.4. Propsの受け渡し](#144-propsの受け渡し)
    - [1.4.5. コンポーネントライフサイクルとフック](#145-コンポーネントライフサイクルとフック)
    - [1.4.6. イベントハンドリング](#146-イベントハンドリング)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. 関数コンポーネントとクラスコンポーネントの比較](#161-関数コンポーネントとクラスコンポーネントの比較)
    - [1.6.2. TypeScriptによる型安全性の向上](#162-typescriptによる型安全性の向上)
    - [1.6.3. Next.jsとReactの関係](#163-nextjsとreactの関係)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. コンポーネント設計の原則](#171-コンポーネント設計の原則)
    - [1.7.2. パフォーマンス最適化テクニック](#172-パフォーマンス最適化テクニック)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: Propsが正しく渡されない](#181-問題1-propsが正しく渡されない)
    - [1.8.2. 問題2: 無限レンダリングループ](#182-問題2-無限レンダリングループ)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)

## 1.2. 【要点】

- JSXはHTMLに似た構文でUIを宣言的に記述する方法で、Reactコンポーネントの基本
- コンポーネントは再利用可能なUI要素で、関数またはクラスとして定義可能
- TypeScriptはJavaScriptに型システムを追加し、開発時のエラー検出を強化
- Propsはコンポーネント間でデータを受け渡す手段として機能
- フック（useState, useEffect等）は関数コンポーネントでの状態管理と副作用の実装に使用
- イベントハンドリングはユーザー操作に対して反応するための仕組み

## 1.3. 【準備】

この講義では、すでに構築済みのプロジェクト（frontend-adminとfrontend-customer）のコードを解説しながらReactの基本概念を学びます。以下のツールと環境が準備されていることを確認してください。

### 1.3.1. チェックリスト

- [ ] Node.js と npm/yarn がインストールされている
- [ ] コードエディタ（VS Code推奨）がインストールされている
- [ ] プロジェクトのコードが最新の状態に更新されている
- [ ] TypeScriptの基本知識がある
- [ ] HTMLとCSSの基本知識がある

## 1.4. 【手順】

### 1.4.1. JSXの基本構文

JSX（JavaScript XML）は、Reactで使用される構文拡張で、HTMLに似た形式でUIを記述できます。これによりUIの構造が視覚的に理解しやすくなります。

**例: DashboardCard.tsx**

```tsx
export default function DashboardCard({
  title,
  icon,
  children,
  className = "",
}: DashboardCardProps) {
  return (
    <div
      className={`bg-white dark:bg-gray-800 rounded-lg shadow-lg overflow-hidden hover-lift transition-smooth card-accent-primary ${className}`}
    >
      <div className="px-6 py-4 border-b border-gray-200 dark:border-gray-700 flex items-center">
        {icon && (
          <span className="mr-3 text-primary dark:text-primary-light">
            {icon}
          </span>
        )}
        <h3 className="font-semibold text-gray-700 dark:text-gray-200">
          {title}
        </h3>
      </div>
      <div className="p-6">{children}</div>
    </div>
  );
}
```

**JSXの主要な特徴:**

1. **HTMLに似た構文**: `<div>`, `<span>`, `<h3>`などのタグを使用
2. **JavaScript式の埋め込み**: 中括弧`{}`内にJavaScriptの式を書ける
3. **条件付きレンダリング**: `{icon && (...)}` のように条件分岐が可能
4. **クラス指定**: `className`属性でCSSクラスを指定（HTMLの`class`と異なる点に注意）
5. **属性の動的な設定**: `className={...}` のように属性値も動的に設定可能

### 1.4.2. コンポーネントの作成

Reactでは、UIを再利用可能なコンポーネントに分割します。各コンポーネントは独立して動作し、必要に応じて組み合わせて複雑なUIを構築できます。

**関数コンポーネントの例: ProductCard.tsx**

```tsx
export default function ProductCard({
  id,
  name,
  description,
  price,
  salePrice = null,
  isNew = false,
  imageUrl = null,
}: ProductCardProps) {
  return (
    <div className="product-card bg-white dark:bg-gray-800 shadow-md">
      {isNew && <div className="sale-badge">新着</div>}
      {/* 省略 */}
      <div className="p-4">
        <Link href={`/products/${id}`} className="block">
          <h3 className="font-semibold text-lg mb-1 hover:text-primary transition-colors">
            {name}
          </h3>
        </Link>
        {/* 省略 */}
      </div>
    </div>
  );
}
```

**コンポーネントの主要な特徴:**

1. **再利用性**: 同じコンポーネントを複数の場所で使用可能
2. **独立性**: 各コンポーネントは独立して動作
3. **カプセル化**: コンポーネント内部のロジックと表示を隠蔽
4. **合成**: 複数のコンポーネントを組み合わせて複雑なUIを構築

### 1.4.3. TypeScriptによる型定義

TypeScriptを使用すると、コンポーネントが受け取るpropsに型を定義できます。これにより、開発時のエラー検出やコード補完が強化されます。

**例: DashboardCard.tsxのProps型定義**

```tsx
type DashboardCardProps = {
  title: string;
  icon?: ReactNode;
  children: ReactNode;
  className?: string;
};
```

**型定義の主要な特徴:**

1. **必須プロパティ**: `title`と`children`は必須
2. **オプショナルプロパティ**: `icon`と`className`は`?`で省略可能に
3. **複雑な型**: `ReactNode`のような複合型の使用
4. **デフォルト値**: `className = ""`のように関数引数でデフォルト値を設定

**例: 認証関連の型定義 (auth.ts)**

```typescript
export interface User {
  id: string;
  name: string;
  email: string;
  role: string;
}

export interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  loading: boolean;
  error: string | null;
}

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}
```

### 1.4.4. Propsの受け渡し

Propsは親コンポーネントから子コンポーネントへデータを渡すための仕組みです。

**例: AdminLayout.tsxでの子コンポーネントへのProps渡し**

```tsx
export default function AdminLayout({ children }: AdminLayoutProps) {
  const [sidebarOpen, setSidebarOpen] = useState(true);

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  return (
    <div className="flex flex-col min-h-screen bg-gray-50 dark:bg-gray-900">
      <AdminHeader toggleSidebar={toggleSidebar} />
      <div className="flex flex-1">
        <AdminSidebar isOpen={sidebarOpen} />
        <main
          className={`flex-1 p-6 md:p-8 overflow-x-auto transition-all duration-300 ${
            !sidebarOpen ? "pl-4" : "pl-6"
          }`}
        >
          {children}
        </main>
      </div>
    </div>
  );
}
```

**Propsの受け渡しの主要な特徴:**

1. **関数の引数として受け取る**: `({ children }: AdminLayoutProps)` のように分割代入で受け取る
2. **子コンポーネントに渡す**: `<AdminHeader toggleSidebar={toggleSidebar} />` のように属性として渡す
3. **子要素としての渡し方**: `{children}` として子要素を配置
4. **計算値の渡し方**: 式や計算結果を中括弧 `{}` で囲んで渡す

### 1.4.5. コンポーネントライフサイクルとフック

フック（Hooks）は関数コンポーネントで状態管理やライフサイクル機能を使用するための仕組みです。

**例: Header.tsxでのuseStateとuseEffectの使用**

```tsx
export default function Header() {
  const [isScrolled, setIsScrolled] = useState(false);
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  // スクロール検出
  useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > 10) {
        setIsScrolled(true);
      } else {
        setIsScrolled(false);
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  // 以下省略
}
```

**フックの主要な特徴:**

1. **useState**: コンポーネントに状態（state）を追加するフック
   - `const [isScrolled, setIsScrolled] = useState(false);` で初期値`false`の状態を作成
   - `setIsScrolled(true)`で状態を更新

2. **useEffect**: 副作用（side effects）を実行するためのフック
   - コンポーネントのレンダリング後に実行される
   - 第2引数の依存配列（`[]`）が空なので、コンポーネントのマウント時に1度だけ実行
   - クリーンアップ関数を返すことで、アンマウント時にイベントリスナーを削除

3. **依存配列**: useEffectの第2引数で、この配列内の値が変わった時だけ再実行される

**例: AdminSidebar.tsxでのuseEffectとuseStateの使用**

```tsx
export default function AdminSidebar({ isOpen = true }: AdminSidebarProps) {
  const pathname = usePathname();
  const [currentTime, setCurrentTime] = useState("");

  useEffect(() => {
    // クライアント側でのみ実行されるように
    setCurrentTime(new Date().toLocaleString("ja-JP"));

    // 1分ごとに更新
    const timer = setInterval(() => {
      setCurrentTime(new Date().toLocaleString("ja-JP"));
    }, 60000);

    return () => clearInterval(timer);
  }, []);

  // 以下省略
}
```

### 1.4.6. イベントハンドリング

Reactでは、イベントハンドリングはDOMイベントに似た方法で処理されますが、いくつかの違いがあります。

**例: AdminHeader.tsxでのイベントハンドリング**

```tsx
export default function AdminHeader({ toggleSidebar }: AdminHeaderProps) {
  const [isProfileMenuOpen, setIsProfileMenuOpen] = useState(false);
  const userData = getUserData() || { name: "Admin User" };

  const toggleProfileMenu = () => {
    setIsProfileMenuOpen(!isProfileMenuOpen);
  };

  return (
    <header className="bg-gradient-primary text-white shadow-lg">
      {/* 省略 */}
      <div className="relative">
        <button
          onClick={toggleProfileMenu}
          className="flex items-center space-x-2 focus:outline-none hover:opacity-80 transition-smooth px-2 py-1 rounded-full"
        >
          {/* ボタンの内容省略 */}
        </button>

        {isProfileMenuOpen && (
          <div className="absolute right-0 mt-2 w-52 bg-white dark:bg-gray-800 rounded-xl shadow-xl py-2 text-gray-800 dark:text-gray-200 z-10 border border-gray-100 dark:border-gray-700 overflow-hidden">
            {/* メニュー内容省略 */}
            <button
              onClick={() => {
                // ログアウト処理
                if (typeof window !== "undefined") {
                  window.location.href = "/admin/login";
                }
              }}
              className="flex w-full items-center px-4 py-2.5 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors group text-left"
            >
              {/* ボタン内容省略 */}
            </button>
          </div>
        )}
      </div>
    </header>
  );
}
```

**イベントハンドリングの主要な特徴:**

1. **イベントハンドラの定義**: `toggleProfileMenu`のような関数を定義
2. **イベントのバインド**: `onClick={toggleProfileMenu}`のようにJSXタグに属性としてバインド
3. **インラインハンドラ**: `onClick={() => { /* 処理 */ }}`のようにインラインでも定義可能
4. **条件付きレンダリングとの組み合わせ**: `{isProfileMenuOpen && (<div>...</div>)}`のようにメニューの表示/非表示を切り替え

## 1.5. 【確認ポイント】

この講義の内容をマスターしたかを確認するためのチェックリストです：

- [ ] JSXの構文を理解し、基本的なコンポーネントを解読できる
- [ ] TypeScriptを使った型定義の書き方を理解している
- [ ] Propsの渡し方と受け取り方を理解している
- [ ] Reactフック（useState, useEffect）の基本的な使い方を理解している
- [ ] イベントハンドリングの方法を理解している
- [ ] 条件付きレンダリングの方法を理解している
- [ ] コンポーネントの合成方法を理解している

## 1.6. 【詳細解説】

### 1.6.1. 関数コンポーネントとクラスコンポーネントの比較

現在のReactでは関数コンポーネントが推奨されていますが、過去のコードではクラスコンポーネントも見かけます。

**関数コンポーネント:**

```tsx
function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}
```

**クラスコンポーネント:**

```tsx
class Welcome extends React.Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }
}
```

**主な違い:**

1. **構文**: 関数コンポーネントはシンプルに関数として定義
2. **状態管理**:
   - 関数コンポーネント: フック（useState）を使用
   - クラスコンポーネント: this.stateとthis.setState()を使用
3. **ライフサイクル**:
   - 関数コンポーネント: useEffectフックで統合
   - クラスコンポーネント: componentDidMount()などの個別メソッド
4. **thisの扱い**: 関数コンポーネントではthisを使わないため、バインドの問題が発生しない

### 1.6.2. TypeScriptによる型安全性の向上

TypeScriptを使用することで、以下のような利点があります：

1. **コンパイル時のエラー検出**:

   ```tsx
   // エラー: name は必須プロパティ
   <ProductCard id="1" price={100} />
   ```

2. **IDEのサポート強化**:
   - コード補完
   - 定義へのジャンプ
   - リファクタリングのサポート

3. **自己文書化コード**:
   - 型定義自体がドキュメントとして機能
   - インターフェースが明確になる

4. **型推論**:

   ```tsx
   // priceの型は自動的にnumberと推論される
   const [price, setPrice] = useState(100);
   ```

### 1.6.3. Next.jsとReactの関係

プロジェクトはNext.jsを使用していますが、ReactとNext.jsの関係を理解することも重要です。

1. **Next.jsとは**: Reactをベースにしたフレームワークで、以下の機能を提供
   - サーバーサイドレンダリング(SSR)
   - 静的サイト生成(SSG)
   - ファイルベースのルーティング
   - APIルート
   - 最適化機能

2. **Next.jsのページコンポーネント**:

   ```tsx
   // app/page.tsx
   export default function Home() {
     return <div>ホームページ</div>;
   }
   ```

3. **Appディレクトリ構造**:
   - `app/layout.tsx`: 共通レイアウト
   - `app/page.tsx`: トップページ
   - `app/products/page.tsx`: 商品一覧ページ

## 1.7. 【補足情報】

### 1.7.1. コンポーネント設計の原則

効果的なReactコンポーネントを設計するための原則：

1. **単一責任の原則**: 一つのコンポーネントは一つの役割だけを持つべき
   - 良い例: `ProductCard` - 商品の表示のみ担当
   - 悪い例: `ProductPageWithFilterAndCart` - 複数の役割を一つのコンポーネントに詰め込んでいる

2. **コンポジション（合成）の活用**:

   ```tsx
   // 小さなコンポーネントを組み合わせる
   <MainLayout>
     <ProductList>
       <ProductCard />
       <ProductCard />
     </ProductList>
   </MainLayout>
   ```

3. **presentationalとcontainerの分離**:
   - Presentational（表示）コンポーネント: UIの見た目のみに集中
   - Container（コンテナ）コンポーネント: データ取得やロジックを担当

4. **再利用性を考慮した設計**:
   - 具体的な値をハードコードしない
   - Propsを使って柔軟性を持たせる

### 1.7.2. パフォーマンス最適化テクニック

Reactアプリケーションのパフォーマンスを向上させるテクニック：

1. **メモ化**: 不要な再レンダリングを防止

   ```tsx
   import { memo } from 'react';

   const MemoizedComponent = memo(function MyComponent(props) {
     // レンダリングロジック
   });
   ```

2. **useMemo**: 計算コストの高い値を再計算を防止

   ```tsx
   const memoizedValue = useMemo(() => {
     return computeExpensiveValue(a, b);
   }, [a, b]);
   ```

3. **useCallback**: 関数の再生成を防止

   ```tsx
   const memoizedCallback = useCallback(() => {
     doSomething(a, b);
   }, [a, b]);
   ```

4. **コンポーネントの分割**: 状態変更の影響範囲を限定する

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: Propsが正しく渡されない

**症状**:

- コンポーネントにPropsが正しく渡されていないようだ
- `undefined is not an object` や `cannot read property of undefined` のようなエラーが出る

**解決策**:

1. Propsの型定義を確認する
2. デフォルト値を設定する

   ```tsx
   function MyComponent({ value = "デフォルト値" }) {
     return <div>{value}</div>;
   }
   ```

3. オプショナルチェイニングを使用する

   ```tsx
   <div>{user?.name}</div>
   ```

### 1.8.2. 問題2: 無限レンダリングループ

**症状**:

- コンポーネントが無限にレンダリングされる
- ブラウザが遅くなる または クラッシュする

**解決策**:

1. useEffectの依存配列を確認する

   ```tsx
   // 誤: 依存配列なし（レンダリング毎に実行）
   useEffect(() => {
     setCount(count + 1);
   });

   // 正: 空の依存配列（マウント時のみ実行）
   useEffect(() => {
     doSomething();
   }, []);
   ```

2. 状態更新の条件を設定する

   ```tsx
   useEffect(() => {
     // 条件付きで状態を更新
     if (data !== prevData) {
       setProcessedData(processData(data));
     }
   }, [data, prevData]);
   ```

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **JSXは宣言的UI記述**: HTMLに似た構文でUIを直感的に記述できる
2. **コンポーネント思考**: UIを独立した再利用可能なコンポーネントに分割する
3. **Propsによるデータフロー**: 親から子へ一方向のデータフローで予測可能性を高める
4. **フックによる状態管理**: useStateとuseEffectで関数コンポーネントに状態と副作用を追加
5. **TypeScriptの活用**: 型定義により開発時のエラー検出とコード補完を強化

これらのポイントは次回以降の実装でも活用されますので、よく理解しておきましょう。

## 1.10. 【次回の準備】

次回（Week 3 Day 1）では、バックエンドAPIとの連携方法、特にデータフェッチングについて学習します。以下の点について事前に確認しておくと良いでしょう：

1. HTTPリクエストの基本概念（GET, POST, PUT, DELETE）
2. フェッチAPIやAxiosの基本的な使い方
3. 非同期処理（Promise, async/await）の基本
4. エラーハンドリングの基本パターン
