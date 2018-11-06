package main

import (
	"fmt"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/mapleFU/QQBot/qqbot/data/group"
	"net/http"
)

const HttpRecvPort = 8085
const HttpUploadPort = 5700

func main()  {
	r := gin.Default()

	r.POST("", func(context *gin.Context) {
		var chatData group.ChatRequestData
		if err := context.ShouldBindJSON(&chatData); err != nil {
			fmt.Println("Error, bad request")
			fmt.Println(err.Error())
			fmt.Println(json.MarshalIndent(context.Request.Body, "", "\t"))
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


	})

	r.Run(fmt.Sprintf(":%d", HttpRecvPort))
}
