package crypto

import (
	"crypto/rsa"
	cryptoUtil "github.com/golang-module/dongle"
	opensslUtil "github.com/golang-module/dongle/openssl"
)

// EncryptRsaToHex
// @Description: 对字符串进行 rsa 编码，输出 hex 字符串
// @param: data
// @param: key
// @return string
func EncryptRsaToHex(data, key string) string {
	return cryptoUtil.Encrypt.FromString(data).ByRsa(key).ToHexString()
}

// EncryptRsaToBase64
// @Description: 对字符串进行 rsa 编码，输出 base64 字符串
// @param: data
// @param: key
// @return string
func EncryptRsaToBase64(data, key string) string {
	return cryptoUtil.Encrypt.FromString(data).ByRsa(key).ToBase64String()
}

// DecryptRsaFromHex
// @Description: 对 hex 字符串进行 rsa 解码，输出字符串
// @param: data
// @param: key
// @return string
func DecryptRsaFromHex(data, key string) string {
	return cryptoUtil.Decrypt.FromHexString(data).ByRsa(key).ToString()
}

// DecryptRsaFromBase64
// @Description: 对 base64 字符串进行 rsa 解码，输出字符串
// @param: data
// @param: key
// @return string
func DecryptRsaFromBase64(data, key string) string {
	return cryptoUtil.Decrypt.FromBase64String(data).ByRsa(key).ToString()
}

// GenKeyPkcs1Pair
// @Description: 生成 PKCS1 格式的 RSA 密钥对
// @return []byte 生成的公钥
// @return []byte 生成的私钥
func GenKeyPkcs1Pair() (publicKey, privateKey []byte) {
	publicKey, privateKey = opensslUtil.RSA.GenKeyPair(opensslUtil.PKCS1, 1024)
	return
}

// GenKeyPkcs8Pair
// @Description: 生成 PKCS8 格式的 RSA 密钥对
// @return []byte 生成的公钥
// @return []byte 生成的私钥
func GenKeyPkcs8Pair() (publicKey, privateKey []byte) {
	publicKey, privateKey = opensslUtil.RSA.GenKeyPair(opensslUtil.PKCS8, 2048)
	return
}

// VerifyKeyPair
// @Description: 验证公钥和私钥是否匹配
// @param: publicKey
// @param: privateKey
// @return bool
func VerifyKeyPair(publicKey, privateKey []byte) bool {
	return opensslUtil.RSA.VerifyKeyPair(publicKey, privateKey)
}

// IsPublicKey
// @Description: 验证是否为公钥
// @param: publicKey 公钥
// @return bool
func IsPublicKey(publicKey []byte) bool {
	return opensslUtil.RSA.IsPublicKey(publicKey)
}

// IsPrivateKey
// @Description: 验证是否为私钥
// @param: privateKey 私钥
// @return bool
func IsPrivateKey(privateKey []byte) bool {
	return opensslUtil.RSA.IsPrivateKey(privateKey)
}

// ParsePublicKey
// @Description: 解析公钥
// @param: publicKey 公钥
// @return *rsa.PublicKey
// @return error
func ParsePublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	return opensslUtil.RSA.ParsePublicKey(publicKey)
}

// ParsePrivateKey
// @Description: 解析私钥
// @param: privateKey 私钥
// @return *rsa.PrivateKey
// @return error
func ParsePrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	return opensslUtil.RSA.ParsePrivateKey(privateKey)
}

// ExportPublicKey
// @Description: 从 RSA 私钥里导出公钥
// @param: privateKey
// @return publicKey
// @return err
func ExportPublicKey(privateKey []byte) (publicKey []byte, err error) {
	return opensslUtil.RSA.ExportPublicKey(privateKey)
}
