package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	handlers "main/handler"
)

func Router() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.LoadHTMLGlob("templates/*.html")
	//r.StaticFile("View/Photo.jpg", "View/Photo.jpg")
	api := r.Group("/api")
	{
		api.POST("/question", handlers.Getquestion)
		//接口层的代码书写在这
		user := api.Group("/user")
		{
			user.GET("login", handlers.Login)
			user.GET("register", handlers.Register)
			user.GET("showbanedlist",handlers.ShowBannedList)
			user.POST("login", handlers.LoginCheck)
			user.POST("register", handlers.RegisterCheck)
			user.POST("banusers",handlers.BanUsers)
		}
		post := api.Group("/post")
		{
			post.POST("getallpost/:num", handlers.GetAllPost)
			post.POST("createonepost", handlers.CreateOnePost)
			post.POST("deleteonepost/:id", handlers.DeleteOnePost)
		}
		comment := api.Group("/comment")
		{
			comment.POST("getallcomment/:id", handlers.GetAllComment)
			comment.POST("createonecomment", handlers.CreateOneComment)
			comment.POST("deletecomment/:id", handlers.DeleteOneComment)
		}
		reply := api.Group("/reply")
		{
			reply.POST("getreply/:page", handlers.GetReply)
			reply.POST("deletereply", handlers.DeleteReply)
			reply.POST("createreply", handlers.CreateReply)
		}
	}
	r.Run(":8081")
}
