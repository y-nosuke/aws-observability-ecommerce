import { ReactNode } from "react";

type DashboardCardProps = {
  title: string;
  icon?: ReactNode;
  children: ReactNode;
  className?: string;
};

export default function DashboardCard({
  title,
  icon,
  children,
  className = "",
}: DashboardCardProps) {
  return (
    <div
      className={`bg-white dark:bg-gray-800 rounded-lg shadow-lg overflow-hidden hover-lift transition-smooth card-accent-primary ${className}`}
    >
      <div className="px-6 py-4 border-b border-gray-200 dark:border-gray-700 flex items-center">
        {icon && (
          <span className="mr-3 text-primary dark:text-primary-light">
            {icon}
          </span>
        )}
        <h3 className="font-semibold text-gray-700 dark:text-gray-200">
          {title}
        </h3>
      </div>
      <div className="p-6">{children}</div>
    </div>
  );
}
