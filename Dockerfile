FROM golang:1.23-alpine

WORKDIR /app

# Goモジュールの依存関係をキャッシュ
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . /app/src

# ビルド
RUN go build -o server /app/src/main.go

# APIサーバ起動
CMD ["./server"]