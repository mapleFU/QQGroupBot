package query

import (
	"strings"

	"github.com/mapleFU/QQGroupBot/qqbot/data/group"
)

type MovieSchedule struct {
	QueryService
}

func (mqs *MovieSchedule) IfAcceptMessage(Request *group.ChatRequestData) bool {
	for _, data := range Request.Message {
		if data.Type == "text" {
			if strings.Contains(data.Data.Text, "[") && strings.Contains(data.Data.Text, "]") {
				return true
			}
		}
	}
	return false
}

func (*MovieSchedule) Run() {
	panic("implement me")
}
