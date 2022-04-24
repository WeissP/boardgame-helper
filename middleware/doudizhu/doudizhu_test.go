package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func newPostReq(urlStr string, jsonFile []byte) *http.Request {
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonFile))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	return req
}

func newGetReq(urlStr string, kv map[string]string) *http.Request {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	for k, v := range kv {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return req
}

func response(r *http.Request, handler func(http.ResponseWriter, *http.Request) handler.Err) (jsonFile []byte, herr handler.Err) {
	w := httptest.NewRecorder()
	herr = handler(w, r)

	resp := w.Result()
	jsonFile, _ = io.ReadAll(resp.Body)
	return
}

func initTest() {
	// init part, we need it because in it will read and set the file path,
	// that will be used in the function we call
	jsonFile, err := os.Open("../../config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println(string(byteValue))

	json.InitConfig(string(byteValue), "", "")
}

type deltaPoints struct {
	Ts  string         `json:"timestamp"`
	Pts map[string]int `json:"final_points"`
}

// testfiles to json
func testToHistory() {
	inputs, err := json.ReadDir[inputItem]("test", "Thu_Mar31_2022")
	if err != nil {
		panic(err)
	}
	for _, val := range inputs {
		jsf, err := json.From(val)
		if err != nil {
			panic(err)
		}
		req := newPostReq("/api/doudizhu/new", jsf)
		_, herr := response(req, AddInput)
		if !herr.Empty() {
			panic(herr.Message)
		}

	}
}

func TestDeltaPoints(t *testing.T) {
	initTest()

	testToHistory()

	dpts, err := json.ReadDir[deltaPoints]("test", "Thu_Mar31_2022")
	if err != nil {
		t.Errorf("error readDir: %v", err)
	}

	for _, x := range dpts {
		t.Run(x.Ts, func(t *testing.T) {
			gotTime, err := timestamp.Parse(x.Ts)
			if err != nil {
				panic(err)
			}

			his, err := historyByDateTime(gotTime)
			if err != nil {
				t.Errorf("Error historyByDate : %v", err)
			}

			for _, playerDetail := range his.PlayerDetails {
				if playerDetail.Deltapoints != x.Pts[playerDetail.Player] {
					t.Errorf("Error deltapoints, %v: want %v, got %v", playerDetail.Player, x.Pts[playerDetail.Player], playerDetail.Deltapoints)
				}
			}
		})
	}
}
