package json

// 引用：https://github.com/bytedance/sonic

import jsonUtil "github.com/bytedance/sonic"

// Marshal
// @Description: 序列化
// @param: val 传入对象
// @return string 序列化后的字符串
// @return error
func Marshal(val interface{}) (string, error) {
	return jsonUtil.MarshalString(val)
}

// Unmarshal
// @Description: 反序列化
// @param: buf 传入字符串
// @param: val 传入指针
// @return error
func Unmarshal(buf string, val interface{}) error {
	return jsonUtil.UnmarshalString(buf, val)
}
