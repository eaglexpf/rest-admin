package context

import (
	"fmt"

	"encoding/json"

	"github.com/eaglexpf/rest-admin/pkg/wechat/util"
)

//用户列表结构体
type ResponseUserList struct {
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
type ResponseBatchUserInfo struct {
	ResponseCommon

	UserInfoList []ResponseUserInfo `json:"user_info_list"`
}

//用户标签列表结构体
type ResponseTagsList struct {
	ResponseCommon

	Tags []struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Count int64  `json:"count"`
	} `json:"tags"`
}

//创建用户标签返回的数据结构体
type ResponseTagsCreate struct {
	ResponseCommon

	Tag struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tag"`
}

//标签下的粉丝列表结构体
type ResponseUserListFromTags struct {
	ResponseCommon

	Count int64 `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

const (
	//用户列表[openid]
	UrlUserList = "https://api.weixin.qq.com/cgi-bin/user/get"
	//单个用户信息
	UrlUserInfo = "https://api.weixin.qq.com/cgi-bin/user/info"
	//批量获取多个用户信息
	UrlBatchUserInfo = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	//设置用户备注名
	UrlSetUserRemark = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark"
	//标签列表
	UrlTagsList = "https://api.weixin.qq.com/cgi-bin/tags/get"
	//创建标签
	UrlTagsCreate = "https://api.weixin.qq.com/cgi-bin/tags/create"
	//修改标签
	UrlTagsUpdate = "https://api.weixin.qq.com/cgi-bin/tags/update"
	//删除标签
	UrlTagsDelete = "https://api.weixin.qq.com/cgi-bin/tags/delete"
	//标签下的粉丝列表[openid]
	UrlUserListFromTags = "https://api.weixin.qq.com/cgi-bin/user/tag/get"
	//批量为用户打标签
	UrlBatchTaggingUser = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging"
	//批量为用户取消标签
	UrlBatchUnTaggingUser = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging"
)

//获取用户列表
func (this *Context) GetUserList(next_openid string) (responseData ResponseUserList, err error) {
	uri := fmt.Sprintf("%s?access_token=%s&next_openid=%s", UrlUserList, this.GetAccessToken(), next_openid)
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
	uri := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", UrlUserInfo, this.GetAccessToken(), openid)
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
func (this *Context) GetUserInfoBatchget(openids []string) (responseData ResponseBatchUserInfo, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UrlBatchUserInfo, this.GetAccessToken())

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
func (this *Context) GetUserTagsList() (responseData ResponseTagsList, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UrlTagsList, this.GetAccessToken())
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
func (this *Context) CreateUserTags(name string) (responseData ResponseTagsCreate, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UrlTagsCreate, this.GetAccessToken())
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
	uri := fmt.Sprintf("%s?access_token=%s", UrlTagsUpdate, this.GetAccessToken())
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
	uri := fmt.Sprintf("%s?access_token=%s", UrlTagsDelete, this.GetAccessToken())
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
func (this *Context) GetUserListFromTags(id int, next_openid string) (responseData ResponseUserListFromTags, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UrlUserListFromTags, this.GetAccessToken())
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

//批量为用户打标签
func (this *Context) BatchTaggingToUser(id int, openids []string) (responseData ResponseCommon, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UrlBatchTaggingUser, this.GetAccessToken())
	postData := map[string]interface{}{
		"tagid":       id,
		"openid_list": openids,
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
		err = fmt.Errorf("get batch_set_user_tag error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//批量为用户取消标签
func (this *Context) BatchUnTaggingToUser(id int, openids []string) (responseData ResponseCommon, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UrlBatchUnTaggingUser, this.GetAccessToken())
	postData := map[string]interface{}{
		"tagid":       id,
		"openid_list": openids,
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
		err = fmt.Errorf("get batch_un_set_user_tag error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}

//设置用户备注名
func (this *Context) SetUserRemark(openid string, remark string) (responseData ResponseCommon, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", UrlSetUserRemark, this.GetAccessToken())
	postData := map[string]interface{}{
		"openid": openid,
		"remark": remark,
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
		err = fmt.Errorf("get set_user_remark error : errcode=%v , errmsg=%v", responseData.ErrCode, responseData.ErrMsg)
		return
	}
	return
}
