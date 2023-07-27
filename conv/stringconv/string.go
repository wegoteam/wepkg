package stringconv

import (
	"encoding/json"
	"strconv"
	"strings"
)

// 驼峰 to 下划线
func CamelToUnderline(target string) string {
	return strings.Join(SplitUpper(target), "_")
}

// 下划线 to 驼峰
func UnderlineToCamel(target string) string {
	arr := strings.Split(target, "_")
	for i, val := range arr {
		arr[i] = UpperFirst(val)
	}
	return strings.Join(arr, "")
}

// 按照大小写分组
func SplitUpper(target string) []string {
	chars := []byte(target)
	result := []string{}
	str := ""
	for i := 0; i < len(target); i++ {
		char := chars[i]
		if i == 0 {

		} else if char >= 65 && char <= 90 {
			result = append(result, str)
			str = ""
		}

		str += string(char)
	}

	result = append(result, str)
	return result
}

// SplitLast 获取字符串被"."分离后的字符串数组中
// ，最后一个字符串。
func SplitLast(str string, sep string) string {
	strArr := strings.Split(str, sep)
	if strArr != nil && len(strArr) > 0 {
		return strArr[len(strArr)-1]
	}

	return str
}

// 首字母是否小写
func IsLowerFirst(str string) bool {
	first := str[0]
	if first >= 97 && first <= 122 {
		return true
	}
	return false
}

// 首字母是否大写
func IsUpperFirst(str string) bool {
	first := str[0]
	if first >= 65 && first <= 90 {
		return true
	}
	return false
}

// LowerFirst 首字母小写
func LowerFirst(str string) string {
	first := str[:1]
	return strings.Replace(str, first, strings.ToLower(first), 1)
}

// UpperFirst 首字母大写
func UpperFirst(str string) string {
	if len(str) == 0 {
		return str
	}

	first := str[:1]

	return strings.Replace(str, first, strings.ToUpper(first), 1)
}

// FormatStrval Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func FormatStrval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
