package clients

import (
    "github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/freedomknight/linewebhooks/configs"
	"github.com/freedomknight/linewebhooks/container"
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

    client, err = linebot.New(config.GetString("channel.secret"), config.GetString("channel.token"))
    container.Provide(func () *linebot.Client {
        return client
    })
}

func GetLineClient() *linebot.Client {
    var c *linebot.Client
    container.Invoke(func (client *linebot.Client) {
        c = client
    });

    return c
}
