package excel

//引用：https://github.com/qax-os/excelize
//官方文档：https://xuri.me/excelize/zh-hans/workbook.html#OpenFile

import excelUtil "github.com/xuri/excelize/v2"

const (
	DefaultSheetName = "Sheet1"
)

// Options
// @Description: excel工具类配置
// @return *excelUtil.Options
func Options() excelUtil.Options {
	return excelUtil.Options{}
}
