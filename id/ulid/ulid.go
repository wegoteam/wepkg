package ulid

//引用：https://github.com/oklog/ulid

import (
	"fmt"
	ulidUtil "github.com/oklog/ulid"
	"math/rand"
	"time"
)

// New
// @Description: 生成ulid
// @return string ulid
func New() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulidUtil.Timestamp(time.Now())
	id, err := ulidUtil.New(ms, entropy)
	if err != nil {
		fmt.Errorf("failed to generate ulid: %v", err)
		return ""
	}
	return id.String()
}
