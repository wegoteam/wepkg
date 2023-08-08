package file

import (
	"os"
)

// Exist
// @Description: 检查给定路径是否存在。
//如果文件是符号链接，它将尝试跟随链接并检查源文件是否存在。
// @param: path
// @return bool
func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// NotExist
// @Description: 检查给定路径是否不存在。
//如果文件是符号链接，它将尝试跟随链接并检查源文件是否不存在。
// @param: path
// @return bool
func NotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// ExistFile
// @Description: 检查给定路径是否存在并且是文件。
//如果文件是符号链接，它将尝试跟随链接并检查源文件是否存在。
// @param: path
// @return bool
func ExistFile(path string) bool {
	return checkPathExist(path, os.Stat, isFileInfo)
}

// ExistDir
// @Description: 检查给定路径是否存在并且是目录。
//如果文件是符号链接，它将尝试跟随链接并检查源文件是否存在。
// @param: path
// @return bool
func ExistDir(path string) bool {
	return checkPathExist(path, os.Stat, isDirFileInfo)
}

// ExistSymlink
// @Description: 检查给定路径是否存在并且是符号链接。
//它只检查路径本身，并不尝试跟随链接。
// @param: path
// @return bool
func ExistSymlink(path string) bool {
	return checkPathExist(path, os.Lstat, isSymlinkFileInfo)
}

// checkPathExist
// @Description: 检查给定路径是否存在。
// @param: path
// @param: stat
// @param: check
// @return bool
func checkPathExist(path string, stat funcStatFileInfo, check funcCheckFileInfo) bool {
	fi, err := stat(path)
	return err == nil && check(&fi)
}
