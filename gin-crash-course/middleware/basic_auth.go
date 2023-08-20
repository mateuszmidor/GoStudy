package middleware

import (
	"github.com/gin-gonic/gin"
)

var accounts gin.Accounts = gin.Accounts{
	"user":  "user",
	"admin": "pass",
}

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(accounts)
}
