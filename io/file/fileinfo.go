package file

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

var (
	errInvalidPath    = errors.New("无效的路径")
	errSameFile       = errors.New("两个文件相同")
	errShortRead      = errors.New("分区的磁盘短读")
	errIsDirectory    = errors.New("是一个目录")
	errNotDirectory   = errors.New("不是一个目录")
	errNotRegularFile = errors.New("非普通文件")
	errNotSymlink     = errors.New("非符号链接")
	errStepOutDir     = errors.New("跳出这个目录")
)

var (
	opnCompare = "compare"
	opnCopy    = "copy"
	opnMove    = "move"
	opnList    = "list"
	opnSize    = "size"
	opnEmpty   = "empty"
	opnChange  = "change"
	opnMake    = "make"
)

// underlyingError
// @Description: 返回给定错误的基础错误。
// @param: err
// @return error
func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.LinkError:
		return err.Err
	case *os.PathError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	}
	return err
}

// opError
// @Description: 返回带有给定细节的错误结构体
// @param: op
// @param: path
// @param: err
// @return *os.PathError
func opError(op, path string, err error) *os.PathError {
	return &os.PathError{
		Op:   op,
		Path: path,
		Err:  underlyingError(err),
	}
}

type (
	funcStatFileInfo  func(name string) (os.FileInfo, error)
	funcCheckFileInfo func(fi *os.FileInfo) bool
	funcRemoveEntry   func(path string) error
	funcCopyEntry     func(src, dest string) error
)

// isFileInfo
// @Description: 指示FileInfo是否为常规文件。
// @param: fi
// @return bool
func isFileInfo(fi *os.FileInfo) bool {
	return fi != nil && (*fi).Mode().IsRegular()
}

// isDirFileInfo
// @Description: 指示FileInfo是否为目录。
// @param: fi
// @return bool
func isDirFileInfo(fi *os.FileInfo) bool {
	return fi != nil && (*fi).Mode().IsDir()
}

// isSymlinkFileInfo
// @Description: 指示FileInfo是否为符号链接。
// @param: fi
// @return bool
func isSymlinkFileInfo(fi *os.FileInfo) bool {
	return fi != nil && ((*fi).Mode()&os.ModeType == os.ModeSymlink)
}

// isLinkErrorCrossDevice
// @Description: 指示LinkError是否为跨设备错误。
// @param: err
// @return bool
func isLinkErrorCrossDevice(err error) bool {
	lerr, ok := err.(*os.LinkError)
	return ok && lerr.Err == syscall.EXDEV
}

// isLinkErrorNotDirectory
// @Description: 指示LinkError是否为非目录错误。
// @param: err
// @return bool
func isLinkErrorNotDirectory(err error) bool {
	lerr, ok := err.(*os.LinkError)
	return ok && lerr.Err == syscall.ENOTDIR
}

// refineOpPaths
// @Description: 验证，清理和调整源和目标路径，以进行诸如复制或移动之类的操作。
// @param: opName
// @param: srcRaw
// @param: destRaw
// @param: followLink
// @return src
// @return dest
// @return err
func refineOpPaths(opName, srcRaw, destRaw string, followLink bool) (src, dest string, err error) {
	// 验证路径
	if strings.TrimSpace(srcRaw) == "" {
		err = opError(opName, srcRaw, errInvalidPath)
	} else if strings.TrimSpace(destRaw) == "" {
		err = opError(opName, destRaw, errInvalidPath)
	}
	if err != nil {
		return
	}

	// 清理路径
	src, dest = filepath.Clean(srcRaw), filepath.Clean(destRaw)

	// 使用操作系统。开始跟随符号链接
	statFunc := os.Lstat
	if followLink {
		statFunc = os.Stat
	}

	// 检查源是否存在
	var srcInfo, destInfo os.FileInfo
	if srcInfo, err = statFunc(src); err != nil {
		return
	}

	// 检查目的地是否存在
	if destInfo, err = statFunc(dest); err != nil {
		// 检查缺失目的地的父节点是否存在
		if os.IsNotExist(err) {
			_, err = os.Stat(filepath.Dir(dest))
		}
	} else {
		if os.SameFile(srcInfo, destInfo) {
			err = opError(opName, dest, errSameFile)
		} else if destInfo.IsDir() {
			// 将源文件名附加到现有目标的路径
			dest = JoinPath(dest, srcInfo.Name())
		}
	}
	return
}

// refineComparePaths
// @Description: 验证，清理文件比较的路径。
// @param: pathRaw1
// @param: pathRaw2
// @return path1
// @return path2
// @return err
func refineComparePaths(pathRaw1, pathRaw2 string) (path1, path2 string, err error) {
	if strings.TrimSpace(pathRaw1) == "" {
		err = opError(opnCompare, pathRaw1, errInvalidPath)
	} else if strings.TrimSpace(pathRaw2) == "" {
		err = opError(opnCompare, pathRaw2, errInvalidPath)
	}

	if err == nil {
		path1, path2 = filepath.Clean(pathRaw1), filepath.Clean(pathRaw2)
	}
	return
}

// resolveDirInfo
// @Description: 如果它是目录或符号链接到目录，则返回路径的文件信息，否则返回错误。
// @param: pathRaw
// @return path
// @return fi
// @return err
func resolveDirInfo(pathRaw string) (path string, fi os.FileInfo, err error) {
	if fi, err = os.Lstat(pathRaw); err == nil {
		// 如果给定的路径是符号链接，则解析为实路径
		if isSymlinkFileInfo(&fi) {
			if path, err = filepath.EvalSymlinks(pathRaw); err == nil {
				// 为实际路径更新文件信息
				fi, err = os.Lstat(path)
			}
			if err != nil {
				path = ""
				return
			}
		} else {
			// 如果原始路径不是要解析的符号链接，只需清除路径
			path = filepath.Clean(pathRaw)
		}

		// 检查最终路径是否为目录
		if !isDirFileInfo(&fi) {
			err, path = errNotDirectory, ""
		}
	}
	return
}

// openFileInfo
// @Description: 如果它是常规文件，则返回路径的文件描述符和信息，否则返回错误。
// @param: path
// @return file
// @return fi
// @return err
func openFileInfo(path string) (file *os.File, fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); err == nil {
		if isFileInfo(&fi) {
			if file, err = os.Open(path); err == nil {
				return
			}
		} else {
			err = errNotRegularFile
		}
	}
	return
}

// compileRegexpList
// @Description: 将字符串模式列表编译为正则表达式模式。
// @param: strPats
// @return rePats
// @return err
func compileRegexpList(strPats []string) (rePats []*regexp.Regexp, err error) {
	var rp *regexp.Regexp
	rePats = make([]*regexp.Regexp, 0, len(strPats))
	for _, sp := range strPats {
		if rp, err = regexp.Compile(sp); err == nil {
			rePats = append(rePats, rp)
		} else {
			err = opError(opnList, sp, err)
			break
		}
	}
	return
}
