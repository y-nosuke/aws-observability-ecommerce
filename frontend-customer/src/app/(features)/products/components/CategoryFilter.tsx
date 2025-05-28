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
      <h2 className="text-xl font-bold mb-4">カテゴリー</h2>
      <ul className="space-y-2">
        <li>
          <button
            onClick={() => onCategoryChange(0)}
            className={`w-full text-left py-2 px-3 rounded-lg transition-colors ${
              activeCategory === 0
                ? "bg-primary text-white font-medium"
                : "hover:bg-gray-100 dark:hover:bg-gray-800"
            }`}
            aria-current={activeCategory === 0 ? "true" : "false"}
          >
            すべて
          </button>
        </li>
        {categories.map((category) => (
          <li key={category.id}>
            <button
              onClick={() => onCategoryChange(category.id)}
              className={`w-full text-left py-2 px-3 rounded-lg transition-colors ${
                activeCategory === category.id
                  ? "bg-primary text-white font-medium"
                  : "hover:bg-gray-100 dark:hover:bg-gray-800"
              }`}
              aria-current={activeCategory === category.id ? "true" : "false"}
            >
              {category.name}
              {category.productCount !== undefined && (
                <span className="ml-2 text-sm text-gray-500 dark:text-gray-400">
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
