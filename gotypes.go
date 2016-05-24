package gotypes

//go:generate goimports -w ./

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

const (
	FieldsSeparator = "."
)

var (
	timeType = reflect.TypeOf(time.Time{})
)

type Converter struct {
	input                 interface{}
	output                interface{}
	allowZeroFields       map[string]bool
	allowZeroFieldsByMask []*regexp.Regexp
	setValueFields        map[string]bool
	invalidFields         []string
	calculated            bool
	validated             bool
}

func NewConverter(input interface{}, output interface{}) *Converter {
	return &Converter{
		input:                 input,
		output:                output,
		allowZeroFields:       map[string]bool{},
		allowZeroFieldsByMask: []*regexp.Regexp{},
		setValueFields:        map[string]bool{},
		invalidFields:         []string{},
	}
}

func (c *Converter) Valid() bool {
	if !c.validated {
		c.calculation()
		c.validate(reflect.ValueOf(c.output), "", "")
		c.validated = true
	}

	return len(c.invalidFields) == 0
}

func (c *Converter) GetInvalidFields() []string {
	c.calculation()

	return c.invalidFields
}

func (c *Converter) GetInput() interface{} {
	return c.input
}

func (c *Converter) GetOutput() interface{} {
	c.calculation()

	return c.output
}

func (c *Converter) getPath(parentPath string, childPath string) string {
	prefix := FieldsSeparator
	if parentPath == "" {
		prefix = ""
	}

	return parentPath + prefix + childPath
}

func (c *Converter) getName(field reflect.StructField) string {
	name := field.Tag.Get("json")

	if name != "" {
		name = strings.Replace(name, ",omitempty", "", 1)
	}

	if name == "" {
		return field.Name
	}

	return name
}

func (c *Converter) calculation() {
	if c.calculated {
		return
	}

	in := reflect.Indirect(reflect.ValueOf(c.input)).Interface()
	out := reflect.Indirect(reflect.ValueOf(c.output))

	c.findAllowZeroFields(out, "")
	c.fillOutput(out, in, "")
	c.calculated = true
}

func (c *Converter) fillOutput(output reflect.Value, input interface{}, path string) {
	switch output.Kind() {

	case reflect.Ptr:
		if output.IsValid() {
			if output.IsNil() {
				output.Set(reflect.New(output.Type().Elem()))
			}

			c.fillOutput(output.Elem(), input, path)
		}

	case reflect.Interface:
		c.fillOutput(output.Elem(), input, path)

	case reflect.Map:
		inputValue := reflect.ValueOf(input)
		values := map[string]interface{}{}

		if inputCast, ok := input.(map[string]interface{}); ok {
			values = inputCast
		} else {
			switch inputValue.Kind() {
			case reflect.Map:
				for _, n := range inputValue.MapKeys() {
					values[n.String()] = inputValue.MapIndex(n).Interface()
				}

			case reflect.Struct:
				for i := 0; i < inputValue.NumField(); i++ {
					field := inputValue.Type().Field(i)
					values[c.getName(field)] = inputValue.FieldByName(field.Name).Interface()
				}
			}
		}

		if len(values) > 0 {
			var value reflect.Value

			keyType := output.Type().Key()
			valueType := output.Type().Elem()

			output.Set(reflect.MakeMap(output.Type()))
			for i := range values {
				key := reflect.New(keyType).Elem()
				c.fillOutput(key, i, path)

				if valueType.Kind() != reflect.Interface {
					value = reflect.New(valueType).Elem()
					childPath := c.getPath(path, fmt.Sprintf("{%q}", i))
					c.fillOutput(value, values[i], childPath)
				} else {
					value = reflect.ValueOf(values[i])
				}

				output.SetMapIndex(key, value)
			}
		}

	case reflect.Slice:
		inputValue := reflect.ValueOf(input)

		if inputValue.Kind() == reflect.Slice {
			output.Set(reflect.MakeSlice(output.Type(), inputValue.Len(), inputValue.Cap()))

			for i := 0; i < output.Len(); i++ {
				c.fillOutput(output.Index(i), inputValue.Index(i).Interface(), c.getPath(path, fmt.Sprintf("[%d]", i)))
			}
		}

	case reflect.Struct:
		// Custom types
		if output.Type() == timeType {
			c.setValueFields[path] = true
			output.Set(reflect.ValueOf(ToTime(input)))

			break
		}

		values := map[string]interface{}{}

		if inputCast, ok := input.(map[string]interface{}); ok {
			values = inputCast
		} else {
			inputValue := reflect.ValueOf(input)

			switch inputValue.Kind() {
			case reflect.Map:
				for _, n := range inputValue.MapKeys() {
					values[n.String()] = inputValue.MapIndex(n).Interface()
				}

			case reflect.Struct:
				for i := 0; i < inputValue.NumField(); i++ {
					field := inputValue.Type().Field(i)
					values[field.Name] = inputValue.FieldByName(field.Name).Interface()
				}
			}
		}

		if len(values) > 0 {
			for i := 0; i < output.NumField(); i++ {
				name := c.getName(output.Type().Field(i))
				childPath := c.getPath(path, name)

				if value, ok := values[name]; ok {
					c.fillOutput(output.Field(i), value, childPath)
				}
			}
		}

	case reflect.Bool:
		c.setValueFields[path] = true
		output.SetBool(ToBool(input))

	case reflect.String:
		c.setValueFields[path] = true
		output.SetString(ToString(input))

	case reflect.Uint:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToUint(input)))

	case reflect.Uint8:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToUint8(input)))

	case reflect.Uint16:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToUint16(input)))

	case reflect.Uint32:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToUint32(input)))

	case reflect.Uint64:
		c.setValueFields[path] = true
		output.SetUint(ToUint64(input))

	case reflect.Int:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToInt(input)))

	case reflect.Int8:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToInt8(input)))

	case reflect.Int16:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToInt16(input)))

	case reflect.Int32:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToInt32(input)))

	case reflect.Int64:
		c.setValueFields[path] = true
		output.SetInt(ToInt64(input))

	case reflect.Float32:
		c.setValueFields[path] = true
		output.Set(reflect.ValueOf(ToFloat32(input)))

	case reflect.Float64:
		c.setValueFields[path] = true
		output.SetFloat(ToFloat64(input))

	}
}

func (c *Converter) setAllowZeroFieldsPath(path string) {
	if strings.Contains(path, "{*}") {
		path = regexp.QuoteMeta(path)
		path = strings.Replace(path, `\{\*\}`, `\{.*?\}`, -1)

		re := regexp.MustCompile("^" + path + "$")
		c.allowZeroFieldsByMask = append(c.allowZeroFieldsByMask, re)
	} else {
		c.allowZeroFields[path] = true
	}
}

func (c *Converter) findAllowZeroFields(output reflect.Value, path string) {
	switch output.Kind() {

	case reflect.Map:
		if len(output.MapKeys()) > 0 {
			for _, n := range output.MapKeys() {
				c.findAllowZeroFields(output.MapIndex(n), c.getPath(path, fmt.Sprintf("{%q}", n.String())))
			}
		} else {
			c.findAllowZeroFields(reflect.New(output.Type().Elem()).Elem(), c.getPath(path, "{*}"))
		}

	case reflect.Ptr:
		if path != "" {
			c.setAllowZeroFieldsPath(path)
		}

		c.findAllowZeroFields(reflect.New(output.Type().Elem()).Elem(), path)

	case reflect.Slice:
		c.findAllowZeroFields(reflect.New(output.Type().Elem()).Elem(), c.getPath(path, "[]"))

	case reflect.Struct:
		for i := 0; i < output.NumField(); i++ {
			field := output.Type().Field(i)
			fieldPath := c.getPath(path, c.getName(field))

			tag := field.Tag.Get("json")
			if tag != "" {
				parts := strings.Split(tag, ",")
				if len(parts) > 1 && parts[1] == "omitempty" {
					c.setAllowZeroFieldsPath(fieldPath)
				}
			}

			c.findAllowZeroFields(output.FieldByName(field.Name), fieldPath)
		}

	}
}

func (c *Converter) validate(output reflect.Value, path string, fieldPath string) {
	output = reflect.Indirect(output)
	valid := true

	switch output.Kind() {
	case reflect.Struct:
		for i := 0; i < output.NumField(); i++ {
			field := output.Type().Field(i)
			val := output.FieldByName(field.Name)

			subPath := c.getPath(path, c.getName(field))
			subFieldPath := c.getPath(fieldPath, c.getName(field))

			c.validate(val, subPath, subFieldPath)
		}

	case reflect.Slice:
		valid = !output.IsNil()

		if valid {
			for i := 0; i < output.Len(); i++ {
				subPath := c.getPath(path, "[]")
				subFieldPath := c.getPath(fieldPath, fmt.Sprintf("[%d]", i))

				c.validate(output.Index(i), subPath, subFieldPath)
			}
		}

	case reflect.Map:
		valid = !output.IsNil()

		if valid {
			for _, n := range output.MapKeys() {
				subPath := c.getPath(path, fmt.Sprintf("{%q}", n.String()))
				subFieldPath := c.getPath(fieldPath, fmt.Sprintf("{%q}", n.String()))

				c.validate(output.MapIndex(n), subPath, subFieldPath)
			}
		}

	case reflect.Chan, reflect.Func, reflect.Interface:
		valid = !output.IsNil()

	default:
		valid = output.IsValid()

		if valid && output.CanInterface() && reflect.Zero(output.Type()).Interface() == output.Interface() {
			_, valid = c.setValueFields[fieldPath]
		}
	}

	if !valid && path != "" {
		_, valid = c.allowZeroFields[path]

		if !valid && len(c.allowZeroFieldsByMask) > 0 {
			for _, re := range c.allowZeroFieldsByMask {
				if valid = re.MatchString(path); valid {
					break
				}
			}
		}

		if !valid {
			c.invalidFields = append(c.invalidFields, fieldPath)
		}
	}
}
