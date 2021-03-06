package user

import (
	"golang.org/x/crypto/bcrypt"
	DB "main/db"
	"math/rand"
	"os"
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

//用户注册
func Register(User User) (int, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	User.Password = string(hash)
	template := "Insert User Set User_Name=?,User_Password=?,User_Email=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(User.Username, User.Password, User.Email)
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

//用户登录
func Login(Username string) (int, string, error) {
	template := "Select User_Id,User_Password From User Where User_Name=? Limit 1"
	rows, err := DB.DB().Query(template, Username)
	if err != nil {
		return -1, "", err
	}
	defer rows.Close()
	var password string
	var id int
	rows.Next()
	err = rows.Scan(&id, &password)
	if err != nil {
		return -1, "", err
	}
	return id, password, nil
}

//通过userID获取UserName
func GetUserNameById(Id int) (string, error) {
	template := "Select User_Name From User Where User_Id = ? Limit 1"
	rows, err := DB.DB().Query(template, Id)
	defer rows.Close()
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
func GetUserIcon(Id int) (string, error) {
	template := `Select Icon_Name From User Where User_Id =? limit 1`
	rows, err := DB.DB().Query(template, Id)
	defer rows.Close()
	rows.Next()
	var filename string
	err = rows.Scan(&filename)
	if err != nil {
		return "", err
	}
	if filename != "nil" {
		filepath := "./Icon/" + filename
		_, err := os.Stat(filepath)
		if err == nil {
			return filename, nil
		}
	}
	template = `Update User Set Icon_Name=? Where User_Id=?`
	iconName := "rand" + strconv.Itoa(rand.Int()%9+1)
	rows1, err := DB.DB().Query(template, iconName, Id)
	rows1.Close()
	if err != nil {
		return "", err
	}
	return iconName, nil
}

//更换头像
func UpdateIcon(filename string, UserId int) error {
	template := `Update User Set Icon_Name = ? Where User_Id = ?`
	rows, err := DB.DB().Query(template, filename, UserId)
	defer rows.Close()
	if err != nil {
		return err
	}
	return nil
}
