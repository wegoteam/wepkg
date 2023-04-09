package copy

import (
	"errors"
	"reflect"
)

func New(val interface{}) (ctx *Context) {
	ref := reflect.ValueOf(val)
	return NewValue(ref)
}

func NewValue(ref reflect.Value) (ctx *Context) {
	if ref.Kind() != reflect.Ptr {
		panic(errors.New("origin must ptr struct or map"))
	}

	unfold := TypeUtiler.UnfoldType(ref.Type())
	if unfold.Kind() != reflect.Struct && unfold.Kind() != reflect.Map {
		panic(errors.New("origin must ptr struct or map"))
	}

	ctx = &Context{
		valueA:  Value(ref),
		Config:  NewConfig(),
		baseMap: NewTypeSet(),
	}
	return
}

type Direction int

const (
	AtoB Direction = iota
	AfromB
)

// 转换模式
type ConvertType int

const (
	// struct->map
	// map->map
	AnyToJsonMap ConvertType = iota
	// map->struct
	JsonMapToStruct
	// struct->struct
	StructToStruct
)

// 数据源的上下文
type Context struct {
	// 值A-New()方法赋值
	valueA Value
	// 值B-To/From()方法赋值
	valueB Value
	// copy方向
	direction Direction
	// 转换类型
	convertType ConvertType
	// 规定的类型
	provideTyp reflect.Type
	// 自定义的参数, 传递给转化函数使用
	params map[string]interface{}
	// 进行中的标签
	inProcess string
	// 配置
	Config *Config
	// 类型的映射
	typeMap *TypeSet
	// 视作base类型的类型
	baseMap *TypeSet
	//
	compareErrors []error
	// 是否比较类型
	compareType bool
	// 是否比较所有字段
	compareAll bool
}

func NewConfig() (obj *Config) {
	obj = &Config{}
	return
}

type Config struct {
	// 转换成map时, 检查FieldTag定义的名字, 例如json/bson, 根据FieldTag转换成对应的Field名字
	// 例如在Id 字段 定义了bson:"_id", 转换后的map["Id"] 变成 map["_id"]
	FieldTag string
	// 当转化成map时, 是否总是携带结构信息, 包括_type和_ptr
	AlwaysStructInfo bool
	// 拷贝时, 如果来源的值等于默认值, 将被忽略
	// 可通过设置属性的tag: value:"" , 设置默认值. 如果没有设置, 根据属性类型的默认值
	IgnoreDefault bool
}

func (ctx *Context) getProvideTyp(src, tar Value) (typ reflect.Type, err error) {
	typ = ctx.provideTyp
	srcref := src.Upper()
	tarref := tar.Upper()
	if typ == nil {
		indirect := reflect.Indirect(tarref)
		if indirect.Kind() == reflect.Struct {
			typ = tarref.Type()
			return
		}

		indirect = reflect.Indirect(srcref)
		if indirect.Kind() == reflect.Struct {
			typ = srcref.Type()
			return
		}
	}

	err = errors.New("not found")
	return
}

func (ctx *Context) WithProvideTyp(val reflect.Type) *Context {
	ctx.provideTyp = val
	return ctx
}

func (ctx *Context) WithTypeMap(val *TypeSet) *Context {
	ctx.typeMap = val
	return ctx
}

func (ctx *Context) WithBaseTypes(val *TypeSet) *Context {
	ctx.baseMap = val
	return ctx
}

func (ctx *Context) WithParams(val map[string]interface{}) *Context {
	ctx.params = val
	return ctx
}

func (ctx *Context) WithConfig(val *Config) *Context {
	ctx.Config = val
	return ctx
}

func (ctx *Context) WithFieldTag(tag string) *Context {
	ctx.Config.FieldTag = tag
	return ctx
}

func (ctx *Context) WithIgnoreDefault() *Context {
	ctx.Config.IgnoreDefault = true
	return ctx
}

func (ctx *Context) InProcess(tag string) *Context {
	ctx.inProcess = tag
	return ctx
}

type CopyContext struct {
	//
	ignore bool

	//
	origin *Context
}

func (ctx *CopyContext) Ignore() {
	ctx.ignore = true
}

// 获取参数
// 可使用maputil操作
func (ctx *CopyContext) Params() map[string]interface{} {
	return ctx.origin.params
}

// 获取参数
// 可使用maputil操作
func (ctx *CopyContext) InProcess(tag string) bool {
	return ctx.origin.inProcess == tag
}

func (ctx *CopyContext) Process() string {
	return ctx.origin.inProcess
}
