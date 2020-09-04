package helper

import "errors"

// VerylyParamsByUserName 验证用户名
func VerylyParamsByUserName(str string) error {
	if len(str) < 6 || len(str) > 10 {
		return errors.New("用户名需要6～10位之间")
	}
	return nil
}

// VerylyParamsByPassWord 验证密码
func VerylyParamsByPassWord(str string) error {
	if len(str) != 8 {
		return errors.New("密码8位")
	}
	return nil
}
