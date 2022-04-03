package timestamp

import (
	"time"

	"github.com/relvacode/iso8601"
)

func Date(t time.Time) string {
	return t.Format("Mon_Jan_02_2006")
}

func DateTime(t time.Time) string {
	return t.Format("15-04-05_Jan_02_2006")
}

func Parse(ts string) (time.Time, error) {
	return iso8601.ParseString(ts)
}
