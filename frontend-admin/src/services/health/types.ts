/**
 * ヘルスチェックAPIの型定義
 * frontend-customer/services/products/types.ts のパターンに準拠
 */

// メモリ統計情報
export interface MemoryStats {
  allocated: number;
  total: number;
  system: number;
}

// システム情報
export interface SystemInfo {
  memory: MemoryStats;
  goroutines: number;
}

// システムリソース情報
export interface SystemResources {
  system: SystemInfo;
}

// バックエンドAPIからのヘルスレスポンス（openapi.yamlに対応）
export interface HealthResponse {
  status: string;
  timestamp: string;
  version: string;
  uptime: number;
  resources: SystemResources;
  components: Record<string, string>;
}

// コンポーネントのステータス（フロントエンド用の変換後データ）
export interface ComponentStatus {
  name: string;
  status: "ok" | "error" | "warning";
  message?: string;
  displayName: string;
}

// システム全体のステータス
export interface OverallStatus {
  status: "healthy" | "warning" | "error";
  message: string;
  lastUpdated: string;
}

// ダッシュボード用の統計情報
export interface HealthStats {
  totalComponents: number;
  healthyComponents: number;
  warningComponents: number;
  errorComponents: number;
  uptime: number;
  lastCheckTime: string;
}

// ヘルスチェック結果の詳細情報
export interface HealthDetails {
  response: HealthResponse;
  components: ComponentStatus[];
  overall: OverallStatus;
  stats: HealthStats;
}

// APIエラーの型定義
export interface HealthCheckError {
  message: string;
  timestamp: string;
  code?: string;
  details?: unknown;
}

// ヘルスチェックのパラメータ型
export interface HealthCheckParams {
  checks?: string;
  timeout?: number;
  includeDetails?: boolean;
}

// コンポーネントフィルタリングの型
export type ComponentStatusFilter = "all" | "ok" | "error" | "warning";

// コンポーネントソートの型
export type ComponentSortOption = "name" | "status" | "displayName";

// ヘルスチェック設定
export interface HealthCheckConfig {
  autoRefresh: boolean;
  refreshInterval: number; // milliseconds
  enableNotifications: boolean;
  criticalComponents: string[];
}

// リアルタイム更新用の型
export interface HealthUpdateEvent {
  type: "status_change" | "component_update" | "system_restart";
  timestamp: string;
  component?: string;
  oldStatus?: ComponentStatus["status"];
  newStatus?: ComponentStatus["status"];
  message?: string;
}
