package service

import (
	"github.com/mapleFU/QQBot/qqbot/data/group"
)

type Servicer interface {
	// Generate Message Channel
	GetChan() (out chan<- *group.ChatRequestData)
	IfAcceptMessage(Request *group.ChatRequestData) bool
	PutRequest(Request *group.ChatRequestData)
	SetOutchan(respChan* chan group.StringRespMessage)
	SendData(data *group.StringRespMessage)

//	logic
	Run()
	Stop()
}

type BaseServicer struct {
	InChan chan *group.ChatRequestData
	// it might be nil
	OutChan *chan group.StringRespMessage
}

func (base *BaseServicer) Stop()  {
	close(base.InChan)
}

func (base *BaseServicer) GetChan() (out chan<- *group.ChatRequestData) {
	return base.InChan
}

func (base *BaseServicer) SendData(data *group.StringRespMessage) {
	if base.OutChan != nil {
		*base.OutChan <- *data
	}
}


func (base *BaseServicer) PutRequest(Request *group.ChatRequestData) {
	base.InChan <- Request
}

func (base *BaseServicer) SetOutchan(respChan* chan group.StringRespMessage)  {
	if respChan == nil {
		return
	} else {
		base.OutChan = respChan
	}
}

func NewBaseServicer() BaseServicer {
	return BaseServicer{
		InChan:make(chan *group.ChatRequestData),
		OutChan:nil,
	}
}

