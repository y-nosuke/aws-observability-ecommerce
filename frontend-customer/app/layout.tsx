import { Inter } from 'next/font/google';
import Link from 'next/link';
import './globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata = {
  title: 'AWS オブザーバビリティ学習用 eコマース',
  description: 'AWSのオブザーバビリティ機能を学習するためのeコマースサイト',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body className={inter.className}>
        <header className="bg-white shadow-md">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="flex justify-between h-16">
              <div className="flex">
                <Link href="/" className="flex-shrink-0 flex items-center">
                  <span className="text-indigo-600 font-bold text-xl">eコマースサイト</span>
                </Link>
                <nav className="ml-6 flex items-center space-x-4">
                  <Link
                    href="/products"
                    className="text-gray-700 hover:text-indigo-600 px-3 py-2 rounded-md text-sm font-medium"
                  >
                    商品一覧
                  </Link>
                  <Link
                    href="/products?category=1"
                    className="text-gray-700 hover:text-indigo-600 px-3 py-2 rounded-md text-sm font-medium"
                  >
                    カテゴリー
                  </Link>
                </nav>
              </div>
            </div>
          </div>
        </header>

        <main>{children}</main>

        <footer className="bg-gray-800 text-white mt-auto">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
            <div className="flex flex-col md:flex-row justify-between items-center">
              <div className="mb-4 md:mb-0">
                <p>&copy; 2025 AWS オブザーバビリティ学習用 eコマースサイト</p>
              </div>
              <div className="flex space-x-6">
                <a href="#" className="text-gray-300 hover:text-white">
                  利用規約
                </a>
                <a href="#" className="text-gray-300 hover:text-white">
                  プライバシーポリシー
                </a>
                <a href="#" className="text-gray-300 hover:text-white">
                  お問い合わせ
                </a>
              </div>
            </div>
          </div>
        </footer>
      </body>
    </html>
  );
}
