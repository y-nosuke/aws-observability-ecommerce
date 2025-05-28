import axios from "axios";

// APIのベースURL（フロントエンドのAPI Routes）
const getApiBaseUrl = () => {
  // ブラウザ環境（クライアントサイド）
  if (typeof window !== "undefined") {
    return process.env.NEXT_PUBLIC_API_URL || "/api";
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
    "Content-Type": "application/json",
  },
});

// リクエストインターセプター
apiClient.interceptors.request.use(
  (config) => {
    // トークンがある場合はヘッダーに追加（ブラウザ環境でのみ）
    if (typeof window !== "undefined") {
      const token = localStorage.getItem("auth_token");
      if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
      }
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
      // 認証エラーの場合（ブラウザ環境でのみ）
      if (error.response.status === 401 && typeof window !== "undefined") {
        // ログアウト処理など
        localStorage.removeItem("auth_token");
      }
    }
    return Promise.reject(error);
  }
);
