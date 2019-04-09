package cache

//import (
//	"time"
//)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, val string, timeout int64) error
	IsExist(key string) bool
	Delete(key string) error
}
