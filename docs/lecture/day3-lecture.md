# 1. Week 1 - Day 3: フロントエンド環境のセットアップ

## 1.1. 目次

- [1. Week 1 - Day 3: フロントエンド環境のセットアップ](#1-week-1---day-3-フロントエンド環境のセットアップ)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. プロジェクトディレクトリの作成](#141-プロジェクトディレクトリの作成)
    - [1.4.2. 顧客向けフロントエンドのNext.jsプロジェクト作成](#142-顧客向けフロントエンドのnextjsプロジェクト作成)
    - [1.4.3. 管理者向けフロントエンドのNext.jsプロジェクト作成](#143-管理者向けフロントエンドのnextjsプロジェクト作成)
    - [1.4.4. 顧客向けフロントエンドのTypeScriptとAPIクライアントの設定](#144-顧客向けフロントエンドのtypescriptとapiクライアントの設定)
    - [1.4.5. 管理者向けフロントエンドのTypeScriptとAPIクライアントの設定](#145-管理者向けフロントエンドのtypescriptとapiクライアントの設定)
    - [1.4.6. 顧客向けフロントエンドの基本レイアウトコンポーネントの作成](#146-顧客向けフロントエンドの基本レイアウトコンポーネントの作成)
    - [1.4.7. 管理者向けフロントエンドの基本レイアウトコンポーネントの作成](#147-管理者向けフロントエンドの基本レイアウトコンポーネントの作成)
    - [1.4.8. Dockerfileの作成](#148-dockerfileの作成)
    - [1.4.9. アプリケーションの起動確認](#149-アプリケーションの起動確認)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. Next.js, TypeScript, TailwindCSSの様々な利点](#161-nextjs-typescript-tailwindcssの様々な利点)
    - [1.6.2. APIクライアントとインターセプターの仕組み](#162-apiクライアントとインターセプターの仕組み)
    - [1.6.3. Dockerマルチステージビルドの利点](#163-dockerマルチステージビルドの利点)
    - [1.6.4. ディレクトリ構成のベストプラクティス](#164-ディレクトリ構成のベストプラクティス)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. フロントエンドのパフォーマンス最適化](#171-フロントエンドのパフォーマンス最適化)
    - [1.7.2. SWRによるデータフェッチの最適化](#172-swrによるデータフェッチの最適化)
    - [1.7.3. レスポンシブデザインの実装](#173-レスポンシブデザインの実装)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. Next.jsプロジェクトビルド時の型エラー](#181-nextjsプロジェクトビルド時の型エラー)
    - [1.8.2. Dockerビルド時の環境変数問題](#182-dockerビルド時の環境変数問題)
    - [1.8.3. TailwindCSSのスタイルが反映されない問題](#183-tailwindcssのスタイルが反映されない問題)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)
  - [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- **Next.js/TypeScript/TailwindCSSによる2つのフロントエンド**: 顧客向けと管理者向けの2つの独立したフロントエンドアプリケーションを構築する基盤を整える
- **App Router採用**: Next.jsの最新機能であるApp Routerを使用した開発方法を学ぶ
- **TypeScript設定**: 型安全性を確保し、開発効率を向上させるためのTypeScript設定を行う
- **モダンなUI/UXフレームワーク導入**: TailwindCSSを使用したモダンなUIコンポーネントの構築方法を学ぶ
- **APIクライアント設定**: バックエンドと通信するための基本的なAPIクライアント構造の実装
- **Dockerコンテナ化**: フロントエンドアプリケーションをコンテナ化し、環境に依存しないデプロイメントを可能にする

## 1.3. 【準備】

以下の項目が準備されていることを確認してください：

### 1.3.1. チェックリスト

- [ ] Node.js v18.x以上がインストールされていること

  ```bash
  node -v
  # v18.x.x以上が表示されることを確認
  ```

- [ ] npmまたはyarnがインストールされていること

  ```bash
  npm -v
  # または
  yarn -v
  ```

- [ ] Dockerがインストールされていること

  ```bash
  docker -v
  ```

- [ ] gitがインストールされていること

  ```bash
  git --version
  ```

- [ ] プロジェクトルートディレクトリが正しく設定されていること

  ```bash
  # プロジェクトルートディレクトリに移動
  cd /path/to/aws-observability-ecommerce
  ```

- バックエンドのAPIエンドポイントが決定されていること（開発中は仮のエンドポイントでも可）

## 1.4. 【手順】

### 1.4.1. プロジェクトディレクトリの作成

プロジェクトのルートディレクトリに移動し、フロントエンド用のディレクトリを作成します。

```bash
# フロントエンド用のディレクトリを作成
mkdir -p {frontend-customer,frontend-admin}
```

### 1.4.2. 顧客向けフロントエンドのNext.jsプロジェクト作成

Next.js、TypeScript、TailwindCSSを使用した顧客向けフロントエンドプロジェクトを作成します。

```bash
cd frontend-customer

# Next.jsプロジェクトを作成 (App Router方式)
npx create-next-app@latest . --typescript --tailwind --eslint
# プロンプトが表示される場合は以下のように選択
# ✔ Would you like your code inside a `src/` directory? … No
# ✔ Would you like to use App Router? (recommended) … Yes
# ✔ Would you like to use Turbopack for `next dev`? … No
# ✔ Would you like to customize the import alias (`@/*` by default)? … No
```

プロジェクトの基本構造を整理します：

```bash
# 必要なディレクトリを作成
mkdir -p src/{components,lib,types}
mkdir -p src/components/{layout,ui}
mkdir -p src/lib/api
```

### 1.4.3. 管理者向けフロントエンドのNext.jsプロジェクト作成

同様に、管理者向けフロントエンドプロジェクトも作成します。

```bash
cd ../frontend-admin

# Next.jsプロジェクトを作成 (App Router方式)
npx create-next-app@latest . --typescript --tailwind --eslint
# プロンプトが表示される場合は同様に選択

# 必要なディレクトリを作成
mkdir -p src/{components,lib,types}
mkdir -p src/components/{layout,ui}
mkdir -p src/lib/{api,auth}
```

### 1.4.4. 顧客向けフロントエンドのTypeScriptとAPIクライアントの設定

顧客向けフロントエンドに移動し、TypeScriptの設定とAPIクライアントを実装します。

```bash
cd ../frontend-customer

# 必要なライブラリをインストール
npm install axios swr
```

tsconfig.jsonを編集して、型定義のパスやその他の設定を行います：

```json
{
  "compilerOptions": {
    "target": "es5",
    "lib": ["dom", "dom.iterable", "esnext"],
    "allowJs": true,
    "skipLibCheck": true,
    "strict": true,
    "noEmit": true,
    "esModuleInterop": true,
    "module": "esnext",
    "moduleResolution": "bundler",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "jsx": "preserve",
    "incremental": true,
    "plugins": [
      {
        "name": "next"
      }
    ],
    "paths": {
      "@/*": ["./src/*"]
    },
    "baseUrl": "."
  },
  "include": ["next-env.d.ts", "**/*.ts", "**/*.tsx", ".next/types/**/*.ts"],
  "exclude": ["node_modules"]
}
```

APIクライアントを実装します：

```bash
touch src/lib/api/client.ts
```

src/lib/api/client.tsの内容：

```typescript
import axios from "axios";

// APIのベースURL
// 本番環境では環境変数から取得する
const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001/api";

// Axiosインスタンスの作成
const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// リクエストインターセプター
apiClient.interceptors.request.use(
  (config) => {
    // トークンがある場合はヘッダーに追加
    const token = localStorage.getItem("auth_token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// レスポンスインターセプター
apiClient.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    // エラーハンドリング
    if (error.response) {
      // 認証エラーの場合
      if (error.response.status === 401) {
        // ログアウト処理など
        localStorage.removeItem("auth_token");
      }
    }
    return Promise.reject(error);
  }
);

export default apiClient;
```

### 1.4.5. 管理者向けフロントエンドのTypeScriptとAPIクライアントの設定

管理者向けフロントエンドに移動し、TypeScriptとAPIクライアントの設定を行います。

```bash
cd ../frontend-admin

# 必要なライブラリをインストール
npm install axios swr
```

tsconfig.jsonの内容は顧客向けフロントエンドと同様ですが、管理者向けには認証機能をより強化します：

```bash
touch src/lib/api/client.ts
touch src/lib/auth/auth.ts
```

src/lib/api/client.tsの内容：

```typescript
import axios from "axios";
import { clearAuthData, getAuthToken } from "../auth/auth";

// APIのベースURL
const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001/api";

// Axiosインスタンスの作成
const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// リクエストインターセプター
apiClient.interceptors.request.use(
  (config) => {
    // 管理者認証トークンをヘッダーに追加
    const token = getAuthToken();
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// レスポンスインターセプター
apiClient.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    // エラーハンドリング
    if (error.response) {
      // 認証エラーの場合はログアウト処理
      if (error.response.status === 401) {
        clearAuthData();
        // ログインページへリダイレクト（ブラウザ環境の場合）
        if (typeof window !== "undefined") {
          window.location.href = "/login";
        }
      }
    }
    return Promise.reject(error);
  }
);

export default apiClient;
```

src/lib/auth/auth.tsの内容：

```typescript
// 管理者認証のためのユーティリティ関数

// 認証トークンの保存キー
const AUTH_TOKEN_KEY = "admin_auth_token";
const USER_DATA_KEY = "admin_user_data";

// 認証トークンを取得
export const getAuthToken = (): string | null => {
  if (typeof window === "undefined") {
    return null;
  }
  return localStorage.getItem(AUTH_TOKEN_KEY);
};

// ユーザーデータを取得
export const getUserData = (): any => {
  if (typeof window === "undefined") {
    return null;
  }
  const userData = localStorage.getItem(USER_DATA_KEY);
  if (userData) {
    try {
      return JSON.parse(userData);
    } catch (e) {
      return null;
    }
  }
  return null;
};

// 認証情報を保存
export const setAuthData = (token: string, userData: any): void => {
  if (typeof window === "undefined") {
    return;
  }
  localStorage.setItem(AUTH_TOKEN_KEY, token);
  localStorage.setItem(USER_DATA_KEY, JSON.stringify(userData));
};

// 認証情報をクリア
export const clearAuthData = (): void => {
  if (typeof window === "undefined") {
    return;
  }
  localStorage.removeItem(AUTH_TOKEN_KEY);
  localStorage.removeItem(USER_DATA_KEY);
};

// ログイン状態の確認
export const isAuthenticated = (): boolean => {
  return !!getAuthToken();
};

// モック認証（開発用）
export const mockLogin = async (
  email: string,
  password: string
): Promise<any> => {
  // 開発環境用の簡易認証（本番環境では使用しないこと）
  if (email === "admin@example.com" && password === "password") {
    const mockToken = "mock-jwt-token-for-development-only";
    const mockUserData = {
      id: "1",
      name: "Admin User",
      email: "admin@example.com",
      role: "admin",
    };
    setAuthData(mockToken, mockUserData);
    return {
      token: mockToken,
      user: mockUserData,
    };
  }
  throw new Error(
    "認証に失敗しました。メールアドレスまたはパスワードが間違っています。"
  );
};
```

また、管理者向けフロントエンドには型定義ファイルを追加します：

```bash
touch src/types/auth.ts
```

src/types/auth.tsの内容：

```typescript
// 認証関連の型定義

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

### 1.4.6. 顧客向けフロントエンドの基本レイアウトコンポーネントの作成

顧客向けフロントエンドの基本レイアウトコンポーネントを作成します。

```bash
touch src/components/layout/{Header.tsx,Footer.tsx,MainLayout.tsx}
```

src/components/layout/Header.tsxの内容（`"use client"`ディレクティブに注意してください）：

```tsx
"use client";

import Link from "next/link";
import { useEffect, useState } from "react";

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

  return (
    <header
      className={`fixed w-full z-50 text-white transition-all duration-300 ${
        isScrolled
          ? "bg-gradient-primary shadow-lg py-2"
          : "bg-transparent py-4"
      }`}
    >
      <div className="container mx-auto px-4 flex justify-between items-center">
        <Link href="/" className="text-2xl font-bold flex items-center">
          <span className="mr-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-8 h-8"
            >
              <path d="M11.25 3v4.046a3 3 0 00-4.277 4.204H1.5v-6A2.25 2.25 0 013.75 3h7.5zM12.75 3v4.011a3 3 0 014.239 4.239H22.5v-6A2.25 2.25 0 0020.25 3h-7.5zM22.5 12.75h-8.983a4.125 4.125 0 004.108 3.75.75.75 0 010 1.5 5.623 5.623 0 01-4.875-2.817V21h7.5a2.25 2.25 0 002.25-2.25v-6zM11.25 21v-5.817A5.623 5.623 0 016.375 18a.75.75 0 010-1.5 4.126 4.126 0 004.108-3.75H1.5v6A2.25 2.25 0 003.75 21h7.5z" />
              <path d="M11.085 10.354c.03.297.038.575.036.805a7.484 7.484 0 01-.805-.036c-.833-.084-1.677-.325-2.195-.843a1.5 1.5 0 012.122-2.12c.517.517.759 1.36.842 2.194zM12.877 10.354c-.03.297-.038.575-.036.805.23.002.508-.006.805-.036.833-.084 1.677-.325 2.195-.843A1.5 1.5 0 0013.72 8.16c-.518.518-.76 1.362-.843 2.194z" />
            </svg>
          </span>
          <span className="bg-clip-text text-transparent bg-gradient-to-r from-white to-indigo-200">
            Observability Shop
          </span>
        </Link>

        {/* デスクトップメニュー */}
        <nav className="hidden md:block">
          <ul className="flex space-x-8">
            <li>
              <Link
                href="/products"
                className="hover:text-indigo-200 transition-colors flex items-center font-medium"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-1"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"
                  />
                </svg>
                商品一覧
              </Link>
            </li>
            <li>
              <Link
                href="/cart"
                className="hover:text-indigo-200 transition-colors flex items-center font-medium"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-1"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
                  />
                </svg>
                カート
              </Link>
            </li>
            <li>
              <Link
                href="/orders"
                className="hover:text-indigo-200 transition-colors flex items-center font-medium"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-1"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
                  />
                </svg>
                注文履歴
              </Link>
            </li>
            <li>
              <Link
                href="/account"
                className="hover:text-indigo-200 transition-colors flex items-center font-medium"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-1"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                  />
                </svg>
                アカウント
              </Link>
            </li>
          </ul>
        </nav>

        {/* 検索ボタン */}
        <div className="hidden md:flex items-center">
          <div className="relative group">
            <button className="flex items-center justify-center w-9 h-9 bg-white/20 backdrop-blur-sm rounded-full hover:bg-white/30 transition-colors">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                />
              </svg>
            </button>
            <div className="absolute -bottom-10 left-1/2 transform -translate-x-1/2 px-3 py-1.5 bg-gray-900 text-white text-xs font-medium rounded-md opacity-0 group-hover:opacity-100 transition-opacity duration-300 whitespace-nowrap pointer-events-none">
              商品を検索
            </div>
          </div>
          <div className="relative group ml-4">
            <Link
              href="/cart"
              className="relative flex items-center justify-center w-9 h-9 bg-white/20 backdrop-blur-sm rounded-full hover:bg-white/30 transition-colors"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
                />
              </svg>
              <span className="absolute -top-1 -right-1 bg-red-500 text-white text-xs w-5 h-5 flex items-center justify-center rounded-full">
                2
              </span>
            </Link>
            <div className="absolute -bottom-10 left-1/2 transform -translate-x-1/2 px-3 py-1.5 bg-gray-900 text-white text-xs font-medium rounded-md opacity-0 group-hover:opacity-100 transition-opacity duration-300 whitespace-nowrap pointer-events-none">
              カートを見る
            </div>
          </div>
        </div>

        {/* モバイルメニューボタン */}
        <button
          className="md:hidden bg-white/20 backdrop-blur-sm p-2 rounded-full hover:bg-white/30 transition-colors"
          onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
        >
          {isMobileMenuOpen ? (
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          ) : (
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M4 6h16M4 12h16m-7 6h7"
              />
            </svg>
          )}
        </button>
      </div>

      {/* モバイルメニュー */}
      {isMobileMenuOpen && (
        <div className="md:hidden bg-gradient-primary shadow-lg mt-2 py-4 px-4 rounded-b-xl">
          <ul className="space-y-4">
            <li>
              <Link
                href="/products"
                className="flex items-center py-2 px-3 rounded-lg hover:bg-white/10 transition-colors"
                onClick={() => setIsMobileMenuOpen(false)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"
                  />
                </svg>
                商品一覧
              </Link>
            </li>
            <li>
              <Link
                href="/cart"
                className="flex items-center py-2 px-3 rounded-lg hover:bg-white/10 transition-colors"
                onClick={() => setIsMobileMenuOpen(false)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
                  />
                </svg>
                カート
                <span className="ml-auto bg-red-500 text-white text-xs px-2 py-1 rounded-full">
                  2
                </span>
              </Link>
            </li>
            <li>
              <Link
                href="/orders"
                className="flex items-center py-2 px-3 rounded-lg hover:bg-white/10 transition-colors"
                onClick={() => setIsMobileMenuOpen(false)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
                  />
                </svg>
                注文履歴
              </Link>
            </li>
            <li>
              <Link
                href="/account"
                className="flex items-center py-2 px-3 rounded-lg hover:bg-white/10 transition-colors"
                onClick={() => setIsMobileMenuOpen(false)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                  />
                </svg>
                アカウント
              </Link>
            </li>
            <li>
              <Link
                href="/search"
                className="flex items-center py-2 px-3 rounded-lg hover:bg-white/10 transition-colors"
                onClick={() => setIsMobileMenuOpen(false)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 mr-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                  />
                </svg>
                検索
              </Link>
            </li>
          </ul>
        </div>
      )}
    </header>
  );
}
```

src/components/layout/Footer.tsxの内容：

```tsx
export default function Footer() {
  return (
    <footer className="bg-gradient-to-b from-gray-800 to-gray-900 dark:from-gray-900 dark:to-black text-white py-16">
      <div className="container mx-auto px-6">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-12">
          <div>
            <h3 className="text-xl font-semibold mb-4">会社情報</h3>
            <p className="mb-2">オブザーバビリティ株式会社</p>
            <p className="mb-2">東京都渋谷区渋谷 1-1-1</p>
            <p>電話: 03-1234-5678</p>
          </div>
          <div>
            <h3 className="text-xl font-semibold mb-4">リンク</h3>
            <ul className="space-y-2">
              <li>
                <a
                  href="/about"
                  className="hover:text-indigo-300 transition-colors"
                >
                  当サイトについて
                </a>
              </li>
              <li>
                <a
                  href="/privacy"
                  className="hover:text-indigo-300 transition-colors"
                >
                  プライバシーポリシー
                </a>
              </li>
              <li>
                <a
                  href="/terms"
                  className="hover:text-indigo-300 transition-colors"
                >
                  利用規約
                </a>
              </li>
              <li>
                <a
                  href="/contact"
                  className="hover:text-indigo-300 transition-colors"
                >
                  お問い合わせ
                </a>
              </li>
            </ul>
          </div>
          <div>
            <h3 className="text-xl font-semibold mb-4">ニュースレター</h3>
            <p className="mb-4">最新のお知らせやセール情報をお届けします。</p>
            <div className="flex">
              <input
                type="email"
                placeholder="メールアドレス"
                className="px-4 py-3 w-full rounded-l-lg text-gray-800 border-0 shadow-sm"
              />
              <button className="btn-primary px-6 py-3 rounded-r-lg shadow-sm hover:shadow-md font-medium">
                登録
              </button>
            </div>
          </div>
        </div>
        <div className="mt-12 pt-8 border-t border-gray-700 text-center">
          <div className="flex justify-center space-x-6 mb-6">
            {/* SNSアイコン */}
            <a
              href="#"
              className="text-gray-400 hover:text-white transition-colors"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="currentColor"
                viewBox="0 0 24 24"
              >
                <path d="M24 4.557c-.883.392-1.832.656-2.828.775 1.017-.609 1.798-1.574 2.165-2.724-.951.564-2.005.974-3.127 1.195-.897-.957-2.178-1.555-3.594-1.555-3.179 0-5.515 2.966-4.797 6.045-4.091-.205-7.719-2.165-10.148-5.144-1.29 2.213-.669 5.108 1.523 6.574-.806-.026-1.566-.247-2.229-.616-.054 2.281 1.581 4.415 3.949 4.89-.693.188-1.452.232-2.224.084.626 1.956 2.444 3.379 4.6 3.419-2.07 1.623-4.678 2.348-7.29 2.04 2.179 1.397 4.768 2.212 7.548 2.212 9.142 0 14.307-7.721 13.995-14.646.962-.695 1.797-1.562 2.457-2.549z" />
              </svg>
            </a>
            <a
              href="#"
              className="text-gray-400 hover:text-white transition-colors"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="currentColor"
                viewBox="0 0 24 24"
              >
                <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z" />
              </svg>
            </a>
            <a
              href="#"
              className="text-gray-400 hover:text-white transition-colors"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="currentColor"
                viewBox="0 0 24 24"
              >
                <path d="M22.675 0h-21.35c-.732 0-1.325.593-1.325 1.325v21.351c0 .731.593 1.324 1.325 1.324h11.495v-9.294h-3.128v-3.622h3.128v-2.671c0-3.1 1.893-4.788 4.659-4.788 1.325 0 2.463.099 2.795.143v3.24l-1.918.001c-1.504 0-1.795.715-1.795 1.763v2.313h3.587l-.467 3.622h-3.12v9.293h6.116c.73 0 1.323-.593 1.323-1.325v-21.35c0-.732-.593-1.325-1.325-1.325z" />
              </svg>
            </a>
          </div>
          <p className="text-gray-400">
            &copy; {new Date().getFullYear()} Observability Shop. All rights
            reserved.
          </p>
        </div>
      </div>
    </footer>
  );
}
```

src/components/layout/MainLayout.tsxの内容：

```tsx
import { ReactNode } from "react";
import Footer from "./Footer";
import Header from "./Header";

type MainLayoutProps = {
  children: ReactNode;
};

export default function MainLayout({ children }: MainLayoutProps) {
  return (
    <div className="flex flex-col min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
      <Header />
      <main className="flex-grow pt-24 md:pt-28 pb-12">{children}</main>
      <Footer />
    </div>
  );
}
```

### 1.4.7. 管理者向けフロントエンドの基本レイアウトコンポーネントの作成

管理者向けフロントエンドの基本レイアウトコンポーネントを作成します。管理向けUIはダッシュボード風の設計にします。

```bash
cd ../frontend-admin
touch src/components/layout/{AdminHeader.tsx,AdminSidebar.tsx,AdminLayout.tsx}
```

src/components/layout/AdminHeader.tsxの内容：

```tsx
"use client";

import { getUserData } from "@/lib/auth/auth";
import Link from "next/link";
import { useState } from "react";

type AdminHeaderProps = {
  toggleSidebar?: () => void;
};

export default function AdminHeader({ toggleSidebar }: AdminHeaderProps) {
  const [isProfileMenuOpen, setIsProfileMenuOpen] = useState(false);
  const userData = getUserData() || { name: "Admin User" };

  const toggleProfileMenu = () => {
    setIsProfileMenuOpen(!isProfileMenuOpen);
  };

  return (
    <header className="bg-gradient-primary text-white shadow-lg">
      <div className="px-4 py-4 flex justify-between items-center">
        <div className="flex items-center">
          <button
            onClick={toggleSidebar}
            className="p-2 rounded-md hover:bg-white/10 transition-colors"
            aria-label="サイドバー切り替え"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M4 6h16M4 12h16M4 18h16"
              />
            </svg>
          </button>
          <Link href="/admin" className="text-xl font-bold flex items-center">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-6 h-6 mr-2"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M2.25 13.5h3.86a2.25 2.25 0 012.012 1.244l.256.512a2.25 2.25 0 002.013 1.244h3.218a2.25 2.25 0 002.013-1.244l.256-.512a2.25 2.25 0 012.013-1.244h3.859m-19.5.338V18a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18v-4.162c0-.224-.034-.447-.1-.661L19.24 5.338a2.25 2.25 0 00-2.15-1.588H6.911a2.25 2.25 0 00-2.15 1.588L2.35 13.177a2.25 2.25 0 00-.1.661z"
              />
            </svg>
            管理ダッシュボード
          </Link>
        </div>

        <div className="relative">
          <button
            onClick={toggleProfileMenu}
            className="flex items-center space-x-2 focus:outline-none hover:opacity-80 transition-smooth px-2 py-1 rounded-full"
          >
            <div className="w-9 h-9 rounded-full bg-white/20 backdrop-blur-sm shadow-inner flex items-center justify-center font-semibold transition-all">
              {userData.name.charAt(0)}
            </div>
            <span className="hidden md:inline font-medium">
              {userData.name}
            </span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={2}
              stroke="currentColor"
              className="w-4 h-4 transition-transform duration-200 ease-in-out"
              style={{
                transform: isProfileMenuOpen
                  ? "rotate(180deg)"
                  : "rotate(0deg)",
              }}
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M19.5 8.25l-7.5 7.5-7.5-7.5"
              />
            </svg>
          </button>

          {isProfileMenuOpen && (
            <div className="absolute right-0 mt-2 w-52 bg-white dark:bg-gray-800 rounded-xl shadow-xl py-2 text-gray-800 dark:text-gray-200 z-10 border border-gray-100 dark:border-gray-700 overflow-hidden">
              <Link
                href="/admin/profile"
                className="flex items-center px-4 py-2.5 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors group"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-5 h-5 mr-3 text-gray-500 group-hover:text-primary"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                  />
                </svg>
                <span className="font-medium">プロフィール</span>
              </Link>
              <Link
                href="/admin/settings"
                className="flex items-center px-4 py-2.5 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors group"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-5 h-5 mr-3 text-gray-500 group-hover:text-primary"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
                  />
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                </svg>
                <span className="font-medium">設定</span>
              </Link>
              <div className="border-t border-gray-200 dark:border-gray-700 my-1"></div>
              <button
                onClick={() => {
                  // ログアウト処理
                  if (typeof window !== "undefined") {
                    window.location.href = "/admin/login";
                  }
                }}
                className="flex w-full items-center px-4 py-2.5 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors group text-left"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-5 h-5 mr-3 text-red-500"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                  />
                </svg>
                <span className="font-medium text-red-600 dark:text-red-500">
                  ログアウト
                </span>
              </button>
            </div>
          )}
        </div>
      </div>
    </header>
  );
}
```

src/components/layout/AdminSidebar.tsxの内容：

```tsx
"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";

type AdminSidebarProps = {
  isOpen?: boolean;
};

export default function AdminSidebar({ isOpen = true }: AdminSidebarProps) {
  const pathname = usePathname();

  const navigationItems = [
    {
      name: "ダッシュボード",
      href: "/admin",
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-5 h-5"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M3.75 6A2.25 2.25 0 016 3.75h2.25A2.25 2.25 0 0110.5 6v2.25a2.25 2.25 0 01-2.25 2.25H6a2.25 2.25 0 01-2.25-2.25V6zM3.75 15.75A2.25 2.25 0 016 13.5h2.25a2.25 2.25 0 012.25 2.25V18a2.25 2.25 0 01-2.25 2.25H6A2.25 2.25 0 013.75 18v-2.25zM13.5 6a2.25 2.25 0 012.25-2.25H18A2.25 2.25 0 0120.25 6v2.25A2.25 2.25 0 0118 10.5h-2.25a2.25 2.25 0 01-2.25-2.25V6zM13.5 15.75a2.25 2.25 0 012.25-2.25H18a2.25 2.25 0 012.25 2.25V18A2.25 2.25 0 0118 20.25h-2.25A2.25 2.25 0 0113.5 18v-2.25z"
          />
        </svg>
      ),
    },
    {
      name: "商品管理",
      href: "/admin/products",
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-5 h-5"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z"
          />
        </svg>
      ),
    },
    {
      name: "カテゴリー管理",
      href: "/admin/categories",
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-5 h-5"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M8.25 6.75h12M8.25 12h12m-12 5.25h12M3.75 6.75h.007v.008H3.75V6.75zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zM3.75 12h.007v.008H3.75V12zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm-.375 5.25h.007v.008H3.75v-.008zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z"
          />
        </svg>
      ),
    },
    {
      name: "注文管理",
      href: "/admin/orders",
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-5 h-5"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M9 12h3.75M9 15h3.75M9 18h3.75m3 .75H18a2.25 2.25 0 002.25-2.25V6.108c0-1.135-.845-2.098-1.976-2.192a48.424 48.424 0 00-1.123-.08m-5.801 0c-.065.21-.1.433-.1.664 0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75 2.25 2.25 0 00-.1-.664m-5.8 0A2.251 2.251 0 0113.5 2.25H15c1.012 0 1.867.668 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m0 0H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V9.375c0-.621-.504-1.125-1.125-1.125H8.25zM6.75 12h.008v.008H6.75V12zm0 3h.008v.008H6.75V15zm0 3h.008v.008H6.75V18z"
          />
        </svg>
      ),
    },
    {
      name: "在庫管理",
      href: "/admin/inventory",
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-5 h-5"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M20.25 6.375c0 2.278-3.694 4.125-8.25 4.125S3.75 8.653 3.75 6.375m16.5 0c0-2.278-3.694-4.125-8.25-4.125S3.75 4.097 3.75 6.375m16.5 0v11.25c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125V6.375m16.5 0v3.75m-16.5-3.75v3.75m16.5 0v3.75C20.25 16.153 16.556 18 12 18s-8.25-1.847-8.25-4.125v-3.75m16.5 0c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125"
          />
        </svg>
      ),
    },
  ];

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

  return (
    <aside
      className={`bg-sidebar text-white min-h-screen transition-all duration-300 overflow-hidden ${
        isOpen ? "w-64 p-4" : "w-0"
      }`}
    >
      {isOpen && (
        <div className="sticky top-0">
          <nav>
            <ul className="space-y-2">
              {navigationItems.map((item) => {
                const isActive =
                  pathname === item.href ||
                  pathname?.startsWith(`${item.href}/`);
                return (
                  <li key={item.name}>
                    <Link
                      href={item.href}
                      className={`flex items-center space-x-3 px-4 py-3 rounded-lg ${
                        isActive
                          ? "bg-indigo-700 text-white shadow-lg shadow-indigo-700/30"
                          : "text-gray-300 hover:bg-indigo-700/20 hover:text-white"
                      } transition-all duration-200 ease-in-out hover:translate-x-1`}
                    >
                      <span
                        className={`${
                          isActive ? "text-white" : "text-indigo-400"
                        }`}
                      >
                        {item.icon}
                      </span>
                      <span className="font-medium">{item.name}</span>
                      {isActive && (
                        <span className="ml-auto bg-white w-1.5 h-1.5 rounded-full"></span>
                      )}
                    </Link>
                  </li>
                );
              })}
            </ul>

            <div className="border-t border-gray-700 my-6"></div>

            <ul className="space-y-2">
              <li>
                <Link
                  href="/admin/settings"
                  className="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-300 hover:bg-indigo-700/20 hover:text-white transition-all duration-200 ease-in-out hover:translate-x-1"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className="w-5 h-5 text-indigo-400"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 011.37.49l1.296 2.247a1.125 1.125 0 01-.26 1.431l-1.003.827c-.293.24-.438.613-.431.992a6.759 6.759 0 010 .255c-.007.378.138.75.43.99l1.005.828c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 01-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 01-.22.128c-.331.183-.581.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 01-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 01-1.369-.49l-1.297-2.247a1.125 1.125 0 01.26-1.431l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 010-.255c.007-.378-.138-.75-.43-.99l-1.004-.828a1.125 1.125 0 01-.26-1.43l1.297-2.247a1.125 1.125 0 011.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.128.332-.183.582-.495.644-.869l.214-1.281z"
                    />
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                    />
                  </svg>
                  <span className="font-medium">設定</span>
                </Link>
              </li>
              <li>
                <Link
                  href="/admin/help"
                  className="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-300 hover:bg-indigo-700/20 hover:text-white transition-all duration-200 ease-in-out hover:translate-x-1"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className="w-5 h-5 text-indigo-400"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9 5.25h.008v.008H12v-.008z"
                    />
                  </svg>
                  <span className="font-medium">ヘルプ</span>
                </Link>
              </li>
            </ul>
          </nav>

          {/* サイドバー底部のエレメント */}
          <div className="mt-auto pt-6 border-t border-gray-700/50 mt-6">
            <div className="bg-indigo-700/20 rounded-lg p-4">
              <div className="flex items-center mb-3">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 text-indigo-400 mr-2"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
                <p className="text-sm font-medium text-indigo-200">
                  管理者ステータス
                </p>
              </div>
              <p className="text-xs text-gray-300">
                システムは正常に稼働しています。最終チェック：{currentTime}
              </p>
            </div>
          </div>
        </div>
      )}
    </aside>
  );
}
```

src/components/layout/AdminLayout.tsxの内容：

```tsx
"use client";

import { ReactNode, useState } from "react";
import AdminHeader from "./AdminHeader";
import AdminSidebar from "./AdminSidebar";

type AdminLayoutProps = {
  children: ReactNode;
};

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

また、管理者向けダッシュボード用のコンポーネントも作成しておきます：

```bash
touch src/components/ui/{DashboardCard.tsx,DashboardStats.tsx}
```

src/components/ui/DashboardCard.tsxの内容：

```tsx
import { ReactNode } from "react";

type DashboardCardProps = {
  title: string;
  icon?: ReactNode;
  children: ReactNode;
  className?: string;
};

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

src/components/ui/DashboardStats.tsxの内容：

```tsx
type StatItemProps = {
  label: string;
  value: string | number;
  change?: string;
  isPositive?: boolean;
};

function StatItem({ label, value, change, isPositive }: StatItemProps) {
  const gradientColorClass = isPositive
    ? "from-emerald-500 to-teal-600"
    : "from-red-500 to-orange-600";

  return (
    <div className="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-5 stat-card relative overflow-hidden">
      {/* 見えないグラデーションアクセント（右上角） */}
      <div
        className={`absolute top-0 right-0 w-20 h-20 bg-gradient-to-br opacity-10 ${gradientColorClass}`}
      ></div>
      {/* アイコンインジケーター */}
      <div className="flex justify-between items-start">
        <p className="text-sm font-medium text-gray-500 dark:text-gray-400">
          {label}
        </p>
        {change && (
          <div
            className={`rounded-full w-8 h-8 flex items-center justify-center ${
              isPositive
                ? "bg-green-100 dark:bg-green-900/30 text-green-600 dark:text-green-400"
                : "bg-red-100 dark:bg-red-900/30 text-red-600 dark:text-red-400"
            }`}
          >
            {isPositive ? (
              <svg
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M5 10l7-7m0 0l7 7m-7-7v18"
                />
              </svg>
            ) : (
              <svg
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M19 14l-7 7m0 0l-7-7m7 7V3"
                />
              </svg>
            )}
          </div>
        )}
      </div>

      <p className="mt-2 text-3xl font-bold text-gray-900 dark:text-gray-100">
        {value}
      </p>

      {change && (
        <div className="mt-4 flex items-center">
          <div
            className={`h-1 w-12 rounded-full ${
              isPositive
                ? "bg-gradient-to-r from-green-400 to-emerald-500"
                : "bg-gradient-to-r from-red-500 to-orange-400"
            }`}
          ></div>
          <p
            className={`ml-2 text-sm font-medium ${
              isPositive
                ? "text-green-600 dark:text-green-400"
                : "text-red-600 dark:text-red-400"
            }`}
          >
            {change}
          </p>
        </div>
      )}
    </div>
  );
}

type DashboardStatsProps = {
  stats: StatItemProps[];
};

export default function DashboardStats({ stats }: DashboardStatsProps) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      {stats.map((stat, index) => (
        <StatItem
          key={index}
          label={stat.label}
          value={stat.value}
          change={stat.change}
          isPositive={stat.isPositive}
        />
      ))}
    </div>
  );
}
```

### 1.4.8. Dockerfileの作成

顧客向けと管理者向けの両方のフロントエンドプロジェクトにDockerfileを作成します。

まず、顧客向けフロントエンド用のDockerfileを作成します。

```bash
cd ../frontend-customer
touch Dockerfile .dockerignore
```

Dockerfileの内容：

```dockerfile
FROM node:23-alpine

WORKDIR /app

# 開発に必要なツールのインストール
RUN apk add --no-cache git curl bash

# タイムゾーンの設定
RUN apk add --no-cache tzdata
ENV TZ=Asia/Tokyo

# 依存関係をインストール
COPY package.json package-lock.json ./
RUN npm ci

# Lintとフォーマット設定ファイルをコピー
# COPY .eslintrc.js .eslintignore .prettierrc.js .prettierignore ./

# アプリケーションのソースコードをコピー
COPY . .

# 開発サーバーを起動
CMD ["npm", "run", "dev"]
```

.dockerignoreの内容：

```text
.git
.github
node_modules
.next
npm-debug.log
README.md
.env*
.dockerignore
Dockerfile
```

次に、管理者向けフロントエンドのDockerfileも作成します。

```bash
cd ../frontend-admin
touch Dockerfile .dockerignore
```

管理者向けのDockerfile、.dockerignoreの内容も顧客向けと同じです。

Docker Composeを使って、両方のフロントエンドアプリケーションを起動します。

`compose.yml`に以下の内容を追加・更新します：

```yaml
・・・

  backend:

・・・

  frontend-customer:
    build:
      context: ./frontend-customer
      dockerfile: Dockerfile
    container_name: frontend-customer
    restart: unless-stopped
    environment:
      - NEXT_PUBLIC_API_URL=http://localhost:8000/api
    volumes:
      - ./frontend-customer:/app # ホットリロード用
      - /app/node_modules # node_modulesはコンテナ内のまま
      - /app/.next # .nextディレクトリはコンテナ内のまま
    ports:
      - "3000:3000"
    depends_on:
      - backend
    networks:
      - ecommerce-network
    healthcheck:
      test:
        [
          "CMD",
          "wget",
          "--no-verbose",
          "--tries=1",
          "--spider",
          "http://localhost:3000",
        ]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 15s
    deploy:
      resources:
        limits:
          memory: 512M

  frontend-admin:
    build:
      context: ./frontend-admin
      dockerfile: Dockerfile
    container_name: frontend-admin
    restart: unless-stopped
    environment:
      - NEXT_PUBLIC_API_URL=http://backend:8080/api
    volumes:
      - ./frontend-admin:/app
      - /app/node_modules
    ports:
      - "3010:3000"
    depends_on:
      - backend
    networks:
      - ecommerce-network
    healthcheck:
      test:
        [
          "CMD",
          "wget",
          "--no-verbose",
          "--tries=1",
          "--spider",
          "http://localhost:3000",
        ]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 15s
    deploy:
      resources:
        limits:
          memory: 512M

・・・
```

### 1.4.9. アプリケーションの起動確認

Docker Composeで両方のフロントエンドを起動します。プロジェクトのルートディレクトリで以下のコマンドを実行します：

```bash
# Docker Composeでサービスをビルドし起動
docker-compose build frontend-customer frontend-admin
docker-compose up -d
```

サービスの起動状況を確認します：

```bash
# サービスの状態確認
docker-compose ps
```

すべてのサービスが起動していることを確認したら、バックエンドのログを確認します：

```bash
# バックエンドのログを表示
docker-compose logs frontend-customer frontend-admin
```

これで以下のURLでアプリケーションにアクセスできるようになります：

- 顧客向けフロントエンド: <http://localhost:3000>
- 管理者向けフロントエンド: <http://localhost:3010>

## 1.5. 【確認ポイント】

実装が正しく完了したことを確認するためのチェックリスト：

- [ ] 顧客向けフロントエンドが <http://localhost:3000> でアクセスできる
- [ ] 管理者向けフロントエンドが <http://localhost:3010> でアクセスできる
- [ ] Docker Composeで両方のフロントエンドが正常に起動している
- [ ] ヘッダーとフッターが正しく表示されている
- [ ] TailwindCSSのスタイルが適用されている
- [ ] TypeScriptの型チェックがエラーなく完了する

TypeScript型チェックの確認コマンド：

```bash
# 顧客向けフロントエンドでの型チェック
cd frontend-customer
npx tsc --noEmit

# 管理者向けフロントエンドでの型チェック
cd ../frontend-admin
npx tsc --noEmit
```

## 1.6. 【詳細解説】

### 1.6.1. Next.js, TypeScript, TailwindCSSの様々な利点

Next.js, TypeScript, TailwindCSSの組み合わせは、現代的なフロントエンド開発で高い生産性を実現するための強力な構成です。

- **Next.js**: Reactフレームワークを拡張したフルスタックフレームワークで、サーバーサイドレンダリング(SSR)や静的サイト生成(SSG)、App Router、ルーティングなどの機能を提供します。これによりパフォーマンスが向上し、SEOにも有利です。

  - **App Router**: Next.js 13以降で導入された新機能で、ファイルシステムベースのルーティングを提供します。従来のPages Routerと比較して、React Server Componentsの活用や、レイアウト、ローディング状態、エラーハンドリングなどの組み込み機能を提供します。

  - **Server Components**: クライアントJavaScriptバンドルを減らし、パフォーマンスを向上させます。データフェッチやサーバーリソースへのアクセスが直接可能です。

- **TypeScript**: JavaScriptに静的型付けを導入した言語で、開発時のエラー捕捉やコードの自己文書化性を向上させます。特に大規模なアプリケーションやチーム開発で威力を発揮します。

- **TailwindCSS**: ユーティリティファーストなCSSフレームワークで、小さなクラスを組み合わせてスタイルを構築します。カスタマイズ性が高く、軸となるデザインシステムを構築しやすい特徴があります。今回の実装では、カスタムのユーティリティクラスとグラデーションスタイルも追加しています。

### 1.6.2. APIクライアントとインターセプターの仕組み

今回実装したAPIクライアントは、Axiosを使用してバックエンドとの通信を抽象化しています。重要な点は以下の通りです：

1. **インターセプターパターン**: すべてのリクエストとレスポンスに共通の処理を推奨するパターンで、模擬データの差し替えや認証トークンの付与などを一元管理できます。

2. **環境変数引数化**: APIエンドポイントの値を環境変数から取得することで、ローカル開発環境と本番環境での切り替えが容易になります。

3. **エラーハンドリング**: レスポンスインターセプターで共通のエラー処理を実装することで、アプリ全体で一貫したエラーハンドリングが可能になります。

### 1.6.3. Dockerマルチステージビルドの利点

本番環境向けのDockerfileでは、マルチステージビルドを使用することで以下の利点があります：

1. **イメージサイズの最適化**: ビルドプロセスで必要なファイルのみを最終イメージに含めることで、イメージサイズを小さく保つことができます。

2. **キャッシュの有効活用**: 依存関係のインストールとビルドを別のステージで行うことで、キャッシュを効率的に利用し、ビルド時間を短縮します。

3. **セキュリティの向上**: 本番ステージでは非ルートユーザーを使用し、最小限の権限でアプリケーションを実行することで、セキュリティリスクを低減します。

4. **ビルド環境と実行環境の切り分け**: ビルドツールや中間ファイルは最終イメージに含まれないため、脆弱性の端末数が減ります。

今回の開発環境では、ホットリロードを活用して開発効率を高めるため、シンプルな単一ステージのDockerfileを採用しています。

### 1.6.4. ディレクトリ構成のベストプラクティス

本講義で提案しているディレクトリ構成は、App Routerを使用したNext.jsとTypeScriptプロジェクトにおける最新のベストプラクティスに従っています。

```text
/
  app/         # アプリケーションのルートとページファイル (App Router)
    layout.tsx # ルートレイアウト
    page.tsx   # ホームページ
    products/  # 商品関連ページ
  src/
    components/ # 再利用可能なUIコンポーネント
      layout/   # ページレイアウトに関連するコンポーネント
      ui/       # 汎用的なUIコンポーネント
    lib/        # ユーティリティ関数やサービス
      api/      # API通信関連のコード
      auth/     # 認証関連の機能
    types/      # TypeScriptの型定義
```

App Routerでは、`app`ディレクトリがルーティングの基盤となります。ページやレイアウト、ローディング状態やエラーページなどの特殊ファイルはここに配置します。一方、コンポーネントやユーティリティ、型定義などの再利用可能なコードは`src`ディレクトリに配置します。

この構造には以下の主要な利点があります：

1. **関心の分離**: UIコンポーネント、APIロジック、型定義が明確に分けられており、各ファイルが特定の責任を持つことで、コードベースの理解と保守が容易になります。これはソフトウェア設計の基本原則「単一責任の原則」に沿っています。

2. **スケーラビリティ**: プロジェクトが拡大しても、新しいコンポーネントやロジックを適切な場所に追加しやすい構造です。ファイル数が増加しても明確なディレクトリ構造がガイドとなり、コードの見通しを維持できます。

3. **再利用性の向上**: コンポーネントが論理的にグループ化されているため、再利用が促進されます。例えば、`layout`コンポーネントはアプリケーション全体で一貫したレイアウトを提供し、`ui`コンポーネントは複数の機能で再利用できます。

4. **絶対パスインポート**: `tsconfig.json`の設定で`@/*`パスエイリアスを使用することで、深い相対パス（`../../../../`など）を避け、コードの可読性とメンテナンス性を向上させています。

5. **産業界での標準**: この構造はVercel（Next.jsの開発会社）自身のテンプレートやサンプルプロジェクトでも推奨されており、多くの企業やオープンソースプロジェクトで採用されている標準的な手法です。

6. **将来の拡張への対応**: App Router（`app/`ディレクトリ）への移行など、Next.jsの将来の機能追加や変更に対して容易に適応できる柔軟な構造となっています。

この構造を採用することで、開発チームのオンボーディングも容易になり、一貫性のあるコーディング慣行を促進することができます。また、テストや自動化も、この明確な構造を活用して効率的に実装できます。

## 1.7. 【補足情報】

### 1.7.1. フロントエンドのパフォーマンス最適化

Next.jsを使用したフロントエンドでパフォーマンスを最適化するためのベストプラクティス：

1. **イメージの最適化**: `next/image`コンポーネントを使用して、自動的なリサイズやWebPなどの最適な形式での提供を行います。

2. **コード分割**: `next/dynamic`を使用したコンポーネントの動的インポートで、初期ロードを軽量化します。

3. **ルートレベルのキャッシュ**: Next.jsのISR (Incremental Static Regeneration)機能を活用し、静的ページの再ビルドの最適化を行います。

### 1.7.2. SWRによるデータフェッチの最適化

SWR (Stale-While-Revalidate) は、データフェッチを効率化するためのReactフックライブラリです。以下のような利点があります：

1. **自動キャッシュと更新**: キャッシュが古くなったデータを表示しながら、背後でデータを再取得します。

2. **複数ページでのデータ共有**: 同じデータを使用する複数のコンポーネント間でリクエストを重複させずに消去します。

3. **リアルタイム更新**: タブのフォーカスやネットワークの回復時に自動的にデータを再取得します。

### 1.7.3. レスポンシブデザインの実装

TailwindCSSを使用したレスポンシブデザインの実装に関するベストプラクティス：

1. **ブレークポイント修飾子**: `sm:`、`md:`、`lg:`などの修飾子を使用して、異なる画面サイズに対応したスタイルを適用します。

2. **グリッドレイアウト**: `grid-cols-1 md:grid-cols-3`のようなクラスを使用して、画面サイズに応じてグリッドの列数を変更します。

3. **フレックスボックス**: 複雑なレイアウトを構築するために、`flex`や`flex-wrap`などのクラスを活用します。

4. **可視性制御**: モバイルでは表示させず、デスクトップでのみ表示する要素には`hidden md:block`のようなクラスを適用します。

## 1.8. 【よくある問題と解決法】

### 1.8.1. Next.jsプロジェクトビルド時の型エラー

**問題**: TypeScriptの型エラーが原因で`npm run build`が失敗する。

**解決策**:

1. `tsconfig.json`で`noEmit`設定を確認する
2. エラーメッセージを確認し、型定義ファイルが必要な場合は作成する
3. 残りの型エラーを個別に解決する

```bash
# 型チェックを実行してエラーを確認
npx tsc --noEmit
```

### 1.8.2. Dockerビルド時の環境変数問題

**問題**: Dockerビルド時に環境変数が正しく反映されない。

**解決策**:

1. `NEXT_PUBLIC_`プレフィックスが付いていることを確認する
2. Dockerfile内で環境変数を正しく設定する
3. Docker Composeで環境変数を設定する

```bash
# Docker Composeで環境変数を設定
environment:
  - NEXT_PUBLIC_API_URL=http://localhost:3001/api
```

### 1.8.3. TailwindCSSのスタイルが反映されない問題

**問題**: TailwindCSSのクラスが正しく反映されず、スタイルが適用されない。

**解決策**:

1. `tailwind.config.js`が正しく設定されているか確認する
2. `globals.css`でTailwindCSSのディレクティブが正しくインポートされているか確認する
3. 開発サーバーを再起動する

```bash
# TailwindCSSの設定ファイルを確認
cat tailwind.config.js

# styles/globals.cssの内容を確認
cat src/app/globals.css

# Docker Composeでのコンテナ再起動
docker-compose -f docker-compose.frontend.yml restart frontend-customer frontend-admin
```

## 1.9. 【今日の重要なポイント】

1. **App Routerの採用**: Next.js 13以降で導入されたApp Routerを使用することで、より直感的なファイルベースのルーティングとReact Server Componentsの利点を活用できます。この最新のアーキテクチャパターンは、パフォーマンスとDX（開発者体験）の両面を向上させます。

2. **フロントエンドの分離**: 顧客向けと管理者向けのフロントエンドを別々のNext.jsプロジェクトとして分離することで、それぞれに最適化されたUIとUXを提供し、独立して開発・デプロイできるようになります。これにより、管理画面の変更が顧客向けサイトに影響を与えることがなくなります。

3. **クライアントコンポーネントとサーバーコンポーネントの区別**: `"use client"`ディレクティブを使用して、クライアントサイドの状態やインタラクティブな機能を必要とするコンポーネントを明示的に指定します。これにより、必要な場所でのみJavaScriptをクライアントに送信し、パフォーマンスを最適化できます。

4. **レスポンシブデザインの実装**: モバイルファーストのアプローチを採用し、さまざまな画面サイズに対応するUIを実現しています。TailwindCSSの`md:`などのブレークポイント修飾子を活用することで、レスポンシブなレイアウトを効率的に構築できます。

5. **認証と安全性**: 管理者向けフロントエンドには、認証ロジックを専用のモジュールに分離し、型定義を活用して型安全性を確保しています。開発段階ではモック認証を使用し、将来的に本番環境向けの認証システムへ置き換えやすい構造になっています。

## 1.10. 【次回の準備】

1. 作成したフロントエンドプロジェクトが正しく動作することを確認してください。両方のフロントエンドアプリケーションが起動し、基本的なUIが表示されることを確認しましょう。

2. 以下の概念について事前に学習しておいてください：
   - App Routerの仕組みとServer Components / Client Componentsの違い
   - CSS変数とTailwindCSSのカスタマイズ方法
   - Traefikの基本概念とホスト名ベースのルーティング

3. ChromeやFirefoxの開発者ツールの使い方を復習しておくと、次回のデバッグが容易になります。特にネットワークタブでのリクエスト確認や、Reactデベロッパーツールの基本操作を確認しておきましょう。

4. 次回は、Traefikを使ったリバースプロキシを設定し、`shop.localhost`と`admin.localhost`というホスト名ベースのルーティングを実装します。ローカルホストファイル（`/etc/hosts`または`C:\Windows\System32\drivers\etc\hosts`）を編集する権限があることを確認しておいてください。

## 1.11. 【.envrc サンプル】

開発環境で必要な環境変数のサンプルです。このファイルはリポジトリにコミットしないでください。

```bash
# .envrc
# このファイルはギットにはコミットしないでください

# バックエンドAPIのURL
export NEXT_PUBLIC_API_URL="http://localhost:8000/api"

# 開発サーバーのポート設定
export PORT=3000
export ADMIN_PORT=3010
```

各プロジェクトのルートディレクトリに`.env.local`ファイルを作成して使用することもできます：

```bash
# frontend-customer/.env.local
# このファイルはギットにはコミットしないでください
NEXT_PUBLIC_API_URL=http://localhost:3001/api
```

```bash
# frontend-admin/.env.local
# このファイルはギットにはコミットしないでください
NEXT_PUBLIC_API_URL=http://localhost:3001/api
```
