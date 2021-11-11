package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Models"
	"net/http"
)

func Router() {
	router := gin.Default()
	router.LoadHTMLFiles("/treeHole/View/index.html")
	router.StaticFile("View/Photo.jpg", "/treeHole/View/Photo.jpg")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
		fmt.Println(models.GetDatetime())
	})
	router.Run(":80")
}
