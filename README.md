ポートフォリオとして開発中の映画情報サービスのバックエンドです。

言語はGo、データベースはMySQLを使用しています。

## 開発環境立ち上げ

`docker-compose -f ./docker-compose.yml up`

## DB Migration

```
docker-compose exec goose create createUsers sql
docker-compose exec golang goose up
```

## Author
Yuki Ikegaya
