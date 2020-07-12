package requester

import (
	"bytes"
	"encoding/json"
	"github.com/mapleFU/QQGroupBot/qqbot/data/group"
	"net/http"
)

type Requester struct {
	// 发送的 Addr
	Addr     string
	respChan chan group.ChatResponseData
	strMsg   chan group.StringRespMessage
}

func (this *Requester) sendResponse(data group.ChatResponseData) {
	jsonValue, _ := json.Marshal(data)
	http.Post(this.Addr, "application/json", bytes.NewBuffer(jsonValue))
}

func (this *Requester) SendResponse(data group.ChatResponseData) {
	this.respChan <- data
}

func (this *Requester) SendGroupChat(data group.StringRespMessage) {
	this.strMsg <- data
}

func (this *Requester) sendGroupChat(data group.StringRespMessage) {
	jsonValue, _ := json.Marshal(data)
	http.Post(this.Addr+"/send_group_msg", "application/json", bytes.NewBuffer(jsonValue))
}

func NewRequester(addr string) *Requester {
	reqter := Requester{
		Addr:     addr,
		respChan: make(chan group.ChatResponseData, 10),
		strMsg:   make(chan group.StringRespMessage, 10),
	}
	go func() {
		for req := range reqter.respChan {
			go reqter.sendResponse(req)
		}
	}()

	go func() {
		for req := range reqter.strMsg {
			go reqter.sendGroupChat(req)
		}
	}()

	return &reqter
}
