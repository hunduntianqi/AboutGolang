package No7_FileOperation

/*
	bufio 包:
		bufio包实现了带缓冲区的读写，是对文件读写的封装, 类似 java 的缓冲流
		bufio包读数据:
			创建对象:
				1. func NewReader(rd io.Reader) *Reader: 创建一个具有默认大小缓冲, 从r读取数据的 *Reader 对象
				2. func NewReaderSize(rd io.Reader, size int) *Reader: 创建一个具有指定大小缓冲, 从r读取数据的 *Reader对象
				例:
					file, err := os.open(path)
					reader := bufio.NewReader(file)
			常用方法:
				1. func (b *Reader) Reset(r io.Reader): 丢弃缓冲区数据, 清空缓冲区, 清除错误
				2. func (b *Reader) Buffered() int: 返回缓冲区现有的可读取数据字节数
				3. func (b *Reader) Read(p []byte) (n int, err error): 读取数据到字节切片 'p' 中, 返回读取数据字节数和错误对象
					当数据读取完毕后, err == io.EOF
				4. func (b *Reader) ReadByte() (c byte, err error): 一次读取一个字节数据并返回字节数据和错误对象, 当数据读取完毕后, err == io.EOF
				5. func (b *Reader) ReadRune() (r rune, size int, err error): 一次读取一个字符数据并返回字符数据, 字节长度和错误对象,
					当数据读取完毕后, err == io.EOF
				6. func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error): 一次读取一行数据, 返回字节切片
					isPerfix ==> 判断读取数据是否超过缓冲大小, 超过缓冲大小则返回true, 否则返回false
				7. func (b *Reader) ReadSlice(delim byte) (line []byte, err error): 读取数据直到第一次遇到 'delim' 字节, 返回包含 'delim' 的字节切片,
					如果读取到delim之前缓冲就被写满, 则读取失败并返回错误ErrBufferFull
				8. func (b *Reader) ReadBytes(delim byte) (line []byte, err error): 读取数据直到第一次遇到 'delim' 字节, 返回包含 'delim' 的字节切片,
					如果读取到delim之前遇到了错误, 会返回在错误之前读取的数据以及该错误, 当且仅当切片不以 delim 结尾时, 会返回一个非nil的错误
				9. func (b *Reader) ReadString(delim byte) (line string, err error): 读取数据直到第一次遇到 'delim' 字节, 返回包含 'delim' 的字符串,
					如果读取到delim之前遇到了错误, 会返回在错误之前读取的数据以及该错误, 当且仅当切片不以 delim 结尾时, 会返回一个非nil的错误
		bufio包写数据:
			创建对象:
				1. func NewWriter(w io.Writer) *Writer: 创建一个具有默认大小缓冲, 向w写入数据的 *Writer 对象
				2. func NewWriterSize(w io.Writer, size int) *Writer: 创建一个具有指定大小缓冲, 向w写入数据的 *Writer 对象
			常用方法:
				1. func (b *Writer) Reset(w io.Writer): 丢弃缓冲区数据, 清空缓冲区, 清除错误
				2. func (b *Writer) Buffered() int: 获取当前缓冲区等待写入数据的字节数
				3. func (b *Writer) Available() int: 获取当前缓冲区中剩余可写入数据的字节数
				4. func (b *Writer) Write(p []byte) (nn int, err error): 将字节切片 'p'数据写入文件, 并返回写入成功字节数和错误说明
				5. func (b *Writer) WriteString(s string) (int, error): 将字符串 's' 数据写入文件, 并返回写入成功字符数和错误说明
				6. func (b *Writer) WriteByte(c byte) error: 将字节 'c' 写入文件
				7. func (b *Writer) WriteRune(r rune) (size int, err error): 将字符 'r' 写入文件, 并返回写入成功字节数和错误说明
				8. func (b *Writer) Flush() error: 刷新缓冲区, 立刻将缓冲区数据全部写入文件
		bufio.Scanner类型:
			提供了方便的读取数据的接口
			创建Scanner类型对象 ==> scannerName := bufio.NewScanner(io.Reader)
			设置分割函数, 根据分割函数指定字符分割字符串 ==> func (s *Scanner) Split(split SplitFunc)
				常用分割函数:
					1. func ScanLines: 将换行符作为分割标记
					2. func ScanWords: 将空白字符作为分割标记
					3. func ScanBytes: 将每个字节作为分割标记
					4. func ScanRunes: 将utf-8编码的每个Unicode字符作为分割标记
				注意: scanner.Split(split SplitFunc)必须在scanner.Scan()之前调用
			获取当前位置的分割标记并移动到下一个分割标记的位置 ==> func (s *Scanner) Scan() bool
				正常返回true,当到达输入流结尾或者遇到错误时会返回false
			读取根据分割标记分割开的字节切片数据 ==> func (s *Scanner) Bytes() []byte
			读取根据分割标记分割开的字符串数据 ==> func (s *Scanner) Text() string
			Scanner使用步骤:
				1. 包装文件对象创建Scanner对象
					scannerName := bufio.NewScanner(io.Reader)
				2. 调用Split函数, 获取分割字符
					scannerName.Split(SplitFunc) ==> 该方法直接调用即可, 无返回值
				3. scannerName.Scan() 结合 for语句读取数据
					for scannerName.Scan() {
						4. 调用Text() / Bytes()方法读取数据
							scannerName.Text() ==> 返回数据字符串形式
							scannerName.Bytes() ==> 返回数据字节切片形式
					}
*/
