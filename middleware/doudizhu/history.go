package doudizhu

import (
	"boardgame-helper/middleware/players"
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"errors"
	"fmt"
	"net/http"
	"os"
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

func (his historyItems) timestamps() (res []string) {
	for _, x := range his {
		res = append(res, x.InputItem.Timestamp)
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

func (his historyItems) lastPlayers() (res [4]string, err error) {
	if len(his) == 0 {
		err = errors.New("no history item can be found")
		return
	}
	var lastHi historyItem
	for _, hi := range his {
		if lastHi.InputItem.Timestamp == "" {
			lastHi = hi
		} else {
			lastHiTs, err := timestamp.Parse(lastHi.InputItem.Timestamp)
			if err != nil {
				panic(err)
			}

			hiTs, err := timestamp.Parse(hi.InputItem.Timestamp)
			if err != nil {
				err = fmt.Errorf("can not parse timestamp in history:%v", hi)
			}

			if lastHiTs.Before(hiTs) {
				lastHi = hi
			}
		}
	}
	res = lastHi.InputItem.Players
	return
}

// filter applies pred to each element in his, if the returned value is true, then add it to res.
// if error is not nil, the loop will be immediately terminated, and this error will be returned.
func (his historyItems) filter(pred func(historyItem) (bool, error)) (res historyItems, err error) {
	for _, hi := range his {
		ok, e := pred(hi)
		if e != nil {
			err = fmt.Errorf("error in filter:%w", err)
			return
		}
		if ok {
			res = append(res, hi)
		}
	}
	return
}

// filterByDateRange only returns the historyItems in his that their timestamp are between oldest and newest.
func (his historyItems) filterByDateRange(from, to time.Time) (res historyItems, err error) {
	if from.After(to) {
		err = fmt.Errorf("the oldest timestamp [%v] is newer than the newest [%v]!", from.String(), to.String())
		return
	}
	pred := func(hi historyItem) (ok bool, err error) {
		t, err := timestamp.Parse(hi.InputItem.Timestamp)
		if err != nil {
			return false, fmt.Errorf("the timestamp in histroy <%v> can not be parsed: %w", hi, err)
		}
		ok = !t.Before(from) && !t.After(to)
		return
	}
	return his.filter(pred)
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

// historyByDateRange returns all historyItems between the given range.
func historyByDateRange(from, to time.Time) (historyItems, error) {
	if from.After(to) {
		return nil, fmt.Errorf("from [%v] is older than to [%v]", from, to)
	}

	var res []historyItem
	for cur := from; !cur.After(to); cur = cur.AddDate(0, 0, 1) {
		curDate := timestamp.Date(cur)
		his, err := json.ReadDir[historyItem]("history", curDate)
		// ignore the error that dir not exists
		if err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("cannot read history on %v: %w", curDate, err)
		}
		res = append(res, his...)
	}
	filteredRes, err := historyItems(res).filterByDateRange(from, to)
	if err != nil {
		return nil, fmt.Errorf("can not find history by date range:%w", err)
	}
	return filteredRes, nil
}

// relatedHistory returns historyItems related to the given t, if t is earlier than 4 o'clock, also returns historyItems one day before t.
func relatedHistory(t time.Time) (his historyItems, err error) {
	if t.Hour() <= 4 {
		his, err = historyByDateRange(t.AddDate(0, 0, -1), t)
	} else {
		his, err = historyByDate(t)
	}

	if err != nil {
		err = fmt.Errorf("error in relatedHistory: %w", err)
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
	his, err := relatedHistory(time.Now())
	if err != nil {
		return handler.CommonErr(nil, "cannot get current history")
	}
	players, err := his.lastPlayers()
	if err != nil {
		return handler.CommonErr(err, "cannot find last players")
	}

	for i, id := range players {
		if id == "" {
			return handler.CommonErr(nil, "missing player - player "+strconv.Itoa(i+1)+" is missing")
		}
	}
	j, err := json.From(struct {
		Players [4]string `json:"players"`
	}{players})
	if err != nil {
		return handler.CommonErr(err, "can not convert players to JSON")
	}
	w.Write(j)
	return
}
