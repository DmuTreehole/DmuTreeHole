package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	handlers "main/handler"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFile("View/Photo.jpg", "View/Photo.jpg")
	api := r.Group("/api")
	{
		//接口层的代码书写在这
		user := api.Group("/user")
		{
			user.GET("login", handlers.Login)
			user.POST("login", handlers.LoginCheck)
			user.GET("register", handlers.Register)
			user.POST("register", handlers.RegisterCheck)
		}
		post := api.Group("/api")
		{
			post.GET("getAllPost", handlers.GetAllPost)
			post.POST("createOnePost", handlers.CreateOnePost)
		}
	}
	r.Run(":8081")
}
