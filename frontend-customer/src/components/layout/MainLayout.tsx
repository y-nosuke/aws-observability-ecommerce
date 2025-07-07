import { ReactNode } from 'react';

import Footer from './Footer';
import Header from './Header';

type MainLayoutProps = {
  children: ReactNode;
};

export default function MainLayout({ children }: MainLayoutProps) {
  return (
    <div className="flex min-h-screen flex-col bg-gray-50 text-gray-900 dark:bg-gray-900 dark:text-gray-100">
      <Header />
      <main className="flex-grow pt-24 pb-12 md:pt-28">{children}</main>
      <Footer />
    </div>
  );
}
