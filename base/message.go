package base

import (
	"bytes"
	"time"
)

const (
	MsgTypeText  = "TEXT"
	MsgTypePic   = "PICTURE"
	MsgTypeVoice = "VOICE"
	MsgTypeTest  = "TEST"
)

type Message struct {
	MessageID   string `json:"messageId" gorm:"comment:消息的唯一标识"`
	MessageType string `json:"messageType" gorm:"comment:消息的类型"`
	Content     string `json:"content" gorm:"comment:消息内容"`

	HubID    string `json:"hubId" gorm:"comment:消息所属聊天id"`
	SenderID string `json:"senderId" gorm:"comment:消息的发出者id"`

	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTime:milli;comment:消息创建时间，毫秒"`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime:milli;comment:消息更新时间，毫秒"`

	Sequence uint32 `json:"sequence" gorm:"comment:消息的序号"`
	Hidden   bool   `json:"hidden" gorm:"comment:是否隐藏"`
}

func NewMessage() Message {
	now := time.Now().UnixMilli()
	return Message{
		MessageID:   "",
		MessageType: "",
		Content:     "",
		HubID:       "",
		SenderID:    "",
		CreatedAt:   now,
		UpdatedAt:   now,
		Sequence:    0,
		Hidden:      false,
	}
}

func (message *Message) Check(c *Client) {
	now := time.Now().UnixMilli()
	message.Content = string(bytes.TrimSpace(bytes.Replace([]byte(message.Content), newline, space, -1)))
	message.CreatedAt = now
	message.UpdatedAt = now
	message.HubID = c.hub.HubID
	message.SenderID = c.kkcrId
}
