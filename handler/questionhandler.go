package handler

import (
	"github.com/gin-gonic/gin"
	"main/models/certification"
	"net/http"
)
// @Summary 查看注册问题
// @Description 查看注册问题，一次三个随机问题
// @Success 200 
// @Accept application/json
// @Produce application/json
// @Tags 树洞相关接口
// @Router /api/question [get]
func GetQuestion(c *gin.Context) {
	question, err := certification.GetThreeQuestions()
	if err != nil {
		c.JSON(400, gin.H{"message": "default"})
	} else {
		c.JSON(http.StatusOK, question)
	}

}
