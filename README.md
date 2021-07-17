ポートフォリオとして開発した映画情報、レビューサービスのバックエンドです。

言語はGo、データベースはMySQLを使用しています。

フロントエンドは[別リポジトリ](https://github.com/ikeyu0806/movie-info-frontend)で管理しています。

## 開発環境立ち上げ

`docker-compose -f ./docker-compose.yml up`


`docker exec -it golang_backend bash`

## DB Migration

マイグレーションファイル作成
```bash
docker-compose exec golang sql-migrate new createReviews sql
```

マイグレーション実行
```bash
docker-compose exec golang sql-migrate up
```

テスト実行
```
go test github.com/ikeyu0806/...
```

## Deploy
herokuにデプロイしています。

```
git push heroku master
```

また、AWS ECSでも実行できるようterraformを用意しています。

https://github.com/ikeyu0806/movie-backend-terraform

## Author
Yuki Ikegaya
