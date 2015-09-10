package gotypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// ToBool
//

func Test_BoolTrue_ToBoolTrue(t *testing.T) {
	val := true

	result := ToBool(val)

	assert.True(t, result)
}

func Test_BoolFalse_ToBoolFalse(t *testing.T) {
	val := false

	result := ToBool(val)

	assert.False(t, result)
}

func Test_StringNonEmpty_ToBoolTrue(t *testing.T) {
	val := "123"

	result := ToBool(val)

	assert.True(t, result)
}

func Test_StringEmpty_ToBoolFalse(t *testing.T) {
	val := ""

	result := ToBool(val)

	assert.False(t, result)
}

func Test_StringZero_ToBoolFalse(t *testing.T) {
	val := "0"

	result := ToBool(val)

	assert.False(t, result)
}

func Test_StringFalse_ToBoolFalse(t *testing.T) {
	val := "false"

	result := ToBool(val)

	assert.False(t, result)
}

func Test_IntOne_ToBoolTrue(t *testing.T) {
	val := 1

	result := ToBool(val)

	assert.True(t, result)
}

func Test_IntZero_ToBoolFalse(t *testing.T) {
	val := 0

	result := ToBool(val)

	assert.False(t, result)
}

func Test_FloatOne_ToBoolTrue(t *testing.T) {
	val := 1.0

	result := ToBool(val)

	assert.True(t, result)
}

//
// ToString
//

func Test_StringNonEmpty_ToString(t *testing.T) {
	val := "1"

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_BoolTrue_ToString(t *testing.T) {
	val := true

	result := ToString(val)

	assert.Equal(t, "true", result)
}

func Test_BoolFalse_ToString(t *testing.T) {
	val := false

	result := ToString(val)

	assert.Equal(t, "false", result)
}

func Test_IntOne_ToString(t *testing.T) {
	val := 1

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_IntZero_ToString(t *testing.T) {
	val := 0

	result := ToString(val)

	assert.Equal(t, "0", result)
}

func Test_Int8One_ToString(t *testing.T) {
	val := int8(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_Int16One_ToString(t *testing.T) {
	val := int16(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_Int32One_ToString(t *testing.T) {
	val := int32(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_Int64One_ToString(t *testing.T) {
	val := int64(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_UintOne_ToString(t *testing.T) {
	val := uint(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_Uint8One_ToString(t *testing.T) {
	val := uint8(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_Uint16One_ToString(t *testing.T) {
	val := uint16(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_Uint32One_ToString(t *testing.T) {
	val := uint32(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_Uint64One_ToString(t *testing.T) {
	val := uint64(1)

	result := ToString(val)

	assert.Equal(t, "1", result)
}

func Test_Float32One_ToString(t *testing.T) {
	val := float32(1)

	result := ToString(val)

	assert.Equal(t, "1.000000", result)
}

func Test_Float64One_ToString(t *testing.T) {
	val := float64(1)

	result := ToString(val)

	assert.Equal(t, "1.000000", result)
}

func Test_Interface_ToString(t *testing.T) {
	var val interface{}

	result := ToString(val)

	assert.Equal(t, "", result)
}

func Test_Bytes_ToString(t *testing.T) {
	val := []byte("1")

	result := ToString(val)

	assert.Equal(t, "1", result)
}

//
// ToInt
//

func Test_StringNonEmpty_ToInt(t *testing.T) {
	val := "123"

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_StringEmpty_ToInt(t *testing.T) {
	val := ""

	result := ToInt(val)

	assert.Exactly(t, result, 0)
}

func Test_StringZero_ToInt(t *testing.T) {
	val := "0"

	result := ToInt(val)

	assert.Exactly(t, result, 0)
}

func Test_StringFalse_ToInt(t *testing.T) {
	val := "false"

	result := ToInt(val)

	assert.Exactly(t, result, 0)
}

func Test_BoolTrue_ToInt(t *testing.T) {
	val := true

	result := ToInt(val)

	assert.Exactly(t, result, 1)
}

func Test_BoolFalse_ToInt(t *testing.T) {
	val := false

	result := ToInt(val)

	assert.Exactly(t, result, 0)
}

func Test_Bytes_ToInt(t *testing.T) {
	val := []byte("1")

	result := ToInt(val)

	assert.Exactly(t, result, 1)
}

func Test_Int_ToInt(t *testing.T) {
	val := 1

	result := ToInt(val)

	assert.Exactly(t, result, 1)
}

func Test_Int8_ToInt(t *testing.T) {
	val := int8(123)

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_Int16_ToInt(t *testing.T) {
	val := int16(123)

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_Int32_ToInt(t *testing.T) {
	val := int32(123)

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_Int64_ToInt(t *testing.T) {
	val := int64(123)

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_Uint_ToInt(t *testing.T) {
	val := uint(1)

	result := ToInt(val)

	assert.Exactly(t, result, 1)
}

func Test_Uit8_ToInt(t *testing.T) {
	val := uint8(123)

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_Uint16_ToInt(t *testing.T) {
	val := uint16(123)

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_Uint32_ToInt(t *testing.T) {
	val := uint32(123)

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_Uint64_ToInt(t *testing.T) {
	val := uint64(123)

	result := ToInt(val)

	assert.Exactly(t, result, 123)
}

func Test_Float32_ToInt(t *testing.T) {
	val := float32(1.0)

	result := ToInt(val)

	assert.Exactly(t, result, 1)
}

func Test_Float64_ToInt(t *testing.T) {
	val := float64(1.0)

	result := ToInt(val)

	assert.Exactly(t, result, 1)
}

//
// ToInt8
//

func Test_BoolTrue_ToInt8(t *testing.T) {
	val := true

	result := ToInt8(val)

	assert.Exactly(t, result, int8(1))
}

//
// ToInt16
//

func Test_BoolTrue_ToInt16(t *testing.T) {
	val := true

	result := ToInt16(val)

	assert.Exactly(t, result, int16(1))
}

//
// ToInt32
//

func Test_BoolTrue_ToInt32(t *testing.T) {
	val := true

	result := ToInt32(val)

	assert.Exactly(t, result, int32(1))
}

//
// ToUint
//

func Test_StringNonEmpty_ToUint(t *testing.T) {
	val := "1"

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_StringEmpty_ToUint(t *testing.T) {
	val := ""

	result := ToUint(val)

	assert.Exactly(t, result, uint(0))
}

func Test_StringZero_ToUint(t *testing.T) {
	val := "0"

	result := ToUint(val)

	assert.Exactly(t, result, uint(0))
}

func Test_StringFalse_ToUint(t *testing.T) {
	val := "false"

	result := ToUint(val)

	assert.Exactly(t, result, uint(0))
}

func Test_BoolTrue_ToUint(t *testing.T) {
	val := true

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_BoolFalse_ToUint(t *testing.T) {
	val := false

	result := ToUint(val)

	assert.Exactly(t, result, uint(0))
}

func Test_Bytes_ToUint(t *testing.T) {
	val := []byte("1")

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Int_ToUint(t *testing.T) {
	val := 1

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Int8_ToUint(t *testing.T) {
	val := int8(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Int16_ToUint(t *testing.T) {
	val := int16(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Int32_ToUint(t *testing.T) {
	val := int32(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Int64_ToUint(t *testing.T) {
	val := int64(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Uint_ToUint(t *testing.T) {
	val := uint(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Uit8_ToUint(t *testing.T) {
	val := uint8(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Uint16_ToUint(t *testing.T) {
	val := uint16(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Uint32_ToUint(t *testing.T) {
	val := uint32(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Uint64_ToUint(t *testing.T) {
	val := uint64(1)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Float32_ToUint(t *testing.T) {
	val := float32(1.0)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

func Test_Float64_ToUint(t *testing.T) {
	val := float64(1.0)

	result := ToUint(val)

	assert.Exactly(t, result, uint(1))
}

//
// ToUint8
//

func Test_BoolTrue_ToUint8(t *testing.T) {
	val := true

	result := ToUint8(val)

	assert.Exactly(t, result, uint8(1))
}

//
// ToUint16
//

func Test_BoolTrue_ToUint16(t *testing.T) {
	val := true

	result := ToUint16(val)

	assert.Exactly(t, result, uint16(1))
}

//
// ToUint32
//

func Test_BoolTrue_ToUint32(t *testing.T) {
	val := true

	result := ToUint32(val)

	assert.Exactly(t, result, uint32(1))
}

//
// Float32
//

func Test_BoolTrue_ToFloat32(t *testing.T) {
	val := true

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(1))
}

func Test_BoolFalse_ToFloat32(t *testing.T) {
	val := false

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(0))
}

func Test_StringNonEmpty_ToFloat32(t *testing.T) {
	val := "123"

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_StringEmpty_ToFloat32(t *testing.T) {
	val := ""

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(0))
}

func Test_StringZero_ToFloat32(t *testing.T) {
	val := "0"

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(0))
}

func Test_StringFalse_ToFloat32(t *testing.T) {
	val := "false"

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(0))
}

func Test_Bytes_ToFloat32(t *testing.T) {
	val := []byte("1")

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(1))
}

func Test_Int_ToFloat32(t *testing.T) {
	val := 1

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(1))
}

func Test_Int8_ToFloat32(t *testing.T) {
	val := int8(123)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_Int16_ToFloat32(t *testing.T) {
	val := int16(123)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_Int32_ToFloat32(t *testing.T) {
	val := int32(123)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_Int64_ToFloat32(t *testing.T) {
	val := int64(123)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_Uint_ToFloat32(t *testing.T) {
	val := uint(1)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(1))
}

func Test_Uit8_ToFloat32(t *testing.T) {
	val := uint8(123)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_Uint16_ToFloat32(t *testing.T) {
	val := uint16(123)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_Uint32_ToFloat32(t *testing.T) {
	val := uint32(123)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_Uint64_ToFloat32(t *testing.T) {
	val := uint64(123)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(123))
}

func Test_Float32_ToFloat32(t *testing.T) {
	val := float32(1.2)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(1.2))
}

func Test_Float64_ToFloat32(t *testing.T) {
	val := float64(1.2)

	result := ToFloat32(val)

	assert.Exactly(t, result, float32(1.2))
}
