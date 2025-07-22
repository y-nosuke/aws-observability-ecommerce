import axios from 'axios';

// APIのベースURL（フロントエンドのAPI Routes）
const getApiBaseUrl = () => {
  // ブラウザ環境（クライアントサイド）
  if (typeof window !== 'undefined') {
    return process.env.NEXT_PUBLIC_API_URL || '/api';
  }

  // サーバー環境（SSR）
  const baseUrl = process.env.VERCEL_URL
    ? `https://${process.env.VERCEL_URL}/api`
    : process.env.BACKEND_API_URL;

  return `${baseUrl}`;
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

    return Promise.reject(error);
  },
);
