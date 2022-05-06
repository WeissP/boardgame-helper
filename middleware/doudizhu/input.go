package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
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
	RawPoints  int            `json:"rawPoints"`
	Winner     string         `json:"winner"`
	Weight     map[string]int `json:"weight"`
	Lord       string         `json:"lord"`
}

func (ii inputItem) isTie() bool {
	for _, x := range ii.Weight {
		if x != 0 {
			return false
		}
	}
	return true
}

func (ii inputItem) valid() bool {
	for _, x := range ii.Players {
		if x == "" {
			return false
		}
	}
	return true
}

func constrain(ii inputItem) (err error) {
	//sum may have some parameter safety problem.
	sum := 0
	for _, x := range ii.Weight {
		sum += x
	}
	if sum != 0 {
		return fmt.Errorf("the sum of weight must be zero")
	}

	if !ii.valid() {
		return fmt.Errorf("Players must not be empty")
	}

	if ii.isTie() {
		return nil
	}

	// winner is lord
	// lord>=3,NonL<0
	if ii.Lord == ii.Winner {
		for id, w := range ii.Weight {
			switch {
			case id != ii.Lord && w >= 0:
				return fmt.Errorf("the peasant %s loss, his weight should under 0", id)
			case id == ii.Lord && w < 3:
				return fmt.Errorf("the lord %s win, his weight should be over/equal 3", id)
			}
		}
		// winner is not lord
		// lord<0, defense>=1, winner>=2,NonL>0
	} else {
		p, err := ii.position()
		if err != nil {
			return fmt.Errorf("can not generate position: %w", err)
		}

		for id, w := range ii.Weight {
			switch {
			case id == ii.Lord && w >= 0:
				return fmt.Errorf("the lord %s loss, his weight should be under 0", id)
			case p[id] == "defense" && w < 1:
				return fmt.Errorf("the defender %s win, his weight should be over/equal 1", id)
			case id == ii.Winner && w < 2:
				return fmt.Errorf("the peasant %s win, his weight should be over/equal 2", id)
			case id != ii.Lord && w <= 0:
				return fmt.Errorf("the peasant %s win, his weight should be over/equal 0", id)

			}

		}

	}

	return nil
}
func (ii inputItem) History() (hi historyItem, err error) {
	hi.InputItem = ii

	err = constrain(ii)
	if err != nil {
		return
	}

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
		hi.PlayerDetails[i] = ii.playerDetails(id, positionMap[id], hi.Deltas[i])
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
	if ii.isTie() {
		return
	}

	lordWeight := ii.Weight[ii.Lord]
	weightSum := 0
	for _, weight := range ii.Weight {
		weightSum += weight
	}
	if weightSum != 0 {
		err = fmt.Errorf("you are large sha bi, sum of weight is not 0!!!")
		return
	}

	ratio := float64(ii.RawPoints*ii.Stake) / math.Abs(float64(lordWeight))
	pointPerWeight := int(math.Ceil(ratio))
	for key, weight := range ii.Weight {
		for i, id := range ii.Players {
			if key == id {
				res[i] = weight * pointPerWeight
			}
		}
	}
	// basic point
	if !exists(ii.Winner, ii.Players) {
		err = fmt.Errorf("no winner <%s> found in %v", ii.Winner, ii.Players)
		return
	}

	if !exists(ii.Lord, ii.Players) {
		err = fmt.Errorf("no lord <%s> found in %v", ii.Lord, ii.Players)
		return
	}

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

func exists(name string, names [4]string) bool {
	for _, x := range names {
		if name == x {
			return true
		}
	}
	return false
}

func (ii inputItem) playerDetails(id, pos string, delta int) playerDetail {
	return playerDetail{
		Player:      id,
		Timestamp:   ii.Timestamp,
		Stake:       ii.Stake,
		BonusTiles:  ii.BonusTiles,
		Lord:        pos == posArray[0],
		Position:    pos,
		Weight:      ii.Weight[id],
		Winner:      ii.Winner,
		Rawpoints:   ii.RawPoints,
		Deltapoints: delta,
		Enabled:     true,
	}
}

func (ii inputItem) JSON() []byte {
	res, err := json.From(ii)
	if err != nil {
		panic(err)
	}
	return res
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
	return
}

func EditInput(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	r.ParseForm()
	tsStr := r.Form.Get("timestamp")
	if tsStr == "" {
		return handler.CommonErr(nil, "timestamp is empty")
	}
	ts, err := timestamp.Parse(tsStr)
	if err != nil {
		return handler.CommonErr(err, "can not parse timestamp")
	}
	hi, err := historyByDateTime(ts)
	if err != nil {
		return handler.CommonErr(err, "can not get history by date time")
	}
	w.Write(hi.InputItem.JSON())
	return
}
