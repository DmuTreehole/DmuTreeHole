package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	post "main/models/post"
	"net/http"
	"strconv"
)

// @Summary 查看所有树洞
// @Description 查看所有树洞
// @Success 200
// @Accept application/json
// @Produce application/json
// @Param id path int true "页数"
// @Tags 树洞相关接口
// @Router /api/post/getallpost/:page [get]
func GetAllPost(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	response, _ := post.ViewPost(page)
	c.JSON(http.StatusOK, response)
}

// @Summary 创建一个树洞
// @Description 创建一个树洞
// @Success 200
// @Accept application/json
// @Produce application/json
// @Param body body post.Post true "树洞请求体"
// @Tags 树洞相关接口
// @Router /api/post/createonepost [post]
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

// @Summary 删除一个树洞
// @Description 删除一个树洞
// @Success 200
// @Accept application/json
// @Produce application/json
// @Param id path int true "postid"
// @Tags 树洞相关接口
// @Router /api/comment/deleteonepost/:id [get]
func DeleteOnePost(c *gin.Context) {
	id := c.Params.ByName("id")
	pid, _ := strconv.Atoi(id)
	if err := post.DeletePost(pid); err != nil {
		c.JSON(400, gin.H{"msg": "树洞删除失败"})
	}
	c.JSON(200, gin.H{"msg": "树洞删除成功"})
}

// @Summary 通过userid查树洞
// @Description 通过userid查树洞
// @Success 200
// @Accept application/json
// @Produce application/json
// @Param body body post.PagePost true "请求体"
// @Tags 树洞相关接口
// @Router /api/post/getpostbyid [post]
func GetPostById(c *gin.Context) {
	var info post.PagePost
	c.ShouldBind(&info)
	session := sessions.Default(c)
	info.Id = session.Get("userid").(int)
	allPost, err := post.QueryPostById(info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, allPost)
}
func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", nil)
}

// @Summary 通过部分内容查树洞
// @Description 通过部分内容查树洞
// @Success 200
// @Accept application/json
// @Produce application/json
// @Param body body Content.Content true "树洞请求体"
// @Tags 树洞相关接口
// @Router /api/post/Search [Post]
func SearchPostByContent(c *gin.Context) {
	var content = post.Content{}
	var data = make(map[string]interface{})
	c.BindJSON(&content)
	result, err := post.Query(content)
	if err != nil {
		data["message"] = err.Error()
	} else {
		data["data"] = result
	}
	c.JSON(http.StatusOK, data)
}
