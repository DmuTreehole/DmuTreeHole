package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	UserModels "main/models/user"
	Utils "main/utils"
	"net/http"
)

//封禁的原因和封禁树洞的id
type BanedRequest struct {
	Post_id int
	Reason  string
}

/*
目前假设管理员id为1
*/
//管理员对应的handler，用于封禁用户
// @Summary 管理员封禁用户
// @Description 管理员封禁用户
// @Success 200
// @Accept application/json
// @Produce application/json
// @Param body body BanedRequest true "封禁请求体"
// @Tags 用户相关接口
// @Router /api/user/banusers [post]
func BanUsers(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userid").(int)
	if userId != Utils.SuperUserId {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.NotSuperUser})
		return
	}
	var bannedRequest BanedRequest
	if c.ShouldBind(&bannedRequest) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.BindDefault})
		return
	}
	//通过树洞的id,查找对应的用户
	_, ok := UserModels.BanUserByPostid(bannedRequest.Post_id, bannedRequest.Reason)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.BannedDefault})
		return
	}
	c.JSON(200, gin.H{"code": Utils.BannedSuccess})
}

// @Summary 展示已经被ban的用户
// @Description 展示已经被ban的用户
// @Success 200
// @Accept application/json
// @Produce application/json
// @Tags 用户相关接口
// @Router /api/user/showbanedlist [get]
func ShowBannedList(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userid").(int)
	if userId != Utils.SuperUserId {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.NotSuperUser})
		return
	}
	bannedList, err := UserModels.ShowBannedUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": Utils.DatabaseDefault})
		return
	}
	c.JSON(http.StatusOK, bannedList)
}
