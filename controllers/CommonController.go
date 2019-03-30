package controllers

import (
	"fmt"

	"net/http"

	"time"

	c_wechat "github.com/eaglexpf/rest-admin/controllers/wechat"
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
	r.GET("/advert", this.getAdvert)
	r.GET("/prize", this.getPrize)
}

var wechatServer = wechat.Server{
	Token:     pkg.LoadData.Wechat.Token,
	AppID:     pkg.LoadData.Wechat.AppID,
	AppSecert: pkg.LoadData.Wechat.AppSecret,
}

//接入微信公众号
func (this *CommonController) instance(c *gin.Context) {
	if !wechatServer.CheckSign(c) {
		return
	}
	c.String(http.StatusOK, "%s", "success")
	var msg wechat.XmlData
	err := c.Bind(&msg)
	if err != nil {
		c.String(http.StatusOK, "%s", err.Error())
		return
	}
	var wechatInit = c_wechat.InitController{}
	wechatInit.InitWechat(msg, wechatServer)

}

var token = "723f8a8bdb6309da6f26e5ab9d8ea005"

/**
 * @apiDefine API 接口：
 */
/**
 * @apiDefine AUTH 被动接口
 */
/**
 * @api {get} /wechat/qr_code 获取二维码
 * @apiDescription 获取公众号二维码地址
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {string} sign 签名
 * @apiParam {string} timespan 时间戳
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 * @apiSuccess {string} data.qr_url 二维码图片地址
 *
 **/
//获取二维码
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

/**
 * @api {get} /wechat/info 获取基础信息
 * @apiDescription 获取基础信息
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {string} sign 签名
 * @apiParam {string} timespan 时间戳
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 * @apiSuccess {string} data.mode 游戏模式【1、正常游戏轮播；4、切红包】
 * @apiSuccess {string} data.use_time 一局游戏时长
 * @apiSuccess {object} data.score 场景积分段
 *
 **/
//获取基础信息
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
			"mode":     4,
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

/**
 * @api {get} /wechat/running 获取是否有正在进行中的游戏
 * @apiDescription 获取是否有正在进行中的游戏
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {string} sign 签名
 * @apiParam {string} timespan 时间戳
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 * @apiSuccess {int} data.user_id 用户id
 * @apiSuccess {int} data.log_id 游戏id
 * @apiSuccess {int} data.countdown 还剩多少秒
 * @apiSuccess {string} data.scene 场景值
 * @apiSuccess {array} data.prize 已获取的优惠券
 *
 **/
//获取是否有正在进行中的游戏
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
			"data": map[string]interface{}{
				"sign": signStr,
			},
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

/**
 * @api {get} /wechat/upload_prize 上传优惠券
 * @apiDescription 上传优惠券
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {string} prize_ids 优惠券id集合
 * @apiParam {string} vr_number 游戏id
 * @apiParam {string} sign 签名
 * @apiParam {string} timespan 时间戳
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
//上传优惠券
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
			var wechatUserPrizeService service.WechatUserPrizeService
			wechatUserPrizeService.InsertUserPrize(log.UserID, log.ID, prize_ids)
			wechatServer.SendText(wechatUser.Openid, "已获得一张优惠券:"+prize_ids)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": make(map[string]string),
	})
}

/**
 * @api {get} /wechat/upload_img 上传图片
 * @apiDescription 上传图片
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {string} uid 用户id
 * @apiParam {string} vr_number 游戏id
 * @apiParam {string} img 图片地址
 * @apiParam {string} sign 签名
 * @apiParam {string} timespan 时间戳
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
//上传图片
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

/**
 * @api {get} /wechat/advert 获取广告logo
 * @apiDescription 获取广告logo
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {string} sign 签名
 * @apiParam {string} timespan 时间戳
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 * @apiSuccess {string} data.logo logo图片地址[515*126]
 * @apiSuccess {array} data.advert 广告图片集合
 * @apiSuccess {int} data.advert.id 广告图片id
 * @apiSuccess {string} data.advert.img 广告图片地址[648*767]
 *
 **/
//获取广告logo
func (this *CommonController) getAdvert(c *gin.Context) {
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
	var advertService service.AdvertService
	logo := advertService.GetLogo()
	advert := advertService.GetAdvert()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": map[string]interface{}{
			"logo":   logo.Img,
			"advert": advert,
		},
	})
}

/**
 * @api {get} /wechat/prize 获取优惠券列表
 * @apiDescription 获取优惠券列表
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {string} scene 场景值字符串【多个场景以英文,拼接】
 * @apiParam {string} sign 签名
 * @apiParam {string} timespan 时间戳
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 * @apiSuccess {array} data.scene 场景值对应数组
 * @apiSuccess {int} data.scene.id 优惠券id
 * @apiSuccess {string} data.scene.name 优惠券名称
 * @apiSuccess {string} data.scene.unit 优惠券单位（张）
 * @apiSuccess {string} data.scene.unity_url 优惠券图片地址
 * @apiSuccess {string} data.scene.icon_url_active 优惠券icon地址（明亮）
 * @apiSuccess {string} data.scene.icon_url_inactive 优惠icon地址（暗色）
 * @apiSuccess {int} data.scene.num 优惠券数量
 * @apiSuccess {int} data.scene.type 优惠券类型（暂未使用，请统一使用默认值1）
 * @apiSuccess {int} data.scene.prob 优惠券概率（暂未使用，请统一使用默认值1）
 *
 **/
//获取优惠券列表
func (this *CommonController) getPrize(c *gin.Context) {
	timestamp := c.Query("timespan")
	sign := c.Query("sign")
	scene := c.Query("scene")
	var signData = make(map[string]interface{})
	signData["scene"] = scene
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

	var prizeService service.PrizeService
	data := prizeService.GetList(scene)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": data,
	})
}
