export default function AdminHeader() {
  return (
    <header className="bg-gray-900 text-white p-4">
      <div className="container mx-auto">
        <div className="flex justify-between items-center">
          <h1 className="text-xl font-bold">管理者ダッシュボード</h1>
          <div>
            <span className="text-sm">管理者</span>
          </div>
        </div>
      </div>
    </header>
  );
}
