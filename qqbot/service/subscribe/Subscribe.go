package subscribe

import (
	"github.com/mapleFU/QQGroupBot/qqbot/data/group"
	"github.com/mapleFU/QQGroupBot/qqbot/service"
)

type Subscribe struct {
	service.BaseServicer
}

func (*Subscribe) IfAcceptMessage(Request *group.ChatRequestData) bool {
	return false
}
