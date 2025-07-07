import axios from 'axios';
import { NextRequest, NextResponse } from 'next/server';

const BACKEND_API_URL = process.env.BACKEND_API_URL || 'http://backend-api:8000/api';

export async function GET(request: NextRequest) {
  try {
    const { searchParams } = new URL(request.url);

    // すべてのクエリパラメータを自動的に取得
    const params = Object.fromEntries(searchParams.entries());

    // backend-apiのhealth エンドポイントにリクエスト
    const response = await axios.get(`${BACKEND_API_URL}/health`, {
      params,
      timeout: 10000,
    });

    return NextResponse.json(response.data);
  } catch (error) {
    console.error('Health API Route error:', error);

    // axiosエラーの詳細なハンドリング
    if (axios.isAxiosError(error)) {
      const status = error.response?.status || 500;
      const message = error.response?.data?.message || error.message;

      return NextResponse.json(
        {
          error: 'Failed to fetch health status',
          message,
          code: 'FETCH_HEALTH_ERROR',
        },
        { status },
      );
    }

    return NextResponse.json(
      {
        error: 'Failed to fetch health status',
        message: 'Unknown error',
        code: 'FETCH_HEALTH_ERROR',
      },
      { status: 500 },
    );
  }
}
