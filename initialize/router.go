package initialize

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
)

func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()

	return Router
}
