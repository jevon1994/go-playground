package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User interface {
	Add(ctx *gin.Context)
}

type Admin struct {
	Name string
}
type Response struct {
	Code int
	Data interface{}
	Msg  string
}

func (n *Admin) Add(ctx *gin.Context) {
	n.Name = "admin"
	ctx.JSON(http.StatusOK, Response{
		200,
		"admin",
		"success",
	})
}
