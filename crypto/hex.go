package crypto

import cryptoUtil "github.com/golang-module/dongle"

// EncodeHex
// @Description: 对字符串进行 hex 编码，输出字符串
// @param: data
// @return string
func EncodeHex(data string) string {
	return cryptoUtil.Encode.FromString(data).ByHex().ToString()
}

// EncodeHexToBytes
// @Description: 对字符串进行 hex 编码，输出字节数组
// @param: data
// @return []byte
func EncodeHexToBytes(data []byte) []byte {
	return cryptoUtil.Encode.FromBytes(data).ByHex().ToBytes()
}

// DecodeHex
// @Description: 对字符串进行 hex 解码，输出字符串
// @param: data
// @return string
func DecodeHex(data string) string {
	return cryptoUtil.Decode.FromString(data).ByHex().ToString()
}

// DecodeHexToBytes
// @Description: 对字符串进行 hex 解码，输出字节数组
// @param: data
// @return []byte
func DecodeHexToBytes(data []byte) []byte {
	return cryptoUtil.Decode.FromBytes(data).ByHex().ToBytes()
}
