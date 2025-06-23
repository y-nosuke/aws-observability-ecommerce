"use client";

import { useCallback, useEffect, useState } from "react";
import {
  getStatusBgColor,
  getStatusColor,
  getStatusIconType,
} from "../../lib/utils/health";
import {
  determineOverallStatus,
  healthApi,
  parseComponentStatuses,
} from "../../services/health/api";
import { ComponentStatus, OverallStatus } from "../../services/health/types";
import StatusIcon from "./StatusIcon";

interface SystemHealthStatusProps {
  /** コンポーネントの表示モード */
  mode?: "compact" | "detailed";
  /** 自動更新間隔（分）。0の場合は自動更新なし */
  autoRefreshMinutes?: number;
}

export default function SystemHealthStatus({
  mode = "compact",
  autoRefreshMinutes = 0,
}: SystemHealthStatusProps) {
  const [components, setComponents] = useState<ComponentStatus[]>([]);
  const [overallStatus, setOverallStatus] = useState<OverallStatus | null>(
    null
  );
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [lastRefresh, setLastRefresh] = useState<string>("");

  // キャッシュ設定
  const CACHE_KEY = "admin_health_status_cache";
  const CACHE_DURATION = 5 * 60 * 1000; // 5分（ミリ秒）

  /**
   * キャッシュからデータを取得
   */
  const getCachedData = useCallback(() => {
    try {
      const cached = sessionStorage.getItem(CACHE_KEY);
      if (!cached) return null;

      const { data, timestamp } = JSON.parse(cached);
      const now = Date.now();

      // キャッシュが有効期限内かチェック
      if (now - timestamp < CACHE_DURATION) {
        return data;
      }

      // 期限切れの場合はキャッシュを削除
      sessionStorage.removeItem(CACHE_KEY);
      return null;
    } catch (error) {
      console.warn("Failed to read health status cache:", error);
      sessionStorage.removeItem(CACHE_KEY);
      return null;
    }
  }, [CACHE_KEY, CACHE_DURATION]);

  /**
   * データをキャッシュに保存
   */
  const setCachedData = useCallback(
    (data: {
      components: ComponentStatus[];
      overallStatus: OverallStatus;
      lastRefresh: string;
    }) => {
      try {
        const cacheData = {
          data,
          timestamp: Date.now(),
        };
        sessionStorage.setItem(CACHE_KEY, JSON.stringify(cacheData));
      } catch (error) {
        console.warn("Failed to save health status cache:", error);
      }
    },
    [CACHE_KEY]
  );

  /**
   * ヘルスステータスを取得する関数（キャッシュ対応）
   */
  const fetchHealthStatus = useCallback(
    async (forceRefresh: boolean = false) => {
      // 常にローディング状態から開始
      setIsLoading(true);
      setError(null);

      // 強制更新でない場合は、まずキャッシュをチェック
      if (!forceRefresh) {
        const cachedData = getCachedData();
        if (cachedData) {
          // キャッシュデータがある場合でも一瞬ローディング表示してから表示
          setTimeout(() => {
            setComponents(cachedData.components);
            setOverallStatus(cachedData.overallStatus);
            setLastRefresh(cachedData.lastRefresh);
            setIsLoading(false);
          }, 200); // 200ms のローディング表示
          return;
        }
      }

      try {
        const healthResponse = await healthApi.getHealthStatus();
        const componentStatuses = parseComponentStatuses(healthResponse);
        const overall = determineOverallStatus(componentStatuses);
        const refreshTime = new Date().toLocaleString("ja-JP");

        setComponents(componentStatuses);
        setOverallStatus(overall);
        setLastRefresh(refreshTime);

        // データをキャッシュに保存
        setCachedData({
          components: componentStatuses,
          overallStatus: overall,
          lastRefresh: refreshTime,
        });
      } catch (err) {
        const errorMessage =
          err instanceof Error ? err.message : "不明なエラーが発生しました";
        setError(errorMessage);
        console.error("Health status fetch error:", err);
      } finally {
        setIsLoading(false);
      }
    },
    [getCachedData, setCachedData]
  );

  /**
   * 手動更新ボタンのハンドラー（強制更新）
   */
  const handleRefresh = useCallback(() => {
    fetchHealthStatus(true); // 強制更新でキャッシュを無視
  }, [fetchHealthStatus]);

  // 初期ロード時にヘルスステータスを取得（キャッシュを使用）
  useEffect(() => {
    fetchHealthStatus(false); // キャッシュを使用
  }, [fetchHealthStatus]);

  // 自動更新の設定（キャッシュを使用）
  useEffect(() => {
    if (autoRefreshMinutes > 0) {
      const interval = setInterval(() => {
        fetchHealthStatus(false); // 自動更新時もキャッシュを使用
      }, autoRefreshMinutes * 60 * 1000);
      return () => clearInterval(interval);
    }
  }, [autoRefreshMinutes, fetchHealthStatus]);

  /**
   * コンパクトモードの表示
   */
  const renderCompactMode = () => (
    <div className="bg-indigo-700/20 rounded-lg p-4">
      {/* ヘッダー部分 */}
      <div className="flex items-center justify-between mb-3">
        <div className="flex items-center">
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
              d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
            />
          </svg>
          <p className="text-sm font-medium text-indigo-200">システムヘルス</p>
        </div>

        {/* 更新ボタン - 改善版 */}
        <div className="flex items-center space-x-2">
          {isLoading && (
            <span className="text-xs text-indigo-300 animate-pulse">
              更新中...
            </span>
          )}
          <button
            onClick={handleRefresh}
            disabled={isLoading}
            className="flex items-center text-indigo-300 hover:text-white transition-colors disabled:opacity-50"
            title="ヘルスステータスを更新"
          >
            <svg
              className={`w-4 h-4 ${isLoading ? "animate-spin" : ""}`}
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
              />
            </svg>
            {isLoading && (
              <span className="ml-1 text-xs">更新中</span>
            )}
          </button>
        </div>
      </div>

      {/* エラー表示 */}
      {error && (
        <div className="mb-3 p-2 bg-red-100 border border-red-300 rounded text-red-700 text-xs">
          {error}
        </div>
      )}

      {/* ステータス表示 */}
      {overallStatus && !error && (
        <div>
          <div className="flex items-center mb-2">
            <span className={getStatusColor(overallStatus.status)}>
              <StatusIcon type={getStatusIconType(overallStatus.status)} />
            </span>
            <span className="ml-2 text-xs text-gray-300">
              {overallStatus.message}
            </span>
          </div>

          {/* コンポーネント一覧（コンパクト表示） */}
          <div className="grid grid-cols-2 gap-1 mb-2">
            {components.map((component) => (
              <div key={component.name} className="flex items-center">
                <span className={`${getStatusColor(component.status)} mr-1`}>
                  <StatusIcon type={getStatusIconType(component.status)} />
                </span>
                <span className="text-xs text-gray-300 truncate">
                  {component.displayName}
                </span>
              </div>
            ))}
          </div>

          <p className="text-xs text-gray-400">最終更新：{lastRefresh}</p>
        </div>
      )}

    </div>
  );

  /**
   * 詳細モードの表示
   */
  const renderDetailedMode = () => (
    <div className="bg-indigo-700/20 rounded-lg shadow-md p-6">
      {/* ヘッダー部分 */}
      <div className="flex items-center justify-between mb-6">
        <h3 className="text-lg font-semibold text-indigo-200">
          システムヘルスステータス
        </h3>
        <div className="flex items-center space-x-3">
          {isLoading && (
            <div className="flex items-center text-indigo-600">
              <svg
                className="w-4 h-4 animate-spin mr-2"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  className="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  strokeWidth="4"
                ></circle>
                <path
                  className="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 0 1 8-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 0 1 4 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
              <span className="text-sm">更新中...</span>
            </div>
          )}
          <button
            onClick={handleRefresh}
            disabled={isLoading}
            className="flex items-center px-3 py-2 text-sm bg-indigo-600 text-white rounded-md hover:bg-indigo-700 disabled:opacity-50 transition-all"
          >
            <svg
              className={`w-4 h-4 mr-2 ${isLoading ? "animate-spin" : ""}`}
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
              />
            </svg>
            {isLoading ? "更新中" : "更新"}
          </button>
        </div>
      </div>

      {/* エラー表示 */}
      {error && (
        <div className="mb-6 p-4 bg-red-50 border border-red-200 rounded-md">
          <div className="flex">
            <svg
              className="w-5 h-5 text-red-400"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                fillRule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                clipRule="evenodd"
              />
            </svg>
            <div className="ml-3">
              <h4 className="text-sm font-medium text-red-800">
                エラーが発生しました
              </h4>
              <p className="mt-1 text-sm text-red-700">{error}</p>
            </div>
          </div>
        </div>
      )}

      {/* 全体ステータス */}
      {overallStatus && !error && (
        <div
          className={`mb-6 p-4 rounded-md ${getStatusBgColor(
            overallStatus.status
          )}`}
        >
          <div className="flex items-center">
            <span className={getStatusColor(overallStatus.status)}>
              <StatusIcon type={getStatusIconType(overallStatus.status)} />
            </span>
            <div className="ml-3">
              <h4 className="text-sm font-medium text-gray-900">
                {overallStatus.status === "healthy"
                  ? "正常"
                  : overallStatus.status === "warning"
                  ? "警告"
                  : "エラー"}
              </h4>
              <p className="text-sm text-gray-700">{overallStatus.message}</p>
            </div>
          </div>
        </div>
      )}

      {/* コンポーネント詳細一覧 */}
      {components.length > 0 && !error && (
        <div>
          <h4 className="text-sm font-medium text-indigo-200 mb-3">
            コンポーネント詳細
          </h4>
          <div className="space-y-3">
            {components.map((component) => (
              <div
                key={component.name}
                className="flex items-center justify-between p-3 bg-indigo-600/10 rounded-md"
              >
                <div className="flex items-center">
                  <span className={getStatusColor(component.status)}>
                    <StatusIcon type={getStatusIconType(component.status)} />
                  </span>
                  <span className="ml-3 text-sm font-medium text-indigo-200">
                    {component.displayName}
                  </span>
                </div>
                <div className="text-right">
                  <span
                    className={`text-sm font-medium ${getStatusColor(
                      component.status
                    )}`}
                  >
                    {component.status === "ok" ? "正常" : "エラー"}
                  </span>
                  {component.message && (
                    <p className="text-xs text-gray-400 mt-1">
                      {component.message}
                    </p>
                  )}
                </div>
              </div>
            ))}
          </div>
        </div>
      )}

      {/* 最終更新時刻 */}
      {lastRefresh && !error && (
        <div className="mt-6 pt-4 border-t border-indigo-600/20">
          <p className="text-xs text-gray-400">最終更新：{lastRefresh}</p>
        </div>
      )}

    </div>
  );

  return mode === "compact" ? renderCompactMode() : renderDetailedMode();
}
