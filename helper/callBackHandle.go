package helper

//CallBack 返回值
type CallBack struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//Successful 返回值
func Successful(data interface{}) interface{} {
	var callBack = CallBack{}
	callBack.Data = data
	callBack.Status = 1
	callBack.Message = "成功"
	return callBack
}

//Error 返回值
func Error(message string, data interface{}) interface{} {
	var callBack = CallBack{}
	callBack.Data = data
	callBack.Status = 0
	callBack.Message = message
	return callBack
}
