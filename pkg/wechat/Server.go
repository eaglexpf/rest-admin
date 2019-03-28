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

var Wechat = Server{}

type Server struct {
	AppID     string
	AppSecert string
	Token     string
}

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
	var c pkg.Controller
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
	var c pkg.Controller
	data, err := c.HttpPostJson(uri, request)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", data["ticket"]), nil
}

func (this *Server) SendText(openid string, content string) {
	var c pkg.Controller
	token, _ := this.GetAccessToken()
	uri := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + token
	c.HttpPostJson(uri, map[string]interface{}{
		"touser":  openid,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": content,
		},
	})
}

type XmlData struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}
