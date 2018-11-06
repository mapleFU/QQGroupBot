package subscribe

import (
	"github.com/mmcdole/gofeed"
	"time"
	"github.com/mapleFU/QQBot/qqbot/data/group"
	"github.com/mapleFU/QQBot/qqbot/service"
	"fmt"
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

func buildService(item *gofeed.Item) group.StringRespMessage {

	Resp := group.StringRespMessage{
		Message:item.Title + " : \n" + item.Description + "\n链接：" + item.Link,
		GroupID:"",
		AutoEscape:true,
	}
	return Resp
}

func (self *WeiboService) Run() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(self.ServiceUrl)
	newest := feed.Items[0]

	if self.OutChan == nil {
		fmt.Println("Bug. self.Outchan is nil")
	} else {
		fmt.Println("Send News")
		*self.OutChan <- buildService(newest)
	}
	fmt.Println("Send News Done")
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
				Resp := buildService(item)
				*self.OutChan <- Resp
			}
		}
	}
}
