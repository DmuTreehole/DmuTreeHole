package handler

//与用户相关的处理器函数
import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"main/models"
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
	//var message string
	//var user = UserModels.User{}
	////Username := c.PostForm("Username")
	////Password := c.PostForm("Password")
	//c.ShouldBind(&user)
	//Id, CurrentPassword, exist := UserModels.Login(user.Username)
	//if exist {
	//	ok := Utils.CobPassWord(user.Password, CurrentPassword)
	//	if ok {
	//		message = "Login Success"
	//		setSessionById(c, Id)
	//	} else {
	//		message = "Wrong PassWord"
	//	}
	//} else {
	//	message = "Account Not Exists"
	//}
	//UserModels.Log(Id, c.ClientIP(), message)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": message,
	//})
	valid := validation.Validation{}
	var user = UserModels.User{}
	var msg string
	c.ShouldBind(&user)
	var a = models.Auth{Username: user.Username, Password: user.Password}
	var data = make(map[string]interface{})
	ok, _ := valid.Valid(&a)
	if ok {
		id, err := UserModels.Login(a.Username, a.Password)
		if err == nil {
			token, err := Utils.CreateToken(a.Username, a.Password)
			if err != nil {
				msg = "TokenErr"
				UserModels.Log(id, c.ClientIP(), "CreateTokenErr")
			} else {
				msg = "Success"
				data["token"] = token
			}
		} else {
			msg = "AuthErr"
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"data":    data,
	})
}

//注册界面
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

//检查注册
func RegisterCheck(c *gin.Context) {
	var message = "Create Default"
	var data = make(map[string]interface{})
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
		token, err := Utils.CreateToken(userinfo.Username, userinfo.Password)
		if err != nil {
			message += "token err"
		} else {
			data["token"] = token
		}
		//setSessionById(c, Id)
	} else {
		Id = -1
		message = "Create Default"
	}
	UserModels.Log(Id, c.ClientIP(), message)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

func ChangeUserProfile(c *gin.Context) {
	var userprofile UserModels.Userprofile
	c.ShouldBindJSON(&userprofile)
	session := sessions.Default(c)
	userprofile.Id = session.Get("userid").(int)
	if !UserModels.UpdateUser(userprofile) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "default"})
}

//显示信息
func ShowUserProfile(c *gin.Context) {
	var userinfo UserModels.User

	session := sessions.Default(c)
	userinfo.Username = session.Get("username").(string)

	id, err := UserModels.Login(userinfo.Username, userinfo.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	} else {
		userprofile, ok := UserModels.QueryUser(id)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"message": "default"})
		}
		c.JSON(http.StatusOK, userprofile)
	}
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
