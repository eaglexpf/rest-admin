package service

import (
	"fmt"

	"strings"

	"time"

	"github.com/eaglexpf/rest-admin/entity"
	"github.com/eaglexpf/rest-admin/pkg"
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

func (PrizeService) Create(name, unit, img_url, icon_on, icon_off, scenes string, num, valid_start, valid_end int) error {
	var prize = &entity.Prize{
		Name:       name,
		Unit:       unit,
		ImgUrl:     img_url,
		IconOn:     icon_on,
		IconOff:    icon_off,
		Num:        num,
		ValidStart: valid_start,
		ValidEnd:   valid_end,
		SceneAlias: scenes,
	}
	err := db.Create(prize).Error
	if prize.ID > 0 {
		var c pkg.Controller
		var uri = "https://wechat.kayunzh.com/mfw/baoshui/Api/prizeCreate"
		var data = make(map[string]interface{})
		data["id"] = prize.ID
		data["name"] = name
		data["unit"] = unit
		data["img_url"] = img_url
		data["icon_on"] = icon_on
		data["icon_off"] = icon_off
		data["scenes"] = scenes
		data["num"] = num
		data["valid_start"] = valid_start
		data["valid_end"] = valid_end
		data["account"] = pkg.LoadData.Wechat.Account
		data["token"] = pkg.LoadData.Wechat.ApiToken
		data["timestamp"] = time.Now().Unix()
		sign := c.Sign(data)
		data["sign"] = sign
		fmt.Println(data)
		response, err := c.HttpPostData(uri, data)
		fmt.Println(response, err)
	}
	fmt.Println(prize.ID, prize.Name, prize)
	return err
}
