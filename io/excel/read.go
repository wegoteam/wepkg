package excel

import (
	"fmt"
	excelUtil "github.com/xuri/excelize/v2"
	"io"
)

// OpenFile
// @Description: 打开excel文件
// @param: filePath 文件路径
// @return *excelUtil.File
// @return error
func OpenFile(filePath string) (*excelUtil.File, error) {
	return excelUtil.OpenFile(filePath)
}

// OpenReader
// @Description: 打开excel文件
// @param: fileReader 文件流
// @return *excelUtil.File
// @return error
func OpenReader(fileReader io.Reader) (*excelUtil.File, error) {
	return excelUtil.OpenReader(fileReader)
}

// GetDatasByFilePath
// @Description: 读取excel文件数据
// @param: filePath 文件路径
// @return [][]interface{}
// @return error
func GetDatasByFilePath(filePath string) ([][]any, error) {
	file, err := OpenFile(filePath)
	if err != nil {
		fmt.Errorf("Open file error: %v \n", err)
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Errorf("Close file error: %v \n", err)
		}
	}()
	var datas = make([][]any, 0)
	rows, err := file.GetRows(DefaultSheetName)
	if err != nil {
		fmt.Errorf("Get Rows error: %v \n", err)
		return datas, err
	}
	for _, row := range rows {
		var rowDatas = make([]any, 0)
		for _, colCell := range row {
			rowDatas = append(rowDatas, colCell)
		}
		datas = append(datas, rowDatas)
	}
	return datas, err
}

// GetDatasByReader
// @Description: 读取excel文件数据
// @param: fileReader 文件流
// @return [][]any
// @return error
func GetDatasByReader(fileReader io.Reader) ([][]any, error) {
	file, err := OpenReader(fileReader)
	if err != nil {
		fmt.Errorf("Open file error: %v \n", err)
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Errorf("Close file error: %v \n", err)
		}
	}()
	var datas = make([][]any, 0)
	rows, err := file.GetRows(DefaultSheetName)
	if err != nil {
		fmt.Errorf("Get Rows error: %v \n", err)
		return datas, err
	}
	for _, row := range rows {
		var rowDatas = make([]any, 0)
		for _, colCell := range row {
			rowDatas = append(rowDatas, colCell)
		}
		datas = append(datas, rowDatas)
	}
	return datas, err
}
