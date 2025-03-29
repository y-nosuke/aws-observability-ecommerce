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
