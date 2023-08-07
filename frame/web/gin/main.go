package main

import (
	"github.com/gin-gonic/gin"
	"go-palyground/frame/web/gin/router"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.InitUserRouter(r.Group("rk-web"))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
