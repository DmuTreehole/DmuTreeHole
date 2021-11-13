package Controller

import (
	"github.com/gin-gonic/gin"
	models "main/Models"
	"net/http"
)

func Router() {
	models.OpenDataBase()
	router := gin.Default()
	router.LoadHTMLGlob("/home/gopath/**/**/*.html")
	router.StaticFile("View/Photo.jpg", "View/Photo.jpg")
	router.GET("/", Default)
	router.GET("/login", Login)
	router.POST("/login", LoginCheck)
	router.GET("/register", Register)
	router.POST("/register", RegisterCheck)
	router.Run(":80")
}
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.html", gin.H{})
}
func LoginCheck(c *gin.Context) {
	var message string
	Username := c.PostForm("Username")
	Password := c.PostForm("Password")
	Id, CurrentPassword, exist := models.Login(Username)
	if exist {
		ok := models.CobPassWord(Password, CurrentPassword)
		if ok {
			message = "Login Success"
		} else {
			message = "Wrong PassWord"
		}
	} else {
		message = "Account Not Exists"
	}
	models.DoLog(Id, c.ClientIP(), message)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})

}
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func RegisterCheck(c *gin.Context) {
	var message string = "Create Default"
	UserName := c.PostForm("Username")
	Password := c.PostForm("Password")
	UserEmail := c.PostForm("Email")
	userinfo := models.User{
		UserName:     UserName,
		UserPassword: Password,
		UserEmail:    UserEmail,
	}
	Id, ok := models.Register(userinfo)
	if ok {
		message = "Create Success"
	} else {
		Id = -1
		message = "Create Default"
	}
	models.DoLog(Id, c.ClientIP(), message)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
func Default(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
