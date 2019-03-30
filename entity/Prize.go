package entity

type Prize struct {
	ID              int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;" json:"id"`
	Name            string `gorm:"Column:name;Type:varchar(200);NOT NULL;" json:"name"`
	Unit            string `gorm:"Column:unit;Type:varchar(30);NOT NULL;" json:"unit"`
	UnityUrl        string `gorm:"Column:unity_url;Type:varchar(255);NOT NULL;" json:"unity_url"`
	IconUrlActive   string `gorm:"Column:icon_url_active;Type:varchar(255);NOT NULL;" json:"icon_url_active"`
	IconUrlInactive string `gorm:"Column:icon_url_inactive;Type:varchar(255);NOT NULL;" json:"icon_url_inactive"`
	Num             int    `gorm:"Column:num;Type:int(11);NOT NULL;" json:"num"`
	Type            int    `gorm:"Column:type;Type:int(4);NOT NULL;" json:"type"`
	Prob            int    `gorm:"Column:prob;Type:int(5);NOT NULL;" json:"prob"`
	SceneAlias      int    `gorm:"Column:scene_alias;Type:varchar(30);NOT NULL;" json:"-"`
}

func (Prize) TableName() string {
	return prefix + "prize"
}

func init() {
	prize := Prize{}
	if !db.HasTable(prize.TableName()) {
		db.CreateTable(prize)
	}
}
