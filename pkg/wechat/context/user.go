package context

import (
	"fmt"

	"encoding/json"

	"github.com/eaglexpf/rest-admin/pkg/wechat/util"
)

//用户列表结构体
type ResponseUserOpenidList struct {
	ResponseCommon

	Total int64 `json:"total"`
	Count int64 `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

//单个用户信息结构体
type ResponseUserInfo struct {
	ResponseCommon

	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	NickName       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"Country"`
	HeadImgUrl     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	Unionid        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupId        int    `json:"groupid"`
	TagIdList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subcribe_scene"`
	QrScene        int64  `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

//多个用户信息结构体
type ResponseUserInfoBatchget struct {
	ResponseCommon

	UserInfoList []ResponseUserInfo `json:"user_info_list"`
}

//用户标签列表结构体
type ResponseUserTagsList struct {
	ResponseCommon

	Tags []struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Count int64  `json:"count"`
	} `json:"tags"`
}

//创建用户标签返回的数据结构体
type ResponseCreateUserTags struct {
	ResponseCommon

	Tag struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tag"`
}

//标签下的粉丝列表结构体
type ResponseTagsOpenidList struct {
	ResponseCommon

	Count int64 `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

const (
	UserOpenidListUrl   = "https://api.weixin.qq.com/cgi-bin/user/get"
	UserInfoUrl         = "https://api.weixin.qq.com/cgi-bin/user/info"
	UserInfoBatchgetUrl = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"

	UserTagsListUrl   = "https://api.weixin.qq.com/cgi-bin/tags/get"
	UserTagsCreateUrl = "https://api.weixin.qq.com/cgi-bin/tags/create"
	UserTagsUpdateUrl = "https://api.weixin.qq.com/cgi-bin/tags/update"
	UserTagsDeleteUrl = "https://api.weixin.qq.com/cgi-bin/tags/delete"

	TagsOpenidListUrl = "https://api.weixin.qq.com/cgi-bin/user/tag/get"
)

//获取用户列表
func (this *Context) GetUserOpenidList(next_openid string) (responseData ResponseUserOpenidList, err error) {
	uri := fmt.Sprintf("%s?access_token=%s&next_openid=%s", UserOpenidListUrl, this.GetAccessToken(), next_openid)
	var body []byte
	body, err = util.HttpGet(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	if responseData.ErrCode != 0 {
		err = fmt.Errorf("get user_list error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//获取单个用户信息
func (this *Context) GetUserInfo(openid string) (responseData ResponseUserInfo, err error) {
	uri := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", UserInfoUrl, this.GetAccessToken(), openid)
	var body []byte
	body, err = util.HttpGet(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	if responseData.ErrCode != 0 {
		err = fmt.Errorf("get user_info error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//获取多个用户信息
func (this *Context) GetUserInfoBatchget(openids []string) (responseData ResponseUserInfoBatchget, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UserInfoBatchgetUrl, this.GetAccessToken())

	type selectUser struct {
		Openid string `json:"openid"`
		Lang   string `json:"lang"`
	}
	var openidList []selectUser
	for _, value := range openids {
		openidList = append(openidList, selectUser{
			Openid: value,
			Lang:   "zh_CN",
		})
	}
	postData := map[string]interface{}{
		"user_list": openidList,
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
		err = fmt.Errorf("get batchget_user_info error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//获取用户标签列表
func (this *Context) GetUserTagsList() (responseData ResponseUserTagsList, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UserTagsListUrl, this.GetAccessToken())
	var body []byte
	body, err = util.HttpGet(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	if responseData.ErrCode != 0 {
		err = fmt.Errorf("get user_tags_list error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//创建用户标签
func (this *Context) CreateUserTags(name string) (responseData ResponseCreateUserTags, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UserTagsCreateUrl, this.GetAccessToken())
	postData := map[string]interface{}{
		"tag": struct {
			Name string
		}{
			Name: name,
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
		err = fmt.Errorf("get create_user_tags error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//修改用户标签
func (this *Context) UpdateUserTags(id int, name string) (responseData ResponseCommon, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UserTagsUpdateUrl, this.GetAccessToken())
	postData := map[string]interface{}{
		"tag": struct {
			Id   int
			Name string
		}{
			Id:   id,
			Name: name,
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
		err = fmt.Errorf("get update_user_tags error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//删除用户标签
func (this *Context) DeleteUserTags(id int) (responseData ResponseCommon, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UserTagsDeleteUrl, this.GetAccessToken())
	postData := map[string]interface{}{
		"tag": struct {
			Id int
		}{
			Id: id,
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
		err = fmt.Errorf("get delete_user_tags error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//获取标签下的粉丝openid列表
func (this *Context) GetTagsOpenidList(id int, next_openid string) (responseData ResponseTagsOpenidList, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", TagsOpenidListUrl, this.GetAccessToken())
	postData := map[string]interface{}{
		"tagid":       id,
		"next_openid": next_openid,
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
		err = fmt.Errorf("get get_tags_openid_list error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}
