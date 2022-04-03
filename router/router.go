package router

import (
	"boardgame-helper/middleware"
	"boardgame-helper/router/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/test/point/{point}", handler.Wrap(middleware.Constant))
	r.HandleFunc("/api/test/int", handler.Wrap(middleware.OnlyInt))
	r.HandleFunc("/api/save", handler.Wrap(middleware.Save))
	return r
}
