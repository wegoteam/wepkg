package file

// SortListByName 实现排序根据FilePathInfo的Info.Name()字段定义的接口。
type SortListByName []*FilePathInfo

func (n SortListByName) Len() int {
	return len(n)
}

func (n SortListByName) Less(i, j int) bool {
	return n[i].Info.Name() < n[j].Info.Name()
}

func (n SortListByName) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// SortListBySize 实现排序接口，基于FilePathInfo的Info.Size()字段。
type SortListBySize []*FilePathInfo

func (s SortListBySize) Len() int {
	return len(s)
}

func (s SortListBySize) Less(i, j int) bool {
	return s[i].Info.Size() < s[j].Info.Size()
}

func (s SortListBySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SortListByModTime 实现了排序基于FilePathInfo的Info.ModTime()字段的接口。
type SortListByModTime []*FilePathInfo

func (t SortListByModTime) Len() int {
	return len(t)
}

func (t SortListByModTime) Less(i, j int) bool {
	return t[i].Info.ModTime().Before(t[j].Info.ModTime())
}

func (t SortListByModTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
