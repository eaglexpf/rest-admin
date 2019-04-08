package wechat

import (
	//	"fmt"
	"net/http"
	//	"sync"
	//	"time"

	"github.com/eaglexpf/rest-admin/pkg/wechat/context"
)

type Config struct {
	AppId     string
	AppSecret string
	Token     string
	Writer    http.ResponseWriter
	Request   *http.Request
}

type Wechat struct {
	AppId     string
	AppSecret string
	Token     string

	Context *context.Context
}

func NewWechat(cfg *Config) *Wechat {
	con := &context.Context{
		AppId:     cfg.AppId,
		AppSecret: cfg.AppSecret,
		Token:     cfg.Token,
		Writer:    cfg.Writer,
		Request:   cfg.Request,
	}
	return &Wechat{Context: con}
}

func (this *Wechat) Server() {}

func (this *Wechat) GetAccessToken() (string, error) {
	return this.Context.GetAccessToken()
}

func (this *Wechat) GetQuery(key string) {

}
