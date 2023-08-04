package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/id/rand"
	"testing"
)

func TestRandom(t *testing.T) {
	randomStr := rand.RandomStr(10)
	randomNum := rand.RandomNum(10)
	fmt.Printf("randomStr: %s\n", randomStr)
	fmt.Printf("randomNum: %s\n", randomNum)
}
