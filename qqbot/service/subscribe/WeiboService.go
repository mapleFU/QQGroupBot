package subscribe

import (
	"github.com/mmcdole/gofeed"
	
	"time"
	"github.com/mapleFU/QQBot/qqbot/data/group"
	"github.com/mapleFU/QQBot/qqbot/service"
	"fmt"
	"strings"
	"bytes"

	"log"
	"golang.org/x/net/html"
	"github.com/grokify/html-strip-tags-go"
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
		ServiceUrl:weiboUrl,
		Subscribe: Subscribe{service.NewBaseServicer()},
	}
}

func buildService(item *gofeed.Item, title string) group.StringRespMessage {
	// handle description

	Resp := group.StringRespMessage{
		Message: title + " : \n" + strip.StripTags(item.Description) + "\n链接：" + item.Link,
		GroupID:"",
		AutoEscape:true,
	}
	return Resp
}

// TODO: make clear how it end.
func (self *WeiboService) Run() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(self.ServiceUrl)
	if err != nil {
		panic(err.Error())

	}
	lastNewest := feed.Items[0]
	title := feed.Title

	*self.OutChan <- buildService(lastNewest, title)
	// 考虑任务如何中止
	for range self.InChan {
		// 10 分钟一次
		time.Sleep(time.Minute * 10)
		feed, _ := fp.ParseURL(self.ServiceUrl)
		if feed.Items == nil {
			fmt.Println("Feed.Items is nil!")
		}
		var curNewest *gofeed.Item
		curNewest = nil
		for _, item := range feed.Items {
			if item == nil {
				fmt.Println("item is nil here")
			}
			if curNewest == nil {
				curNewest = item
			}
			if item.Title == lastNewest.Title {
				lastNewest = curNewest
				break
			} else {
				Resp := buildService(item, title)
				*self.OutChan <- Resp
			}
		}
	}
}
