FROM golang:1.24-alpine

WORKDIR /app

# 開発に必要なツールのインストール
RUN apk add --no-cache git curl bash

# airのインストール（ホットリロード用）
RUN go install github.com/air-verse/air@latest

# デバッグツールのインストール
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# モックジェネレーターのインストール
RUN go install go.uber.org/mock/mockgen@latest

# goimportsのインストール
RUN go install golang.org/x/tools/cmd/goimports@latest

# golangci-lintのインストール
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# AWS CLI v2とawscli-localのインストール（LocalStackとの連携用）
RUN apk add --no-cache python3 py3-pip unzip
# 仮想環境を作成してawscli-localをインストール
RUN python3 -m venv /opt/venv && \
    . /opt/venv/bin/activate && \
    pip3 install awscli-local && \
    ln -s /opt/venv/bin/aws /usr/local/bin/aws && \
    ln -s /opt/venv/bin/awslocal /usr/local/bin/awslocal

# タイムゾーンの設定
RUN apk add --no-cache tzdata
ENV TZ=Asia/Tokyo

# アプリケーションの依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# airの設定
COPY .air.toml .

# ポートを公開
EXPOSE 8000

# ホットリロードでアプリケーションを起動
CMD ["air", "-c", ".air.toml"]
