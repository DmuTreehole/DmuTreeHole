package user

import (
	"database/sql"
	"log"

	// DB "main/db"
	Tools "main/utils"

	"golang.org/x/crypto/bcrypt"
)

//用户登录和注册信息
type User struct {
	Id       int64  `json:'User_id'`
	Username string `json:'User_name'`
	Password string `json:'User_password'"`
	Email    string `json:'User_Email'`
}

//用户注册
func Register(User User, db *sql.DB) (int64, bool) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	User.Password = string(hash)
	template := "Insert User Set User_Name=?,User_Password=?,User_Email=?"
	if db==nil{
		log.Print("指针为空")
	}
	stmt, err := db.Prepare(template)
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
	return id, true
}

//用户登录
func Login(Username string, db *sql.DB) (int64, string, bool) {
	template := "Select User_Id,User_Password From User Where User_Name=?"
	rows, err := db.Query(template, Username)
	if err != nil {
		log.Print(err)
		return -1, "SQL Err!", false
	}
	var password string
	var Id int64
	rows.Next()
	err = rows.Scan(&Id, &password)
	if err != nil {
		log.Print(err)
		return -1, "login default", false
	}
	return Id, password, true
}

//记录日志
func Log(User User, Ip string, Info string, db *sql.DB) (string, bool) {
	template := "Insert Logs Set User_id=?,Log_Time=?,Log_Ip=?,Log_Info=?"
	stmt, err := db.Prepare(template)
	if err != nil {
		return err.Error(), false
	}
	userid := User.Id
	logTime := Tools.GetDatetime()
	_, err = stmt.Exec(userid, logTime, Ip, Info)
	if err != nil {
		log.Print(err)
		return err.Error(), false
	}
	return "", true
}
func DoLog(Id int64, Ip string, Info string, db *sql.DB) string {
	_, ok := Log(User{Id: Id}, Ip, Info, db)
	if ok {
		return "Log OK"
	}
	return "Log Default"
}
