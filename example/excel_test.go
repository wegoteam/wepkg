package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/io/excel"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestCreateExcelFile(t *testing.T) {
	//https://github.com/qax-os/excelize
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	//_, err := f.NewSheet("Sheet2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	// Set value of a cell.
	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	//f.SetCellValue("Sheet1", "B2", 100)

	for idx, row := range [][]interface{}{
		{nil, "Apple", "Orange", "Pear"}, {"Small", 2, 3, 3},
		{"Normal", 5, 2, 4}, {"Large", 6, 7, 8},
	} {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetSheetRow("Sheet1", cell, &row)
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(0)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("./testdata/Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func TestReadExcel(t *testing.T) {
	f, err := excelize.OpenFile("./testdata/Book1.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func TestExcelFile(t *testing.T) {
	var datas = make([][]interface{}, 0)
	datas = append(datas, []interface{}{"姓名", "年龄", "性别"})
	datas = append(datas, []interface{}{"张三", 18, "男"})
	datas = append(datas, []interface{}{"李四", 20, "女"})
	datas = append(datas, []interface{}{"王五", 22, "男"})
	err := excel.SaveExcelPath("./testdata/Book2.xlsx", datas)
	if err != nil {
		fmt.Println(err)
	}

	err = excel.SaveExcelWriter(nil, datas)
	if err != nil {
		fmt.Println(err)
	}

}
