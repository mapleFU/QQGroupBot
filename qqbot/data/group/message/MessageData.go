package message

//import (
//	"encoding/json"
//	"fmt"
//)

// 定义的 QQ 消息的数据结构, 表示单条的消息
type MessageData struct {
//	field for file
	File string `json:"file,omitempty"`
	Url string `json:"url,omitempty"`
//	field for text
	Text string `json:"text,omitempty"`
//	field for face
	Id string `json:"id,omitempty"`
//	field for at
	QQ string `json:"qq"`
}

type MessageSegment struct {
	Type string `json:"type,omitempty"`
	// may be message type ?
	Data MessageData `json:"data,omitempty"`
}

type Message []MessageSegment
//type Message struct {
//	Message []MessageSegment `json:"message,omitempty,string"`
//}

//// 这是什么几把坑
//func (message *Message) UnmarshalJSON(data []byte) error {
//	var v []string
//	if err:= json.Unmarshal(data, &v);err!=nil{
//		fmt.Println(err)
//		return err
//	}
//	message.Message = make([]MessageSegment, 0)
//	for str := range v {
//		message.Message = append(json.)
//	}
//}