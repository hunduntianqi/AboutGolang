package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

/*
	Excelize:
		Go 语言编写的用于操作 Office Excel 文档基础库, 可以使用它来读取, 写入由 Microsoft Excel™ 2007 及以上版本创建的电子表格文档;
		支持 XLAM / XLSM / XLSX / XLTM / XLTX 等多种文档格式; 高度兼容带有样式、图片(表)、透视表、切片器等复杂组件的文档, 并提供流式
		读写 API, 用于处理包含大规模数据的工作簿; 本类库要求使用的 Go 语言为 1.16 或更高版本
	安装第三方包:
		go get github.com/xuri/excelize
		或 go get github.com/xuri/excelize/v2
		更新命令 ==> go get -u github.com/xuri/excelize/v2
	创建一个新的Excel工作簿 ==> workBook := excelize.NewFile()
	打开一个已有的工作簿 ==> workBook, err := excelize.OpenFile(filePath)
	保存工作簿 ==> workBook.SaveAs(filePath)
	创建一个工作表 ==> sheetName, err := workBook.NewSheet("SheetName")
	删除一个工作表 ==> err := workBook.DeleteSheet("SheetName")
	设置单元格的值 ==> workBook.SetCellValue("SheetName", "CellName", value)
*/

func main() {
	// 调用函数, 创建工作簿并保存
	//creatWorkBook()
	// 调用函数, 打开已有工作簿
	openWorkBook()
}

// 定义一个函数, 创建一个工作簿并写入数据
func creatWorkBook() {
	// 新建一个工作簿对象
	workBook := excelize.NewFile()
	// 新建工作表
	sheetName, errSheet := workBook.NewSheet("新建工作表格")
	if errSheet != nil {
		fmt.Println("新建工作表失败 ==> ", errSheet)
		return
	}
	fmt.Println(sheetName)
	// 向指定单元格写入数据
	workBook.SetCellValue("Sheet1", "A1", "一条测试数据")
	workBook.SetCellValue("Sheet1", "A2", "一条测试数据")
	workBook.SetCellValue("Sheet1", "A3", "一条测试数据")
	workBook.SetCellValue("Sheet1", "A4", "一条测试数据")
	workBook.SetCellValue("新建工作表格", "A4", "一条测试数据")
	// 保存工作簿到指定路径
	workBook.SaveAs("No15_ExcelOperator/TestCreateFile/newWorkBook.xlsx")
}

// 定义函数, 打开已有工作簿
func openWorkBook() {
	// 打开工作簿
	workBook, _ := excelize.OpenFile("No15_ExcelOperator/TestCreateFile/newWorkBook.xlsx")
	// 删除工作表
	workBook.DeleteSheet("新建工作表格")
	// 保存工作簿
	workBook.SaveAs("No15_ExcelOperator/TestCreateFile/newWorkBook.xlsx")
}
