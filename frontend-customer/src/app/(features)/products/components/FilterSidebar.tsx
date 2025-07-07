import CategoryFilter from './CategoryFilter';
import PriceFilter from './PriceFilter';
import StatusFilter from './StatusFilter';

interface FilterSidebarProps {
  categories: Array<{
    id: number;
    name: string;
  }>;
  activeCategory: number;
  onCategoryChange: (categoryId: number) => void;
  selectedPriceRanges: string[];
  onPriceFilterChange: (priceRanges: string[]) => void;
  selectedStatuses: string[];
  onStatusFilterChange: (statuses: string[]) => void;
}

export default function FilterSidebar({
  categories,
  activeCategory,
  onCategoryChange,
  selectedPriceRanges,
  onPriceFilterChange,
  selectedStatuses,
  onStatusFilterChange,
}: FilterSidebarProps) {
  return (
    <div className="w-full shrink-0 md:w-64">
      <div className="sticky top-28">
        <CategoryFilter
          categories={categories}
          activeCategory={activeCategory}
          onCategoryChange={onCategoryChange}
        />

        <PriceFilter
          selectedPriceRanges={selectedPriceRanges}
          onPriceFilterChange={onPriceFilterChange}
        />

        <StatusFilter
          selectedStatuses={selectedStatuses}
          onStatusFilterChange={onStatusFilterChange}
        />
      </div>
    </div>
  );
}
