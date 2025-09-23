# study_abroad_journal_back

バックエンド用のリポジトリです．

## 必要な環境

- Docker
- Docker Compose

## 手順

1. リポジトリをクローンする．front と back を同じ階層に置く．

2. back ディレクトリの中で，Docker コンテナをビルド & 起動する．

```
docker compose up --build
```

3. サーバが起動したら接続を確認．

```
http://localhost:8080/health
```

->OK from Go backend と表示される．

```
http://localhost:3000
```

->フロントページが表示される．
