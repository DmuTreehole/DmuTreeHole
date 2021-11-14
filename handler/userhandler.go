package hander

//与用户相关的处理器函数
import (
	"github.com/gin-gonic/gin"
	UserModels "main/models/user"
	Utils "main/utils"
	"net/http"
)

//用户登录显示的页面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.html", gin.H{})
}

//检查用户登录
func LoginCheck(c *gin.Context) {
	var message string
	Username := c.PostForm("Username")
	Password := c.PostForm("Password")
	Id, CurrentPassword, exist := UserModels.Login(Username)
	if exist {
		ok := Utils.CobPassWord(Password, CurrentPassword)
		if ok {
			message = "Login Success"
		} else {
			message = "Wrong PassWord"
		}
	} else {
		message = "Account Not Exists"
	}
	UserModels.DoLog(Id, c.ClientIP(), message)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

//注册界面
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func RegisterCheck(c *gin.Context) {
	message := "Create Default"
	UserName := c.PostForm("Username")
	Password := c.PostForm("Password")
	UserEmail := c.PostForm("Email")
	userinfo := UserModels.User{
		Username: UserName,
		Password: Password,
		Email:    UserEmail,
	}
	Id, ok := UserModels.Register(userinfo)
	if ok {
		message = "Create Success"
	} else {
		Id = -1
		message = "Create Default"
	}
	UserModels.DoLog(Id, c.ClientIP(), message)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
