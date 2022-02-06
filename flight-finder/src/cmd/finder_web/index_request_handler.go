package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newIndexHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "map.html", nil)
	}
}
