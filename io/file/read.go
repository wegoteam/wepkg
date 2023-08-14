package file

import "os"

// ReadFile
// @Description: 从文件中读取数据
// @param: filename 文件名
// @return []byte 读取到的数据
// @return error
func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

// ReadDir
// @Description: 从目录中读取数据
// @param: name
// @return []os.DirEntry
// @return error
func ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(name)
}

// Readlink
// @Description: 从符号链接中读取数据
// @param: name
// @return string
// @return error
func Readlink(name string) (string, error) {
	return os.Readlink(name)
}
