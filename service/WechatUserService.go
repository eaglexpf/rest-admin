package service

import (
	//	"fmt"

	"errors"

	"github.com/eaglexpf/rest-admin/entity"
	"github.com/eaglexpf/rest-admin/pkg"
)

type WechatUserService struct{}

var db = pkg.DB

func (service *WechatUserService) ExistWechatUserByOpenid(openid string) entity.WechatUser {
	var user entity.WechatUser
	db.Where("openid=?", openid).First(&user)
	return user
}
func (service *WechatUserService) ExistWechatUserByID(id int) entity.WechatUser {
	var user entity.WechatUser
	db.Where("id=?", id).First(&user)
	return user
}

func (service *WechatUserService) CreateWechatUser(openid string) error {
	if service.ExistWechatUserByOpenid(openid).ID > 0 {
		return errors.New("该用户已存在")
	}
	var user entity.WechatUser
	user.Openid = openid
	user.Status = 1
	return db.Create(&user).Error
}

func (service *WechatUserService) GetWechatUserIDByOpenid(openid string) int {
	var user entity.WechatUser
	db.Where("openid=?", openid).First(&user)
	return user.ID
}
