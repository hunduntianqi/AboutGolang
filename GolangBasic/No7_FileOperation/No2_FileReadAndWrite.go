package No7_FileOperation

/*
	读取文件数据:
		1. func (file *File) Read(b []byte) (n int, err Error): 读取文件数据到一个字节切片中
			返回值:
				n ==> 读取数据字节数
				err ==> Error对象
		2. func (file *File) ReadAt(b []byte, off int64) (n int, err Error): 从指定位置开始读取文件数据到一个字节切片中
			返回值:
				n ==> 读取数据字节数
				err ==> Error对象
	向文件写入数据:
		1. func (file *File) Write(b []byte) (n int, err Error): 将一个字节切片数据写入文件
			返回值:
				n ==> 写入文件数据的字节数
				err ==> Error对象
		2. func (file *File) WriteAt(b []byte, off int64) (n int, err Error): 从指定位置开始写入字节切片中的数据
			返回值:
				n ==> 写入文件数据的字节数
				err ==> Error对象
		3. func (file *File) WriteString(s string) (ret int, err Error): 讲一个字符串数据写入文件对象
			返回值:
				ret ==> 写入文件数据的字符数
				err ==> Error对象
*/
