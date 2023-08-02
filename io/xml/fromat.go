package xml

//https://github.com/beevik/etree

import "encoding/xml"

// Marshal
// @Description: 序列化
// @param: val
// @return string
// @return error
func Marshal(val any) (string, error) {
	marshal, err := xml.Marshal(val)
	return string(marshal), err
}

// Unmarshal
// @Description: 反序列化
// @param: buf
// @param: val
// @return error
func Unmarshal(buf string, val any) error {
	return xml.Unmarshal([]byte(buf), val)
}
