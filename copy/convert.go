package copy

import (
	"reflect"
	"strconv"
)

// 转化成map类型的值
func Convert2MapValue(val reflect.Value) (r reflect.Value) {
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32:
		r = reflect.ValueOf(float64(val.Int()))
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		r = reflect.ValueOf(float64(val.Uint()))
	case reflect.Int, reflect.Int64:
		a := val.Int()
		if a > 2147483647 {
			r = reflect.ValueOf(strconv.Itoa(int(a)))
		} else {
			r = reflect.ValueOf(float64(a))
		}
	case reflect.Uint, reflect.Uint64:
		a := val.Uint()
		if a > 4294967295 {
			r = reflect.ValueOf(strconv.Itoa(int(a)))
		} else {
			r = reflect.ValueOf(float64(a))
		}
	case reflect.Float32, reflect.Float64:
		r = reflect.ValueOf(val.Float())
	case reflect.Bool:
		r = reflect.ValueOf(val.Bool())
	case reflect.String:
		r = reflect.ValueOf(val.String())
	default:
		r = val
	}
	return
}

func Convert2String(val reflect.Value) (r reflect.Value) {
	switch val.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r = reflect.ValueOf(strconv.Itoa(int(val.Uint())))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		r = reflect.ValueOf(strconv.Itoa(int(val.Int())))
	case reflect.Float32, reflect.Float64:
		r = reflect.ValueOf(strconv.FormatFloat(val.Float(), 'f', -1, 64))
	case reflect.String:
		r = val
	default:
		r = reflect.ValueOf("")
	}
	return
}

// 转换成Int型
func Convert2Int(val reflect.Value) (r reflect.Value) {
	switch val.Kind() {
	case reflect.String:
		i, err := strconv.Atoi(val.Interface().(string))
		if err != nil {
			r = reflect.ValueOf(0)
		} else {
			r = reflect.ValueOf(i)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r = reflect.ValueOf(int(val.Uint()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		r = reflect.ValueOf(val.Int())
	case reflect.Float32, reflect.Float64:
		r = reflect.ValueOf(int(val.Float()))
	default:
		r = reflect.ValueOf(0)
	}

	return
}

func Convert2Float(val reflect.Value) (r reflect.Value) {
	switch val.Kind() {
	case reflect.String:
		i, err := strconv.ParseFloat(val.Interface().(string), 64)
		if err != nil {
			r = reflect.ValueOf(0.0)
		} else {
			r = reflect.ValueOf(i)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r = reflect.ValueOf(float64(val.Uint()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		r = reflect.ValueOf(float64(val.Int()))
	case reflect.Float32, reflect.Float64:
		r = reflect.ValueOf(float64(val.Float()))
	default:
		r = reflect.ValueOf(0.0)
	}

	return
}

func convert2Bool(val reflect.Value) (r reflect.Value) {
	switch val.Kind() {
	case reflect.String:
		data := val.Interface().(string)
		if data == "true" || data == "1" {
			r = reflect.ValueOf(true)
		} else {
			r = reflect.ValueOf(false)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		r = reflect.ValueOf(float64(val.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r = reflect.ValueOf(val.Int() > 0)
	case reflect.Float32, reflect.Float64:
		r = reflect.ValueOf(val.Int() > 0)
	default:
		r = reflect.ValueOf(false)
	}
	return
}

func convert2StringNotReflect(val interface{}) (r string) {
	switch a := val.(type) {
	case uint:
		r = strconv.Itoa(int(a))
	case uint8:
		r = strconv.Itoa(int(a))
	case uint16:
		r = strconv.Itoa(int(a))
	case uint32:
		r = strconv.Itoa(int(a))
	case uint64:
		r = strconv.Itoa(int(a))
	case int:
		r = strconv.Itoa(int(a))
	case int8:
		r = strconv.Itoa(int(a))
	case int16:
		r = strconv.Itoa(int(a))
	case int32:
		r = strconv.Itoa(int(a))
	case int64:
		r = strconv.Itoa(int(a))
	case float32:
		r = strconv.FormatFloat(float64(a), 'f', -1, 64)
	case float64:
		r = strconv.FormatFloat(float64(a), 'f', -1, 64)
	default:
		r = ""
	}
	return
}

// 转换成Int型
func convert2IntNotReflect(val interface{}) (r int) {
	switch a := val.(type) {
	case string:
		i, err := strconv.Atoi(a)
		if err != nil {
			r = 0
		} else {
			r = i
		}
	case uint:
		r = int(a)
	case uint8:
		r = int(a)
	case uint16:
		r = int(a)
	case uint32:
		r = int(a)
	case uint64:
		r = int(a)
	case int:
		r = int(a)
	case int8:
		r = int(a)
	case int16:
		r = int(a)
	case int32:
		r = int(a)
	case int64:
		r = int(a)
	case float32:
		r = int(a)
	case float64:
		r = int(a)
	default:
		r = 0
	}

	return
}

func convert2FloatNotReflect(val interface{}) (r float64) {
	switch a := val.(type) {
	case string:
		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			r = 0.0
		} else {
			r = i
		}
	case uint:
		r = float64(a)
	case uint8:
		r = float64(a)
	case uint16:
		r = float64(a)
	case uint32:
		r = float64(a)
	case uint64:
		r = float64(a)
	case int:
		r = float64(a)
	case int8:
		r = float64(a)
	case int16:
		r = float64(a)
	case int32:
		r = float64(a)
	case int64:
		r = float64(a)
	case float32:
		r = float64(a)
	case float64:
		r = float64(a)
	default:
		r = a.(float64)
	}

	return
}

func convert2BoolNotReflect(val interface{}) (r bool) {
	switch a := val.(type) {
	case uint:
		r = a > 0
	case uint8:
		r = a > 0
	case uint16:
		r = a > 0
	case uint32:
		r = a > 0
	case uint64:
		r = a > 0
	case int:
		r = a > 0
	case int8:
		r = a > 0
	case int16:
		r = a > 0
	case int32:
		r = a > 0
	case int64:
		r = a > 0
	case float32:
		r = a > 0
	case float64:
		r = a > 0
	case string:
		if a == "true" || a == "1" {
			r = true
		} else {
			r = false
		}

	default:
		r = false
	}
	return
}
