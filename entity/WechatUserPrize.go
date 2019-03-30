package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type WechatUserPrize struct {
	ID       int `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	UserID   int `gorm:"Column:user_id;Type:int(11);NOT NULL;"`
	LogID    int `gorm:"Column:log_id;Type:int(11);NOT NULL;"`
	PrizeID  int `gorm:"Column:prize_id;Type:int(11);NOT NULL;"`
	CreateAt int `gorm:"Column:crate_at;Type:int(11);NOT NULL;"`
}

func (WechatUserPrize) TableName() string {
	return prefix + "wechat_user_prize"
}

func (WechatUserPrize) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateAt", time.Now().Unix())
	return nil
}

func init() {
	var userPrize = &WechatUserPrize{}
	if !db.HasTable(userPrize.TableName()) {
		db.CreateTable(userPrize)
	}
}
