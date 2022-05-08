package main

import (
	"boardgame-helper/middleware/players"
	"boardgame-helper/router"
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"embed"
	"flag"
	"log"
	"net/http"
	"path"
	"time"

	_ "embed"
)

//go:embed config.json
var configJson []byte

//go:embed client/build
var client embed.FS

func main() {
	argPath := flag.String("path", "", "the path to the program")
	argWritable := flag.Bool("writable", true, "if config files can be written")
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
	json.SetWritable(*argWritable)

	err := players.Import()
	if err != nil {
		panic(err)
	}

	r := router.Router()

	spa := handler.New(client)
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
