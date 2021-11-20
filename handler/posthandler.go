package handler

import (
	DB "main/db"
	post "main/models/post"
	"net/http"

	"github.com/gin-gonic/gin"
)
//现在只实现了对0-5个数据的分页查询，用于前端的测试
func GetAllPost(c *gin.Context){
	response,_:=post.ViewPost(DB.Dbs)
	c.JSON(http.StatusOK,response)
}
//创建树洞
func CreateOnePost(c *gin.Context){
	var requestpost post.Post
	//bind data
	if err:=c.ShouldBindJSON(&requestpost);err!=nil{
		c.JSON(400,gin.H{"error":"Json绑定错误"})
	}
	_,err:=post.CreatePost(requestpost,DB.Dbs)
	if err!=nil{
		c.JSON(400,gin.H{"msg":"创建树洞失败"})
	}
	c.JSON(200,gin.H{"msg":"树洞创建成功"})
}