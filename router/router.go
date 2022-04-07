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
	r.HandleFunc("/api/doudizhu/view/date", handler.Wrap(doudizhu.GetViewByDate))
	return r
}
