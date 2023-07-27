package uuid

import (
	"github.com/google/uuid"
	"strings"
)

// New
// @Description: 生成uuid
// @return string
func New() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
