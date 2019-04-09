package wechat

import (
	//	"fmt"
	"net/http"
	//	"sync"
	//	"time"

	"github.com/eaglexpf/rest-admin/pkg/wechat/cache"
	"github.com/eaglexpf/rest-admin/pkg/wechat/context"
	"github.com/eaglexpf/rest-admin/pkg/wechat/message"
)

type Config struct {
	AppId     string
	AppSecret string
	Token     string
	Writer    http.ResponseWriter
	Request   *http.Request
	Cache     cache.Cache
}

type Wechat struct {
	//	AppId     string
	//	AppSecret string
	//	Token     string

	Context *context.Context
}

func NewWechat(cfg *Config) *Wechat {
	con := &context.Context{
		AppId:     cfg.AppId,
		AppSecret: cfg.AppSecret,
		Token:     cfg.Token,
		Writer:    cfg.Writer,
		Request:   cfg.Request,
		Cache:     cfg.Cache,
	}
	return &Wechat{Context: con}
}

func (this *Wechat) Server() {
	this.Context.Serve()
}

func (this *Wechat) HandleFunc(f func(message.RequestMessage) interface{}) {
	this.Context.MessageHandleFunc = f
}

func (this *Wechat) GetAccessToken() string {
	return this.Context.GetAccessToken()
}

func (this *Wechat) GetQuery(key string) {

}
