package context

import (
	"fmt"
	"net/http"
	"sync"

	"errors"

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
func (this *Context) Serve() error {
	if this.Validate() == false {
		return errors.New("请求效验失败")
	}
	if echo_str, ok := this.GetQuery("echostr"); ok {
		this.Strings(echo_str)
		return nil
	}
	msg, err := this.GetMessage()
	if err != nil {
		return err
	}
	this.RequestMsg = msg
	fmt.Println(msg)
	fmt.Println(this.MessageHandleFunc)
	if this.MessageHandleFunc == nil {
		this.Strings("success")
		return nil
	}
	response := this.MessageHandleFunc(msg)
	if response != nil {
		this.XML(response)
	} else {
		this.Strings("success")
	}

	return nil
}

func (this *Context) GetMessage() (msg message.RequestMessage, err error) {
	var body []byte
	body, err = ioutil.ReadAll(this.Request.Body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(body, &msg)
	return
}
