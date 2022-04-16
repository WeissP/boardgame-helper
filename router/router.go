package router

import (
	"boardgame-helper/middleware"
	"boardgame-helper/middleware/doudizhu"
	"boardgame-helper/router/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/test/point/{point}", handler.Wrap(middleware.PointTest))
	r.HandleFunc("/api/test/int", handler.Wrap(middleware.OnlyIntTest))
	r.HandleFunc("/api/save", handler.Wrap(middleware.SaveTest))
	r.HandleFunc("/api/doudizhu/view/now", handler.Wrap(doudizhu.GetViewNow))
	r.HandleFunc("/api/doudizhu/view/date", handler.Wrap(doudizhu.GetViewByDate))
	r.HandleFunc("/api/doudizhu/view/update", handler.Wrap(doudizhu.Update))
	r.HandleFunc("/api/doudizhu/disable", handler.Wrap(doudizhu.DisableHistory))
	r.HandleFunc("/api/doudizhu/enable", handler.Wrap(doudizhu.EnableHistory))
	r.HandleFunc("/api/doudizhu/new", handler.Wrap(doudizhu.AddInput))
	r.HandleFunc("/api/doudizhu/edit", handler.Wrap(doudizhu.EditInput))
	r.HandleFunc("/api/doudizhu/curPlayers", handler.Wrap(doudizhu.CurPlayers))
	return r
}
