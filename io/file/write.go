package file

import "os"

// WriteFile
// @Description: 将数据写入文件
// @param: name 文件名
// @param: data 数据
// @param: perm 权限
// @return error
func WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}
