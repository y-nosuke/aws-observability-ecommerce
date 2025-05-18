# 2.6. フロントエンドのログ収集

## 2.6.1. フロントエンドログの設計と考慮点

フロントエンドでのログ収集はバックエンドとは異なる課題と機会を持っています。適切に実装することで、ユーザー体験の向上やバグの早期発見、パフォーマンス最適化に大きく貢献します。

### 主要な考慮点

1. **クライアント・サーバー環境の二重性**
   - Next.jsのようなフレームワークではサーバーサイドとクライアントサイドの両方でログが必要
   - それぞれの環境で異なるアプローチが必要

2. **ネットワークとストレージの制約**
   - バックエンドと異なり、帯域幅とストレージには厳しい制限がある
   - すべてをログ記録するのではなく、重要な情報の選択が必要

3. **プライバシーとセキュリティ**
   - ユーザーの個人情報や機密データをログに含めない配慮
   - ブラウザ環境での安全なログ送信の確保

4. **ユーザーコンテキストの把握**
   - デバイス情報、ブラウザ情報、画面サイズなどの環境情報
   - ユーザーのインタラクションパターンとナビゲーション

5. **エラー検出と診断**
   - クライアントサイドのエラーを効果的に捕捉する仕組み
   - エラー発生時のコンテキスト情報の確保

### ログ設計の原則

- **必要性の精査**: 本当に必要なログのみを収集
- **構造化ログの採用**: JSON形式など解析しやすい形式を使用
- **一貫性の確保**: バックエンドとフロントエンド間で一貫したログフォーマット
- **段階的なログレベル**: ERROR、WARN、INFO、DEBUGの適切な使い分け
- **リクエスト追跡**: リクエストIDなどを使った、バックエンドとの連携

## 2.6.2. Next.jsでのpinoログライブラリの導入

[Pino](https://getpino.io/)はNode.js向けの超高速ロギングライブラリで、Next.jsアプリケーションでの使用に適しています。特に軽量で、JSONログ出力に最適化されています。

### インストール

```bash
npm install pino pino-pretty
# または
yarn add pino pino-pretty
```

### 基本設定

まず、共通のロガーモジュールを作成します：

```typescript
// src/lib/logger.ts
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

## 2.6.3. クライアント側とサーバー側のログ戦略

Next.jsの特性を活かした効果的なログ戦略を実装しましょう。

### サーバー側のログ

サーバーコンポーネントやAPI Routesでは、直接pinoロガーを使用できます：

```typescript
// src/app/api/products/route.ts
import { logger } from '@/lib/logger';
import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export async function GET(request: NextRequest) {
  const requestId = request.headers.get('x-request-id') || crypto.randomUUID();
  const log = logger.child({ requestId });

  log.info({ path: request.url }, 'Handling products request');

  try {
    // 商品データの取得ロジック
    const products = await fetchProducts();
    log.debug({ productCount: products.length }, 'Products fetched successfully');

    return NextResponse.json({ products });
  } catch (error) {
    log.error({ error: error.message }, 'Failed to fetch products');
    return NextResponse.json({ error: 'Failed to fetch products' }, { status: 500 });
  }
}
```

### クライアント側のログ

クライアントでは、ブラウザのコンソールとサーバーへの送信を組み合わせます：

```typescript
// src/lib/client-logger.ts
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

  // コンソールにも出力（開発時に便利）
  if (process.env.NODE_ENV !== 'production') {
    console[level](message, data);
  }

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

#### クライアントログ受信API

```typescript
// src/app/api/logs/route.ts
import { logger } from '@/lib/logger';
import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export async function POST(request: NextRequest) {
  try {
    const body = await request.json();
    const { logs } = body;

    if (Array.isArray(logs)) {
      logs.forEach(log => {
        const { level, message, data, context } = log;
        if (level && typeof logger[level] === 'function') {
          logger[level]({ ...data, ...context }, message);
        }
      });
    }

    return NextResponse.json({ success: true });
  } catch (error) {
    logger.error({ error: error.message }, 'Failed to process client logs');
    return NextResponse.json({ error: 'Failed to process logs' }, { status: 400 });
  }
}
```

## 2.6.4. フロントエンドエラーのキャプチャ方法

フロントエンドでのエラー捕捉は、品質向上の鍵となります。複数の方法を組み合わせて包括的なエラー捕捉を実現しましょう。

### グローバルエラーハンドリング

```typescript
// src/lib/error-tracking.ts
import { clientLogger } from './client-logger';

export function setupErrorTracking() {
  if (typeof window === 'undefined') return;

  // 未処理のPromiseエラーをキャプチャ
  window.addEventListener('unhandledrejection', (event) => {
    clientLogger.error('Unhandled Promise Rejection', {
      reason: event.reason?.toString(),
      stack: event.reason?.stack,
    });
  });

  // 未処理のJSエラーをキャプチャ
  window.addEventListener('error', (event) => {
    clientLogger.error('Uncaught JavaScript Error', {
      message: event.message,
      filename: event.filename,
      lineno: event.lineno,
      colno: event.colno,
      stack: event.error?.stack,
    });

    // イベントの伝播を止めない（他のハンドラーも実行可能に）
    return false;
  });
}
```

### コンポーネントレベルのエラーバウンダリ

React Error Boundaryを使用して、コンポーネントレベルでのエラーをキャプチャします：

```tsx
// src/components/ErrorBoundary.tsx
import React from 'react';
import { clientLogger } from '@/lib/client-logger';

interface ErrorBoundaryProps {
  fallback?: React.ReactNode;
  children: React.ReactNode;
  componentName: string;
}

interface ErrorBoundaryState {
  hasError: boolean;
  error?: Error;
}

export class ErrorBoundary extends React.Component<ErrorBoundaryProps, ErrorBoundaryState> {
  constructor(props: ErrorBoundaryProps) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error: Error) {
    return { hasError: true, error };
  }

  componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
    clientLogger.error(`Error in component: ${this.props.componentName}`, {
      error: error.toString(),
      stack: error.stack,
      componentStack: errorInfo.componentStack,
    });
  }

  render() {
    if (this.state.hasError) {
      return this.props.fallback || (
        <div className="error-container">
          <h2>Something went wrong</h2>
          <p>Please try refreshing the page</p>
        </div>
      );
    }

    return this.props.children;
  }
}
```

使用例：

```tsx
// src/app/products/page.tsx
import { ErrorBoundary } from '@/components/ErrorBoundary';
import ProductList from '@/components/ProductList';

export default function ProductsPage() {
  return (
    <div className="products-page">
      <h1>Our Products</h1>
      <ErrorBoundary componentName="ProductList">
        <ProductList />
      </ErrorBoundary>
    </div>
  );
}
```

## 2.6.5. グローバルエラーハンドラーの実装

Next.jsでは、アプリケーション全体のエラーハンドリングのためのメカニズムを提供しています。これを使ってエラーログとフォールバックUIを実装します。

### グローバルエラーページ

```tsx
// src/app/error.tsx
'use client';

import { useEffect } from 'react';
import { clientLogger } from '@/lib/client-logger';

export default function GlobalError({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    // エラー発生時にログを記録
    clientLogger.error('Global error caught', {
      message: error.message,
      stack: error.stack,
      digest: error.digest,
    });
  }, [error]);

  return (
    <html>
      <body>
        <div className="error-container">
          <h2>Something went wrong!</h2>
          <button
            onClick={() => reset()}
          >
            Try again
          </button>
        </div>
      </body>
    </html>
  );
}
```

### 特定ルートのエラーハンドリング

```tsx
// src/app/products/error.tsx
'use client';

import { useEffect } from 'react';
import { clientLogger } from '@/lib/client-logger';

export default function ProductsError({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
  useEffect(() => {
    clientLogger.error('Error in products page', {
      message: error.message,
      stack: error.stack,
    });
  }, [error]);

  return (
    <div className="route-error">
      <h3>Sorry, there was a problem loading the products</h3>
      <button onClick={() => reset()}>Try again</button>
    </div>
  );
}
```

### API エラーハンドリング

ユーザー体験を損なわずに、APIエラーを適切に処理するためのカスタムフックを実装します：

```typescript
// src/hooks/useFetch.ts
import { useState, useEffect } from 'react';
import { clientLogger } from '@/lib/client-logger';

interface FetchState<T> {
  data: T | null;
  isLoading: boolean;
  error: Error | null;
}

export function useFetch<T>(url: string, options?: RequestInit) {
  const [state, setState] = useState<FetchState<T>>({
    data: null,
    isLoading: true,
    error: null,
  });

  useEffect(() => {
    const controller = new AbortController();
    const { signal } = controller;

    const fetchData = async () => {
      try {
        const startTime = performance.now();

        const response = await fetch(url, {
          ...options,
          signal,
          headers: {
            ...options?.headers,
          },
        });

        const endTime = performance.now();
        const duration = endTime - startTime;

        // API呼び出しのパフォーマンスをログ
        clientLogger.info(`API Call to ${url}`, {
          duration: `${duration.toFixed(2)}ms`,
          status: response.status,
        });

        if (!response.ok) {
          const errorText = await response.text();
          throw new Error(`API error ${response.status}: ${errorText}`);
        }

        const data = await response.json();
        setState({ data, isLoading: false, error: null });
      } catch (error) {
        if (error.name !== 'AbortError') {
          clientLogger.error(`API fetch error for ${url}`, {
            message: error.message,
            stack: error.stack,
          });
          setState({ data: null, isLoading: false, error });
        }
      }
    };

    fetchData();

    return () => controller.abort();
  }, [url, JSON.stringify(options)]);

  return state;
}
```

## 2.6.6. ユーザーインタラクションのログ記録

ユーザーインタラクションのログ記録は、UX改善や問題診断に役立ちます。ただし、プライバシーに配慮してください。

### 共通のイベントトラッカー

```typescript
// src/lib/event-tracking.ts
import { clientLogger } from './client-logger';

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
  // 個人情報をフィルタリング（例：メールアドレス）
  const sanitizedMetadata = { ...metadata };

  // JSONに変換できるデータのみを許可
  try {
    // 循環参照をチェック
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

### 実装例

```tsx
// src/components/ProductCard.tsx
import { trackEvent } from '@/lib/event-tracking';

interface ProductCardProps {
  id: string;
  name: string;
  price: number;
  // その他のプロパティ
}

export default function ProductCard({ id, name, price }: ProductCardProps) {
  const handleAddToCart = () => {
    // 商品をカートに追加するロジック

    // イベントをトラッキング
    trackEvent({
      category: 'interaction',
      action: 'add_to_cart',
      label: name,
      value: price,
      metadata: {
        productId: id,
        productPrice: price,
      },
    });
  };

  const handleProductClick = () => {
    trackEvent({
      category: 'navigation',
      action: 'product_detail_view',
      label: name,
      metadata: {
        productId: id,
      },
    });
  };

  return (
    <div className="product-card">
      <h3 onClick={handleProductClick}>{name}</h3>
      <p>${price.toFixed(2)}</p>
      <button onClick={handleAddToCart}>Add to Cart</button>
    </div>
  );
}
```

### ナビゲーショントラッキング

Next.jsのルーティングイベントを使用して、ページナビゲーションを追跡します：

```tsx
// src/app/layout.tsx
'use client';

import { useEffect } from 'react';
import { usePathname, useSearchParams } from 'next/navigation';
import { trackEvent } from '@/lib/event-tracking';
import { setupErrorTracking } from '@/lib/error-tracking';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const pathname = usePathname();
  const searchParams = useSearchParams();

  useEffect(() => {
    // エラートラッキングをセットアップ
    setupErrorTracking();
  }, []);

  useEffect(() => {
    // ページビューをトラッキング
    trackEvent({
      category: 'navigation',
      action: 'page_view',
      label: pathname,
      metadata: {
        path: pathname,
        query: Object.fromEntries(searchParams.entries()),
        referrer: document.referrer,
      },
    });
  }, [pathname, searchParams]);

  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  );
}
```

## 2.6.7. パフォーマンス関連情報の収集

Webビタルなどのパフォーマンス情報を収集して、ユーザー体験を継続的に向上させましょう。

### Web Vitalsの収集

```typescript
// src/lib/performance-tracking.ts
import { clientLogger } from './client-logger';
import { trackEvent } from './event-tracking';

interface WebVitalMetric {
  id: string;
  name: string;
  value: number;
  rating: 'good' | 'needs-improvement' | 'poor';
  delta: number;
}

// 結果をWeb Vitalsの評価範囲に基づいて分類
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

export function setupPerformanceTracking() {
  if (typeof window === 'undefined') return;

  // ページロード完了時の基本メトリクスを収集
  window.addEventListener('load', () => {
    // 少し遅延させてすべてのリソースが読み込まれるのを待つ
    setTimeout(() => {
      if (performance && 'getEntriesByType' in performance) {
        const navigation = performance.getEntriesByType('navigation')[0] as PerformanceNavigationTiming;

        if (navigation) {
          trackEvent({
            category: 'performance',
            action: 'page_load',
            metadata: {
              dnsTime: Math.round(navigation.domainLookupEnd - navigation.domainLookupStart),
              tcpTime: Math.round(navigation.connectEnd - navigation.connectStart),
              ttfb: Math.round(navigation.responseStart - navigation.requestStart),
              domLoadTime: Math.round(navigation.domComplete - navigation.domLoading),
              loadTime: Math.round(navigation.loadEventEnd - navigation.loadEventStart),
              totalTime: Math.round(navigation.loadEventEnd),
            },
          });
        }
      }
    }, 0);
  });

  // Web Vitalsと互換性のあるメトリクス収集関数
  function reportWebVital(metric: WebVitalMetric) {
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

  // Next.jsのreportWebVitalsフック用に公開
  return reportWebVital;
}

// Custom Elementsのパフォーマンス測定
export function measureElementPerformance(elementId: string, eventName: string, metadata: Record<string, any> = {}) {
  return () => {
    // 要素の表示時間を測定
    if (typeof window !== 'undefined' && window.performance) {
      const now = performance.now();
      trackEvent({
        category: 'performance',
        action: eventName,
        label: elementId,
        value: Math.round(now),
        metadata,
      });
    }
  };
}
```

### Next.jsでのWeb Vitals統合

```tsx
// src/app/layout.tsx（既存のレイアウトに追加）
import { setupPerformanceTracking } from '@/lib/performance-tracking';

export function reportWebVitals(metric) {
  // Web Vitalsレポート設定
  const reportWebVital = setupPerformanceTracking();
  reportWebVital(metric);
}
```

### コンポーネントレベルのパフォーマンス測定

```tsx
// src/components/ProductList.tsx
import { useEffect, useRef } from 'react';
import { measureElementPerformance } from '@/lib/performance-tracking';

export default function ProductList({ products }) {
  const listRef = useRef(null);

  useEffect(() => {
    if (listRef.current) {
      // 製品リストが表示された時間を測定
      const onLoad = measureElementPerformance('product-list', 'component_loaded', {
        productCount: products.length,
      });
      onLoad();
    }
  }, [products]);

  return (
    <div className="product-list" ref={listRef}>
      {/* 製品リストの内容 */}
    </div>
  );
}
```

### まとめ

フロントエンドのログ収集は、アプリケーションの品質向上とユーザー体験の改善に不可欠です。ここで紹介した技術と方法を組み合わせることで、包括的なログ基盤を構築できます。

- **効率的なログ設計**: 必要最小限のデータを優先し、バッチ処理でサーバー負荷を軽減
- **エラー検出**: 複数レベルでのエラーキャプチャにより問題を早期に発見
- **ユーザー行動分析**: 意味のあるインタラクション情報を収集しUXを向上
- **パフォーマンスモニタリング**: Web Vitalsなど重要指標の継続的な追跡

次のセクションでは、これらのフロントエンドログをCloudWatch Logsに効率的に送信する方法について学びます。
