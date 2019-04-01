package pkg

import (
	//	"fmt"

	"time"

	"log"

	"github.com/go-ini/ini"
)

type dbLoad struct {
	DBType     string
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	DBPrefix   string
}

type wechatLoad struct {
	Token      string
	AppID      string
	AppSecret  string
	ApiToken   string
	Account    string
	ApiMyToken string
}

type load struct {
	Cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string

	DB     dbLoad
	Wechat wechatLoad
}

var LoadData = &load{}

func init() {
	var err error
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadData.Cfg = cfg
	LoadData.loadBase()
	LoadData.loadServe()
	LoadData.loadDB()
	LoadData.loadWechat()
}

func (this *load) loadBase() {
	this.RunMode = this.Cfg.Section("").Key("RunMode").MustString("debug")
}

func (this *load) loadServe() {
	sec, err := this.Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	this.HttpPort, err = sec.Key("HttpPort").Int()
	if err != nil {
		log.Fatalf("Fail to get section 'server.HttpPort': %v", err)
	}
	readTimeOut, err := sec.Key("ReadTimeOut").Int()
	if err != nil {
		log.Fatalf("Fail to get section 'server.ReadTimeOut': %v", err)
	}
	this.ReadTimeout = time.Duration(readTimeOut) * time.Second
	writeTimeOut, err := sec.Key("WriteTimeOut").Int()
	if err != nil {
		log.Fatalf("Fail to get section 'server.WriteTimeOut': %v", err)
	}
	this.WriteTimeout = time.Duration(writeTimeOut) * time.Second
}

func (this *load) loadDB() {
	sec, err := this.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	this.DB.DBType = sec.Key("DBType").String()
	this.DB.DBHost = sec.Key("DBHost").String()
	this.DB.DBName = sec.Key("DBName").String()
	this.DB.DBUser = sec.Key("DBUser").String()
	this.DB.DBPassword = sec.Key("DBPassword").String()
	this.DB.DBPrefix = sec.Key("DBPrefix").String()
}

func (this *load) loadWechat() {
	sec, err := this.Cfg.GetSection("wechat")
	if err != nil {
		log.Fatalf("Fail to get section 'wechat': %v", err)
	}
	this.Wechat.Token = sec.Key("Token").String()
	this.Wechat.AppID = sec.Key("AppID").String()
	this.Wechat.AppSecret = sec.Key("AppSecret").String()
	this.Wechat.ApiToken = sec.Key("ApiToken").String()
	this.Wechat.Account = sec.Key("Account").String()
	this.Wechat.ApiMyToken = sec.Key("ApiMyToken").String()
}
