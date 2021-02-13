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

goose の設定ファイルは`db/dbconf.yml`です。

## test

テスト実行

```
go test github.com/ikeyu0806/...
```

## Author

Yuki Ikegaya
