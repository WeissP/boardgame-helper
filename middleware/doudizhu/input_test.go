package doudizhu

import (
	"reflect"
	"testing"
)

func Test_inputItem_position(t *testing.T) {
	genInputItem := func(players [4]string, lord string) (ii inputItem) {
		ii.Players = players
		ii.Lord = lord
		return
	}
	tests := []struct {
		name    string
		players [4]string
		lord    string
		wantRes [4]string // the same order as posArray in input.go
	}{{
		"test func postion in inputItem",
		[4]string{"a", "b", "c", "d"},
		"b",
		[4]string{"b", "c", "d", "a"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ii := genInputItem(tt.players, tt.lord)
			wantResMap := map[string]string{}
			for i, id := range tt.wantRes {
				wantResMap[id] = posArray[i]
			}
			gotRes, _ := ii.position()
			if !reflect.DeepEqual(gotRes, wantResMap) {
				t.Errorf("inputItem.position() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
