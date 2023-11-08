## Install

### run mongo on docker

```
make build-mongo
```

### Install Dependency

```
go mod tidy
```

### Run API Server

Default run on `localhost:8080`

```
go run main
```

or

```
make run
```

## Setup Your LINE secret and access token

fill up your secret on configs/line.yaml

You can find these information on LINE Developer Dashboard

1. Secret: Dashboard > {Your LINE Message Channel} > Basic settings Tab > Channel secret
2. Access Token: Dashboard > {Your LINE Message Channel} > Messaging API Tab > Channel access token (long-lived) 

## Simple API Document

### Reply By ID

由於文件說明需要一隻 API 做 reply 因此獨立一隻 API ，透過之前儲存於 mongo 的 id 去做回覆訊息

POST {host}/api/v1/messages

Request
```json
{
    "id": "xxxx" // mongo document id
    "text": "what you want to reply"
}
```

### List By User ID
GET {host}/api/v1/messages?source="{line-user-id}"

Response
```json
{
    "data": [
        {
            "ID": "{string}", // mongo object id
            "webhookEventId": "", // line web hook id
            "sourceId": "{string}", // line user id
            "type": "text", // text
            "messageId": "",
            "quotedMessageId": "",
            "quoteToken": "",
            "text": "{string}", // message
            "timestamp": "{string}", // ISO8601 timestamp string
            "replyToken": "String" // reply token
        }
    ]
}
```

## Demo Video

[Demo Video Link](https://capture.dropbox.com/wWHxBjKHf2523JGB)

[PostMan Collection](postmans/line-webhooks.postman_collection.json)
