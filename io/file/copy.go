package file

import (
	"io"
	"io/ioutil"
	"math/bits"
	"os"
)

const (
	defaultDirectoryPermMode = os.FileMode(0755)
	defaultBufferSize        = 256 * 1024
	defaultNewFileFlag       = os.O_RDWR | os.O_CREATE | os.O_TRUNC
)

// CopyFile
// @Description: 将文件复制到目标文件或目录。符号链接将被跟踪。
//如果目标文件是一个存在的文件，则目标文件将被源文件覆盖
//如果目标文件是已存在的目录，则将源文件复制到相同文件名的目录中。
//如果目标不存在，但其父目录存在，则源文件将被复制到以目标名称命名的父目录中。
//如果有错误，它的类型将是*os. pathror。
// @param: src
// @param: dest
// @return err
func CopyFile(src, dest string) (err error) {
	if src, dest, err = refineOpPaths(opnCopy, src, dest, true); err == nil {
		err = bufferCopyFile(src, dest, defaultBufferSize)
	}
	return
}

// CopyDir
// @Description: 递归地将目录复制到目标目录。目录中的符号链接将被复制而不是被跟踪。
//如果目标是一个存在的文件，则会返回错误
//如果目标是一个已存在的目录，则将源目录复制到同名目录中。
//如果目标不存在，但其父目录存在，则源目录将被复制到具有目标名称的父目录中。
//如果发生错误，它会立即停止并返回，错误类型为*os.PathError。
// @param: src
// @param: dest
// @return err
func CopyDir(src, dest string) (err error) {
	if src, dest, err = refineOpPaths(opnCopy, src, dest, true); err == nil {
		err = copyDir(src, dest)
	}
	return
}

// CopySymlink
// @Description: 将符号链接复制到目标文件或目录。
// CopySymlink只复制链接中的内容，而不会尝试读取链接指向的文件。
//如果有错误，它的类型将是*os. pathror。
// @param: src
// @param: dest
// @return err
func CopySymlink(src, dest string) (err error) {
	if src, dest, err = refineOpPaths(opnCopy, src, dest, false); err == nil {
		err = copySymlink(src, dest)
	}
	return
}

// bufferCopyFile
// @Description: 从源文件读取内容并使用缓冲区写入目标文件。
//如果目标文件是一个存在的文件，则目标文件将被源文件覆盖
// @param: src
// @param: dest
// @param: bufferSize
// @return err
func bufferCopyFile(src, dest string, bufferSize int64) (err error) {
	var (
		srcFile, destFile *os.File
		srcInfo, destInfo os.FileInfo
	)

	// 检查源文件是否存在并打开以便读取
	if srcFile, srcInfo, err = openFileInfo(src); err == nil {
		defer srcFile.Close()
	} else {
		err = opError(opnCopy, src, err)
		return
	}

	// 检查源文件和目标文件是否相同
	if destInfo, err = os.Stat(dest); err == nil {
		if !isFileInfo(&destInfo) {
			err = opError(opnCopy, dest, errNotRegularFile)
		} else if os.SameFile(srcInfo, destInfo) {
			err = opError(opnCopy, dest, errSameFile)
		}
	} else if os.IsNotExist(err) {
		err = nil
	}
	if err != nil {
		return
	}

	// 如果源文件不够大，使用较小的缓冲区
	fileSize := srcInfo.Size()
	if bufferSize > fileSize {
		bufferSize = 1 << uint(bits.Len64(uint64(fileSize)))
	}

	if destFile, err = os.OpenFile(dest, defaultNewFileFlag, srcInfo.Mode()); err != nil {
		return
	}
	defer func() {
		if fe := destFile.Close(); fe != nil {
			err = fe
		}
		// 有错误删除目标文件
		if err != nil {
			_ = os.Remove(dest)
		}
	}()

	var nr, nw int
	buf := make([]byte, bufferSize)
	for {
		if nr, err = srcFile.Read(buf); err != nil || nr == 0 {
			if err == io.EOF && nr > 0 {
				err = opError(opnCopy, src, io.ErrUnexpectedEOF)
			}
			break
		}

		if nw, err = destFile.Write(buf[:nr]); err != nil {
			break
		} else if nw != nr {
			err = opError(opnCopy, dest, io.ErrShortWrite)
			break
		}
	}

	if err == io.EOF {
		err = nil
	}

	// err = destFile.Sync()
	return
}

// copySymlink
// @Description: 从源符号链接读取内容并写入目标符号链接。
// @param: src
// @param: dest
// @return err
func copySymlink(src, dest string) (err error) {
	var destInfo os.FileInfo
	if destInfo, err = os.Lstat(dest); err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
	} else {
		if isDirFileInfo(&destInfo) {
			// 避免覆盖目录
			err = opError(opnCopy, dest, errIsDirectory)
		} else {
			err = os.Remove(dest)
		}
	}
	if err != nil {
		return
	}

	var link string
	if link, err = os.Readlink(src); err != nil {
		err = opError(opnCopy, src, err)
	} else if err = os.Symlink(link, dest); err != nil {
		err = opError(opnCopy, dest, err)
	}
	return
}

// copyDir
// @Description: 递归地将源目录的所有条目复制到目标目录。
//如果目标目录是一个存在的文件，则会发生错误。
// @param: src
// @param: dest
// @return err
func copyDir(src, dest string) (err error) {
	var srcInfo, destInfo os.FileInfo

	// 检查source是否存在并且是一个目录
	if srcInfo, err = os.Stat(src); err == nil {
		if !isDirFileInfo(&srcInfo) {
			err = opError(opnCopy, src, errNotDirectory)
		}
	}
	if err != nil {
		return
	}

	// 检查目标是否不存在，或者不是文件或源本身
	if destInfo, err = os.Stat(dest); err == nil {
		if !isDirFileInfo(&destInfo) {
			err = opError(opnCopy, dest, errNotDirectory)
		} else if os.SameFile(srcInfo, destInfo) {
			err = opError(opnCopy, dest, errSameFile)
		}
	} else if os.IsNotExist(err) {
		err = nil
		if err = os.MkdirAll(dest, defaultDirectoryPermMode); err == nil {
			originMode := srcInfo.Mode()
			defer os.Chmod(dest, originMode)
		}
	}
	if err != nil {
		return
	}

	// 循环遍历源目录中的条目
	var entries []os.FileInfo
	if entries, err = ioutil.ReadDir(src); err != nil {
		return
	}

IterateEntry:
	for _, entry := range entries {
		srcPath, destPath := JoinPath(src, entry.Name()), JoinPath(dest, entry.Name())

		switch entry.Mode() & os.ModeType {
		case os.ModeDir:
			if err = copyDir(srcPath, destPath); err != nil {
				break IterateEntry
			}
		case os.ModeSymlink:
			if err = copySymlink(srcPath, destPath); err != nil {
				break IterateEntry
			}
		case 0:
			if err = bufferCopyFile(srcPath, destPath, defaultBufferSize); err != nil {
				break IterateEntry
			}
		}
	}

	return
}
