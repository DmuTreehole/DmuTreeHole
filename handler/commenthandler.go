package handler

import (
	post "main/models/post"
	"net/http"
	"strconv"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @Summary 打开树洞下面所有的评论
// @Description 打开树洞下面所有的评论
// @Success 200 
// @Accept application/json
// @Produce application/json
// @Param id path int true "postid"
// @Tags 评论相关接口
// @Router /api/comment/getallcomment/:id [get]
func GetAllComment(c *gin.Context) {
	//var pid struct{
	//	Id int `json:"id"`
	//}
	//if err:=c.ShouldBindJSON(&pid) ;err!=nil{
	//	c.JSON(400, gin.H{"error": "Json绑定错误"})
	//}
	//response,_:=post.ShowComment(pid.Id)
	var comment post.Comment
	c.ShouldBind(&comment) //Uid 和 page
	response, err := post.ShowComment(comment)
	if err != nil {
		c.JSON(400, gin.H{"message": "default"})
	}
	c.JSON(http.StatusOK, response)
}
// @Summary 创建一个评论
// @Description 创建一个评论
// @Success 200 
// @Accept application/json
// @Produce application/json
// @Param body body post.Comment true "评论请求体"
// @Tags 评论相关接口
// @Router /api/comment/createonecomment [post]
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
// @Summary 删除一个评论
// @Description 删除一个评论
// @Success 200 
// @Accept application/json
// @Produce application/json
// @Param id path int true "commentid"
// @Tags 评论相关接口
// @Router /api/comment/deletecomment/:id [get]
func DeleteOneComment(c *gin.Context) {
	id := c.Params.ByName("id")
	//这里暂时采用comment id 来删除
	cid, _ := strconv.Atoi(id)
	if err := post.DeleteComment(cid); err != nil {
		c.JSON(400, gin.H{"msg": "评论删除失败"})
	}
	c.JSON(200, gin.H{"msg": "评论删除成功"})
}
