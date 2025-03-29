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
