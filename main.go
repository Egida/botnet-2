package main

import (
	"botnet/configs"
	"botnet/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})
	configs.ConnectDB()
	routes.BotnetRoute(router)
	router.Run("localhost:4444")
}
