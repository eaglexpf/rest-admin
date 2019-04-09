package context

import (
	"fmt"

	"encoding/json"

	"time"

	"github.com/eaglexpf/rest-admin/pkg/wechat/util"
)

const (
	AccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token"
)

type ResponseAccessToken struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//刷新access_token
func (this *Context) ResetAccessTokenFromServer() (access_token string) {
	uri := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", AccessTokenUrl, this.AppId, this.AppSecret)
	var body []byte
	var err error
	body, err = util.HttpGet(uri)
	if err != nil {
		access_token = err.Error()
		return
	}
	var responseAccessToken ResponseAccessToken
	err = json.Unmarshal(body, &responseAccessToken)
	if err != nil {
		access_token = err.Error()
		return
	}
	if responseAccessToken.ErrCode != 0 {
		access_token = fmt.Sprintf("get access_token error : errcode=%v , errmsg=%v", responseAccessToken.ErrCode, responseAccessToken.ErrMsg)
		return
	}
	accessTokenKey := fmt.Sprintf("access_token_%s", this.AppId)
	access_token = responseAccessToken.AccessToken
	expires_in := responseAccessToken.ExpiresIn - 1200 + time.Now().Unix()
	err = this.Cache.Set(accessTokenKey, responseAccessToken.AccessToken, expires_in)
	if err != nil {
		access_token = err.Error()
	}
	return
}

//获取access_token
func (this *Context) GetAccessToken() (access_token string) {
	//	this.AccessTokenLock.Lock()
	//	defer this.AccessTokenLock.Unlock()
	var err error
	accessTokenKey := fmt.Sprintf("access_token_%s", this.AppId)
	access_token, err = this.Cache.Get(accessTokenKey)
	if access_token == "" || err != nil {
		access_token = this.ResetAccessTokenFromServer()
	}
	return
}
