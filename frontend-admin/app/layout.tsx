import { Inter } from 'next/font/google';
import Link from 'next/link';
import './globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata = {
  title: '管理者ダッシュボード | AWS オブザーバビリティ学習用 eコマース',
  description: 'AWSのオブザーバビリティ機能を学習するためのeコマースサイト管理者機能',
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
                  <span className="text-emerald-600 font-bold text-xl">管理者ダッシュボード</span>
                </Link>
                <nav className="ml-6 flex items-center space-x-4">
                  <Link
                    href="/products"
                    className="text-gray-700 hover:text-emerald-600 px-3 py-2 rounded-md text-sm font-medium"
                  >
                    商品管理
                  </Link>
                  <span className="text-gray-400 px-3 py-2 rounded-md text-sm font-medium">
                    在庫管理
                  </span>
                  <span className="text-gray-400 px-3 py-2 rounded-md text-sm font-medium">
                    注文管理
                  </span>
                </nav>
              </div>
              <div className="flex items-center">
                <span className="text-gray-700 px-3 py-2 rounded-md text-sm font-medium">
                  管理者
                </span>
              </div>
            </div>
          </div>
        </header>

        <main>{children}</main>

        <footer className="bg-gray-800 text-white mt-auto">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
            <div className="flex flex-col md:flex-row justify-between items-center">
              <div className="mb-2 md:mb-0">
                <p>&copy; 2025 AWS オブザーバビリティ学習用 eコマースサイト 管理システム</p>
              </div>
            </div>
          </div>
        </footer>
      </body>
    </html>
  );
}
