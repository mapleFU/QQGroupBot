package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mapleFU/QQBot/qqbot/data/group"

	"github.com/mapleFU/QQBot/qqbot/service"
	"github.com/mapleFU/QQBot/qqbot/service/subscribe"
	"github.com/mapleFU/QQBot/qqbot/service/query"
)

const HttpRecvPort = 8085
const HttpUploadPort = 5700
const HttpManagerPort = 6324

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

func runCQHttpServer(manager *service.Manager)  {

}

func main() {

	manager := service.NewManager("http://cqhttp:5700")
	weiboService := subscribe.NewWeiboService("http://101.132.121.41:1200/weibo/user/5628238455")
	imageSearch := query.NewSauceNaoQuery()
	hitoService := query.NewHitoService()

	manager.AddService(weiboService, "weibo")
	manager.AddService(imageSearch, "image-search")
	manager.AddService(hitoService, "hitokoto")

	//manager.AddManagedGroups("117440534")
	manager.AddManagedGroups("702208467")

	r := gin.Default()
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

	// 处理 manager
	r.POST("/manager/group/add/:groupId", func(context *gin.Context) {

	})

	r.POST("/manager/group/delete/:groupId", func(context *gin.Context) {

	})

	r.POST("/manager/service/start/:serviceName", func(context *gin.Context) {
		
	})
	
	r.POST("/manager/service/stop/:serviceName", func(context *gin.Context) {
		
	})

	r.Run(fmt.Sprintf(":%d", HttpRecvPort))
	go runCQHttpServer(manager)
	
}
