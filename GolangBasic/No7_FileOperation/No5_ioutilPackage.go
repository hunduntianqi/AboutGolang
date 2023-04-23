package No7_FileOperation

/*
	ioutil 工具包:
		实现了一些 I/O 操作的实用函数
			1. func ReadAll(r io.Reader) ([]byte, error): 从 r 读取数据直到 EOF 或遇到 error, 返回读取数据的字节切片和遇到的错误
			2. func ReadFile(filename string) ([]byte, error): 从指定路径的文件中读取所有数据并返回字节切片和错误描述
			3. func WriteFile(filename string, data []byte, perm fs.FileMode) error: 向指定文件中写入数据data, 如果文件不存在则按照
				给定的 perm 创建文件, 文件存在则清空文件原有数据
			4. func ReadDir(dirname string) ([]os.FileInfo, error): 返回指定的文件夹中所有文件 FileInfo 对象的切片
*/
