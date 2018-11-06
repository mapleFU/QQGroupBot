package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mapleFU/QQBot/qqbot/data/group"

	"github.com/mapleFU/QQBot/qqbot/service"
	"github.com/mapleFU/QQBot/qqbot/service/subscribe"
)

const HttpRecvPort = 8085
const HttpUploadPort = 5700
const robotQQ = "3187545268"

func checkAtData(chatData *group.ChatRequestData, robotQQ string) bool {
	ok := false
	for _, detailMessage := range chatData.Message {
		if detailMessage.Type == "at" {
			if detailMessage.Data.QQ == robotQQ {
				ok = true
			}
		}
	}

	return ok
}

func main() {
	r := gin.Default()
	manager := service.NewManager("http://cqhttp:5700")
	weiboService := subscribe.NewWeiboService("https://rsshub.app/weibo/user/5628238455")

	manager.AddService(weiboService, "weibo")

	manager.AddManagedGroups("117440534")

	r.POST("", func(context *gin.Context) {
		var chatData group.ChatRequestData
		//var mydata interface{}
		if err := context.ShouldBindJSON(&chatData); err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// add middle ware
		if !checkAtData(&chatData, robotQQ) {
			return
		}
		fmt.Println("ok, our robot is at, let's call our manager")
	//	来定义我们的服务了
		manager.RecvRequest(&chatData)
	})

	r.Run(fmt.Sprintf(":%d", HttpRecvPort))
}
