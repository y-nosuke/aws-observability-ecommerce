FROM node:23-alpine

WORKDIR /app

# 開発に必要なツールのインストール
RUN apk add --no-cache git curl bash

# タイムゾーンの設定
RUN apk add --no-cache tzdata
ENV TZ=Asia/Tokyo

# 依存関係をインストール
COPY package.json package-lock.json ./
RUN npm ci

# Lintとフォーマット設定ファイルをコピー
# COPY .eslintrc.js .eslintignore .prettierrc.js .prettierignore ./

# アプリケーションのソースコードをコピー
COPY . .

# 開発サーバーを起動
CMD ["npm", "run", "dev"]
