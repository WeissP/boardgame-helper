package doudizhu

import (
	"boardgame-helper/middleware/players"
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"fmt"
	"net/http"
	"sort"
	"strconv"
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
	updateCurView()
	return err
}

func (hi *historyItem) toggle(status bool) error {
	hi.Enabled = status
	for _, item := range hi.PlayerDetails {
		item.Enabled = status
	}
	err := hi.write()
	if err != nil {
		fmt.Println("can not write to json file:")
	}
	return err
}

type historyItems []historyItem

func IDsToNames(ids [4]string) (names [4]string, err error) {
	for i, x := range ids {
		name, e := players.IDToName(x)
		if e != nil {
			err = fmt.Errorf("error IDsToNames: %w ", e)
			return
		}
		names[i] = name
	}
	return
}

func (his historyItems) View() (res view, err error) {
	sort.Slice(his, func(i, j int) bool {
		resI, err := timestamp.Parse(his[i].InputItem.Timestamp)
		if err != nil {
			panic(err)
		}
		resJ, err := timestamp.Parse(his[j].InputItem.Timestamp)
		if err != nil {
			panic(err)
		}
		return resI.Before(resJ)
	})
	if len(his) == 0 {
		return res, nil
	} else {
		names, e := IDsToNames(his[0].InputItem.Players) // find function to change ID to Name
		if e != nil {
			err = fmt.Errorf("error IDsToNames: %w ", e)
		}
		res.PlayerNames = names
		currentRound := 0
		for _, item := range his {
			dpi := DeltaPointsItem{}
			dpi.Enabled = item.Enabled
			dpi.Deltas = item.Deltas
			dpi.Timestamp = item.InputItem.Timestamp
			if item.Enabled {
				currentRound++
				dpi.Round = currentRound
			}
			res.DeltaPoints = append(res.DeltaPoints, dpi)
		}
		var finalPoints [4]int
		for _, dp := range res.DeltaPoints {
			if !dp.Enabled {
				continue
			}
			for i := range finalPoints {
				finalPoints[i] += dp.Deltas[i]
			}
		}
		res.FinalPoints = finalPoints
		return res, nil
	}
}

func (his historyItems) write() {
	for _, x := range his {
		err := x.write()
		if err != nil {
			panic(err)
		}
	}
}

func historyByDate(t time.Time) (hi historyItems, err error) {
	date := timestamp.Date(t)
	hi, err = json.ReadDir[historyItem]("history", date)
	if err != nil {
		err = fmt.Errorf("cannot read from json file!!! %w", err)
	}
	return
}

func historyByDateTime(t time.Time) (hi historyItem, err error) {
	date := timestamp.Date(t)
	dateTime := timestamp.DateTime(t)
	hi, err = json.ReadFile[historyItem]("history", date, dateTime+".json")
	if err != nil {
		err = fmt.Errorf("cannot read from json file!!! %w", err)
	}
	return
}

func ToggleHistory(status bool) func(w http.ResponseWriter, r *http.Request) handler.Err {
	return func(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
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
		err = hi.toggle(status)
		if err != nil {
			return handler.CommonErr(err, "can not toggle history")
		}
		return
	}
}

func CurPlayers(w http.ResponseWriter, r *http.Request) (herr handler.Err) {
	res := struct {
		Players [4]string `json:"players"`
	}{}
	now := time.Now()
	hiss, err := historyByDate(now)
	if err != nil {
		handler.CommonErr(nil, "cannot get history by date!!!")
	}
	var lastHis historyItem
	if len(hiss) == 0 {
		return handler.CommonErr(nil, "no game today")
	}

	for _, his := range hiss {
		if lastHis.InputItem.Timestamp == "" {
			lastHis = his
		} else {
			lastHisTs, err := timestamp.Parse(lastHis.InputItem.Timestamp)
			if err != nil {
				panic(err)
			}

			hisTs, err := timestamp.Parse(his.InputItem.Timestamp)
			if err != nil {
				return handler.CommonErr(nil, "can not parse timestamp in history")
			}

			if lastHisTs.Before(hisTs) {
				lastHis = his
			}
		}
	}
	res.Players = lastHis.InputItem.Players

	for i, id := range res.Players {
		if id == "" {
			return handler.CommonErr(nil, "missing player - player "+strconv.Itoa(i+1)+" is missing")
		}
	}
	j, err := json.From(res)
	if err != nil {
		return handler.CommonErr(err, "can not convert players to JSON")
	}
	w.Write(j)
	return
}
