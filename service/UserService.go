package service

import (
	"fmt"

	"github.com/eaglexpf/rest-admin/entity"
	//	"github.com/eaglexpf/rest-admin/pkg"
)

type UserService struct{}

//var db = pkg.DB

func (service *UserService) Count() int {
	var user entity.User
	var count int
	db.Table(user.TableName()).Count(&count)
	fmt.Println(count)
	return count
}

func (service *UserService) FindOne(id int) (entity.User, error) {
	var user entity.User
	err := db.Where("id=?", id).First(&user).Error
	return user, err
}
