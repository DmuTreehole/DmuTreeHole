package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	post "main/models/post"
	"net/http"
	"strconv"
)

//查看所有的树洞
func GetAllPost(c *gin.Context) {
	//session := sessions.Default(c)
	//id := session.Get("userid").(int)
	//type page struct {
	//	id int `json:"id"`
	//}
	//p := page{}
	//if err := c.BindJSON(&p); err != nil {
	//	c.JSON(400, gin.H{"error": "Json绑定错误"})
	//}
	page, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		response, _ := post.ViewPost(page)
		c.JSON(http.StatusOK, response)
	}
}

//创建树洞
func CreateOnePost(c *gin.Context) {
	var requestpost post.Post
	session := sessions.Default(c)
	requestpost.Uid = session.Get("userid").(int)
	//bind data
	if err := c.ShouldBindJSON(&requestpost); err != nil {
		c.JSON(400, gin.H{"error": "Json绑定错误"})
	}
	_, err := post.CreatePost(requestpost)
	if err != nil {
		c.JSON(400, gin.H{"msg": "创建树洞失败"})
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
