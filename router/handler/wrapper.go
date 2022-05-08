package handler

import (
	"fmt"
	"net/http"
)

type Err struct {
	Error   error
	Message string
	Code    int
}

func (e Err) Empty() bool {
	return e.Message == ""
}

func (e Err) String() string {
	if e.Error == nil {
		return e.Message
	}
	return fmt.Errorf("%s\n%w", e.Message, e.Error).Error()
}

func NewError(error error, msg string, errCode int) Err {
	return Err{error, msg, errCode}
}

func CommonErr(error error, msg string) Err {
	return Err{error, msg, 500}
}

func NoImplementErr() Err {
	return Err{nil, "not yet implemented", 501}
}

func Wrap(fn func(http.ResponseWriter, *http.Request) Err) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if e := fn(w, r); !e.Empty() {
			fmt.Printf("%v\n", e)
			http.Error(w, e.String(), e.Code)
		}
	}
}
