package gotypes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TimeStruct struct {
	Time time.Time
}

func Test_String_ToTime(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	outTime := time.Date(2015, time.December, 8, 23, 4, 2, 0, loc)

	output := TimeStruct{}
	input := map[string]interface{}{
		"Time": "2015-12-08T23:04:02+03:00",
	}

	converter := NewConverter(input, &output)
	result := converter.GetOutput()

	assert.Equal(t, result.(*TimeStruct).Time.String(), outTime.String())
}
