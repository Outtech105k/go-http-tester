# go-http-tester

Chatting API Tool.

## Overview

Go 言語で構成された、Web API を気軽に試せるサーバーコンテナです。

## Usage

Docker 実行環境が必要です。

Docker コンテナ内に入り、/server/app.go を実行してください。

エンドポイントは以下の通りです。

| Path  | Method | Description    |
| ----- | ------ | -------------- |
| /post | POST   | 投稿を追加     |
| /post | GET    | 投稿履歴を確認 |

投稿は以下のフォーマットで POST します。

```JSON
{
    "body": "Post sentence here."
}

```

投稿一覧は以下の通りに取得されます。なお、時刻は UTC で取得されます。

```JSON
{
    "status": "OK",
    "body": [
        {
            "id": 1,
            "body": "Post sentence here.",
            "posted": "2024-01-01 00:00:00"
        }
    ]
}
```

また、それ以外のレスポンスは以下の通りです。(一例)

```JSON
{
    "status": "Not Found",
    "body": null
}
```

## License

MIT License を適用します。

## Author

Outtech105

[Homepage](http://outtech105.com)

[X](http://x.com/105techno)
