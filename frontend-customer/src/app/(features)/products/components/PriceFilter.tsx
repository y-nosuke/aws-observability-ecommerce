interface PriceFilterProps {
  onPriceFilterChange: (priceRanges: string[]) => void;
  selectedPriceRanges: string[];
}

export default function PriceFilter({
  onPriceFilterChange,
  selectedPriceRanges,
}: PriceFilterProps) {
  // 価格帯の定義
  const priceRanges = [
    { id: 'price-1', label: '¥10,000以下', value: 'under-10000' },
    { id: 'price-2', label: '¥10,000〜¥30,000', value: '10000-30000' },
    { id: 'price-3', label: '¥30,000〜¥50,000', value: '30000-50000' },
    { id: 'price-4', label: '¥50,000以上', value: 'over-50000' },
  ];

  // チェックボックスの状態変更ハンドラー
  const handlePriceFilterChange = (value: string) => {
    if (selectedPriceRanges.includes(value)) {
      // すでに選択されている場合は削除
      onPriceFilterChange(selectedPriceRanges.filter((range) => range !== value));
    } else {
      // 選択されていない場合は追加
      onPriceFilterChange([...selectedPriceRanges, value]);
    }
  };

  return (
    <div className="mb-8">
      <h2 className="mb-4 text-xl font-bold">価格</h2>
      <div className="space-y-2">
        {priceRanges.map((range) => (
          <div key={range.id} className="flex items-center">
            <input
              type="checkbox"
              id={range.id}
              className="mr-3"
              checked={selectedPriceRanges.includes(range.value)}
              onChange={() => handlePriceFilterChange(range.value)}
              aria-label={`${range.label}の商品を表示`}
            />
            <label htmlFor={range.id}>{range.label}</label>
          </div>
        ))}
      </div>
    </div>
  );
}
