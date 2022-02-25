package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "main/docs"
	handlers "main/handler"
	"main/middleware"
	"net/http"
)

func Router() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("session", store))
	r.LoadHTMLGlob("templates/*.html")
	//r.StaticFS("/Icon",http.Dir("./Icon"))
	api := r.Group("/api")
	{
		api.StaticFS("/Icon", http.Dir("./Icon"))
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
			user.POST("getusericon", handlers.ShowUserIcon)
			user.POST("getusername", handlers.GetUserName)
			api.Group("", middleware.JWT())
			{
				user.POST("createuserprofile", handlers.CreateUserProfile)
				user.POST("showuserprofile", handlers.ShowUserProfile)
				user.POST("changeuserprofile", handlers.ChangeUserProfile)
				user.POST("uploadicon", handlers.UploadUserIcon)
				user.POST("sendcode", handlers.EmailAuth)
				user.POST("checkcode", handlers.EmailCheck)
			}
		}
		post := api.Group("/post", middleware.JWT())
		{
			post.GET("getallpost/:page", handlers.GetAllPost)
			post.POST("createonepost", handlers.CreateOnePost)
			post.GET("deleteonepost/:id", handlers.DeleteOnePost)
			post.POST("getpostbyId", handlers.GetPostById)
			post.POST("search", handlers.SearchPostByContent)
			post.POST("createfeedback", handlers.CreateOneFeedback)
		}
		comment := api.Group("/comment", middleware.JWT())
		{
			comment.POST("getallcomment", handlers.GetAllComment)
			comment.POST("createonecomment", handlers.CreateOneComment)
			comment.GET("deletecomment/:id", handlers.DeleteOneComment)
		}
	}
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8081")
}

/*
┏┛┻━━━━━┛┻┓
┃｜｜｜｜｜|┃
┃     ━   ┃
┃  ┳┛  ┗┳ ┃
┃　　　　　 ┃
┃　　　┻　　┃
┃　　　　　 ┃
┗━┓　　　┏━┛
  ┃　代　┃
  ┃　码　┃
  ┃　之　┃
  ┃　神　┃
  ┃　　　┗━━━━━┓
  ┃  No Error ┣┓
  ┃  No Bug   ┃
  ┗┓┓┏━┳┓┏━━━━┛
   ┃┫┫ ┃┫┫
   ┗┻┛ ┗┻┛
*/
