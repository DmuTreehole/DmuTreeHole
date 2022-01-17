package handler

import (
	"github.com/gin-gonic/gin"
	"main/models/certification"
	Utils "main/utils"
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
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.DatabaseDefault})
	} else {
		c.JSON(http.StatusOK, question)
	}
}
