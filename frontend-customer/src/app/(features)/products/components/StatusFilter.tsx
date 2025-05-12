interface StatusFilterProps {
  onStatusFilterChange: (statuses: string[]) => void;
  selectedStatuses: string[];
}

export default function StatusFilter({
  onStatusFilterChange,
  selectedStatuses,
}: StatusFilterProps) {
  // 商品状態の定義
  const statuses = [
    { id: "status-new", label: "新着商品", value: "new" },
    { id: "status-sale", label: "セール中", value: "sale" },
  ];
  
  // チェックボックスの状態変更ハンドラー
  const handleStatusFilterChange = (value: string) => {
    if (selectedStatuses.includes(value)) {
      // すでに選択されている場合は削除
      onStatusFilterChange(selectedStatuses.filter((status) => status !== value));
    } else {
      // 選択されていない場合は追加
      onStatusFilterChange([...selectedStatuses, value]);
    }
  };

  return (
    <div className="mb-8">
      <h2 className="text-xl font-bold mb-4">状態</h2>
      <div className="space-y-2">
        {statuses.map((status) => (
          <div key={status.id} className="flex items-center">
            <input
              type="checkbox"
              id={status.id}
              className="mr-3"
              checked={selectedStatuses.includes(status.value)}
              onChange={() => handleStatusFilterChange(status.value)}
              aria-label={`${status.label}を表示`}
            />
            <label htmlFor={status.id}>{status.label}</label>
          </div>
        ))}
      </div>
    </div>
  );
}
