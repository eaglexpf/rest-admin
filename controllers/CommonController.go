package controllers

import (
	"fmt"

	"net/http"

	"time"

	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/eaglexpf/rest-admin/pkg/wechat"
	"github.com/eaglexpf/rest-admin/service"
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	pkg.Controller
}

func (this *CommonController) RegisterRouter(router *gin.Engine) {
	r := router.Group("/wechat")
	r.GET("/instance", this.instance)
	r.POST("/instance", this.instance)
	r.GET("/qr_code", this.getQrCode)
	r.GET("/info", this.GetInfo)
	r.GET("/running", this.GetRunning)
	r.GET("/upload_img", this.UploadImg)
	r.GET("/upload_prize", this.UploadPrize)
	r.POST("/upload_img", this.UploadImg)
	r.POST("/upload_prize", this.UploadPrize)
}

var wechatServer = wechat.Server{
	Token:     "***",
	AppID:     "***",
	AppSecert: "***",
}

func (this *CommonController) instance(c *gin.Context) {
	wechatServer.CheckSign(c)
	c.String(http.StatusOK, "%s", "success")
	var msg wechat.XmlData
	err := c.Bind(&msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch msg.MsgType {
	case "event":
		this.Events(c, msg)
		break
	}
}

func (this *CommonController) Events(c *gin.Context, msg wechat.XmlData) {
	//	wechatServer.SendText(msg.FromUserName, "已获得一张图片，<a href='https://www.baidu.com'>图片</a>")
	switch msg.Event {
	case "subscribe": //用户关注
		this.Start(c, msg)
		break
	case "unsubscribe": //取消关注
		break
		//	case "subscribe": //未关注时扫描事件
		//		this.Start(c, msg)
		//		break
	case "SCAN": //已关注时扫描事件
		this.Start(c, msg)
		break
	}
}

func (this *CommonController) Start(c *gin.Context, msg wechat.XmlData) {
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

var token = "***"

func (this *CommonController) getQrCode(c *gin.Context) {
	timestamp := c.Query("timespan")
	sign := c.Query("sign")
	var signData = make(map[string]interface{})
	signData["token"] = token
	signData["timespan"] = timestamp
	signStr := this.Sign(signData)
	if signStr != sign {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "签名错误",
			"data": map[string]interface{}{
				"sign": signStr,
			},
		})
		return
	}
	ticketData := make(map[string]interface{})
	ticketData["scene_str"] = "1231231231"
	ticket, err := wechatServer.GetTicket(600, "QR_STR_SCENE", ticketData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": map[string]interface{}{
			"qr_url": "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + ticket,
		},
	})
}

func (this *CommonController) GetInfo(c *gin.Context) {
	timestamp := c.Query("timespan")
	sign := c.Query("sign")
	var signData = make(map[string]interface{})
	signData["token"] = token
	signData["timespan"] = timestamp
	signStr := this.Sign(signData)
	if signStr != sign {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "签名错误",
			"data": map[string]interface{}{
				"sign": signStr,
			},
		})
		return
	}
	fmt.Println("aaa")
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": map[string]interface{}{
			"use_time": 180,
			"mode":     1,
			"score": map[string]interface{}{
				"Christmas": [3]int{100, 200, 300},
				"Kongfu":    [3]int{100, 200, 300},
				"Pool":      [3]int{100, 200, 300},
				"Photo":     [3]int{100, 200, 300},
				"Rain":      [3]int{100, 200, 300},
				"Firework":  [3]int{100, 200, 300},
				"ocean01":   [3]int{100, 200, 300},
				"forest01":  [3]int{100, 200, 300},
				"traffic01": [3]int{100, 200, 300},
				"Video":     [3]int{100, 200, 300},
			},
		},
	})
}

func (this *CommonController) GetRunning(c *gin.Context) {
	timestamp := c.Query("timespan")
	sign := c.Query("sign")
	var signData = make(map[string]interface{})
	signData["token"] = token
	signData["timespan"] = timestamp
	signStr := this.Sign(signData)
	if signStr != sign {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "签名错误",
			"data": "",
		})
		return
	}
	var wechatLogService service.WechatLogService
	log := wechatLogService.ExistWechatLog()
	if log.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "success",
			"data": map[string]interface{}{
				"user_id":   log.UserID,
				"log_id":    log.ID,
				"countdown": log.EndAt - time.Now().Unix(),
				"scene":     "broadcast",
				"prize":     []int{},
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "",
	})
}

func (this *CommonController) UploadPrize(c *gin.Context) {
	timestamp := c.Query("timespan")
	sign := c.Query("sign")
	prize_ids := c.Query("prize_ids")
	vr_number := c.Query("vr_number")
	var signData = make(map[string]interface{})
	signData["token"] = token
	signData["timespan"] = timestamp
	signData["vr_number"] = vr_number
	signData["prize_ids"] = prize_ids
	signStr := this.Sign(signData)
	if signStr != sign {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "签名错误",
			"data": map[string]interface{}{
				"sign": signStr,
			},
		})
		return
	}

	var wechatLogService service.WechatLogService
	log := wechatLogService.ExistWechatLog()
	if log.ID > 0 {
		var wechatUserService service.WechatUserService
		wechatUser := wechatUserService.ExistWechatUserByID(log.UserID)
		if wechatUser.ID > 0 {
			wechatServer.SendText(wechatUser.Openid, "已获得一张优惠券:"+prize_ids)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": make(map[string]string),
	})
}

func (this *CommonController) UploadImg(c *gin.Context) {
	timestamp := c.Query("timespan")
	uid := c.Query("uid")
	vr_number := c.Query("vr_number")
	img := c.Query("img")
	sign := c.Query("sign")
	var signData = make(map[string]interface{})
	signData["token"] = token
	signData["timespan"] = timestamp
	signData["uid"] = uid
	signData["vr_number"] = vr_number
	signData["img"] = img
	signStr := this.Sign(signData)
	if signStr != sign {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "签名错误",
			"data": map[string]interface{}{
				"sign": signStr,
			},
		})
		return
	}

	fmt.Println(img)
	var wechatLogService service.WechatLogService
	log := wechatLogService.ExistWechatLog()
	if log.ID > 0 {
		var wechatUserService service.WechatUserService
		wechatUser := wechatUserService.ExistWechatUserByID(log.UserID)
		if wechatUser.ID > 0 {
			wechatServer.SendText(wechatUser.Openid, "已获得一张图片，<a href='"+img+"'>图片</a>")
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": make(map[string]string),
	})
}
