package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const bearerSchema = "BEARER "
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) < len(bearerSchema) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}
		tokenString := authHeader[len(bearerSchema):]
		token, err := service.NewJWTService().ValidateToken(tokenString)

		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[name]: ", claims["username"])
			log.Println("Claims[admin]: ", claims["admin"])
			log.Println("Claims[issuer]: ", claims["iss"])
			log.Println("Claims[issuedAt]: ", claims["iat"])
			log.Println("Claims[expiresAt]: ", claims["exp"])
		} else {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
		}
	}
}
