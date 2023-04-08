package arrayconv

import (
	"sort"
)

type ArrayList []interface{}

// 添加, 如果存在, 不处理
func (arr *ArrayList) Add(items ...interface{}) {
	for _, item := range items {
		if arr.Contains(item) {
			return
		}
		arr.Push(item)
	}
}

func (arr *ArrayList) At(index int) interface{} {
	return (*arr)[index]
}

func (arr *ArrayList) Clear() {
	*arr = (*arr)[:0]
}

func (arr *ArrayList) Contains(item interface{}) bool {
	return arr.IndexOf(item) != -1
}

func (arr *ArrayList) ContainsCond(cond func(item interface{}) bool) bool {
	return arr.IndexOf(cond) != -1
}

// 每一个项都符合条件就返回true
func (arr *ArrayList) Every(cond func(item interface{}) bool) bool {
	s := *arr
	for i := 0; i < len(s); i++ {
		val := s[i]
		if cond(val) == false {
			return false
		}
	}
	return true
}

func (arr *ArrayList) First(cond func(item interface{}) bool) (val interface{}, has bool) {
	s := *arr
	for i := 0; i < len(s); i++ {
		val := s[i]
		if cond(val) {
			return val, true
		}
	}
	return nil, false
}

func (arr *ArrayList) Filter(cond func(index int, elem interface{}) bool) (r ArrayList) {
	for i, x := range *arr {
		if cond(i, x) {
			r = append(r, x)
		}
	}
	return r
}

func (arr *ArrayList) ForRange(handler func(item interface{})) {
	for _, x := range *arr {
		handler(x)
	}
}

func (arr *ArrayList) IndexOfConf(cond func(item interface{}) bool) int {
	s := *arr
	for i := 0; i < len(s); i++ {
		if cond(s[i]) {
			return i
		}
	}
	return -1
}

func (arr *ArrayList) IndexOf(item interface{}) int {
	s := *arr
	for i := 0; i < len(s); i++ {
		if s[i] == item {
			return i
		}
	}
	return -1
}

func (arr *ArrayList) Last(cond func(item interface{}) bool) interface{} {
	s := *arr
	for i := len(s) - 1; i >= 0; i-- {
		val := s[i]
		if cond(val) {
			return val
		}
	}
	return nil
}

func (arr *ArrayList) Length() int {
	return len(*arr)
}

func (arr *ArrayList) Pop() interface{} {
	s := *arr
	last := s[len(s)-1]
	s[len(s)-1] = nil // GC
	s2 := s[:len(s)-1]
	*arr = s2

	return last
}

func (arr *ArrayList) Push(item interface{}) {
	s := *arr
	s = append(s, item)
	*arr = s
}

func (arr *ArrayList) PushList(list ArrayList) {
	s := *arr
	s = append(s, list...)
	*arr = s
}

func (arr *ArrayList) Remove(item interface{}) {
	i := arr.IndexOf(item)
	if i != -1 {
		arr.RemoveAt(i)
	}
}

func (arr *ArrayList) RemoveAt(i int) {
	s := *arr
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil // GC
	s2 := s[:len(s)-1]
	*arr = s2
}

func (arr *ArrayList) Replace(i int, item interface{}) {
	s := *arr
	over := i - len(s)
	if over > -1 {
		ss := make([]interface{}, i+1)
		copy(ss[0:], s[:])
		s = ss
	}
	s[i] = item
	*arr = s
}

func (arr *ArrayList) Reverse() {
	for i := len(*arr)/2 - 1; i >= 0; i-- {
		opp := len(*arr) - 1 - i
		(*arr)[i], (*arr)[opp] = (*arr)[opp], (*arr)[i]
	}
}

func (arr *ArrayList) Shift() interface{} {
	s := *arr
	top := s[0]
	s[0] = nil // GC
	s2 := s[1:]
	*arr = s2

	return top
}

func (arr *ArrayList) Slice() []interface{} {
	return []interface{}(*arr)
}

func (arr *ArrayList) Sort(compare func(a, b interface{}) int) {
	l := *arr
	sort.Slice(l, func(i, j int) bool {
		return compare(l[i], l[j]) >= 0
	})
}

func (arr *ArrayList) Unshift(item interface{}) {
	s := *arr
	l := len(s) + 1
	ss := make([]interface{}, l, l)
	ss[0] = item
	copy(ss[1:], s[:])
	*arr = ss
}

// 去重操作, 返回去重后的数组
func (arr *ArrayList) Unique(getKey func(a interface{}) string) (r ArrayList) {
	l := *arr
	m := map[string]interface{}{} // 存放不重复主键
	for _, e := range l {
		length := len(m)
		m[getKey(e)] = 0
		if len(m) != length { // 加入map后，map长度变化，则元素不重复
			r = append(r, e)
		}
	}
	return
}

// 并集
func (arr *ArrayList) Union(a ArrayList, getKey func(a interface{}) string) ArrayList {
	arr.PushList(a)
	return arr.Unique(getKey)
}
