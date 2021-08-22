package model

type Message struct {
	ChannelID string `json:"channelID"`
	UserID    string `json:"userID"`
	Content   string `json:"content"`
	SentAt    string `json:"sentAt"`
}

func (m Message) GetChannelID() string {
	return m.ChannelID
}

func (m Message) GetUserID() string {
	return m.UserID
}

func (m Message) GetContent() string {
	return m.Content
}

func (m Message) GetSentAt() string {
	return m.SentAt
}
