package handler

//与用户相关的处理器函数
import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	DB "main/db"
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
	var user = UserModels.User{}
	//Username := c.PostForm("Username")
	//Password := c.PostForm("Password")
	c.ShouldBind(user)
	Id, CurrentPassword, exist := UserModels.Login(user.Username, DB.Dbs)
	if exist {
		ok := Utils.CobPassWord(user.Password, CurrentPassword)
		if ok {
			message = "Login Success"
			session := sessions.Default(c)
			session.Options(sessions.Options{
				MaxAge: 60 * 60, //1h
			})
			session.Set("username", user.Username)
			session.Save()
		} else {
			message = "Wrong PassWord"
		}
	} else {
		message = "Account Not Exists"
	}
	UserModels.DoLog(Id, c.ClientIP(), message, DB.Dbs)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

//注册界面
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

//检查注册
func RegisterCheck(c *gin.Context) {
	var message string = "Create Default"
	var userinfo = UserModels.User{}
	//UserName := c.PostForm("Username")
	//Password := c.PostForm("Password")
	//UserEmail := c.PostForm("Email")
	//userinfo := UserModels.User{
	//  Username:  UserName,
	//  Password: Password,
	//  Email:    UserEmail,
	//}
	c.ShouldBind(&userinfo)
	Id, ok := UserModels.Register(userinfo, DB.Dbs)
	if ok {
		message = "Create Success"
		session := sessions.Default(c)
		session.Options(sessions.Options{
			MaxAge: 60 * 60, //1h
		})
		session.Set("username", userinfo.Username)
		session.Save()
	} else {
		Id = -1
		message = "Create Default"
	}
	UserModels.DoLog(Id, c.ClientIP(), message, DB.Dbs)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

//显示信息
func UserProfile(c *gin.Context) {
	var userinfo UserModels.User
	c.ShouldBind(&userinfo)
	id, _, ok := UserModels.Login(userinfo.Username, DB.Dbs)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "err",
		})
	} else {
		userprofile := UserModels.QueryUser(id, DB.Dbs)
		c.JSON(http.StatusOK, userprofile)
	}
}
