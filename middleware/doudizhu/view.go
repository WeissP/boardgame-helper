package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/timestamp"
	"net/http"
	"time"
)

type view struct{}

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
