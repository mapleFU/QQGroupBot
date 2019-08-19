package interact

import (
	"strings"
	"sync"

	"github.com/mapleFU/QQBot/qqbot/data/group"
	"github.com/mapleFU/QQBot/qqbot/data/group/message"
	"github.com/mapleFU/QQBot/qqbot/service"
)

type RollService struct {
	service.BaseServicer
	innerData map[string][]string // 内部的 roll 数据
	innerLock sync.Mutex
}

func (self *RollService) checkRollOrInner(segment *message.MessageSegment) {

}

func (self *RollService) IfAcceptMessage(Request *group.ChatRequestData) bool {
	for _, data := range Request.Message {
		if data.Type == "text" {
			strData := data.Data.Text
			// 开始 roll 的数据
			if strings.Contains(strData, "roll") {
				return true
			}
			// 类型为 参加 ROLL 的数据
			if val, ok := self.innerData[strData]; ok {
				return true
			}
		}
	}
	return false
}

func (self *RollService) Run() {
	for data := range self.InChan {

	}
}
