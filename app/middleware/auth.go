package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/adamnasrudin03/merchant-service/app/service"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
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

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			response := utils.APIResponse("no token found", http.StatusUnauthorized, "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := utils.APIResponse("token is not valid", http.StatusUnauthorized, "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
