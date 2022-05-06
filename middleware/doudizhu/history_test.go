package doudizhu

import (
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"reflect"
	"testing"
	"time"
)

func Test_historyItems_filterByDateRange(t *testing.T) {
	initTest()
	type args struct {
		oldest time.Time
		newest time.Time
	}
	oldT, err1 := timestamp.Parse("2022-04-20T17:31:12.768Z")
	if err1 != nil {
		t.Errorf("error in case generation: cannot parse tring to ts: oldest")
	}
	newT, err2 := timestamp.Parse("2022-04-20T17:31:32.768Z")
	if err2 != nil {
		t.Errorf("error in case generation: cannot parse tring to ts: newest")
	}
	hiss, err := json.ReadDir[historyItem]("history", "Wed_Apr_20_2022")
	if err != nil {
		t.Errorf("error in case generation: cannot read data dir")
	}
	tests := []struct {
		name    string
		hiss    historyItems
		args    args
		wantRes historyItems
	}{{
		name: "Test function filterByDateRange in historyItems",
		hiss: hiss,
		args: args{oldest: oldT, newest: newT},
		wantRes: []historyItem{
			{
				InputItem: inputItem{
					Timestamp:  "2022-04-20T17:31:22.768Z",
					Stake:      3,
					BonusTiles: 8,
					Players:    [4]string{"yunfan", "xiao", "bai", "jintian"},
					RawPoints:  25,
					Winner:     "xiao",
					Weight:     map[string]int{"bai": 0, "jintian": 0, "xiao": 3, "yunfan": -3},
					Lord:       "xiao",
				},
				Deltas:  [4]int{-83, 99, -8, -8},
				Enabled: true,
				PlayerDetails: [4]playerDetail{
					{
						Player:      "yunfan",
						Timestamp:   "2022-04-20T17:31:22.768Z",
						Stake:       3,
						BonusTiles:  8,
						Lord:        false,
						Weight:      -3,
						Winner:      "xiao",
						Rawpoints:   25,
						Deltapoints: -83,
						Position:    "defense",
						Enabled:     true,
					},
					{
						Player:      "xiao",
						Timestamp:   "2022-04-20T17:31:22.768Z",
						Stake:       3,
						BonusTiles:  8,
						Lord:        true,
						Weight:      3,
						Winner:      "xiao",
						Rawpoints:   25,
						Deltapoints: 99,
						Position:    "lord",
						Enabled:     true,
					},
					{
						Player:      "bai",
						Timestamp:   "2022-04-20T17:31:22.768Z",
						Stake:       3,
						BonusTiles:  8,
						Lord:        false,
						Weight:      0,
						Winner:      "xiao",
						Rawpoints:   25,
						Deltapoints: -8,
						Position:    "support",
						Enabled:     true,
					},
					{
						Player:      "jintian",
						Timestamp:   "2022-04-20T17:31:22.768Z",
						Stake:       3,
						BonusTiles:  8,
						Lord:        false,
						Weight:      0,
						Winner:      "xiao",
						Rawpoints:   25,
						Deltapoints: -8,
						Position:    "carry",
						Enabled:     true,
					},
				},
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, _ := tt.hiss.filterByDateRange(tt.args.oldest, tt.args.newest)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("\ngot: %v \nwant: %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_relatedHistory(t *testing.T) {
	initTest()
	time1, err1 := timestamp.Parse("2022-04-20T17:31:12.768Z")
	if err1 != nil {
		t.Errorf("error in case generation: cannot parse tring to ts: oldest")
	}
	time2, err2 := timestamp.Parse("2022-04-21T02:31:12.768Z")
	if err2 != nil {
		t.Errorf("error in case generation: cannot parse tring to ts: oldest")
	}
	hiss, err := json.ReadDir[historyItem]("history", "Wed_Apr_20_2022")
	if err != nil {
		t.Errorf("error in case generation: cannot read data dir")
	}
	tests := []struct {
		name     string
		relatedT time.Time
		wantHis  historyItems
	}{{
		"ts:2022-04-20T17:31:12.768Z",
		time1,
		hiss,
	}, {
		"ts:2022-04-20T02:31:12.768Z",
		time2,
		hiss,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHis, err := relatedHistory(tt.relatedT)
			if err != nil {
				t.Errorf("error happens: %v", err)
			}
			if !reflect.DeepEqual(gotHis, tt.wantHis) {
				t.Errorf("got: %v\nwant: %v", gotHis.timestamps(), tt.wantHis.timestamps())
			}
		})
	}
}

func Test_historyItems_lastPlayers(t *testing.T) {
	initTest()
	hiss, err := json.ReadDir[historyItem]("history", "Wed_Apr_20_2022")
	if err != nil {
		t.Errorf("error in case generation: cannot read data dir")
	}
	tests := []struct {
		name    string
		his     historyItems
		wantRes [4]string
	}{{
		"Wed_Apr_20_2022",
		hiss,
		[4]string{"yunfan", "xiao", "bai", "jintian"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, _ := tt.his.lastPlayers()
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("historyItems.lastPlayers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
