package handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	post "main/models/post"
	"net/http"
	"strconv"
)

//查看所有的树洞
func GetAllPost(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	fmt.Println(page)
	response, _ := post.ViewPost(page)
	c.JSON(http.StatusOK, response)
}

//创建树洞
func CreateOnePost(c *gin.Context) {
	requestpost := post.Post{}
	session := sessions.Default(c)
	//bind data
	if c.ShouldBind(&requestpost) != nil {
		c.JSON(400, gin.H{"error": "Json绑定错误"})
		return
	}
	tmp := session.Get("userid")
	if tmp == nil {
		c.JSON(http.StatusOK, gin.H{"message": "NotLogin"})
		return
	}
	requestpost.Uid = tmp.(int)
	_, err := post.CreatePost(requestpost)
	if err != nil {
		c.JSON(400, gin.H{"msg": "创建树洞失败"})
		return
	}
	c.JSON(200, gin.H{"msg": "树洞创建成功"})
}

//删除树洞
func DeleteOnePost(c *gin.Context) {
	id := c.Params.ByName("id")
	pid, _ := strconv.Atoi(id)
	if err := post.DeletePost(pid); err != nil {
		c.JSON(400, gin.H{"msg": "树洞删除失败"})
	}
	c.JSON(200, gin.H{"msg": "树洞删除成功"})
}
