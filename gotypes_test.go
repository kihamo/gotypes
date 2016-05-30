package gotypes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TimeStruct struct {
	Time time.Time
}

type MapStruct struct {
	StructField string
}

type MapStructWithOptionalFieldValue struct {
	StructField string `json:",omitempty"`
}

type MapStructWithJsonTagForField struct {
	StructField string `json:"field"`
}

type MapStructWithJsonTagSkip struct {
	StructField string `json:"-"`
}

func Test_StringToTimeWithDateInRFC3339_ResultIsValid(t *testing.T) {
	outTime := time.Date(2015, time.December, 8, 23, 4, 2, 0, time.FixedZone("", 4*60*60))

	output := TimeStruct{}
	input := map[string]interface{}{
		"Time": "2015-12-08T23:04:02+04:00",
	}

	converter := NewConverter(input, &output)
	result := converter.GetOutput()

	assert.Equal(t, result.(*TimeStruct).Time.String(), outTime.String())
}

func Test_StringToTime_ResultIsValid(t *testing.T) {
	outTime := time.Date(2016, time.August, 19, 18, 55, 0, 0, time.UTC)

	output := TimeStruct{}
	input := map[string]interface{}{
		"Time": "19.08.2016 18:55:00",
	}

	converter := NewConverter(input, &output)
	result := converter.GetOutput()

	assert.Equal(t, result.(*TimeStruct).Time.String(), outTime.String())
}

func Test_MapMapsToMapStruct_ResultIsValid(t *testing.T) {
	output := map[string]MapStruct{}
	input := map[string]map[string]string{
		"MapField": {
			"StructField": "Value",
		},
	}

	converter := NewConverter(input, &output)
	valid := converter.Valid()

	assert.True(t, valid)
	assert.Equal(t, output["MapField"].StructField, "Value")
}

func Test_MapMapsWithEmptyFieldValueToMapStruct_ResultIsNotValid(t *testing.T) {
	output := map[string]MapStruct{}
	input := map[string]map[string]string{
		"MapField": {},
	}

	converter := NewConverter(input, &output)
	valid := converter.Valid()

	assert.False(t, valid)
	assert.Equal(t, converter.GetInvalidFields()[0], "{\"MapField\"}.StructField")
	assert.Equal(t, output["MapField"].StructField, "")
}

func Test_MapMapsToMapStructWithOptionalFieldValue_ResultIsValid(t *testing.T) {
	output := map[string]MapStructWithOptionalFieldValue{}
	input := map[string]map[string]string{
		"MapField": {},
	}

	converter := NewConverter(input, &output)
	valid := converter.Valid()

	assert.True(t, valid)
	assert.Equal(t, output["MapField"].StructField, "")
}

func Test_MapToStructWithJsonSkipField_ResultIsValid(t *testing.T) {
	output := MapStructWithJsonTagSkip{}
	input := map[string]string{
		"StructField": "testValue",
	}

	converter := NewConverter(input, &output)
	valid := converter.Valid()

	assert.True(t, valid)
	assert.Equal(t, output.StructField, "")
}

func Test_SliceBoolInInterfaceToSliceBool_ResultIsValid(t *testing.T) {
	output := []bool{}
	input := []interface{}{
		false,
		"0",
		"false",
		true,
		"1",
		"true",
	}

	converter := NewConverter(input, &output)
	valid := converter.Valid()

	assert.True(t, valid)
	assert.False(t, output[0])
	assert.False(t, output[1])
	assert.False(t, output[2])
	assert.True(t, output[3])
	assert.True(t, output[4])
	assert.True(t, output[5])
}

func Test_MapStringInterfaceToMapStringBool_ResultIsValid(t *testing.T) {
	output := map[string]bool{}
	input := map[string]interface{}{
		"OneFalse":  false,
		"TwoFalse":  "0",
		"TreeFalse": "false",
		"OneTrue":   true,
		"TwoTrue":   "1",
		"TreeTrue":  "true",
	}

	converter := NewConverter(input, &output)
	valid := converter.Valid()

	assert.True(t, valid)
	assert.False(t, output["OneFalse"])
	assert.False(t, output["TwoFalse"])
	assert.False(t, output["TreeFalse"])
	assert.True(t, output["OneTrue"])
	assert.True(t, output["TwoTrue"])
	assert.True(t, output["TreeTrue"])
}

func Test_StructToMapStringInterface_ResultIsValid(t *testing.T) {
	output := map[string]interface{}{}
	input := MapStruct{
		StructField: "test",
	}

	converter := NewConverter(input, &output)
	valid := converter.Valid()

	assert.True(t, valid)
	assert.Equal(t, output["StructField"], "test")
}

func Test_StructByPointerToMapStringInterface_ResultIsValid(t *testing.T) {
	output := map[string]interface{}{}
	input := MapStruct{
		StructField: "test",
	}

	converter := NewConverter(&input, &output)
	valid := converter.Valid()

	assert.True(t, valid)
	assert.Equal(t, output["StructField"], "test")
}

func Test_StructWithJsonTagToMapStringInterface_ResultIsValid(t *testing.T) {
	output := map[string]interface{}{}
	input := MapStructWithJsonTagForField{
		StructField: "test",
	}

	converter := NewConverter(input, &output)
	valid := converter.Valid()

	assert.True(t, valid)
	assert.Equal(t, output["field"], "test")
}
