package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adamnasrudin03/merchant-service/app/dto"
	"github.com/adamnasrudin03/merchant-service/app/service"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

//TransactionController interface is a contract what this controller can do
type TransactionController interface {
	ListTransactionReport(ctx *gin.Context)
}

type transactionController struct {
	TransactionService service.TransactionService
	jwtService         service.JWTService
}

//NewTransactionController creates a new instance of TransactionController
func NewTransactionController(TransactionService service.TransactionService, jwtService service.JWTService) TransactionController {
	return &transactionController{
		TransactionService: TransactionService,
		jwtService:         jwtService,
	}
}

func (c *transactionController) ListTransactionReport(ctx *gin.Context) {
	var (
		paramPage  int = 0
		paramLimit int = 10
	)

	if ctx.Param("merchantID") == "" {
		response := utils.APIResponse("param ID Merchant not found", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	merchantID, err := strconv.Atoi(ctx.Param("merchantID"))
	if err != nil {
		response := utils.APIResponse("param ID Merchant not valid", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	paramPage, err = strconv.Atoi(ctx.Query("page"))
	if err != nil {
		response := utils.APIResponse("query param page not found or invalid", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if paramLimit, err = strconv.Atoi(ctx.Query("limit")); err != nil {
		response := utils.APIResponse("query param limit not found or invalid", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	startAt := ctx.Query("start_at")
	_, err = time.Parse("2006-01-02", startAt)
	if err != nil {
		response := utils.APIResponse("query param start_at not found or invalid", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	endAt := ctx.Query("end_at")
	_, err = time.Parse("2006-01-02", endAt)
	if err != nil {
		response := utils.APIResponse("query param end_at not found or invalid", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	param := dto.ParamTransaction{
		Page:       paramPage,
		Limit:      paramLimit,
		MerchantID: merchantID,
		StartAt:    startAt,
		EndAt:      endAt,
	}

	transactions, err := c.TransactionService.GetIncomeReport(param)
	if err != nil {
		response := utils.APIResponse("Error to get transaction", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("List of transaction", http.StatusOK, "success", transactions)
	ctx.JSON(http.StatusOK, response)
}
