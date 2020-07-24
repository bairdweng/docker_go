package helper

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTClaims 对象
type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
}

// SecretKey key
const SecretKey = "hello ios services"

// GetToken 获取token
func GetToken(userName string) (string, error) {
	// 生成token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// 24小时后过期，这里也可以存储必要的数据。
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userName"] = userName
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SecretKey))
	return tokenString, err
}

// ValiteToken 验证
func ValiteToken(strToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(strToken, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	return token, err
}

// GetTokenValue GetKey
func GetTokenValue(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
