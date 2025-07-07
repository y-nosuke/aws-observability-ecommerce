type StatItemProps = {
  label: string;
  value: string | number;
  change?: string;
  isPositive?: boolean;
};

function StatItem({ label, value, change, isPositive }: StatItemProps) {
  const gradientColorClass = isPositive
    ? 'from-emerald-500 to-teal-600'
    : 'from-red-500 to-orange-600';

  return (
    <div className="stat-card relative overflow-hidden rounded-lg bg-white p-5 shadow-lg dark:bg-gray-800">
      {/* 見えないグラデーションアクセント（右上角） */}
      <div
        className={`absolute top-0 right-0 h-20 w-20 bg-gradient-to-br opacity-10 ${gradientColorClass}`}
      ></div>
      {/* アイコンインジケーター */}
      <div className="flex items-start justify-between">
        <p className="text-sm font-medium text-gray-500 dark:text-gray-400">{label}</p>
        {change && (
          <div
            className={`flex h-8 w-8 items-center justify-center rounded-full ${
              isPositive
                ? 'bg-green-100 text-green-600 dark:bg-green-900/30 dark:text-green-400'
                : 'bg-red-100 text-red-600 dark:bg-red-900/30 dark:text-red-400'
            }`}
          >
            {isPositive ? (
              <svg className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M5 10l7-7m0 0l7 7m-7-7v18"
                />
              </svg>
            ) : (
              <svg className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M19 14l-7 7m0 0l-7-7m7 7V3"
                />
              </svg>
            )}
          </div>
        )}
      </div>

      <p className="mt-2 text-3xl font-bold text-gray-900 dark:text-gray-100">{value}</p>

      {change && (
        <div className="mt-4 flex items-center">
          <div
            className={`h-1 w-12 rounded-full ${
              isPositive
                ? 'bg-gradient-to-r from-green-400 to-emerald-500'
                : 'bg-gradient-to-r from-red-500 to-orange-400'
            }`}
          ></div>
          <p
            className={`ml-2 text-sm font-medium ${
              isPositive ? 'text-green-600 dark:text-green-400' : 'text-red-600 dark:text-red-400'
            }`}
          >
            {change}
          </p>
        </div>
      )}
    </div>
  );
}

type DashboardStatsProps = {
  stats: StatItemProps[];
};

export default function DashboardStats({ stats }: DashboardStatsProps) {
  return (
    <div className="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-4">
      {stats.map((stat, index) => (
        <StatItem
          key={index}
          label={stat.label}
          value={stat.value}
          change={stat.change}
          isPositive={stat.isPositive}
        />
      ))}
    </div>
  );
}
