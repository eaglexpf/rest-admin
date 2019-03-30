package service

import (
	//	"fmt"

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
