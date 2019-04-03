package wechat

import (
	"fmt"

	"time"

	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/eaglexpf/rest-admin/pkg/wechat"
	"github.com/eaglexpf/rest-admin/service"
	"github.com/gin-gonic/gin"
)

type InitController struct {
	pkg.Controller
}

func (this *InitController) InitWechat(msg wechat.XmlData, w wechat.Server, c *gin.Context) {
	fmt.Println(msg)
	switch msg.MsgType {
	case "event":
		this.Events(msg, w)
		break
	}
	var data = w.ResponseText(msg.FromUserName, "Hello World!")
	//	c.Writer.
	c.XML(200, data)
}
func (this *InitController) Events(msg wechat.XmlData, w wechat.Server) {
	//	wechatServer.SendText(msg.FromUserName, "已获得一张图片，<a href='https://www.baidu.com'>图片</a>")
	switch msg.Event {
	case "subscribe": //用户关注 //未关注时扫描事件
		this.Start(msg, w)
		break
	case "unsubscribe": //取消关注
		break
		//	case "subscribe": //未关注时扫描事件
		//		this.Start(c, msg)
		//		break
	case "SCAN": //已关注时扫描事件
		this.Start(msg, w)
		break
	}
}

//开始游戏，用户扫描事件
func (this *InitController) Start(msg wechat.XmlData, wechatServer wechat.Server) {
	var wechatUserService service.WechatUserService
	var wechatLogService service.WechatLogService
	wechatUser := wechatUserService.ExistWechatUserByOpenid(msg.FromUserName)
	if wechatUser.ID <= 0 {
		wechatUserService.CreateWechatUser(msg.FromUserName)
		wechatUser = wechatUserService.ExistWechatUserByOpenid(msg.FromUserName)
	}
	log := wechatLogService.ExistWechatLog()
	if log.ID > 0 {
		fmt.Println("游戏正在进行中")
		wechatServer.SendText(msg.FromUserName, "游戏正在进行中,请等待游戏结束后再扫码")
		return
	}
	hasTicket := wechatLogService.ExistWechatLogByTicket(msg.Ticket)
	if hasTicket {
		fmt.Println("二维码已被使用")
		wechatServer.SendText(msg.FromUserName, "二维码已被使用，请使用最新的二维码")
		return
	}
	wechatLogService.CreateWechatLog(wechatUser.ID, msg.Ticket)
	log = wechatLogService.ExistWechatLog()
	fmt.Println("success")
	startData := make(map[string]interface{})
	startData["user_id"] = wechatUser.ID
	startData["log_id"] = log.ID
	startData["timestamp"] = time.Now().Unix()
	startData["account"] = pkg.LoadData.Wechat.Account
	startData["token"] = pkg.LoadData.Wechat.ApiToken
	sign := this.Sign(startData)
	startData["sign"] = sign
	res, err := this.HttpPostData("https://wechat.kayunzh.com/mfw/baoshui/Api/start", startData)
	fmt.Println("asdasdadad", res)
	if err != nil {
		wechatServer.SendText(msg.FromUserName, err.Error())
		return
	}
	if res["code"] != 0 {
		wechatServer.SendText(msg.FromUserName, fmt.Sprintf("%v", res["msg"]))
		return
	}
	wechatServer.SendText(msg.FromUserName, "游戏已开始")
}
