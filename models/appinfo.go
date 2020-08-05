package models

// AppInfo 应用信息
type AppInfo struct {
	BaseModel
	// DeletedAt *time.Time
	AppID   string `json:"app_id" form:"app_id" binding:"required"`
	AppName string `json:"app_name" form:"app_name" binding:"required"`
	Status  string `json:"status" form:"status"`
}
