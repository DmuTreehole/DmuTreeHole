package handler

//与用户相关的处理器函数
import (
	"fmt"
	"github.com/gin-contrib/sessions"
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
	var user = UserModels.User{}
	//Username := c.PostForm("Username")
	//Password := c.PostForm("Password")
	c.ShouldBind(&user)
	Id, CurrentPassword, exist := UserModels.Login(user.Username)
	if exist {
		ok := Utils.CobPassWord(user.Password, CurrentPassword)
		if ok {
			message = "Login Success"
			setSessionById(c, Id)
		} else {
			message = "Wrong PassWord"
		}
	} else {
		message = "Account Not Exists"
	}
	UserModels.Log(Id, c.ClientIP(), message)
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
	var message = "Create Default"
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
	Id, ok := UserModels.Register(userinfo)
	if ok {
		message = "Create Success"
		setSessionById(c, Id)
	} else {
		Id = -1
		message = "Create Default"
	}
	UserModels.Log(Id, c.ClientIP(), message)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

//json UserSex,User_NickName,User_Addr
func CreateUserProfile(c *gin.Context) {
	var userprofile UserModels.Userprofile
	c.ShouldBind(&userprofile)
	session := sessions.Default(c)
	userprofile.Id = session.Get("userid").(int)
	err := UserModels.CreateUser(userprofile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

//json UserSex,User_NickName,User_Addr
func ChangeUserProfile(c *gin.Context) {
	var userprofile UserModels.Userprofile
	c.ShouldBind(&userprofile)
	session := sessions.Default(c)
	userprofile.Id = session.Get("userid").(int)
	err := UserModels.UpdateUser(userprofile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

//json UserId
//显示信息
func ShowUserProfile(c *gin.Context) {
	var userinfo UserModels.User
	c.ShouldBind(&userinfo)
	userprofile, err := UserModels.QueryUser(userinfo.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userprofile)
}

func GetUserName(c *gin.Context) {
	var user = UserModels.User{}
	c.ShouldBind(&user)
	username, err := UserModels.GetUserNameById(user.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "dafault"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"username": username})
}

func setSessionById(c *gin.Context, Id int) {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge: 60 * 60 * 24 * 30 * 3, //3month
		Path:   "/",
	})
	session.Set("userid", Id)
	session.Save()
}
