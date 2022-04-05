package doudizhu

import (
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"fmt"
	"sort"
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
	// TODO: Implement
	hi.Enabled = status
	for _, item := range hi.PlayerDetails {
		item.Enabled = status
	}
	err := hi.write()
	if err != nil {
		fmt.Println("can not write to json file:")
	}
}

type historyItems []historyItem

func (his historyItems) View() view {
	// TODO: Implement
	sort.Slice(his, func(i, j int) bool { return his[i].InputItem.Timestamp < his[j].InputItem.Timestamp })
	if len(his) == 0 {
		panic("No game today!!!")
	} else {
		var currentView view
		currentView.PlayerNames = his[0].InputItem.Players // find function to change ID to Name
		deltaSlice := []DeltaPointsItem{}
		currentRound := 1
		for _, hisItem := range his {
			deltaPointsinstance := DeltaPointsItem{}
			deltaPointsinstance.Enabled = hisItem.Enabled
			deltaPointsinstance.Deltas = hisItem.Deltas
			deltaPointsinstance.Timestamp = hisItem.InputItem.Timestamp
			if hisItem.Enabled {
				currentRound++
				deltaPointsinstance.Round = currentRound
			} else {
				deltaPointsinstance.Round = 0
			}
			deltaSlice = append(deltaSlice, deltaPointsinstance)
		}
		currentView.DeltaPoints = deltaSlice
		finalPoints := [4]int{0, 0, 0, 0}
		for _, deltaPoints := range currentView.DeltaPoints {
			for i := 0; i < 4; i++ {
				finalPoints[i] += deltaPoints.Deltas[i]
			}
		}
		currentView.FinalPoints = finalPoints
		return currentView
	}

}

func (his historyItems) write() {
	for _, x := range his {
		x.write()
	}
}

func historyByDate(t time.Time) historyItems {
	date := timestamp.Date(t)
	// TODO: Implement
	currentHistoryItems, err := json.ReadDir[historyItem]("history", date)
	if err != nil {
		fmt.Println("cannot read from json file!!!")
	} else {
		return currentHistoryItems
	}
}

func historyByDateTime(t time.Time) historyItem {
	date := timestamp.Date(t)
	dateTime := timestamp.DateTime(t)
	// TODO: Implement
	currentHistoryItem, err := json.ReadFile[historyItem]("history", date, dateTime+".json")
	if err != nil {
		fmt.Println("cannot read from json file!!!")
	} else {
		return currentHistoryItem
	}
}
