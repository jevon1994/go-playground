package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Printf("%s \n", "po something!")
	}

}
