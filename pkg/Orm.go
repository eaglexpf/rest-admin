package pkg

import (
	"fmt"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(LoadData.DB.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		LoadData.DB.DBUser,
		LoadData.DB.DBPassword,
		LoadData.DB.DBHost,
		LoadData.DB.DBName))
	//	defer DB.Close()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(1000)
}

//func DB() *gorm.DB {
//	return DB
//}
