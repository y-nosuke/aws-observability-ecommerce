import axios from 'axios';
import { NextRequest, NextResponse } from 'next/server';

const BACKEND_API_URL = process.env.BACKEND_API_URL || 'http://backend-api:8000/api';

export async function GET(request: NextRequest) {
  try {
    const { searchParams } = new URL(request.url);

    // すべてのクエリパラメータを自動的に取得
    const params = Object.fromEntries(searchParams.entries());

    const response = await axios.get(`${BACKEND_API_URL}/categories`, {
      params,
      timeout: 10000,
    });

    return NextResponse.json(response.data);
  } catch (error) {
    console.error('API Route error:', error);

    if (axios.isAxiosError(error)) {
      const status = error.response?.status || 500;
      const message = error.response?.data?.message || error.message;

      return NextResponse.json(
        {
          error: 'Failed to fetch categories',
          message,
          code: 'FETCH_CATEGORIES_ERROR',
        },
        { status },
      );
    }

    return NextResponse.json(
      {
        error: 'Failed to fetch categories',
        message: 'Unknown error',
        code: 'FETCH_CATEGORIES_ERROR',
      },
      { status: 500 },
    );
  }
}
