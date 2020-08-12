package helper

import (
	"strconv"
	"time"
)

//GetCurrentDateTime 2
func GetCurrentDateTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 03:04:05 PM")
}

// StrToInt string 转int
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

// StrToUInt string 转int
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}
