package v1

import "github.com/gin-gonic/gin"

type Menu struct {
	Name string
}

func (m *Menu) Add(ctx *gin.Context) string {
	m.Name = "menu"
	return m.Name
}
