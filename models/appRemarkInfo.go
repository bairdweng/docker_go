package models

// AppRemarkInfo app备注信息
type AppRemarkInfo struct {
	BaseModel
	HostName string `json:"host_name" form:"host_name"`
	IP       string `json:"ip" form:"ip"`
	BaseURL  string `json:"base_url" form:"base_url"`
	Other    string `json:"other" form:"other"`
}

// AppRemarkInfoResult 返回
type AppRemarkInfoResult struct {
	HostName string `json:"host_name" form:"host_name"`
	IP       string `json:"ip" form:"ip"`
	BaseURL  string `json:"base_url" form:"base_url"`
	Other    string `json:"other" form:"other"`
}
