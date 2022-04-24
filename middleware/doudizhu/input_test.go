package doudizhu

import (
	"reflect"
	"testing"
)

func Test_inputItem_position(t *testing.T) {
	type fields struct {
		Timestamp  string
		Stake      int
		BonusTiles int
		Players    [4]string
		RawPoints  int
		Winner     string
		Weight     map[string]int
		Lord       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantRes map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ii := inputItem{
				Timestamp:  tt.fields.Timestamp,
				Stake:      tt.fields.Stake,
				BonusTiles: tt.fields.BonusTiles,
				Players:    tt.fields.Players,
				RawPoints:  tt.fields.RawPoints,
				Winner:     tt.fields.Winner,
				Weight:     tt.fields.Weight,
				Lord:       tt.fields.Lord,
			}
			gotRes, err := ii.position()
			if (err != nil) != tt.wantErr {
				t.Errorf("inputItem.position() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("inputItem.position() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
