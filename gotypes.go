package gotypes

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	FieldsSeparator = "."
)

type Converter struct {
	input           interface{}
	output          interface{}
	allowZeroFields map[string]bool
	invalidFields   []string
	calculated      bool
	validated       bool
}

func NewConverter(input interface{}, output interface{}) *Converter {
	return &Converter{
		input:           input,
		output:          output,
		allowZeroFields: map[string]bool{},
		invalidFields:   []string{},
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

	value := reflect.Indirect(reflect.ValueOf(c.output))

	c.findAllowZeroFields(value, "")
	c.fillOutput(value, c.input, "")
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

		if input, ok := input.(map[string]interface{}); ok {
			values = input
		} else {
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
			var value reflect.Value

			keyType := output.Type().Key()
			valueType := output.Type().Elem()

			output.Set(reflect.MakeMap(output.Type()))
			for i := range values {
				key := reflect.New(keyType).Elem()
				c.fillOutput(key, i, path)

				if valueType.Kind() != reflect.Interface {
					value = reflect.New(valueType).Elem()
					childPath := c.getPath(path, fmt.Sprintf("[%q]", i))
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
		inputValue := reflect.ValueOf(input)
		values := map[string]interface{}{}

		if input, ok := input.(map[string]interface{}); ok {
			values = input
		} else {
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
		output.SetBool(ToBool(input))

	case reflect.String:
		output.SetString(ToString(input))

	case reflect.Uint:
		output.Set(reflect.ValueOf(ToUint(input)))

	case reflect.Uint8:
		output.Set(reflect.ValueOf(ToUint8(input)))

	case reflect.Uint16:
		output.Set(reflect.ValueOf(ToUint16(input)))

	case reflect.Uint32:
		output.Set(reflect.ValueOf(ToUint32(input)))

	case reflect.Uint64:
		output.SetUint(ToUint64(input))

	case reflect.Int:
		output.Set(reflect.ValueOf(ToInt(input)))

	case reflect.Int8:
		output.Set(reflect.ValueOf(ToInt8(input)))

	case reflect.Int16:
		output.Set(reflect.ValueOf(ToInt16(input)))

	case reflect.Int32:
		output.Set(reflect.ValueOf(ToInt32(input)))

	case reflect.Int64:
		output.SetInt(ToInt64(input))

	case reflect.Float32:
		output.Set(reflect.ValueOf(ToFloat32(input)))

	case reflect.Float64:
		output.SetFloat(ToFloat64(input))

	}
}

func (c *Converter) findAllowZeroFields(output reflect.Value, path string) {
	switch output.Kind() {

	case reflect.Map:
		for _, n := range output.MapKeys() {
			c.findAllowZeroFields(output.MapIndex(n), c.getPath(path, fmt.Sprintf("[%q]", n.String())))
		}

	case reflect.Ptr:
		if path != "" {
			c.allowZeroFields[path] = true
		}

	case reflect.Slice:
		value := reflect.New(output.Type().Elem()).Elem()
		c.findAllowZeroFields(value, c.getPath(path, "[]"))

	case reflect.Struct:
		for i := 0; i < output.NumField(); i++ {
			field := output.Type().Field(i)
			fieldPath := c.getPath(path, c.getName(field))

			tag := field.Tag.Get("json")
			if tag != "" {
				parts := strings.Split(tag, ",")
				if len(parts) > 1 && parts[1] == "omitempty" {
					c.allowZeroFields[fieldPath] = true
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
				subPath := c.getPath(path, fmt.Sprintf("[%q]", n.String()))
				subFieldPath := c.getPath(fieldPath, fmt.Sprintf("[%q]", n.String()))

				c.validate(output.MapIndex(n), subPath, subFieldPath)
			}
		}

	case reflect.Chan, reflect.Func, reflect.Interface:
		valid = !output.IsNil()

	default:
		if !output.IsValid() || output.CanInterface() && reflect.Zero(output.Type()).Interface() == output.Interface() {
			valid = false
		}
	}

	if !valid && path != "" {
		if _, ok := c.allowZeroFields[path]; !ok {
			c.invalidFields = append(c.invalidFields, fieldPath)
		}
	}
}
