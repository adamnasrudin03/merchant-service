package main

import (
	"fmt"
	"net/http"

	"github.com/adamnasrudin03/merchant-service/app/configs"
	"github.com/adamnasrudin03/merchant-service/app/controller"
	"github.com/adamnasrudin03/merchant-service/app/middleware"
	"github.com/adamnasrudin03/merchant-service/app/repository"
	routers "github.com/adamnasrudin03/merchant-service/app/router"
	"github.com/adamnasrudin03/merchant-service/app/service"
	"github.com/adamnasrudin03/merchant-service/pkg/gormdb"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                    *gorm.DB                         = gormdb.SetupDbConnection()
	userRepository        repository.UserRepository        = repository.NewUserRepository(db)
	transactionRepository repository.TransactionRepository = repository.NewTransactionRepository(db)

	jwtService         service.JWTService         = service.NewJWTService()
	authService        service.AuthService        = service.NewAuthService(userRepository)
	transactionService service.TransactionService = service.NewTransactionService(transactionRepository)

	transactionController controller.TransactionController = controller.NewTransactionController(transactionService, jwtService)
	authController        controller.AuthController        = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer gormdb.CloseDbConnection(db)

	authMiddleware := middleware.NewAuthMiddleware(jwtService)

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
