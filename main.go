package main

import (
	"fmt"

	"github.com/freedomknight/linewebhooks/configs"
	"github.com/freedomknight/linewebhooks/routes"
)

func main() {
    config := configs.Get()
    engine := routes.New()
    config.SetConfigName("app")
    config.SetConfigType("yaml")
    err := config.ReadInConfig()
    if err != nil {
        panic(err.Error())
    }
    host, port := config.Get("app.host"), config.Get("app.port")
    engine.Run(fmt.Sprintf("%s:%d", host, port))
}
