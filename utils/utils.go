package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"main/db"
	"net/smtp"
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

func SendCode(context string, email string) error {
	user := "dmutreehole@163.com"
	password := "DLCHYHPHXZVTIIGJ"
	host := "smtp.163.com:25"
	to := email
	subject := `海大树洞`
	body := `<html><body><h3>` + context + `</h3></body></html>`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		return err
	}
	fmt.Println("Send mail success!")
	return nil
}

type AuthCode struct {
	ToEmail string `json:"email"`
	Code    string `json:"code"`
}

//SendToMail 发送邮件的函数
func SendToMail(user, password, host, to, subject, body, mailType string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}
