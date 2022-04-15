package players

import (
	"boardgame-helper/utils/json"
)

var playerMap map[string]string

func Import() (err error) {
	playerMap, err = json.ReadFile[map[string]string]("players.json")
	return
}

func NameToID(name string) (id string) {
	panic("not implemented") // TODO: Implement
}
