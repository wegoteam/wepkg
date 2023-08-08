package file

import (
	"os"
	"path/filepath"
)

// IsFileEmpty
// @Description: 检查给定的文件是否为空。
// @param: path
// @return empty
// @return err
func IsFileEmpty(path string) (empty bool, err error) {
	var fi os.FileInfo
	if fi, err = os.Stat(path); err == nil {
		if isFileInfo(&fi) {
			empty = fi.Size() == 0
		} else {
			err = opError(opnEmpty, path, errNotRegularFile)
		}
	}
	return
}

// IsDirEmpty
// @Description: 检查给定的目录是否为空。
// @param: path
// @return empty
// @return err
func IsDirEmpty(path string) (empty bool, err error) {
	var (
		rootFi os.FileInfo
		root   string
	)
	if root, rootFi, err = resolveDirInfo(path); err == nil {
		err = filepath.Walk(root, func(itemPath string, itemFi os.FileInfo, errItem error) error {
			if os.SameFile(rootFi, itemFi) || errItem != nil {
				return errItem
			}
			// 强制退出根目录之外的第一个条目
			return errStepOutDir
		})

		if err == nil {
			empty = true
		} else if err == errStepOutDir {
			err = nil
		}
	} else {
		err = opError(opnEmpty, path, err)
	}
	return
}
