package main

import (
	"botnet/configs"
	"botnet/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})
	configs.ConnectDB()
	routes.BotnetRoute(router)
	router.Run(os.Getenv("PORT"))

}
