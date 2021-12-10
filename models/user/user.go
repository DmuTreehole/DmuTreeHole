package user

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	DB "main/db"
	Tools "main/utils"
	"math/rand"
	"strconv"
	// DB "main/db"
)

//用户登录和注册信息
type User struct {
	Id       int    `json:"UserId"`
	Username string `json:"Username",form:"Username"`
	Password string `json:"Userpassword",form:"Password"`
	Email    string `json:"UserEmail"`
}

type Icon struct {
	UserId int    `json:"UserId"`
	Url    string `json:"Url"`
}
type IconGet struct {
	UserIds []int `json:"UserId"`
}

//用户注册
func Register(User User) (int, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	User.Password = string(hash)
	template := "Insert User Set User_Name=?,User_Password=?,User_Email=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		return -1, err
	}
	result, err := stmt.Exec(User.Username, User.Password, User.Email)
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

//用户登录
func Login(Username string) (int, string, error) {
	template := "Select User_Id,User_Password From User Where User_Name=?"
	rows, err := DB.DB().Query(template, Username)
	if err != nil {
		return -1, "", err
	}
	var password string
	var id int
	rows.Next()
	err = rows.Scan(&id, &password)
	if err != nil {
		return -1, "", err
	}
	return id, password, nil
}

//记录日志
func Log(Id int, Ip string, Info string) bool {
	template := "Insert Logs Set User_id=?,Log_Time=?,Log_Ip=?,Log_Info=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		return false
	}
	logTime := Tools.GetDatetime()
	_, err = stmt.Exec(Id, logTime, Ip, Info)
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}

//通过userID获取UserName
func GetUserNameById(Id int) (string, error) {
	template := "Select User_Name From User Where User_Id = ?"
	rows, err := DB.DB().Query(template, Id)
	if err != nil {
		return "", err
	}
	rows.Next()
	userName := ""
	err = rows.Scan(&userName)
	if err != nil {
		return "", err
	}
	return userName, nil
}

//查询用户头像
func GetUserIcon(Ids []int) ([]Icon, error) {
	template1 := `Select Icon_Name From User Where User_Id =?`
	template2 := `Update User Set Icon_Name=? Where User_Id=?`
	db := DB.DB()
	var result = []Icon{}
	var icon = Icon{}
	for _, id := range Ids {
		rows, err := db.Query(template1, id)
		if err != nil {
			return nil, err
		}
		rows.Next()
		err = rows.Scan(&icon.Url)
		if err != nil {
			return nil, err
		}
		if icon.Url != "nil" {
			icon.UserId = id
			result = append(result, icon)
		} else {
			icon.Url = "/Icon/rand" + strconv.Itoa(rand.Int()%9+1) + ".jpg"
			icon.UserId = id
			_, err = db.Query(template2, icon.Url, id)
			if err != nil {
				return nil, err
			}
			result = append(result, icon)
		}
	}
	return result, nil
}
