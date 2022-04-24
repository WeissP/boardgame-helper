package doudizhu

import (
	"boardgame-helper/router/handler"
	"net/http"
	"reflect"
	"testing"
)

func TestCurPlayers(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantHerr handler.Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHerr := CurPlayers(tt.args.w, tt.args.r); !reflect.DeepEqual(gotHerr, tt.wantHerr) {
				t.Errorf("CurPlayers() = %v, want %v", gotHerr, tt.wantHerr)
			}
		})
	}
}
