package controllers

import (
	"github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/freedomknight/linewebhooks/clients"
	"github.com/freedomknight/linewebhooks/models"
	"github.com/freedomknight/linewebhooks/forms"
	db "github.com/freedomknight/linewebhooks/database"
)

type Message struct {}

// TODO if you have more time, you should make pagination
func (m *Message) Index(c *gin.Context) {
    messages := make([]models.Message, 0)

    println("query", c.DefaultQuery("destination", ""))
    filter := bson.M{"destinationId": c.DefaultQuery("destination", "")}

    cursor, err := db.Get().Collection("messages").Find(db.GetCtx(), filter)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    for cursor.Next(db.GetCtx()) {
        var message models.Message
        cursor.Decode(&message)
        messages = append(messages, message)
    }

    c.JSON(200, gin.H{
        "data": messages,
    })
}

func (m *Message) Reply(c *gin.Context) {
    client := clients.GetLineClient()
    var message models.Message

    var request forms.MessageReply

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    objId, _ := primitive.ObjectIDFromHex(request.ID)
    filter := bson.M{"_id": objId}
    err := db.Get().Collection("messages").FindOne(db.GetCtx(), filter).Decode(&message)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    _, err = client.ReplyMessage(message.ReplyToken, linebot.NewTextMessage(request.Text)).Do()
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "message": "Success",
    })
}
