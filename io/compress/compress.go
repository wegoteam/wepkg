package compress

import compressUtil "github.com/golang/snappy"

//引用：https://github.com/golang/snappy
//snappy压缩算法具有以下特性：
//快速：压缩速度大概在250MB/秒及更快的速度进行压缩。
//稳定：在过去的几年中，Snappy在Google的生产环境中压缩并解压缩了数P字节（petabytes）的数据。Snappy位流格式是稳定的，不会在版本之间发生变化
//健壮性：Snappy解压缩器设计为不会因遇到损坏或恶意输入而崩溃

// Encode
// @Description: 对字节切片进行压缩
// @param: dst
// @param: src
// @return []byte
func Encode(dst, src []byte) []byte {
	return compressUtil.Encode(dst, src)
}

// Decode
// @Description: 对字节切片进行解压缩
// @param: dst
// @param: src
// @return []byte
// @return error
func Decode(dst, src []byte) ([]byte, error) {
	return compressUtil.Decode(dst, src)
}
