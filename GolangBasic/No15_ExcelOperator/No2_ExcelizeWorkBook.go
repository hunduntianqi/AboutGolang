package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

/*
	Excelize 工作簿相关:
		Options ==> 定义了读写电子表格时的选项
			type Options struct {
				Password          string ==> 以明文形式指定打开和保存工作簿时所使用的密码，默认值为空
				RawCellValue      bool ==> 用以指定读取单元格值时是否获取原始值，默认值为 false(应用数字格式)
				UnzipSizeLimit    int64 ==> 用以指定打开电子表格文档时的解压缩大小限制, 默认大小限制为 16GB
				UnzipXMLSizeLimit int64 ==> 用以指定解压每个工作表以及共享字符表时的内存限制, 默认大小限制为 16MB
			}
		func NewFile() *File ==> 新建 Excel 工作薄, 默认包含一个名为 Sheet1 的工作表
		func OpenFile(filename string, opts ...Options) (*File, error) ==> 打开已有 Excel 工作簿
			例: 打开带有密码保护的工作簿
				f, err := excelize.OpenFile("Book1.xlsx", excelize.Options{Password: "password"})
				if err != nil {
					return
				}
		func (f *File) Save(opts ...Options) error ==> 保存工作簿
		func (f *File) SaveAs(name string, opts ...Options) error ==> 保存工作簿到指定路径(新建工作簿必须用此方法保存)
		func (f *File) Close() error ==> 关闭工作簿并清理打开文档时可能产生的系统磁盘缓存
		func (f *File) SetActiveSheet(index int) ==> 根据给定的索引值设置活动工作表, 索引的值应该大于等于 0 且小于工作簿所包含的累积工作表总数
		func (f *File) GetActiveSheetIndex() int ==> 获取活动工作表索引
		func (f *File) NewSheet(sheet string) (int, error) ==> 根据指定名称新建工作表
		func (f *File) DeleteSheet(sheet string) error ==> 删除指定名称的工作表, 当工作簿中仅包含一个工作表时, 调用此方法无效
		func (f *File) CopySheet(from, to int) error ==> 根据给定的 被复制工作表 与 目标工作表 索引 复制工作表, 目前支持仅包含
														 单元格值和公式的工作表间的复制
		func (f *File) SetSheetBackground(sheet, picture string) error ==> 根据给定的工作表名称和图片文件路径为指定的工作表设置平铺效果的背景图片
			支持的图片文件格式为：BMP、EMF、EMZ、GIF、JPEG、JPG、PNG、SVG、TIF、TIFF、WMF 和 WMZ
		func (f *File) SetSheetBackgroundFromBytes(sheet, extension string, picture []byte) error ==> 根据给定的工作表名称、图片格式扩展名
			和图片格式数据为指定的工作表设置平铺效果的背景图片, 支持的图片文件格式为：BMP、EMF、EMZ、GIF、JPEG、JPG、PNG、SVG、TIF、TIFF、WMF 和 WMZ
		func (f *File) SetSheetProps(sheet string, opts *SheetPropsOptions) error ==> 根据给定的工作表名称和属性参数设置工作表属性
		func (f *File) SetSheetVisible(sheet string, visible bool, veryHidden ...bool) error ==> 根据给定的工作表名称和可见性参数设置工作表的可见性,
			一个工作簿中至少包含一个可见工作表, 如果给定的工作表为默认工作表, 则对其可见性设置无效, 可选参数 veryHidden 仅在 visible 参数值为 false 时有效
		func (f *File) GetSheetVisible(sheet string) (bool, error) ==> 根据给定的工作表名称获取工作表可见性设置
		func (f *File) ProtectWorkbook(opts *WorkbookProtectionOptions) error ==> 使用密码保护工作簿的结构，以防止其他用户查看隐藏的工作表，添加、移动
			或隐藏工作表以及重命名工作表
			WorkbookProtectionOptions 定义了保护工作簿的设置选项
			type WorkbookProtectionOptions struct {
				AlgorithmName string ==> 指定加密算法, 支持指定哈希算法 XOR、MD4、MD5、SHA-1、SHA-256、SHA-384 或 SHA-512, 默认为  XOR 算法
				Password      string ==> 使用密码保护工作簿
				LockStructure bool ==> 是否锁定工作簿结构
				LockWindows   bool ==> 是否锁定工作簿窗口
			}
		func (f *File) UnprotectWorkbook(password ...string) error ==> 指定可选密码参数以通过密码验证来取消工作簿保护
*/

func main() {
	// 打开工作簿
	workBook, errOpenWorkBook := excelize.OpenFile("No15_ExcelOperator/TestCreateFile/newWorkBook.xlsx")
	if errOpenWorkBook != nil {
		fmt.Println("打开工作簿失败 ==> ", errOpenWorkBook)
		return
	}
	// 新建工作表
	workBook.NewSheet("新建工作表")
	// 设置第一张表格为默认工作表, index = 0
	workBook.SetActiveSheet(0)
	// 设置新建工作表背景图片
	workBook.SetSheetBackground("新建工作表", "E:\\Users\\Administrator\\IdeaProjects\\Pictures\\pickure\\1675088424171-234394110.jpg")
	// 设置保护工作簿
	err := workBook.ProtectWorkbook(&excelize.WorkbookProtectionOptions{
		AlgorithmName: "MD5",
		Password:      "2251789949gpt",
		LockStructure: true,
		LockWindows:   true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	// 保存工作簿
	errSave := workBook.Save()
	if err != nil {
		fmt.Println(errSave)
		return
	}
	// 关闭工作簿
	workBook.Close()
}
