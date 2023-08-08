package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/io/file"
	"sort"
	"testing"
)

func TestFile(t *testing.T) {
	contentType, ext, parent := file.GetFileType("./testdata/a.txt")
	fmt.Printf("contentType=%v, ext=%v, parent=%v \n", contentType, ext, parent)
	fileType := file.GetFileExt("./testdata/a.txt")
	fmt.Printf("fileType=%v \n", fileType)
	isOn32bitArch := file.IsOn32bitArch()
	fmt.Printf("是否32位系统架构：%v \n", isOn32bitArch)
	isOn64bitArch := file.IsOn64bitArch()
	fmt.Printf("是否64位系统架构：%v \n", isOn64bitArch)
	isOnLinux := file.IsOnLinux()
	fmt.Printf("是否Linux系统：%v \n", isOnLinux)
	isOnMacOS := file.IsOnMacOS()
	fmt.Printf("是否MacOS系统：%v \n", isOnMacOS)
	isOnWindows := file.IsOnWindows()
	fmt.Printf("是否Windows系统：%v \n", isOnWindows)
	//file.ChangeExeDir()
	existDir := file.ExistDir("./testdata")
	fmt.Printf("是否存在目录：%v \n", existDir)
	existFile := file.ExistFile("./testdata/a.txt")
	fmt.Printf("是否存在文件：%v \n", existFile)
	existSymlink := file.ExistSymlink("./testdata/a.txt")
	fmt.Printf("是否存在软连接：%v \n", existSymlink)
	isDirEmpty, _ := file.IsDirEmpty("./testdata")
	fmt.Printf("目录是否为空：%v \n", isDirEmpty)
	isFileEmpty, _ := file.IsFileEmpty("./testdata/a.txt")
	fmt.Printf("文件是否为空：%v \n", isFileEmpty)
	size, _ := file.GetDirSize("./testdata")
	fmt.Printf("目录大小：%v \n", size)
	fileSize, _ := file.GetFileSize("./testdata/a.txt")
	fmt.Printf("文件大小：%v \n", fileSize)
	symlinkSize, _ := file.GetSymlinkSize("./testdata/a.txt")
	fmt.Printf("软连接大小：%v \n", symlinkSize)
	sameDirEntries, _ := file.SameDirEntries("./testdata", "./testdata")
	fmt.Printf("目录是否相同：%v \n", sameDirEntries)
	sameFileContent, _ := file.SameFileContent("./testdata/a.txt", "./testdata/a.txt")
	fmt.Printf("文件是否相同：%v \n", sameFileContent)
	sameSymlinkContent, _ := file.SameSymlinkContent("./testdata/a.txt", "./testdata/a.txt")
	fmt.Printf("软连接是否相同：%v \n", sameSymlinkContent)
	listDir, _ := file.ListDir("./testdata")
	fmt.Printf("目录列表：%v \n", listDir)
	listFile, _ := file.ListFile("./testdata")
	fmt.Printf("文件列表：%v \n", listFile)
	listSymlink, _ := file.ListSymlink("./testdata")
	fmt.Printf("软连接列表：%v \n", listSymlink)
	err := file.CopyDir("./testdata", "./testdata2")
	fmt.Printf("复制目录错误：%v \n", err)
	err = file.CopyFile("./testdata/a.txt", "./testdata2/a.txt")
	fmt.Printf("复制文件错误：%v \n", err)
	err = file.CopySymlink("./testdata/a.txt", "./testdata2/a.txt")
	fmt.Printf("复制软连接错误：%v \n", err)
	err = file.MoveDir("./testdata", "./testdata2")
	fmt.Printf("移动目录错误：%v \n", err)
	err = file.MoveFile("./testdata/a.txt", "./testdata2/a.txt")
	fmt.Printf("移动文件错误：%v \n", err)
	err = file.MoveSymlink("./testdata/a.txt", "./testdata2/a.txt")
	fmt.Printf("移动软连接错误：%v \n", err)
	listMatch, _ := file.ListMatch("./testdata", file.ListIncludeAll, "*.txt")
	fmt.Printf("匹配列表：%v \n", listMatch)
	joinPath := file.JoinPath("./testdata", "a.txt")
	fmt.Printf("拼接路径：%v \n", joinPath)
	exist := file.Exist("./testdata/a.txt")
	fmt.Printf("是否存在：%v \n", exist)
	notExist := file.NotExist("./testdata/a.txt")
	fmt.Printf("是否不存在：%v \n", notExist)
	err = file.MakeDir("./testdata2")
	fmt.Printf("创建目录错误：%v \n", err)
	//根据名称排序
	sort.Stable(file.SortListByName(listFile))
	//根据大小排序
	sort.Stable(file.SortListBySize(listFile))
	//根据修改时间排序
	sort.Stable(file.SortListByModTime(listFile))
}
