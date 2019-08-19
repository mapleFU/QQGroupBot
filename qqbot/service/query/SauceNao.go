package query

import (
	"github.com/mapleFU/QQBot/qqbot/data/group"
	"github.com/mapleFU/QQBot/qqbot/data/tracemoe/search"
	"github.com/mapleFU/QQBot/qqbot/image"
	"github.com/mapleFU/QQBot/qqbot/service"

	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bytes"
	"encoding/json"

	"github.com/polds/imgbase64"
)

type SauceNaoQuery struct {
	QueryService
}

func NewSauceNaoQuery() *SauceNaoQuery {
	return &SauceNaoQuery{
		QueryService{service.NewBaseServicer()},
	}
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
		}
	}
	return false
}

func (snq *SauceNaoQuery) Run() {
	for data := range snq.InChan {
		for _, seg := range data.Message {
			if seg.Type == "image" {
				go func() {
					imageLink, ok := image.GetImage(&seg)
					img := imgbase64.FromRemote(imageLink)
					//fmt.Println(img)
					if !ok {
						return
					}
					//resp, err := http.Get(imageLink)
					//
					//
					//fmt.Println("Debug: 走到这了3")
					//if err != nil {
					//	fmt.Println("http error")
					//	fmt.Println(err.Error())
					//	return
					//}
					//defer resp.Body.Close()
					//bytesData, err := ioutil.ReadAll(resp.Body)
					//fmt.Println("Debug: 走到这了2")
					//if err != nil {
					//	fmt.Println("read file error")
					//	fmt.Println(err.Error())
					//	return
					//}
					//fmt.Println("Debug: 走到这了1")
					//ext := filepath.Ext(seg.Data.File)[1:]
					//
					//mimeData := fmt.Sprintf("data:image/%s;base64,%s", ext, base64.StdEncoding.EncodeToString(bytesData))
					//
					bytesData, err := json.Marshal(map[string]string{
						"image": img,
					})
					////b64_out, err := os.Create("/home/user/log/base64.log")
					////outData := bufio.NewReader(b64_out)
					////w := bytes.NewReader(bytesData)
					////defer b64_out.Close()
					////io.Copy(b64_out, w)
					//
					//fmt.Println("Debug: 走到这了0")
					//if err != nil {
					//	fmt.Println(err.Error())
					//	return
					//}
					//
					//fmt.Println("Debug: 走到这了-1")
					respSearch, err := http.Post("https://trace.moe/api/search", "application/json", bytes.NewBuffer(bytesData))
					if err != nil {
						fmt.Println("http POST 请求异常")
						fmt.Println(err.Error())
						return
					}
					if respSearch.StatusCode != http.StatusOK {
						fmt.Println(fmt.Sprintf("http resp code %d", respSearch.StatusCode))
						return
					}

					defer respSearch.Body.Close()
					strData, err := ioutil.ReadAll(respSearch.Body)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					//fmt.Println(string(strData))

					var target search.SearchResult
					err = json.NewDecoder(bytes.NewBuffer(strData)).Decode(&target)
					if err != nil {
						fmt.Println("search.SearchResult 解码异常")
						fmt.Println(err.Error())
					}

					//if err = json.NewDecoder(respSearch.Body).Decode(&target); err != nil {
					//	fmt.Println("search.SearchResult 解码异常")
					//	// DEBUG
					//	out, err := os.Open("/home/user/log/http-response.log")
					//	w, err := ioutil.ReadAll(respSearch.Body)
					//	if err != nil {
					//		// panic?
					//		fmt.Println(err.Error())
					//
					//		fmt.Println(string(w))
					//	}
					//	fmt.Println(string(w))
					//	defer out.Close()
					//	io.Copy(out, respSearch.Body)
					//	//fmt.Println(err.Error())
					//	return
					//}
					//jsonData, err := json.Marshal(target)
					//if err != nil {
					//	fmt.Println("search.SearchResult 解码异常")
					//	fmt.Println(err.Error())
					//	return
					//}

					Resp := group.StringRespMessage{
						Message:    target.String(),
						GroupID:    "",
						AutoEscape: true,
					}
					fmt.Println("Ready to send")
					*snq.OutChan <- Resp
					fmt.Println("Send done")
				}()

			}
		}
	}
}
