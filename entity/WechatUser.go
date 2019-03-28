package entity

import (
	"time"

	//	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/jinzhu/gorm"
)

type WechatUser struct {
	ID       int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	Openid   string `gorm:"Column:openid;Type:varchar(255);NOT NULL;UNIQUE_INDEX;"`
	Status   int    `gorm:"Column:status;Type:int(11);NOT NULL;"`
	CreateAt int    `gorm:"Column:crate_at;Type:int(11);NOT NULL;"`
	UpdateAt int    `gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

//var db = pkg.DB

//var prefix = pkg.LoadData.DB.DBPrefix

func (WechatUser) TableName() string {
	return prefix + "wechat_user"
}

func (WechatUser) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateAt", time.Now().Unix())
	return nil
}

func (WechatUser) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateAt", time.Now().Unix())
	return nil
}

func init() {
	user := &WechatUser{}
	if !db.HasTable(user.TableName()) {
		db.CreateTable(user)
	}
}
