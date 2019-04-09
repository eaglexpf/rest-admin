package context

import (
	"fmt"
	"net/http"
	"sync"

	//	"errors"

	"io/ioutil"

	"encoding/xml"

	"github.com/eaglexpf/rest-admin/pkg/wechat/cache"
	"github.com/eaglexpf/rest-admin/pkg/wechat/message"
	"github.com/eaglexpf/rest-admin/pkg/wechat/util"
)

type Context struct {
	AppId     string
	AppSecret string
	Token     string

	Writer  http.ResponseWriter
	Request *http.Request

	AccessTokenLock *sync.RWMutex

	Cache cache.Cache

	RequestMsg message.RequestMessage

	MessageHandleFunc func(message.RequestMessage) interface{}
}

type ResponseCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (this *Context) Query(key string) string {
	value, _ := this.GetQuery(key)
	return value
}

func (this *Context) GetQuery(key string) (string, bool) {
	if values, ok := this.Request.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

//验证签名
func (this *Context) Validate() bool {
	timestamp := this.Query("timestamp")
	nonce := this.Query("nonce")
	signature := this.Query("signature")
	return signature == util.Signature(this.Token, timestamp, nonce)
}

//数据接入
func (this *Context) Serve() {
	if this.Validate() == false {
		this.Strings("请求效验失败")
		return
	}
	if echo_str, ok := this.GetQuery("echostr"); ok {
		this.Strings(echo_str)
		return
	}
	msg, err := this.GetMessage()
	if err != nil {
		this.Strings(err.Error())
		return
	}
	this.RequestMsg = msg
	fmt.Println(msg)
	fmt.Println(this.MessageHandleFunc)
	if this.MessageHandleFunc == nil {
		this.Strings("success")
		return
	}
	response := this.MessageHandleFunc(msg)
	if response != nil {
		this.XML(response)
	} else {
		this.Strings("success")
	}

}

//解析微信返回的数据
func (this *Context) GetMessage() (msg message.RequestMessage, err error) {
	var body []byte
	body, err = ioutil.ReadAll(this.Request.Body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(body, &msg)
	return
}
