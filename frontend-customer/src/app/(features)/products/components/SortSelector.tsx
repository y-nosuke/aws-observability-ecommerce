import { SortOption } from '@/services/products/types';

interface SortSelectorProps {
  sortOption: SortOption;
  onSortChange: (option: SortOption) => void;
}

export default function SortSelector({ sortOption, onSortChange }: SortSelectorProps) {
  return (
    <div className="flex items-center">
      <label htmlFor="sort" className="mr-2 text-sm">
        並び替え:
      </label>
      <select
        id="sort"
        value={sortOption}
        onChange={(e) => onSortChange(e.target.value as SortOption)}
        className="focus:ring-primary rounded-lg border border-gray-300 bg-white px-3 py-2 text-sm focus:ring-2 focus:outline-none dark:border-gray-700 dark:bg-gray-800"
        aria-label="商品の並び替え選択"
      >
        <option value="recommended">おすすめ順</option>
        <option value="newest">新着順</option>
        <option value="price-asc">価格: 安い順</option>
        <option value="price-desc">価格: 高い順</option>
      </select>
    </div>
  );
}
