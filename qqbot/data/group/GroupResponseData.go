package group

import "github.com/mapleFU/QQGroupBot/qqbot/data/group/message"

type ChatResponseData struct {
	Reply      message.Message `json:"reply"`
	AutoEscape bool            `json:"auto_escape"`
	AtSender   bool            `json:"at_sender"`
}
