package doudizhu

import (
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"fmt"
	"time"
)

type historyItem struct {
	InputItem     inputItem       `json:"inputItem"`
	Deltas        [4]int          `json:"deltas"`
	Enabled       bool            `json:"enabled"`
	PlayerDetails [4]playerDetail `json:"playerDetails"`
}
type playerDetail struct {
	Player      string `json:"player"`
	Timestamp   string `json:"timestamp"`
	Stake       int    `json:"stake"`
	BonusTiles  int    `json:"bonusTiles"`
	Lord        bool   `json:"lord"`
	Weight      int    `json:"weight"`
	Winner      string `json:"winner"`
	Rawpoints   int    `json:"rawpoints"`
	Deltapoints int    `json:"deltapoints"`
	Position    string `json:"position"`
	Enabled     bool   `json:"enabled"`
}

func (hi historyItem) write() error {
	jsonFile, err := json.From(hi)
	if err != nil {
		return err
	}
	t, err := timestamp.Parse(hi.InputItem.Timestamp)
	if err != nil {
		return fmt.Errorf("can not parse timestamp [%s]:%w", hi.InputItem.Timestamp, err)
	}
	date := timestamp.Date(t)
	dateTime := timestamp.DateTime(t)
	err = json.WriteNew(jsonFile, "history", date, dateTime+".json")
	return err
}

func (hi *historyItem) toggle(status bool) {
	panic("not implemented") // TODO: Implement
}

type historyItems []historyItem

func (his historyItems) View() view {
	panic("not implemented") // TODO: Implement
}

func (his historyItems) write() {
	for _, x := range his {
		x.write()
	}
}

func historyByDate(t time.Time) historyItems {
	date := timestamp.Date(t)
	_ = date
	panic("not implemented") // TODO: Implement
}

func historyByDateTime(t time.Time) historyItem {
	dateTime := timestamp.DateTime(t)
	_ = dateTime
	panic("not implemented") // TODO: Implement
}
