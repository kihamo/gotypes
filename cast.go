package gotypes

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToBool(in interface{}) bool {
	var castIn bool

	switch v := in.(type) {
	case bool:
		castIn = v

	case []byte:
		castIn = ToBool(string(v))

	case string:
		castIn = v != "" && v != "0" && v != "false"

	case int:
		castIn = v != 0

	case int8:
		castIn = v != 0

	case int16:
		castIn = v != 0

	case int32:
		castIn = v != 0

	case int64:
		castIn = v != 0

	case uint:
		castIn = v != 0

	case uint8:
		castIn = v != 0

	case uint16:
		castIn = v != 0

	case uint32:
		castIn = v != 0

	case uint64:
		castIn = v != 0

	case float32:
		castIn = v != 0

	case float64:
		castIn = v != 0

	default:
		if reflect.ValueOf(v).Kind() == reflect.Ptr {
			return ToBool(reflect.Indirect(reflect.ValueOf(v)).Interface())
		}
	}

	return castIn
}

func ToString(in interface{}) string {
	var castIn string

	switch v := in.(type) {
	case string:
		castIn = v

	case bool:
		castIn = strconv.FormatBool(v)

	case []byte:
		castIn = string(v)

	case int:
		castIn = strconv.Itoa(v)

	case int8:
		castIn = strconv.FormatInt(int64(v), 10)

	case int16:
		castIn = strconv.FormatInt(int64(v), 10)

	case int32:
		castIn = strconv.FormatInt(int64(v), 10)

	case int64:
		castIn = strconv.FormatInt(v, 10)

	case uint:
		castIn = strconv.FormatUint(uint64(v), 10)

	case uint8:
		castIn = strconv.FormatUint(uint64(v), 10)

	case uint16:
		castIn = strconv.FormatUint(uint64(v), 10)

	case uint32:
		castIn = strconv.FormatUint(uint64(v), 10)

	case uint64:
		castIn = strconv.FormatUint(v, 10)

	case float32:
		castIn = strconv.FormatFloat(float64(v), 'f', 6, 32)

	case float64:
		castIn = strconv.FormatFloat(v, 'f', 6, 64)

	default:
		if stringer, ok := v.(fmt.Stringer); ok {
			return stringer.String()
		}

		if stringer, ok := v.(fmt.GoStringer); ok {
			return stringer.GoString()
		}

		if cast, ok := in.(string); ok {
			castIn = cast
		} else if reflect.ValueOf(v).Kind() == reflect.Ptr {
			return ToString(reflect.Indirect(reflect.ValueOf(v)).Interface())
		}
	}

	return castIn
}

func ToUint(in interface{}) uint {
	return uint(ToUint64(in))
}

func ToUint8(in interface{}) uint8 {
	return uint8(ToUint64(in))
}

func ToUint16(in interface{}) uint16 {
	return uint16(ToUint64(in))
}

func ToUint32(in interface{}) uint32 {
	return uint32(ToUint64(in))
}

func ToUint64(in interface{}) uint64 {
	var castIn uint64

	switch v := in.(type) {
	case string:
		castIn, _ = strconv.ParseUint(v, 10, 64)

	case bool:
		if v {
			castIn = uint64(1)
		}

	case []byte:
		castIn = ToUint64(string(v))

	case int64:
		castIn = uint64(v)

	case int:
		castIn = uint64(v)

	case int8:
		castIn = uint64(v)

	case int16:
		castIn = uint64(v)

	case int32:
		castIn = uint64(v)

	case uint:
		castIn = uint64(v)

	case uint8:
		castIn = uint64(v)

	case uint16:
		castIn = uint64(v)

	case uint32:
		castIn = uint64(v)

	case uint64:
		castIn = v

	case float32:
		castIn = uint64(v)

	case float64:
		castIn = uint64(v)

	default:
		if reflect.ValueOf(v).Kind() == reflect.Ptr {
			return ToUint64(reflect.Indirect(reflect.ValueOf(v)).Interface())
		}
	}

	return castIn
}

func ToInt(in interface{}) int {
	return int(ToInt64(in))
}

func ToInt8(in interface{}) int8 {
	return int8(ToInt64(in))
}

func ToInt16(in interface{}) int16 {
	return int16(ToInt64(in))
}

func ToInt32(in interface{}) int32 {
	return int32(ToInt64(in))
}

func ToInt64(in interface{}) int64 {
	var castIn int64

	switch v := in.(type) {
	case string:
		t, _ := strconv.Atoi(v)
		castIn = int64(t)

	case bool:
		if v {
			castIn = 1
		}

	case []byte:
		castIn = ToInt64(string(v))

	case int64:
		castIn = v

	case int:
		castIn = int64(v)

	case int8:
		castIn = int64(v)

	case int16:
		castIn = int64(v)

	case int32:
		castIn = int64(v)

	case uint:
		castIn = int64(v)

	case uint8:
		castIn = int64(v)

	case uint16:
		castIn = int64(v)

	case uint32:
		castIn = int64(v)

	case uint64:
		castIn = int64(v)

	case float32:
		castIn = int64(v)

	case float64:
		castIn = int64(v)

	default:
		if reflect.ValueOf(v).Kind() == reflect.Ptr {
			return ToInt64(reflect.Indirect(reflect.ValueOf(v)).Interface())
		}
	}

	return castIn
}

func ToFloat32(in interface{}) float32 {
	return float32(ToFloat64(in))
}

func ToFloat64(in interface{}) float64 {
	var castIn float64

	switch v := in.(type) {
	case string:
		castIn, _ = strconv.ParseFloat(v, 64)

	case bool:
		if v {
			castIn = float64(1)
		}

	case []byte:
		castIn = ToFloat64(string(v))

	case float64:
		castIn = v

	case int:
		castIn = float64(v)

	case int8:
		castIn = float64(v)

	case int16:
		castIn = float64(v)

	case int32:
		castIn = float64(v)

	case int64:
		castIn = float64(v)

	case uint:
		castIn = float64(v)

	case uint8:
		castIn = float64(v)

	case uint16:
		castIn = float64(v)

	case uint32:
		castIn = float64(v)

	case uint64:
		castIn = float64(v)

	case float32:
		castIn = float64(v)

	default:
		if reflect.ValueOf(v).Kind() == reflect.Ptr {
			return ToFloat64(reflect.Indirect(reflect.ValueOf(v)).Interface())
		}
	}

	return castIn
}
