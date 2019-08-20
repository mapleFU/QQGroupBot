package subscribe

import (
	"github.com/mapleFU/QQBot/qqbot/data/group"
	"github.com/mapleFU/QQBot/qqbot/logger"
	"github.com/mapleFU/QQBot/qqbot/service"
	"time"

	"bytes"
	"golang.org/x/net/html"
	"log"
	"strings"

	"github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed"
)

type WeiboService struct {
	Subscribe
	ServiceUrl string
}

func removeScript(n *html.Node) {
	// if note is script tag
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "br" || n.Data == "body" || n.Data == "html") {
		n.Parent.RemoveChild(n)
		return // script tag is gone...
	}
	// traverse DOM
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		removeScript(c)
	}
}

func extractLink(linkText string) string {
	doc, err := html.Parse(strings.NewReader(linkText))
	if err != nil {
		log.Fatal(err)
	}
	removeScript(doc)
	buf := bytes.NewBuffer([]byte{})
	if err := html.Render(buf, doc); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func NewWeiboService(weiboUrl string) *WeiboService {
	return &WeiboService{
		ServiceUrl: weiboUrl,
		Subscribe:  Subscribe{service.NewBaseServicer()},
	}
}

func buildService(item *gofeed.Item, title string) group.StringRespMessage {
	// handle description
	Resp := group.StringRespMessage{
		Message:    title + " : \n" + strip.StripTags(item.Description) + "\n链接：" + item.Link,
		GroupID:    "",
		AutoEscape: true,
	}
	return Resp
}

// TODO: make clear how it end.
// TODO: consider using redis or other persistence techniques
func (self *WeiboService) Run() {
	fp := gofeed.NewParser()
	var lastNewest *gofeed.Item
	var title string

	// 考虑任务如何中止
	for {
		// 10 分钟一次
		feed, _ := fp.ParseURL(self.ServiceUrl)
		if feed.Items == nil {
			logger.SLogger.Info("Feed.Items is nil!")
		}
		if lastNewest == nil {
			lastNewest = feed.Items[0]
			logger.SLogger.Info("LastNewest Inited!")
			title = "[测试消息]" + lastNewest.Title
			Resp := buildService(lastNewest, title)
			logger.SLogger.Info("ready to send resp ", "resp", Resp)
			*self.OutChan <- Resp
		}

		for _, item := range feed.Items {
			if item == nil {
				logger.SLogger.Info("item is nil here")
			}

			if item.Title == lastNewest.Title {
				logger.SLogger.Info("item.Title == lastNewest.Title")
				// set latest news to the newer latest field
				lastNewest = feed.Items[0]
				break
			} else {

				title = item.Title
				Resp := buildService(item, title)
				logger.SLogger.Info("ready to send resp ", "resp", Resp)
				*self.OutChan <- Resp
			}
		}
		// TODO: debug set it 1
		time.Sleep(time.Minute * 5)
	}
}
