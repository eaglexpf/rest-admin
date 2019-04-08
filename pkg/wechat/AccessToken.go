package wechat

//import (
//	"encoding/json"
//	"fmt"
//	"time"
//)

//const (
//	AccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token"
//)

//type ResponseAccessToken struct {
//	ErrCode int64  `json:"errcode"`
//	ErrMsg  string `json:"errmsg"`

//	AccessToken string `json:"access_token"`
//	ExpiresIn   int64  `json:"expires_in"`
//}

//func (this *Context) GetAccessTokenFromServer() (responseAccessToken ResponseAccessToken, err error) {
//	uri := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", AccessTokenUrl, this.AppId, this.AppSecret)
//	var body []byte
//	body, err = Wechat.HttpGet(uri)
//	if err != nil {
//		return
//	}

//	err = json.Unmarshal(body, &responseAccessToken)
//	if err != nil {
//		return
//	}
//	if responseAccessToken.ErrCode != 0 {
//		err = fmt.Errorf("get access_token error : errcode=%v , errmsg=%v", responseAccessToken.ErrCode, responseAccessToken.ErrMsg)
//		return
//	}
//	accessTokenKey := fmt.Sprintf("access_token_%s", this.AppId)
//	expires_in := responseAccessToken.ExpiresIn - 1200
//	err = this.Cache.Set(accessTokenKey, responseAccessToken.AccessToken, time.Duration(expires_in)*time.Second)
//	return
//}

//func (this *Context) GetAccessToken() (access_token string, err error) {
//	this.AccessTokenLock.Lock()
//	defer this.AccessTokenLock.Unlock()

//	accessTokenKey := fmt.Sprintf("access_token_%s", this.AppId)
//	val := this.Cache.Get(accessTokenKey)
//	if val != nil {
//		access_token = val.(string)
//		return
//	}
//	var responseAccessToken ResponseAccessToken
//	responseAccessToken, err = this.GetAccessTokenFromServer()
//	if err != nil {
//		return
//	}
//	access_token = responseAccessToken.AccessToken
//	return
//}
