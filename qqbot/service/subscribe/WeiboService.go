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

func (self *WeiboService) Run() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(self.ServiceUrl)
	if err != nil {
		panic(err.Error())

	}
	newest := feed.Items[0]
	title := feed.Title
	if self.OutChan == nil {
		fmt.Println("Bug. self.Outchan is nil")
	} else if newest == nil {
		fmt.Println("Bug, Newest is nil")
	} else {
		fmt.Println("Send News")
		//*self.OutChan <- buildService(newest, title)
	}
	fmt.Println("Send News Done")
	// 考虑任务如何中止
	for true  {
		// 10 分钟一次
		time.Sleep(time.Minute * 10)
		feed, _ := fp.ParseURL(self.ServiceUrl)
		if feed.Items == nil {
			fmt.Println("Feed.Items is nil!")
		}

		for _, item := range feed.Items {
			if item == nil {
				fmt.Println("item is nil here")
			}
			if item.Title == newest.Title {
				newest = item
				break
			} else {
				Resp := buildService(item, title)
				*self.OutChan <- Resp
			}
		}
	}
}
