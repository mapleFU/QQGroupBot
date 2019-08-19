package message

type MessageType string

const (
	Text  MessageType = "text"
	Image MessageType = "image"
	// 表情
	Face MessageType = "face"
)
