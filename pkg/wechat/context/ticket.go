package context

import (
	"fmt"

	"encoding/json"

	"github.com/eaglexpf/rest-admin/pkg/wechat/util"
)

const (
	TicketUrl = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	QrCodeUrl = "https://mp.weixin.qq.com/cgi-bin/showqrcode"
)

type ResponseTicket struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`

	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	Url           string `json:"url"`
}

func (this *Context) QrCodeUrl(ticket string) string {
	return fmt.Sprintf("%s?ticket=%s", QrCodeUrl, ticket)
}

//临时二维码ticket--根据id生成
func (this *Context) TicketShortTimeSceneId(expire_seconds, scene_id int64) (responseData ResponseTicket, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", TicketUrl, this.GetAccessToken())

	postData := map[string]interface{}{
		"expire_seconds": expire_seconds,
		"action_name":    "QR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_id": scene_id,
			},
		},
	}
	var body []byte
	body, err = util.HttpPostJson(uri, postData)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	if responseData.ErrCode != 0 {
		err = fmt.Errorf("get ticket error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//临时二维码ticket--根据字符串生成
func (this *Context) TicketShortTimeSceneStr(expire_seconds int64, scene_str string) (responseData ResponseTicket, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", TicketUrl, this.GetAccessToken())

	postData := map[string]interface{}{
		"expire_seconds": expire_seconds,
		"action_name":    "QR_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": scene_str,
			},
		},
	}
	var body []byte
	body, err = util.HttpPostJson(uri, postData)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	if responseData.ErrCode != 0 {
		err = fmt.Errorf("get ticket error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//永久二维码ticket--根据id生成
func (this *Context) TicketLimitSceneId(expire_seconds, scene_id int64) (responseData ResponseTicket, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", TicketUrl, this.GetAccessToken())

	postData := map[string]interface{}{
		"action_name": "QR_LIMIT_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_id": scene_id,
			},
		},
	}
	var body []byte
	body, err = util.HttpPostJson(uri, postData)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	if responseData.ErrCode != 0 {
		err = fmt.Errorf("get ticket error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//永久二维码ticket--根据字符串生成
func (this *Context) TicketLimitSceneStr(expire_seconds int64, scene_str string) (responseData ResponseTicket, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", TicketUrl, this.GetAccessToken())

	postData := map[string]interface{}{
		"action_name": "QR_LIMIT_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": scene_str,
			},
		},
	}
	var body []byte
	body, err = util.HttpPostJson(uri, postData)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	if responseData.ErrCode != 0 {
		err = fmt.Errorf("get ticket error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}
