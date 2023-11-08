package controllers

import (

	"github.com/gin-gonic/gin"

	"github.com/freedomknight/linewebhooks/forms"
	"github.com/freedomknight/linewebhooks/models"
	db "github.com/freedomknight/linewebhooks/database"
)

type Webhook struct {}

func (l *Webhook) Receive(c *gin.Context) {
    var request forms.WebhookRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    messages := make([]interface{}, 0)

    for _, event := range request.Events {
        message := models.Message {
            WebhookEventID  : event.WebhookEventID,
            DestinationID   : request.Destination,
            SourceId        : event.Source.UserID,
            Type            : event.Message.Type,
            MessageID       : event.Message.ID,
            QuotedMessageID : event.Message.QuotedMessageID,
            QuoteToken      : event.Message.QuoteToken,
            Text            : event.Message.Text,
            Timestamp       : event.Timestamp,
            ReplyToken      : event.ReplyToken,
        }
        messages = append(messages, message)
    }

    _, err := db.Get().Collection("messages").InsertMany(db.GetCtx(), messages)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "message": "Success",
    })
}
