package base

import "time"

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
	now := time.Now().Unix()
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
