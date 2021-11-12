package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

func init() {
	var err error
	DB, err = sql.Open("mysql", "Server:treehole@tcp(127.0.0.1:3306)/Server")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
func Resginer(User User) bool {
	stmt, err := DB.Prepare("Insert User Set User_id = 1,User_Name=?,User_Password=?")
	if err != nil {
		return false
	}
	e, err := stmt.Exec(User.UserName, User.UserPassword)
	if err != nil {
		return false
	}
	template := "Insert UserInfo Set UserInfo_Id=1,User_Id=?,User_Phone=?,User_Email=?"
	stmt1, err := DB.Prepare(template)
	if err != nil {
		return false
	}
	User.UserId, _ = e.LastInsertId()
	_, err = stmt1.Exec(User.UserId, User.UserPhone, User.UserEmail)
	if err != nil {
		return false
	}
	return true
}

func fin() {
	DB.Close()
}
