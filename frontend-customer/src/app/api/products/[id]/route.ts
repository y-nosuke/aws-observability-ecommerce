import axios from "axios";
import { NextRequest, NextResponse } from "next/server";

const BACKEND_API_URL =
  process.env.BACKEND_API_URL || "http://backend-api:8000/api";

interface RouteParams {
  params: {
    id: string;
  };
}

export async function GET(request: NextRequest, { params }: RouteParams) {
  try {
    const { id } = params;

    // IDのバリデーション
    if (!id || isNaN(Number(id))) {
      return NextResponse.json(
        {
          error: "Invalid product ID",
          message: "Product ID must be a valid number",
          code: "INVALID_PRODUCT_ID",
        },
        { status: 400 }
      );
    }

    const response = await axios.get(`${BACKEND_API_URL}/products/${id}`, {
      params,
      timeout: 10000,
    });

    return NextResponse.json(response.data);
  } catch (error) {
    console.error("API Route error:", error);

    if (axios.isAxiosError(error)) {
      const status = error.response?.status || 500;
      const message = error.response?.data?.message || error.message;

      return NextResponse.json(
        {
          error: "Failed to fetch product",
          message,
          code: "FETCH_PRODUCT_ERROR",
        },
        { status }
      );
    }

    return NextResponse.json(
      {
        error: "Failed to fetch product",
        message: "Unknown error",
        code: "FETCH_PRODUCT_ERROR",
      },
      { status: 500 }
    );
  }
}
