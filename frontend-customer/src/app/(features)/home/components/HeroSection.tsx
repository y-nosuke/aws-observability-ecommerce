import Link from "next/link";
import AnimateInView from "@/components/ui/AnimateInView";
import { Product } from "@/services/products/types";

interface HeroSectionProps {
  primaryButtonLink: string;
  primaryButtonText: string;
  secondaryButtonLink: string;
  secondaryButtonText: string;
}

export default function HeroSection({
  primaryButtonLink,
  primaryButtonText,
  secondaryButtonLink,
  secondaryButtonText,
}: HeroSectionProps) {
  return (
    <AnimateInView>
      <div className="hero-gradient text-white py-20 md:py-32 mb-12 relative overflow-hidden">
        <div className="absolute top-0 left-0 w-full h-full bg-black opacity-40"></div>
        <div className="container mx-auto px-6 relative z-10">
          <div className="max-w-3xl mx-auto text-center">
            <AnimateInView direction="down" delay={200}>
              <h1 className="text-4xl md:text-5xl font-bold mb-6 leading-tight">
                <span className="block mb-2">高品質な商品を</span>
                <span className="bg-clip-text text-transparent bg-gradient-to-r from-white to-indigo-200">
                  お求めやすい価格で
                </span>
              </h1>
            </AnimateInView>

            <AnimateInView direction="up" delay={400}>
              <p className="text-xl md:text-2xl mb-10 text-gray-100">
                オブザーバビリティショップで最高のショッピング体験をお楽しみください
              </p>

              <div className="flex flex-col sm:flex-row justify-center gap-4">
                <Link
                  href={primaryButtonLink}
                  className="btn-primary text-white py-3 px-8 rounded-full font-medium text-lg shadow-lg transition-all hover:shadow-xl"
                >
                  {primaryButtonText}
                </Link>
                <Link
                  href={secondaryButtonLink}
                  className="bg-white/20 backdrop-blur-sm text-white py-3 px-8 rounded-full font-medium text-lg hover:bg-white/30 transition-colors"
                >
                  {secondaryButtonText}
                </Link>
              </div>
            </AnimateInView>
          </div>
        </div>

        {/* 装飾要素 */}
        <div className="absolute -bottom-16 -left-16 w-64 h-64 bg-indigo-600/20 rounded-full blur-3xl"></div>
        <div className="absolute -top-20 -right-20 w-72 h-72 bg-purple-600/20 rounded-full blur-3xl"></div>
      </div>
    </AnimateInView>
  );
}
