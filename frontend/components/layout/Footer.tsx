export default function Footer() {
  return (
    <footer className="bg-gray-800 text-white py-6">
      <div className="container mx-auto px-4">
        <div className="text-center">
          <p>
            © {new Date().getFullYear()}{" "}
            AWSオブザーバビリティ学習用eコマースアプリ
          </p>
          <p className="text-sm mt-2">
            このアプリケーションは学習目的で作成されています
          </p>
        </div>
      </div>
    </footer>
  );
}
