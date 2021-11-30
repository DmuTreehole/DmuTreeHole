package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	post "main/models/post"
	"net/http"
	"strconv"
)

//打开一个树洞下面所有的评论
func GetAllComment(c *gin.Context) {
	//var pid struct{
	//	Id int `json:"id"`
	//}
	//if err:=c.ShouldBindJSON(&pid) ;err!=nil{
	//	c.JSON(400, gin.H{"error": "Json绑定错误"})
	//}
	//response,_:=post.ShowComment(pid.Id)
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := post.ShowComment(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "default"})
	}
	c.JSON(http.StatusOK, response)
}

//创建一个评论
func CreateOneComment(c *gin.Context) {
	var comment post.Comment
	session := sessions.Default(c)
	comment.Uid = session.Get("userid").(int)
	//绑定树洞编号
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(400, gin.H{"error": "Json绑定错误"})
	}
	_, err := post.CreateComment(comment)
	if err != nil {
		c.JSON(400, gin.H{"msg": "创建评论失败"})
	}
	c.JSON(200, gin.H{"msg": "评论创建成功"})
}
func DeleteOneComment(c *gin.Context) {
	id := c.Params.ByName("id")
	//这里暂时采用comment id 来删除
	cid, _ := strconv.Atoi(id)
	if err := post.DeleteComment(cid); err != nil {
		c.JSON(400, gin.H{"msg": "评论删除失败"})
	}
	c.JSON(200, gin.H{"msg": "评论删除成功"})
}
