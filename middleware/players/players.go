package players

import (
	"boardgame-helper/utils/json"
	"fmt"
)

var playerMap map[string]string

func Import() (err error) {
	playerMap, err = json.ReadFile[map[string]string]("players.json")
	return
}

func IDToName(id string) (name string, err error) {

	if x, ok := playerMap[id]; ok {
		name = x
	} else {
		err = fmt.Errorf("cannot find player id:  %v", id)

	}
	return
}
