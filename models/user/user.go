package user

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	DB "main/db"
	Tools "main/utils"
	// DB "main/db"
)

//用户登录和注册信息
type User struct {
	Id       int    `json:'User_id'`
	Username string `json:'User_name',form:'Username'`
	Password string `json:'User_password',form:'Password"`
	Email    string `json:'User_Email'`
}

//用户注册
func Register(User User) (int, bool) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	User.Password = string(hash)
	template := "Insert User Set User_Name=?,User_Password=?,User_Email=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		log.Print(err)
		return -1, false
	}
	result, err := stmt.Exec(User.Username, User.Password, User.Email)
	if err != nil {
		log.Print(err)
		return -1, false
	}
	id, _ := result.LastInsertId()
	return int(id), true
}

//用户登录
func Login(Username string) (int, string, bool) {
	template := "Select User_Id,User_Password From User Where User_Name=?"
	rows, err := DB.DB().Query(template, Username)
	if err != nil {
		log.Print(err)
		return -1, "SQL Err!", false
	}
	var password string
	var Id int
	rows.Next()
	err = rows.Scan(&Id, &password)
	if err != nil {
		log.Print(err)
		return -1, "login default", false
	}
	return Id, password, true
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
