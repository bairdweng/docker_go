package helper

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTClaims 对象
type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
}

// SecretKey key
const SecretKey = "hello ios services"

// GetToken 根据用户ID获取Token
func GetToken(userID uint) (string, error) {
	// 生成token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// 24小时后过期，这里也可以存储必要的数据。
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userID"] = userID
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

// GetUserIDByToken 根据token获取用户ID
func GetUserIDByToken(token string) uint {
	deToken2, err := ValiteToken(token)
	if err != nil {
		return 0
	}
	userID := GetTokenValue("userID", deToken2.Claims)
	i, e := strconv.Atoi(userID)
	if e != nil {
		return 0
	}
	return uint(i)
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
