package main

import (
	"boardgame-helper/middleware/players"
	"boardgame-helper/router"
	"boardgame-helper/utils/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	_ "embed"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		fmt.Printf("can not find file:%v\n", err)
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

//go:embed config.json
var configJson []byte

func main() {
	argPath := flag.String("path", "", "the path to the program")
	flag.Parse()
	dataPath := *argPath
	if dataPath == "" {
		m, err := json.Parse[map[string]string](configJson)
		if err != nil {
			panic("the data path is not set by argument, and config.json can not be correctly parsed")
		}
		if dir, ok := m["Dir"]; ok {
			dataPath = dir
		}
	}
	if dataPath == "" {
		panic("the data path is neither set by argument nor by config.json")
	}
	log.Printf("The program path is: %s", dataPath)

	json.SetRootPath(path.Join(dataPath, "data"))

	err := players.Import()
	if err != nil {
		panic(err)
	}

	r := router.Router()

	spa := spaHandler{staticPath: path.Join(dataPath, "client", "build"), indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8888",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
