ポートフォリオとして開発した映画情報、レビューサービスのバックエンドです。

言語はGo、データベースはMySQLを使用しています。

フロントエンドは[別リポジトリ](https://github.com/ikeyu0806/movie-info-frontend)で管理しています。

## 開発環境立ち上げ

`docker-compose -f ./docker-compose.yml up`

## DB Migration

マイグレーションファイル作成
```bash
docker-compose exec golang goose create createReviews sql
```

マイグレーション実行
```bash
docker-compose exec golang goose up
```

テスト実行
```
go test github.com/ikeyu0806/...
```

## Author
Yuki Ikegaya
