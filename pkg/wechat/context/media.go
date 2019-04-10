package context

import (
	"fmt"

	"encoding/json"

	"github.com/eaglexpf/rest-admin/pkg/wechat/util"
)

type ResponseUploadMedia struct {
	ResponseCommon

	Type     string `json:"type"`
	MediaId  string `json:"media_id"`
	CreateAt int64  `json:"create_at"`
}

const (
	UrlMediaUpload = "https://api.weixin.qq.com/cgi-bin/media/upload"
)

//新增临时素材
func (this *Context) UploadMedia(media_type, file_name string, file_data []byte) (responseData ResponseUploadMedia, err error) {
	uri := fmt.Sprintf("%s?access_token=%s&type=%s", UrlMediaUpload, this.GetAccessToken(), media_type)

	body, err := util.HttpPostFile(uri, "media", file_name, file_data, make(map[string]string))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	if responseData.ErrCode != 0 {
		err = fmt.Errorf("get set_user_remark error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}
