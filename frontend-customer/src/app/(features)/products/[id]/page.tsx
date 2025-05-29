import { fetchProductById } from "@/services/products/api";
import { notFound } from "next/navigation";
import ProductDetailClient from "./client";

// 動的レンダリングを強制（ビルド時の静的生成を無効化）
export const dynamic = "force-dynamic";

// 商品IDの型をより厳密に定義
type ProductId = number & { readonly brand: unique symbol };

interface ProductDetailPageProps {
  params: Promise<{
    id: string;
  }>;
}

// 商品IDのバリデーションと型変換
function validateAndConvertProductId(id: string): ProductId | null {
  const parsedId = parseInt(id, 10);
  if (isNaN(parsedId) || parsedId <= 0) {
    return null;
  }
  return parsedId as ProductId;
}

// サーバーコンポーネント：商品詳細データを取得し、クライアントコンポーネントに渡す
export default async function ProductDetailPage({
  params,
}: ProductDetailPageProps) {
  const resolvedParams = await params;
  const productId = validateAndConvertProductId(resolvedParams.id);

  if (!productId) {
    notFound();
  }

  try {
    // サーバーサイドで商品詳細データを取得
    const product = await fetchProductById(productId);

    if (!product) {
      notFound();
    }

    // クライアントコンポーネントにデータを渡す
    return <ProductDetailClient product={product} />;
  } catch (error) {
    console.error("Failed to fetch product:", error);
    // エラーが発生した場合はNext.jsのエラーバウンダリに委ねる
    throw error;
  }
}

// メタデータの生成（SEO対応）
export async function generateMetadata({ params }: ProductDetailPageProps) {
  const resolvedParams = await params;
  const productId = validateAndConvertProductId(resolvedParams.id);

  if (!productId) {
    return {
      title: "商品が見つかりません",
    };
  }

  try {
    const product = await fetchProductById(productId);

    if (!product) {
      return {
        title: "商品が見つかりません",
      };
    }

    return {
      title: `${product.name} | オンラインストア`,
      description: product.description || `${product.name}の詳細ページです。`,
      openGraph: {
        title: product.name,
        description: product.description || `${product.name}の詳細ページです。`,
        images: product.imageUrl ? [product.imageUrl] : [],
      },
    };
  } catch (error) {
    console.error("Failed to generate metadata:", error);
    return {
      title: "商品詳細",
    };
  }
}
