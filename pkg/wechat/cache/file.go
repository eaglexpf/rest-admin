package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type File struct {
	Path string
}

type FileAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func NewFileCache(path string) *File {
	return &File{
		Path: path,
	}
}

func (this *File) Get(key string) (access_token string, err error) {
	file_path := this.Path + key
	var contents []byte
	fmt.Println(file_path)
	contents, err = ioutil.ReadFile(file_path)
	if err != nil {
		return
	}
	var file_response FileAccessToken
	err = json.Unmarshal(contents, &file_response)
	if err != nil {
		fmt.Println(err)
		return
	}
	if file_response.ExpiresIn > 0 && file_response.ExpiresIn > time.Now().Unix() {
		access_token = file_response.AccessToken
		return
	} else {
		err = errors.New(fmt.Sprintf("access_token失效，当前时间：%v , expire_in：%v", time.Now().Unix(), file_response.ExpiresIn))
	}
	return "", err
}

func (this *File) Set(key string, val string, timeout int64) error {
	file_path := this.Path + key
	//	var file_response FileAccessToken
	response := FileAccessToken{
		AccessToken: val,
		ExpiresIn:   timeout,
	}
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}
	//	fmt.Println(val.([]byte))
	err = ioutil.WriteFile(file_path, data, 0644)
	if err != nil {
		return fmt.Errorf("accessToken写入文件失败：%v", err.Error())
	}
	return nil
}

func (this *File) IsExist(key string) bool {
	return true
}

func (this *File) Delete(key string) error {
	return nil
}
