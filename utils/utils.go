package utils

import (
	"github.com/go-ini/ini"
	"golang.org/x/crypto/bcrypt"
	"main/db"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

type AuthCode struct {
	ToEmail string `json:"email",form:"email"`
	Code    string `json:"code",form:"code"`
	GenTime time.Time
}

type mail struct {
	user     string `ini:"user"`
	password string `ini:"password"`
	mailType string `ini:"type"`
}

var rootmail mail

func init() {
	cfg, _ := ini.Load("conf/app.ini")
	cfg.Section("email").MapTo(&rootmail)
}
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
	user := rootmail.user
	password := rootmail.password
	host := rootmail.mailType
	to := email
	subject := `海大树洞`
	body := `<html><body><a>您的验证码为</a><h3>` + context + `</h3><a><br/>验证码有效期为1小时，请在1小时内完成验证<br/>如果不是您本人操作，请忽略本条邮件</a></body></html>`
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		return err
	}
	return nil
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

//获取六位随机验证码
func GenVerCode() (string, time.Time) {
	result := ""
	directory := `0123456789ABCDEFGHJKLMNPQRSTUVWXYZ`
	for i := 0; i < 6; i++ {
		a := rand.Int() % 34
		result += directory[a : a+1]
	}
	return result, time.Now()
}
func AuthCodeRegister(email AuthCode) error {
	template := `Select * from AuthCode Where email=?`
	rows, err := db.DB().Query(template, email.ToEmail)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		template = `Update AuthCode Set code = ?,time=? Where email =?`
		rows, err = db.DB().Query(template, email.Code, email.GenTime, email.ToEmail)
		if err != nil {
			return err
		}
		defer rows.Close()
	} else {
		template = `Insert into AuthCode Set email=?,code=?,time=?`
		rows, err = db.DB().Query(template, email.ToEmail, email.Code, email.GenTime)
		if err != nil {
			return err
		}
		defer rows.Close()
	}
	return nil
}

func AuthCodeCheck(email AuthCode) (bool, error) {
	template := `Select code,time from AuthCode Where email=?`
	rows, err := db.DB().Query(template, email.ToEmail)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if !rows.Next() {
		return false, nil
	}
	current := new(AuthCode)
	rows.Scan(&current.Code, &current.GenTime)
	if email.Code == current.Code {
		return time.Now().Sub(current.GenTime).Hours() > 1, nil
	} else {
		return false, nil
	}
}
