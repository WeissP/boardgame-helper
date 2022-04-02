package router

import (
	"boardgame-helper/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/test/{point}", middleware.Constant)
	r.HandleFunc("/api/save", middleware.Save)
	return r
}
