package handler

import (
	"log"
	UserModels "main/models/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)
//封禁的原因和封禁树洞的id
type BanedRequest struct{
	Post_id int
	Reason string
}
/*
目前假设管理员id为1
*/
//管理员对应的handler，用于封禁用户
// @Summary 管理员封禁用户
// @Description 管理员封禁用户
// @Success 200 
// @Accept application/json
// @Produce application/json
// @Param body body BanedRequest true "封禁请求体"
// @Tags 用户相关接口
// @Router /api/user/banusers [post]
func BanUsers(c*gin.Context)  {
	session:=sessions.Default(c)
	if IssuperUser:=session.Get("userid").(int);IssuperUser!=1{
		c.JSON(400,gin.H{"msg":"当前用户不具有封禁权限"})
	}
	var banedrequest BanedRequest
	if err:=c.ShouldBindJSON(&banedrequest);err!=nil{
		log.Print("绑定json失败")
	};
	var message string="用户封禁失败"
	//通过树洞的id,查找对应的用户
	_,ok:=UserModels.BanUserByPostid(banedrequest.Post_id,banedrequest.Reason)
	if ok {
		message = "用户封禁成功"
	} else {
		message = "用户封禁失败"
	}
	c.JSON(200,gin.H{"msg":message})
}
// @Summary 展示已经被ban的用户
// @Description 展示已经被ban的用户
// @Success 200 
// @Accept application/json
// @Produce application/json
// @Tags 用户相关接口
// @Router /api/user/showbanedlist [get]
func ShowBannedList(c*gin.Context){
	session:=sessions.Default(c)
	if IssuperUser:=session.Get("userid").(int);IssuperUser!=1{
		c.JSON(400,gin.H{"msg":"当前用户不具有封禁权限"})
	}
	banedlist,_:=UserModels.ShowBannedUsers()
	c.JSON(200,banedlist)
}