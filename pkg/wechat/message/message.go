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

type ResponseMessage struct {
	MsgType string
	MsgData interface{}
}

type ResponseTextMessage struct {
	CommonMessage
	Content string `xml:"Content"`
}

func NewText(from_user_name, to_user_name, content string) *ResponseTextMessage {
	text := new(ResponseTextMessage)
	text.FromUserName = from_user_name
	text.ToUserName = to_user_name
	text.CreateTime = time.Now().Unix()
	text.MsgType = "text"
	text.Content = content
	return text
}
