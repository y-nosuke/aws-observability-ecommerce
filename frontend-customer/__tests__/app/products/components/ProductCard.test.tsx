import { Product } from '@/lib/api/products';
import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import React from 'react'; // Reactをインポート

// ProductCardコンポーネントをモック化
jest.mock('next/link', () => {
  const LinkComponent = ({ children, href }: { children: React.ReactNode; href: string }) => {
    return <a href={href}>{children}</a>;
  };
  LinkComponent.displayName = 'MockNextLink';
  return LinkComponent;
});

// ProductCardコンポーネントをインポート
// 注: テスト対象のコンポーネントを実際にインポートする際のパスは
// アプリケーションの構造によって異なる場合があります
import ProductCard from '@/app/products/components/ProductCard';

describe('ProductCard', () => {
  const mockProduct: Product = {
    id: 1,
    name: 'テスト商品',
    description: 'これはテスト用の商品です',
    price: 1000,
    image_url: 'https://example.com/test.jpg',
    category_id: 1,
  };

  it('商品名が正しく表示される', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText('テスト商品')).toBeInTheDocument();
  });

  it('商品説明が正しく表示される', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText('これはテスト用の商品です')).toBeInTheDocument();
  });

  it('商品価格が正しくフォーマットされて表示される', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText('¥1,000')).toBeInTheDocument();
  });

  it('商品詳細へのリンクが正しく設定される', () => {
    render(<ProductCard product={mockProduct} />);
    const link = screen.getByText('詳細を見る');
    expect(link).toBeInTheDocument();
    expect(link.closest('a')).toHaveAttribute('href', '/products/1');
  });

  it('画像がない場合は代替テキストが表示される', () => {
    const productWithoutImage = { ...mockProduct, image_url: '' };
    render(<ProductCard product={productWithoutImage} />);
    expect(screen.getByText('商品画像なし')).toBeInTheDocument();
  });
});
