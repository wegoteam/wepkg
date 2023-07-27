package crypto

import cryptoUtil "github.com/golang-module/dongle"

// EncodeBase64
// @Description: 对字符串进行 base64 编码，输出字符串
// @param: data
// @return string
func EncodeBase64(data string) string {
	return cryptoUtil.Encode.FromString(data).ByBase64().ToString()
}

// EncodeBase64ToBytes
// @Description: 对字符串进行 base64 编码，输出字节数组
// @param: data
// @return []byte
func EncodeBase64ToBytes(data []byte) []byte {
	return cryptoUtil.Encode.FromBytes(data).ByBase64().ToBytes()
}

// DecodeBase64
// @Description: 对字符串进行 base64 解码，输出字符串
// @param: data
// @return string
func DecodeBase64(data string) string {
	return cryptoUtil.Decode.FromString(data).ByBase64().ToString()
}

// DecodeBase64ToBytes
// @Description: 对字符串进行 base64 解码，输出字节数组
// @param: data
// @return []byte
func DecodeBase64ToBytes(data []byte) []byte {
	return cryptoUtil.Decode.FromBytes(data).ByBase64().ToBytes()
}
