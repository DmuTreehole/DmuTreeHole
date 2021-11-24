package handler

import (
	post "main/models/post"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//打开一个树洞下面所有的评论
func GetAllComment(c *gin.Context){
	session := sessions.Default(c)
	session.Get("userid")
	var pid int
	if err:=c.ShouldBindJSON(&pid) ;err!=nil{
		c.JSON(400, gin.H{"error": "Json绑定错误"})
	}
	response,_:=post.ShowComment(pid)
	c.JSON(http.StatusOK,response)
}
//创建一个评论
func CreateOneComment(c*gin.Context){
	session:=sessions.Default(c)
	uid:=session.Get("userid").(int)
	//绑定树洞编号
	var comment post.Comment
	comment.Uid=uid
	if err:=c.ShouldBindJSON(&comment) ;err!=nil{
		c.JSON(400, gin.H{"error": "Json绑定错误"})
	}
	_,err:=post.CreateComment(comment)
	if err!=nil{
		c.JSON(400, gin.H{"msg": "创建评论失败"})
	}
	c.JSON(200, gin.H{"msg": "评论创建成功"})
}
func DeleteOneComment(c *gin.Context){
	//TODO：某用户是否有删除的权限
}