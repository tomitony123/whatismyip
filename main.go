package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()

	r.GET("/whatismyip", whatismyipController)
}

func whatismyipController(c *gin.Context) {
	var output = map[string]interface{}{}

	output["ip"] = c.RemoteIP()

	c.JSON(200, output)
}