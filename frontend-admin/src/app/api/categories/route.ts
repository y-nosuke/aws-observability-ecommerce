import { NextRequest, NextResponse } from 'next/server';

const BACKEND_API_URL = process.env.BACKEND_API_URL || 'http://backend-api:8000/api';

export async function GET(request: NextRequest) {
  try {
    console.log('Proxying to backend:', `${BACKEND_API_URL}/categories`);

    // バックエンドAPIにリクエストを転送
    const response = await fetch(`${BACKEND_API_URL}/categories`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        // 必要に応じて認証ヘッダーを追加
        ...(request.headers.get('authorization') && {
          Authorization: request.headers.get('authorization')!,
        }),
      },
    });

    if (!response.ok) {
      console.error(`Backend API error: ${response.status} ${response.statusText}`);
      const errorText = await response.text();
      console.error('Error details:', errorText);
      return NextResponse.json(
        {
          error: 'Failed to fetch categories from backend',
          details: errorText,
        },
        { status: response.status },
      );
    }

    const data = await response.json();
    return NextResponse.json(data);
  } catch (error) {
    console.error('API Route error:', error);
    return NextResponse.json(
      {
        error: 'Internal server error',
        details: error instanceof Error ? error.message : 'Unknown error',
      },
      { status: 500 },
    );
  }
}
