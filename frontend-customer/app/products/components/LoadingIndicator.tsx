export default function LoadingIndicator() {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      {[...Array(6)].map((_, i) => (
        <div key={i} className="bg-white rounded-lg overflow-hidden shadow animate-pulse">
          <div className="h-48 bg-gray-300"></div>
          <div className="p-4">
            <div className="h-6 bg-gray-300 rounded mb-2"></div>
            <div className="h-12 bg-gray-200 rounded mb-4"></div>
            <div className="flex justify-between items-center">
              <div className="h-6 w-20 bg-gray-300 rounded"></div>
              <div className="h-10 w-24 bg-gray-300 rounded"></div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
