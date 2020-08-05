package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Gdb 222我是
var Gdb *gorm.DB

//InitDataBaseWithDataBase 都是 docker.for.mac.host.internal
func InitDataBaseWithDataBase(name string) {
	var baseURL = "localhost"
	// baseURL = "docker.for.mac.host.internal"
	var password = "jiangye089"
	// password = "root"
	var url = fmt.Sprintf("root:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=%s", password, baseURL, name, "Asia%2FShanghai")
	db, err := gorm.Open("mysql", url)
	if err != nil {
		println("数据库连接失败", err.Error())
	}
	db.SingularTable(true)
	db.LogMode(true)
	Gdb = db
}
