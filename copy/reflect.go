package copy

import (
	"errors"
	"github.com/wegoteam/wepkg/conv/arrayconv"
	"reflect"
	"strconv"
	"strings"
)

var (
	errTyp    = reflect.TypeOf((*error)(nil)).Elem()
	bytesTyp  = reflect.TypeOf([]byte{})
	stringTyp = reflect.TypeOf("")
)

func NewTypeSet(types ...reflect.Type) (r *TypeSet) {
	r = &TypeSet{}
	r.Add(types...)
	return
}

type TypeSet struct {
	arrayconv.ArrayList
}

func (m *TypeSet) Add(items ...reflect.Type) {
	for _, val := range items {
		m.ArrayList.Add(val)
	}
}

func (m *TypeSet) GetByName(n string) (typ reflect.Type, has bool) {
	ityp, has := m.First(func(item interface{}) bool {
		t := item.(reflect.Type)
		return t.Name() == n
	})

	if has {
		typ = ityp.(reflect.Type)
	}

	return
}

type Value reflect.Value

func (val Value) Upper() reflect.Value {
	return reflect.Value(val)
}

func (val Value) Indirect() Value {
	return Value(reflect.Indirect(val.Upper()))
}

func (val Value) unfoldInterface() (r Value) {
	ref := val.Upper()
	if ref.Kind() == reflect.Interface {
		return Value(ref.Elem())
	}
	return val
}

// 如果val是map类型, 检查键值为_type和_ptr的值, 获取反射类型
func (val Value) parseMapType(ctx *Context) (x reflect.Type, err error) {
	srcunfold := val.unfoldInterface()
	if srcunfold.Upper().Kind() != reflect.Map {
		err = errors.New("not map")
		return
	}
	src := val.Upper().Interface().(map[string]interface{})
	// 处理类型
	sttype := func() (y reflect.Type) {
		istr, srcok := src["_type"]
		if srcok == false {
			return
		}

		str := istr.(string)
		t, typok := ctx.typeMap.GetByName(str)
		if typok == false {
			println("stcopy: parse map type: not found: " + str)
			return
		}

		delete(src, "_type")
		y = t
		return
	}()
	// 处理指针

	isPtr := func() (x bool) {
		if sttype == nil {
			return false
		}
		istr, ok := src["_ptr"]
		if ok == false {
			return false
		}
		delete(src, "_ptr")
		x = istr.(bool)
		return
	}()

	//
	if sttype == nil {
		x = srcunfold.Upper().Type()
	} else {
		x = sttype
	}
	if isPtr {
		x = reflect.PtrTo(x)
	}
	return
}

// 为map对象添加结构类型 {"_type":Name}
func (val Value) updateMapStructTypeBy(source Value) (err error) {
	indirect := source.Indirect()
	if indirect.Upper().Kind() != reflect.Struct {
		//
		return
	}

	ref := val.Indirect()
	if ref.Upper().Kind() != reflect.Map {
		err = errors.New("not map")
		return
	}

	ref.Upper().SetMapIndex(reflect.ValueOf("_type"), reflect.ValueOf(indirect.Upper().Type().Name()))
	return
}

// 为map对象添加结构类型 {"_ptr":boolean}
func (val Value) updateMapStructPtrBy(source Value) (err error) {
	indirect := source.Indirect()
	if indirect.Upper().Kind() != reflect.Struct {
		return
	}
	ref := val.Indirect()
	if ref.Upper().Kind() != reflect.Map {
		err = errors.New("not map")
		return
	}

	if source.Upper().Kind() == reflect.Ptr {
		ref.Upper().SetMapIndex(reflect.ValueOf("_ptr"), reflect.ValueOf(true))
	}
	return
}

func (val Value) updateMapPtrBy(source Value) (err error) {
	ref := val.Indirect()
	if ref.Upper().Kind() != reflect.Map {
		err = errors.New("not map")
		return
	}

	if source.Upper().Kind() == reflect.Ptr {
		ref.Upper().SetMapIndex(reflect.ValueOf("_ptr"), reflect.ValueOf(true))
	}
	return
}

func (val Value) GetTypeString() (y string) {
	if val.Upper().IsValid() {
		y = val.Upper().Type().String()
	} else {
		y = "is nil"
	}
	return
}

func (val Value) IsNil() bool {
	if isHard(val.Upper().Kind()) {
		return val.IsNil()
	}
	return false
}

var (
	TypeUtiler = typeUtil(0)
)

type typeUtil int

// 获取正确的反射对象，如果nil，创建新的
func (*typeUtil) UnfoldType(typ reflect.Type) reflect.Type {
	switch typ.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		typ = typ.Elem()
		return typ
	}

	return typ
}

func (*typeUtil) CompareEqualDefault(value reflect.Value, field *reflect.StructField) bool {
	x := Convert2String(value)
	tag, tagok := field.Tag.Lookup("value")
	if tagok {
		return x.String() == tag
	}

	zero := reflect.Zero(field.Type)

	return value.Interface() == zero.Interface()
}

func (*typeUtil) HasTagIgnore(field reflect.StructField) bool {
	flags, has := field.Tag.Lookup("stcopy")
	if has == false {
		return false
	}

	return strings.Contains(flags, "ignore")
}

// 获取
func (sv *typeUtil) GetFieldRecursion(typ reflect.Type) (r []*reflect.StructField) {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		switch field.Type.Kind() {
		case reflect.Struct:
			if TypeUtiler.HasTagIgnore(field) {
				continue
			}
			if field.Anonymous == true {
				r = append(r, sv.GetFieldRecursion(field.Type)...)
			} else {
				r = append(r, &field)
			}

		default:
			r = append(r, &field)
		}
	}
	return
}

func (*typeUtil) Call(target reflect.Value, mname string, ctx *Context, args ...reflect.Value) (err error) {
	mtype, ok := target.Type().MethodByName(mname)
	if ok == true {
		methodVal := target.MethodByName(mname)
		argslen := 2 + len(args)
		if mtype.Type.NumIn() > argslen {
			err = errors.New("func " + mname + ": number of argument is not " + strconv.Itoa(mtype.Type.NumIn()))
			return
		}
		results := func() (x []reflect.Value) {
			if mtype.Type.NumIn() == 1 {
				x = methodVal.Call([]reflect.Value{})
			} else {
				a := append([]reflect.Value{reflect.ValueOf(ctx)}, args...)
				x = methodVal.Call(a)
			}
			return
		}()
		// 包含error
		if mtype.Type.NumOut() > 0 {
			if results[0].Type().Kind() == reflect.Bool {
				if results[0].Bool() == false {
					err = errors.New("failed")
				}
			} else { // error
				if results[0].IsNil() == false {
					err = results[0].Interface().(error)
				}
			}
			return
		}
	}
	return
}
