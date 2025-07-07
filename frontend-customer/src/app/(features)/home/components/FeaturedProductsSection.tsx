import Link from 'next/link';

import AnimateInView from '@/components/ui/AnimateInView';
import ProductCard from '@/components/ui/ProductCard';
import { Product } from '@/services/products/types';

interface FeaturedProductsSectionProps {
  products: Product[];
  title: string;
  viewAllLink: string;
  viewAllText: string;
}

export default function FeaturedProductsSection({
  products,
  title,
  viewAllLink,
  viewAllText,
}: FeaturedProductsSectionProps) {
  return (
    <section className="mb-16">
      <div className="container mx-auto px-6">
        <AnimateInView direction="up" delay={100}>
          <div className="mb-8 flex items-center justify-between">
            <h2 className="text-2xl font-bold">{title}</h2>
            <Link
              href={viewAllLink}
              className="text-primary hover:text-primary-dark flex items-center font-medium transition-colors"
            >
              {viewAllText}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="ml-1 h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fillRule="evenodd"
                  d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                  clipRule="evenodd"
                />
              </svg>
            </Link>
          </div>
        </AnimateInView>

        <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
          {products.map((product, index) => (
            <AnimateInView key={product.id} direction="up" delay={200 + index * 100}>
              <ProductCard
                id={product.id}
                name={product.name}
                description={product.description}
                price={product.price}
                salePrice={product.salePrice}
                isNew={product.isNew}
                imageUrl={product.imageUrl}
              />
            </AnimateInView>
          ))}
        </div>
      </div>
    </section>
  );
}
