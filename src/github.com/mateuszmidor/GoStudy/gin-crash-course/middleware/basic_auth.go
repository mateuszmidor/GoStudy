package middleware

import (
	"github.com/gin-gonic/gin"
)

var accounts gin.Accounts = gin.Accounts{
	"user":  "user",
	"admin": "admin",
}

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(accounts)
}
