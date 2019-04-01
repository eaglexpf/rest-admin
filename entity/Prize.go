package entity

type Prize struct {
	ID         int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;" json:"id"`
	Name       string `gorm:"Column:name;Type:varchar(200);NOT NULL;" json:"name" form:"name"`
	Unit       string `gorm:"Column:unit;Type:varchar(30);NOT NULL;" json:"unit" form:"unit"`
	ImgUrl     string `gorm:"Column:img_url;Type:varchar(255);NOT NULL;" json:"img_url" form:"img_url"`
	IconOn     string `gorm:"Column:icon_on;Type:varchar(255);NOT NULL;" json:"icon_on" form:"icon_on"`
	IconOff    string `gorm:"Column:icon_off;Type:varchar(255);NOT NULL;" json:"icon_off" form:"icon_off"`
	Num        int    `gorm:"Column:num;Type:int(11);NOT NULL;" json:"num" form:"num"`
	Type       int    `gorm:"Column:type;Type:int(4);NOT NULL;" json:"type" form:"type"`
	Prob       int    `gorm:"Column:prob;Type:int(5);NOT NULL;" json:"prob" form:"prob"`
	ValidStart int    `gorm:"Column:valid_start;Type:int(11);NOT NULL;" json:"valid_start" form:"valid_start"`
	ValidEnd   int    `gorm:"Column:valid_end;Type:int(11);NOT NULL;" json:"valid_end" form:"valid_start"`
	SceneAlias string `gorm:"Column:scene_alias;Type:varchar(255);NOT NULL;" json:"scene_alias" form:"scene_alias"`
}

func (Prize) TableName() string {
	return prefix + "wechat_prize"
}

func init() {
	prize := Prize{}
	if !db.HasTable(prize.TableName()) {
		db.CreateTable(prize)
	}
}
