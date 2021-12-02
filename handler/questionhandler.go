package handler

import (
	"github.com/gin-gonic/gin"
	"main/models/certification"
	"net/http"
)

func GetQuestion(c *gin.Context) {
	question, err := certification.GetThreeQuestions()
	if err != nil {
		c.JSON(400, gin.H{"message": "default"})
	} else {
		c.JSON(http.StatusOK, question)
	}

}
