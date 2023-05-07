package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

/*
	Excelize 工作表相关:
		func (f *File) SetColVisible(sheet, col string, visible bool) error ==> 根据给定的工作表名称和列名称设置列可见性
			例: err := f.SetColVisible("Sheet1", "D", false) ==> 隐藏名为 Sheet1 工作表上的 D 列
				err := f.SetColVisible("Sheet1", "D:F", false) ==> 隐藏名称为 Sheet1 的工作表中的 D 至 F 列
		func (f *File) SetColWidth(sheet, startCol, endCol string, width float64) error ==> 根据给定的工作表名称、列范围和宽度值设置单个或多个列的宽度
		func (f *File) SetRowHeight(sheet string, row int, height float64) error ==> 根据给定的工作表名称、行号和高度值设置单行高度
		func (f *File) SetRowVisible(sheet string, row int, visible bool) error ==> 根据给定的工作表名称和行号设置行可见性
			例: err := f.SetRowVisible("Sheet1", 2, false) ==> 隐藏名为 Sheet1 工作表上第二行
		func (f *File) GetSheetName(index int) string ==> 根据给定的工作表索引获取工作表名称, 如果工作表不存在将返回空字符
		func (f *File) GetColVisible(sheet, column string) (bool, error) ==> 根据给定的工作表名称和列名获取工作表中指定列的可见性,
			可见返回值为 true, 否则为 false
		func (f *File) GetRowVisible(sheet string, row int) (bool, error) ==> 根据给定的工作表名称和行号获取工作表中指定行的可见性
		func (f *File) GetColWidth(sheet, col string) (float64, error) ==> 根据给定的工作表和列名获取工作表中指定列的宽度
		func (f *File) GetRowHeight(sheet string, row int) (float64, error) ==> 根据给定的工作表名称和行号获取工作表中指定行的高度
		func (f *File) GetSheetIndex(sheet string) (int, error) ==> 根据给定的工作表名称获取该工作表的索引，如果工作表不存在将返回 -1
		func (f *File) GetSheetMap() map[int]string ==> 获取工作簿中以 ID 和名称构成的全部工作表、图表工作表和对话工作表映射表
		func (f *File) GetSheetList() []string ==> 获取与工作簿内顺序相一致的, 包含工作表、图表工作表、对话工作表在内的工作表切片
		func (f *File) SetSheetName(source, target string) error ==> 根据给定的新旧工作表名称重命名工作表, 工作表名称最多允许使用 31 个字符
		func (f *File) InsertCols(sheet, col string, n int) error ==> 根据给定的工作表名称、列名称和要插入的列数, 在指定列前插入空白列
		func (f *File) InsertRows(sheet string, row, n int) error ==> 根据给定的工作表名称、行号和要插入的行数, 在指定行前插入空白行
		func (f *File) RemoveCol(sheet, col string) error ==> 根据给定的工作表名称和列名称删除指定列
		func (f *File) RemoveRow(sheet string, row int) error ==> 根据给定的工作表名称和行号删除指定行
		func (f *File) DuplicateRow(sheet string, row int) error ==> 根据给定的工作表名称和行号, 将该行数据剪切到下一行
		func (f *File) DuplicateRowTo(sheet string, row, row2 int) error ==> 根据给定的工作表名称和行号, 将该行数据剪切到对应行
		func (f *File) SetRowOutlineLevel(sheet string, row int, level uint8) error ==> 根据给定的工作表名称、行号和分级参数创建组
		func (f *File) SetColOutlineLevel(sheet, col string, level uint8) error ==> 根据给定的工作表名称、列名称和分级参数创建组
		func (f *File) GetRowOutlineLevel(sheet string, row int) (uint8, error) ==> 根据给定的工作表名称和行号获取分组级别
		func (f *File) GetColOutlineLevel(sheet, col string) (uint8, error) ==> 根据给定的工作表名称和列名称获取分组分级
		列迭代器:
			func (f *File) Cols(sheet string) (*Cols, error) ==> 根据给定的工作表名称获取该工作表的列迭代器, 使用列迭代器进行流式读取遍历单元格
			func (cols *Cols) Rows(opts ...Options) ([]string, error) ==> 返回当前列所有行的值
			func (cols *Cols) Next() bool ==> 下一列有值存在将返回 true
		行迭代器:
			func (f *File) Rows(sheet string) (*Rows, error) ==> 根据给定的工作表名称获取该工作表的行迭代器, 使用行迭代器进行流式读取遍历单元格
			func (rows *Rows) Columns(opts ...Options) ([]string, error) ==> 返回当前行中各列单元格的值
			func (rows *Rows) Next() bool ==> 下一行有值存在将返回 true
			func (rows *Rows) GetRowOpts() RowOpts ==> 返回当前行的行高、可见性和样式 ID 属性
			func (rows *Rows) Close() error ==> 关闭数据流并清理打开工作表时可能产生的系统磁盘缓存
		func (f *File) SetSheetCol(sheet, cell string, slice interface{}) error ==> 根据给定的工作表名称、起始坐标和 slice 类型引用按列赋值
		func (f *File) SetSheetRow(sheet, cell string, slice interface{}) error ==> 根据给定的工作表名称、起始坐标和 slice 类型引用按行赋值
		func (f *File) GetCols(sheet string, opts ...Options) ([][]string, error) ==> 根据给定的工作表名称, 获取所有非空列数据
		func (f *File) GetRows(sheet string, opts ...Options) ([][]string, error) ==> 根据给定的工作表名称, 获取所有非空行数据
		func (f *File) GetSheetDimension(sheet string) (string, error) ==> 根据给定的工作表名称获取指定工作表的已用区域
		func (f *File) SetSheetDimension(sheet string, rangeRef string) error ==> 根据给定的工作表名称和单元格坐标或单元格坐标区域设置或移除工作表的已用区域
			使用的单元格包括具有公式、文本内容和单元格格式的单元格, 当给定的单元格坐标区域为空字符时, 将移除工作表的已用区域
*/

func main() {
	// 打开已有工作簿
	workBook, errWork := excelize.OpenFile("No15_ExcelOperator/OperatiorData/codes/practice1.xlsx")
	if errWork != nil {
		fmt.Println("打开工作簿失败 ==> ", errWork)
		return
	}
	// 设置第二张工作表 A:D 列列宽为30
	workBook.SetColWidth(workBook.GetSheetName(1), "A", "D", 30)
	// 设置第二张工作表 1~22 行行高为20
	for i := 1; i <= 22; i++ {
		workBook.SetRowHeight(workBook.GetSheetName(1), i, 20)
	}
	// 获取工作簿中所有工作表 map 对象
	mapSheet := workBook.GetSheetMap()
	fmt.Println("工作簿中的表格状况为 ==> ", mapSheet)
	// 获取工作簿中所有工作表 名称 Slice
	sheetNameList := workBook.GetSheetList()
	fmt.Println("工作簿中的表格为 ==> ", sheetNameList)
	// 获取第二张表格全部非空行数据
	rowData, _ := workBook.GetRows(sheetNameList[1])
	for _, data := range rowData {
		fmt.Println(data)
	}
	// 获取第二张表格全部非空列数据
	colData, _ := workBook.GetCols(sheetNameList[1])
	for _, data := range colData {
		fmt.Println(data)
	}
	fmt.Println("第二张表格中最后一行数据行数为 ==> ", len(rowData))
	// 新建工作表, 并按行赋值将第二张工作表数据添加到新建工作表中
	workBook.NewSheet("TestData")
	for i := 0; i < len(rowData); i++ {
		// 按行赋值
		workBook.SetSheetRow("TestData", "A"+strconv.Itoa(i+1), &rowData[i])
	}
	// 设置第二张表格已用区域
	workBook.SetSheetDimension(sheetNameList[1], "A1:D21")
	// 获取第二张表格已用区域
	fmt.Println(workBook.GetSheetDimension(sheetNameList[1]))
	// 保存并关闭工作簿
	workBook.Save()
	workBook.Close()
}
