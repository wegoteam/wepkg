package file

import (
	"os"
	"path/filepath"
)

// GetFileSize
// @Description: 返回常规文件的大小（以字节为单位）。
//如果给定的路径是符号链接，它将被跟随。
// @param: path
// @return size
// @return err
func GetFileSize(path string) (size int64, err error) {
	var fi os.FileInfo
	if fi, err = os.Stat(path); err == nil {
		if isFileInfo(&fi) {
			size = fi.Size()
		} else {
			err = opError(opnSize, path, errNotRegularFile)
		}
	}
	return
}

// GetSymlinkSize
// @Description: 返回符号链接的大小（以字节为单位）。
// @param: path
// @return size
// @return err
func GetSymlinkSize(path string) (size int64, err error) {
	var fi os.FileInfo
	if fi, err = os.Lstat(path); err == nil {
		if isSymlinkFileInfo(&fi) {
			size = fi.Size()
		} else {
			err = opError(opnSize, path, errNotSymlink)
		}
	}
	return
}

// GetDirSize
// @Description: 返回目录中所有常规文件和符号链接的总大小（以字节为单位）。
//如果给定的路径是符号链接，它将被跟随，但目录中的符号链接不会。
// @param: path
// @return size
// @return err
func GetDirSize(path string) (size int64, err error) {
	var (
		rootFi os.FileInfo
		root   string
	)
	if root, rootFi, err = resolveDirInfo(path); err == nil {
		err = filepath.Walk(root, func(itemPath string, itemFi os.FileInfo, errIn error) (errOut error) {
			errOut = errIn
			if os.SameFile(rootFi, itemFi) || errOut != nil {
				return
			}
			if isFileInfo(&itemFi) || isSymlinkFileInfo(&itemFi) {
				size += itemFi.Size()
			}
			return
		})
	} else {
		err = opError(opnSize, path, err)
	}
	return
}
