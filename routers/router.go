package routers

import (
	"game_sys/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", controller.GetInfo)
	r.GET("/test", controller.UseUtil)

	return r
}
