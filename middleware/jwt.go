package middleware

import (
	"github.com/gin-gonic/gin"
	"main/utils"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg string
		token := c.Query("token")
		if token == "" {
			msg = `NotLogin`
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				msg = `CheckFail`
			} else if time.Now().Unix() > claims.ExpiresAt {
				msg = `TimeOut`
			}
		}
		if msg != "" {
			c.JSON(http.StatusOK, gin.H{
				"message": msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
