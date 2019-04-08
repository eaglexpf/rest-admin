package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//json序列化（禁止转义）
func (this *Wechat) JSON(data map[string]interface{}) []byte {
	bf := bytes.NewBuffer([]byte{})
	jsonEncode := json.NewEncoder(bf)
	jsonEncode.SetEscapeHTML(false)
	jsonEncode.Encode(data)
	return bf.Bytes()
}

//发起http get请求，返回数据为json，返回json解析出来的数据
func (this *Wechat) HttpGet(uri string) ([]byte, error) {
	req, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , cause : %s", uri, err.Error())
	}
	defer req.Body.Close()
	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%s , statusCode=%v", uri, req.StatusCode)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , response nil : %s", uri, err.Error())
	}
	return body, err
	//	var jsonData map[string]interface{}
	//	err = json.Unmarshal(body, &jsonData)
	//	if err != nil {
	//		return nil, fmt.Errorf("http get error : uri=%s , json decode error : %s", uri, err.Error())
	//	}
	//	return jsonData, err
}

//发起http post请求，请求数据为json，返回数据为json，返回json解析出来的数据
func (this *Wechat) HttpPostJson(uri string, data map[string]interface{}) (map[string]interface{}, error) {
	msg := this.JSON(data)
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(msg))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , cause : %s", uri, err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , cause : %s", uri, err.Error())
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , cause : %s", uri, err.Error())
	}
	return jsonData, err
}

//发起http post请求，请求数据为x-www-form-urlencoded，返回数据为json，返回json解析出来的数据
func (this *Wechat) HttpPostData(uri string, data map[string]interface{}) (map[string]interface{}, error) {
	v := url.Values{}
	for key, value := range data {
		v.Add(key, fmt.Sprintf("%v", value))
	}
	msg := v.Encode()
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(msg)))
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , cause : %s", uri, err.Error())
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , cause : %s", uri, err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , cause : %s", uri, err.Error())
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("http get error : uri=%s , cause : %s", uri, err.Error())
	}
	return jsonData, err
}
