interface CategoryFilterProps {
  categories: Array<{
    id: number;
    name: string;
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
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
