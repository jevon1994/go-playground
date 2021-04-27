package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-palyground/frame/web/gin/api/v1"
	"go-palyground/frame/web/gin/middleware"
)

func InitUserRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	admin := v1.Admin{}
	group := router.Group("/user").Use(middleware.Handler())
	{
		group.GET("/hi", admin.Add)
	}
	return group
}
