package shenjian

import "fmt"

type sjCrawler struct {
	sjBase
	AppId int
}

func NewSjCrawer(key, secret string, appId int) *sjCrawler {
	s := newShenjian(key, secret)
	return &sjCrawler{
		sjBase: s,
		AppId: appId,
	}
}

func (s *sjCrawler) request(api string, params map[string]string, method string,
	data map[string]string) (interface{}, error) {

	api = fmt.Sprintf("crawler/%d%s", s.AppId, api)
	return s.sjBase.request(api, params, method, data)
}

func (s *sjCrawler) Edit(appName, appInfo string) error {

	return nil
}

func (s *sjCrawler) ConfigProxy(proxyType int) error {

	return nil
}

func (s *sjCrawler) ConfigHost(hostType int, image, text, audio, video, application bool) error {
	return nil
}

func (s *sjCrawler) ConfigCustom(bodystr string) error {
	return nil
}

type StartTimer struct {
	// 通用|Daily
	DateStart string
	DateEnd string
	TimeStart string
	TimeEnd string
	// Once
	OnceDateStart string
	// Weekly
	WeeklyDay []string
	// Cyclically
	Duration int
	Interval int
}
func (s *sjCrawler) Start(node int, followNew, followChange bool, dupType, changeType, timerType string,
		stimer StartTimer) (string, error) {

	return "", nil
}

func (s *sjCrawler) Stop() (string, error) {
	return "", nil
}

func (s *sjCrawler) Pause() (string, error) {
	return "", nil
}

func (s *sjCrawler) Resume(node int) (string, error) {
	return "", nil
}

func (s *sjCrawler) GetStatus() (string, error) {
	return "", nil
}

func (s *sjCrawler) GetSpeed() (float32, error) {
	return 0.0, nil
}

func (s *sjCrawler) AddNode(node int) (map[string]int, error) {
	return nil, nil
}

func (s *sjCrawler) ReduceNode(node int) (map[string]int, error) {
	return nil, nil
}

func (s *sjCrawler) GetSource() (map[string]int, error) {
	return nil, nil
}

func (s *sjCrawler) GetWebhook() {

}

func (s *sjCrawler) SetWebhook() {

}

func (s *sjCrawler) DeleteWebhook() {

}

func (s *sjCrawler) StartPublish() {

}

func (s *sjCrawler) StopPublish() {

}

func (s *sjCrawler) GetPublishStaus() {

}

func (s *sjCrawler) Delete() {

}

