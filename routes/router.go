package routes

import (
	"botnet/controllers"

	"github.com/gin-gonic/gin"
)

func BotnetRoute(router *gin.Engine) {
	router.GET("/ping", controllers.Ping())
}
