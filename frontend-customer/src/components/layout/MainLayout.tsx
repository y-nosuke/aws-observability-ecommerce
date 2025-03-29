import { ReactNode } from "react";
import Footer from "./Footer";
import Header from "./Header";

type MainLayoutProps = {
  children: ReactNode;
};

export default function MainLayout({ children }: MainLayoutProps) {
  return (
    <div className="flex flex-col min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
      <Header />
      <main className="flex-grow pt-24 md:pt-28 pb-12">{children}</main>
      <Footer />
    </div>
  );
}
