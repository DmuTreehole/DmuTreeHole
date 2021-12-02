package models

import "main/db"

type Auth struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	template := `Select Id From Auth Where Username=? And Password=?`
	rows, err := db.DB().Query(template, username, password)
	if err != nil {
		return false
	}
	rows.Next()
	rows.Scan(&auth.Id)
	if auth.Id > 0 {
		return true
	}
	return false
}
