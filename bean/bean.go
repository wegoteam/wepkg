package bean

import beanUtil "gopkg.in/jeevatkm/go-model.v1"

// BeanCopy
// @Description: 结构体拷贝
// @param: dst 目标结构体, 传入指针
// @param: src 源结构体, 传入值
// @return []error
func BeanCopy(dst, src interface{}) []error {
	return beanUtil.Copy(dst, src)
}

// BeanClone
// @Description: 结构体克隆
// @param: src 源结构体, 传入值
// @return interface{}
// @return error
func BeanClone(src interface{}) (interface{}, error) {
	return beanUtil.Clone(src)
}

// BeanToMap
// @Description: 结构体转map
// @param: src 源结构体, 传入值
// @return map[string]interface{}
// @return error
func BeanToMap(src interface{}) (map[string]interface{}, error) {
	return beanUtil.Map(src)
}
