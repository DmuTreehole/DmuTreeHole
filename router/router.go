package router

import (
	docs "main/docs"
	handlers "main/handler"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.LoadHTMLGlob("templates/*.html")
	//r.StaticFile("View/Photo.jpg", "View/Photo.jpg")
	api := r.Group("/api")
	{
		api.GET("/question", handlers.GetQuestion)
		api.GET("/test", handlers.Test)
		//接口层的代码书写在这
		user := api.Group("/user")
		{
			user.GET("login", handlers.Login)
			user.GET("register", handlers.Register)
			user.GET("showbanedlist", handlers.ShowBannedList)
			user.POST("login", handlers.LoginCheck)
			user.POST("register", handlers.RegisterCheck)
			user.POST("banusers", handlers.BanUsers)
			user.POST("getusername", handlers.GetUserName)
			user.POST("createuserprofile", handlers.CreateUserProfile)
			user.POST("showuserprofile", handlers.ShowUserProfile)
			user.POST("changeuserprofile", handlers.ChangeUserProfile)
		}
		post := api.Group("/post")
		{
			post.GET("getallpost/:page", handlers.GetAllPost)
			post.POST("createonepost", handlers.CreateOnePost)
			post.GET("deleteonepost/:id", handlers.DeleteOnePost)
			post.POST("getpostbyId", handlers.GetPostById)
			post.POST("search", handlers.SearchPostByContent)
		}
		comment := api.Group("/comment")
		{
			comment.GET("getallcomment/:id", handlers.GetAllComment)
			comment.POST("createonecomment", handlers.CreateOneComment)
			comment.GET("deletecomment/:id", handlers.DeleteOneComment)
		}
	}
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8081")
}
