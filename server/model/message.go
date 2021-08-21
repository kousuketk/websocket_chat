package model

type Message struct {
	channelID string `json:"channelID"`
	userID    string `json:"userID"`
	content   string `json:"content"`
	sentAt    string `json:"sentAt"`
}

func (m Message) GetChannelID() string {
	return m.channelID
}

func (m Message) GetUserID() string {
	return m.userID
}

func (m Message) GetContent() string {
	return m.content
}

func (m Message) GetSentAt() string {
	return m.sentAt
}
