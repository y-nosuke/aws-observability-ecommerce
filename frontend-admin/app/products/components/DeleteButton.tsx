'use client';

import { useState } from 'react';

// 削除ボタンのクライアントコンポーネント
export default function DeleteButton({ productId }: { productId: number }) {
  const [isDeleting, setIsDeleting] = useState(false);

  const handleDelete = () => {
    if (window.confirm(`商品ID ${productId} を削除しますか？`)) {
      setIsDeleting(true);
      // ここに実際の削除処理を実装
      console.log(`商品ID ${productId} の削除処理を実行します`);

      // 実際のAPIコールは以下のようになります
      /*
      fetch(`/api/products/${productId}`, {
        method: 'DELETE',
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('削除に失敗しました');
          }
          window.location.reload();
        })
        .catch(error => {
          console.error('削除エラー:', error);
          alert('削除に失敗しました: ' + error.message);
        })
        .finally(() => {
          setIsDeleting(false);
        });
      */

      // デモ用のタイムアウト
      setTimeout(() => {
        setIsDeleting(false);
        alert(`商品ID ${productId} を削除しました（デモ）`);
      }, 1000);
    }
  };

  return (
    <button
      onClick={handleDelete}
      disabled={isDeleting}
      className="text-red-600 hover:text-red-900 disabled:text-red-300 disabled:cursor-not-allowed"
    >
      {isDeleting ? '削除中...' : '削除'}
    </button>
  );
}
