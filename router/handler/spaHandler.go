package handler

import (
	"errors"
	"io/fs"
	"net/http"
	"strings"
)

type spaHandler struct {
	fs fs.FS
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fsys, err := fs.Sub(h.fs, "client/build")
	if err != nil {
		panic(err)
	}

	path := r.URL.Path
	_, err = fs.Stat(fsys, strings.Trim(path, "/"))
	switch {
	case err == nil, path == "/", strings.HasPrefix(path, "/api/"):
		http.FileServer(http.FS(fsys)).ServeHTTP(w, r)
	case errors.Is(err, fs.ErrNotExist):
		// use react router, redirect to index 
		content, err := fs.ReadFile(fsys, "index.html")
		if err != nil {
			panic(err)
		}
		w.Write(content)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func New(fs fs.FS) spaHandler {
	return spaHandler{fs}
}
