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
	input      interface{}
	output     interface{}
	zeros      map[string]bool
	errors     []string
	calculated bool
	validated  bool
}

func NewConverter(input interface{}, output interface{}) *Converter {
	return &Converter{
		input:  input,
		output: output,
		zeros:  map[string]bool{},
		errors: []string{},
	}
}

func (c *Converter) Valid() bool {
	if !c.validated {
		c.validateAny(reflect.Indirect(reflect.ValueOf(c.output)), "")
		c.validated = true
	}

	return len(c.errors) == 0
}

func (c *Converter) GetInvalidFields() []string {
	c.calculation()

	return c.errors
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

func (c *Converter) isZero(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return value.IsNil()

	case reflect.Struct:
		z := true
		for i := 0; i < value.NumField(); i++ {
			z = z && c.isZero(value.Field(i))
		}
		return z
	}

	if value.CanInterface() {
		return reflect.Zero(value.Type()).Interface() == value.Interface()
	}

	return false
}

func (c *Converter) setZero(val reflect.Value, path string) {
	if path != "" && c.isZero(val) {
		c.zeros[path] = true
	}
}

func (c *Converter) calculation() {
	if c.calculated {
		return
	}

	c.fill(reflect.Indirect(reflect.ValueOf(c.output)), c.input, "")
}

func (c *Converter) fill(out reflect.Value, in interface{}, path string) {
	if out.Kind() != reflect.Ptr && out.Kind() != reflect.Interface {
		c.setZero(out, path)
	}

	switch out.Kind() {
	case reflect.Ptr:
		if out.IsValid() {
			if out.IsNil() {
				out.Set(reflect.New(out.Type().Elem()))
			}

			c.fill(out.Elem(), in, path)
		}

	case reflect.Interface:
		c.fill(out.Elem(), in, path)

	case reflect.Struct:
		if in, ok := in.(map[string]interface{}); ok {
			for i := 0; i < out.NumField(); i++ {
				name := c.getName(out.Type().Field(i))
				childPath := c.getPath(path, name)

				if value, ok := in[name]; ok {
					c.fill(out.Field(i), value, childPath)
				}
			}
		}

	case reflect.Map:
		if in, ok := in.(map[string]interface{}); ok {
			keyType := out.Type().Key()
			valueType := out.Type().Elem()

			out.Set(reflect.MakeMap(out.Type()))
			for mapKey, mapValue := range in {
				key := reflect.New(keyType).Elem()
				c.fill(key, mapKey, path)

				value := reflect.New(valueType).Elem()
				childPath := c.getPath(path, fmt.Sprintf("[%q]", mapKey))
				c.fill(value, mapValue, childPath)

				out.SetMapIndex(key, value)
			}
		}

	case reflect.Slice:
		if in, ok := in.([]interface{}); ok {
			out.Set(reflect.MakeSlice(out.Type(), len(in), cap(in)))
			for i := range in {
				c.fill(out.Index(i), in[i], c.getPath(path, fmt.Sprintf("[%d]", i)))
			}
		}

	case reflect.Bool:
		out.SetBool(ToBool(in))

	case reflect.String:
		out.SetString(ToString(in))

	case reflect.Uint:
		out.Set(reflect.ValueOf(ToUint(in)))

	case reflect.Uint8:
		out.Set(reflect.ValueOf(ToUint8(in)))

	case reflect.Uint16:
		out.Set(reflect.ValueOf(ToUint16(in)))

	case reflect.Uint32:
		out.Set(reflect.ValueOf(ToUint32(in)))

	case reflect.Uint64:
		out.SetUint(ToUint64(in))

	case reflect.Int:
		out.Set(reflect.ValueOf(ToInt(in)))

	case reflect.Int8:
		out.Set(reflect.ValueOf(ToInt8(in)))

	case reflect.Int16:
		out.Set(reflect.ValueOf(ToInt16(in)))

	case reflect.Int32:
		out.Set(reflect.ValueOf(ToInt32(in)))

	case reflect.Int64:
		out.SetInt(ToInt64(in))

	case reflect.Float32:
		out.Set(reflect.ValueOf(ToFloat32(in)))

	case reflect.Float64:
		out.SetFloat(ToFloat64(in))

	default:
	}
}

func (c *Converter) validateAny(out reflect.Value, path string) {
	out = reflect.Indirect(out)
	if !out.IsValid() {
		return
	}

	switch out.Kind() {
	case reflect.Struct:
		c.validateStruct(out, path)

	case reflect.Slice:
		for i := 0; i < out.Len(); i++ {
			c.validateAny(out.Index(i), c.getPath(path, fmt.Sprintf("[%d]", i)))
		}

	case reflect.Map:
		for _, n := range out.MapKeys() {
			c.validateAny(out.MapIndex(n), c.getPath(path, fmt.Sprintf("[%q]", n.String())))
		}
	}
}

func (c *Converter) validateStruct(value reflect.Value, path string) {
	for i := 0; i < value.Type().NumField(); i++ {
		f := value.Type().Field(i)
		val := value.FieldByName(f.Name)

		name := c.getName(f)
		childPath := c.getPath(path, name)
		valid := true

		if val.Kind() != reflect.Ptr {
			switch val.Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
				if val.IsNil() {
					valid = false
				}
			default:
				if !val.IsValid() {
					valid = false
				} else if c.isZero(val) {
					_, valid = c.zeros[childPath]
				}
			}
		}

		if valid {
			c.validateAny(val, childPath)
		} else {
			c.errors = append(c.errors, childPath)
		}
	}
}
