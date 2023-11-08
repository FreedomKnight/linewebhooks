package controllers

import (
	"github.com/gin-gonic/gin"
    "github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/freedomknight/linewebhooks/clients"
	db "github.com/freedomknight/linewebhooks/database"
	"github.com/freedomknight/linewebhooks/models"
)

type Webhook struct {}

func ParseMessages(events []*linebot.Event) []interface{} {
    messages := make([]interface{}, 0)

    for _, event := range events {
        if event.Type == linebot.EventTypeMessage {
            switch m := event.Message.(type) {
                case  *linebot.TextMessage:
                    message := models.Message {
                        WebhookEventID  : event.WebhookEventID,
                        SourceId        : event.Source.UserID,
                        Type            : "text",
                        MessageID       : m.ID,
                        Text            : m.Text,
                        Timestamp       : event.Timestamp,
                        ReplyToken      : event.ReplyToken,
                    }

                    messages = append(messages, message)
            }
        }
    }

    return messages
}

func (l *Webhook) Receive(c *gin.Context) {
    // Use Client to parse will verify webhook event signature.
    bot := clients.GetLineClient()
    events, err := bot.ParseRequest(c.Request)

    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    _, err = db.Get().Collection("messages").InsertMany(db.GetCtx(), ParseMessages(events))
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "message": "Success",
    })
}
