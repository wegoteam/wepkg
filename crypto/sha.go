package crypto

import cryptoUtil "github.com/golang-module/dongle"

// EncryptSha1ToHex
// @Description: 对字符串进行 sha1 编码，输出 hex 字符串
// @param: data
// @return string
func EncryptSha1ToHex(data string) string {
	return cryptoUtil.Encrypt.FromString(data).BySha1().ToHexString()
}

// EncryptSha1ToBase64
// @Description: 对字符串进行 sha1 编码，输出 base64 字符串
// @param: data
// @return string
func EncryptSha1ToBase64(data string) string {
	return cryptoUtil.Encrypt.FromString(data).BySha1().ToBase64String()
}

// EncryptSha3ToHex
// @Description: 对字符串进行 sha3 编码，输出 hex 字符串
// 包含 sha3-224, sha3-256, sha3-384, sha3-512
// @param: data
// @param: size 224, 256, 384, 512
// @return string
func EncryptSha3ToHex(data string, size ...int) string {
	var num int
	if len(size) == 0 {
		num = 256
	} else {
		num = size[0]
	}
	return cryptoUtil.Encrypt.FromString(data).BySha3(num).ToHexString()
}

// EncryptSha3ToBase64
// @Description: 对字符串进行 sha3 编码，输出 hex 字符串
// 包含 sha3-224, sha3-256, sha3-384, sha3-512
// @param: data
// @param: size 224, 256, 384, 512
// @return string
func EncryptSha3ToBase64(data string, size ...int) string {
	var num int
	if len(size) == 0 {
		num = 256
	} else {
		num = size[0]
	}
	return cryptoUtil.Encrypt.FromString(data).BySha3(num).ToBase64String()
}

// EncryptSha256ToHex
// @Description: 对字符串进行 sha256 编码，输出 hex 字符串
// @param: data
// @return string
func EncryptSha256ToHex(data string) string {
	return cryptoUtil.Encrypt.FromString(data).BySha256().ToHexString()
}

// EncryptSha256ToBase64
// @Description: 对字符串进行 sha256 编码，输出 base64 字符串
// @param: data
// @return string
func EncryptSha256ToBase64(data string) string {
	return cryptoUtil.Encrypt.FromString(data).BySha256().ToBase64String()
}
