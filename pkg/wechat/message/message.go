package message

import (
	"encoding/xml"
	"time"
)

type CommonMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

//获取的微信消息结构体
type RequestMessage struct {
	CommonMessage

	Content string `xml:Content`
	MsgId   int    `xml:"MsgId"`

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

//返回的文本消息结构体
type ResponseTextMessage struct {
	CommonMessage
	Content string `xml:"Content"`
}

//初始化一个自动回复的文本消息
func NewText(from_user_name, to_user_name, content string) *ResponseTextMessage {
	text := new(ResponseTextMessage)
	text.FromUserName = from_user_name
	text.ToUserName = to_user_name
	text.CreateTime = time.Now().Unix()
	text.MsgType = "text"
	text.Content = content
	return text
}

type ResponseImageMessage struct {
	CommonMessage
	Image struct {
		MediaId int `xml:"MediaId"`
	} `xml:"Image"`
}

func NewImage(from_user_name, to_user_name string, media_id int) *ResponseImageMessage {
	image := new(ResponseImageMessage)
	image.FromUserName = from_user_name
	image.ToUserName = to_user_name
	image.CreateTime = time.Now().Unix()
	image.MsgType = "image"
	image.Image.MediaId = media_id
	return image
}
