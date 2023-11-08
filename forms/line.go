package forms

type WebhookRequest struct {
	Destination string   `json:"destination"`
	Events      []Events `json:"events"`
}

type Message struct {
	Type            string `json:"type"`
	ID              string `json:"id"`
	QuotedMessageID string `json:"quotedMessageId"`
	QuoteToken      string `json:"quoteToken"`
	Text            string `json:"text"`
}

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type Source struct {
	Type    string `json:"type"`
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}

type Events struct {
	Type            string          `json:"type"`
	Message         Message         `json:"message"`
	WebhookEventID  string          `json:"webhookEventId"`
	DeliveryContext DeliveryContext `json:"deliveryContext"`
	Timestamp       int64           `json:"timestamp"`
	Source          Source          `json:"source"`
	ReplyToken      string          `json:"replyToken"`
	Mode            string          `json:"mode"`
}
