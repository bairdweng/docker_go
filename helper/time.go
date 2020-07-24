package helper

import "time"

//GetCurrentDateTime 2
func GetCurrentDateTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 03:04:05 PM")
}
