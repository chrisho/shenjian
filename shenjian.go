package shenjian

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/chrisho/sd-helper/timex"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	ProxyTypeNone = 0 + iota
	ProxyTypeBasic
	ProxyTypeBetter
	ProxyTypeVip
	ProxyTypeOversea
)

const (
	HostTypeNone = 0 + iota
	HostTypeAlioss
	HostTypeQiniu
	HostTypeShenjianshou
	HostTypeUpyun

	RestfulUrl = "http://www.shenjian.io/rest/v3/"
)


type sjBase struct {
	params map[string]string
}

func newShenjian(key, secret string) sjBase {
	timestamp := timex.UnixTime()
	sign := fmt.Sprintf("%s%d%s", key, timestamp, secret)
	hash := md5.Sum([]byte(sign))
	sign = fmt.Sprintf("%x",hash)

	return sjBase{
		params: map[string]string{
			"user_key": key,
			"timestamp": fmt.Sprintf("%d", timestamp),
			"sign": sign,
		},
	}
}

type Response struct {
	Code int `json:"code"`
	Reason string `json:"reason"`
	Data interface{}
}

func (s *sjBase) request(api string, params map[string]string, method string,
		data map[string]string) (interface{}, error) {

	// 合并自定义参数
	if params != nil {
		for i, v := range params {
			if _, ok := s.params[i]; !ok {
				s.params[i] = v
			}
		}
	}
	query := ""
	for i,v := range s.params {
		query = fmt.Sprintf("%s%s=%s&", query, i, v)
	}
	query = query[:len(query)-1]
	api = fmt.Sprintf("%s%s?%s", RestfulUrl, api, query)

	// 请求数据
	var bodystr string
	if data != nil {
		var r http.Request
		r.ParseForm()
		for i, v := range data {
			r.Form.Add(i, v)
		}
		bodystr = strings.TrimSpace(r.Form.Encode())
	}

	// 请求API
	req, err := http.NewRequest(method, api, strings.NewReader(bodystr))
	if err != nil {
		return nil, err
	}
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Connection", "Keep-Alive")
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON
	var sr *Response
	if err = json.Unmarshal(b, &sr); err != nil {
		return nil, err
	}
	if sr.Code != 0 {
		return nil, errors.New(sr.Reason)
	}
	return sr.Data, nil
}

type Container struct {
	AppId int `mapstructure:"app_id"`
	Name string
	Info string
	Type string
	Status string
	TimeCreate int32 `mapstructure:"time_create"`
}
