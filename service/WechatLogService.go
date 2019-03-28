package service

import (
	//	"fmt"

	//	"errors"

	"time"

	"github.com/eaglexpf/rest-admin/entity"
	//	"github.com/eaglexpf/rest-admin/pkg"
)

type WechatLogService struct{}

//var db = pkg.DB

func (service *WechatLogService) ExistWechatLog() entity.WechatLog {
	var log entity.WechatLog
	db.Where("create_at<=? and end_at>?", time.Now().Unix(), time.Now().Unix()).First(&log)
	return log
}

func (service *WechatLogService) ExistWechatLogByTicket(ticket string) bool {
	var log entity.WechatLog
	db.Where("Ticket=?", ticket).First(&log)
	if log.ID > 0 {
		return true
	}
	return false
}

func (service *WechatLogService) CreateWechatLog(user_id int, ticket string) error {
	return db.Create(&entity.WechatLog{
		UserID:   user_id,
		Ticket:   ticket,
		CreateAt: time.Now().Unix(),
		EndAt:    time.Now().Unix() + 180,
	}).Error
}
