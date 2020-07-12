package group

import "github.com/mapleFU/QQGroupBot/qqbot/data/group/message"

type ArrayRespMessage struct {
	GroupID    string          `json:"group_id"`
	Message    message.Message `json:"message"`
	AutoEscape bool            `json:"auto_escape"`
}

type StringRespMessage struct {
	GroupID    string `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
