package handler

//
//import (
//  "fmt"
//  "github.com/astaxie/beego/validation"
//  "github.com/gin-gonic/gin"
//  "log"
//  "main/models"
//  UserModels "main/models/user"
//  "main/utils"
//  "net/http"
//)
//
//
//func GetAuth(c *gin.Context){
//  valid := validation.Validation{}
//  var user = UserModels.User{}
//  var msg string
//  c.ShouldBind(&user)
//  fmt.Println(user)
//  var a  = auth{Username: user.Username,Password: user.Password}
//  var data = make(map[string]interface{})
//  ok,_ := valid.Valid(&a)
//  if ok {
//    isExist := models.CheckAuth(a.Username,a.Password)
//    if !isExist{
//      token,err := utils.CreateToken(a.Username,a.Password)
//      if err != nil {
//        msg = "TokenErr"
//      }else{
//        msg = "Success"
//        data["token"] = token
//      }
//    }else{
//      msg = "AuthErr"
//    }
//  }else{
//    for _,err := range valid.Errors{
//      log.Println(err.Key,err.Message)
//    }
//  }
//  c.JSON(http.StatusOK,gin.H{
//    "message":msg,
//    "data":data,
//  })
//}
