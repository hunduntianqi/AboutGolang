package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

/*
	Excelize 单元格相关:
		func (f *File) SetCellValue(sheet, cell string, value interface{}) error ==> 根据给定的工作表名和单元格坐标设置单元格的值
		func (f *File) GetCellValue(sheet, cell string, opts ...Options) (string, error) ==> 根据给定的工作表和单元格坐标获取单元格的值, 返回值将转换为 string 类型
		func (f *File) GetCellType(sheet, cell string) (CellType, error) ==> 根据给定的工作表、单元格坐标获取指定单元格的数据类型
		func (f *File) GetCols(sheet string, opts ...Options) ([][]string, error) ==> 根据给定的工作表名按列获取该工作表上全部单元格的值, 以二维数组形式返回
		func (f *File) GetRows(sheet string, opts ...Options) ([][]string, error) ==> 根据给定的工作表名按行获取该工作表上全部单元格的值, 以二维数组形式返回
		func (f *File) MergeCell(sheet, hCell, vCell string) error ==> 根据给定的工作表名和单元格坐标区域合并单元格
		func (f *File) UnmergeCell(sheet string, hCell, vCell string) error ==> 根据给定的工作表名和单元格坐标区域取消合并单元格
		func (f *File) GetMergeCells(sheet string) ([]MergeCell, error) ==> 根据给定的工作表名获取全部合并单元格的坐标区域和值
		func (m *MergeCell) GetCellValue() string ==> 获取合并单元格的值
*/

func main() {
	// 打开已有工作表
	workBook, errOpen := excelize.OpenFile("No15_ExcelOperator/OperatiorData/codes/practice1.xlsx")
	if errOpen != nil {
		fmt.Println(errOpen)
		return
	}
	// 获取工作簿所有表格名称列表
	sheetAllName := workBook.GetSheetList()
	fmt.Println(sheetAllName)
	// 按行获取 "TsetData" 表所有数据
	dataList, errData := workBook.GetRows(sheetAllName[2])
	if errData != nil {
		fmt.Println("获取表格数据失败 ==> ", errData)
		return
	}
	// 遍历数据
	for _, data := range dataList {
		fmt.Println(data)
	}
	// 设置单元格的值, 新增一行数据
	//workBook.SetCellValue(sheetAllName[2], "A"+strconv.Itoa(len(dataList)+1), "S10000")
	//workBook.SetCellValue(sheetAllName[2], "B"+strconv.Itoa(len(dataList)+1), "郭鹏涛")
	//workBook.SetCellValue(sheetAllName[2], "C"+strconv.Itoa(len(dataList)+1), "6000")
	//workBook.SetCellValue(sheetAllName[2], "D"+strconv.Itoa(len(dataList)+1), "开发部")
	// 保存工作簿
	workBook.Save()
	// 关闭工作簿
	workBook.Close()
}
