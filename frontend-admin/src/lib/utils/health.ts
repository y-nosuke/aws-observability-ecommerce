/**
 * ヘルスステータス関連のユーティリティ関数
 */

export type HealthStatus = 'ok' | 'error' | 'warning' | 'healthy';

/**
 * ステータスに対応するCSSクラスを取得
 */
export const getStatusColor = (status: HealthStatus): string => {
  switch (status) {
    case 'ok':
    case 'healthy':
      return 'text-green-500';
    case 'warning':
      return 'text-yellow-500';
    case 'error':
      return 'text-red-500';
    default:
      return 'text-gray-500';
  }
};

/**
 * ステータスに対応する背景色クラスを取得
 */
export const getStatusBgColor = (status: HealthStatus): string => {
  switch (status) {
    case 'ok':
    case 'healthy':
      return 'bg-green-100';
    case 'warning':
      return 'bg-yellow-100';
    case 'error':
      return 'bg-red-100';
    default:
      return 'bg-gray-100';
  }
};

/**
 * ステータスに対応するアイコン種別を取得
 */
export const getStatusIconType = (
  status: HealthStatus,
): 'success' | 'warning' | 'error' | 'default' => {
  switch (status) {
    case 'ok':
    case 'healthy':
      return 'success';
    case 'warning':
      return 'warning';
    case 'error':
      return 'error';
    default:
      return 'default';
  }
};
