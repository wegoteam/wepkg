package excel

import (
	"errors"
	"fmt"
	excelUtil "github.com/xuri/excelize/v2"
	"io"
)

// NewFile
// @Description: 创建一个新的excel文件
// @return *excelUtil.File
func NewFile() *excelUtil.File {
	return excelUtil.NewFile()
}

// SaveExcelPath
// @Description: 保存excel文件到指定路径
// @param: filePath 文件路径
// @param: datas 数据
// @return error
func SaveExcelPath(filePath string, datas [][]interface{}) error {
	if datas == nil || len(datas) == 0 {
		return fmt.Errorf("datas is empty \n")
		return errors.New("数据为空")
	}
	if filePath == "" {
		return fmt.Errorf("filePath is empty \n")
		return errors.New("文件路径为空")
	}
	file := NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Errorf("Close file error: %v \n", err)
		}
	}()
	for idx, row := range datas {
		cell, err := excelUtil.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Errorf("Coordinates Cell  error: %v \n", err)
			return err
		}
		file.SetSheetRow(DefaultSheetName, cell, &row)
	}
	file.Path = filePath
	file.SetActiveSheet(0)
	if err := file.Save(); err != nil {
		fmt.Errorf("Save file error: %v \n", err)
		return err
	}
	return nil
}

// SaveAsExcelPath
// @Description: 另存excel文件到指定路径
// @param: filePath
// @param: datas
// @return error
func SaveAsExcelPath(filePath string, datas [][]interface{}) error {
	if datas == nil || len(datas) == 0 {
		return fmt.Errorf("datas is empty \n")
		return errors.New("数据为空")
	}
	if filePath == "" {
		return fmt.Errorf("filePath is empty \n")
		return errors.New("文件路径为空")
	}
	file := NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Errorf("Close file error: %v \n", err)
		}
	}()
	for idx, row := range datas {
		cell, err := excelUtil.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Errorf("Coordinates Cell  error: %v \n", err)
			return err
		}
		file.SetSheetRow(DefaultSheetName, cell, &row)
	}
	file.SetActiveSheet(0)
	if err := file.SaveAs(filePath); err != nil {
		fmt.Errorf("Save file error: %v \n", err)
		return err
	}
	return nil
}

// SaveExcelWriter
// @Description: 保存excel文件到指定路径
// @param: fileWriter
// @param: datas
// @return error
func SaveExcelWriter(fileWriter io.Writer, datas [][]interface{}) error {
	if datas == nil || len(datas) == 0 {
		return fmt.Errorf("datas is empty \n")
		return errors.New("数据为空")
	}
	if fileWriter == nil {
		return fmt.Errorf("fileWriter is empty \n")
		return errors.New("文件流为空")
	}
	file := NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Errorf("Close file error: %v \n", err)
		}
	}()
	for idx, row := range datas {
		cell, err := excelUtil.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Errorf("Coordinates Cell  error: %v \n", err)
			return err
		}
		file.SetSheetRow(DefaultSheetName, cell, &row)
	}
	file.SetActiveSheet(0)
	if err := file.Write(fileWriter); err != nil {
		fmt.Errorf("Save file error: %v \n", err)
		return err
	}
	return nil
}

// SaveAsExcelWriter
// @Description: 另存excel文件到指定路径
// @param: fileWriter
// @param: datas
// @return error
func SaveAsExcelWriter(fileWriter io.Writer, datas [][]interface{}) error {
	if datas == nil || len(datas) == 0 {
		return fmt.Errorf("datas is empty \n")
		return errors.New("数据为空")
	}
	if fileWriter == nil {
		return fmt.Errorf("fileWriter is empty \n")
		return errors.New("文件流为空")
	}
	file := NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Errorf("Close file error: %v \n", err)
		}
	}()
	for idx, row := range datas {
		cell, err := excelUtil.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Errorf("Coordinates Cell  error: %v \n", err)
			return err
		}
		file.SetSheetRow(DefaultSheetName, cell, &row)
	}
	file.SetActiveSheet(0)
	if _, err := file.WriteTo(fileWriter); err != nil {
		fmt.Errorf("Save file error: %v \n", err)
		return err
	}
	return nil
}
