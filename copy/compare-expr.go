package copy

import "regexp"

type Expr func(val interface{}) bool

// 大于
// scale表示, 比较前对源值进行放大的倍数
func GreaterThan(number int, scale ...float64) (r Expr) {
	s := func() (x float64) {
		if len(scale) > 0 {
			return scale[0]
		}
		return 1
	}()

	r = func(val interface{}) bool {
		return int(s*convert2FloatNotReflect(val)) > number
	}
	return
}

// 大于等于
// scale表示, 比较前对源值进行放大的倍数
func GreaterThanOrEqualTo(number int, scale ...float64) (r Expr) {
	s := func() (x float64) {
		if len(scale) > 0 {
			return scale[0]
		}
		return 1
	}()

	r = func(val interface{}) bool {
		return int(s*convert2FloatNotReflect(val)) >= number
	}
	return
}

// 小于
// scale表示, 比较前对源值进行放大的倍数
func LessThan(number int, scale ...float64) (r Expr) {
	s := func() (x float64) {
		if len(scale) > 0 {
			return scale[0]
		}
		return 1
	}()

	r = func(val interface{}) bool {
		return int(s*convert2FloatNotReflect(val)) < number
	}
	return
}

// 小于等于
// scale表示, 比较前对源值进行放大的倍数
func LessThanOrEqualTo(number int, scale ...float64) (r Expr) {
	s := func() (x float64) {
		if len(scale) > 0 {
			return scale[0]
		}
		return 1
	}()

	r = func(val interface{}) bool {
		return int(s*convert2FloatNotReflect(val)) <= number
	}
	return
}

func NotEqualTo(number int, scale ...float64) (r Expr) {
	s := func() (x float64) {
		if len(scale) > 0 {
			return scale[0]
		}
		return 1
	}()

	r = func(val interface{}) bool {
		return int(s*convert2FloatNotReflect(val)) != number
	}
	return
}

func EqualTo(number int, scale ...float64) (r Expr) {
	s := func() (x float64) {
		if len(scale) > 0 {
			return scale[0]
		}
		return 1
	}()

	r = func(val interface{}) bool {
		return int(s*convert2FloatNotReflect(val)) == number
	}
	return
}

func Regexp(exp string) (r Expr) {
	re := regexp.MustCompile(exp)
	r = func(val interface{}) bool {
		return re.MatchString(convert2StringNotReflect(val))
	}
	return
}
