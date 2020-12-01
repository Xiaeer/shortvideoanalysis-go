package routers

import (
	"webtemp/shortvideoanalysis/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine) {
	// 设置根路由
	setCommonRouter(router)
	// 设置短视频路由
	setShortVideoRouter(router)
}

func setCommonRouter(router *gin.Engine) {
	// 设置favicon（浏览器默认请求favicon）
	router.GET("/favicon.ico", controllers.Favicon)
}
