package wechat

import (
	"fmt"

	"net/http"

	"sort"

	"strings"

	"crypto/sha1"

	"io/ioutil"

	"time"

	"log"
	//	"os"

	"errors"

	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/gin-gonic/gin"
)

var c pkg.Controller

type Server struct {
	AppID     string
	AppSecert string
	Token     string
}

func New(app_id, app_secret, token string) Server {
	return Server{
		AppID:     app_id,
		AppSecert: app_secret,
		Token:     token,
	}
}

//验证签名
func (this *Server) CheckSign(c *gin.Context) bool {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	var sortString []string
	sortString = append(sortString, timestamp, nonce, this.Token)
	sort.Strings(sortString)

	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(strings.Join(sortString, "")))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	sign := fmt.Sprintf("%x", bs)
	if sign != signature {
		c.String(http.StatusOK, "%s", "签名错误")
		return false
	}
	if echostr != "" {
		c.String(http.StatusOK, "%s", echostr)
		return false
	}
	return true
}

//获取微信AccessToken
func (this *Server) GetAccessToken() (string, error) {
	path := "accessToken"
	lastTime, err := c.GetFileModTime(path)
	if time.Now().Unix()-lastTime > 6000 || err != nil {
		uri := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + this.AppID + "&secret=" + this.AppSecert
		response, err := c.HttpGet(uri)
		if err != nil {
			return "", err
		}
		if errcode, ok := response["errcode"]; ok {
			return "", errors.New(fmt.Sprintf("%s%s", "获取accessToken失败，错误编号：", errcode))
		}
		accessToken, ok := response["access_token"]
		if !ok {
			return "", errors.New(fmt.Sprintf("%s", "获取accessToken失败，返回数据没有access_token参数"))
		}
		value, ok := accessToken.(string)
		if !ok {
			return "", errors.New(fmt.Sprintf("%s", "获取accessToken失败，返回access_token参数无法转成字符串"))
		}
		err = ioutil.WriteFile(path, []byte(value), 0644)
		if err != nil {
			log.Panicln("accessToken写入文件失败：", err.Error())
		}
		return value, err
	}
	contents, err := ioutil.ReadFile(path)
	if err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		return result, nil
	}
	return "", err

}

//获取ticket
func (this *Server) GetTicket(expire_seconds int, action_name string, action_info map[string]interface{}) (string, error) {
	token, err := this.GetAccessToken()
	if err != nil {
		return "", err
	}
	uri := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + token
	var request = make(map[string]interface{})
	request["expire_seconds"] = expire_seconds
	request["action_name"] = action_name
	request["action_info"] = action_info
	data, err := c.HttpPostJson(uri, request)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", data["ticket"]), nil
}

//发送文本消息
func (this *Server) SendText(openid string, content string) (map[string]interface{}, error) {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + token
	return c.HttpPostJson(uri, map[string]interface{}{
		"touser":  openid,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": content,
		},
	})
}

//发送图片消息
func (this *Server) SendImg(openid string, media_id int) (map[string]interface{}, error) {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + token
	return c.HttpPostJson(uri, map[string]interface{}{
		"touser":  openid,
		"msgtype": "image",
		"image": map[string]interface{}{
			"media_id": media_id,
		},
	})
}

//发送语音消息
func (this *Server) SendVoice(openid string, media_id int) (map[string]interface{}, error) {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + token
	return c.HttpPostJson(uri, map[string]interface{}{
		"touser":  openid,
		"msgtype": "voice",
		"voice": map[string]interface{}{
			"media_id": media_id,
		},
	})
}

//发送视频消息
func (this *Server) SendVideo(openid, title, description string, media_id, thumb_media_id int) (map[string]interface{}, error) {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + token
	return c.HttpPostJson(uri, map[string]interface{}{
		"touser":  openid,
		"msgtype": "video",
		"video": map[string]interface{}{
			"media_id":       media_id,
			"thumb_media_id": thumb_media_id,
			"title":          title,
			"description":    description,
		},
	})
}

//发送音乐消息
func (this *Server) SendMusic(openid, title, description, musicurl, hqmusicurl string, thumb_media_id int) (map[string]interface{}, error) {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + token
	return c.HttpPostJson(uri, map[string]interface{}{
		"touser":  openid,
		"msgtype": "music",
		"music": map[string]interface{}{
			"title":          title,
			"description":    description,
			"musicurl":       musicurl,
			"hqmusicurl":     hqmusicurl,
			"thumb_media_id": thumb_media_id,
		},
	})
}

//发送图文消息（点击跳转到外链）
func (this *Server) SendNews(openid string, articles []struct {
	title       string
	description string
	url         string
	picurl      string
}) (map[string]interface{}, error) {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + token
	return c.HttpPostJson(uri, map[string]interface{}{
		"touser":  openid,
		"msgtype": "news",
		"news": map[string]interface{}{
			"articles": articles,
		},
	})
}

func (this *Server) SendMpNews(openid string, media_id int) (map[string]interface{}, error) {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + token
	return c.HttpPostJson(uri, map[string]interface{}{
		"touser":  openid,
		"msgtype": "mpnews",
		"mpnews": map[string]interface{}{
			"media_id": media_id,
		},
	})
}

//创建菜单
func (this *Server) MenuCreate(data map[string]interface{}) error {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + token
	res, err := c.HttpPostJson(uri, data)
	fmt.Println(res)
	if err != nil {
		return err
	}
	errcode, ok := res["errcode"]
	if !ok {
		return errors.New("创建失败，没有返回errcode参数")
	}
	code, err := c.InterfaceToInt(errcode)
	if err != nil {
		return errors.New(fmt.Sprintf("创建失败，返回的errcode参数非int类型%v", errcode))
	}
	if code != 0 {
		return errors.New(fmt.Sprintf("创建失败，errcode:%v", code))
	}
	return nil
}

//查询菜单
func (this *Server) MenuSelect() (map[string]interface{}, error) {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=" + token
	res, err := c.HttpGet(uri)
	if err != nil {
		return make(map[string]interface{}), err
	}
	errcode, ok := res["errcode"]
	if !ok {
		return res, errors.New("查询失败，没有返回errcode参数")
	}
	code, err := c.InterfaceToInt(errcode)
	if err != nil {
		return make(map[string]interface{}), errors.New(fmt.Sprintf("查询失败，返回的errcode参数非int类型%v", errcode))
	}
	if code == 46003 {
		return make(map[string]interface{}), errors.New("查询失败，不存在的菜单数据")
	}
	return make(map[string]interface{}), errors.New("查询失败，失败原因未知")
}

//清除菜单
func (this *Server) MenuClear() error {
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" + token
	res, err := c.HttpGet(uri)
	if err != nil {
		return err
	}
	errcode, ok := res["errcode"]
	if !ok {
		return errors.New("清除失败，没有返回errcode参数")
	}
	code, err := c.InterfaceToInt(errcode)
	if err != nil {
		return errors.New(fmt.Sprintf("清除失败，返回的errcode参数非int类型%v", errcode))
	}
	if code != 0 {
		return errors.New(fmt.Sprintf("清除失败，errcode:%v", code))
	}
	return nil
}

//被动回复文本消息
func (this *Server) ResponseText(openid, content string) interface{} {
	type xml struct {
		ToUserName   string `xml:"ToUserName"`
		FromUserName string `xml:"FromUserName"`
		CreateTime   int64  `xml:"CreateTime"`
		MsgType      string `xml:"MsgType"`
		Content      string `xml:"Content"`
	}
	var data = xml{
		ToUserName: openid,
		CreateTime: time.Now().Unix(),
		MsgType:    "text",
		Content:    content,
	}
	return data
}

//被动回复图片消息
func (this *Server) ResponseImage(openid string, media_id int) interface{} {
	type xml struct {
		ToUserName   string `xml:"ToUserName"`
		FromUserName string `xml:"FromUserName"`
		CreateTime   int64  `xml:"CreateTime"`
		MsgType      string `xml:"MsgType"`
		MediaId      int    `xml:"MediaId"`
	}
	var data = xml{
		ToUserName: openid,
		CreateTime: time.Now().Unix(),
		MsgType:    "image",
		MediaId:    media_id,
	}
	return data
}

//被动回复语音消息
func (this *Server) ResponseVoice(openid string, media_id int) interface{} {
	type xml struct {
		ToUserName   string `xml:"ToUserName"`
		FromUserName string `xml:"FromUserName"`
		CreateTime   int64  `xml:"CreateTime"`
		MsgType      string `xml:"MsgType"`
		MediaId      int    `xml:"MediaId"`
	}
	var data = xml{
		ToUserName: openid,
		CreateTime: time.Now().Unix(),
		MsgType:    "voice",
		MediaId:    media_id,
	}
	return data
}

//被动回复视频消息
func (this *Server) ResponseVideo(openid, title, description string, media_id int) interface{} {
	type video struct {
		MediaId     int    `xml:"MediaId"`
		Title       string `xml:"Title"`
		Description string `xml:"Description"`
	}
	type xml struct {
		ToUserName   string `xml:"ToUserName"`
		FromUserName string `xml:"FromUserName"`
		CreateTime   int64  `xml:"CreateTime"`
		MsgType      string `xml:"MsgType"`
		Video        video  `xml:"Video"`
	}
	var data = xml{
		ToUserName: openid,
		CreateTime: time.Now().Unix(),
		MsgType:    "video",
		Video: video{
			MediaId:     media_id,
			Title:       title,
			Description: description,
		},
	}
	return data
}

//被动回复音乐消息
func (this *Server) ResponseMusic(openid, title, description, musicurl, hqmusicurl string, thumb_media_id int) interface{} {
	type music struct {
		Title        string `xml:"Title"`
		Description  string `xml:"Description"`
		MusicUrl     string `xml:"MusicUrl"`
		HQMusicUrl   string `xml:"HQMusicUrl"`
		ThumbMediaId int    `xml:"ThumbMediaId"`
	}
	type xml struct {
		ToUserName   string `xml:"ToUserName"`
		FromUserName string `xml:"FromUserName"`
		CreateTime   int64  `xml:"CreateTime"`
		MsgType      string `xml:"MsgType"`
		Music        music  `xml:"Music"`
	}
	var data = xml{
		ToUserName: openid,
		CreateTime: time.Now().Unix(),
		MsgType:    "music",
		Music: music{
			Title:        title,
			Description:  description,
			MusicUrl:     musicurl,
			HQMusicUrl:   hqmusicurl,
			ThumbMediaId: thumb_media_id,
		},
	}
	return data
}

//func (this *Server) ResponseNew(openid string, article []map[string]interface{}) (interface{}, error) {
//	type item struct {
//		Title       string `xml:"Title"`
//		Description string `xml:"Description"`
//		PicUrl      string `xml:"PicUrl"`
//		Url         string `xml:"Url"`
//	}
//	type xml struct {
//		ToUserName   string `xml:"ToUserName"`
//		FromUserName string `xml:"FromUserName"`
//		CreateTime   int64  `xml:"CreateTime"`
//		MsgType      string `xml:"MsgType"`
//		ArticleCount int    `xml:"ArticleCount"`
//		Articles     []item `xml:"Articles"`
//	}
//	items := make([]item)
//	for _, value := range article {
//		if title, ok := value["title"]; !ok {
//			return nil, errors.New("没有title")
//		}
//		items = append(items, item{
//			Title:       value["title"],
//			Description: value["description"],
//			PicUrl:      value["pic_url"],
//			Url:         value["url"],
//		})
//	}
//	var data = xml{
//		ToUserName:   openid,
//		CreateTime:   time.Now().Unix(),
//		MsgType:      "news",
//		ArticleCount: len(items),
//		Articles:     items,
//	}
//	return data, nil
//}

type XmlData struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:Content`
	MsgId        int    `xml:"MsgId"`

	PicUrl  string `xml:"PicUrl"`
	MediaId int    `xml:"MediaId"`

	Format      string `xml:"Format"`
	Recognition string `xml:"Recognition"`

	ThumbMediaId string `xml:"ThumbMediaId"`

	Location_X string `xml:"Location_X"`
	Location_Y string `xml:"Location_Y"`
	Scale      string `xml:"Scale"`
	Label      string `xml:"Label"`
	Latitude   string `xml:"Latitude"`
	Longitude  string `xml:"Longitude"`
	Precision  string `xml:"Precision"`

	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	Url         string `xml:"Url"`

	Event    string `xml:"Event"`
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

type XmlImgData struct {
	MediaId int `xml:"MediaId"`
}

type XmlVoiceData struct {
	MediaId int `xml:"MediaId"`
}

type XmlVideoData struct {
	MediaId     int    `xml:"MediaId"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

type XmlMusicData struct {
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	MusicUrl     string `xml:"MusicUrl"`
	HQMusicUrl   string `xml:"HQMusicUrl"`
	ThumbMediaId int    `xml:"ThumbMediaId"`
}

type XmlArticleItemData struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

type XmlNewsData struct {
	Item []XmlArticleItemData `xml:"item"`
}
