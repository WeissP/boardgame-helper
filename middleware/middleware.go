package middleware

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func PointTest(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	vars := mux.Vars(r)
	fmt.Printf("point:%v\n", vars["point"])
	w.Write([]byte(`{"point123":"` + vars["point"] + `"}`))
	return
}

func SaveTest(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	fmt.Println("new")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.CommonErr(err, "can not parse request")
	}
	t := time.Now()
	json.WriteNew(reqBody, "history", timestamp.Date(t), timestamp.DateTime(t)+".json")
	return
}

func OnlyIntTest(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	fmt.Println("only int")
	r.ParseForm()
	num := r.Form.Get("number")
	if num == "" {
		return handler.CommonErr(nil, "number is empty")
	}
	i, err := strconv.Atoi(num)
	if err != nil {
		return handler.CommonErr(err, "can not parse number")
	}
	s := struct{ Num int }{i * i}
	res, err := json.From(s)
	if err != nil {
		return handler.CommonErr(err, fmt.Sprintf("can not generate JSON by struct %v", s))
	}
	fmt.Printf("res:%v\n", string(res))
	w.Write(res)
	return
}
