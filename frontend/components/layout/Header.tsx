"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Header() {
  const pathname = usePathname();

  return (
    <header className="bg-blue-700 text-white shadow-md">
      <div className="container mx-auto px-4 py-4">
        <div className="flex justify-between items-center">
          <div>
            <Link href="/" className="text-xl font-bold">
              ECアプリ
            </Link>
          </div>

          <nav>
            <ul className="flex space-x-6">
              <li>
                <Link href="/" className={pathname === "/" ? "font-bold" : ""}>
                  ホーム
                </Link>
              </li>
              <li>
                <Link
                  href="/products"
                  className={
                    pathname.startsWith("/products") ? "font-bold" : ""
                  }
                >
                  商品一覧
                </Link>
              </li>
              <li>
                <Link
                  href="/cart"
                  className={pathname === "/cart" ? "font-bold" : ""}
                >
                  カート
                </Link>
              </li>
            </ul>
          </nav>
        </div>
      </div>
    </header>
  );
}
