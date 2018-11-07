package query

import (
	"github.com/mapleFU/QQBot/qqbot/service"
	"github.com/mapleFU/QQBot/qqbot/data/group"
)

type QueryService struct {
	service.BaseServicer

}

// 自己还是没法判断是不是要查询吧...
func (*QueryService) IfAcceptMessage(Request *group.ChatRequestData) bool {
	panic("implement me")
}


