package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type User struct {
	UserId       int64  `json:'User_id'`
	UserName     string `json:'User_name'`
	UserPassword string `json:'User_password'"`
	UserEmail    string `json:'User_Email'`
	UserPhone    string `json:'User_Phone'`
}

var DB *sql.DB

func OpenDataBase() {
	var err error
	DB, err = sql.Open("mysql", "TreeHole:treehole@tcp(127.0.0.1:3306)/Server")
	if err != nil {
		os.Exit(-1)
	}
}
func Register(User User) (int64, bool) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(User.UserPassword), bcrypt.DefaultCost)
	User.UserPassword = string(hash)
	template := "Insert User Set User_Name=?,User_Password=?,User_Email=?"
	stmt, err := DB.Prepare(template)
	if err != nil {
		return -1, false
	}
	result, err := stmt.Exec(User.UserName, User.UserPassword, User.UserEmail)
	if err != nil {
		return -1, false
	}
	id, _ := result.LastInsertId()
	return id, true
}
func Login(UserName string) (int64, string, bool) {
	template := "Select User_Id,User_Password From User Where User_Name=?"
	rows, err := DB.Query(template, UserName)
	if err != nil {
		return -1, "SQL Err!", false
	}
	var password string
	var Id int64
	rows.Next()
	err = rows.Scan(&Id, &password)
	if err != nil {
		return -1, "login default", false
	}
	return Id, password, true
}
func Log(User User, Ip string, Info string) (string, bool) {
	template := "Insert Logs Set User_id=?,Log_Time=?,Log_Ip=?,Log_Info=?"
	stmt, err := DB.Prepare(template)
	if err != nil {
		return err.Error(), false
	}
	userid := User.UserId
	logTime := GetDatetime()
	_, err = stmt.Exec(userid, logTime, Ip, Info)
	if err != nil {
		return err.Error(), false
	}
	return "", true
}
