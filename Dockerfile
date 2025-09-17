FROM golang:1.23-alpine

WORKDIR /app

# Goモジュールの依存関係をキャッシュ
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . ./

# ビルド
RUN go build -o server .

# APIサーバ起動
CMD ["./server"]