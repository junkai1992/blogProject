package common

import (
	"github.com/dgrijalva/jwt-go"
	"blogBack/model"
	"time"
)

// 定义jwt密钥
var jwtKey = []byte("a_secret_creat")

// 定义用户加密结构体
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User)(string,error){
	// token过期时间:7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),// 过期时间设置
			IssuedAt: time.Now().Unix(),//token发放时间
			Issuer: "ginEssential.tech",//发型人
			Subject: "user token",//主题
		},
	}
	// 生成jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	//转换string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil{
		return "",err
	}
	return tokenString, nil
}
// 解析token
func ParseToken(tokenString string)(*jwt.Token, *Claims, error) {
	claims := &Claims{}
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString,claims,func(token *jwt.Token)(i interface{}, err error){
		return jwtKey,nil
	})
	return token,claims,err
}