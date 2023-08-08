package file

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// FilePathInfo
// @Description: FilePathInfo描述文件或目录的路径和状态。
type FilePathInfo struct {
	Path string
	Info os.FileInfo
}

// ListAll
// @Description: 返回给定目录中所有条目的列表。给定的目录不包含在列表中。
//它递归搜索，但不会跟随给定路径以外的符号链接。
// @param: root
// @return entries
// @return err
func ListAll(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return true, nil })
}

// ListFile
// @Description: 返回给定目录中所有文件条目的列表。给定的目录不包含在列表中。它会递归搜索，但不会跟踪给定路径以外的符号链接。
// @param: root
// @return entries
// @return err
func ListFile(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return isFileInfo(&info), nil })
}

// ListSymlink
// @Description: 返回给定目录中所有符号链接条目的列表。给定的目录不包含在列表中。它会递归搜索，但不会跟踪给定路径以外的符号链接。
// @param: root
// @return entries
// @return err
func ListSymlink(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return isSymlinkFileInfo(&info), nil })
}

// ListDir
// @Description: 返回给定目录中所有嵌套目录条目的列表。给定的目录不包含在列表中。它会递归搜索，但不会跟踪给定路径以外的符号链接。
// @param: root
// @return entries
// @return err
func ListDir(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return isDirFileInfo(&info), nil })
}

//这些标志由ListMatch方法使用
const (
	// ListRecursive 表示使用ListMatch递归地列出遇到的目录项。
	ListRecursive int = 1 << iota
	// ListToLower ListMatch用于在模式匹配之前将文件名转换为小写
	ListToLower
	// ListUseRegExp 指定ListMatch使用正则表达式进行模式匹配。
	ListUseRegExp
	// ListIncludeDir 通过ListMatch将匹配的目录包含在返回的列表中。
	ListIncludeDir
	// ListIncludeFile 通过ListMatch将匹配的文件包含在返回的列表中。
	ListIncludeFile
	// ListIncludeSymlink 指定ListMatch，将匹配的符号链接包含在返回的列表中。
	ListIncludeSymlink
)

const (
	// ListIncludeAll indicates ListMatch to include all the matched in the returned list.
	ListIncludeAll = ListIncludeDir | ListIncludeFile | ListIncludeSymlink
)

// ListMatch
// @Description: ListMatch返回一个目录条目的列表，该列表中的条目按词法顺序匹配目录中的任何给定模式。
//指定路径以外的符号链接不会被跟踪给定的目录没有包含在列表中。
// ListMatch要求模式匹配完整的文件名，而不仅仅是子字符串如果任何模式的格式不正确，则返回Errors。
//支持两种模式类型:
// 1)在filepath.Match()中描述的通配符，这是默认值;
// @param: root
// @param: flag
// @param: patterns
// @return entries
// @return err
func ListMatch(root string, flag int, patterns ...string) (entries []*FilePathInfo, err error) {
	var (
		rePatterns   []*regexp.Regexp
		typeFlag     = flag & ListIncludeAll
		useRegExp    = flag&ListUseRegExp != 0
		useLowerName = flag&ListToLower != 0
	)
	if useRegExp {
		if rePatterns, err = compileRegexpList(patterns); err != nil {
			return
		}
	}

	return listCondEntries(root, func(info os.FileInfo) (ok bool, err error) {
		fileName := info.Name()
		if useLowerName {
			fileName = strings.ToLower(fileName)
		}

		if isFileTypeMatched(&info, typeFlag) {
			if useRegExp {
				for _, pat := range rePatterns {
					if ok = pat.MatchString(fileName); ok {
						break
					}
				}
			} else {
				for _, pat := range patterns {
					if ok, err = filepath.Match(pat, fileName); ok || err != nil {
						break
					}
				}
			}
		}

		if err == nil && (flag&ListRecursive == 0) && isDirFileInfo(&info) {
			err = filepath.SkipDir
		}
		return
	})
}

// listCondEntries
// @Description: listCondEntries返回一个条件目录条目的列表。
// @param: root
// @param: cond
// @return entries
// @return err
func listCondEntries(root string, cond func(os.FileInfo) (bool, error)) (entries []*FilePathInfo, err error) {
	var (
		rootFi   os.FileInfo
		rootPath string
	)
	if rootPath, rootFi, err = resolveDirInfo(root); err != nil {
		err = opError(opnList, root, err)
		return
	}

	err = filepath.Walk(rootPath, func(itemPath string, itemFi os.FileInfo, errIn error) (errOut error) {
		errOut = errIn
		if os.SameFile(rootFi, itemFi) || errOut != nil {
			return
		}
		var ok bool
		if ok, errOut = cond(itemFi); ok {
			entries = append(entries, &FilePathInfo{
				Path: itemPath,
				Info: itemFi,
			})
		}
		return
	})
	return
}

// isFileTypeMatched
// @Description: 检查文件类型是否与标志匹配。
// @param: info
// @param: flag
// @return match
func isFileTypeMatched(info *os.FileInfo, flag int) (match bool) {
	switch {
	case flag == ListIncludeAll:
		match = true
	case flag&ListIncludeDir != 0 && isDirFileInfo(info):
		match = true
	case flag&ListIncludeFile != 0 && isFileInfo(info):
		match = true
	case flag&ListIncludeSymlink != 0 && isSymlinkFileInfo(info):
		match = true
	}
	return
}
