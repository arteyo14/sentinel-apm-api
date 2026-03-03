package routes

import (
	"sentinel-apm-api/internal/modules/auth"
	"sentinel-apm-api/internal/modules/role"
	"sentinel-apm-api/internal/modules/user"
	"sentinel-apm-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api/v1")

	auth.RouteRegister(api, db)

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		user.RouteRegister(protected, db)
		role.RouteRegister(protected, db)
	}
}
