package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/crypto"
	"testing"
)

func TestCrypto(t *testing.T) {
	fmt.Printf("hex 编码: %v\n", crypto.EncodeHex("hello world"))
	fmt.Printf("hex 编码: %s\n", crypto.EncodeHexToBytes([]byte("hello world")))
	fmt.Printf("hex 解码: %v\n", crypto.DecodeHex("68656c6c6f20776f726c64"))
	fmt.Printf("hex 解码: %s\n", crypto.DecodeHexToBytes([]byte("68656c6c6f20776f726c64")))
	fmt.Printf("base64 编码: %v\n", crypto.EncodeBase64("hello world"))
	fmt.Printf("base64 编码: %s\n", crypto.EncodeBase64ToBytes([]byte("hello world")))
	fmt.Printf("base64 解码: %v\n", crypto.DecodeBase64("aGVsbG8gd29ybGQ="))
	fmt.Printf("base64 解码: %s\n", crypto.DecodeBase64ToBytes([]byte("aGVsbG8gd29ybGQ=")))
	fmt.Printf("md5 编码: %v\n", crypto.EncryptMd5ToHex("hello world"))
	fmt.Printf("md5 编码: %v\n", crypto.EncryptMd5ToBase64("hello world"))
	fmt.Printf("md5 编码: %s\n", crypto.EncryptMd5ToHexBytes("hello world"))
	fmt.Printf("md5 编码: %s\n", crypto.EncryptMd5ToBase64Bytes("hello world"))
	fmt.Printf("sha1 编码: %v\n", crypto.EncryptSha1ToHex("hello world"))
	fmt.Printf("sha1 编码: %v\n", crypto.EncryptSha1ToBase64("hello world"))
	fmt.Printf("sha3-224 编码: %v\n", crypto.EncryptSha3ToHex("hello world", 224))
	fmt.Printf("sha3-224 编码: %v\n", crypto.EncryptSha3ToBase64("hello world", 224))
	fmt.Printf("sha3-256 编码: %v\n", crypto.EncryptSha3ToHex("hello world", 256))
	fmt.Printf("sha3-256 编码: %v\n", crypto.EncryptSha3ToBase64("hello world", 256))
	fmt.Printf("sha3-384 编码: %v\n", crypto.EncryptSha3ToHex("hello world", 384))
	fmt.Printf("sha3-384 编码: %v\n", crypto.EncryptSha3ToBase64("hello world", 384))
	fmt.Printf("sha3-512 编码: %v\n", crypto.EncryptSha3ToHex("hello world", 512))
	fmt.Printf("sha3-512 编码: %v\n", crypto.EncryptSha3ToBase64("hello world", 512))
	fmt.Printf("sha256 编码: %v\n", crypto.EncryptSha256ToHex("hello world"))
	fmt.Printf("sha256 编码: %v\n", crypto.EncryptSha256ToBase64("hello world"))

	publicKeyPkcs1, privateKeyPkcs1 := crypto.GenKeyPkcs1Pair()
	fmt.Printf("生成 PKCS1 格式的 RSA 密钥对: publicKeyPkcs1=%s privateKeyPkcs1=%s\n", publicKeyPkcs1, privateKeyPkcs1)
	publicKeyPkcs8, privateKeyPkcs8 := crypto.GenKeyPkcs8Pair()
	fmt.Printf("生成 PKCS8 格式的 RSA 密钥对: publicKeyPkcs8=%s privateKeyPkcs8=%s\n", publicKeyPkcs8, privateKeyPkcs8)

	fmt.Printf("验证 RSA 密钥对是否匹配 :%v\n", crypto.VerifyKeyPair(publicKeyPkcs1, privateKeyPkcs1))
	fmt.Printf("验证 RSA 密钥对是否匹配 :%v\n", crypto.VerifyKeyPair(publicKeyPkcs8, privateKeyPkcs8))
	fmt.Printf("验证是否是 RSA 公钥 :%v\n", crypto.IsPublicKey(publicKeyPkcs1))
	fmt.Printf("验证是否是 RSA 私钥 :%v\n", crypto.IsPrivateKey(privateKeyPkcs8))

	parsePublicKey, _ := crypto.ParsePublicKey(publicKeyPkcs1)
	fmt.Printf("解析公钥 :%v\n", parsePublicKey)

	parsePrivateKey, _ := crypto.ParsePrivateKey(privateKeyPkcs1)
	fmt.Printf("解析私钥 :%v\n", parsePrivateKey)

	exportPrivateKey, exportPrivateKeyErr := crypto.ExportPublicKey(publicKeyPkcs1)
	if exportPrivateKeyErr != nil {
		fmt.Errorf("exportPrivateKeyErr:%s\n", exportPrivateKeyErr.Error())
	}
	fmt.Printf("从 RSA 私钥里导出公钥 :%v\n", exportPrivateKey)

}
