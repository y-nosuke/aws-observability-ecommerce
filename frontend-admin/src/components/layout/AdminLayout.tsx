"use client";

import { ReactNode, useState } from "react";
import AdminHeader from "./AdminHeader";
import AdminSidebar from "./AdminSidebar";

type AdminLayoutProps = {
  children: ReactNode;
};

export default function AdminLayout({ children }: AdminLayoutProps) {
  const [sidebarOpen, setSidebarOpen] = useState(true);

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  return (
    <div className="flex flex-col min-h-screen bg-gray-50 dark:bg-gray-900">
      <AdminHeader toggleSidebar={toggleSidebar} />
      <div className="flex flex-1">
        <AdminSidebar isOpen={sidebarOpen} />
        <main
          className={`flex-1 p-6 md:p-8 overflow-x-auto transition-all duration-300 ${
            !sidebarOpen ? "pl-4" : "pl-6"
          }`}
        >
          {children}
        </main>
      </div>
    </div>
  );
}
