package helper

import "golang.org/x/crypto/bcrypt"

// EnPassword 密码加密
func EnPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// ValationPassWord 密码验证
func ValationPassWord(password string, enPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(enPassword), []byte(password))
	return err == nil
}
