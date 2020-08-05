package models

// User 用户信息。
type User struct {
	BaseModel
	UserName string `json:"userName" form:"userName" binding:"required"`
	Password string `json:"passWord" form:"passWord" binding:"required"`
	Time     int64  `json:"time" form:"time"`
}
