#!/bin/bash

# APIのベースURL
BASE_URL="http://backend-api.localhost/api"

# 色の定義
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# ヘルスチェック
echo -e "${GREEN}Testing health endpoint:${NC}"
curl -s "${BASE_URL}/health" | jq
echo ""

# 商品一覧の取得（デフォルトページ）
echo -e "${GREEN}Testing products list (default page):${NC}"
curl -s "${BASE_URL}/products" | jq
echo ""

# 商品一覧の取得（ページサイズ指定）
echo -e "${GREEN}Testing products list with pageSize=5:${NC}"
curl -s "${BASE_URL}/products?pageSize=5" | jq
echo ""

# 商品一覧の取得（ページネーション）
echo -e "${GREEN}Testing products list page 2:${NC}"
curl -s "${BASE_URL}/products?page=2&pageSize=3" | jq
echo ""

# カテゴリー別商品一覧の取得
echo -e "${GREEN}Testing products by category (ID=1):${NC}"
curl -s "${BASE_URL}/products?categoryId=1" | jq
echo ""

# 単一商品の取得
echo -e "${GREEN}Testing get product by ID:${NC}"
curl -s "${BASE_URL}/products/1" | jq
echo ""

# カテゴリー一覧の取得
echo -e "${GREEN}Testing categories list:${NC}"
curl -s "${BASE_URL}/categories" | jq
echo ""

# 存在しない商品の取得（エラーレスポンスのテスト）
echo -e "${GREEN}Testing error response (product not found):${NC}"
curl -s "${BASE_URL}/products/999" | jq
echo ""

echo -e "${GREEN}All tests completed!${NC}"
