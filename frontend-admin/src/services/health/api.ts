import { apiClient } from '../api-client';
import {
  ComponentSortOption,
  ComponentStatus,
  ComponentStatusFilter,
  HealthResponse,
  OverallStatus,
} from './types';

// ヘルスチェック関連のAPI関数
export const healthApi = {
  /**
   * システムヘルスステータスを取得
   * @param checks 実行するチェック項目（カンマ区切り）
   * @returns ヘルスレスポンス
   */
  async getHealthStatus(checks: string = 'db,s3,iam'): Promise<HealthResponse> {
    const response = await apiClient.get('/health', {
      params: { checks },
    });
    return response.data;
  },

  /**
   * 基本的なヘルスチェック（チェック項目指定なし）
   */
  async getBasicHealth(): Promise<HealthResponse> {
    const response = await apiClient.get('/health');
    return response.data;
  },

  /**
   * 特定のコンポーネントのヘルスチェック
   * @param component チェックするコンポーネント名
   */
  async getComponentHealth(component: string): Promise<HealthResponse> {
    const response = await apiClient.get('/health', {
      params: { checks: component },
    });
    return response.data;
  },
};

// Server Components用のfetch関数
export async function fetchHealthStatus(checks?: string): Promise<HealthResponse | null> {
  try {
    return await healthApi.getHealthStatus(checks);
  } catch (error) {
    console.error('Failed to fetch health status:', error);
    return null;
  }
}

// 基本的なヘルスチェック取得
export async function fetchBasicHealth(): Promise<HealthResponse | null> {
  try {
    return await healthApi.getBasicHealth();
  } catch (error) {
    console.error('Failed to fetch basic health:', error);
    return null;
  }
}

// システムステータスの詳細分析
export async function fetchSystemStatus(): Promise<{
  health: HealthResponse | null;
  components: ComponentStatus[];
  overall: OverallStatus;
}> {
  try {
    const health = await healthApi.getHealthStatus();
    const components = parseComponentStatuses(health);
    const overall = determineOverallStatus(components);

    return { health, components, overall };
  } catch (error) {
    console.error('Failed to fetch system status:', error);
    return {
      health: null,
      components: [],
      overall: {
        status: 'error',
        message: 'システムステータスの取得に失敗しました',
        lastUpdated: new Date().toLocaleString('ja-JP'),
      },
    };
  }
}

// ユーティリティ関数

/**
 * ヘルスレスポンスをコンポーネントステータスに変換
 * @param healthResponse ヘルスレスポンス
 * @returns コンポーネントステータスの配列
 */
export function parseComponentStatuses(healthResponse: HealthResponse): ComponentStatus[] {
  const componentMapping: Record<string, string> = {
    api_server: 'APIサーバー',
    database: 'データベース',
    s3_connectivity: 'S3接続',
    iam_auth: 'IAM認証',
  };

  return Object.entries(healthResponse.components).map(([key, value]) => {
    // バックエンドからのレスポンス形式: "ok" または "ng: エラーメッセージ"
    const isError = value.startsWith('ng:');
    const status = isError ? 'error' : 'ok';
    const message = isError ? value.replace('ng: ', '') : undefined;

    return {
      name: key,
      status: status as 'ok' | 'error' | 'warning',
      message,
      displayName: componentMapping[key] || key,
    };
  });
}

/**
 * 全体的なシステムステータスを判定
 * @param components コンポーネントステータスの配列
 * @returns 全体ステータス
 */
export function determineOverallStatus(components: ComponentStatus[]): OverallStatus {
  const errorComponents = components.filter((c) => c.status === 'error');
  const warningComponents = components.filter((c) => c.status === 'warning');

  if (errorComponents.length > 0) {
    return {
      status: 'error',
      message: `${errorComponents.length}個のコンポーネントでエラーが発生しています`,
      lastUpdated: new Date().toLocaleString('ja-JP'),
    };
  }

  if (warningComponents.length > 0) {
    return {
      status: 'warning',
      message: `${warningComponents.length}個のコンポーネントで警告が発生しています`,
      lastUpdated: new Date().toLocaleString('ja-JP'),
    };
  }

  return {
    status: 'healthy',
    message: 'すべてのコンポーネントが正常に稼働しています',
    lastUpdated: new Date().toLocaleString('ja-JP'),
  };
}

// フィルタリング・ソート用のユーティリティ関数

/**
 * コンポーネントをステータスでフィルタリング
 */
export function filterComponentsByStatus(
  components: ComponentStatus[],
  status: ComponentStatusFilter = 'all',
): ComponentStatus[] {
  if (status === 'all') {
    return components;
  }
  return components.filter((component) => component.status === status);
}

/**
 * コンポーネントをソート
 */
export function sortComponents(
  components: ComponentStatus[],
  sortBy: ComponentSortOption = 'name',
): ComponentStatus[] {
  return [...components].sort((a, b) => {
    if (sortBy === 'status') {
      // エラー > 警告 > 正常 の順でソート
      const statusOrder = { error: 0, warning: 1, ok: 2 };
      return statusOrder[a.status] - statusOrder[b.status];
    }

    const keyA = a[sortBy].toLowerCase();
    const keyB = b[sortBy].toLowerCase();
    return keyA.localeCompare(keyB);
  });
}
