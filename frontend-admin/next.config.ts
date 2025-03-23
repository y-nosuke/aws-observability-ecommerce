import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  /* config options here */
  // 画像最適化の設定（最新の推奨方法）
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'placehold.co',
      },
      {
        protocol: 'https',
        hostname: 'non-existent-domain.example',
      },
      {
        protocol: 'http',
        hostname: 'localhost',
      },
      {
        protocol: 'http',
        hostname: 'backend',
      },
    ],
    dangerouslyAllowSVG: true,
    contentDispositionType: 'attachment',
    contentSecurityPolicy: "default-src 'self'; script-src 'none'; sandbox;",
  },
  // 開発環境でのクロスオリジンリクエストを許可
  allowedDevOrigins: ['shop.localhost', 'admin.localhost', 'api.localhost'],
  // 外部API接続のための設定
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://backend:8080/api/:path*', // コンテナ名を使用
      },
    ];
  },
};

export default nextConfig;
