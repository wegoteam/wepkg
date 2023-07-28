package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/net/http"
	"testing"
)

func TestDefaultClientPOST(t *testing.T) {
	client := http.BuildDefaultClient()
	var res string
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"roleName":""}`).
		SetResult(res).
		Post("http://localhost:18080/weflow/role/list")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", resp)
	fmt.Println("Response Info:", res)
}

type Response[T any] struct {
	Code int    `json:"code"` // 0:成功，其他：失败
	Msg  string `json:"msg"`  // 错误信息
	Data T      `json:"data"` // 数据
}

type RoleInfoResult struct {
	ID         int64  `json:"id"`         // 唯一id
	RoleID     string `json:"roleID"`     // 角色id
	ParentID   string `json:"parentID"`   // 角色父id
	RoleName   string `json:"roleName"`   // 角色名称
	Status     int32  `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string `json:"remark"`     // 描述
	CreateUser string `json:"createUser"` // 创建人
	UpdateUser string `json:"updateUser"` // 更新人
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}

func TestGet(t *testing.T) {
	res1, err := http.Get[Response[[]RoleInfoResult]]("http://localhost:18080/weflow/role/list",
		map[string]string{
			"roleName": "",
		})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res1)

	res2, err := http.GetString("http://localhost:18080/weflow/role/list",
		map[string]string{
			"roleName": "",
		})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res2)
}

func TestPost(t *testing.T) {
	type Role struct {
		RoleName string `json:"roleName"`
	}
	var param = &Role{}
	res1, err := http.Post[Response[[]RoleInfoResult]]("http://localhost:18080/weflow/role/list", param)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res1)

	res2, err := http.PostString("http://localhost:18080/weflow/role/list", param)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res2)

	res3, err := http.PostForm[Response[[]RoleInfoResult]]("http://localhost:18080/weflow/role/list",
		map[string]string{
			"roleName": "",
		})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res3)

	res4, err := http.PostFile[Response[any]]("http://localhost:18080/weflow/upload/file", "a.txt", "./testdata/a.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res4)

	res5, err := http.PostFiles[Response[any]]("http://localhost:18080/weflow/upload/file", map[string]string{
		"a.txt": "./testdata/a.txt",
	})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res5)
}
