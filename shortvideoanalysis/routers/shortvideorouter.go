package routers

import (
	"webtemp/shortvideoanalysis/controllers"

	"github.com/gin-gonic/gin"
)

// setShortVideoRouter ...
func setShortVideoRouter(router *gin.Engine) {

	// 设置短视频路由
	router.GET("parseshortvideo", controllers.ParseShortVideo)

	// 设置短视频链接解析接口路由
	router.POST("parseshortvideobyurl", controllers.ParseShortVideoByURL)

}
