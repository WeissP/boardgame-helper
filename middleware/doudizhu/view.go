package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"fmt"
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
	view, err := his.View()
	if err != nil {
		return handler.CommonErr(err, "cannot get view by date")
	}
	w.Write(view.JSON())
	return
}

func Update(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	err := updateCurView()
	if err != nil {
		return handler.CommonErr(err, "can not update current view")
	}
	return
}

func updateCurView() error {
	his, err := relatedHistory(time.Now())
	if err != nil {
		return err
	}
	view, err := his.View()
	if err != nil {
		return fmt.Errorf("cannot update cuurentview %w", err)
	}
	currentView = view

	return nil
}
