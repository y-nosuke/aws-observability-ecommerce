// 管理者認証のためのユーティリティ関数

// 認証トークンの保存キー
const AUTH_TOKEN_KEY = 'admin_auth_token';
const USER_DATA_KEY = 'admin_user_data';

// ユーザーデータの型定義
interface UserData {
  id: string;
  name: string;
  email: string;
  role: string;
}

// 認証トークンを取得
export const getAuthToken = (): string | null => {
  if (typeof window === 'undefined') {
    return null;
  }
  return localStorage.getItem(AUTH_TOKEN_KEY);
};

// ユーザーデータを取得
export const getUserData = (): UserData | null => {
  if (typeof window === 'undefined') {
    return null;
  }
  const userData = localStorage.getItem(USER_DATA_KEY);
  if (userData) {
    try {
      return JSON.parse(userData);
    } catch {
      return null;
    }
  }
  return null;
};

// 認証情報を保存
export const setAuthData = (token: string, userData: UserData): void => {
  if (typeof window === 'undefined') {
    return;
  }
  localStorage.setItem(AUTH_TOKEN_KEY, token);
  localStorage.setItem(USER_DATA_KEY, JSON.stringify(userData));
};

// 認証情報をクリア
export const clearAuthData = (): void => {
  if (typeof window === 'undefined') {
    return;
  }
  localStorage.removeItem(AUTH_TOKEN_KEY);
  localStorage.removeItem(USER_DATA_KEY);
};

// ログイン状態の確認
export const isAuthenticated = (): boolean => {
  return !!getAuthToken();
};

// モック認証（開発用）
export const mockLogin = async (
  email: string,
  password: string,
): Promise<{ token: string; user: UserData }> => {
  // 開発環境用の簡易認証（本番環境では使用しないこと）
  if (email === 'admin@example.com' && password === 'password') {
    const mockToken = 'mock-jwt-token-for-development-only';
    const mockUserData = {
      id: '1',
      name: 'Admin User',
      email: 'admin@example.com',
      role: 'admin',
    };
    setAuthData(mockToken, mockUserData);
    return {
      token: mockToken,
      user: mockUserData,
    };
  }
  throw new Error('認証に失敗しました。メールアドレスまたはパスワードが間違っています。');
};
