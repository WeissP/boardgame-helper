package doudizhu

import (
	"boardgame-helper/utils/timestamp"
	"time"
)

type historyItem struct{}

func (hi historyItem) write() {
	panic("not implemented") // TODO: Implement
}

func (hi *historyItem) toggle(status bool) {
	panic("not implemented") // TODO: Implement
}

type historyItems []historyItem

func (his historyItems) View() view {
	panic("not implemented") // TODO: Implement
}

func (his historyItems) write() {
	for _, x := range his {
		x.write()
	}
}

func historyByDate(t time.Time) historyItems {
	date := timestamp.Date(t)
	_ = date
	panic("not implemented") // TODO: Implement
}

func historyByDateTime(t time.Time) historyItem {
	dateTime := timestamp.DateTime(t)
	_ = dateTime
	panic("not implemented") // TODO: Implement
}

