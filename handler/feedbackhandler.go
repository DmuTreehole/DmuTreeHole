package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	post "main/models/post"
	Utils "main/utils"
	"net/http"
)
// @Summary 创建一个反馈
// @Description 用户创建反馈
// @Success 200
// @Accept application/json
// @Produce application/json
// @Param body body post.Feedback true "反馈请求体"
// @Tags 树洞相关接口
// @Router /api/post/createonefeedback [post]
func CreateOneFeedback(c *gin.Context) {
	requestfeed := post.Feedback{}
	session := sessions.Default(c)
	//bind data
	if c.ShouldBind(&requestfeed) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.BindDefault})
		return
	}
	tmp := session.Get("userid")
	if tmp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.NotLogin})
		return
	}
	requestfeed.Uid = tmp.(int)
	_, err := post.CreateFeedBack(requestfeed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.CreateFeedBackDefault})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": Utils.CreateFeedBackSuccess})
}