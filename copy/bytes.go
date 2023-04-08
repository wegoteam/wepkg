package copy

import (
	"encoding/binary"
	"errors"
	"github.com/wegoteam/wepkg/conv/binconv"
	"io"
	"reflect"
	"strconv"
)

func (ctx *Context) toBytes(source Value, provideTyp reflect.Type, w io.Writer, depth int) (err error) {
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
		err = binary.Write(w, binary.BigEndian, int16(srcref.Len()))
		if srcref.Len() == 0 {
			return
		}
		for i := 0; i < srcref.Len(); i++ {
			srcitem := srcref.Index(i)
			err = ctx.toBytes(Value(srcitem), provideTyp.Elem(), w, depth+1)
			if err != nil {
				err = errors.New("at " + strconv.Itoa(i) + ": " + err.Error())
				return
			}
		}
	case reflect.Interface:
		err = ctx.toBytes(Value(srcref.Elem()), srcref.Elem().Type(), w, depth+1)
		if err != nil {
			return
		}
	case reflect.Ptr:
		err = ctx.toBytes(Value(srcref.Elem()), provideTyp.Elem(), w, depth+1)
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
			err = ctx.toBytes(Value(srcfield), field.Type, w, depth+1)
			if err != nil {
				err = errors.New(field.Name + ": " + err.Error())
				return
			}
		}
	case reflect.Map:
		err = binary.Write(w, binary.BigEndian, int16(srcref.Len()))
		if srcref.Len() == 0 {
			return
		}

		for _, k := range srcref.MapKeys() {
			err = binconv.WriteUTF(w, binary.BigEndian, k.String())
			val1 := srcref.MapIndex(k)
			if val1.IsValid() == false {
				continue
			}
			//fmt.Println("||| copy map key: ", k, ", fieldtyp=", val1.Type())
			//fmt.Println("src=", val1, ", typ=", val2)

			err = ctx.toBytes(Value(val1), val1.Type(), w, depth+1)
			if err != nil {
				err = errors.New("at " + k.String() + ": " + err.Error())
				return
			}
		}

	case reflect.Func:
		panic("not suppor")
	default:
		err = ctx.toBytesProp(provideTyp.Kind(), srcref, w)
		if err != nil {
			return
		}
	}

	//fmt.Println("resut >", result.Upper())
	return
}

func (ctx *Context) toBytesProp(kind reflect.Kind, val reflect.Value, w io.Writer) (err error) {
	switch kind {
	case reflect.Bool:
		err = binary.Write(w, binary.BigEndian, val.Bool())
	case reflect.Int:
		err = binary.Write(w, binary.BigEndian, int64(val.Int()))
	case reflect.Int8:
		err = binary.Write(w, binary.BigEndian, int8(val.Int()))
	case reflect.Int16:
		err = binary.Write(w, binary.BigEndian, int16(val.Int()))
	case reflect.Int32:
		err = binary.Write(w, binary.BigEndian, int32(val.Int()))
	case reflect.Int64:
		err = binary.Write(w, binary.BigEndian, int64(val.Int()))
	case reflect.Uint:
		err = binary.Write(w, binary.BigEndian, uint64(val.Uint()))
	case reflect.Uint8:
		err = binary.Write(w, binary.BigEndian, uint8(val.Uint()))
	case reflect.Uint16:
		err = binary.Write(w, binary.BigEndian, uint16(val.Uint()))
	case reflect.Uint32:
		err = binary.Write(w, binary.BigEndian, uint32(val.Uint()))
	case reflect.Uint64:
		err = binary.Write(w, binary.BigEndian, uint64(val.Uint()))
	case reflect.Float32:
		err = binary.Write(w, binary.BigEndian, float32(val.Float()))
	case reflect.Float64:
		err = binary.Write(w, binary.BigEndian, val.Float())
	case reflect.String:
		err = binconv.WriteUTF(w, binary.BigEndian, val.String())
	default:
		err = errors.New("unknown kind")
	}

	return
}

func (ctx *Context) fromBytes(source Value, provideTyp reflect.Type, r io.Reader, depth int) (err error) {
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
		length := int(binconv.ReadInt16(r, binary.BigEndian))
		if length == 0 {
			return
		}
		for i := 0; i < length; i++ {
			srcitem := reflect.New(provideTyp.Elem())
			err = ctx.fromBytes(Value(srcitem), provideTyp.Elem(), r, depth+1)
			if err != nil {
				err = errors.New("at " + strconv.Itoa(i) + ": " + err.Error())
				return
			}
		}
	case reflect.Interface:
		err = ctx.fromBytes(Value(srcref.Elem()), srcref.Elem().Type(), r, depth+1)
		if err != nil {
			return
		}
	case reflect.Ptr:
		err = ctx.fromBytes(Value(srcref.Elem()), provideTyp.Elem(), r, depth+1)
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
			err = ctx.fromBytes(Value(srcfield), field.Type, r, depth+1)
			if err != nil {
				err = errors.New(field.Name + ": " + err.Error())
				return
			}
		}
	case reflect.Map:
		length := int(binconv.ReadInt16(r, binary.BigEndian))
		if length == 0 {
			return
		}
		for i := 0; i < length; i++ {
			k := reflect.ValueOf(binconv.ReadUTF(r, binary.BigEndian))

			val := reflect.New(provideTyp.Elem())
			err = ctx.fromBytes(Value(val), val.Type(), r, depth+1)
			if err != nil {
				err = errors.New("at " + k.String() + ": " + err.Error())
				return
			}

			srcref.SetMapIndex(k, val)
			//fmt.Println("||| copy map key: ", k, ", fieldtyp=", val1.Type())
			//fmt.Println("src=", val1, ", typ=", val2)

		}

	case reflect.Func:
		panic("not suppor")
	default:
		val, fromerr := ctx.fromBytesProp(provideTyp.Kind(), r)
		if fromerr != nil {
			err = fromerr
			return
		}

		srcref.Set(val)
	}

	//fmt.Println("resut >", result.Upper())
	return
}

// 获取正确的反射对象，如果nil，创建新的
func unfoldValue(val reflect.Value) reflect.Value {
	var typ = val.Type()
	switch typ.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		typ = typ.Elem()
		if val.IsNil() {
			var obj = reflect.New(typ)
			val.Set(obj)
		}

		return val.Elem()
	}

	return val
}

// 获取正确的反射对象，如果nil，创建新的
func unfoldType(typ reflect.Type) reflect.Type {
	switch typ.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		typ = typ.Elem()
		return typ
	}

	return typ
}

func (ctx *Context) fromBytesProp(kind reflect.Kind, r io.Reader) (result reflect.Value, err error) {
	var data interface{}
	switch kind {
	case reflect.Bool:
		data = binconv.ReadBool(r, binary.BigEndian)
	case reflect.Int:
		data = int(binconv.ReadInt64(r, binary.BigEndian))
	case reflect.Int8:
		data = binconv.ReadInt8(r, binary.BigEndian)
	case reflect.Int16:
		data = binconv.ReadInt16(r, binary.BigEndian)
	case reflect.Int32:
		data = binconv.ReadInt32(r, binary.BigEndian)
	case reflect.Int64:
		data = binconv.ReadInt64(r, binary.BigEndian)
	case reflect.Uint:
		data = uint(binconv.ReadUint64(r, binary.BigEndian))
	case reflect.Uint8:
		data = binconv.ReadUint8(r, binary.BigEndian)
	case reflect.Uint16:
		data = binconv.ReadUint16(r, binary.BigEndian)
	case reflect.Uint32:
		data = binconv.ReadUint32(r, binary.BigEndian)
	case reflect.Uint64:
		data = binconv.ReadUint64(r, binary.BigEndian)
	case reflect.Float32:
		data = binconv.ReadFloat32(r, binary.BigEndian)
	case reflect.Float64:
		data = binconv.ReadFloat64(r, binary.BigEndian)
	case reflect.String:
		data = binconv.ReadUTF(r, binary.BigEndian)
	default:
		err = errors.New("not support default kind: " + kind.String())
	}

	result = reflect.ValueOf(data)
	return
}
