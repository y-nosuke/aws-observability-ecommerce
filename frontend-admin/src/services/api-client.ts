import axios from 'axios';

import { clearAuthData, getAuthToken } from '../lib/auth/auth';

// APIのベースURL（フロントエンドのAPI Routes）
const getApiBaseUrl = () => {
  // ブラウザ環境（クライアントサイド）
  if (typeof window !== 'undefined') {
    return process.env.NEXT_PUBLIC_API_URL || '/api';
  }

  // サーバー環境（SSR）での自分自身のAPI Routes呼び出し
  const baseUrl =
    process.env.NEXTAUTH_URL ||
    (process.env.VERCEL_URL
      ? `https://${process.env.VERCEL_URL}`
      : `http://localhost:${process.env.PORT || 3000}`);

  return `${baseUrl}/api`;
};

// Axiosインスタンスの作成
export const apiClient = axios.create({
  baseURL: getApiBaseUrl(),
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// リクエストインターセプター
apiClient.interceptors.request.use(
  (config) => {
    // デバッグ用ログ（開発環境のみ）
    if (process.env.NODE_ENV === 'development') {
      console.log(`API Request: ${config.method?.toUpperCase()} ${config.baseURL}${config.url}`);
    }

    // 管理者認証トークンをヘッダーに追加
    const token = getAuthToken();
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    console.error('Request interceptor error:', error);
    return Promise.reject(error);
  },
);

// レスポンスインターセプター
apiClient.interceptors.response.use(
  (response) => {
    // デバッグ用ログ（開発環境のみ）
    if (process.env.NODE_ENV === 'development') {
      console.log(`API Response: ${response.status} ${response.config.url}`);
    }
    return response;
  },
  (error) => {
    // エラーハンドリング
    console.error('API Error:', {
      message: error.message,
      config: error.config,
      response: error.response?.data,
      status: error.response?.status,
    });

    if (error.response) {
      // 認証エラーの場合はログアウト処理
      if (error.response.status === 401) {
        clearAuthData();
        // ログインページへリダイレクト（ブラウザ環境の場合）
        if (typeof window !== 'undefined') {
          window.location.href = '/login';
        }
      }
    }
    return Promise.reject(error);
  },
);
