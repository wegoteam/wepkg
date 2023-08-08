package file

import (
	"os"
	"path/filepath"
)

// JoinPath
// @Description: 将任意数量的路径元素连接成一个路径，必要时添加分隔符。
// @param: elem
// @return string
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

// ChangeExeDir
// @Description: 将当前工作目录更改为启动当前进程的可执行文件所在的目录。
//如果使用符号链接启动进程，当前工作目录将变为该链接所指向的可执行文件的目录。
// @return err
func ChangeExeDir() (err error) {
	var (
		ap string
		fi os.FileInfo
	)
	// 获取启动当前进程的可执行文件的路径
	if ap, err = os.Executable(); err != nil {
		err = opError(opnChange, ap, err)
		return
	}

	// 获取可执行文件的文件信息，如果它是符号链接，则解析路径
	if fi, err = os.Lstat(ap); err == nil && isSymlinkFileInfo(&fi) {
		ap, err = filepath.EvalSymlinks(ap)
	}
	if err != nil {
		err = opError(opnChange, ap, err)
		return
	}

	// 获取可执行目录并将当前工作目录更改为该目录
	if err = os.Chdir(filepath.Dir(ap)); err != nil {
		err = opError(opnChange, ap, err)
	}
	return
}

// MakeDir
// @Description: MakeDir创建一个名为path的目录，该目录具有0755个权限位，以及所有必要的父目录。
// 权限位表示所有者可以读、写和执行，其他所有人都可以读和执行，但不能修改。
//如果路径已经是一个目录，MakeDir什么也不做，返回nil
// @param: path
// @return err
func MakeDir(path string) (err error) {
	if err = os.MkdirAll(path, defaultDirectoryPermMode); err != nil {
		err = opError(opnMake, path, err)
	}
	return
}
