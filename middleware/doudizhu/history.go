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
	sort.Slice(his, func(i, j int) bool {
		res1, err := timestamp.Parse(his[i].InputItem.Timestamp)
		if err != nil {
			fmt.Println("cannot pasre string to time.Time!!!")
		}
		res2, err := timestamp.Parse(his[j].InputItem.Timestamp)
		if err != nil {
			fmt.Println("cannot pasre string to time.Time!!!")
		}
		return res1.Before(res2)
	})
	var currentView view
	if len(his) == 0 {
		print("No Game Today!!!")
		return currentView
	} else {
		//TODO players ID to Name
		currentView.PlayerNames = his[0].InputItem.Players // find function to change ID to Name
		deltaSlice := []DeltaPointsItem{}
		currentRound := 0
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
		var finalPoints [4]int
		for _, deltaPoints := range currentView.DeltaPoints {
			for i, finalPoint := range finalPoints {
				finalPoint += deltaPoints.Deltas[i]
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

func historyByDate(t time.Time) (historyItems, error) {
	date := timestamp.Date(t)
	// TODO: Implement
	currentHistoryItems, err := json.ReadDir[historyItem]("history", date)
	if err != nil {
		err = fmt.Errorf("cannot read from json file!!! %w", err)
	}
	return currentHistoryItems, err
	panic("not implemented") // TODO: Implement
}

func historyByDateTime(t time.Time) (historyItem, error) {
	date := timestamp.Date(t)
	dateTime := timestamp.DateTime(t)
	// TODO: Implement
	currentHistoryItem, err := json.ReadFile[historyItem]("history", date, dateTime+".json")
	if err != nil {
		err = fmt.Errorf("cannot read from json file!!! %w", err)
	}
	return currentHistoryItem, err
	panic("not implemented") // TODO: Implement
}
