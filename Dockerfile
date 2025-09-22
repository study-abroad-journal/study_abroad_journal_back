# FROM golang:1.24-alpine

# WORKDIR /app

# # Goモジュールの依存関係をキャッシュ
# COPY go.mod ./
# #COPY go.su[m] ./
# RUN go mod download
# RUN go mod tidy

# # ソースコードをコピー
# COPY . ./

# # ビルド
# RUN go build -o server .

# # APIサーバ起動
# CMD ["./server"]

FROM golang:1.24-alpine

WORKDIR /app

# まずgo.modとソースコードを一緒にコピー
COPY go.mod ./
COPY . ./

# 依存関係を解決
RUN go mod tidy
RUN go mod download

# ビルド
RUN go build -o server .

CMD ["./server"]