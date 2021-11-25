package handler

import (
	"github.com/gin-gonic/gin"
	"main/models/post"
	"net/http"
	"strconv"
)

//获取5条评论
func GetReply(c *gin.Context) {
	reply := post.Reply{}
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "default"})
	}
	err = c.ShouldBindJSON(&reply) //只要传comment_Id或者Reply_Id就好另外一个传-1，可递归
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "default"})
	}
	replys := post.ShowReply(reply, page)
	c.JSON(http.StatusOK, replys)
}

//删除评论
func DeleteReply(c *gin.Context) {
	reply := post.Reply{}
	err := c.ShouldBindJSON(&reply) //只要reply_id
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "default"})
	}
	if !post.DeleteReply(reply) {
		c.JSON(http.StatusOK, gin.H{"message": "default"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

//创建评论
func CreateReply(c *gin.Context) {
	reply := post.Reply{}
	err := c.ShouldBindJSON(&reply) //跟show一样comment的回复proreply置-1
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "default"})
	}
	if !post.CreateReply(reply) {
		c.JSON(http.StatusOK, gin.H{"message": "default"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
