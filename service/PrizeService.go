package service

import (
	"fmt"

	"strings"

	"github.com/eaglexpf/rest-admin/entity"
)

type PrizeService struct{}

func (PrizeService) GetList(scenes string) map[string]interface{} {
	str := strings.Split(scenes, ",")
	var data = make(map[string]interface{})
	for _, value := range str {
		var item = []entity.Prize{}
		db.Where("scene_alias=?", value).Find(&item)
		data[value] = item
	}
	return data
}

func (PrizeService) Create(name, unit, img_url, icon_on, icon_off string, num, valid_start, valid_end int) error {
	var prize = &entity.Prize{
		Name:            name,
		Unit:            unit,
		UnityUrl:        img_url,
		IconUrlActive:   icon_on,
		IconUrlInactive: icon_off,
		Num:             num,
		ValidStart:      valid_start,
		ValidEnd:        valid_end,
	}
	err := db.Create(prize).Error
	fmt.Println(prize.ID, prize.Name, prize)
	return err
}
