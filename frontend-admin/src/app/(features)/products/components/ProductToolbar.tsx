'use client';

import { usePathname, useRouter, useSearchParams } from 'next/navigation';
import { useCallback, useEffect, useState } from 'react';

import { Category } from '@/services/products/types';

type ProductToolbarProps = {
  categories: Category[];
};

export default function ProductToolbar({ categories }: ProductToolbarProps) {
  const searchParams = useSearchParams();
  const router = useRouter();
  const pathname = usePathname();
  const [searchInput, setSearchInput] = useState('');
  const [isComposing, setIsComposing] = useState(false);

  // URLから現在の検索条件を取得
  const currentKeyword = searchParams.get('keyword') || '';
  const currentCategoryId = searchParams.get('categoryId') || '';
  // URLパラメータは文字列なので、数値への変換が必要
  const currentCategoryIdNumber = currentCategoryId ? parseInt(currentCategoryId) : null;

  // URLの検索パラメータを更新する関数
  const updateSearchParams = useCallback(
    (updates: Record<string, string | null>, shouldCreateHistory: boolean = false) => {
      const params = new URLSearchParams(searchParams);

      Object.entries(updates).forEach(([key, value]) => {
        if (value === null || value === '') {
          params.delete(key);
        } else {
          params.set(key, value);
        }
      });

      // ページをリセット（検索条件変更時）
      if (updates.keyword !== undefined || updates.categoryId !== undefined) {
        params.delete('page');
      }

      // 完全なURLパスを構築（履歴が正しく記録されるように）
      const queryString = params.toString();
      const newUrl = queryString ? `${pathname}?${queryString}` : pathname;
      const currentQueryString = searchParams.toString();
      const currentUrl = currentQueryString ? `${pathname}?${currentQueryString}` : pathname;

      // 同じURLへの更新を避ける
      if (newUrl !== currentUrl) {
        if (shouldCreateHistory) {
          // 意図的な操作は履歴に残す（カテゴリ変更、手動検索、フィルタークリアなど）
          router.push(newUrl);
        } else {
          // 自動的な操作は履歴を汚染しない（デバウンス処理など）
          router.replace(newUrl);
        }
      }
    },
    [searchParams, router, pathname],
  );

  // 初期化時にURLの検索キーワードを検索入力フィールドに設定
  useEffect(() => {
    setSearchInput(currentKeyword);
  }, [currentKeyword]);

  // インクリメンタルサーチ：IME入力中は実行せず、入力値が変化したときに即座に検索を実行
  useEffect(() => {
    // IME入力中は検索を実行しない
    if (isComposing) {
      return;
    }

    // 入力値が現在のURLキーワードと異なる場合のみ更新
    if (searchInput !== currentKeyword) {
      // インクリメンタルサーチは履歴に残さない（自動的な処理のため）
      updateSearchParams({ keyword: searchInput }, false);
    }
  }, [searchInput, currentKeyword, updateSearchParams, isComposing]);

  const handleSearchInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchInput(e.target.value);
  };

  const handleSearchInputKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      // Enterキー押下は意図的な検索なので履歴に残す
      updateSearchParams({ keyword: searchInput }, true);
    }
  };

  const handleClearSearch = () => {
    setSearchInput('');
    // 検索クリアは意図的な操作なので履歴に残す
    updateSearchParams({ keyword: null }, true);
  };

  const handleCategoryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const value = e.target.value;
    // カテゴリ変更は意図的な操作なので履歴に残す
    updateSearchParams({ categoryId: value || null }, true);
  };

  const handleClearFilters = () => {
    setSearchInput('');
    // フィルタークリアは意図的な操作なので履歴に残す
    updateSearchParams({ keyword: null, categoryId: null }, true);
  };

  const hasActiveFilters = currentKeyword || currentCategoryId;

  // IME入力開始
  const handleCompositionStart = () => {
    setIsComposing(true);
  };

  // IME入力確定
  const handleCompositionEnd = () => {
    setIsComposing(false);
  };

  return (
    <div className="space-y-4 rounded-lg border border-gray-200 bg-white p-4 shadow-sm dark:border-gray-700 dark:bg-gray-800">
      <div className="flex flex-col gap-4 lg:flex-row">
        {/* 検索バー */}
        <div className="flex-1">
          <div className="relative">
            <input
              type="text"
              placeholder="商品名で検索..."
              value={searchInput}
              onChange={handleSearchInputChange}
              onKeyDown={handleSearchInputKeyDown}
              onCompositionStart={handleCompositionStart}
              onCompositionEnd={handleCompositionEnd}
              className="w-full rounded-md border border-gray-300 py-2 pr-10 pl-10 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500"
            />
            {/* 検索アイコン */}
            <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
              <svg
                className="h-5 w-5 text-gray-400"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z"
                />
              </svg>
            </div>
            {/* クリアボタン */}
            {searchInput && (
              <button
                type="button"
                onClick={handleClearSearch}
                className="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-400 hover:text-gray-600"
              >
                <svg
                  className="h-5 w-5"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth={1.5}
                  stroke="currentColor"
                >
                  <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            )}
          </div>
        </div>

        {/* カテゴリフィルター */}
        <div className="lg:w-64">
          <select
            value={currentCategoryId}
            onChange={handleCategoryChange}
            className="w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-gray-900 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500"
          >
            <option value="" className="bg-white text-gray-900">
              全カテゴリ
            </option>
            {categories.map((category) => (
              <option
                key={category.id}
                value={category.id.toString()}
                className="bg-white text-gray-900"
              >
                {category.name}
              </option>
            ))}
          </select>
        </div>

        {/* フィルタークリアボタン */}
        {hasActiveFilters && (
          <button
            type="button"
            onClick={handleClearFilters}
            className="w-full rounded-md border border-gray-300 bg-gray-100 px-4 py-2 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500 lg:w-auto"
          >
            <span className="flex items-center justify-center gap-2">
              <svg
                className="h-4 w-4"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
              >
                <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
              フィルタークリア
            </span>
          </button>
        )}
      </div>

      {/* アクティブフィルター表示 */}
      {hasActiveFilters && (
        <div className="flex flex-wrap gap-2">
          {currentKeyword && (
            <span className="inline-flex items-center rounded-full bg-indigo-100 px-2.5 py-0.5 text-xs font-medium text-indigo-800">
              検索: &quot;{currentKeyword}&quot;
            </span>
          )}
          {currentCategoryIdNumber && (
            <span className="inline-flex items-center rounded-full bg-green-100 px-2.5 py-0.5 text-xs font-medium text-green-800">
              カテゴリ: {categories.find((c) => c.id === currentCategoryIdNumber)?.name}
            </span>
          )}
        </div>
      )}
    </div>
  );
}
