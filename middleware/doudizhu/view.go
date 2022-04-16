package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"net/http"
	"time"
)

type DeltaPointsItem struct {
	Round     int    `json:"round"`
	Timestamp string `json:"timestamp"`
	Enabled   bool   `json:"enabled"`
	Deltas    [4]int `json:"deltas"`
}

type view struct {
	PlayerNames [4]string         `json:"playerNames"`
	DeltaPoints []DeltaPointsItem `json:"deltaPoints"`
	FinalPoints [4]int            `json:"finalPoints"`
}

func JsonToStruct() view {
	Testview, err := json.ReadFile[view]("test", "Testview.json")
	if err != nil {
		panic(err)
	} else {
		return Testview
	}
}

var currentView view

func (v view) JSON() []byte {
	res, err := json.From(v)
	if err != nil {
		panic(err)
	}
	return res
}

func GetViewNow(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	w.Write(currentView.JSON())
	return
}

func GetViewByDate(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	r.ParseForm()
	ts := r.Form.Get("timestamp")
	if ts == "" {
		return handler.CommonErr(nil, "timestamp is empty")
	}
	t, err := timestamp.Parse(ts)
	if err != nil {
		return handler.CommonErr(err, "can not parse timestamp")
	}
	his, err := historyByDate(t)
	if err != nil {
		panic(err)
	}
	w.Write(his.View().JSON())
	return
}

func Update(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	updateCurView()
	return
}

func updateCurView() {
	his, err := historyByDate(time.Now())
	if err != nil {
		panic(err)
	}
	currentView = his.View()
}
