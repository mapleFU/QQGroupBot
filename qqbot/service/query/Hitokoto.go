package query

import (
	"encoding/json"
	"fmt"
	"github.com/mapleFU/QQBot/qqbot/data/group"
	hitokoto2 "github.com/mapleFU/QQBot/qqbot/data/hitokoto"
	"github.com/mapleFU/QQBot/qqbot/service"
	"net/http"
	"strings"
)

type HitokotoService struct {
	QueryService
}

func NewHitoService() *HitokotoService {
	return &HitokotoService{
		QueryService{service.NewBaseServicer()},
	}
}

func (hqs *HitokotoService) IfAcceptMessage(Request *group.ChatRequestData) bool {
	for _, data := range Request.Message {
		if data.Type == "text" {
			if strings.Contains(data.Data.Text, "二次元名言") {
				return true
			}
		}
	}
	return false
}

func (hqs *HitokotoService) Run() {
	for range hqs.InChan {
		go func() {
			resp, err := http.Get("https://v1.hitokoto.cn")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			defer resp.Body.Close()
			var hitokoto hitokoto2.HitoResp
			if err = json.NewDecoder(resp.Body).Decode(&hitokoto); err != nil {
				fmt.Println(err.Error())
				return
			}

			*hqs.OutChan <- group.StringRespMessage{
				Message:    hitokoto.String(),
				GroupID:    "",
				AutoEscape: true,
			}
		}()
	}
}
