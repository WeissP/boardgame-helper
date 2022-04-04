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
	PlayerNames [4]string          `json:"playerNames"`
	DeltaPoints [4]DeltaPointsItem `json:"deltaPoints"`
	FinalPoints [4]int             `json:"finalPoints"`
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
	panic("not implemented") // TODO: Implement
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
	view := historyByDate(t).View()
	w.Write(view.JSON())
	return
}

func updateCurView() {
	currentView = historyByDate(time.Now()).View()
}
