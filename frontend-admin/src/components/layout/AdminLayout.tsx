'use client';

import { ReactNode, useState } from 'react';

import AdminHeader from './AdminHeader';
import AdminSidebar from './AdminSidebar';

type AdminLayoutProps = {
  children: ReactNode;
};

export default function AdminLayout({ children }: AdminLayoutProps) {
  const [sidebarOpen, setSidebarOpen] = useState(true);

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  return (
    <div className="flex min-h-screen flex-col bg-gray-50 dark:bg-gray-900">
      <AdminHeader toggleSidebar={toggleSidebar} />
      <div className="flex flex-1">
        <AdminSidebar isOpen={sidebarOpen} />
        <main
          className={`flex-1 overflow-x-auto p-6 transition-all duration-300 md:p-8 ${
            !sidebarOpen ? 'pl-4' : 'pl-6'
          }`}
        >
          {children}
        </main>
      </div>
    </div>
  );
}
