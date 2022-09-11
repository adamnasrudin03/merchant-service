package main

import (
	"fmt"
	"net/http"

	"github.com/adamnasrudin03/merchant-service/app"
	"github.com/adamnasrudin03/merchant-service/app/configs"
	"github.com/adamnasrudin03/merchant-service/app/controller"
	"github.com/adamnasrudin03/merchant-service/app/middleware"
	routers "github.com/adamnasrudin03/merchant-service/app/router"
	"github.com/adamnasrudin03/merchant-service/app/service"
	"github.com/adamnasrudin03/merchant-service/pkg/gormdb"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB           = gormdb.SetupDbConnection()
	repo                          = app.WiringRepository(db)
	services                      = app.WiringService(repo)
	jwtService service.JWTService = service.NewJWTService()

	transactionController controller.TransactionController = controller.NewTransactionController(services, jwtService)
	authController        controller.AuthController        = controller.NewAuthController(services.Auth, jwtService)
	authMiddleware                                         = middleware.NewAuthMiddleware(jwtService, services.Auth)
)

func main() {
	defer gormdb.CloseDbConnection(db)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		response := utils.APIResponse("Welcome my application", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, response)
	})

	routers.AuthRouter(router, authController)
	router.Use(authMiddleware.AuthorizeJWT())
	routers.TransactionRouter(router, transactionController)

	router.NoRoute(func(c *gin.Context) {
		response := utils.APIResponse("Page not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
	})

	config := configs.GetInstance()
	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
