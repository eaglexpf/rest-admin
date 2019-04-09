package controllers

import (
	"fmt"

	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/eaglexpf/rest-admin/pkg/wechat"
	"github.com/eaglexpf/rest-admin/pkg/wechat/cache"
	"github.com/eaglexpf/rest-admin/pkg/wechat/message"
	"github.com/gin-gonic/gin"
)

func getWx(c *gin.Context) *wechat.Wechat {
	//	var ch cache.Cache
	ch := cache.NewFileCache("")
	cfg := &wechat.Config{
		AppId:     pkg.LoadData.Wechat.AppID,
		AppSecret: pkg.LoadData.Wechat.AppSecret,
		Token:     pkg.LoadData.Wechat.Token,
		Writer:    c.Writer,
		Request:   c.Request,
		Cache:     ch,
	}
	wx := wechat.NewWechat(cfg)
	return wx
}

type WechatController struct {
	pkg.Controller
}

func (this *WechatController) RegisterRouter(router *gin.Engine) {
	r := router.Group("/wx")
	r.Any("/instance", this.instance)
	r.GET("/access_token", this.GetQrCode)
	r.GET("/user", this.GetUserList)
}

func (this *WechatController) instance(c *gin.Context) {

	wx := getWx(c)
	wx.HandleFunc(this.WxMessage)
	wx.Server()
}

func (this *WechatController) WxMessage(request message.RequestMessage) interface{} {
	fmt.Println(request)
	return nil
}

func (this *WechatController) GetUserList(c *gin.Context) {
	wx := getWx(c)
	//"ofks8uPRzuOGWfPKHMrkCh_kk8Ag","ofks8uIDZYNVRE_ZstjHkQu0ods8"
	list := []string{"ofks8uPRzuOGWfPKHMrkCh_kk8Ag", "ofks8uIDZYNVRE_ZstjHkQu0ods8"}
	data, err := wx.Context.GetUserInfoBatchget(list)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  err,
		"data": data,
	})
}

func (this *WechatController) GetQrCode(c *gin.Context) {
	wx := getWx(c)
	ticket, err := wx.Context.TicketShortTimeSceneStr(600, "test123123")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
	}
	if ticket.ErrCode != 0 {
		c.JSON(200, gin.H{
			"code": ticket.ErrCode,
			"msg":  ticket.ErrMsg,
		})
	}
	fmt.Println(ticket, err)
	shortData, err := wx.Context.LongToShortUrl(wx.Context.QrCodeUrl(ticket.Ticket))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
	}
	if shortData.ErrCode != 0 {
		c.JSON(200, gin.H{
			"code": shortData.ErrCode,
			"msg":  shortData.ErrMsg,
		})
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{
			"url": shortData.ShortUrl,
		},
	})
}
