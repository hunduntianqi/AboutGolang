package No7_FileOperation

/*
	CSV 文件读写:
		encoding/csv包
		type Reader ==> 从csv编码的文件中读取记录
			type Reader struct {
				Comma            rune // 字段分隔符（NewReader将之设为','）
				Comment          rune // 一行开始位置的注释标识符
				FieldsPerRecord  int  // 每条记录期望的字段数
				LazyQuotes       bool // 允许懒引号
				TrailingComma    bool // 忽略，出于后端兼容性而保留
				TrimLeadingSpace bool // 去除前导的空白
				// 内含隐藏或非导出字段
			}
			func NewReader(r io.Reader) *Reader ==> 创建一个指定文件创建 *Reader 对象
			func (r *Reader) Read() (record []string, err error) ==> 从 r 读取一行数据, 返回值是字符串的切片, 每个字符串代表一个字段
			func (r *Reader) ReadAll() (records [][]string, err error) ==> 从 r 中读取剩余的所有数据, 返回值是二维字符串切片
		type Writer ==> 将数据写入一个csv编码的文件
			type Writer struct {
				Comma   rune // 字段分隔符(NewWriter将之设为',')
				UseCRLF bool // 如为真, 则换行时使用\r\n
				// 内含隐藏或非导出字段
			}
			func NewWriter(w io.Writer) *Writer ==> 根据指定文件创建 *Writer 对象
			func (w *Writer) Write(record []string) (err error) ==> 向指定文件中写入一行数据, 会自行添加必需的引号
			func (w *Writer) WriteAll(records [][]string) (err error) ==> 向指定文件中写入多行数据, 会自行添加必需的引号, 并在最后调用Flush方法清空缓存
			func (w *Writer) Flush() ==> 缓存中的数据写入底层的io.Writer
*/
