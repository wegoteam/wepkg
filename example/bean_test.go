package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/bean"
	"testing"
)

type A struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type B struct {
	Name string
}

func TestBeanCopy(t *testing.T) {
	a1 := A{Age: 100}
	a2 := A{Name: "C"}
	//拷贝属性：a2拷贝a1的属性
	errors := bean.BeanCopy(&a1, a2)
	a2.Name = "D"
	a1.Name = "E"
	fmt.Println("Errors:", errors)
	fmt.Println("a1的地址:", &a1)
	fmt.Println("a1:", a1)
	fmt.Println("a2的地址:", &a2)
	fmt.Println("a2:", a2)
}

func TestBeanToMap(t *testing.T) {
	a := A{Age: 100, Name: "C"}
	toMap, err := bean.BeanToMap(a)
	fmt.Println("Errors:", err)
	fmt.Println("toMap:", toMap)
}

func TestBeanClone(t *testing.T) {
	a := &A{Age: 100, Name: "C"}
	clone, err := bean.BeanClone(a)
	clone.(*A).Name = "D"
	fmt.Println("Errors:", err)
	fmt.Println("a:", a)
	fmt.Println("clone:", clone)
	fmt.Println("a的地址:", &a)
	fmt.Println("clone的地址:", &clone)
}

func TestIsZero(t *testing.T) {
	a := A{}
	zero := bean.IsZero(a)
	fmt.Println("zero:", zero)
}

func TestHasZero(t *testing.T) {
	existZero := bean.HasZero(A{Name: "1"})
	existZero2 := bean.HasZero(A{Name: "test", Age: 1})
	fmt.Println("existZero:", existZero)
	fmt.Println("existZero2:", existZero2)
}

func TestHasFieldsZero(t *testing.T) {
	field, existFieldsZero := bean.HasFieldsZero(A{}, "Name", "Age")
	fmt.Println("existFieldsZero:", existFieldsZero)
	fmt.Println("field:", field)
}

func TestGetFields(t *testing.T) {
	field, err := bean.GetFields(A{})
	fmt.Println("err:", err)
	fmt.Println("field:", field)
}

func TestGetKind(t *testing.T) {
	kind, err := bean.GetKind(A{}, "Name")
	fmt.Println("err:", err)
	fmt.Println("kind:", kind)
}

func TestGetTag(t *testing.T) {
	tag, err := bean.GetTag(A{}, "Name")
	fmt.Println("err:", err)
	fmt.Println("tag:", tag.Get("json"))
}

func TestGetTags(t *testing.T) {
	tagsMap, err := bean.GetTags(A{})
	fmt.Println("err:", err)
	fmt.Println("tagsMap:", tagsMap)
}

func TestGetFiled(t *testing.T) {
	field, err := bean.GetFieldVal(A{Name: "test"}, "Name")
	fmt.Println("err:", err)
	fmt.Println("field:", field)
}

func TestSetFieldVal(t *testing.T) {
	var a = A{Name: "test"}
	fmt.Println("a:", a)
	err := bean.SetFieldVal(&a, "Name", "test2")
	fmt.Println("err:", err)
	fmt.Println("修改后a:", a)
}
