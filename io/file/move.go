package file

import (
	"os"
)

// MoveFile
// @Description: 将文件移动到目标文件或目录。符号链接将不会被跟踪。
//如果目标文件是一个存在的文件，则目标文件将被源文件覆盖
//如果目标文件是一个已经存在的目录，源文件将被移动到同名目录中。
//如果目标文件不存在，但它的父目录存在，源文件将被移动到具有目标名称的父目录。
//如果有错误，它的类型将是*os. pathror。
// @param: src
// @param: dest
// @return err
func MoveFile(src, dest string) (err error) {
	return moveEntry(
		src, dest,
		isFileInfo, errNotRegularFile,
		os.Remove,
		func(src, dest string) error { return bufferCopyFile(src, dest, defaultBufferSize) })
}

// MoveSymlink
// @Description: 将符号链接移动到目标文件或目录。
//如果目标是一个存在的文件或链接，目标将被源链接覆盖。
//如果目标是一个存在的目录，源链接将被移动到同名的目录中。
//如果目标不存在，但其父目录存在，则源链接将被移动到具有目标名称的父目录。
//如果有错误，它的类型将是*os. pathror。
// @param: src
// @param: dest
// @return err
func MoveSymlink(src, dest string) (err error) {
	return moveEntry(
		src, dest,
		isSymlinkFileInfo, errNotSymlink,
		os.Remove,
		func(src, dest string) error { return copySymlink(src, dest) })
}

// MoveDir
// @Description: 递归地将目录移动到目标目录。目录中的符号链接将不会被跟踪。
//如果目标是一个存在的文件，它将被源目录取代。
//如果目标目录是一个已经存在的目录，则源目录将被移动到同名目录中。
//如果目标不存在，但其父目录存在，则源目录将被移动到具有目标名称的父目录中。
//如果发生任何错误，MoveDir将立即停止并返回，错误类型为*os. pathror。
// @param: src
// @param: dest
// @return err
func MoveDir(src, dest string) (err error) {
	return moveEntry(
		src, dest,
		isDirFileInfo, errNotDirectory,
		os.RemoveAll,
		func(src, dest string) error { return copyDir(src, dest) })
}

// moveEntry
// @Description: 移动源到目标通过重命名或复制。
// @param: src
// @param: dest
// @param: check
// @param: errMode
// @param: remove
// @param: copy
// @return err
func moveEntry(src, dest string, check funcCheckFileInfo, errMode error, remove funcRemoveEntry, copy funcCopyEntry) (err error) {
	// 验证和优化路径
	if src, dest, err = refineOpPaths(opnMove, src, dest, false); err != nil {
		return
	}

	// 检查source是否存在以及它的文件模式
	var srcInfo os.FileInfo
	if srcInfo, err = os.Lstat(src); err == nil && !check(&srcInfo) {
		err = opError(opnMove, src, errMode)
	}
	if err != nil {
		return
	}

	// 试图通过重命名链接来移动文件
	if err = os.Rename(src, dest); os.IsExist(err) || isLinkErrorNotDirectory(err) {
		// 如果目标目录存在或不存在，则删除目标
		_ = remove(dest)
		err = os.Rename(src, dest)
	}

	switch {
	case err == nil:
		// 重命名成功
	case isLinkErrorCrossDevice(err):
		//移动设备==移除dest +复制到dest +移除SRC
		//删除目标文件，忽略文件不存在错误
		if err = remove(dest); err != nil && !os.IsNotExist(err) {
			err = opError(opnMove, dest, err)
			return
		}
		if err = copy(src, dest); err == nil {
			err = remove(src)
		}
	case os.IsNotExist(err):
		err = opError(opnMove, src, err)
	case os.IsExist(err):
		err = opError(opnMove, dest, err)
	}
	return
}
