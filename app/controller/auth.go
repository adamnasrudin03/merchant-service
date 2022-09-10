package controller

import (
	"net/http"
	"strconv"

	"github.com/adamnasrudin03/merchant-service/app/dto"
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"github.com/adamnasrudin03/merchant-service/app/service"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var input dto.LoginRequest

	//Validation input user
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("failed to process request", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	authResult := c.authService.VerifyCredential(input.Username, input.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatInt(v.ID, 10))
		v.Token = generatedToken

		response := utils.APIResponse("login Success", http.StatusOK, "success", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := utils.APIResponse("please check again your username or password, invalid credential", http.StatusUnauthorized, "error", nil)
	ctx.JSON(http.StatusUnauthorized, response)
}
