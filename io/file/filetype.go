package file

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/wegoteam/wepkg/log"
	"path"
)

// GetFileType
// @Description: 获取文件类型
// @param: path
func GetFileType(filesPaths string) (contentType, ext, parent string) {
	mtype, err := mimetype.DetectFile(filesPaths)
	if err != nil {
		log.Errorf("File type err %v \n", err)
		return
	}
	contentType = mtype.String()
	ext = mtype.Extension()
	parent = mtype.String()
	return
}

// GetFileExt
// @Description: 获取文件的后缀
// @param: filesPaths
// @return fileType
func GetFileExt(filesPaths string) (fileType string) {
	fileType = path.Ext(filesPaths)
	return
}
