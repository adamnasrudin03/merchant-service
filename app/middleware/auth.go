package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/adamnasrudin03/merchant-service/app/service"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareController interface {
	AuthorizeJWT()
}
type authMiddleware struct {
	jwtService  service.JWTService
	authService service.AuthService
}

func NewAuthMiddleware(jwtService service.JWTService, authService service.AuthService) *authMiddleware {
	return &authMiddleware{
		jwtService:  jwtService,
		authService: authService,
	}
}

//AuthorizeJWT validates the token user given, return 401 if not valid
func (auth *authMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := utils.APIResponse("inappropriate procedure", http.StatusUnauthorized, "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		if tokenString == "" {
			response := utils.APIResponse("no token found", http.StatusUnauthorized, "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := auth.jwtService.ValidateToken(tokenString)
		if err != nil {
			response := utils.APIResponse("invalid token", http.StatusUnauthorized, "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		playload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := utils.APIResponse("token is not valid", http.StatusUnauthorized, "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		str := fmt.Sprintf("%v", playload["user_id"])
		userID, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			fmt.Printf("%d of type %T", userID, userID)
		}

		user, err := auth.authService.GetUserByID(userID)
		if err != nil {
			response := utils.APIResponse("please check again if the username and password are registered", http.StatusUnauthorized, "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//set context isinya user
		c.Set("currentUser", user)

	}
}
