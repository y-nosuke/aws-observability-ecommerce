"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

export default function AdminSidebar() {
  const pathname = usePathname();

  const menuItems = [
    { name: "ダッシュボード", path: "/dashboard" },
    { name: "商品管理", path: "/products" },
    { name: "在庫管理", path: "/inventory" },
  ];

  return (
    <aside className="w-64 bg-gray-800 text-white p-4">
      <nav>
        <ul className="space-y-2">
          {menuItems.map((item) => (
            <li key={item.path}>
              <Link
                href={`/admin${item.path}`}
                className={`block p-2 rounded ${
                  pathname === `/admin${item.path}`
                    ? "bg-blue-700"
                    : "hover:bg-gray-700"
                }`}
              >
                {item.name}
              </Link>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  );
}
