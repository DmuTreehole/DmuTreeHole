package handler

import (
	post "main/models/post"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

//查看所有的树洞
func GetAllPost(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("userid")
	var page struct{
		Id int `json:"id"`
	}
	if err:=c.ShouldBindJSON(&page);err!=nil{
		c.JSON(400, gin.H{"error": "Json绑定错误"})
	}
	response, _ := post.ViewPost(page.Id)
	c.JSON(http.StatusOK, response)
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
func DeleteOnePost(c *gin.Context){
	id:=c.Params.ByName("id")
	pid,_:=strconv.Atoi(id)
	if err:=post.DeletePost(pid);err!=nil{
		c.JSON(400,gin.H{"msg":"树洞删除失败"})
	}
	c.JSON(200, gin.H{"msg": "树洞删除成功"})
}
