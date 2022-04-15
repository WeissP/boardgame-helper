package doudizhu

import (
	"boardgame-helper/router/handler"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
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

// req := newGetReq("/api/doudizhu/view/date", map[string]string{
// 	"timestamp": "2020",
// })
// resJson, herr := response(req, AddInput)
