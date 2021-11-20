package router

import (
	handlers "main/handler"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFile("View/Photo.jpg", "View/Photo.jpg")
	api := r.Group("/api")
	{
		//接口层的代码书写在这
		api.GET("user/login", handlers.Login)
		api.POST("user/login", handlers.LoginCheck)
		api.GET("user/register", handlers.Register)
		api.POST("user/register", handlers.RegisterCheck)
		api.GET("post/getAllPost", handlers.GetAllPost)
		api.POST("post/createOnePost", handlers.CreateOnePost)
	}
	r.Run(":8081")
}
