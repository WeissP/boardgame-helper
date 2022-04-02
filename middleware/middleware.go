package middleware

import (
	"boardgame-helper/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetScore(w http.ResponseWriter, r *http.Request) {
}

func Constant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("point:%v\n", vars["point"])
	w.Write([]byte(`{"point123":"` + vars["point"] + `"}`))
}

func Save(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	t := time.Now()
	json.WriteNew(reqBody, "history", t.Format("Mon_Jan_2_2006"), t.Format("15:04:05_Jan_2_2006")+".json")
}
