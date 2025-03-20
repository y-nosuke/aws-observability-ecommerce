import Link from 'next/link';

interface ProductTablePaginationProps {
  pageParam: number;
  totalPages: number;
  totalProducts: number;
}

export default function ProductTablePagination({
  pageParam,
  totalPages,
  totalProducts,
}: ProductTablePaginationProps) {
  return (
    <div className="flex justify-between items-center mt-5">
      <div className="text-sm text-gray-700">
        全{totalProducts}件中 {(pageParam - 1) * 10 + 1}-{Math.min(pageParam * 10, totalProducts)}
        件を表示
      </div>
      <div className="flex space-x-2">
        <Link
          href={`/products?page=${Math.max(1, pageParam - 1)}`}
          className={`px-3 py-1 border rounded ${
            pageParam <= 1
              ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
              : 'bg-white text-blue-600 hover:bg-blue-50'
          }`}
        >
          前へ
        </Link>
        <Link
          href={`/products?page=${Math.min(totalPages, pageParam + 1)}`}
          className={`px-3 py-1 border rounded ${
            pageParam >= totalPages
              ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
              : 'bg-white text-blue-600 hover:bg-blue-50'
          }`}
        >
          次へ
        </Link>
      </div>
    </div>
  );
}
