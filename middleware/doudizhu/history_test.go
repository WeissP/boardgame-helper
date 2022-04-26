package doudizhu

import (
	"reflect"
	"testing"
	"time"
)

func Test_historyItems_filterByDateRange(t *testing.T) {
	type args struct {
		oldest time.Time
		newest time.Time
	}
	tests := []struct {
		name    string
		hiss    historyItems
		args    args
		wantRes historyItems
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := tt.hiss.filterByDateRange(tt.args.oldest, tt.args.newest)
			if (err != nil) != tt.wantErr {
				t.Errorf("historyItems.filterByDateRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("historyItems.filterByDateRange() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_curHistory(t *testing.T) {
	tests := []struct {
		name    string
		wantHis historyItems
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHis, err := curHistory()
			if (err != nil) != tt.wantErr {
				t.Errorf("curHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHis, tt.wantHis) {
				t.Errorf("curHistory() = %v, want %v", gotHis, tt.wantHis)
			}
		})
	}
}

func Test_historyItems_lastPlayers(t *testing.T) {
	tests := []struct {
		name    string
		his     historyItems
		wantRes [4]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := tt.his.lastPlayers()
			if (err != nil) != tt.wantErr {
				t.Errorf("historyItems.lastPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("historyItems.lastPlayers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
