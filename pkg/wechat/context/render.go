package context

import (
	"encoding/xml"
	"net/http"
)

var plainContentType = "text/plain; charset=utf-8"
var xmlContentType = "application/xml; charset=utf-8"

func writeHeaderType(w http.ResponseWriter, value string) {
	w.Header().Set("Content-Type", value)
}

func (this *Context) Render(bytes []byte) {
	this.Writer.WriteHeader(http.StatusOK)
	_, err := this.Writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func (this *Context) Strings(value string) {
	writeHeaderType(this.Writer, plainContentType)
	this.Render([]byte(value))
}

func (this *Context) XML(obj interface{}) {
	writeHeaderType(this.Writer, xmlContentType)
	bytes, err := xml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	this.Render(bytes)
}
