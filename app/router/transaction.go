package router

import (
	"github.com/adamnasrudin03/merchant-service/app/controller"
	"github.com/gin-gonic/gin"
)

func TransactionRouter(e *gin.Engine, transactionController controller.TransactionController) {
	transactionRoutes := e.Group("/api/v1/transaction")
	{
		transactionRoutes.GET("/:param/:id", transactionController.ListTransactionReport)
	}
}
