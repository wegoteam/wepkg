package copy

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var types = NewTypeSet(
	reflect.TypeOf(Struct{}),
	reflect.TypeOf(StructOnCopyed{}),
	reflect.TypeOf(StructOnCopyed2{}),
	reflect.TypeOf(StructAny{}),
)

var baseTypes = NewTypeSet(
	reflect.TypeOf(time.Time{}),
)

type TestEnum int8

const (
	TestEnum_A1 TestEnum = iota + 1
	TestEnum_A2
)

func (s TestEnum) To() (r string) {
	switch s {
	case TestEnum_A1:
		return "A1"
	case TestEnum_A2:
		return "A2"
	}
	return
}

func (s TestEnum) From(val string) (e TestEnum, err error) {
	switch val {
	case "A1":
		return TestEnum_A1, nil
	case "A2":
		return TestEnum_A2, nil
	}
	return
}

type ClassAny struct {
	Any interface{}
}

type ClassStruct struct {
	Struct *Struct
}

type ClassBase struct {
	Int    int64
	Uint   uint
	String string
	Float  float64
	Bool   bool
	Bytes  []byte

	ConvertString ConvertString
	//

	local int
}

type ClassIgnore struct {
	ClassBase `stcopy:"ignore"`

	String string
}

type ClassCombination struct {
	ClassBase
	ClassStruct
}

type ClassArray struct {
	Array       []string
	ArrayStruct []*Struct
}

type ClassConvert struct {
	Convert       *Convert
	ConvertDefine ConvertDefine
	Enum          TestEnum
}

type ClassConvert2 struct {
	Convert string
	Enum    string
}

type ClassMap struct {
	Map       map[string]string
	MapStruct map[string]*Struct
}

type ClassEnum struct {
	Enum    TestEnum
	EnumMap map[TestEnum]bool
}

type Struct struct {
	String string `value:"default"`
	Int    int    `value:"1"`
	Bool   bool
	Float  float64
	Uint   uint
}

type StructOnCopyed struct {
	Int int
}

func (s *StructOnCopyed) CopyAfter() {
	s.Int = s.Int * 2
}

func (s *StructOnCopyed) CopyBefore() {
	s.Int = s.Int / 2
}

type StructOnCopyed2 struct {
	Struct *StructOnCopyed
}

func (s *StructOnCopyed2) CopyBefore() {
	s.Struct.Int = s.Struct.Int / 2
}

func (s *StructOnCopyed2) CopyAfter() {
	s.Struct.Int = s.Struct.Int * 2
}

type StructAny struct {
	Any interface{}
}

type Convert struct {
	Int int16
}

func (s *Convert) To(ctx *CopyContext) (r string) {
	b, _ := json.Marshal(s)
	r = string(b)
	return
}

func (s *Convert) From(ctx *CopyContext, val interface{}) (err error) {
	err = json.Unmarshal([]byte(val.(string)), &s)
	if err != nil {
		return
	}
	return
}

type ConvertMap struct {
	Int int16
}

func (s *ConvertMap) ToMap(ctx *CopyContext) (r map[string]interface{}) {
	return
}

func (s *ConvertMap) FromMap(ctx *CopyContext, val map[string]interface{}) (err error) {
	return
}

type ConvertString string

type ConvertDefine int

func (s ConvertDefine) To() (r string) {
	return strconv.Itoa(int(s))
}

func (s ConvertDefine) From(val string) (r ConvertDefine, err error) {
	t, err := strconv.Atoi(val)
	if err != nil {
		return
	}
	r = ConvertDefine(t)
	return
}

type A struct {
	Int    int
	String string
}

type B struct {
	Int    int
	String string
}

func TestBeanCopy(t *testing.T) {
	a1 := &A{Int: 100}
	a2 := &A{}
	New(a1).To(a2)
	fmt.Println(a1, a2)
}

func TestCopyMapToStructArray(t *testing.T) {
	sources := []interface{}{
		// next
		&map[string]interface{}{"Array": map[string]interface{}{"0": "A1", "2": "A3"}}, // source
		&ClassArray{}, // result
		&ClassArray{Array: []string{"A1", "", "A3"}}, // result
		// next
		&map[string]interface{}{"ArrayStruct": map[string]interface{}{
			"0": map[string]interface{}{"String": "test struct", "Bool": false, "Int": int(100)},
		}}, // source
		&ClassArray{ArrayStruct: []*Struct{
			{String: "test struct", Bool: true, Int: 200.0},
			{String: "test struct", Bool: true, Int: 200.0},
		}}, // result
		&ClassArray{ArrayStruct: []*Struct{
			{String: "test struct", Bool: false, Int: 100.0},
			{String: "test struct", Bool: true, Int: 200.0},
		}}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		err := New(sources[i+1]).From(sources[i])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}
		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
			break
		}
	}
}

func TestCopyMapToStructEnum(t *testing.T) {
	sources := []interface{}{
		// next
		&map[string]interface{}{"Enum": "A1"}, // source
		&ClassEnum{},                          // target
		&ClassEnum{Enum: TestEnum_A1},         // result
		// next
		&map[string]interface{}{"EnumMap": map[string]interface{}{"A2": true}}, // source
		&ClassEnum{}, // target
		&ClassEnum{EnumMap: map[TestEnum]bool{TestEnum_A2: true}}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		err := New(sources[i+1]).From(sources[i])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}
		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
		}
	}
}

func TestCopyMapToStructCombination(t *testing.T) {

	sources := []interface{}{
		// next
		&map[string]interface{}{"String": "test 1", "Int": 1, "Struct": &map[string]interface{}{"String": "test struct", "Bool": false, "Int": int(100)}}, // source
		&ClassCombination{}, // target
		&ClassCombination{ClassBase: ClassBase{String: "test 1", Int: 1}, ClassStruct: ClassStruct{Struct: &Struct{String: "test struct", Int: 100}}}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		origin := New(sources[i])
		err := origin.To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}
		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
		}
	}
}

func TestCopyStructToStructCombination(t *testing.T) {

	sources := []interface{}{
		&ClassCombination{ClassBase: ClassBase{String: "test 1", Int: 1}}, // source
		&ClassCombination{}, // target
		&ClassCombination{ClassBase: ClassBase{String: "test 1", Int: 1}}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		origin := New(sources[i]).WithBaseTypes(baseTypes)
		err := origin.To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}

		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
		}
	}
}

func TestCopyStructToStruct(t *testing.T) {

	sources := []interface{}{
		&ClassBase{String: "test 1", Int: 1}, // source
		&ClassBase{String: "test 2", Int: 2}, // target
		&ClassBase{String: "test 1", Int: 1}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		origin := New(sources[i]).WithBaseTypes(baseTypes)
		err := origin.To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}

		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
		}
	}
}

func TestCopyStructToStructConvert(t *testing.T) {

	sources := []interface{}{
		&ClassConvert{Convert: &Convert{Int: 100}, Enum: TestEnum_A1}, // source
		&ClassConvert2{Convert: `{"Int":99}`, Enum: "A2"},             // target
		&ClassConvert2{Convert: `{"Int":100}`, Enum: "A1"},            // result
	}

	for i := 0; i < len(sources); i += 3 {
		origin := New(sources[i])
		err := origin.To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}

		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal: ", i)
			break
		}
	}
}

func TestCopyMapToStructBase(t *testing.T) {

	sources := []interface{}{
		//// next
		//&map[string]interface{}{"Uint": 0}, // source
		//&ClassBase{Uint: 2},                // target
		//&ClassBase{Uint: 0},                // result
		//// next
		//&map[string]interface{}{"Float": "3.1"}, // source
		//&ClassBase{Float: 2},                    // target
		//&ClassBase{Float: 3.1},                  // result
		////next
		//&map[string]interface{}{"String": "test 1", "Int": 1}, // source
		//&ClassBase{String: "test 2", Int: 2},                  // target
		//&ClassBase{String: "test 1", Int: 1},                  // result
		//
		//&map[string]interface{}{"Bytes": "dGVzdA=="}, // source
		//&ClassBase{},                      // target
		//&ClassBase{Bytes: []byte("test")}, // result
		//
		//&map[string]interface{}{"Bytes": []byte("test")}, // source
		//&ClassBase{},                      // target
		//&ClassBase{Bytes: []byte("test")}, // result
		////next
		//&map[string]interface{}{"ConvertString": "test string"}, // source
		//&ClassBase{},                             // target
		//&ClassBase{ConvertString: "test string"}, // result
		//
		&map[string]interface{}{"Int": "2432359120053993472"}, // source
		&ClassBase{Int: 2},                   // target
		&ClassBase{Int: 2432359120053993472}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		origin := New(sources[i]).WithBaseTypes(baseTypes)
		err := origin.To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}

		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
		}
	}
}

func TestCopyStructFromMapConvert(t *testing.T) {
	sources := []interface{}{
		&map[string]interface{}{"Convert": `{"Int":100}`},              // source
		&ClassConvert{Convert: &Convert{Int: 99}, Enum: TestEnum(99)},  // target
		&ClassConvert{Convert: &Convert{Int: 100}, Enum: TestEnum(99)}, // result
		// next
		&map[string]interface{}{"Enum": "A2"}, // source
		&ClassConvert{Enum: TestEnum_A1},      // target
		&ClassConvert{Enum: TestEnum_A2},      // result
		//// next
		&map[string]interface{}{"ConvertDefine": "100"},                      // source
		&ClassConvert{ConvertDefine: ConvertDefine(99), Enum: TestEnum(99)},  // target
		&ClassConvert{ConvertDefine: ConvertDefine(100), Enum: TestEnum(99)}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		err := New(sources[i+1]).From(sources[i])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}

		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal: ", i)
			break
		}
	}
}

func TestCopyStructToMapConvert(t *testing.T) {

	sources := []interface{}{
		//
		&ClassConvert{Convert: &Convert{Int: 100}, ConvertDefine: ConvertDefine(200), Enum: TestEnum_A1}, // source
		&map[string]interface{}{}, // target
		&map[string]interface{}{"Convert": "{\"Int\":100}", "Enum": "A1", "ConvertDefine": "200"}, // result
		//
		&ClassBase{Bytes: []byte("test")},           // source
		&map[string]interface{}{"String": "test 1"}, // target
		&map[string]interface{}{"Bytes": "dGVzdA==", "String": "", "Float": 0.0, "Uint": 0.0, "Int": 0.0, "Bool": false, "ConvertString": ""}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		err := New(sources[i]).WithBaseTypes(baseTypes).To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
		}
		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal: ", i)
			break
		}
	}
}

func TestCopyStructToMapBase(t *testing.T) {

	sources := []interface{}{
		&ClassStruct{&Struct{String: "test struct", Int: 100}}, // source
		&map[string]interface{}{},                              // target
		&map[string]interface{}{"Struct": map[string]interface{}{"String": "test struct", "Bool": false, "Float": 0.0, "Uint": 0.0, "Int": 100.0, "_ptr": true}}, // result
		// next 结构没有赋值的字段, 也会覆盖掉目标字段
		&ClassBase{Bytes: []byte("test"), Int: 1, ConvertString: ConvertString("convert string")},                                                           // source
		&map[string]interface{}{"String": "test 2", "Int": 2.0, "Bool": true},                                                                               // target
		&map[string]interface{}{"Bytes": "dGVzdA==", "String": "", "Int": 1.0, "Float": 0.0, "Uint": 0.0, "Bool": false, "ConvertString": "convert string"}, // result
		//
		&ClassBase{Int: 2432359120053993472}, // source
		&map[string]interface{}{},            // target
		&map[string]interface{}{"String": "", "Int": "2432359120053993472", "Float": 0.0, "Uint": 0.0, "Bool": false, "ConvertString": ""}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		err := New(sources[i]).WithBaseTypes(baseTypes).To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
			return
		}

		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])

		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
			break
		}
	}
}

func TestCopyStructToMapIgnore(t *testing.T) {

	sources := []interface{}{
		&ClassIgnore{ClassBase: ClassBase{Bytes: []byte("test"), Int: 1, ConvertString: ConvertString("convert string")}, String: "test ignore"}, // source
		&map[string]interface{}{},                        // target
		&map[string]interface{}{"String": "test ignore"}, // result
		//
		&Struct{Uint: 1, Float: 0.0, String: "default", Int: 1, Bool: true},
		&map[string]interface{}{},                          // target
		&map[string]interface{}{"Bool": true, "Uint": 1.0}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		err := New(sources[i]).WithBaseTypes(baseTypes).WithIgnoreDefault().To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
			return
		}

		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])

		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
			break
		}
	}
}

func TestCopyStructToMapAlwaysStructInfo(t *testing.T) {

	sources := []interface{}{
		&ClassStruct{&Struct{String: "test struct", Int: 100}}, // source
		&map[string]interface{}{},                              // target
		&map[string]interface{}{"_type": "ClassStruct", "Struct": map[string]interface{}{"String": "test struct", "Bool": false, "Float": 0.0, "Uint": 0.0, "Int": 100.0, "_ptr": true, "_type": "Struct"}}, // result
		// next 结构没有赋值的字段, 也会覆盖掉目标字段
		&ClassBase{Bytes: []byte("test"), Int: 1, ConvertString: ConvertString("convert string")},                                                                                 // source
		&map[string]interface{}{"String": "test 2", "Int": 2.0, "Bool": true},                                                                                                     // target
		&map[string]interface{}{"_type": "ClassBase", "Bytes": "dGVzdA==", "String": "", "Float": 0.0, "Uint": 0.0, "Int": 1.0, "Bool": false, "ConvertString": "convert string"}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		err := New(sources[i]).WithBaseTypes(baseTypes).WithConfig(&Config{
			AlwaysStructInfo: true,
		}).To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
			return
		}

		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])

		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
			break
		}
	}
}

// 根据provideType转化map
//func TestCopyMapToMapBase(t *testing.T) {
//
//	sources := []interface{}{
//		&map[string]interface{}{"String": "test 1", "Int": 1, "Bool": false}, // source
//		&map[string]interface{}{}, // target
//		&map[string]interface{}{"String": "test 1", "Int": int(1), "Bool": false}, // result
//		//next struct
//		//&map[string]interface{}{"Struct": &map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": int(1)}}, // source
//		//&map[string]interface{}{"Struct": &map[string]interface{}{"String": "test struct", "Bool": false}},                   // target
//		//&map[string]interface{}{"Struct": &map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": int(1)}}, // result
//	}
//
//	for i := 0; i < len(sources); i += 3 {
//		err := New(sources[i]).WithProvideTyp(reflect.TypeOf(&ClassBase{})).To(sources[i+1])
//		if err != nil {
//			t.Error("stcopy: " + err.Error())
//		}
//		debugutil.PrintJson("result=", sources[i+2])
//		debugutil.PrintJson("target=", sources[i+1])
//		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
//			t.Error("not equal")
//		}
//	}
//}

// 根据provideType转化map
//func TestCopyMapToMapArray(t *testing.T) {
//
//	sources := []interface{}{
//		//next
//		&map[string]interface{}{"Array": []string{"c", "d"}},      // source
//		&map[string]interface{}{},                                 // target
//		&map[string]interface{}{"Array": []interface{}{"c", "d"}}, // result
//		////next
//		&map[string]interface{}{"Array": []string{"c", "d"}},      // source
//		&map[string]interface{}{"Array": []interface{}{"a"}},      // target
//		&map[string]interface{}{"Array": []interface{}{"c", "d"}}, // result
//		//next array struct
//		&map[string]interface{}{"ArrayStruct": []interface{}{&map[string]interface{}{"String": "1"}}},           // source
//		&map[string]interface{}{"ArrayStruct": []interface{}{&map[string]interface{}{"String": "1", "Int": 1}}}, // target
//		&map[string]interface{}{"ArrayStruct": []interface{}{&map[string]interface{}{"String": "1"}}},           // result
//		//next array struct
//		&map[string]interface{}{"ArrayStruct": []interface{}{&map[string]interface{}{"String": "1"}, &map[string]interface{}{"String": "2"}}}, // source
//		&map[string]interface{}{"ArrayStruct": []interface{}{&map[string]interface{}{"String": "1", "Int": 1}}},                               // target
//		&map[string]interface{}{"ArrayStruct": []interface{}{&map[string]interface{}{"String": "1"}, &map[string]interface{}{"String": "2"}}}, // result
//	}
//
//	for i := 0; i < len(sources); i += 3 {
//		origin, err := NewContext(sources[i])
//		if err != nil {
//			panic(err)
//		}
//		origin.WithProvideTyp(reflect.TypeOf(&ClassArray{}))
//		err = origin.To(sources[i+1])
//		if err != nil {
//			panic(err)
//		}
//		debugutil.PrintJson("result=", sources[i+2])
//		debugutil.PrintJson("target=", sources[i+1])
//		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
//			t.Error("not equal")
//		}
//	}
//}

// 根据provideType转化map
//func TestCopyMapToMapMap(t *testing.T) {
//
//	sources := []interface{}{
//		// next map
//		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1"}}, // source
//		&map[string]interface{}{"Map": map[string]interface{}{"a": "a2"}}, // target
//		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1"}}, // result
//		//// next map
//		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1", "c": "c3"}},            // source
//		&map[string]interface{}{"Map": map[string]interface{}{"a": "a2", "b": "b2"}},            // target
//		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1", "b": "b2", "c": "c3"}}, // result
//		//// next map struct
//		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": &map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": int(1)}}}, // source
//		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": &map[string]interface{}{"String": "test struct", "Bool": false}}},                 // target
//		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": &map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": int(1)}}}, // result
//		// next map struct
//		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": &map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": int(1)}, "c": &map[string]interface{}{"String": "test struct 1", "Int": int(1)}}},                                                                       // source
//		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": &map[string]interface{}{"String": "test struct", "Bool": false}, "b": &map[string]interface{}{"String": "test struct", "Bool": false}}},                                                                                         // target
//		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": &map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": int(1)}, "b": &map[string]interface{}{"String": "test struct", "Bool": false}, "c": &map[string]interface{}{"String": "test struct 1", "Int": int(1)}}}, // result
//	}
//
//	for i := 0; i < len(sources); i += 3 {
//		origin, err := NewContext(sources[i])
//		if err != nil {
//			panic(err)
//		}
//		origin.WithProvideTyp(reflect.TypeOf(&ClassMap{}))
//		err = origin.To(sources[i+1])
//		if err != nil {
//			panic(err)
//		}
//		debugutil.PrintJson("result=", sources[i+2])
//		debugutil.PrintJson("target=", sources[i+1])
//		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
//			t.Error("not equal")
//		}
//	}
//}

func TestCopyAnyToJsonMap(t *testing.T) {
	sources := []interface{}{
		//"test string type",
		//"test string type",
		//int(10),
		//float64(10),
		//int32(10),
		//float64(10),
		//uint(10),
		//float64(10),
		//// 5
		//uint32(10),
		//float64(10),
		//true,
		//true,
		//map[string]interface{}{"String": "test struct"},
		//map[string]interface{}{"String": "test struct"},
		//&map[string]interface{}{"String": "test struct"},
		//map[string]interface{}{"String": "test struct", "_ptr": true},
		//map[string]interface{}{
		//	"a": "test a",
		//	"b": "test b",
		//},
		//map[string]interface{}{
		//	"a": "test a",
		//	"b": "test b",
		//},
		// 10
		//map[string]interface{}{
		//	"a": &map[string]interface{}{"String": "test map struct a"},
		//	"b": &map[string]interface{}{"String": "test map struct b"},
		//},
		//map[string]interface{}{
		//	"a": map[string]interface{}{"String": "test map struct a", "_ptr": true},
		//	"b": map[string]interface{}{"String": "test map struct b", "_ptr": true},
		//},
		//Struct{String: "test struct", Int: 100, Bool: true},
		//map[string]interface{}{"String": "test struct", "Float": 0.0, "Uint": 0.0, "Int": 100.0, "Bool": true, "_type": "Struct"},
		//&Struct{String: "test struct", Int: 100, Bool: true},
		//map[string]interface{}{"String": "test struct", "Float": 0.0, "Uint": 0.0, "Int": 100.0, "Bool": true, "_type": "Struct", "_ptr": true},
		//[]string{"1", "2"},
		//[]interface{}{"1", "2"},
		//[]byte{0, 1, 2},
		//"AAEC",
		//[]int{1, 2},
		//[]interface{}{1.0, 2.0},
		//[]map[string]interface{}{
		//	{"String": "test struct"},
		//	{"String": "test struct"},
		//},
		//[]interface{}{
		//	map[string]interface{}{"String": "test struct"},
		//	map[string]interface{}{"String": "test struct"},
		//},
		//[]*map[string]interface{}{
		//	{"String": "test struct"},
		//	{"String": "test struct"},
		//},
		//[]interface{}{
		//	map[string]interface{}{"String": "test struct", "_ptr": true},
		//	map[string]interface{}{"String": "test struct", "_ptr": true},
		//},
		//[]Struct{
		//	{String: "test struct", Int: 100, Bool: true},
		//	{String: "test struct 2", Int: 200, Bool: true},
		//},
		//[]interface{}{
		//	map[string]interface{}{"String": "test struct", "Float": 0.0, "Uint": 0.0, "Int": 100.0, "Bool": true, "_type": "Struct"},
		//	map[string]interface{}{"String": "test struct 2", "Float": 0.0, "Uint": 0.0, "Int": 200.0, "Bool": true, "_type": "Struct"},
		//},
		//[]*Struct{
		//	{String: "test struct", Int: 100, Bool: true},
		//	{String: "test struct 2", Int: 200, Bool: true},
		//},
		//[]interface{}{
		//	map[string]interface{}{"String": "test struct", "Float": 0.0, "Uint": 0.0, "Int": 100.0, "Bool": true, "_type": "Struct", "_ptr": true},
		//	map[string]interface{}{"String": "test struct 2", "Float": 0.0, "Uint": 0.0, "Int": 200.0, "Bool": true, "_type": "Struct", "_ptr": true},
		//},
		//测试 CopyAfter
		//&StructOnCopyed{Int: 20},
		//map[string]interface{}{"_ptr": true, "_type": "StructOnCopyed", "Int": 10.0},
		// 测试CopyAfter 多层
		&StructOnCopyed2{Struct: &StructOnCopyed{Int: 40}},
		map[string]interface{}{"_ptr": true, "_type": "StructOnCopyed2", "Struct": map[string]interface{}{"Int": 10.0, "_ptr": true, "_type": "StructOnCopyed"}},
	}

	// map转换成json map
	for i := 0; i < len(sources); i += 2 {
		source := &map[string]interface{}{
			"Any": sources[i],
		}
		target := &map[string]interface{}{}
		origin := New(source)
		origin.WithProvideTyp(reflect.TypeOf(&ClassAny{}))
		err := origin.To(target)
		if err != nil {
			panic(err)
		}

		resultMap := sources[i+1]
		targetAny := (*target)["Any"]
		fmt.Printf("result=%v\n", resultMap)
		fmt.Printf("target=%v\n", targetAny)
		if reflect.DeepEqual(resultMap, targetAny) == false {
			t.Error("not equal: " + strconv.Itoa(i))
			return
		}
	}

}

func TestCopyJsonMapToStruct(t *testing.T) {
	sources := []interface{}{
		"test string type",
		"test string type",
		true,
		true,
		map[string]interface{}{
			"String": "test struct",
		},
		map[string]interface{}{
			"String": "test struct",
		},
		map[string]interface{}{
			"a": "test a",
			"b": "test b",
		},
		map[string]interface{}{
			"a": "test a",
			"b": "test b",
		},
		map[string]interface{}{
			"a": &map[string]interface{}{"String": "test map struct a"},
			"b": &map[string]interface{}{"String": "test map struct b"},
		},
		map[string]interface{}{
			"a": map[string]interface{}{"String": "test map struct a"},
			"b": map[string]interface{}{"String": "test map struct b"},
		},
		map[string]interface{}{"String": "test struct", "Int": float64(100), "Bool": true, "_type": "Struct"},
		Struct{String: "test struct", Int: 100, Bool: true},
		map[string]interface{}{"String": "test struct", "Int": float64(100), "Bool": true, "_type": "Struct", "_ptr": true},
		&Struct{String: "test struct", Int: 100, Bool: true},
		// next
		map[string]interface{}{"Any": map[string]interface{}{"String": "test struct", "Int": float64(100), "Bool": true, "_type": "Struct", "_ptr": true}, "_type": "StructAny", "_ptr": true},
		&StructAny{Any: &Struct{String: "test struct", Int: 100, Bool: true}},
		// next
		map[string]interface{}{"Any": []interface{}{
			map[string]interface{}{"String": "test struct", "Int": float64(100), "Bool": true, "_type": "Struct", "_ptr": true},
		}, "_type": "StructAny", "_ptr": true},
		&StructAny{Any: []interface{}{
			&Struct{String: "test struct", Int: 100, Bool: true},
		}},
		// next
		map[string]interface{}{"String": "test struct", "Int": float64(100), "Bool": true, "_type": "NoStruct", "_ptr": true},
		map[string]interface{}{"String": "test struct", "Int": float64(100), "Bool": true, "_type": "NoStruct", "_ptr": true},
		[]interface{}{"1", "2"},
		[]interface{}{"1", "2"},
		[]interface{}{
			map[string]interface{}{"String": "test struct"},
			map[string]interface{}{"String": "test struct"},
		},
		[]interface{}{
			map[string]interface{}{"String": "test struct"},
			map[string]interface{}{"String": "test struct"},
		},
		[]interface{}{
			&map[string]interface{}{"String": "test struct"},
			&map[string]interface{}{"String": "test struct"},
		},
		[]interface{}{
			map[string]interface{}{"String": "test struct"},
			map[string]interface{}{"String": "test struct"},
		},
		[]interface{}{
			map[string]interface{}{"String": "test struct", "Int": 100.0, "Bool": true, "_type": "Struct"},
			map[string]interface{}{"String": "test struct 2", "Int": 200.0, "Bool": true, "_type": "Struct"},
		},
		[]interface{}{
			Struct{String: "test struct", Int: 100, Bool: true},
			Struct{String: "test struct 2", Int: 200, Bool: true},
		},
		[]interface{}{
			map[string]interface{}{"String": "test struct", "Int": 100.0, "Bool": true, "_type": "Struct", "_ptr": true},
			map[string]interface{}{"String": "test struct 2", "Int": 200.0, "Bool": true, "_type": "Struct", "_ptr": true},
		},
		[]interface{}{
			&Struct{String: "test struct", Int: 100, Bool: true},
			&Struct{String: "test struct 2", Int: 200, Bool: true},
		},
		[]interface{}{
			map[string]interface{}{"String": "test struct", "Int": 100.0, "Bool": true, "_type": "NoStruct", "_ptr": true},
			map[string]interface{}{"String": "test struct 2", "Int": 200.0, "Bool": true, "_type": "NoStruct", "_ptr": true},
		},
		[]interface{}{
			map[string]interface{}{"String": "test struct", "Int": 100.0, "Bool": true, "_type": "NoStruct", "_ptr": true},
			map[string]interface{}{"String": "test struct 2", "Int": 200.0, "Bool": true, "_type": "NoStruct", "_ptr": true},
		},
		//测试 CopyAfter
		map[string]interface{}{"_ptr": true, "_type": "StructOnCopyed", "Int": 10},
		&StructOnCopyed{Int: 20},
		// 测试CopyAfter 多层
		map[string]interface{}{"_ptr": true, "_type": "StructOnCopyed2", "Struct": map[string]interface{}{"Int": 10}},
		&StructOnCopyed2{Struct: &StructOnCopyed{Int: 40}},
	}

	// json map转换成struct
	for i := 0; i < len(sources); i += 2 {
		source := &map[string]interface{}{
			"Any": sources[i],
		}
		target := &ClassAny{}
		origin := New(target).WithTypeMap(types)
		err := origin.From(source)
		if err != nil {
			panic(err)
		}
		fmt.Printf("result=%v\n", sources[i+1])
		fmt.Printf("target=%v\n", target.Any)
		//fmt.Printf("copy_test[862]> %T, %T\n", sources[i+1], target.Any)
		if reflect.DeepEqual(sources[i+1], target.Any) == false {
			t.Error("not equal: " + strconv.Itoa(i))
			break
		}
	}
}

func TestCopyMapToMap(t *testing.T) {
	sources := []interface{}{
		// next map
		&map[string]interface{}{"Map": nil},                               // source
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a2"}}, // target
		&map[string]interface{}{"Map": nil},                               // result
		// next map
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1"}}, // source
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a2"}}, // target
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1"}}, // result
		// next map
		&map[string]interface{}{"Map": &map[string]interface{}{"a": "a1"}},              // source
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a2", "_ptr": true}}, // target
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1", "_ptr": true}}, // result
		//// next map
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1", "c": "c3"}},            // source
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a2", "b": "b2"}},            // target
		&map[string]interface{}{"Map": map[string]interface{}{"a": "a1", "b": "b2", "c": "c3"}}, // result
		//// next map struct
		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": &map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": int(1)}}},           // source
		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": map[string]interface{}{"String": "test struct", "Bool": false, "_ptr": true}}},              // target
		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": 1.0, "_ptr": true}}}, // result
		// next map struct
		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": &map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": int(1)}, "c": &map[string]interface{}{"String": "test struct 1", "Int": int(1)}}},                                                                                          // source
		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": map[string]interface{}{"String": "test struct", "Bool": false}, "b": map[string]interface{}{"String": "test struct", "Bool": false}}},                                                                                                              // target
		&map[string]interface{}{"MapStruct": map[string]interface{}{"a": map[string]interface{}{"String": "test struct 1", "Bool": true, "Int": 1.0, "_ptr": true}, "b": map[string]interface{}{"String": "test struct", "Bool": false}, "c": map[string]interface{}{"String": "test struct 1", "Int": 1.0, "_ptr": true}}}, // result
	}

	for i := 0; i < len(sources); i += 3 {
		err := New(sources[i]).To(sources[i+1])
		if err != nil {
			t.Error("stcopy: " + err.Error())
			return
		}
		fmt.Printf("result=%v\n", sources[i+2])
		fmt.Printf("target=%v\n", sources[i+1])
		if reflect.DeepEqual(sources[i+2], sources[i+1]) == false {
			t.Error("not equal")
		}
	}
}
