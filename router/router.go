package router

import (
	"github.com/gin-gonic/gin"
	handlers "main/handler"
)

func Router() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	api := r.Group("/api")
	{
		//接口层的代码书写在这
		api.GET("user/login", handlers.Login)
		api.POST("user/login", handlers.LoginCheck)
		api.GET("user/register", handlers.Register)
		api.POST("user/register", handlers.RegisterCheck)
	}
	r.Run(":80")
}
