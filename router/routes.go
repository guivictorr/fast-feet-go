package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	initializePackageRoutes(v1)
	initializeRecipientsRoutes(v1)
	initializeUsersRoutes(v1)
}
