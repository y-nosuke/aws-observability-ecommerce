'use client';

import Image from 'next/image';
import { useEffect, useRef } from 'react';

interface ImageModalProps {
  isOpen: boolean;
  onClose: () => void;
  imageUrl: string;
  altText: string;
}

export default function ImageModal({ isOpen, onClose, imageUrl, altText }: ImageModalProps) {
  const modalRef = useRef<HTMLDivElement>(null);

  // Escキーでモーダルを閉じる
  useEffect(() => {
    const handleEscape = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        onClose();
      }
    };

    if (isOpen) {
      document.addEventListener('keydown', handleEscape);
      document.body.style.overflow = 'hidden'; // 背景のスクロールを無効化

      // モーダルが開いたときに最初のフォーカス可能な要素にフォーカス
      if (modalRef.current) {
        const focusableElements = modalRef.current.querySelectorAll(
          'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])',
        );
        if (focusableElements.length > 0) {
          (focusableElements[0] as HTMLElement).focus();
        }
      }
    }

    return () => {
      document.removeEventListener('keydown', handleEscape);
      document.body.style.overflow = 'unset'; // スクロールを有効化
    };
  }, [isOpen, onClose]);

  if (!isOpen) return null;

  return (
    <div
      ref={modalRef}
      className="fixed inset-0 z-50 flex items-center justify-center p-4"
      role="dialog"
      aria-modal="true"
      aria-label="商品画像の拡大表示"
    >
      {/* オーバーレイ */}
      <div
        className="bg-opacity-75 absolute inset-0 bg-black transition-opacity"
        onClick={onClose}
        aria-hidden="true"
      />

      {/* モーダルコンテンツ */}
      <div className="relative flex h-full max-h-full w-full max-w-4xl items-center justify-center">
        {/* 閉じるボタン */}
        <button
          onClick={onClose}
          className="bg-opacity-50 hover:bg-opacity-75 absolute top-4 right-4 z-10 rounded-full bg-black p-2 text-white transition-all focus:ring-2 focus:ring-white focus:outline-none"
          aria-label="画像を閉じる"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            aria-hidden="true"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>

        {/* 画像 */}
        <div className="relative h-full max-h-[80vh] w-full max-w-3xl">
          <Image
            src={imageUrl}
            alt={altText}
            fill
            className="object-contain"
            sizes="(max-width: 768px) 100vw, 80vw"
            priority
          />
        </div>
      </div>

      {/* 操作ヒント */}
      <div
        className="absolute bottom-4 left-1/2 -translate-x-1/2 transform text-sm text-white opacity-75"
        aria-live="polite"
      >
        <p>Escキーまたはクリックで閉じる</p>
      </div>
    </div>
  );
}
