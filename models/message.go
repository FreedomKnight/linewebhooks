package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
    ID              primitive.ObjectID `bson:"_id,omitempty"`
    WebhookEventID  string `json:"webhookEventId"  bson:"webhookEventId"`
    DestinationID   string `json:"destinationId" bson:"destinationId"`
    SourceId        string `json:"sourceId" bson:"sourceId"`
	Type            string `json:"type" bson:"type"`
	MessageID       string `json:"messageId" bson:"messageId"`
	QuotedMessageID string `json:"quotedMessageId" bson:"quotedMessageId"`
	QuoteToken      string `json:"quoteToken" bson:"quoteToken"`
	Text            string `json:"text" bson:"text"`
	Timestamp       time.Time  `json:"timestamp" bson:"timestamp"`
    ReplyToken      string `json:"replyToken" bson:"replyToken"`
}
