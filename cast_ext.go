package gotypes

import (
	"time"
)

func ToTime(in interface{}) (t time.Time) {
	var err error
	v := ToString(in)

	t, err = time.Parse(time.RFC3339, v)

	if err != nil {
		t, err = time.Parse("02.01.2006 15:04:05", v)
	}

	return t
}
