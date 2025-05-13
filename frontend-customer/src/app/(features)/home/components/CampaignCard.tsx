import Link from "next/link";

interface CampaignCardProps {
  title: string;
  description: string;
  badgeText: string;
  badgeColor: "blue" | "red" | "green" | "purple" | "yellow";
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
    blue: "bg-blue-500",
    red: "bg-red-500",
    green: "bg-gradient-to-r from-emerald-400 to-green-500",
    purple: "bg-purple-500",
    yellow: "bg-gradient-to-r from-yellow-400 to-orange-500",
  };

  return (
    <div
      className={`${bgGradient} rounded-2xl shadow-lg overflow-hidden card-hover`}
    >
      <div className="p-6 md:p-8 flex flex-col h-full">
        <div className={`badge mb-4 self-start ${badgeClasses[badgeColor]}`}>
          {badgeText}
        </div>
        <h3 className="text-xl md:text-2xl font-bold text-white mb-2">
          {title}
        </h3>
        <p className="text-white/80 mb-6">{description}</p>
        <Link
          href={linkUrl}
          className="mt-auto inline-block bg-white text-primary px-6 py-2 rounded-full font-medium hover:bg-opacity-90 transition-colors self-start"
        >
          {linkText}
        </Link>
      </div>
    </div>
  );
}
