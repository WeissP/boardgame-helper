package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"io/ioutil"
	"net/http"
)

type inputItem struct {
	Timestamp  string         `json:"timestamp"`
	Stake      int            `json:"stake"`
	BonusTiles int            `json:"bonusTiles"`
	Players    [4]string      `json:"players"`
	Points     int            `json:"points"`
	Winner     string         `json:"winner"`
	Weight     map[string]int `json:"weight"`
	Lord       string         `json:"lord"`
}

func TestInput() string {
	TestStruct := inputItem{"20220405...", 8, 3, [4]string{"bai", "xiao", "jintian", "yunfan"}, 88, "bai", map[string]int{"bai": 1}, "bai"}
	item, err := json.From(TestStruct)
	if err != nil {
		panic(err)
	} else {
		return string(item)
	}
}

func (ii inputItem) History() historyItem {
	panic("not implemented") // TODO: Implement
}

func AddInput(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.CommonErr(err, "can not parse request")
	}
	ii, err := json.Parse[inputItem](reqBody)
	if err != nil {
		return handler.CommonErr(err, "can not parse input: "+string(reqBody))
	}
	ii.History().write()
	updateCurView()
	return
}
