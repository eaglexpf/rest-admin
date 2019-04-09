package context

import (
	"fmt"

	"encoding/json"

	"github.com/eaglexpf/rest-admin/pkg/wechat/util"
)

const (
	ShortUrl = "https://api.weixin.qq.com/cgi-bin/shorturl"
)

type ResponseShortUrl struct {
	ErrCode  int64  `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	ShortUrl string `json:"short_url"`
}

//长地址转短地址
func (this *Context) LongToShortUrl(long_url string) (responseData ResponseShortUrl, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", ShortUrl, this.GetAccessToken())

	postData := map[string]interface{}{
		"action":   "long2short",
		"long_url": long_url,
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
		err = fmt.Errorf("get short2url error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}
