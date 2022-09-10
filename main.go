package main

import (
	"fmt"
	"net/http"

	"github.com/adamnasrudin03/merchant-service/app/configs"
	"github.com/adamnasrudin03/merchant-service/pkg/gormdb"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = gormdb.SetupDbConnection()
)

func main() {
	defer gormdb.CloseDbConnection(db)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		response := utils.APIResponse("Welcome my application", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, response)
	})

	router.NoRoute(func(c *gin.Context) {
		response := utils.APIResponse("Page not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
	})

	config := configs.GetInstance()
	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
