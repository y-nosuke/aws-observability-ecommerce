'use client';

import { useRouter } from 'next/navigation';
import { useState } from 'react';

// 商品検索フォームのクライアントコンポーネント
export default function SearchForm() {
  const router = useRouter();
  const [keyword, setKeyword] = useState('');

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    // ここに検索処理を実装
    console.log('検索キーワード:', keyword);

    // 検索パラメータ付きでページを更新
    if (keyword.trim()) {
      router.push(`/products?search=${encodeURIComponent(keyword.trim())}`);
    } else {
      router.push('/products');
    }
  };

  return (
    <div className="bg-white p-4 rounded-lg shadow mb-6">
      <h3 className="text-lg font-medium text-gray-700 mb-3">商品検索</h3>
      <form onSubmit={handleSearch} className="flex flex-wrap gap-4">
        <div className="flex-1 min-w-[200px]">
          <label htmlFor="search" className="block text-sm font-medium text-gray-700 mb-1">
            キーワード
          </label>
          <input
            type="text"
            id="search"
            value={keyword}
            onChange={e => setKeyword(e.target.value)}
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            placeholder="商品名や説明で検索"
          />
        </div>
        <div className="flex-none self-end">
          <button
            type="submit"
            className="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-300"
          >
            検索
          </button>
        </div>
      </form>
    </div>
  );
}
