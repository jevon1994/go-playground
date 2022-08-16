package router

import (
	"github.com/gin-gonic/gin"
	"go-palyground/frame/web/gin/middleware"
	v1 "go-palyground/frame/web/rk-gin/api/v1"
)

func InitUserRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	admin := v1.Admin{}
	group := router.Group("/user").Use(middleware.Handler())
	{
		group.GET("/hi", admin.Add)
	}
	return group
}
