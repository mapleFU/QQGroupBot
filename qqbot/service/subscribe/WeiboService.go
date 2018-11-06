package subscribe

import (
	"github.com/mmcdole/gofeed"
	"time"
	"github.com/mapleFU/QQBot/qqbot/data/group"
	"github.com/mapleFU/QQBot/qqbot/service"
)

type WeiboService struct {
	Subscribe
	ServiceUrl string
}



func NewWeiboService(weiboUrl string) *WeiboService {
	return &WeiboService{
		ServiceUrl:weiboUrl,
		Subscribe: Subscribe{service.NewBaseServicer()},
	}
}

func (self *WeiboService) Run() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(self.ServiceUrl)
	newest := feed.Items[0]

	// 考虑任务如何中止
	for true  {
		// 10 分钟一次
		time.Sleep(time.Minute * 10)
		feed, _ := fp.ParseURL(self.ServiceUrl)
		for _, item := range feed.Items {
			if item.Title == newest.Title {
				newest = item
				break
			} else {
				Resp := group.StringRespMessage{
					Message:item.Title + " : \n" + item.Content,
					GroupID:"",
					AutoEscape:true,
				}
				*self.OutChan <- &Resp
			}
		}
	}
}
