package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

var privateKey string

const MaxAge = 60 * 60 * 24 * 30 // 一个月

type CustomClaims struct {
	UserId int
	jwt.StandardClaims
}

func init() {
	cfg, _ := ini.Load("conf/app.ini")
	privateKey = cfg.Section("jwt").Key("privateKey").String()
}

//CreateToken 生成一个新的token
func CreateToken(userId int) string {
	c := &CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(MaxAge) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(tokenString)
	return tokenString
}

//ParseToken 解析Token
//时间超过了token生存时间的二倍，则返回错误
//时间超过了token的生存时间但是未超过其二倍的话，则会为其生成一个新的token
func ParseToken(tokenString string) (*CustomClaims, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(privateKey), nil
	})
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, "", nil
	}
	t := time.Unix(claims.StandardClaims.ExpiresAt, 0)
	timeExceed := int(time.Now().Sub(t).Seconds())
	if timeExceed < MaxAge {
		CreateToken(claims.UserId)
		return claims, "", nil
	}
	return nil, "", err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		_, newToken, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"message": "请重新登录"})
			c.Abort()
			return
		}
		c.Next()
		if newToken != "" {
			c.Header("token", newToken)
		}
	}
}
