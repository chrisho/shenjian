package shenjian

import "testing"

const (
	UserKey = "1dcf777f29-YjFlYzA4ZW"
	UserSecret = "IxZGNmNzc3ZjI5MW-1e0e924f7fb1ec0"
)

var (
	s *sjService
)

func init() {
	s = NewSjService(UserKey, UserSecret)
}

func TestService_GetMoneyInfo(t *testing.T) {
	if datas, err := s.GetMoneyInfo(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(datas)
	}
}

func TestService_GetNodeInfo(t *testing.T) {
	if datas, err := s.GetNodeInfo(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(datas)
	}
}

func TestService_GetAppList(t *testing.T) {
	if datas, page, total, err := s.GetAppList(1, 50); err != nil {
		t.Fatal(err)
	} else {
		t.Log(page, total, datas)
	}
}

func TestService_GetCrawlerList(t *testing.T) {
	if datas, page, total, err := s.GetCrawlerList(1, 50); err != nil {
		t.Fatal(err)
	} else {
		t.Log(page, total, datas)
	}
}

//func TestService_GetCreateCrawler(t *testing.T) {
//	if data, err := s.GetCreateCrawler("Test", "Info", "abcdefg"); err != nil {
//		t.Fatal(err)
//	} else {
//		t.Log(data)
//	}
//}