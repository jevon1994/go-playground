package inject

import "github.com/gin-gonic/gin"

//Index 注入IStartService
type Index struct {
	Service IStartService `inject:""`
}

//GetName 调用IStartService的Say方法
func (i *Index) GetName(ctx *gin.Context) {
	var message = ctx.Param("msg")
	ctx.JSON(200, i.Service.Say(message))
}
