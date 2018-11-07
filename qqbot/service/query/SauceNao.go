package query

import (
	"github.com/mapleFU/QQBot/qqbot/data/group"
	"strings"
	"github.com/mapleFU/QQBot/qqbot/image"
)

type SauceNaoQuery struct {
	QueryService
}

func (snq *SauceNaoQuery) IfAcceptMessage(Request *group.ChatRequestData) bool {
	var is_search = false
	for _, seg := range Request.Message {
		if seg.Type == "text" {
			strings.Contains(seg.Data.Text, "搜图")
			is_search = true
		}
		if seg.Type == "image" && is_search {
			return true
			//imageLink, ok := image.GetImage(&seg)
			//if !ok {
			//	return false
			//}
		}
	}
}

func (snq *SauceNaoQuery) Run() {
	for data := range snq.InChan {
		for _, seg := range data.Message {
			if seg.Type == "image" {
				go func() {
					imageLink, ok := image.GetImage(&seg)
					if !ok {
						return
					}

				}()

			}
		}
	}
}
