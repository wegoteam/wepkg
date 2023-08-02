package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/io/compress"
	"testing"
)

func TestCompress(t *testing.T) {
	var dst []byte
	var source = []byte("test")
	encode := compress.Encode(dst, source)
	fmt.Printf("encode:%s\n", encode)
	fmt.Printf("dst encode:%s\n", dst)
	var src []byte
	decode, err := compress.Decode(encode, src)
	if err != nil {
		fmt.Errorf("err:%s\n", err.Error())
	}
	fmt.Printf("decode:%s\n", decode)
	fmt.Printf("src decode:%s\n", src)
}
