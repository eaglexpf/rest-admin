package context

import (
	"fmt"

	"encoding/json"

	"github.com/eaglexpf/rest-admin/pkg/wechat/util"
)

type ResponseMaterialUpload struct {
	ResponseCommon

	MediaId string `json:"media_id"`
	Url     string `json:"url"`
}

const (
	UrlAddNews        = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	UrlNewxUploadImg  = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	UrlMaterialUpload = "https://api.weixin.qq.com/cgi-bin/material/add_material"
)

//新增其他类型的永久素材
func (this *Context) UploadMaterial(media_type, file_name string, file_data []byte, post_data map[string]string) (responseData ResponseMaterialUpload, err error) {
	uri := fmt.Sprintf("%s?access_token=%s&type=%s", UrlMaterialUpload, this.GetAccessToken(), media_type)

	body, err := util.HttpPostFile(uri, "media", file_name, file_data, post_data)
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
