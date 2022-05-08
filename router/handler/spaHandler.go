package handler

import (
	"io/fs"
	"net/http"
)

type spaHandler struct {
	fs fs.FS
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fsys, err := fs.Sub(h.fs, "client/build")
	if err != nil {
		panic(err)
	}
	http.FileServer(http.FS(fsys)).ServeHTTP(w, r)
}

func New(fs fs.FS) spaHandler {
	return spaHandler{fs}
}
