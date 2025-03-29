type StatItemProps = {
  label: string;
  value: string | number;
  change?: string;
  isPositive?: boolean;
};

function StatItem({ label, value, change, isPositive }: StatItemProps) {
  const gradientColorClass = isPositive
    ? "from-emerald-500 to-teal-600"
    : "from-red-500 to-orange-600";

  return (
    <div className="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-5 stat-card relative overflow-hidden">
      {/* 見えないグラデーションアクセント（右上角） */}
      <div
        className={`absolute top-0 right-0 w-20 h-20 bg-gradient-to-br opacity-10 ${gradientColorClass}`}
      ></div>
      {/* アイコンインジケーター */}
      <div className="flex justify-between items-start">
        <p className="text-sm font-medium text-gray-500 dark:text-gray-400">
          {label}
        </p>
        {change && (
          <div
            className={`rounded-full w-8 h-8 flex items-center justify-center ${
              isPositive
                ? "bg-green-100 dark:bg-green-900/30 text-green-600 dark:text-green-400"
                : "bg-red-100 dark:bg-red-900/30 text-red-600 dark:text-red-400"
            }`}
          >
            {isPositive ? (
              <svg
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M5 10l7-7m0 0l7 7m-7-7v18"
                />
              </svg>
            ) : (
              <svg
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
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

      <p className="mt-2 text-3xl font-bold text-gray-900 dark:text-gray-100">
        {value}
      </p>

      {change && (
        <div className="mt-4 flex items-center">
          <div
            className={`h-1 w-12 rounded-full ${
              isPositive
                ? "bg-gradient-to-r from-green-400 to-emerald-500"
                : "bg-gradient-to-r from-red-500 to-orange-400"
            }`}
          ></div>
          <p
            className={`ml-2 text-sm font-medium ${
              isPositive
                ? "text-green-600 dark:text-green-400"
                : "text-red-600 dark:text-red-400"
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
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
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
