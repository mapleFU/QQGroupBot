package subscribe

import (
	"github.com/mapleFU/QQBot/qqbot/service"
	"github.com/mapleFU/QQBot/qqbot/data/group"
)

type Subscribe struct {

	service.BaseServicer
}


func (*Subscribe) IfAcceptMessage(Request *group.ChatRequestData) bool {
	return false
}
