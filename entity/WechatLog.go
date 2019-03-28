package entity

import (
	"time"

	//	"github.com/eaglexpf/rest-admin/pkg"
	"fmt"

	"github.com/jinzhu/gorm"
)

type WechatLog struct {
	ID       int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	UserID   int    `gorm:"Column:user_id;Type:int(11);NOT NULL;"`
	Ticket   string `gorm:"Column:ticket;Type:varchar(255);NOT NULL;"`
	CreateAt int64  `gorm:"Column:create_at;Type:int(11);NOT NULL;"`
	EndAt    int64  `gorm:"Column:end_at;Type:int(11);NOT NULL;"`
	UpdateAt int    `gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

//var db = pkg.DB

//var prefix = pkg.LoadData.DB.DBPrefix

func (WechatLog) TableName() string {
	return prefix + "wechat_log"
}

//func (WechatLog) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreateAt", time.Now().Unix())
//	return nil
//}

func (WechatLog) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateAt", time.Now().Unix())
	return nil
}

func init() {
	log := &WechatLog{}
	fmt.Println(db.HasTable(log.TableName()))
	if !db.HasTable(log.TableName()) {
		db.CreateTable(log)
	}
}
