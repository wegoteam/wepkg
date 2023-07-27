package bean

//引用：https://github.com/fatih/structs
//引用：https://github.com/jeevatkm/go-model

import (
	beanUtil "gopkg.in/jeevatkm/go-model.v1"
	"reflect"
)

// IsZero
// @Description: 判断结构体是否为空
// @param: src 源结构体, 传入值
// @return bool 是否为空
func IsZero(src interface{}) bool {
	return beanUtil.IsZero(src)
}

// HasZero
// @Description: 判断结构体是否有空值
// @param: src 源结构体, 传入值
// @return bool 是否有空值
func HasZero(src interface{}) bool {
	return beanUtil.HasZero(src)
}

// HasFieldsZero
// @Description: 判断结构体指定字段是否有空值
// @param: src 源结构体, 传入值
// @param: fields 字段名
// @return string 字段名
// @return bool 是否有空值
func HasFieldsZero(src interface{}, fields ...string) (string, bool) {
	return beanUtil.IsZeroInFields(src, fields...)
}

// GetFields
// @Description: 获取结构体所有字段
// @param: src 源结构体, 传入值
// @return []reflect.StructField
// @return error
func GetFields(src interface{}) ([]reflect.StructField, error) {
	return beanUtil.Fields(src)
}

// GetKind
// @Description: 获取结构体指定字段的类型
// @param: src 源结构体, 传入值
// @param: field 字段名
// @return reflect.Kind
// @return error
func GetKind(src interface{}, field string) (reflect.Kind, error) {
	return beanUtil.Kind(src, field)
}

// GetTag
// @Description: 获取结构体指定字段的tag
// @param: src 源结构体, 传入值
// @param: field 字段名
// @return reflect.StructTag
// @return error
func GetTag(src interface{}, field string) (reflect.StructTag, error) {
	return beanUtil.Tag(src, field)
}

// GetTags
// @Description: 获取结构体所有字段的tag
// @param: src 源结构体, 传入值
// @return map[string]reflect.StructTag
// @return error
func GetTags(src interface{}) (map[string]reflect.StructTag, error) {
	return beanUtil.Tags(src)
}

// GetFieldVal
// @Description: 获取结构体指定字段的值
// @param: src
// @param: field
// @return interface{}
// @return error
func GetFieldVal(src interface{}, field string) (interface{}, error) {
	return beanUtil.Get(src, field)
}

// SetFieldVal
// @Description: 设置结构体指定字段的值
// @param: src 源结构体, 传入指针
// @param: field 字段名
// @param: value 字段值
// @return error
func SetFieldVal(src interface{}, field string, value interface{}) error {
	return beanUtil.Set(src, field, value)
}
