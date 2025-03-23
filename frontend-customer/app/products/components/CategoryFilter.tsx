import { getCategories } from '@/lib/api/products';
import Link from 'next/link';

export default async function CategoryFilter({
  selectedCategoryId,
}: {
  selectedCategoryId?: number;
}) {
  const categories = await getCategories();

  return (
    <div className="mb-6">
      <h3 className="text-lg font-medium text-gray-700 mb-2">カテゴリー</h3>
      <div className="flex flex-wrap gap-2">
        <Link
          href="/products"
          className={`px-3 py-1 rounded-full text-sm font-medium ${
            !selectedCategoryId
              ? 'bg-indigo-600 text-white'
              : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
          }`}
        >
          すべて
        </Link>

        {categories.map(category => (
          <Link
            key={category.id}
            href={`/products?category=${category.id}`}
            className={`px-3 py-1 rounded-full text-sm font-medium ${
              selectedCategoryId === category.id
                ? 'bg-indigo-600 text-white'
                : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
            }`}
          >
            {category.name}
          </Link>
        ))}
      </div>
    </div>
  );
}
