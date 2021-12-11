package utils

import (
	"golang.org/x/crypto/bcrypt"
	"main/db"
	"strings"
	"time"
)

func GetDatetime() (datetime string) {
	template := "2006-01-02 15:04:05"
	datetime = time.Now().Format(template)
	return
}
func CobPassWord(password string, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password2), []byte(password))
	if err != nil {
		return false
	}
	return true
}

//屏蔽关键字
func Fuck(context string) string {
	changed := "***"
	var word string
	template := `Select Content From Keyword`
	rows, err := db.DB().Query(template)
	defer rows.Close()
	if err != nil {
		return context
	}
	for rows.Next() {
		rows.Scan(&word)
		context = strings.ReplaceAll(context, word, changed)
	}
	return context
}
