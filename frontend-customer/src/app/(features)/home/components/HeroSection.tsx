import Link from 'next/link';

import AnimateInView from '@/components/ui/AnimateInView';

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
      <div className="hero-gradient relative mb-12 overflow-hidden py-20 text-white md:py-32">
        <div className="absolute top-0 left-0 h-full w-full bg-black opacity-40"></div>
        <div className="relative z-10 container mx-auto px-6">
          <div className="mx-auto max-w-3xl text-center">
            <AnimateInView direction="down" delay={200}>
              <h1 className="mb-6 text-4xl leading-tight font-bold md:text-5xl">
                <span className="mb-2 block">高品質な商品を</span>
                <span className="bg-gradient-to-r from-white to-indigo-200 bg-clip-text text-transparent">
                  お求めやすい価格で
                </span>
              </h1>
            </AnimateInView>

            <AnimateInView direction="up" delay={400}>
              <p className="mb-10 text-xl text-gray-100 md:text-2xl">
                オブザーバビリティショップで最高のショッピング体験をお楽しみください
              </p>

              <div className="flex flex-col justify-center gap-4 sm:flex-row">
                <Link
                  href={primaryButtonLink}
                  className="btn-primary rounded-full px-8 py-3 text-lg font-medium text-white shadow-lg transition-all hover:shadow-xl"
                >
                  {primaryButtonText}
                </Link>
                <Link
                  href={secondaryButtonLink}
                  className="rounded-full bg-white/20 px-8 py-3 text-lg font-medium text-white backdrop-blur-sm transition-colors hover:bg-white/30"
                >
                  {secondaryButtonText}
                </Link>
              </div>
            </AnimateInView>
          </div>
        </div>

        {/* 装飾要素 */}
        <div className="absolute -bottom-16 -left-16 h-64 w-64 rounded-full bg-indigo-600/20 blur-3xl"></div>
        <div className="absolute -top-20 -right-20 h-72 w-72 rounded-full bg-purple-600/20 blur-3xl"></div>
      </div>
    </AnimateInView>
  );
}
