package models

// AccessRecords 应用信息
type AccessRecords struct {
	BaseModel
	AppID uint   `json:"app_id" form:"app_id" binding:"required"`
	IP    string `json:"ip" form:"ip" binding:"required"`
}
