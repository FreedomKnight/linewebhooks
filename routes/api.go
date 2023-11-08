package routes

import (
    // "github.com/gin-gonic/gin"
	c "github.com/freedomknight/linewebhooks/controllers"
)

func init() {
    api := engine.Group("api/v1")

    webhook := new(c.Webhook)
    line := new(c.Message)

    api.POST("/webhooks", webhook.Receive)
    api.POST("/messages", line.Reply)
    api.GET("/messages", line.Index)
}
