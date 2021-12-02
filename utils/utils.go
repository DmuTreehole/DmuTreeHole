package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GetDatetime() (datetime string) {
	template := "2006-01-02 15:04:05"
	datetime = time.Now().Format(template)
	return
}
func BcrPassWord(password string, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password2), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

var jwtSecret = []byte("treeHole")

func CreateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "treeHole",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

//屏蔽关键字
//func fuck(context string) string {
//  changed := "***"
//  for _,word := range dictionary {
//    strings.ReplaceAll(context,word,changed)
//  }
//  return context
//}
