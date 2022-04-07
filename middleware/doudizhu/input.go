package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type inputItem struct {
	Timestamp  string         `json:"timestamp"`
	Stake      int            `json:"stake"`
	BonusTiles int            `json:"bonusTiles"`
	Players    [4]string      `json:"players"`
	Points     int            `json:"points"`
	Winner     string         `json:"winner"`
	Weight     map[string]int `json:"weight"`
	Lord       string         `json:"lord"`
}

func TestInput() string {
	TestStruct := inputItem{"20220405...", 8, 3, [4]string{"bai", "xiao", "jintian", "yunfan"}, 88, "bai", map[string]int{"bai": 1}, "bai"}
	item, err := json.From(TestStruct)
	if err != nil {
		panic(err)
	} else {
		return string(item)
	}
}

func (ii inputItem) History() (historyItem, error) {
	// TODO: Implement
	var Item historyItem
	Item.InputItem = ii
	// deltas
	lordWeight := Item.InputItem.Weight[Item.InputItem.Lord]
	weightSum := 0
	for _, weight := range Item.InputItem.Weight {
		weightSum += weight
	}
	ratio := Item.InputItem.Points * Item.InputItem.Stake / lordWeight
	pointPerWeight := int(math.Ceil(float64(ratio)))
	for key, weight := range Item.InputItem.Weight {
		for i, id := range Item.InputItem.Players {
			if key == id {
				Item.Deltas[i] = weight * pointPerWeight
			}
		}
	}
	//basic delta
	for i, id := range Item.InputItem.Players {
		// lord win
		if Item.InputItem.Winner == Item.InputItem.Lord {
			if id == Item.InputItem.Lord {
				Item.Deltas[i] += 24
			} else {
				Item.Deltas[i] -= 8
			}
		} else { //lord lose
			if id == Item.InputItem.Lord {
				Item.Deltas[i] -= 24
			} else {
				Item.Deltas[i] += 8
			}
		}
	}
	Item.Enabled = true
	// players detail
	for i, id := range Item.InputItem.Players {
		Item.PlayerDetails[i].Player = id
		Item.PlayerDetails[i].Timestamp = Item.InputItem.Timestamp
		Item.PlayerDetails[i].Stake = Item.InputItem.Stake
		Item.PlayerDetails[i].BonusTiles = Item.InputItem.BonusTiles
		// bool lord
		if id == Item.InputItem.Lord {
			Item.PlayerDetails[i].Lord = true
		} else {
			Item.PlayerDetails[i].Lord = false
		}
		Item.PlayerDetails[i].Weight = Item.InputItem.Weight[id]
		Item.PlayerDetails[i].Winner = Item.InputItem.Winner
		Item.PlayerDetails[i].Rawpoints = Item.InputItem.Points
		Item.PlayerDetails[i].Deltapoints = Item.Deltas[i]
		Item.PlayerDetails[i].Enabled = Item.Enabled
		// calculate position
		var positionMap map[string]string
		lordId := -1
		for i, id := range Item.InputItem.Players {
			if Item.InputItem.Lord == id {
				positionMap[id] = "Lord"
				lordId = i
			}
		}
		currentPosition := -1
		for i, _ := range Item.InputItem.Players {
			if i <= lordId {
				currentPosition = i + 4 - lordId
			} else {
				currentPosition = i - lordId
			}
			switch currentPosition {
			case 0:
				Item.PlayerDetails[i].Position = "lord"
			case 1:
				Item.PlayerDetails[i].Position = "support"
			case 2:
				Item.PlayerDetails[i].Position = "carry"
			case 3:
				Item.PlayerDetails[i].Position = "defense"
			}
		}
	}
	// check error
	var err error
	if weightSum != 0 {
		err = fmt.Errorf("sum of weight is not 0")
		Item.Enabled = false
		for i, _ := range Item.InputItem.Players {
			Item.PlayerDetails[i].Enabled = Item.Enabled
		}
	}
	return Item, err
}

func AddInput(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.CommonErr(err, "can not parse request")
	}
	ii, err := json.Parse[inputItem](reqBody)
	if err != nil {
		return handler.CommonErr(err, "can not parse input: "+string(reqBody))
	}
	res, err := ii.History()
	if err != nil {
		return handler.CommonErr(err, "cannot generate the correct history!!!")
	}
	res.write()
	updateCurView()
	return
}
