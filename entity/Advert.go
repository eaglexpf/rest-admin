package entity

//	"fmt"

//	"github.com/jinzhu/gorm"

type Advert struct {
	ID   int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;" json:"id"`
	Img  string `gorm:"Column:img;Type:varchar(255);NOT NULL;" json:"img"`
	Type int    `gorm:"Column:type;Type:varchar(30);NOT NULL;" json:"-"`
}

func (Advert) TableName() string {
	return prefix + "advert"
}

func init() {
	advert := &Advert{}
	if !db.HasTable(advert.TableName()) {
		db.CreateTable(advert)
	}
}
