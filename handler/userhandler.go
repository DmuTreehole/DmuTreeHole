package handler

//与用户相关的处理器函数
import (
	"fmt"
	UserModels "main/models/user"
	Utils "main/utils"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//用户登录显示的页面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.html", gin.H{})
}

// @Summary 验证用户名和密码
// @Description 检查用户登陆，验证用户名和密码
// @Success 200
// @Accept application/json
// @Produce application/json
// @Tags 用户相关接口
// @Params body body UserModels.User "用户请求体"
// @Router /api/user/logincheck [post]
func LoginCheck(c *gin.Context) {
	var user = UserModels.User{}
	var code int
	//Username := c.PostForm("Username")
	//Password := c.PostForm("Password")
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.BindDefault})
		return
	}
	fmt.Println(user)
	Id, CurrentPassword, err := UserModels.Login(user.Username)
	if err == nil {
		ok := Utils.CobPassWord(user.Password, CurrentPassword)
		if ok {
			setSessionById(c, Id)
			code = Utils.LoginSuccess
		} else {
			code = Utils.PasswordWrong
		}
	} else {
		code = Utils.UserNameNotExists
	}
	UserModels.Log(Id, c.ClientIP(), strconv.Itoa(code))
	if code == Utils.LoginSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code":   code,
			"UserId": Id,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"code": code})
}

//注册界面
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// @Summary 用户注册
// @Description 用户注册
// @Success 200
// @Accept application/json
// @Produce application/json
// @Tags 用户相关接口
// @Params body body UserModels.User "用户请求体"
// @Router /api/user/registercheck [post]
func RegisterCheck(c *gin.Context) {
	var code int
	var userinfo = UserModels.User{}
	c.ShouldBind(&userinfo)
	Id, err := UserModels.Register(userinfo)
	if err != nil {
		code = Utils.RegisterDefault
	} else {
		code = Utils.RegisterSuccess
		setSessionById(c, Id)
	}
	UserModels.Log(Id, c.ClientIP(), strconv.Itoa(code))
	if code == Utils.RegisterSuccess {
		c.JSON(http.StatusOK, gin.H{"code": code})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"code": code})
}

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

// @Summary 得到用户名称
// @Description 通过id得到用户名称
// @Success 200
// @Accept application/json
// @Produce application/json
// @Tags 用户相关接口
// @Params body body UserModels.User "用户请求体"
// @Router /api/user/getusername [post]
func GetUserName(c *gin.Context) {
	var user = UserModels.User{}
	c.ShouldBind(&user)
	username, err := UserModels.GetUserNameById(user.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.UserNameNotExists})
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

// @Summary 获取用户头像
// @Description 获取用户头像
// @Success 200
// @Accept application/json
// @Produce application/json
// @Tags 用户相关接口
// @Params body body UserModels.User "用户请求体"
// @Router /api/user/getusericon [post]
func ShowUserIcon(c *gin.Context) {
	var user UserModels.User
	c.ShouldBind(&user)
	filename, err := UserModels.GetUserIcon(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.UserIconNotFound})
		return
	}
	filepath := `./Icon/` + filename + `.jpg`
	//c.Writer.Header().Add("content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	//c.Writer.Header().Set("content-Type","application/jpeg")
	c.File(filepath)
	//c.JSON(http.StatusOK,gin.H{"message":"ok"})
}

// @Summary 得到用户名称
// @Description 通过id得到用户名称
// @Success 200
// @Accept application/json
// @Produce application/json
// @Tags 用户相关接口
// @Params body body UserModels.User "用户请求体"
// @Router /api/user/getusername [post]
func UploadUserIcon(c *gin.Context) {

}
