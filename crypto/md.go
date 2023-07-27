package crypto

import cryptoUtil "github.com/golang-module/dongle"

// EncryptMd5ToHex
// @Description: 对字符串进行 md5 编码，输出 hex 字符串
// @param: data
// @return string
func EncryptMd5ToHex(data string) string {
	return cryptoUtil.Encrypt.FromString(data).ByMd5().ToHexString()
}

// EncryptMd5ToBase64
// @Description: 对字符串进行 md5 编码，输出 base64 字符串
// @param: data
// @return string
func EncryptMd5ToBase64(data string) string {
	return cryptoUtil.Encrypt.FromString(data).ByMd5().ToBase64String()
}

// EncryptMd5ToHexBytes
// @Description: 对字符串进行 md5 编码，输出 hex 字符串
// @param: data
// @return string
func EncryptMd5ToHexBytes(data string) []byte {
	return cryptoUtil.Encrypt.FromString(data).ByMd5().ToHexBytes()
}

// EncryptMd5ToBase64Bytes
// @Description: 对字符串进行 md5 编码，输出 base64 字符串
// @param: data
// @return string
func EncryptMd5ToBase64Bytes(data string) []byte {
	return cryptoUtil.Encrypt.FromString(data).ByMd5().ToBase64Bytes()
}
