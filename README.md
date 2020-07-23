ポートフォリオとして開発中の映画情報サービスのバックエンドです。

言語はGo、データベースはMySQLを使用しています。

フロントエンドは[別リポジトリ](https://github.com/ikeyu0806/movie-info-frontend)で管理しています。

## 開発環境立ち上げ

`docker-compose -f ./docker-compose.yml up`

## DB Migration

マイグレーションファイル作成
```
docker-compose exec golang goose create createReviews sql
```

マイグレーション実行
```
docker-compose exec golang goose up
```

## Author
Yuki Ikegaya
