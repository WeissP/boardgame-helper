package router

import (
	"boardgame-helper/middleware/doudizhu"
	"boardgame-helper/router/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/doudizhu/view/now", handler.Wrap(doudizhu.GetViewNow))
	r.HandleFunc("/api/doudizhu/view/date", handler.Wrap(doudizhu.GetViewByDate))
	r.HandleFunc("/api/doudizhu/view/update", handler.Wrap(doudizhu.Update))
	r.HandleFunc("/api/doudizhu/disable", handler.Wrap(doudizhu.ToggleHistory(false)))
	r.HandleFunc("/api/doudizhu/enable", handler.Wrap(doudizhu.ToggleHistory(true)))
	r.HandleFunc("/api/doudizhu/new", handler.Wrap(doudizhu.AddInput))
	r.HandleFunc("/api/doudizhu/edit", handler.Wrap(doudizhu.EditInput))
	r.HandleFunc("/api/doudizhu/curPlayers", handler.Wrap(doudizhu.CurPlayers))
	return r
}
