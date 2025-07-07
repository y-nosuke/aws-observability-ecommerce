'use client';

import { usePathname, useRouter, useSearchParams } from 'next/navigation';
import { useCallback } from 'react';

type PaginationProps = {
  currentPage: number;
  totalPages: number;
  hasNextPage: boolean;
  hasPreviousPage: boolean;
};

export default function Pagination({
  currentPage,
  totalPages,
  hasNextPage,
  hasPreviousPage,
}: PaginationProps) {
  const searchParams = useSearchParams();
  const router = useRouter();
  const pathname = usePathname();

  // ページ変更時にURLを更新する関数
  const handlePageChange = useCallback(
    (page: number) => {
      const params = new URLSearchParams(searchParams);

      if (page === 1) {
        params.delete('page');
      } else {
        params.set('page', page.toString());
      }

      // 完全なURLパスを構築（履歴が正しく記録されるように）
      const queryString = params.toString();
      const newUrl = queryString ? `${pathname}?${queryString}` : pathname;
      const currentQueryString = searchParams.toString();
      const currentUrl = currentQueryString ? `${pathname}?${currentQueryString}` : pathname;

      // 同じURLへの更新を避ける
      if (newUrl !== currentUrl) {
        // ページネーションは履歴を汚染しないようにreplaceを使用
        router.replace(newUrl);
      }
    },
    [searchParams, router, pathname],
  );

  // ページ番号の配列を生成
  const getPageNumbers = () => {
    const pages: (number | string)[] = [];
    const maxVisiblePages = 7; // 表示する最大ページ数

    if (totalPages <= maxVisiblePages) {
      // 総ページ数が少ない場合は全て表示
      for (let i = 1; i <= totalPages; i++) {
        pages.push(i);
      }
    } else {
      // 総ページ数が多い場合は省略表示
      const startPage = Math.max(1, currentPage - 2);
      const endPage = Math.min(totalPages, currentPage + 2);

      // 最初のページ
      if (startPage > 1) {
        pages.push(1);
        if (startPage > 2) {
          pages.push('...');
        }
      }

      // 中間のページ
      for (let i = startPage; i <= endPage; i++) {
        pages.push(i);
      }

      // 最後のページ
      if (endPage < totalPages) {
        if (endPage < totalPages - 1) {
          pages.push('...');
        }
        pages.push(totalPages);
      }
    }

    return pages;
  };

  const pageNumbers = getPageNumbers();

  return (
    <div className="flex items-center justify-between rounded-lg border-t border-gray-200 bg-white px-4 py-3 shadow-sm sm:px-6 dark:border-gray-700 dark:bg-gray-800">
      <div className="flex flex-1 justify-between sm:hidden">
        {/* モバイル表示 */}
        <button
          onClick={() => handlePageChange(currentPage - 1)}
          disabled={!hasPreviousPage}
          className="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-300 dark:hover:bg-gray-600"
        >
          前へ
        </button>
        <button
          onClick={() => handlePageChange(currentPage + 1)}
          disabled={!hasNextPage}
          className="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-300 dark:hover:bg-gray-600"
        >
          次へ
        </button>
      </div>
      <div className="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
        <div>
          <p className="text-sm text-gray-700 dark:text-gray-300">
            <span className="font-medium">{currentPage}</span> /{' '}
            <span className="font-medium">{totalPages}</span> ページ
          </p>
        </div>
        <div>
          <nav
            className="relative z-0 inline-flex -space-x-px rounded-md shadow-sm"
            aria-label="Pagination"
          >
            {/* 前のページボタン */}
            <button
              onClick={() => handlePageChange(currentPage - 1)}
              disabled={!hasPreviousPage}
              className="relative inline-flex items-center rounded-l-md border border-gray-300 bg-white px-2 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-400 dark:hover:bg-gray-600"
            >
              <span className="sr-only">Previous</span>
              <svg
                className="h-5 w-5"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
              >
                <path
                  fillRule="evenodd"
                  d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                  clipRule="evenodd"
                />
              </svg>
            </button>

            {/* ページ番号 */}
            {pageNumbers.map((page, index) => {
              if (page === '...') {
                return (
                  <span
                    key={`ellipsis-${index}`}
                    className="relative inline-flex items-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-300"
                  >
                    ...
                  </span>
                );
              }

              const pageNum = page as number;
              const isCurrentPage = pageNum === currentPage;

              return (
                <button
                  key={pageNum}
                  onClick={() => handlePageChange(pageNum)}
                  disabled={isCurrentPage}
                  className={`relative inline-flex items-center border px-4 py-2 text-sm font-medium ${
                    isCurrentPage
                      ? 'z-10 border-indigo-500 bg-indigo-50 text-indigo-600 dark:border-indigo-400 dark:bg-indigo-900 dark:text-indigo-300'
                      : 'border-gray-300 bg-white text-gray-500 hover:bg-gray-50 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-400 dark:hover:bg-gray-600'
                  } ${isCurrentPage ? 'cursor-default' : 'cursor-pointer'}`}
                >
                  {pageNum}
                </button>
              );
            })}

            {/* 次のページボタン */}
            <button
              onClick={() => handlePageChange(currentPage + 1)}
              disabled={!hasNextPage}
              className="relative inline-flex items-center rounded-r-md border border-gray-300 bg-white px-2 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-400 dark:hover:bg-gray-600"
            >
              <span className="sr-only">Next</span>
              <svg
                className="h-5 w-5"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
              >
                <path
                  fillRule="evenodd"
                  d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                  clipRule="evenodd"
                />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>
  );
}
