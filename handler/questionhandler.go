package handler

import (
	"github.com/gin-gonic/gin"
	"main/models/certification"
	"net/http"
)

func Getquestion(c *gin.Context) {
	question, err := certification.Getonequestion()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "default"})
	} else {
		c.JSON(http.StatusOK, question)
	}

}