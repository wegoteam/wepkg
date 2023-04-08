package copy

import (
	"encoding/base64"
	"errors"
	"github.com/wegoteam/wepkg/conv/stringconv"
	"reflect"
	"strconv"
)

/*
*
bean属性src拷贝给dst
*/
func BeanCopy(dst, src interface{}) (err error) {
	return New(src).To(dst)
}

func isHard(k reflect.Kind) bool {
	switch k {
	case reflect.Map, reflect.Slice, reflect.Ptr, reflect.Interface:
		return true
	}
	return false
}

func getFieldVal(val reflect.Value, field *reflect.StructField) (x reflect.Value) {
	if val.Kind() == reflect.Map {
		x = val.MapIndex(reflect.ValueOf(field.Name))
	} else {
		x = val.FieldByName(field.Name)
	}
	return
}

func getTargetMode(source, target Value) (r ConvertType) {
	if target.Upper().Kind() == reflect.Map {
		return AnyToJsonMap
	}
	if source.Upper().Kind() == reflect.Map && target.Upper().Kind() == reflect.Struct {
		return JsonMapToStruct
	}
	if source.Upper().Kind() == reflect.Struct && target.Upper().Kind() == reflect.Struct {
		return StructToStruct
	}
	panic("not support ConvertType: src kind=" + source.Upper().Kind().String() + ", tar kind=" + target.Upper().Kind().String())
	return
}

func (ctx *Context) To(val interface{}) (err error) {
	ref := reflect.ValueOf(val)
	if ref.Kind() != reflect.Ptr {
		err = errors.New("target must ptr map")
		return
	}
	ctx.valueB = Value(ref)
	ctx.direction = AtoB

	if ctx.valueA.Indirect().Upper().Kind() == reflect.Map && ctx.valueB.Indirect().Upper().Kind() == reflect.Map {
		if ctx.provideTyp == nil {
			ctx.provideTyp = ctx.valueA.Upper().Type()
			//err = errors.New("must set provide type")
			//return
		}
	} else {
		provideTyp, geterr := ctx.getProvideTyp(ctx.valueA, ctx.valueB)
		if geterr != nil {
			return
		}
		ctx.provideTyp = provideTyp
	}
	ctx.convertType = getTargetMode(ctx.valueA.Indirect(), ctx.valueB.Indirect())

	_, err = ctx.copy(ctx.valueA.Indirect(), ctx.valueB.Indirect().Indirect(), ctx.provideTyp.Elem(), false, 0)
	if err != nil {
		return
	}
	return
}

func (ctx *Context) From(val interface{}) (err error) {
	ref := reflect.ValueOf(val)
	if ref.Kind() != reflect.Ptr {
		err = errors.New("target must ptr map")
		return
	}
	ctx.valueB = Value(ref)
	ctx.direction = AfromB
	if ctx.valueA.Indirect().Upper().Kind() == reflect.Map && ctx.valueB.Indirect().Upper().Kind() == reflect.Map {
		if ctx.provideTyp == nil {
			ctx.provideTyp = ctx.valueA.Upper().Type()
			//err = errors.New("must set provide type")
			//return
		}
	} else {
		ctx.provideTyp, err = ctx.getProvideTyp(ctx.valueB, ctx.valueA)
		if err != nil {
			err = errors.New("must set provide type")
			return
		}
	}
	ctx.convertType = getTargetMode(ctx.valueB.Indirect(), ctx.valueA.Indirect())

	_, err = ctx.copy(ctx.valueB, ctx.valueA, ctx.provideTyp, false, 0)
	if err != nil {
		return
	}

	return
}

func (ctx *Context) getMethodType(val reflect.Value) (r string) {
	r = stringconv.UpperFirst(val.Kind().String())
	return
}

func (ctx *Context) copy(source, target Value, provideTyp reflect.Type, inInterface bool, depth int) (result Value, err error) {

	srcref := source.Upper()
	tarref := target.Upper()
	//
	//prefix := strings.Repeat("----", depth)
	//fmt.Println(prefix+"> copy:", "provide typ=", provideTyp, "kind=", provideTyp.Kind())
	//fmt.Println(prefix+"copy: srctyp=", srcref.Type(), "src=", srcref)
	//fmt.Println(prefix+"copy: tartyp=", target.GetTypeString(), "tar=", tarref, "nil=", ",  canset=", tarref.CanSet(), func() (x string) {
	//	if isHard(tarref.Kind()) && tarref.IsNil() {
	//		x = "isnil=true"
	//	} else {
	//		x = "isnil=false"
	//	}
	//	return
	//}())

	// 源是否空
	if srcref.IsValid() == false {
		return
	}
	if isHard(srcref.Kind()) && srcref.IsNil() {
		return
	}

	// 处理base map
	if ctx.baseMap.Contains(provideTyp) {
		if tarref.CanSet() == false {
			tarref = reflect.New(provideTyp)
			if provideTyp.Kind() != reflect.Ptr {
				tarref = tarref.Elem()
			}
		}
		tarref.Set(func() (x reflect.Value) {
			// 规定的类型跟源类型不一致的情况
			if srcref.Type() != provideTyp {
				switch srcref.Type().Kind() {
				case reflect.Interface:
					x = srcref.Elem().Convert(provideTyp)
				default:
					if provideTyp.Kind().String() == provideTyp.String() {
						x = srcref.Convert(provideTyp)
					} else {
						// 枚举
						err = errors.New("enum convert function not found")
						return
					}
				}
			} else {
				x = srcref
			}
			return
		}())
		result = Value(tarref)
		return
	}

	// 接口处理
	if provideTyp.Kind() != reflect.Interface {
		if srcref.Kind() == reflect.Interface {
			srcref = srcref.Elem()
		}
		if tarref.Kind() == reflect.Interface {
			tarref = tarref.Elem()
		}
	}

	// 修正来源值
	{
		// 创建新的值
		unfold := TypeUtiler.UnfoldType(provideTyp)
		switch unfold.Kind() {
		case reflect.Array, reflect.Slice:
			switch ctx.convertType {
			case JsonMapToStruct:
				if srcref.Kind() == reflect.Map { // 处理来源为map的情况
					var max = 0
					for _, key := range srcref.MapKeys() {
						idx, _ := strconv.Atoi(key.String())
						if idx > max {
							max = idx
						}
					}
					max = max + 1
					srcref = func() (x reflect.Value) {
						x = reflect.MakeSlice(reflect.SliceOf(srcref.Type().Elem()), max, max)
						for _, key := range srcref.MapKeys() {
							idx, _ := strconv.Atoi(key.String())
							x.Index(idx).Set(srcref.MapIndex(key))
						}
						return
					}()

				}
			}
		}

		source = Value(srcref)
		//fmt.Println(prefix+"copy: srctyp=", srcref.Type(), "src=", srcref, ",  canset=", srcref.CanSet(), func() (x string) {
		//	if isHard(srcref.Kind()) && srcref.IsNil() {
		//		x = "isnil=true"
		//	} else {
		//		x = "isnil=false"
		//	}
		//	return
		//}(), "<last>")
	}

	// 检查目标是否需要创建新的值
	checkNewTarget := func() (x bool) {
		if tarref.IsValid() == false {
			return true
		}
		switch tarref.Kind() {
		case reflect.Array, reflect.Slice:
			x = tarref.IsNil()
		case reflect.Struct:

		case reflect.Map, reflect.Ptr:
			x = tarref.IsNil()
		case reflect.Interface:
			if isHard(tarref.Elem().Kind()) {
				x = tarref.IsNil()
			} else {
				x = tarref.IsNil() || tarref.CanSet() == false
			}

		default:
			x = tarref.CanSet() == false
		}

		return
	}()

	if checkNewTarget {
		// 创建新的值
		unfold := TypeUtiler.UnfoldType(provideTyp)

		switch unfold.Kind() {
		case reflect.Array:
			tarref = reflect.New(provideTyp).Elem()
		case reflect.Slice:
			tarref = func() (x reflect.Value) {
				switch ctx.convertType {
				case AnyToJsonMap:
					if provideTyp == bytesTyp {
						if srcref.Kind() == reflect.String {
							x = reflect.New(stringTyp).Elem()
						} else {
							x = reflect.MakeSlice(provideTyp, srcref.Len(), srcref.Cap())
						}
					} else {
						slice := make([]interface{}, srcref.Len(), srcref.Cap())
						x = reflect.ValueOf(slice)
					}
				case StructToStruct:
					x = reflect.MakeSlice(provideTyp, srcref.Len(), srcref.Cap())
				case JsonMapToStruct:
					if srcref.Kind() == reflect.Map { // 处理来源为map的情况
						var max = 0
						for _, key := range srcref.MapKeys() {
							idx, _ := strconv.Atoi(key.String())
							if idx > max {
								max = idx
							}
						}
						max = max + 1
						x = reflect.MakeSlice(provideTyp, max, max)
						a := reflect.MakeSlice(reflect.SliceOf(srcref.Type().Elem()), max, max)
						for _, key := range srcref.MapKeys() {
							idx, _ := strconv.Atoi(key.String())
							a.Index(idx).Set(srcref.MapIndex(key))
						}
						srcref = a

					} else {
						if provideTyp == bytesTyp {
							if srcref.Kind() == reflect.String {
								x = reflect.New(stringTyp).Elem()
							} else {
								x = reflect.MakeSlice(provideTyp, srcref.Len(), srcref.Cap())
							}
						} else {
							if srcref.Kind() == reflect.Slice || srcref.Kind() == reflect.Array {
								x = reflect.MakeSlice(provideTyp, srcref.Len(), srcref.Cap())
							} else {
								x = reflect.New(provideTyp).Elem()
							}
						}
					}

				}

				return
			}()
		case reflect.Map:
			switch ctx.convertType {
			case AnyToJsonMap:
				if provideTyp.Kind() == reflect.Ptr {
					tarref = reflect.ValueOf(&map[string]interface{}{})
				} else {
					tarref = reflect.ValueOf(map[string]interface{}{})
				}
			case StructToStruct:
				tarref = reflect.MakeMap(unfold)
			case JsonMapToStruct:
				tarref = func() (x reflect.Value) {
					if provideTyp.Kind() == reflect.Ptr {
						x = reflect.New(provideTyp.Elem()).Elem()
						a := reflect.MakeMap(unfold)
						x.Set(a)
						x = x.Addr()
					} else {
						x = reflect.MakeMap(unfold)
					}
					return
				}()
				//}
			}
		case reflect.Struct:
			if ctx.convertType == AnyToJsonMap {
				if provideTyp.Kind() == reflect.Ptr {
					tarref = reflect.ValueOf(&map[string]interface{}{})
				} else {
					tarref = reflect.ValueOf(map[string]interface{}{})
				}
			} else {
				tarref = reflect.New(unfold)
				if provideTyp.Kind() != reflect.Ptr {
					tarref = tarref.Elem()
				}
			}
		default:
			tarref = reflect.New(provideTyp)
			if provideTyp.Kind() != reflect.Ptr {
				tarref = tarref.Elem()
			}
		}

		target = Value(tarref)
		//
		//fmt.Println(prefix+"copy: tartyp=", tarref.Type(), "tar=", tarref, ",  canset=", tarref.CanSet(), func() (x string) {
		//	if isHard(tarref.Kind()) && tarref.IsNil() {
		//		x = "isnil=true"
		//	} else {
		//		x = "isnil=false"
		//	}
		//	return
		//}(), "<last>")
	}

	// 如果存在To/From函数, 执行并返回结果
	switch ctx.direction {
	case AtoB:
		mtype, ok := srcref.Type().MethodByName("To")
		if ok == true {
			copyctx := &CopyContext{
				origin: ctx,
			}
			r, callerr := ctx.callToMethod(copyctx, srcref, "To", mtype)
			if callerr != nil {
				err = callerr
				return
			}
			if copyctx.ignore == false {
				if r.Kind() == reflect.Interface {
					r = r.Elem()
				}
				result = Value(r)
				return
			}
		}
	case AfromB:
		if tarref.IsValid() == true {
			mtype, ok := tarref.Type().MethodByName("From")
			if ok == true {
				copyctx := &CopyContext{
					origin: ctx,
				}
				r, callerr := ctx.callFromMethod(copyctx, srcref, tarref, "From", mtype)
				if callerr != nil {
					err = callerr
					return
				}
				if copyctx.ignore == false {
					if r.Kind() == reflect.Interface {
						r = r.Elem()
					}
					result = Value(r)
					return
				}
			}
		}
	}

	var retval Value
	switch provideTyp.Kind() {
	case reflect.Slice, reflect.Array:
		if srcref.Len() == 0 {
			return
		}

		switch ctx.convertType {
		case AnyToJsonMap:
			if provideTyp == bytesTyp {
				tarref = reflect.ValueOf(base64.StdEncoding.EncodeToString(srcref.Interface().([]byte)))
			} else {
				for i := 0; i < srcref.Len(); i++ {
					srcitem := srcref.Index(i)
					taritem := tarref.Index(i)
					retval, copyerr := ctx.copy(Value(srcitem), Value(taritem), provideTyp.Elem(), inInterface, depth+1)
					if copyerr != nil {
						err = copyerr
						return
					}

					tarref.Index(i).Set(retval.Upper())
				}
			}
		case JsonMapToStruct:
			if provideTyp == bytesTyp {
				if srcref.Kind() == reflect.String {
					b, _ := base64.StdEncoding.DecodeString(srcref.String())
					tarref = reflect.ValueOf(b)
				} else {
					reflect.Copy(tarref, srcref)
				}

			} else {
				if isHard(provideTyp.Elem().Kind()) || provideTyp.Elem().Kind() != srcref.Type().Kind() {
					for i := 0; i < srcref.Len(); i++ {
						srcitem := srcref.Index(i)
						taritem := tarref.Index(i)
						if srcitem.IsValid() == false || srcitem.IsNil() {
							continue
						}

						retval, copyerr := ctx.copy(Value(srcitem), Value(taritem), provideTyp.Elem(), inInterface, depth+1)
						if copyerr != nil {
							err = copyerr
							return
						}
						tarref.Index(i).Set(retval.Upper())
					}
				} else {
					reflect.Copy(tarref, srcref)
				}
			}

		}
	case reflect.Interface:
		switch ctx.convertType {
		case AnyToJsonMap:
			retval, err = ctx.copy(Value(srcref.Elem()), Value(tarref.Elem()), srcref.Elem().Type(), true, depth+1)
			if err != nil {
				return
			}
		case JsonMapToStruct:
			// 矫正一下map
			provideTyp = func() (x reflect.Type) {
				srcunfold := source.unfoldInterface()
				checkMap := func() (y bool) {
					switch srcunfold.Upper().Kind() {
					case reflect.Map:
						y = true
						return
					}
					return
				}()

				if checkMap {
					x, _ = source.parseMapType(ctx)
				} else {
					x = srcref.Elem().Type()
				}

				return
			}()

			retval, err = ctx.copy(Value(srcref.Elem()), Value(tarref.Elem()), provideTyp, true, depth+1)
			if err != nil {
				return
			}
		}

		if tarref.CanSet() {
			tarref.Set(retval.Upper())
		}
	case reflect.Ptr:
		switch ctx.convertType {
		case AnyToJsonMap:
			err = TypeUtiler.Call(srcref, "CopyBefore", ctx)
			if err != nil {
				return
			}
		}

		srcptr := func() (x Value) {
			if srcref.Kind() == reflect.Ptr {
				x = Value(srcref.Elem())
			} else {
				x = Value(srcref)
			}
			return
		}()

		tarptr := func() (x Value) {
			if tarref.Kind() == reflect.Ptr {
				x = Value(tarref.Elem())
			} else {
				x = Value(tarref)
			}
			return
		}()
		retval, err = ctx.copy(srcptr, tarptr, provideTyp.Elem(), inInterface, depth+1)
		if err != nil {
			return
		}
		if tarref.Kind() == reflect.Ptr {
			tarptr.Upper().Set(retval.Upper())
		}

		switch ctx.convertType {
		case AnyToJsonMap:
			if retval.Indirect().Upper().Kind() == reflect.Map {
				err = retval.updateMapPtrBy(source.unfoldInterface())
				if err != nil {
					return
				}

				if tarref.Kind() == reflect.Ptr {
					tarref = tarref.Elem()
				}
			}
		case JsonMapToStruct:
			if tarref.Kind() == reflect.Ptr {
				if tarref.Elem().Kind() == reflect.Map {
					tarref = tarref.Elem()
				}
			}
			err = TypeUtiler.Call(tarref, "CopyAfter", ctx)
			if err != nil {
				return
			}
		}

	case reflect.Struct:
		for _, field := range TypeUtiler.GetFieldRecursion(provideTyp) {
			if stringconv.IsLowerFirst(field.Name) {
				// 私有变量不作处理
				continue
			}

			key := reflect.ValueOf(field.Name)
			srcfield := func() (x reflect.Value) {
				if srcref.Kind() == reflect.Map {
					if ctx.Config.FieldTag != "" {
						tag, tagok := field.Tag.Lookup(ctx.Config.FieldTag)
						if tagok == false {
							x = srcref.MapIndex(reflect.ValueOf(field.Name))
						} else {
							x = srcref.MapIndex(reflect.ValueOf(tag))
						}
					} else {
						x = srcref.MapIndex(reflect.ValueOf(field.Name))
					}

				} else {
					x = srcref.FieldByName(field.Name)
				}

				return
			}()
			if srcref.Kind() == reflect.Map {
				if srcfield.IsValid() == false || srcfield.IsNil() {
					continue
				}
			}

			if ctx.Config.IgnoreDefault {
				if TypeUtiler.CompareEqualDefault(srcfield, field) {
					continue
				}
			}

			//fmt.Println(prefix+"struct: field=", field.Name, ", fieldtyp=", field.Type)
			// 获取目标值
			tarfield := getFieldVal(tarref, field)
			retval, err = ctx.copy(Value(srcfield), Value(tarfield), field.Type, inInterface, depth+1)
			if err != nil {
				return
			}

			switch tarref.Kind() {
			case reflect.Map:
				tarref.SetMapIndex(func() (x reflect.Value) {
					if ctx.Config.FieldTag != "" {
						tag, tagok := field.Tag.Lookup(ctx.Config.FieldTag)
						if tagok == false {
							x = key
						} else {
							x = reflect.ValueOf(tag)
						}
					} else {
						x = key
					}
					return
				}(), retval.Upper())
			case reflect.Struct:
				if retval.Upper().IsValid() {
					tarfield.Set(retval.Upper())
				}
			default:
				panic("not support in struct")
			}
			//if retval.Upper().IsValid() {
			//	fmt.Println(prefix+"struct[AF]: field=", field.Name, ", fieldval=", retval.Upper().Interface())
			//} else {
			//	fmt.Println(prefix+"struct[AF]: field=", field.Name, ", fieldval=", "not valid")
			//}
		}

		switch ctx.convertType {
		case AnyToJsonMap:
			if ctx.Config.AlwaysStructInfo || (inInterface && tarref.Kind() == reflect.Map) {
				err = Value(tarref).updateMapStructTypeBy(source.unfoldInterface())
				if err != nil {
					return
				}
			}
		case JsonMapToStruct:
		}

	case reflect.Map:
		for _, keySrc := range srcref.MapKeys() {
			//fmt.Println(prefix+"map: before copy key source: type=", keySrc.Type(), ", val=", keySrc.Interface())
			keyTar := reflect.New(provideTyp.Key()).Elem()
			keyTarVal, copyerr := ctx.copy(Value(keySrc), Value(keyTar), provideTyp.Key(), inInterface, depth+1)
			if copyerr != nil {
				err = copyerr
				return
			}
			//fmt.Println(prefix+"map: after copy key target: type=", keyTarVal.Upper().Type(), ", val=", keyTarVal.Upper().Interface())

			valTar := tarref.MapIndex(keyTarVal.Upper())
			valSrc := srcref.MapIndex(keySrc)
			if valSrc.IsValid() == false {
				continue
			}

			if valSrc.IsNil() {
				// 初始化目标
				tarref.SetMapIndex(keyTarVal.Upper(), valSrc)
			} else {
				//fmt.Println(prefix+"map: before copy value source : type=", valSrc.Type(), ", val=", valSrc.Interface())
				if valTar.IsValid() && valTar.IsNil() == false {
					//fmt.Println(prefix+"before copy value target: type=", valTar.Type(), ", val=", valTar.Interface())
				}
				varTarVal, copyerr := ctx.copy(Value(valSrc), Value(valTar), provideTyp.Elem(), inInterface, depth+1)
				if copyerr != nil {
					err = copyerr
					return
				}
				//fmt.Println(prefix+"map: after copy value target: type=", varTarVal.Upper().Type(), ", val=", varTarVal.Upper().Interface())
				tarref.SetMapIndex(keyTarVal.Upper(), varTarVal.Upper())
				//fmt.Println(prefix+"map: after copy value map: ", tarref.Interface())
			}

		}

	case reflect.Func:
		panic("function not support")
	default:
		if tarref.CanSet() == false {
			return
		}
		tarref.Set(func() (x reflect.Value) {
			// 规定的类型跟源类型不一致的情况
			if srcref.Type() != provideTyp {
				switch provideTyp.Kind() {
				case reflect.String:
					x = Convert2String(srcref)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					fallthrough
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					x = Convert2Int(srcref)
				case reflect.Float32, reflect.Float64:
					x = Convert2Float(srcref)
				case reflect.Bool:
					x = convert2Bool(srcref)
				default:
					err = errors.New("convert fail")
					return
				}

				if x.Type().ConvertibleTo(provideTyp) {
					x = x.Convert(provideTyp)
				}
			} else {
				x = srcref
			}

			return
		}())
		if ctx.convertType == AnyToJsonMap {
			tarref = Convert2MapValue(tarref)
		}
	}

	result = Value(tarref)

	//fmt.Println("resut >", result.Upper())
	return
}

func (ctx *Context) callToMethod(copyctx *CopyContext, srcref reflect.Value, mname string, mtype reflect.Method) (result reflect.Value, err error) {
	methodVal := srcref.MethodByName(mname)
	if mtype.Type.NumIn() > 2 {
		err = errors.New("func " + mname + " NumIn() must 1 or 0")
		return
	}
	results := func() (x []reflect.Value) {
		if mtype.Type.NumIn() == 2 {
			x = methodVal.Call([]reflect.Value{reflect.ValueOf(copyctx)})

		} else {
			x = methodVal.Call([]reflect.Value{})
		}
		return
	}()
	// 包含error
	if mtype.Type.NumOut() > 1 {
		if results[1].IsNil() == false {
			err = results[1].Interface().(error)
			return
		}
	}
	result = results[0]
	return
}

func (ctx *Context) callFromMethod(copyctx *CopyContext, srcref, tarref reflect.Value, mname string, mtype reflect.Method) (result reflect.Value, err error) {
	methodVal := tarref.MethodByName(mname)
	if mtype.Type.NumIn() > 3 || mtype.Type.NumIn() == 1 {
		err = errors.New("func " + mname + " NumIn() must 2 or 1")
		return
	}

	results := func() (x []reflect.Value) {
		if mtype.Type.NumIn() == 3 {
			x = methodVal.Call([]reflect.Value{reflect.ValueOf(copyctx), srcref})
		} else {
			x = methodVal.Call([]reflect.Value{srcref})
		}
		return
	}()
	switch mtype.Type.NumOut() {
	case 0:
		result = tarref
	case 1:
		if mtype.Type.Out(0).Implements(errTyp) {
			if tarref.Kind() != reflect.Ptr {
				err = errors.New("目标非指针类型, 必须返回运行结果")
				return
			}
			if results[0].IsNil() == false { // From返回错误
				err = results[0].Interface().(error)
				return
			}
			result = tarref
		} else {
			result = results[0]
		}

	case 2:
		if results[1].IsNil() == false { // From返回错误
			err = results[1].Interface().(error)
			return
		}

		result = results[0]
	}

	return
}
