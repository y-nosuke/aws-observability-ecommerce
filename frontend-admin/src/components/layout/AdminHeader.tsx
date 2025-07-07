'use client';

import Link from 'next/link';
import { useState } from 'react';

import { getUserData } from '@/lib/auth/auth';

type AdminHeaderProps = {
  toggleSidebar?: () => void;
};

export default function AdminHeader({ toggleSidebar }: AdminHeaderProps) {
  const [isProfileMenuOpen, setIsProfileMenuOpen] = useState(false);
  const userData = getUserData() || { name: 'Admin User' };

  const toggleProfileMenu = () => {
    setIsProfileMenuOpen(!isProfileMenuOpen);
  };

  return (
    <header className="bg-gradient-primary text-white shadow-lg">
      <div className="flex items-center justify-between px-4 py-4">
        <div className="flex items-center">
          <button
            onClick={toggleSidebar}
            className="rounded-md p-2 transition-colors hover:bg-white/10"
            aria-label="サイドバー切り替え"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M4 6h16M4 12h16M4 18h16"
              />
            </svg>
          </button>
          <Link href="/" className="flex items-center text-xl font-bold">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="mr-2 h-6 w-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M2.25 13.5h3.86a2.25 2.25 0 012.012 1.244l.256.512a2.25 2.25 0 002.013 1.244h3.218a2.25 2.25 0 002.013-1.244l.256-.512a2.25 2.25 0 012.013-1.244h3.859m-19.5.338V18a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18v-4.162c0-.224-.034-.447-.1-.661L19.24 5.338a2.25 2.25 0 00-2.15-1.588H6.911a2.25 2.25 0 00-2.15 1.588L2.35 13.177a2.25 2.25 0 00-.1.661z"
              />
            </svg>
            管理ダッシュボード
          </Link>
        </div>

        <div className="relative">
          <button
            onClick={toggleProfileMenu}
            className="transition-smooth flex items-center space-x-2 rounded-full px-2 py-1 hover:opacity-80 focus:outline-none"
          >
            <div className="flex h-9 w-9 items-center justify-center rounded-full bg-white/20 font-semibold shadow-inner backdrop-blur-sm transition-all">
              {userData.name.charAt(0)}
            </div>
            <span className="hidden font-medium md:inline">{userData.name}</span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={2}
              stroke="currentColor"
              className="h-4 w-4 transition-transform duration-200 ease-in-out"
              style={{
                transform: isProfileMenuOpen ? 'rotate(180deg)' : 'rotate(0deg)',
              }}
            >
              <path strokeLinecap="round" strokeLinejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
            </svg>
          </button>

          {isProfileMenuOpen && (
            <div className="absolute right-0 z-10 mt-2 w-52 overflow-hidden rounded-xl border border-gray-100 bg-white py-2 text-gray-800 shadow-xl dark:border-gray-700 dark:bg-gray-800 dark:text-gray-200">
              <Link
                href="/profile"
                className="group flex items-center px-4 py-2.5 transition-colors hover:bg-gray-50 dark:hover:bg-gray-700"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="group-hover:text-primary mr-3 h-5 w-5 text-gray-500"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                  />
                </svg>
                <span className="font-medium">プロフィール</span>
              </Link>
              <Link
                href="/settings"
                className="group flex items-center px-4 py-2.5 transition-colors hover:bg-gray-50 dark:hover:bg-gray-700"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="group-hover:text-primary mr-3 h-5 w-5 text-gray-500"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
                  />
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                </svg>
                <span className="font-medium">設定</span>
              </Link>
              <div className="my-1 border-t border-gray-200 dark:border-gray-700"></div>
              <button
                onClick={() => {
                  // ログアウト処理
                  if (typeof window !== 'undefined') {
                    window.location.href = '/login';
                  }
                }}
                className="group flex w-full items-center px-4 py-2.5 text-left transition-colors hover:bg-gray-50 dark:hover:bg-gray-700"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="mr-3 h-5 w-5 text-red-500"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                  />
                </svg>
                <span className="font-medium text-red-600 dark:text-red-500">ログアウト</span>
              </button>
            </div>
          )}
        </div>
      </div>
    </header>
  );
}
