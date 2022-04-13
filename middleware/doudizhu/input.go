package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

var posArray = [4]string{"lord", "support", "carry", "defense"}

type inputItem struct {
	Timestamp  string         `json:"timestamp"`
	Stake      int            `json:"stake"`
	BonusTiles int            `json:"bonusTiles"`
	Players    [4]string      `json:"players"`
	RawPoints  int            `json:"points"`
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

func (ii inputItem) History() (hi historyItem, err error) {
	hi.InputItem = ii
	// deltas
	hi.Deltas, err = ii.deltas()
	if err != nil {
		return
	}

	hi.Enabled = true

	positionMap, err := ii.position()
	if err != nil {
		return
	}

	for i, id := range hi.InputItem.Players {
		hi.PlayerDetails[i] = playerDetails(ii, id, positionMap[id], hi.Deltas[i])
	}

	return hi, err
}

func (ii inputItem) position() (res map[string]string, err error) {
	lordIdx := -1
	for i, id := range ii.Players {
		if ii.Lord == id {
			lordIdx = i
		}
	}
	if lordIdx == -1 {
		err = errors.New("can not find lord")
		return
	}

	dupPlayers := append(ii.Players[:], ii.Players[:]...)
	orderedPlayers := dupPlayers[lordIdx : lordIdx+4]
	res = make(map[string]string)
	for i, id := range orderedPlayers {
		res[id] = posArray[i]
	}
	return
}

func (ii inputItem) deltas() (res [4]int, err error) {
	lordWeight := ii.Weight[ii.Lord]
	weightSum := 0
	for _, weight := range ii.Weight {
		weightSum += weight
	}
	if weightSum != 0 {
		err = fmt.Errorf("you are large sha bi, sum of weight is not 0!!!")
		return
	}

	ratio := float64(ii.RawPoints*ii.Stake) / float64(lordWeight)
	pointPerWeight := int(math.Ceil(ratio))
	for key, weight := range ii.Weight {
		for i, id := range ii.Players {
			if key == id {
				res[i] = weight * pointPerWeight
			}
		}
	}
	// basic point
	for i, id := range ii.Players {
		// lord win
		if ii.Winner == ii.Lord {
			if id == ii.Lord {
				res[i] += 24
			} else {
				res[i] -= 8
			}
		} else { // lord lose
			if id == ii.Lord {
				res[i] -= 24
			} else {
				res[i] += 8
			}
		}
	}
	return
}

func playerDetails(ii inputItem, id, pos string, delta int) (res playerDetail) {
	res.Player = id
	res.Timestamp = ii.Timestamp
	res.Stake = ii.Stake
	res.BonusTiles = ii.BonusTiles

	res.Lord = pos == posArray[0]
	res.Position = pos

	res.Weight = ii.Weight[id]
	res.Winner = ii.Winner
	res.Rawpoints = ii.RawPoints
	res.Deltapoints = delta
	res.Enabled = true

	return
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
