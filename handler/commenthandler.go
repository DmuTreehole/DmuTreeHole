package handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	post "main/models/post"
	Utils "main/utils"
	"net/http"
	"strconv"
)

// @Summary 打开树洞下面所有的评论
// @Description 打开树洞下面所有的评论
// @Success 200
// @Accept application/json
// @Produce application/json
// @Param body body post.Comment "评论请求体"
// @Tags 评论相关接口
// @Router /api/comment/getallcomment [post]
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
	fmt.Println(comment)
	response, err := post.ShowComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.DatabaseDefault})
		return
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
	if comment.Uid == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.NotLogin})
		return
	}
	//绑定树洞编号
	if c.ShouldBind(&comment) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.BindDefault})
		return
	}
	err := post.CreateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.DatabaseDefault})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": Utils.CreateCommentSuccess})
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
	if post.DeleteComment(cid) != nil {
		c.JSON(400, gin.H{"code": Utils.DeleteCommentDefault})
		return
	}
	c.JSON(200, gin.H{"code": Utils.DeleteCommentSuccess})
}
