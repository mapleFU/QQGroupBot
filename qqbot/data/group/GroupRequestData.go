package group

import "github.com/mapleFU/QQBot/qqbot/data/group/message"

//import "github.com/mapleFU/QQBot/qqbot/data/group/message"

type ChatRequestData struct {
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	// 消息子类型，正常消息是 normal，匿名消息是 anonymous，系统提示（如「管理员已禁止群内匿名聊天」）是 notice
	SubType   string `json:"sub_type"`
	MessageID int32  `json:"message_id"`

	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	//Anonymous
	Message message.Message `json:"message"`
}
