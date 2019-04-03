package pkg

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
)

type Controller struct {
	Data interface{}
}

//json序列化（禁止转义）
func (Controller) JSON(data map[string]interface{}) []byte {
	bf := bytes.NewBuffer([]byte{})
	jsonEncode := json.NewEncoder(bf)
	jsonEncode.SetEscapeHTML(false)
	jsonEncode.Encode(data)
	return bf.Bytes()
}

//发起http get请求，返回数据为json，返回json解析出来的数据
func (Controller) HttpGet(uri string) (map[string]interface{}, error) {
	req, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	fmt.Println(jsonData)
	return jsonData, err
}

//发起http post请求，请求数据为json，返回数据为json，返回json解析出来的数据
func (this *Controller) HttpPostJson(uri string, data map[string]interface{}) (map[string]interface{}, error) {
	msg := this.JSON(data)
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(msg))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	return jsonData, err
}

//发起http post请求，请求数据为x-www-form-urlencoded，返回数据为json，返回json解析出来的数据
func (this *Controller) HttpPostData(uri string, data map[string]interface{}) (map[string]interface{}, error) {
	v := url.Values{}
	for key, value := range data {
		v.Add(key, fmt.Sprintf("%v", value))
	}
	msg := v.Encode()
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(msg)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	return jsonData, err
}

//获取签名字符串 按参数名升序排序 将排序后的参数用&符号链接 将获得的字符串使用md5加密
func (this *Controller) Sign(data map[string]interface{}) string {
	var sslice []string
	for key, _ := range data {
		sslice = append(sslice, key)
	}
	sort.Strings(sslice)
	v := url.Values{}
	for _, key := range sslice {
		v.Add(key, fmt.Sprintf("%v", data[key]))
	}
	body := v.Encode()
	str, _ := url.QueryUnescape(body)
	return this.Md5(str)
}

//字符串MD5加密
func (Controller) Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//获取文件修改时间 返回unix时间戳
func (Controller) GetFileModTime(path string) (int64, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return 0, err
	}

	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return fi.ModTime().Unix(), nil
}

func (Controller) InterfaceToInt(data interface{}) (int, error) {
	switch data.(type) {
	case int:
		return data.(int), nil
	case string:
		return strconv.Atoi(data.(string))
	case int64:
		return int(data.(int64)), nil
	case float64:
		strInt64 := strconv.FormatFloat(data.(float64), 'f', -1, 64)
		return strconv.Atoi(strInt64)
	}
	return 0, errors.New("非int类型")
}
