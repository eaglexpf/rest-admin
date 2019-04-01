define({ "api": [
  {
    "type": "get",
    "url": "/wechat/advert",
    "title": "获取广告logo",
    "description": "<p>获取广告logo</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "sign",
            "description": "<p>签名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timespan",
            "description": "<p>时间戳</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.logo",
            "description": "<p>logo图片地址[515*126]</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "data.advert",
            "description": "<p>广告图片集合</p>"
          },
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "data.advert.id",
            "description": "<p>广告图片id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.advert.img",
            "description": "<p>广告图片地址[648*767]</p>"
          }
        ]
      }
    },
    "filename": "controllers/CommonController.go",
    "groupTitle": "被动接口",
    "name": "GetWechatAdvert",
    "sampleRequest": [
      {
        "url": "https://ssl.xupengfei.net/wechat/advert"
      }
    ]
  },
  {
    "type": "get",
    "url": "/wechat/create_prize",
    "title": "创建优惠券",
    "description": "<p>创建优惠券</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>优惠券名称</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "unit",
            "description": "<p>优惠券单位</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "img_url",
            "description": "<p>优惠券地址</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "icon_on",
            "description": "<p>优惠券激活地址</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "icon_off",
            "description": "<p>优惠券关闭地址</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "num",
            "description": "<p>优惠券数量</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "valid_start",
            "description": "<p>优惠券开始时间</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "valid_end",
            "description": "<p>优惠券结束时间</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          }
        ]
      }
    },
    "filename": "controllers/CommonController.go",
    "groupTitle": "被动接口",
    "name": "GetWechatCreate_prize",
    "sampleRequest": [
      {
        "url": "https://ssl.xupengfei.net/wechat/create_prize"
      }
    ]
  },
  {
    "type": "get",
    "url": "/wechat/info",
    "title": "获取基础信息",
    "description": "<p>获取基础信息</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "sign",
            "description": "<p>签名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timespan",
            "description": "<p>时间戳</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.mode",
            "description": "<p>游戏模式【1、正常游戏轮播；4、切红包】</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.use_time",
            "description": "<p>一局游戏时长</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data.score",
            "description": "<p>场景积分段</p>"
          }
        ]
      }
    },
    "filename": "controllers/CommonController.go",
    "groupTitle": "被动接口",
    "name": "GetWechatInfo",
    "sampleRequest": [
      {
        "url": "https://ssl.xupengfei.net/wechat/info"
      }
    ]
  },
  {
    "type": "get",
    "url": "/wechat/prize",
    "title": "获取优惠券列表",
    "description": "<p>获取优惠券列表</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "scene",
            "description": "<p>场景值字符串【多个场景以英文,拼接】</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "sign",
            "description": "<p>签名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timespan",
            "description": "<p>时间戳</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "data.scene",
            "description": "<p>场景值对应数组</p>"
          },
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "data.scene.id",
            "description": "<p>优惠券id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.scene.name",
            "description": "<p>优惠券名称</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.scene.unit",
            "description": "<p>优惠券单位（张）</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.scene.unity_url",
            "description": "<p>优惠券图片地址</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.scene.icon_url_active",
            "description": "<p>优惠券icon地址（明亮）</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.scene.icon_url_inactive",
            "description": "<p>优惠icon地址（暗色）</p>"
          },
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "data.scene.num",
            "description": "<p>优惠券数量</p>"
          },
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "data.scene.type",
            "description": "<p>优惠券类型（暂未使用，请统一使用默认值1）</p>"
          },
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "data.scene.prob",
            "description": "<p>优惠券概率（暂未使用，请统一使用默认值1）</p>"
          }
        ]
      }
    },
    "filename": "controllers/CommonController.go",
    "groupTitle": "被动接口",
    "name": "GetWechatPrize",
    "sampleRequest": [
      {
        "url": "https://ssl.xupengfei.net/wechat/prize"
      }
    ]
  },
  {
    "type": "get",
    "url": "/wechat/qr_code",
    "title": "获取二维码",
    "description": "<p>获取公众号二维码地址</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "sign",
            "description": "<p>签名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timespan",
            "description": "<p>时间戳</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.qr_url",
            "description": "<p>二维码图片地址</p>"
          }
        ]
      }
    },
    "filename": "controllers/CommonController.go",
    "groupTitle": "被动接口",
    "name": "GetWechatQr_code",
    "sampleRequest": [
      {
        "url": "https://ssl.xupengfei.net/wechat/qr_code"
      }
    ]
  },
  {
    "type": "get",
    "url": "/wechat/running",
    "title": "获取是否有正在进行中的游戏",
    "description": "<p>获取是否有正在进行中的游戏</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "sign",
            "description": "<p>签名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timespan",
            "description": "<p>时间戳</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          },
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "data.user_id",
            "description": "<p>用户id</p>"
          },
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "data.log_id",
            "description": "<p>游戏id</p>"
          },
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "data.countdown",
            "description": "<p>还剩多少秒</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "data.scene",
            "description": "<p>场景值</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "data.prize",
            "description": "<p>已获取的优惠券</p>"
          }
        ]
      }
    },
    "filename": "controllers/CommonController.go",
    "groupTitle": "被动接口",
    "name": "GetWechatRunning",
    "sampleRequest": [
      {
        "url": "https://ssl.xupengfei.net/wechat/running"
      }
    ]
  },
  {
    "type": "get",
    "url": "/wechat/upload_img",
    "title": "上传图片",
    "description": "<p>上传图片</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uid",
            "description": "<p>用户id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "vr_number",
            "description": "<p>游戏id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "img",
            "description": "<p>图片地址</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "sign",
            "description": "<p>签名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timespan",
            "description": "<p>时间戳</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          }
        ]
      }
    },
    "filename": "controllers/CommonController.go",
    "groupTitle": "被动接口",
    "name": "GetWechatUpload_img",
    "sampleRequest": [
      {
        "url": "https://ssl.xupengfei.net/wechat/upload_img"
      }
    ]
  },
  {
    "type": "get",
    "url": "/wechat/upload_prize",
    "title": "上传优惠券",
    "description": "<p>上传优惠券</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "prize_ids",
            "description": "<p>优惠券id集合</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "vr_number",
            "description": "<p>游戏id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "sign",
            "description": "<p>签名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timespan",
            "description": "<p>时间戳</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          }
        ]
      }
    },
    "filename": "controllers/CommonController.go",
    "groupTitle": "被动接口",
    "name": "GetWechatUpload_prize",
    "sampleRequest": [
      {
        "url": "https://ssl.xupengfei.net/wechat/upload_prize"
      }
    ]
  }
] });
