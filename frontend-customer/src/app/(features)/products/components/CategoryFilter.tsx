interface CategoryFilterProps {
  categories: Array<{
    id: number;
    name: string;
    slug?: string;
    productCount?: number;
  }>;
  activeCategory: number;
  onCategoryChange: (categoryId: number) => void;
}

export default function CategoryFilter({
  categories,
  activeCategory,
  onCategoryChange,
}: CategoryFilterProps) {
  return (
    <div className="mb-8">
      <h2 className="mb-4 text-xl font-bold">カテゴリー</h2>
      <ul className="space-y-2" role="listbox" aria-label="商品カテゴリー">
        <li role="option" aria-selected={activeCategory === 0}>
          <button
            onClick={() => onCategoryChange(0)}
            className={`w-full rounded-lg border-2 px-4 py-3 text-left transition-all duration-200 ${
              activeCategory === 0
                ? 'scale-[1.02] border-blue-700 bg-blue-600 font-bold text-white shadow-lg'
                : 'border-transparent hover:bg-gray-100 dark:hover:bg-gray-800'
            }`}
            aria-selected={activeCategory === 0}
            role="option"
          >
            すべて
          </button>
        </li>
        {categories.map((category) => (
          <li key={category.id} role="option" aria-selected={activeCategory === category.id}>
            <button
              onClick={() => onCategoryChange(category.id)}
              className={`w-full rounded-lg border-2 px-4 py-3 text-left transition-all duration-200 ${
                activeCategory === category.id
                  ? 'scale-[1.02] border-blue-700 bg-blue-600 font-bold text-white shadow-lg'
                  : 'border-transparent hover:bg-gray-100 dark:hover:bg-gray-800'
              }`}
              aria-selected={activeCategory === category.id}
              role="option"
            >
              {category.name}
              {category.productCount !== undefined && (
                <span
                  className={`ml-2 text-sm ${
                    activeCategory === category.id
                      ? 'text-blue-100'
                      : 'text-gray-500 dark:text-gray-400'
                  }`}
                >
                  ({category.productCount})
                </span>
              )}
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
