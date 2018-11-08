package service

import (
	"github.com/mapleFU/QQBot/qqbot/data/group"
	"github.com/mapleFU/QQBot/qqbot/Requester"
	"fmt"
	"sync"
)

type Manager struct {
	serviceMap map[string]Servicer
	requester Requester.Requester
	receiver chan group.ChatResponseData
	strReceiver chan group.StringRespMessage
	// 管理的群组
	managedGroups []string
	// 同步的 lock
	serviceLock sync.Mutex
}

// deprecated
// copy read only array, should be readonly
func (manager *Manager) ListManagedGroups() []string {
	return manager.managedGroups
}

// deprecated
// copy read only array, should be readonly
func (manager *Manager) GetServiceMap() *map[string]Servicer {
	return &manager.serviceMap
}

func (manager *Manager) copyServiceMap() map[string]Servicer {
	panic("impl me")
	return nil
}

func (manager *Manager) AddManagedGroups(groupId string)  {

	manager.managedGroups = append(manager.managedGroups, groupId)
}

func (manager *Manager) DeleteManagedGroups(groupId string)  {
	a := manager.managedGroups
	i := 0
	var value string
	find := false
	for i, value = range a {
		if value == groupId {
			find = true
			break
		}
	}
	if !find {
		return
	}
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
}

func (manager *Manager) AddService(servicer Servicer, name string)  {
	manager.serviceLock.Lock()
	defer manager.serviceLock.Unlock()

	manager.serviceMap[name] = servicer
	servicer.SetOutchan(&manager.strReceiver)
	go servicer.Run()
}

func (manager *Manager) RemoveService(name string) bool {
	manager.serviceLock.Lock()
	defer manager.serviceLock.Unlock()
	_, ok := manager.serviceMap[name]
	delete(manager.serviceMap, name)
	return ok
}

func (manager *Manager) RecvRequest(request *group.ChatRequestData) {
	manager.serviceLock.Lock()
	defer manager.serviceLock.Unlock()

	for k, v := range manager.serviceMap {
		if v.IfAcceptMessage(request) {
			fmt.Println("Call service " + k)
			v.PutRequest(request)
		}
	}
}



func NewManager(Addr string) *Manager {
	this := &Manager{
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

	go func(manager *Manager) {
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
