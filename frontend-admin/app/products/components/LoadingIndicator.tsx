export default function LoadingIndicator() {
  return (
    <div className="bg-white rounded-lg shadow overflow-hidden">
      <div className="animate-pulse">
        <div className="h-12 bg-gray-200"></div>
        <div className="h-64 bg-gray-100"></div>
      </div>
    </div>
  );
}
