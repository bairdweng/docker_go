package models

// DeviceInfo 游戏信息 form:"pid" 可以进行转换，太牛逼了。
type DeviceInfo struct {
	BaseModel
	Pid        int    `json:"pid" form:"pid"`
	Idfa       string `json:"idfa" form:"idfa" binding:"required"`
	Idfv       string `json:"idfv" form:"idfv" binding:"required"`
	Imei       string `json:"imei" form:"imei" binding:"required"`
	Oaid       string `json:"oaid" form:"oaid" binding:"required"`
	Androidid  string `json:"androidid" form:"androidid" binding:"required"`
	DeviceID   string `json:"device_id" form:"device_id" binding:"required"`
	OsType     string `json:"os_type" form:"os_type" binding:"required"`
	OsVersion  string `json:"os_version" form:"os_version" binding:"required"`
	AppVersion string `json:"app_version" form:"app_version" binding:"required"`
	SdkVersion string `json:"sdk_version" form:"sdk_version" binding:"required"`
	Time       int    `json:"time" form:"time" binding:"required"`
}
