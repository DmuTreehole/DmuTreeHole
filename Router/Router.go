package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello,World!")
	})
	router.Run(":80")
}
