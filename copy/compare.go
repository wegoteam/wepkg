package copy

import (
	"errors"
	"reflect"
	"strconv"
)

func (ctx *Context) addCompareError(err error) {
	ctx.compareErrors = append(ctx.compareErrors, err)
}

// 比较目标中[val]包含的表达式
func (ctx *Context) Compare(val interface{}) []error {
	ctx.compareErrors = ctx.compareErrors[:0]
	ctx.compare(ctx.valueA, Value(reflect.ValueOf(val)), "ROOT", 0)
	return ctx.compareErrors
}

// 深度比较, 与目标的类型和值必须完全一致
func (ctx *Context) CompareDeep(val interface{}) []error {
	ctx.compareErrors = ctx.compareErrors[:0]
	ctx.compareType = true
	ctx.compareAll = true
	ctx.compare(ctx.valueA, Value(reflect.ValueOf(val)), "ROOT", 0)
	return ctx.compareErrors
}

func (ctx *Context) compare(source, target Value, path string, depth int) {
	srcref := source.Upper()
	tarref := target.Upper()

	//prefix := strings.Repeat("----", depth)
	//fmt.Println(prefix+"> compare: srctyp=", srcref.Type(), "src=", srcref)
	//fmt.Println(prefix+"compare: tartyp=", target.GetTypeString(), "tar=", tarref, "nil=", ",  canset=", tarref.CanSet(), func() (x string) {
	//	if isHard(tarref.Kind()) && tarref.IsNil() {
	//		x = "isnil=true"
	//	} else {
	//		x = "isnil=false"
	//	}
	//	return
	//}())

	if srcref.IsValid() == false && tarref.IsValid() == false {
		return
	} else {
		if srcref.IsValid() != tarref.IsValid() {
			ctx.addCompareError(errors.New(path + ": valid not match: " + strconv.FormatBool(srcref.IsValid()) + " !=" + strconv.FormatBool(tarref.IsValid()) + "(s/t)"))
			return
		}
	}

	if ctx.compareType {
		// 检查类型是否匹配
		if srcref.Type() != tarref.Type() {
			ctx.addCompareError(errors.New(path + ": type not match: " + srcref.Type().String() + " !=" + tarref.Type().String() + "(s/t)"))
			return
		}
	}
	if srcref.Kind() == reflect.Interface {
		srcref = srcref.Elem()
	}
	if tarref.Kind() == reflect.Interface {
		tarref = tarref.Elem()
	}

	switch srcref.Type().Kind() {
	case reflect.Slice, reflect.Array:
		if srcref.Len() != tarref.Len() {
			ctx.addCompareError(errors.New(path + ": length not equal: " + strconv.Itoa(srcref.Len()) + " !=" + strconv.Itoa(tarref.Len()) + "(s/t)"))
			return
		}
		for i := 0; i < srcref.Len(); i++ {
			srcitem := srcref.Index(i)
			taritem := tarref.Index(i)
			ctx.compare(Value(srcitem), Value(taritem), path+"/"+strconv.Itoa(i), depth+1)
		}
	case reflect.Interface:
		ctx.compare(Value(srcref.Elem()), Value(tarref.Elem()), path, depth+1)
	case reflect.Ptr:
		ctx.compare(Value(srcref.Elem()), Value(tarref.Elem()), path, depth+1)
	case reflect.Struct:
		if tarref.Kind() != reflect.Struct && tarref.Kind() != reflect.Map {
			ctx.addCompareError(errors.New(path + ": target type must be struct or map: type=" + tarref.Type().String() + ""))
			return
		}

		for _, field := range TypeUtiler.GetFieldRecursion(srcref.Type()) {
			srcfield := getFieldVal(srcref, field)
			if srcref.Kind() == reflect.Map {
				if srcfield.IsValid() == false || srcfield.IsNil() {
					continue
				}
			}
			//

			tarfield := getFieldVal(tarref, field)
			switch tarref.Kind() {
			case reflect.Map:
				if tarfield.IsValid() == false || tarfield.IsNil() {
					continue
				}
			case reflect.Struct:
				_, has := tarref.Type().FieldByName(field.Name)
				if has == false {
					if ctx.compareAll {
						ctx.addCompareError(errors.New(path + ": target not field: " + field.Name))
					}
					continue
				}
				if ctx.Config.IgnoreDefault {
					if TypeUtiler.CompareEqualDefault(tarfield, field) {
						continue
					}
				}
			}

			//fmt.Println(">>> compare struct field: ", field.Name, ", fieldtyp=", field.Type)
			ctx.compare(Value(srcfield), Value(tarfield), path+"/"+field.Name, depth+1)
		}
	case reflect.Map:
		if ctx.compareAll {
			if len(srcref.MapKeys()) != len(tarref.MapKeys()) {
				ctx.addCompareError(errors.New(path + ": keys not equal: " + strconv.Itoa(len(srcref.MapKeys())) + " !=" + strconv.Itoa(len(tarref.MapKeys())) + "(s/t)"))
				return
			}
		}
		for _, k := range srcref.MapKeys() {
			val1 := srcref.MapIndex(k)
			val2 := tarref.MapIndex(k)
			//fmt.Println("||| copy map key: ", k, ", fieldtyp=", val1.Type())
			//fmt.Println("src=", val1, ", typ=", val2)

			ctx.compare(Value(val1), Value(val2), path+"/"+k.String(), depth+1)
		}
	case reflect.Func:
		panic("not suppor")
	default:

		if tarref.Kind() == reflect.Func {
			exprResult := tarref.Call([]reflect.Value{srcref})
			bRef := exprResult[0]
			if bRef.Bool() == false {
				ctx.addCompareError(errors.New(path + ": not equal expr: src=" + Convert2String(srcref).String()))
				return
			}
		} else {
			if reflect.DeepEqual(srcref.Interface(), tarref.Interface()) == false {
				ctx.addCompareError(errors.New(path + ": not equal: " + Convert2String(srcref).String() + " !=" + Convert2String(tarref).String() + "(s/t)"))
				return
			}
		}
	}
	return
}
