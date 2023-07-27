package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/id/ulid"
	"github.com/wegoteam/wepkg/id/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	fmt.Printf("uuid: %s\n", uuid.New())
	fmt.Printf("ulid: %s\n", ulid.New())
}
