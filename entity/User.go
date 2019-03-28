package entity

import (
	"time"

	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID             int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	UserName       string `gorm:"Column:username;Type:varchar(36);NOT NULL;UNIQUE_INDEX;"`
	AuthKey        string `gorm:"Column:auth_key;Type:varchar(36);NOT NULL;"`
	PassHash       string `gorm:"Column:pass_hash;Type:varchar(256);NOT NULL;"`
	ResetPassToken string `gorm:"Column:reset_pass_token;Type:varchar(256);"`
	Email          string `gorm:"Column:email;Type:varchar(256);NOT NULL;UNIQUE_INDEX;"`
	Status         int    `gorm:"Column:status;Type:int(11);NOT NULL;"`
	CreateAt       int    `gorm:"Column:crate_at;Type:int(11);NOT NULL;"`
	UpdateAt       int    `gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

var db = pkg.DB

var prefix = pkg.LoadData.DB.DBPrefix

func (User) TableName() string {
	return prefix + "user"
}

func (User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateAt", time.Now().Unix())
	return nil
}

func (User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateAt", time.Now().Unix())
	return nil
}

func init() {
	user := &User{}
	if !db.HasTable(user.TableName()) {
		db.CreateTable(user)
	}
}
