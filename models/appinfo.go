package models

// AppInfo 应用信息
type AppInfo struct {
	BaseModel
	AppID    string `json:"app_id" form:"app_id" binding:"required"`
	AppName  string `json:"app_name" form:"app_name" binding:"required"`
	Status   string `json:"status" form:"status"`
	BundleID string `json:"bundle_id" form:"bundle_id"`
	IsCrash  bool   `json:"is_crash" form:"is_crash"`
	CreateID uint   `json:"create_id" form:"create_id" binding:"required"`
}

// AppInfoResult 返回
type AppInfoResult struct {
	ID       uint                `json:"id" form:"id"`
	AppID    string              `json:"app_id" form:"app_id" binding:"required"`
	AppName  string              `json:"app_name" form:"app_name" binding:"required"`
	Status   string              `json:"status" form:"status"`
	BundleID string              `json:"bundle_id" form:"bundle_id"`
	IsCrash  bool                `json:"is_crash" form:"is_crash"`
	Remark   AppRemarkInfoResult `json:"remark" form:"remark"`
}
