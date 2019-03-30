package service

import (
	//	"fmt"

	"github.com/eaglexpf/rest-admin/entity"
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
