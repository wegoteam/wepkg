package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/io/json"
	"testing"
)

func TestFormat(t *testing.T) {
	type Role struct {
		RoleName string `json:"roleName"`
	}
	var param = &Role{
		RoleName: "admin",
	}
	marshal, err := json.Marshal(param)
	if err != nil {
		fmt.Errorf("json.Marshal err: %v", err)
	}
	fmt.Println(marshal)

	var role = &Role{}
	err = json.Unmarshal(marshal, role)
	if err != nil {
		fmt.Errorf("json.Unmarshal err: %v", err)
	}
	fmt.Println(role)
}
