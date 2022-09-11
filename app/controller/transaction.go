package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adamnasrudin03/merchant-service/app/dto"
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"github.com/adamnasrudin03/merchant-service/app/service"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

//TransactionController interface is a contract what this controller can do
type TransactionController interface {
	ListTransactionReport(ctx *gin.Context)
}

type transactionController struct {
	Service    *service.Services
	jwtService service.JWTService
}

//NewTransactionController creates a new instance of TransactionController
func NewTransactionController(srv *service.Services, jwtService service.JWTService) TransactionController {
	return &transactionController{
		Service:    srv,
		jwtService: jwtService,
	}
}

func (c *transactionController) ListTransactionReport(ctx *gin.Context) {
	var (
		paramPage  int = 0
		paramLimit int = 10
		isMerchant     = false
	)

	paramID := ctx.Param("id")

	switch ctx.Param("param") {
	case "merchant":
		isMerchant = true
	case "outlet":
		isMerchant = false
	default:
		response := utils.APIResponse("page not found", http.StatusNotFound, "error", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if paramID == "" {
		response := utils.APIResponse("param ID not found", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(paramID)
	if err != nil {
		response := utils.APIResponse("param ID not valid", http.StatusBadRequest, "error", nil)
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

	currentUser := ctx.MustGet("currentUser").(entity.User)

	param := dto.ParamTransaction{
		Page:    paramPage,
		Limit:   paramLimit,
		StartAt: startAt,
		EndAt:   endAt,
	}

	if isMerchant {
		param.MerchantID = ID
		merchant, _ := c.Service.Merchant.GetMerchantByID(int64(ID))
		if merchant.UserID != currentUser.ID {
			response := utils.APIResponse("Access to that resource is forbidden", http.StatusForbidden, "error", nil)
			ctx.JSON(http.StatusForbidden, response)
			return
		}
	} else {
		param.OutletID = ID
		outlet, _ := c.Service.Outlet.GetOutletByID(int64(ID))
		if outlet.UserID != currentUser.ID {
			response := utils.APIResponse("Access to that resource is forbidden", http.StatusForbidden, "error", nil)
			ctx.JSON(http.StatusForbidden, response)
			return
		}
	}

	transactions, err := c.Service.Transaction.GetIncomeReport(param)
	if err != nil {
		response := utils.APIResponse("Error to get transaction", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("List of transaction", http.StatusOK, "success", transactions)
	ctx.JSON(http.StatusOK, response)
}
