package gotypes

import (
	"time"
)

func ToTime(in interface{}) (t time.Time) {
	switch v := in.(type) {
	default:
		t, _ = time.Parse(time.RFC3339, ToString(v))
	}

	return t
}
