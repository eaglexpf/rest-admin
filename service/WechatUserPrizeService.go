package service

import (
	//	"fmt"

	"strings"

	"strconv"

	"github.com/eaglexpf/rest-admin/entity"
)

type WechatUserPrizeService struct{}

func (WechatUserPrizeService) GetUserPrizeListByLogID(user_id, log_id int) []entity.Prize {
	var userPrize = []entity.WechatUserPrize{}
	db.Where("user_id=? and log_id=?", user_id, log_id).Find(&userPrize)
	var ids []int
	for _, value := range userPrize {
		ids = append(ids, value.PrizeID)
	}
	var data = []entity.Prize{}
	db.Where("id in (?)", ids).Find(&data)
	return data
}

func (WechatUserPrizeService) InsertUserPrize(user_id, log_id int, prize_ids string) error {
	str := strings.Split(prize_ids, ",")
	for _, value := range str {
		prize_id, _ := strconv.Atoi(value)
		if prize_id <= 0 {
			continue
		}
		db.Create(&entity.WechatUserPrize{
			UserID:  user_id,
			LogID:   log_id,
			PrizeID: prize_id,
		})
	}
	return nil
}

func (WechatUserPrizeService) GetPrizeListByIds(prize_ids string) []entity.Prize {
	str := strings.Split(prize_ids, ",")
	var ids []int
	for _, value := range str {
		prize_id, _ := strconv.Atoi(value)
		if prize_id <= 0 {
			continue
		}
		ids = append(ids, prize_id)
	}
	var data = []entity.Prize{}
	db.Where("id in (?)", ids).Find(&data)
	return data
}
