package message
// 定义的 QQ 消息的数据结构, 表示单条的消息
type MessageData struct {
//	field for file
	File string `json:"file"`
	Url string `json:"url"`
//	field for text
	Text string `json:"text"`
//	field for face
	Id string `json:"id"`
}

type MessageSegment struct {
	Type string `json:"type"`
	Data MessageType `json:"data"`
}

type Message struct {
	Message []MessageSegment `json:"message"`
} 