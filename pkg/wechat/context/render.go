package context

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

var plainContentType = "text/plain; charset=utf-8"
var xmlContentType = "application/xml; charset=utf-8"

//设置header中的Content-Type
func writeHeaderType(w http.ResponseWriter, value string) {
	w.Header().Set("Content-Type", value)
}

//返回数据
func (this *Context) Render(bytes []byte) {
	fmt.Println("aaa", string(bytes))
	this.Writer.WriteHeader(http.StatusOK)
	aaa, err := this.Writer.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(aaa)
}

//返回字符串数据
func (this *Context) Strings(value string) {
	writeHeaderType(this.Writer, plainContentType)
	this.Render([]byte(value))
}

//以xml的格式返回数据
func (this *Context) XML(obj interface{}) {
	writeHeaderType(this.Writer, xmlContentType)
	bytes, err := xml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	this.Render(bytes)
}
