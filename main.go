package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/config"
	"github.com/guivictorr/fast-feet-go/routes"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	router := SetupAppRouter()
	logger.Info(router.Run(":" + "3000"))
}

func SetupAppRouter() *gin.Engine {
	err := config.Init()
	if err != nil {
		logger.Errorf("error initializing configuration: %v", err)
		return nil
	}
	db := config.GetSQLite()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	api := router.Group("api/v1")

	users := api.Group("/users")
	routes.InitUserRoutes(db, users)
	auth := api.Group("/auth")
	routes.InitAuthRoutes(db, auth)

	return router
}
