import Link from 'next/link';

export default function Pagination({
  currentPage,
  totalPages,
}: {
  currentPage: number;
  totalPages: number;
}) {
  return (
    <div className="flex justify-center items-center space-x-2 mt-8">
      <Link
        href={`/products?page=${Math.max(1, currentPage - 1)}`}
        className={`px-4 py-2 rounded-lg ${
          currentPage <= 1
            ? 'bg-gray-200 text-gray-400 cursor-not-allowed'
            : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
        }`}
      >
        前へ
      </Link>

      <div className="text-gray-700">
        {currentPage} / {totalPages}
      </div>

      <Link
        href={`/products?page=${Math.min(totalPages, currentPage + 1)}`}
        className={`px-4 py-2 rounded-lg ${
          currentPage >= totalPages
            ? 'bg-gray-200 text-gray-400 cursor-not-allowed'
            : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
        }`}
      >
        次へ
      </Link>
    </div>
  );
}
