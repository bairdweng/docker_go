package mydatabase

import "github.com/jinzhu/gorm"

//Gdb 222我是
var Gdb *gorm.DB

//InitDataBaseWithDataBase 都是 docker.for.mac.host.internal
func InitDataBaseWithDataBase(name string) {
	var url = "root:root@tcp(docker.for.mac.host.internal:3306)/" + name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)
	if err != nil {
		println("数据库连接失败", err.Error())
	}
	db.SingularTable(true)
	db.LogMode(true)
	Gdb = db
	// defer db.Close()
}
