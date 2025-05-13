import AnimateInView from "@/components/ui/AnimateInView";
import CampaignCard from "./CampaignCard";

interface CampaignSectionProps {
  title: string;
  campaigns: Array<{
    id: string;
    title: string;
    description: string;
    badgeText: string;
    badgeColor: 'blue' | 'red' | 'green' | 'purple' | 'yellow';
    bgGradient: string;
    linkText: string;
    linkUrl: string;
  }>;
}

export default function CampaignSection({ title, campaigns }: CampaignSectionProps) {
  return (
    <section className="mb-16">
      <AnimateInView direction="up" delay={200}>
        <div className="container mx-auto px-6">
          <h2 className="text-2xl font-bold mb-8">{title}</h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {campaigns.map((campaign) => (
              <CampaignCard
                key={campaign.id}
                title={campaign.title}
                description={campaign.description}
                badgeText={campaign.badgeText}
                badgeColor={campaign.badgeColor}
                bgGradient={campaign.bgGradient}
                linkText={campaign.linkText}
                linkUrl={campaign.linkUrl}
              />
            ))}
          </div>
        </div>
      </AnimateInView>
    </section>
  );
}
