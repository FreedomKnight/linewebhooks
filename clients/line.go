package clients

import (
    "github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/freedomknight/linewebhooks/configs"
)

var client *linebot.Client

func init() {
    config := configs.Get()
    config.SetConfigName("line")
    config.SetConfigType("yaml")
    err := config.ReadInConfig()
    if err != nil {
        panic(err.Error())
    }

    println("line", config.Get("channel.secret"), config.Get("channel.token"))
    client, err = linebot.New(config.GetString("channel.secret"), config.GetString("channel.token"))
}

func GetLineClient() *linebot.Client {
    return client
}
