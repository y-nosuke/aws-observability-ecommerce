import Link from 'next/link';

interface CampaignCardProps {
  title: string;
  description: string;
  badgeText: string;
  badgeColor: 'blue' | 'red' | 'green' | 'purple' | 'yellow';
  bgGradient: string;
  linkText: string;
  linkUrl: string;
}

export default function CampaignCard({
  title,
  description,
  badgeText,
  badgeColor,
  bgGradient,
  linkText,
  linkUrl,
}: CampaignCardProps) {
  // バッジの色に応じたクラス
  const badgeClasses = {
    blue: 'bg-blue-500',
    red: 'bg-red-500',
    green: 'bg-gradient-to-r from-emerald-400 to-green-500',
    purple: 'bg-purple-500',
    yellow: 'bg-gradient-to-r from-yellow-400 to-orange-500',
  };

  return (
    <div className={`${bgGradient} card-hover overflow-hidden rounded-2xl shadow-lg`}>
      <div className="flex h-full flex-col p-6 md:p-8">
        <div className={`badge mb-4 self-start ${badgeClasses[badgeColor]}`}>{badgeText}</div>
        <h3 className="mb-2 text-xl font-bold text-white md:text-2xl">{title}</h3>
        <p className="mb-6 text-white/80">{description}</p>
        <Link
          href={linkUrl}
          className="text-primary hover:bg-opacity-90 mt-auto inline-block self-start rounded-full bg-white px-6 py-2 font-medium transition-colors"
        >
          {linkText}
        </Link>
      </div>
    </div>
  );
}
