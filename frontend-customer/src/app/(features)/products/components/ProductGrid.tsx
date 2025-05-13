import ProductCard from "@/components/ui/ProductCard";
import { Product } from "@/services/products/types";

interface ProductGridProps {
  products: Product[];
}

export default function ProductGrid({ products }: ProductGridProps) {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      {products.map((product) => (
        <ProductCard
          key={product.id}
          id={product.id}
          name={product.name}
          description={product.description}
          price={product.price}
          salePrice={product.salePrice}
          isNew={product.isNew}
          imageUrl={product.imageUrl}
        />
      ))}
    </div>
  );
}
