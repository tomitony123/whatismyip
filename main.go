package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	port        = flag.Int("p", 9092, "set server port")
)

func main()  {
	r := gin.Default()

	r.GET("/whatismyip", whatismyipController)

	var listenPort = fmt.Sprintf(":%d", *port)
	r.Run(listenPort)
}

func whatismyipController(c *gin.Context) {
	var output = map[string]interface{}{}

	output["ip"] = c.RemoteIP()

	c.JSON(200, output)
}