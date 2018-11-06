package service

import (
	"github.com/mapleFU/QQBot/qqbot/data/group"
	"github.com/mapleFU/QQBot/qqbot/Requester"
)

type manager struct {
	serviceMap map[string]Servicer
	requester Requester.Requester
	receiver chan group.ChatResponseData
	strReceiver chan group.StringRespMessage
	// 管理的群组
	managedGroups []string
}

func (manager *manager) AddManagedGroups(groupId string)  {
	manager.managedGroups = append(manager.managedGroups, groupId)
}
func (manager *manager) AddService(servicer Servicer, name string)  {
	manager.serviceMap[name] = servicer
	servicer.SetOutchan(&manager.strReceiver)
	go servicer.Run()
}

func (manager *manager) RemoveService(name string) {
	delete(manager.serviceMap, name)
}

func (manager *manager) RecvRequest(request *group.ChatRequestData) {
	for _, v := range manager.serviceMap {
		if v.IfAcceptMessage(request) {
			v.PutRequest(request)
		}
	}
}

func NewManager(Addr string) *manager {
	this := &manager{
		serviceMap:make(map[string]Servicer),
		requester:*Requester.NewRequester(Addr),
		receiver:make(chan group.ChatResponseData, 5),
		strReceiver: make(chan group.StringRespMessage, 5),
		managedGroups: make([]string, 0),
	}

	go func() {
		for data := range this.receiver {
			this.requester.SendResponse(data)
		}
	}()

	go func(manager *manager) {
		for data := range this.strReceiver {
			for _, groupID := range manager.managedGroups {
				data.GroupID = groupID
				manager.requester.SendGroupChat(data)
			}
			//this.requester.SendGroupChat(&data)
		}
	}(this)
	return this
}