package subscribe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/mapleFU/QQGroupBot/qqbot/data/group"
	"github.com/mapleFU/QQGroupBot/qqbot/logger"
	"github.com/mapleFU/QQGroupBot/qqbot/service"

	"github.com/PuerkitoBio/goquery"
	"github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed"
)

type WeiboService struct {
	Subscribe
	ServiceUrl string
}

//func removeScript(n *html.Node) {
//	// if note is script tag
//	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "br" || n.Data == "body" || n.Data == "html") {
//		n.Parent.RemoveChild(n)
//		return // script tag is gone...
//	}
//	// traverse DOM
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		removeScript(c)
//	}
//}
//
//func extractLink(linkText string) string {
//	doc, err := html.Parse(strings.NewReader(linkText))
//	if err != nil {
//		log.Fatal(err)
//	}
//	removeScript(doc)
//	buf := bytes.NewBuffer([]byte{})
//	if err := html.Render(buf, doc); err != nil {
//		log.Fatal(err)
//	}
//	return buf.String()
//}

// Get image from web link
// TODO: move this function to html utils
func ExtractImages(rawDoc string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawDoc))
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)
	doc.Find("img").Each(func(_ int, selection *goquery.Selection) {
		val, exists := selection.Attr("src")
		if exists {
			links = append(links, val)
		}
	})
	logger.SLogger.Info(links)
	return links, nil
}

func CqLink(s string) string {
	return fmt.Sprintf("[CQ:image,file=%s]", s)
}

func NewWeiboService(weiboUrl string) *WeiboService {
	return &WeiboService{
		ServiceUrl: weiboUrl,
		Subscribe:  Subscribe{service.NewBaseServicer()},
	}
}

func buildService(item *gofeed.Item) group.StringRespMessage {
	// handle description
	arr, err := ExtractImages(item.Description)
	if err != nil {
		logger.SLogger.Error(err)
	}
	links := make([]string, len(arr))
	for i, v := range arr {
		links[i] = CqLink(v)
	}

	logger.SLogger.Info("Send with title", "title", item.Title)
	Resp := group.StringRespMessage{
		Message: strip.StripTags(item.Description) + "\n链接：" + item.Link + "\n" + strings.Join(links, " "),
		GroupID: "",
	}
	return Resp
}

func (self *WeiboService) Run() {
	fp := gofeed.NewParser()
	var lastNewest *gofeed.Item
	var title string

	// 考虑任务如何中止
	for {
		func() {

			// 10 分钟一次
			feed, err := fp.ParseURL(self.ServiceUrl)
			if err != nil {
				logger.SLogger.Info("fp.ParseURL(self.ServiceUrl) failed with reason ", "error", err.Error())
				time.Sleep(time.Second * 5)
				return
			}

			defer time.Sleep(time.Minute * 5)

			if feed.Items == nil {
				logger.SLogger.Info("Feed.Items is nil!")
				return
			}
			sort.Slice(feed.Items, func(i, j int) bool {
				return feed.Items[i].PublishedParsed.After(*feed.Items[j].PublishedParsed)
			})

			// initialize lastNewest
			// 指向 feed.Items[0]
			if lastNewest == nil {
				lastNewest = feed.Items[0]

				logger.SLogger.Info("LastNewest Inited!")
				Resp := buildService(lastNewest)
				logger.SLogger.Info("ready to send resp ", "resp", Resp)
				*self.OutChan <- Resp
				return
			}

			for _, item := range feed.Items {
				// item 是 nil
				if item == nil {
					logger.SLogger.Info("item is nil here")
					continue
				}

				if item.Title == lastNewest.Title {
					logger.SLogger.Info("item.Title == lastNewest.Title")
					// set latest news to the newer latest field
					lastNewest = feed.Items[0]
					break
				} else {
					title = item.Title
					Resp := buildService(item)
					logger.SLogger.Info("ready to send resp ", "resp", Resp)
					*self.OutChan <- Resp
				}
			}
		}()
	}
}
