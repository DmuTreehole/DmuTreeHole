package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetDatetime() (datetime string) {
	templage := "2006-01-02 15:04:05"
	datetime = time.Now().Format(templage)
	return
}
func CobPassWord(Password string, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password2), []byte(Password))
	if err != nil {
		return false
	}
	return true
}
func DoLog(Id int64, Ip string, Info string) string {
	_, ok := Log(User{UserId: Id}, Ip, Info)
	if ok {
		return "Log OK"
	}
	return "Log Default"
}
