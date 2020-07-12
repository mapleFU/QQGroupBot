package main

import (
	"fmt"
	"net/http"

	"github.com/mapleFU/QQGroupBot/qqbot/data/group"
	"github.com/mapleFU/QQGroupBot/qqbot/service"
	"github.com/mapleFU/QQGroupBot/qqbot/service/query"
	"github.com/mapleFU/QQGroupBot/qqbot/service/subscribe"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const HttpRecvPort = 8085
const HttpUploadPort = 5700
const HttpManagerPort = 6324

const robotQQ = "3187545268"

const RssHubAddress = "http://rsshub:1200"

// testing
const RssHubTestingAddress = "http://localhost:1200"

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

	manager := service.NewManager("http://cqhttp:5700")

	weiboService := subscribe.NewWeiboService(RssHubAddress + "/weibo/user/5628238455")
	weiboService2 := subscribe.NewWeiboService(RssHubAddress + "/weibo/user/1195908387")

	//weiboService := subscribe.NewWeiboService(RssHubAddress + "/weibo/user/5628238455")
	//weiboService2 := subscribe.NewWeiboService(RssHubAddress + "/weibo/user/1195908387")

	imageSearch := query.NewSauceNaoQuery()
	hitoService := query.NewHitoService()

	// 写死的 map
	serviceMap := map[string]service.Servicer{
		"RSS Searcher":                   weiboService,
		"hitokoto provider":              hitoService,
		"trace.moe image search service": imageSearch,
		"weibo2":                         weiboService2,
	}

	revMap := make(map[service.Servicer]string)
	for k, v := range serviceMap {
		revMap[v] = k
	}

	manager.AddService(weiboService, "weibo")
	manager.AddService(imageSearch, "image-search")
	manager.AddService(hitoService, "hitokoto")
	manager.AddService(weiboService2, "weibo2")

	// TODO: set this config by
	manager.AddManagedGroups("117440534")
	//manager.AddManagedGroups("247437988")
	//manager.AddManagedGroups("702208467")

	r := gin.Default()
	r.Use(cors.Default())

	// ping of the robot manager, show the robot is alive
	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"pong": robotQQ,
		})
	})

	r.POST("", func(context *gin.Context) {
		var chatData group.ChatRequestData
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

	// 本来应该这样处理，但是时间有点不够了，先写后面的
	//// 处理 manager 的 groups
	//r.POST("/manager/service/:serviceName/group/:groupId", func(context *gin.Context) {
	//})
	//
	//r.GET("/manager/service/:serviceName/group", func(context *gin.Context) {
	//})
	//
	//r.DELETE("/manager/service/:serviceName/group/:groupId", func(context *gin.Context) {
	//})

	r.POST("/manager/group/:groupId", func(context *gin.Context) {
		groudId := context.Param("groupId")
		manager.AddManagedGroups(groudId)
		context.Writer.WriteHeader(http.StatusNoContent)
	})

	r.GET("/manager/group", func(context *gin.Context) {
		groups := manager.ListManagedGroups()

		context.JSON(200, map[string][]string{
			"groups": groups,
		})
	})

	r.DELETE("/manager/group/:groupId", func(context *gin.Context) {
		groudId := context.Param("groupId")
		manager.DeleteManagedGroups(groudId)
		context.Writer.WriteHeader(http.StatusNoContent)
	})

	// 添加、删除服务
	r.POST("/manager/service/:serviceName", func(context *gin.Context) {
		serviceName := context.Param("serviceName")

		servicer, ok := serviceMap[serviceName]
		if !ok {
			context.JSON(404, map[string]string{
				"error": fmt.Sprintf("Service '%s' not found", serviceName),
			})
			return
		}
		type Var struct {
			AddName string `json:"addName"`
		}
		var curAdd Var
		context.BindJSON(&curAdd)
		manager.AddService(servicer, curAdd.AddName)
		context.JSON(200, map[string]string{
			"serviceName": serviceName,
			"addName":     curAdd.AddName,
		})
	})

	r.DELETE("/manager/service/:addName", func(context *gin.Context) {
		serviceName := context.Param("addName")
		exists := manager.RemoveService(serviceName)
		if exists {
			context.Writer.WriteHeader(http.StatusNoContent)
		} else {
			context.Writer.WriteHeader(http.StatusNotFound)
		}
		return
	})

	// 获得服务的service状态
	r.GET("/manager/service", func(context *gin.Context) {
		sMap := manager.GetServiceMap()
		// retMap : appName --> serviceName
		retMap := make(map[string]string)
		for k, v := range *sMap {
			retMap[k] = revMap[v]
		}

		context.JSON(200, map[string]map[string]string{
			"groups": retMap,
		})
	})

	r.Run(fmt.Sprintf(":%d", HttpRecvPort))
}
