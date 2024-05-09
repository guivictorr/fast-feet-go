package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/handler"
)

func initializeRouter(r *gin.Engine) {
	handler.InitializeHandler()
	v1 := r.Group("/api/v1")
	initializePackageRoutes(v1)
	initializeRecipientsRoutes(v1)
	initializeUsersRoutes(v1)
}
