package utils

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

//屏蔽关键字
//func fuck(context string) string {
//  changed := "***"
//  for _,word := range dictionary {
//    strings.ReplaceAll(context,word,changed)
//  }
//  return context
//}
