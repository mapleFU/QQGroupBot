package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
)

const HttpRecvPort = 8085
const HttpUploadPort = 5700

func main()  {
	r := gin.Default()

	r.POST("", func(context *gin.Context) {
		fmt.Println(json.MarshalIndent(context.Request.Body, "", "    "))
	})

	r.Run(fmt.Sprintf(":%d", HttpRecvPort))
}
