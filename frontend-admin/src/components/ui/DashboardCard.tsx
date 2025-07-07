import { ReactNode } from 'react';

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
  className = '',
}: DashboardCardProps) {
  return (
    <div
      className={`hover-lift transition-smooth card-accent-primary overflow-hidden rounded-lg bg-white shadow-lg dark:bg-gray-800 ${className}`}
    >
      <div className="flex items-center border-b border-gray-200 px-6 py-4 dark:border-gray-700">
        {icon && <span className="text-primary dark:text-primary-light mr-3">{icon}</span>}
        <h3 className="font-semibold text-gray-700 dark:text-gray-200">{title}</h3>
      </div>
      <div className="p-6">{children}</div>
    </div>
  );
}
