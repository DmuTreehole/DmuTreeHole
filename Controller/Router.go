package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	models "main/Models"
	"net/http"
)

func Router() {
	models.OpenDataBase()
	router := gin.Default()
	router.LoadHTMLGlob("/home/gopath/**/**/*.html")
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
	Username := c.PostForm("Username")
	Password := c.PostForm("Password")
	CurrentPassword, _ := models.Login(Username)
	fmt.Println(Password, CurrentPassword)
	if Password == CurrentPassword {
		c.JSON(http.StatusOK, gin.H{
			"message": "true",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码错误",
		})
	}
}
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func RegisterCheck(c *gin.Context) {
	UserName := c.PostForm("Username")
	Password := c.PostForm("Password")
	UserPhone := c.PostForm("Phone")
	UserEmail := c.PostForm("Email")
	userinfo := models.User{
		UserName:     UserName,
		UserPassword: Password,
		UserEmail:    UserEmail,
		UserPhone:    UserPhone,
	}
	fmt.Println(userinfo)
	if models.Register(userinfo) {
		c.JSON(http.StatusOK, "True")
	} else {
		c.JSON(http.StatusOK, "False")
	}
}
func Default(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
