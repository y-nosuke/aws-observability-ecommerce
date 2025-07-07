import axios from 'axios';
import { NextRequest, NextResponse } from 'next/server';

const BACKEND_API_URL = process.env.BACKEND_API_URL || 'http://backend-api:8000/api';

export async function GET(request: NextRequest, { params }: { params: { id: string } }) {
  try {
    // クエリパラメータを取得
    const searchParams = request.nextUrl.searchParams;
    const page = searchParams.get('page');
    const pageSize = searchParams.get('pageSize');
    const { id: categoryId } = await params;

    // バックエンドAPIにリクエスト
    const response = await axios.get(`${BACKEND_API_URL}/categories/${categoryId}/products`, {
      params: {
        page,
        pageSize,
      },
    });

    // レスポンスを返す
    return NextResponse.json(response.data);
  } catch (error) {
    console.error('Failed to fetch products by category:', error);
    return NextResponse.json({ error: 'Failed to fetch products by category' }, { status: 500 });
  }
}
