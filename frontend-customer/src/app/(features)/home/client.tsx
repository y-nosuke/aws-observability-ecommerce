"use client";

import { Product } from "@/services/products/types";
import HeroSection from "./components/HeroSection";
import FeaturedProductsSection from "./components/FeaturedProductsSection";
import CampaignSection from "./components/CampaignSection";
import NewsletterSignup from "./components/NewsletterSignup";

interface HomeClientProps {
  popularProducts: Product[];
}

export default function HomeClient({ popularProducts }: HomeClientProps) {
  // キャンペーン情報
  const campaigns = [
    {
      id: "summer-sale",
      title: "夏のビッグセール",
      description: "期間限定で全商品20%オフ。今すぐチェック！",
      badgeText: "SALE",
      badgeColor: "blue" as const,
      bgGradient: "bg-gradient-to-r from-indigo-500 to-purple-600",
      linkText: "詳細を見る",
      linkUrl: "/sale",
    },
    {
      id: "new-arrivals",
      title: "新商品登場",
      description: "最新モデルが登場。先行予約で特典プレゼント！",
      badgeText: "限定商品",
      badgeColor: "yellow" as const,
      bgGradient: "bg-gradient-to-r from-rose-400 to-red-500",
      linkText: "チェックする",
      linkUrl: "/new-arrivals",
    },
  ];

  return (
    <>
      {/* ヒーローセクション */}
      <HeroSection
        primaryButtonLink="/products"
        primaryButtonText="商品を見る"
        secondaryButtonLink="/about"
        secondaryButtonText="詳細を見る"
      />

      {/* 人気商品セクション */}
      <FeaturedProductsSection
        products={popularProducts}
        title="人気の商品"
        viewAllLink="/products"
        viewAllText="すべて見る"
      />

      {/* キャンペーンセクション */}
      <CampaignSection title="特集・キャンペーン" campaigns={campaigns} />

      {/* メールマガジン登録セクション */}
      <NewsletterSignup
        title="お得な情報をお届けします"
        description="メールマガジンに登録すると、新商品情報や限定セールのお知らせをいち早くお届けします。"
        buttonText="登録する"
        placeholderText="メールアドレス"
      />
    </>
  );
}
