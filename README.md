ポートフォリオとして開発した映画情報、レビューサービスのバックエンドです。

言語は Go、データベースは MySQL を使用しています。

フロントエンドは[別リポジトリ](https://github.com/ikeyu0806/movie-info-frontend)で管理しています。

## 開発環境立ち上げ

`docker-compose -f ./docker-compose.yml up`

## DB Migration

開発環境立ち上げ後に手動で migration を実行する必要があります。

マイグレーションファイル作成

```bash
docker-compose exec golang goose create createReviews sql
```

マイグレーション実行

```bash
docker-compose exec golang goose up
```

production 環境

```
goose -env production up
```

goose の設定ファイルは`db/dbconf.yml`です。

## test

テスト実行

```
go test github.com/ikeyu0806/...
```

## デプロイ準備

1. `db/db_connect.yml`と`db/dbconf.yml`に RDS の情報を追記

2. `docker build -t movie-info-backend -f ./Dockerfile.production .`

3. できたイメージを ECR に push

4. 残りは terraform で構築。1-3 も自動化したいが未対応

5. prodution での migration は現状手動

### ローカルでの production 用 docker image の確認

`docker run -it {image_id} sh`

## Author

Yuki Ikegaya
