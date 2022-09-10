package router

import (
	"github.com/adamnasrudin03/merchant-service/app/controller"
	"github.com/gin-gonic/gin"
)

func AuthRouter(e *gin.Engine, authController controller.AuthController) {
	authRoutes := e.Group("/api/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}
}
