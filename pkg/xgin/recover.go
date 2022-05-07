package xgin

import (
	"github.com/gin-gonic/gin"
)

func XRecoveryHandler() gin.RecoveryFunc {
	return func(c *gin.Context, err interface{}) {
		c.JSON(200, result(nil, "服务器异常", 500))
	}
}
