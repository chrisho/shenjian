package shenjian

import (
	"encoding/base64"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

const (
	ApiMoneyInfo = "user/money"
	ApiNodeInfo = "user/node"
	ApiAppList = "app/list"
	ApiCrawlerList = "crawler/list"
	ApiCreateCrawler = "crawler/create"
)

type sjService struct {
	sjBase
}

func NewSjService(key, secret string) *sjService {
	s := newShenjian(key, secret)
	return &sjService{
		s,
	}
}
func (s *sjService) GetMoneyInfo() (map[string]interface{}, error) {
	if res, err := s.request(ApiMoneyInfo,nil,"", nil); err != nil {
		return nil, err
	} else {
		datas := res.(map[string]interface{})
		return datas, nil
	}
}

func (s *sjService) GetNodeInfo() (map[string]interface{}, error) {
	if res, err := s.request(ApiNodeInfo, nil,"", nil); err != nil {
		return nil, err
	} else {
		datas := res.(map[string]interface{})
		return datas, nil
	}
}

func (s *sjService) GetAppList(page, pageSize int) ([]Container, int, int, error) {
	res, err := s.request(ApiAppList, nil,http.MethodPost, map[string]string{
		"page": fmt.Sprint(page),
		"pageSize":  fmt.Sprint(pageSize),
	})
	if err != nil {
		return nil, 0, 0, err
	}
	datas := res.(map[string]interface{})


	var apps struct {
		List []Container
		Page int
		Total int
	}
	if err = mapstructure.Decode(datas, &apps); err != nil {
		return nil, 0, 0, err
	}

	return apps.List, apps.Page, apps.Total, nil
}


func (s *sjService) GetCrawlerList(page, pageSize int) ([]Container, int, int, error) {
	res, err := s.request(ApiCrawlerList, nil, http.MethodPost, map[string]string{
		"page": fmt.Sprint(page),
		"pageSize":  fmt.Sprint(pageSize),
	})
	if err != nil {
		return nil, 0, 0, err
	}
	datas := res.(map[string]interface{})

	var crawlers struct {
		List []Container
		Page int
		Total int
	}
	if err = mapstructure.Decode(datas, &crawlers); err != nil {
		return nil, 0, 0, err
	}
	return crawlers.List, crawlers.Page, crawlers.Total, nil
}

func (s *sjService) GetCreateCrawler(appName, appInfo, code string) (*Container, error) {
	res, err := s.request(ApiCreateCrawler, nil, http.MethodPost, map[string]string{
		"app_name": appName,
		"app_info":  appInfo,
		"code": base64.StdEncoding.EncodeToString([]byte(code)),
	})
	if err != nil {
		return nil, err
	}
	datas := res.(map[string]interface{})
	datas["app_id"], _ = strconv.Atoi(fmt.Sprint(datas["app_id"]))
	var crawler *Container
	if err = mapstructure.Decode(datas, &crawler); err != nil {
		return nil, err
	}

	return crawler, err
}
