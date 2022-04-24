package doudizhu

import (
	"boardgame-helper/router/handler"
	"boardgame-helper/utils/json"
	"boardgame-helper/utils/timestamp"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//!!!!!waring: most of those Testfunction are rubbishes!!!!!!!
func newPostReq(urlStr string, jsonFile []byte) *http.Request {
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonFile))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	return req
}

func newGetReq(urlStr string, kv map[string]string) *http.Request {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	for k, v := range kv {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return req
}

func response(r *http.Request, handler func(http.ResponseWriter, *http.Request) handler.Err) (jsonFile []byte, herr handler.Err) {
	w := httptest.NewRecorder()
	herr = handler(w, r)

	resp := w.Result()
	jsonFile, _ = io.ReadAll(resp.Body)
	return
}

func initTest() {
	//init part, we need it because in it will read and set the file path,
	//that will be used in the function we call
	jsonFile, err := os.Open("../../config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println(string(byteValue))

	json.InitConfig(string(byteValue), "", "")
}

func viewCompaire(testView view, t *testing.T) error {
	for ind, val := range testView.PlayerNames {
		if val != currentView.PlayerNames[ind] {
			t.Errorf("GetViewNow PlayerNames error, wanted %v, get %v", currentView.PlayerNames[ind], val)
		}
	}

	for ind1, val1 := range testView.DeltaPoints {
		if val1.Round != currentView.DeltaPoints[ind1].Round {
			t.Errorf("GetViewNow Round error, wanted %v, get %v", currentView.DeltaPoints[ind1].Round, val1.Round)
		}

		if val1.Timestamp != currentView.DeltaPoints[ind1].Timestamp {
			t.Errorf("GetViewNow Timestamp error, wanted %v, get %v", currentView.DeltaPoints[ind1].Timestamp, val1.Timestamp)
		}

		if val1.Enabled != currentView.DeltaPoints[ind1].Enabled {
			t.Errorf("GetViewNow Enabled error, wanted %v, get %v", currentView.DeltaPoints[ind1].Enabled, val1.Enabled)
		}

		for ind2, val2 := range testView.DeltaPoints[ind1].Deltas {
			if val2 != currentView.DeltaPoints[ind1].Deltas[ind2] {
				t.Errorf("GetViewNow Deltas error, wanted %v, get %v", currentView.DeltaPoints[ind1].Deltas[ind2], val2)
			}
		}
	}

	for ind, val := range testView.FinalPoints {
		if val != currentView.FinalPoints[ind] {
			t.Errorf("GetViewNow FinalPoints error, wanted %v, get %v", currentView.FinalPoints[ind], val)
		}
	}

	return nil
}

func makeCurView(t *testing.T) {
	//to remember: historyitem can not be called from other pakage, also the view, may need some change
	realItem, err := json.ReadDir[historyItem]("history", "Thu_Mar_31_2022")
	if err != nil {
		//why kann Errorf called by fmt and also by t?
		t.Errorf("Readfile error  = %v", err)
	}

	var realItems historyItems
	realItems = append(realItem)
	currentView = realItems.View()
}

func inputCompair(testInput inputItem, realInput inputItem, t *testing.T) error {
	if testInput.Timestamp != realInput.Timestamp {
		t.Errorf("error Timestamp, wanted %v, get %v", realInput.Timestamp, testInput.Timestamp)
	}

	if testInput.Stake != realInput.Stake {
		t.Errorf("error Stake, wanted %v, get %v", realInput.Stake, testInput.Stake)
	}

	if testInput.BonusTiles != realInput.BonusTiles {
		t.Errorf("error BonusTiles, wanted %v, get %v", realInput.BonusTiles, testInput.BonusTiles)
	}

	for ind, val := range testInput.Players {
		if val != realInput.Players[ind] {
			t.Errorf("error Players, wanted %v, get %v", realInput.Players[ind], val)
		}
	}

	if testInput.RawPoints != realInput.RawPoints {
		t.Errorf("error RawPoints, wanted %v, get %v", realInput.RawPoints, testInput.RawPoints)
	}

	if testInput.Winner != realInput.Winner {
		t.Errorf("error Winner, wanted %v, get %v", realInput.Winner, testInput.Winner)
	}

	for nam, val := range testInput.Weight {
		if val != realInput.Weight[nam] {
			t.Errorf("error Weights, wanted %v, get %v", realInput.Weight[nam], val)
		}
	}

	if testInput.Lord != realInput.Lord {
		t.Errorf("error Lord, wanted %v, get %v", realInput.Lord, testInput.Lord)
	}

	return nil
}

func TestGetViewNow(t *testing.T) {
	initTest()

	req := newGetReq("/api/doudizhu/view/now", map[string]string{})

	makeCurView(t)
	testByte, herr := response(req, GetViewNow)
	if herr.Error != nil {
		t.Errorf("response error = %v", herr.Error)
		panic("")
	}
	fmt.Println("the testview", string(testByte))

	testView, err1 := json.Parse[view](testByte)
	if err1 != nil {
		t.Errorf("parse error = %v", err1)
	}
	err2 := viewCompaire(testView, t)
	if err2 != nil {
		t.Errorf("GetViewNow error , view = %v", string(testByte))
	}
}

func TestGetViewByDate(t *testing.T) {
	initTest()

	req := newGetReq("/api/doudizhu/date", map[string]string{
		"timestamp": "2022-03-31T14:29:06.480Z",
	})

	testByte, herr := response(req, GetViewByDate)
	if herr.Error != nil {
		t.Errorf("response error = %v", herr.Error)
		panic("")
	}
	fmt.Println("the testview", string(testByte))

	makeCurView(t)
	testView, err1 := json.Parse[view](testByte)
	if err1 != nil {
		t.Errorf("parse error = %v", err1)
	}
	err2 := viewCompaire(testView, t)
	if err2 != nil {
		t.Errorf("GetViewNow error , view = %v", string(testByte))
	}
}

func TestUpdate(t *testing.T) {
	//the problem is, that time.Now() will gives us the current time, and wenn i am doing the test
	//the current time is different from the time stamp in Testdata.

	//And there is a risk. If the time goes through 24:00 wenn we palying the game, it will be in another day,
	//the update funktion will not be able to find the correct time.
}

func TestDisableHistory(t *testing.T) {
	initTest()

	req := newGetReq("/api/doudizhu/disable", map[string]string{
		"timestamp": "2022-03-31T14:29:06.480Z",
	})

	//the test daten should be setet "true" first.
	_, herr := response(req, DisableHistory)
	if herr.Error != nil {
		t.Errorf("response error = %v", herr.Error)
		panic("")
	}

	testbyte, err1 := response(req, GetViewByDate)
	if err1.Error != nil {
		t.Errorf("response error = %v", err1.Error)
	}

	fmt.Println("the view is :", string(testbyte))

	testview, err2 := json.Parse[view](testbyte)
	if err2 != nil {
		t.Errorf("Parse err = %v", err2)
	}

	for _, val := range testview.DeltaPoints {
		if val.Timestamp == "2022-03-31T14:29:06.480Z" && val.Enabled {
			t.Errorf("DisableHistory failed, want: false, get %v", val.Enabled)
		}
	}
}

func TestEnableHistory(t *testing.T) {
	initTest()

	req := newGetReq("/api/doudizhu/enable", map[string]string{
		"timestamp": "2022-03-31T14:29:06.480Z",
	})

	//the test daten should be setet "false" first.
	_, herr := response(req, EnableHistory)
	if herr.Error != nil {
		t.Errorf("response error = %v", herr.Error)
		panic("")
	}

	testbyte, err1 := response(req, GetViewByDate)
	if err1.Error != nil {
		t.Errorf("response error = %v", err1.Error)
	}

	fmt.Println("the view is :", string(testbyte))

	testview, err2 := json.Parse[view](testbyte)
	if err2 != nil {
		t.Errorf("Parse err = %v", err2)
	}

	for _, val := range testview.DeltaPoints {
		if val.Timestamp == "2022-03-31T14:29:06.480Z" && !val.Enabled {
			t.Errorf("DisableHistory failed, want: true, get %v", val.Enabled)
		}
	}
}

func TestAddInput(t *testing.T) {
	initTest()

	inputItemTest, err := json.ReadDir[historyItem]("history", "Thu_Mar_31_2022")
	if err != nil {
		t.Errorf("Readfile error  = %v", err)
	}

	inputByteTest, err := json.From(inputItemTest)
	if err != nil {
		t.Errorf("json.form error = %v", err)
	}

	req := newPostReq("/api/doudizhu/new", inputByteTest)

	//response error = json: cannot unmarshal array into Go value of type doudizhu.inputItem
	testByte, herr := response(req, AddInput)
	if !herr.Empty() {
		t.Errorf("response error = %v", herr.Error)
	}

	fmt.Println("the inputwanted", string(inputByteTest))
	fmt.Println("the ", string(testByte))

}

func TestEditInput(t *testing.T) {
	initTest()

	req := newGetReq("/api/doudizhu/edit", map[string]string{
		"timestamp": "2022-03-31T14:29:06.480Z",
	})

	realHistory, err := json.ReadFile[historyItem]("history", "Thu_Mar_31_2022", "14-29-06_Mar_31_2022.json")
	if err != nil {
		t.Errorf("Readfile error  = %v", err)
	}
	fmt.Println("real input timestamp = ", realHistory.InputItem.Timestamp)

	testByte, herr := response(req, EditInput)
	if herr.Error != nil {
		t.Errorf("response error = %v", herr.Error)
		panic("")
	}
	testInput, err := json.Parse[inputItem](testByte)
	if err != nil {
		t.Errorf("Parse err = %v", err)
	}
	fmt.Println("test input = ", string(testByte))

	err = inputCompair(testInput, realHistory.InputItem, t)
	if err != nil {
		t.Errorf("compair err = %v", err)
	}

}

func TestCurPlayers(t *testing.T) {
	initTest()

	req := newGetReq("/api/doudizhu/edit", map[string]string{
		"timestamp": "2022-03-31T14:29:06.480Z",
	})

	testByte, herr := response(req, CurPlayers)
	if herr.Error != nil {
		t.Errorf("response error = %v", herr.Error)
		panic("")
	}
	fmt.Println("testByte is: ", string(testByte))
	if string(testByte) != "{\"players\":[\"bai\",\"xiao\",\"jintian\",\"yunfan\"]}" {
		t.Errorf("error CurPlayers, wanted: %v, get %v ", "{\"players\":[\"bai\",\"xiao\",\"jintian\",\"yunfan\"]}", string(testByte))
	}

}

type deltaPoints struct {
	Ts  string         `json:"timestamp"`
	Pts map[string]int `json:"final_points"`
}

// testfiles to json
func testToHistory() {
	inputs, err := json.ReadDir[inputItem]("test", "Thu_Mar31_2022")
	if err != nil {
		panic(err)
	}
	for _, val := range inputs {
		jsf, err := json.From(val)
		if err != nil {
			panic(err)
		}
		req := newPostReq("/api/doudizhu/new", jsf)
		_, herr := response(req, AddInput)
		if !herr.Empty() {
			panic(herr.Message)
		}

	}
}

func TestDeltaPoints(t *testing.T) {
	initTest()

	testToHistory()

	dpts, err := json.ReadDir[deltaPoints]("test", "Thu_Mar31_2022")
	if err != nil {
		t.Errorf("error readDir: %v", err)
	}

	for _, x := range dpts {
		t.Run(x.Ts, func(t *testing.T) {
			gotTime, err := timestamp.Parse(x.Ts)
			if err != nil {
				panic(err)
			}

			his, err := historyByDateTime(gotTime)
			if err != nil {
				t.Errorf("Error historyByDate : %v", err)
			}

			for _, playerDetail := range his.PlayerDetails {
				if playerDetail.Deltapoints != x.Pts[playerDetail.Player] {
					t.Errorf("Error deltapoints, %v: want %v, got %v", playerDetail.Player, x.Pts[playerDetail.Player], playerDetail.Deltapoints)
				}

			}
		})
	}
}

// if !reflect.DeepEqual(gotHi, tt.wantHi) {
// 	t.Errorf("inputItem.History() = %v, want %v", gotHi, tt.wantHi)
// }

// req := newGetReq("/api/doudizhu/view/date", map[string]string{
// 	"timestamp": "2020",
// })
// resJson, herr := response(req, AddInput)
