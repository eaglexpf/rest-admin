package service

import (
	//	"fmt"

	"github.com/eaglexpf/rest-admin/entity"
	//	"github.com/eaglexpf/rest-admin/pkg"
)

type AdvertService struct{}
type adverList struct {
	List []entity.Advert
}

func (this *AdvertService) GetLogo() entity.Advert {
	var logo = entity.Advert{}
	db.Select("id,img").Where("type=?", "logo").First(&logo)
	return logo
}

func (this *AdvertService) GetAdvert() []entity.Advert {
	var advert = []entity.Advert{}
	db.Select("id,img").Where("type=?", "advert").Find(&advert)
	return advert
}

func (this *AdvertService) Create(img string, img_type string) error {
	//	err := db.Create(&entity.Advert{
	//		Img:  img,
	//		Type: img_type,
	//	}).Error
	//	var c pkg.Controller
	//	uri := "http://127.0.0.1:21001/mfw/baoshui/Api/createAdvert"
	//	data, err := c.HttpPostData()

	return nil
}
