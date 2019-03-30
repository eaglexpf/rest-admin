package wechat

import (
	"fmt"

	"time"

	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/eaglexpf/rest-admin/pkg/wechat"
	"github.com/eaglexpf/rest-admin/service"
)

type InitController struct {
	pkg.Controller
}

func (this *InitController) InitWechat(msg wechat.XmlData, w wechat.Server) {
	switch msg.MsgType {
	case "event":
		this.Events(msg, w)
		break
	}
}
func (this *InitController) Events(msg wechat.XmlData, w wechat.Server) {
	//	wechatServer.SendText(msg.FromUserName, "已获得一张图片，<a href='https://www.baidu.com'>图片</a>")
	switch msg.Event {
	case "subscribe": //用户关注
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
	startData["scene"] = "broadcast"
	startData["broadcast_time"] = 180
	startData["timestamp"] = time.Now().Unix()
	startData["token"] = "87d7b590a440da860adc7069fdc3c2ef"
	sign := this.Sign(startData)
	startData["sign"] = sign
	startData["mode"] = 1
	startData["uid"] = "6045CB85C000"
	res, err := this.HttpPostData("https://wechat.kayunzh.com/mfw/baoshui/Control/startCeshi", startData)
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
