package copy

import (
	"errors"
	"reflect"
	"strconv"
)

func (ctx *Context) Valid() error {
	return ctx.valid(ctx.valueA, ctx.valueA.Upper().Type(), 0)
}

func (ctx *Context) valid(source Value, provideTyp reflect.Type, depth int) (err error) {
	srcref := source.Upper()
	//fmt.Println("\n||| to", "provide=", provideTyp)
	//fmt.Println("srctyp=", srcref.Type(), "src=", srcref)

	// 源是否空
	if srcref.IsValid() == false {
		return
	}
	if isHard(srcref.Kind()) && srcref.IsNil() {
		return
	}

	// 接口处理
	if provideTyp.Kind() != reflect.Interface {
		if srcref.Kind() == reflect.Interface {
			srcref = srcref.Elem()
		}
	}
	//fmt.Println("last target=", tarref, tarref.Type(), tarref.CanSet())

	switch provideTyp.Kind() {
	case reflect.Slice, reflect.Array:
		if srcref.Len() == 0 {
			return
		}
		for i := 0; i < srcref.Len(); i++ {
			srcitem := srcref.Index(i)
			err = ctx.valid(Value(srcitem), provideTyp.Elem(), depth+1)
			if err != nil {
				err = errors.New("at " + strconv.Itoa(i) + ": " + err.Error())
				return
			}
		}
	case reflect.Interface:
		err = ctx.valid(Value(srcref.Elem()), srcref.Elem().Type(), depth+1)
		if err != nil {
			return
		}
	case reflect.Ptr:
		err = ctx.valid(Value(srcref.Elem()), provideTyp.Elem(), depth+1)
		if err != nil {
			return
		}
	case reflect.Struct:
		for _, field := range TypeUtiler.GetFieldRecursion(provideTyp) {
			srcfield := getFieldVal(srcref, field)
			if srcref.Kind() == reflect.Map {
				if srcfield.IsValid() == false || srcfield.IsNil() {
					continue
				}
			}
			//fmt.Println(">>> copy struct field: ", field.Name, ", fieldtyp=", field.Type)
			err = ctx.valid(Value(srcfield), field.Type, depth+1)
			if err != nil {
				err = errors.New(field.Name + ": " + err.Error())
				return
			}
		}
	case reflect.Map:
		for _, k := range srcref.MapKeys() {

			val1 := srcref.MapIndex(k)
			if val1.IsValid() == false {
				continue
			}
			//fmt.Println("||| copy map key: ", k, ", fieldtyp=", val1.Type())
			//fmt.Println("src=", val1, ", typ=", val2)

			err = ctx.valid(Value(val1), val1.Type(), depth+1)
			if err != nil {
				err = errors.New("at " + k.String() + ": " + err.Error())
				return
			}
		}

	case reflect.Func:
		panic("not suppor")
	default:
	}

	mname := "Valid"
	mtype, ok := srcref.Type().MethodByName(mname)
	if ok == true {
		methodVal := srcref.MethodByName(mname)
		if mtype.Type.NumIn() > 2 {
			err = errors.New("func " + mname + " NumIn() must 1 or 0")
			return
		}
		results := func() (x []reflect.Value) {
			if mtype.Type.NumIn() == 2 {
				x = methodVal.Call([]reflect.Value{reflect.ValueOf(ctx)})

			} else {
				x = methodVal.Call([]reflect.Value{})
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
	//fmt.Println("resut >", result.Upper())
	return
}
